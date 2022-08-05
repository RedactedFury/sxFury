package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	esmtypes "github.com/comdex-official/comdex/x/esm/types"
	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/comdex-official/comdex/x/locker/types"
)

//get locker lookup table.

func (k Keeper) SetLockerProductAssetMapping(ctx sdk.Context, lockerProductMapping types.LockerProductAssetMapping) {
	var (
		store = k.Store(ctx)
		key   = types.LockerProductAssetMappingKey(lockerProductMapping.AppId)
		value = k.cdc.MustMarshal(&lockerProductMapping)
	)

	store.Set(key, value)
}

func (k Keeper) SetLockerTotalRewardsByAssetAppWise(ctx sdk.Context, lockerRewardsMapping types.LockerTotalRewardsByAssetAppWise) error {
	var (
		store = k.Store(ctx)
		key   = types.LockerTotalRewardsByAssetAppWiseKey(lockerRewardsMapping.AppId, lockerRewardsMapping.AssetId)
		value = k.cdc.MustMarshal(&lockerRewardsMapping)
	)

	store.Set(key, value)
	return nil
}

func (k Keeper) GetLockerTotalRewardsByAssetAppWise(ctx sdk.Context, appID, assetID uint64) (lockerRewardsMapping types.LockerTotalRewardsByAssetAppWise, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.LockerTotalRewardsByAssetAppWiseKey(appID, assetID)
		value = store.Get(key)
	)

	if value == nil {
		return lockerRewardsMapping, false
	}

	k.cdc.MustUnmarshal(value, &lockerRewardsMapping)
	return lockerRewardsMapping, true
}

func (k Keeper) GetAllLockerTotalRewardsByAssetAppWise(ctx sdk.Context) (lockerTotalRewardsByAssetAppWise []types.LockerTotalRewardsByAssetAppWise) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.LockerTotalRewardsByAssetAppWiseKeyPrefix)
	)

	defer func(iter sdk.Iterator) {
		err := iter.Close()
		if err != nil {
			return
		}
	}(iter)

	for ; iter.Valid(); iter.Next() {
		var lock types.LockerTotalRewardsByAssetAppWise
		k.cdc.MustUnmarshal(iter.Value(), &lock)
		lockerTotalRewardsByAssetAppWise = append(lockerTotalRewardsByAssetAppWise, lock)
	}
	return lockerTotalRewardsByAssetAppWise
}

func (k Keeper) GetLockerProductAssetMapping(ctx sdk.Context, appMappingID uint64) (lockerProductMapping types.LockerProductAssetMapping, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.LockerProductAssetMappingKey(appMappingID)
		value = store.Get(key)
	)

	if value == nil {
		return lockerProductMapping, false
	}

	k.cdc.MustUnmarshal(value, &lockerProductMapping)
	return lockerProductMapping, true
}

func (k Keeper) GetAllLockerProductAssetMapping(ctx sdk.Context) (lockerProductAssetMapping []types.LockerProductAssetMapping) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.LockerProductAssetMappingKeyPrefix)
	)

	defer func(iter sdk.Iterator) {
		err := iter.Close()
		if err != nil {
			return
		}
	}(iter)

	for ; iter.Valid(); iter.Next() {
		var lock types.LockerProductAssetMapping
		k.cdc.MustUnmarshal(iter.Value(), &lock)
		lockerProductAssetMapping = append(lockerProductAssetMapping, lock)
	}
	return lockerProductAssetMapping
}

func (k Keeper) SetLockerLookupTable(ctx sdk.Context, lockerLookupData types.LockerLookupTable) {
	var (
		store = k.Store(ctx)
		key   = types.LockerLookupTableKey(lockerLookupData.AppId)
		value = k.cdc.MustMarshal(&lockerLookupData)
	)

	store.Set(key, value)
}

func (k Keeper) GetLockerLookupTable(ctx sdk.Context, appMappingID uint64) (lockerLookupData types.LockerLookupTable, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.LockerLookupTableKey(appMappingID)
		value = store.Get(key)
	)

	if value == nil {
		return lockerLookupData, false
	}

	k.cdc.MustUnmarshal(value, &lockerLookupData)
	return lockerLookupData, true
}

func (k Keeper) GetAllLockerLookupTable(ctx sdk.Context) (lockerLookupTable []types.LockerLookupTable) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.LockerLookupTableKeyPrefix)
	)

	defer func(iter sdk.Iterator) {
		err := iter.Close()
		if err != nil {
			return
		}
	}(iter)

	for ; iter.Valid(); iter.Next() {
		var lock types.LockerLookupTable
		k.cdc.MustUnmarshal(iter.Value(), &lock)
		lockerLookupTable = append(lockerLookupTable, lock)
	}
	return lockerLookupTable
}

func (k Keeper) CheckLockerProductAssetMapping(ctx sdk.Context, assetID uint64, lockerProductMapping types.LockerProductAssetMapping) (found bool) {
	for _, id := range lockerProductMapping.AssetIds {
		if id == assetID {
			return true
		}
		continue
	}
	return false
}

// UpdateTokenLockerMapping For updating token locker mapping in lookup table.
func (k Keeper) UpdateTokenLockerMapping(ctx sdk.Context, lockerLookupData types.LockerLookupTable, counter uint64, userLockerData types.Locker) {
	for _, lockerData := range lockerLookupData.Lockers {
		if lockerData.AssetId == userLockerData.AssetDepositId {
			lockerData.DepositedAmount = lockerData.DepositedAmount.Add(userLockerData.NetBalance)
			lockerData.LockerIds = append(lockerData.LockerIds, userLockerData.LockerId)
		}
	}
	lockerLookupData.Counter = counter
	k.SetLockerLookupTable(ctx, lockerLookupData)
}

// UpdateAmountLockerMapping For updating token locker mapping in lookup table.
func (k Keeper) UpdateAmountLockerMapping(ctx sdk.Context, lockerLookupData types.LockerLookupTable, assetID uint64, amount sdk.Int, changeType bool) { //if Change type true = Add to deposits
	//If change type false = Subtract from the deposits

	for _, lockerData := range lockerLookupData.Lockers {
		if lockerData.AssetId == assetID {
			if changeType {
				lockerData.DepositedAmount = lockerData.DepositedAmount.Add(amount)
			} else {
				lockerData.DepositedAmount = lockerData.DepositedAmount.Sub(amount)
			}
		}
	}
	k.SetLockerLookupTable(ctx, lockerLookupData)
}

// SetUserLockerAssetMapping User Locker Functions.
func (k Keeper) SetUserLockerAssetMapping(ctx sdk.Context, userLockerAssetData types.UserLockerAssetMapping) {
	var (
		store = k.Store(ctx)
		key   = types.UserLockerAssetMappingKey(userLockerAssetData.Owner)
		value = k.cdc.MustMarshal(&userLockerAssetData)
	)

	store.Set(key, value)
}

func (k Keeper) GetUserLockerAssetMapping(ctx sdk.Context, address string) (userLockerAssetData types.UserLockerAssetMapping, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.UserLockerAssetMappingKey(address)
		value = store.Get(key)
	)

	if value == nil {
		return userLockerAssetData, false
	}

	k.cdc.MustUnmarshal(value, &userLockerAssetData)
	return userLockerAssetData, true
}

func (k Keeper) GetAllUserLockerAssetMapping(ctx sdk.Context) (userLockerAssetMapping []types.UserLockerAssetMapping) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.UserLockerAssetMappingKeyPrefix)
	)

	defer func(iter sdk.Iterator) {
		err := iter.Close()
		if err != nil {
			return
		}
	}(iter)

	for ; iter.Valid(); iter.Next() {
		var lock types.UserLockerAssetMapping
		k.cdc.MustUnmarshal(iter.Value(), &lock)
		userLockerAssetMapping = append(userLockerAssetMapping, lock)
	}
	return userLockerAssetMapping
}

// CheckUserAppToAssetMapping Checking if for a certain user for the app type , whether there exists a certain asset or not and if it contains a locker id or not.
func (k Keeper) CheckUserAppToAssetMapping(ctx sdk.Context, userLockerAssetData types.UserLockerAssetMapping, assetID uint64, appID uint64) (lockerID string, found bool) {
	for _, lockerAppMapping := range userLockerAssetData.LockerAppMapping {
		if lockerAppMapping.AppId == appID {
			for _, assetToLockerIDMapping := range lockerAppMapping.UserAssetLocker {
				if assetToLockerIDMapping.AssetId == assetID && len(assetToLockerIDMapping.LockerId) > 0 {
					lockerID = assetToLockerIDMapping.LockerId
					return lockerID, true
				}
			}
		}
	}
	return lockerID, false
}

func (k Keeper) CheckUserToAppMapping(ctx sdk.Context, userLockerAssetData types.UserLockerAssetMapping, appID uint64) (found bool) {
	for _, lockerAppMapping := range userLockerAssetData.LockerAppMapping {
		if lockerAppMapping.AppId == appID {
			return true
		}
	}
	return false
}

func (k Keeper) SetLocker(ctx sdk.Context, locker types.Locker) {
	var (
		store = k.Store(ctx)
		key   = types.LockerKey(locker.LockerId)
		value = k.cdc.MustMarshal(&locker)
	)

	store.Set(key, value)
}

func (k Keeper) GetLocker(ctx sdk.Context, lockerID string) (locker types.Locker, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.LockerKey(lockerID)
		value = store.Get(key)
	)

	if value == nil {
		return locker, false
	}

	k.cdc.MustUnmarshal(value, &locker)
	return locker, true
}

func (k Keeper) GetLockers(ctx sdk.Context) (locker []types.Locker) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.LockerKeyPrefix)
	)

	defer func(iter sdk.Iterator) {
		err := iter.Close()
		if err != nil {
			return
		}
	}(iter)

	for ; iter.Valid(); iter.Next() {
		var lock types.Locker
		k.cdc.MustUnmarshal(iter.Value(), &lock)
		locker = append(locker, lock)
	}
	return locker
}

func (k Keeper) UpdateLocker(ctx sdk.Context, locker types.Locker) {
	var (
		store = k.Store(ctx)
		key   = types.LockerKey(locker.LockerId)
		value = k.cdc.MustMarshal(&locker)
	)

	store.Set(key, value)
}

//Target
//user sends create request
//it comes to the function and check if user data exists or not. if not create locker
//if user data exists- check app mapping , from app mapping check asset id . if it does then fail tx.
// else user locker id  exists use that to create this struct and set it.

func QueryState(addr, denom, blockHeight, target string) (*sdk.Coin, error) {
	myAddress, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return nil, err
	}

	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		target,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}
	defer func(grpcConn *grpc.ClientConn) {
		err := grpcConn.Close()
		if err != nil {
			return
		}
	}(grpcConn)

	bankClient := banktypes.NewQueryClient(grpcConn)
	bankRes, err := bankClient.Balance(
		context.Background(),
		&banktypes.QueryBalanceRequest{Address: myAddress.String(), Denom: denom},
	)
	if err != nil {
		return nil, err
	}

	var header metadata.MD
	bankRes, err = bankClient.Balance(
		metadata.AppendToOutgoingContext(context.Background(), grpctypes.GRPCBlockHeightHeader, blockHeight), // Add metadata to request
		&banktypes.QueryBalanceRequest{Address: myAddress.String(), Denom: denom},
		grpc.Header(&header),
	)

	if err != nil {
		return nil, err
	}

	return bankRes.GetBalance(), nil
}

func (k Keeper) WasmAddWhiteListedAssetQuery(ctx sdk.Context, appMappingID, AssetID uint64) (bool, string) {
	_, found := k.GetApp(ctx, appMappingID)
	if !found {
		return false, types.ErrorAppMappingDoesNotExist.Error()
	}
	_, found = k.GetAsset(ctx, AssetID)
	if !found {
		return false, types.ErrorAssetDoesNotExist.Error()
	}
	lockerProductAssetMapping, found := k.GetLockerProductAssetMapping(ctx, appMappingID)

	if found {
		found := k.CheckLockerProductAssetMapping(ctx, AssetID, lockerProductAssetMapping)
		if found {
			return false, types.ErrorLockerProductAssetMappingExists.Error()
		}
	}
	return true, ""
}


func (k Keeper) AddWhiteListedAsset(c context.Context, msg *types.MsgAddWhiteListedAssetRequest) (*types.MsgAddWhiteListedAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	esmStatus, found := k.GetESMStatus(ctx, msg.AppId)
	status := false
	if found {
		status = esmStatus.Status
	}
	if status {
		return nil, esmtypes.ErrESMAlreadyExecuted
	}
	klwsParams, _ := k.GetKillSwitchData(ctx, msg.AppId)
	if klwsParams.BreakerEnable {
		return nil, esmtypes.ErrCircuitBreakerEnabled
	}
	appMapping, found := k.GetApp(ctx, msg.AppId)
	if !found {
		return nil, types.ErrorAppMappingDoesNotExist
	}
	asset, found := k.GetAsset(ctx, msg.AssetId)
	if !found {
		return nil, types.ErrorAssetDoesNotExist
	}
	lockerProductAssetMapping, found := k.GetLockerProductAssetMapping(ctx, msg.AppId)

	if !found {
		//Set a new instance of Locker Product Asset  Mapping

		var locker types.LockerProductAssetMapping
		locker.AppId = appMapping.Id
		locker.AssetIds = append(locker.AssetIds, asset.Id)
		k.SetLockerProductAssetMapping(ctx, locker)

		//Also Create a LockerLookup table Instance and set it with the new asset id
		var lockerLookupData types.LockerLookupTable
		var lockerAssetData types.TokenToLockerMapping

		lockerAssetData.AssetId = asset.Id
		lockerLookupData.Counter = 0
		lockerLookupData.AppId = appMapping.Id
		lockerLookupData.Lockers = append(lockerLookupData.Lockers, &lockerAssetData)
		k.SetLockerLookupTable(ctx, lockerLookupData)
		return &types.MsgAddWhiteListedAssetResponse{}, nil
	}
	// Check if the asset from msg exists or not ,
	found = k.CheckLockerProductAssetMapping(ctx, msg.AssetId, lockerProductAssetMapping)
	if found {
		return nil, types.ErrorLockerProductAssetMappingExists
	}
	lockerProductAssetMapping.AssetIds = append(lockerProductAssetMapping.AssetIds, asset.Id)
	k.SetLockerProductAssetMapping(ctx, lockerProductAssetMapping)
	lockerLookupTableData, _ := k.GetLockerLookupTable(ctx, appMapping.Id)
	var lockerAssetData types.TokenToLockerMapping
	lockerAssetData.AssetId = asset.Id
	lockerLookupTableData.Lockers = append(lockerLookupTableData.Lockers, &lockerAssetData)
	k.SetLockerLookupTable(ctx, lockerLookupTableData)
	return &types.MsgAddWhiteListedAssetResponse{}, nil
}
