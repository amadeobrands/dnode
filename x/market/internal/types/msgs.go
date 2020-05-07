package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = MsgCreateMarket{}
)

type MsgCreateMarket struct {
	Nominee           sdk.AccAddress `json:"nominee" yaml:"nominee"`
	BaseAssetDenom    string         `json:"base_asset_denom" yaml:"base_asset_denom"`
	QuoteAssetDenom   string         `json:"quote_asset_denom" yaml:"quote_asset_denom"`
	BaseAssetDecimals uint8          `json:"base_asset_decimals" yaml:"base_asset_decimals"`
}

func NewMsgCreateMarket(nominee sdk.AccAddress, baseAsset string, quoteAsset string, baseDecimals uint8) MsgCreateMarket {
	return MsgCreateMarket{
		Nominee:           nominee,
		BaseAssetDenom:    baseAsset,
		QuoteAssetDenom:   quoteAsset,
		BaseAssetDecimals: baseDecimals,
	}
}
func (msg MsgCreateMarket) Route() string { return ModuleName }

func (msg MsgCreateMarket) Type() string { return "createMarket" }

func (msg MsgCreateMarket) ValidateBasic() error {
	if msg.Nominee.Empty() {
		return sdkErrors.Wrap(ErrWrongNominee, "empty")
	}
	if msg.BaseAssetDenom == "" {
		return sdkErrors.Wrap(ErrWrongAssetDenom, "BaseAsset")
	}
	if msg.QuoteAssetDenom == "" {
		return sdkErrors.Wrap(ErrWrongAssetDenom, "QuoteAsset")
	}

	return nil
}

func (msg MsgCreateMarket) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgCreateMarket) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Nominee}
}