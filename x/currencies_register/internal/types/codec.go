package types

import "github.com/cosmos/cosmos-sdk/codec"

// module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	ModuleCdc.Seal()
}
