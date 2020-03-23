package test

import (
	"testing"

	"github.com/bitontop/gored/exchange"
	"github.com/bitontop/gored/pair"
)

// Copyright (c) 2015-2019 Bitontop Technologies Inc.
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

/********************Public API********************/

func Test_Dragonex(t *testing.T) {
	e := InitEx(exchange.DRAGONEX)
	pair := pair.GetPairByKey("ETH|WAN")

	Test_Coins(e)
	Test_Pairs(e)
	Test_Pair(e, pair)
	Test_Orderbook(e, pair)
	Test_ConstraintFetch(e, pair)
	Test_Constraint(e, pair)

	Test_Balance(e, pair)
	// Test_Trading(e, pair, 0.0015, 1)
	// Test_Trading_Sell(e, pair, 0.0022, 1)
	// Test_Withdraw(e, pair.Base, 1, "ADDRESS")
}
