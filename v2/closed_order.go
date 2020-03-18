// Copyright 2019 Tokenomy Technologies Ltd. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package v2

import "github.com/shuLhan/share/lib/math/big"

//
// ClosedOrder represent combination of closed ask and bid.
//
type ClosedOrder struct {
	ID         int64    `json:"id"`
	Type       string   `json:"type"`
	Method     string   `json:"method"`
	Asset      string   `json:"asset"`
	Price      *big.Rat `json:"price"`
	Amount     *big.Rat `json:"amount"`
	Remain     *big.Rat `json:"remain,omitempty"`
	Status     string   `json:"status"`
	SubmitTime int64    `json:"submit_time"`
	FinishTime int64    `json:"finish_time"`
}
