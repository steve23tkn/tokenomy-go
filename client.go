// Copyright 2020 Tokenomy Technologies Ltd. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package tokenomy

import "github.com/shuLhan/share/lib/math/big"

//
// Client is an interface for API v1 and v2.
//
type Client interface {
	TradeAsk(method, pairName string, amount, price *big.Rat) (*TradeResponse, error)
	TradeBid(method, pairName string, amount, price *big.Rat) (*TradeResponse, error)
	TradeCancelAsk(pairName string, orderID int64) (*TradeResponse, error)
	TradeCancelBid(pairName string, orderID int64) (*TradeResponse, error)
}
