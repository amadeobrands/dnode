package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/dfinance/dnode/x/order/internal/types"
)

func (k Keeper) LockOrderCoins(ctx sdk.Context, order types.Order) error {
	coin, err := order.LockCoin()
	if err != nil {
		return sdkErrors.Wrap(err, "creating lock coin")
	}

	if err = k.supplyKeeper.SendCoinsFromAccountToModule(ctx, order.Owner, types.ModuleName, sdk.NewCoins(coin)); err != nil {
		return sdkErrors.Wrapf(types.ErrInternal, "locking coins: %v", err)
	}

	return nil
}

func (k Keeper) UnlockOrderCoins(ctx sdk.Context, order types.Order) error {
	coin, err := order.LockCoin()
	if err != nil {
		return sdkErrors.Wrap(err, "creating unlock coin")
	}

	if err = k.supplyKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, order.Owner, sdk.NewCoins(coin)); err != nil {
		return sdkErrors.Wrapf(types.ErrInternal, "unlocking coins: %v", err)
	}

	return nil
}

func (k Keeper) ExecuteOrderFills(ctx sdk.Context, orderFills types.OrderFills) {
	for _, orderFill := range orderFills {
		fillCoin, err := orderFill.FillCoin()
		if err != nil {
			k.GetLogger(ctx).Debug(orderFill.String())
			k.GetLogger(ctx).Error(fmt.Sprintf("creating fill coin: %v", err))
			continue
		}
		if _, err = k.bankKeeper.AddCoins(ctx, orderFill.Order.Owner, sdk.NewCoins(fillCoin)); err != nil {
			k.GetLogger(ctx).Debug(orderFill.String())
			panic(fmt.Sprintf("transfering fill coins: %v", err))
		}

		doRefund, refundCoin, err := orderFill.RefundCoin()
		if err != nil {
			k.GetLogger(ctx).Debug(orderFill.String())
			k.GetLogger(ctx).Error(fmt.Sprintf("creating refund coin: %v", err))
			continue
		}
		if doRefund {
			if refundCoin != nil {
				if _, err = k.bankKeeper.AddCoins(ctx, orderFill.Order.Owner, sdk.NewCoins(*refundCoin)); err != nil {
					k.GetLogger(ctx).Debug(orderFill.String())
					panic(fmt.Sprintf("adding refund coins: %v", err))
				}
			} else {
				k.GetLogger(ctx).Debug(orderFill.String())
				k.GetLogger(ctx).Info(fmt.Sprintf("order refund amount is too small: %s", orderFill.Order.ID))
			}
		}

		if orderFill.QuantityUnfilled.IsZero() {
			k.GetLogger(ctx).Info(fmt.Sprintf("order completely filled: %s", orderFill.Order.ID))
			k.Del(ctx, orderFill.Order.ID)
		} else {
			k.GetLogger(ctx).Info(fmt.Sprintf("order partially filled: %s", orderFill.Order.ID))
			orderFill.Order.Quantity = orderFill.QuantityUnfilled
			k.Set(ctx, orderFill.Order)
		}
	}
}
