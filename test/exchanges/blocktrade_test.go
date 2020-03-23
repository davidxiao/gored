package test

// Copyright (c) 2015-2019 Bitontop Technologies Inc.
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

import (
	"testing"

	"github.com/bitontop/gored/exchange"
	"github.com/bitontop/gored/pair"
	// "../exchange/blocktrade"
	// "./conf"
)

/********************Public API********************/
func Test_Blocktrade(t *testing.T) {
	e := InitEx(exchange.BLOCKTRADE)
	pair := pair.GetPairByKey("EUR|BTC")

	Test_Coins(e)
	Test_Pairs(e)
	Test_Pair(e, pair)
	Test_Orderbook(e, pair)
	Test_ConstraintFetch(e, pair)
	Test_Constraint(e, pair)

	// Test_Balance(e, pair)
	// Test_Trading(e, pair, 0.00000001, 100)
	// Test_Trading_Sell(e, pair, 0.00000001, 667)
	// Test_Withdraw(e, pair.Base, 1, "address")
}
