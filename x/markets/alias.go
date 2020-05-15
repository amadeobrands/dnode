package markets

import (
	"github.com/dfinance/dnode/x/markets/internal/keeper"
	"github.com/dfinance/dnode/x/markets/internal/types"
)

type (
	Keeper         = keeper.Keeper
	Market         = types.Market
	Markets        = types.Markets
	MarketExtended = types.MarketExtended
)

const (
	ModuleName        = types.ModuleName
	DefaultParamspace = types.DefaultParamspace
)

var (
	// variable aliases
	ModuleCdc = types.ModuleCdc
	// function aliases
	NewKeeper         = keeper.NewKeeper
	NewMarket         = types.NewMarket
	NewMarketExtended = types.NewMarketExtended
)