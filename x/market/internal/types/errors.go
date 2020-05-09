package types

import sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrInternal = sdkErrors.Register(ModuleName, 100, "internal")
	// ID is invalid or not found.
	ErrWrongID = sdkErrors.Register(ModuleName, 101, "wrong ID")
	// AssetDenom is empty.
	ErrWrongAssetDenom = sdkErrors.Register(ModuleName, 102, "wrong asset denom, should not be empty")
	// Nominee not exists or empty.
	ErrWrongNominee = sdkErrors.Register(ModuleName, 103, "wrong nominee")
	// Market already exists.
	ErrMarketExists = sdkErrors.Register(ModuleName, 104, "market exists")
	// Base to Quote asset quantity convert failed.
	ErrInvalidQuantity = sdkErrors.Register(ModuleName, 105, "base to quote asset quantity normalization failed")
)
