package main

import (
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/debug"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/snapshots"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authcli "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"
	tmdb "github.com/tendermint/tm-db"

	petrichor "github.com/redactedfury/petri/app"
)

func NewRootCmd() (*cobra.Command, petri.EncodingConfig) {
	var (
		config  = petri.MakeEncodingConfig()
		context = client.Context{}.
			WithCodec(config.Marshaler).
			WithInterfaceRegistry(config.InterfaceRegistry).
			WithTxConfig(config.TxConfig).
			WithLegacyAmino(config.Amino).
			WithInput(os.Stdin).
			WithAccountRetriever(authtypes.AccountRetriever{}).
			WithBroadcastMode(flags.BroadcastBlock).
			WithHomeDir(petri.DefaultNodeHome)
	)

	cobra.EnableCommandSorting = false
	root := &cobra.Command{
		Use:   "petri",
		Short: "Petrichor - Decentralised Synthetic Asset Management",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			if err := client.SetCmdClientContextHandler(context, cmd); err != nil {
				return err
			}

			return server.InterceptConfigsPreRunHandler(cmd, "", nil)
		},
	}

	initRootCmd(root, config)
	return root, config
}

func initRootCmd(rootCmd *cobra.Command, encoding petri.EncodingConfig) {
	rootCmd.AddCommand(
		genutilcli.InitCmd(petri.ModuleBasics, petri.DefaultNodeHome),
		genutilcli.CollectGenTxsCmd(banktypes.GenesisBalancesIterator{}, petri.DefaultNodeHome),
		genutilcli.GenTxCmd(petri.ModuleBasics, encoding.TxConfig, banktypes.GenesisBalancesIterator{}, petri.DefaultNodeHome),
		genutilcli.ValidateGenesisCmd(petri.ModuleBasics),
		AddGenesisAccountCmd(petri.DefaultNodeHome),
		AddGenesisWasmMsgCmd(petri.DefaultNodeHome),
		tmcli.NewCompletionCmd(rootCmd, true),
		testnetCmd(petri.ModuleBasics, banktypes.GenesisBalancesIterator{}),
		debug.Cmd(),
	)

	server.AddCommands(rootCmd, petri.DefaultNodeHome, appCreatorFunc, appExportFunc, addModuleInitFlags)
	rootCmd.AddCommand(
		rpc.StatusCommand(),
		queryCommand(),
		txCommand(),
		keys.Commands(petri.DefaultNodeHome),
	)
}

func addModuleInitFlags(cmd *cobra.Command) {
	crisis.AddModuleInitFlags(cmd)
}

func queryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "query",
		Aliases:                    []string{"q"},
		Short:                      "Querying subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcli.GetAccountCmd(),
		rpc.ValidatorCommand(),
		rpc.BlockCommand(),
		authcli.QueryTxsByEventsCmd(),
		authcli.QueryTxCmd(),
	)

	petri.ModuleBasics.AddQueryCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

func txCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "tx",
		Short:                      "Transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcli.GetSignCommand(),
		authcli.GetSignBatchCommand(),
		authcli.GetMultiSignCommand(),
		authcli.GetMultiSignBatchCmd(),
		authcli.GetValidateSignaturesCommand(),
		authcli.GetBroadcastCommand(),
		authcli.GetEncodeCommand(),
		authcli.GetDecodeCommand(),
	)

	petri.ModuleBasics.AddTxCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

func appCreatorFunc(logger log.Logger, db tmdb.DB, tracer io.Writer, options servertypes.AppOptions) servertypes.Application {
	var cache sdk.MultiStorePersistentCache
	if cast.ToBool(options.Get(server.FlagInterBlockCache)) {
		cache = store.NewCommitKVStoreCacheManager()
	}

	skipUpgradeHeights := make(map[int64]bool)
	for _, height := range cast.ToIntSlice(options.Get(server.FlagUnsafeSkipUpgrades)) {
		skipUpgradeHeights[int64(height)] = true
	}

	pruningOptions, err := server.GetPruningOptionsFromFlags(options)
	if err != nil {
		panic(err)
	}

	snapshotDir := filepath.Join(cast.ToString(options.Get(flags.FlagHome)), "data", "snapshots")
	snapshotDB, err := sdk.NewLevelDB("metadata", snapshotDir)
	if err != nil {
		panic(err)
	}
	snapshotStore, err := snapshots.NewStore(snapshotDB, snapshotDir)
	if err != nil {
		panic(err)
	}
	var wasmOpts []wasm.Option
	if cast.ToBool(options.Get("telemetry.enabled")) {
		wasmOpts = append(wasmOpts, wasmkeeper.WithVMCacheMetrics(prometheus.DefaultRegisterer))
	}
	return petri.New(
		logger, db, tracer, true, skipUpgradeHeights,
		cast.ToString(options.Get(flags.FlagHome)),
		cast.ToUint(options.Get(server.FlagInvCheckPeriod)),
		petri.MakeEncodingConfig(),
		options,
		petri.GetWasmEnabledProposals(),
		wasmOpts,
		baseapp.SetPruning(pruningOptions),
		baseapp.SetMinGasPrices(cast.ToString(options.Get(server.FlagMinGasPrices))),
		baseapp.SetHaltHeight(cast.ToUint64(options.Get(server.FlagHaltHeight))),
		baseapp.SetHaltTime(cast.ToUint64(options.Get(server.FlagHaltTime))),
		baseapp.SetMinRetainBlocks(cast.ToUint64(options.Get(server.FlagMinRetainBlocks))),
		baseapp.SetInterBlockCache(cache),
		baseapp.SetTrace(cast.ToBool(options.Get(server.FlagTrace))),
		baseapp.SetIndexEvents(cast.ToStringSlice(options.Get(server.FlagIndexEvents))),
		baseapp.SetSnapshotStore(snapshotStore),
		baseapp.SetSnapshotInterval(cast.ToUint64(options.Get(server.FlagStateSyncSnapshotInterval))),
		baseapp.SetSnapshotKeepRecent(cast.ToUint32(options.Get(server.FlagStateSyncSnapshotKeepRecent))),
	)
}

func appExportFunc(logger log.Logger, db tmdb.DB, tracer io.Writer, height int64,
	forZeroHeight bool, jailAllowedAddrs []string, options servertypes.AppOptions,
) (servertypes.ExportedApp, error) {
	config := petri.MakeEncodingConfig()
	config.Marshaler = codec.NewProtoCodec(config.InterfaceRegistry)
	homePath, ok := options.Get(flags.FlagHome).(string)
	if !ok || homePath == "" {
		return servertypes.ExportedApp{}, errors.New("application home is not set")
	}
	var emptyWasmOpts []wasm.Option
	var app *petri.App
	if height != -1 {
		app = petri.New(logger, db, tracer, false, map[int64]bool{}, homePath, cast.ToUint(options.Get(server.FlagInvCheckPeriod)), config, options, petri.GetWasmEnabledProposals(), emptyWasmOpts)

		if err := app.LoadHeight(height); err != nil {
			return servertypes.ExportedApp{}, err
		}
	} else {
		app = petri.New(logger, db, tracer, true, map[int64]bool{}, homePath, cast.ToUint(options.Get(server.FlagInvCheckPeriod)), config, options, petri.GetWasmEnabledProposals(), emptyWasmOpts)
	}

	return app.ExportAppStateAndValidators(forZeroHeight, jailAllowedAddrs)
}
