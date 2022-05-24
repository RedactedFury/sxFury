package keeper

import (
	"fmt"
	lockertypes "github.com/comdex-official/comdex/x/locker/types"
	"github.com/comdex-official/comdex/x/rewards/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) Iterate(ctx sdk.Context, appMappingId uint64, assetIds []uint64) error {
	CollectorAppAsset, _ := k.GetAppToDenomsMapping(ctx, appMappingId)
	for i := range assetIds {
		found := uint64InSlice(assetIds[i], CollectorAppAsset.AssetIds)
		if !found {
			return types.ErrAssetIdDoesNotExist
		}
		CollectorLookup, _ := k.GetCollectorLookupByAsset(ctx, appMappingId, assetIds[i])
		for _, j := range CollectorLookup.AssetrateInfo {
			LockerProductAssetMapping, _ := k.GetLockerLookupTable(ctx, appMappingId)
			lockers := LockerProductAssetMapping.Lockers
			for _, v := range lockers {
				if v.AssetId == assetIds[i] {
					lockerIds := v.LockerIds
					for w := range lockerIds {
						locker, _ := k.GetLocker(ctx, lockerIds[w])
						balance := locker.NetBalance
						rewards, err := k.CalculateRewards(ctx, balance, *j.LockerSavingRate)
						if err != nil {
							return nil
						}
						// update the lock position
						returnsAcc := locker.ReturnsAccumulated
						updatedReturnsAcc := rewards.Add(returnsAcc)
						netBalance := locker.NetBalance.Add(rewards)
						updatedLocker := lockertypes.Locker{
							LockerId:           locker.LockerId,
							Depositor:          locker.Depositor,
							ReturnsAccumulated: updatedReturnsAcc,
							NetBalance:         netBalance,
							CreatedAt:          locker.CreatedAt,
							AssetDepositId:     locker.AssetDepositId,
							IsLocked:           locker.IsLocked,
							AppMappingId:       locker.AppMappingId,
						}
						k.UpdateLocker(ctx, updatedLocker)
					}
				}
			}
		}
	}
	return nil
}

func (k Keeper) CalculateRewards(ctx sdk.Context, amount sdk.Int, lsr sdk.Dec) (sdk.Int, error) {

	LockerSavingsRate := lsr.Quo(sdk.OneDec())
	currentTime := ctx.BlockTime().Unix()

	prevInterestTime := k.GetLastInterestTime(ctx)
	if prevInterestTime == 0 {
		prevInterestTime = currentTime
	}

	secondsElapsed := currentTime - prevInterestTime
	if secondsElapsed < 0 {
		return sdk.ZeroInt(), sdkerrors.Wrap(types.ErrNegativeTimeElapsed, fmt.Sprintf("%d seconds", secondsElapsed))
	}
	yearsElapsed := sdk.NewDec(secondsElapsed).QuoInt64(types.SecondsPerYear)

	newAmount := sdk.NewDecFromInt(amount.Mul(sdk.Int(LockerSavingsRate))).Mul(yearsElapsed).QuoInt64(100)

	err := k.SetLastInterestTime(ctx, currentTime)
	if err != nil {
		return sdk.ZeroInt(), err
	}

	return sdk.Int(newAmount), nil
}

func (k Keeper) IterateVaults(ctx sdk.Context, appMappingId uint64) error {
	extVaultMapping, _ := k.GetAppExtendedPairVaultMapping(ctx, appMappingId)
	for _, v := range extVaultMapping.ExtendedPairVaults {
		vaultIds := v.VaultIds
		for j, _ := range vaultIds {
			vault, _ := k.GetVault(ctx, vaultIds[j])
			ExtPairVault, _ := k.GetPairsVault(ctx, vault.ExtendedPairVaultID)
			StabilityFee := ExtPairVault.StabilityFee

			if StabilityFee != sdk.ZeroDec() {
				interest, _ := k.CalculateRewards(ctx, vault.AmountOut, StabilityFee)
				intAcc := vault.InterestAccumulated
				updatedIntAcc := (intAcc).Add(interest)
				vault.InterestAccumulated = &updatedIntAcc
				//update vault
				k.SetVault(ctx, vault)
			}
		}
	}
	return nil
}

func (k Keeper) DistributeExtRewardCollector(ctx sdk.Context) error {
	extRewards := k.GetExternalRewardsLockers(ctx)
	for i, v := range extRewards {
		epochTime, _ := k.GetEpochTime(ctx, extRewards[i].AppMappingId)
		et := epochTime
		timeNow := ctx.BlockTime().Unix()

		if et > timeNow {

			if extRewards[i].IsActive == true {
				count, _ := k.GetExternalRewardsLockersCounter(ctx, extRewards[i].Id)

				if count < uint64(extRewards[i].DurationDays) {
					lockerLookup, _ := k.GetLockerLookupTable(ctx, v.AppMappingId)
					for _, u := range lockerLookup.Lockers {
						if u.AssetId == v.AssetId {
							lockerIds := u.LockerIds
							totalShare := u.DepositedAmount
							for w, _ := range lockerIds {
								locker, _ := k.GetLocker(ctx, lockerIds[w])
								userShare := locker.NetBalance.Quo(totalShare)
								totalRewards := k.GetExternalRewardsLocker(ctx, v.Id).TotalRewards
								Duration := k.GetExternalRewardsLocker(ctx, v.Id).DurationDays
								rewardsPerEpoch := (totalRewards.Amount).Quo(sdk.NewInt(Duration))
								dailyRewards := userShare.Mul(rewardsPerEpoch)
								err := k.SendCoinFromModuleToModule(ctx, types.ModuleName, locker.Depositor, sdk.NewCoins(sdk.NewCoin(totalRewards.Denom, dailyRewards)))
								if err != nil {
									return err
								}
								k.SetExternalRewardsLockersCounter(ctx, extRewards[i].Id, count+1)
							}
						}
					}
				} else {
					extRewards[i].IsActive = false
					k.SetExternalRewardsLockers(ctx, extRewards[i])
				}
			}
		}
		k.SetEpochTime(ctx, extRewards[i].AppMappingId, et+84600)
	}
	return nil
}
