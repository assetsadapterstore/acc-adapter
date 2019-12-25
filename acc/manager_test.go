/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package acc

import (
	"testing"

	"github.com/blocktree/openwallet/openwallet"
	eos "github.com/eoscanada/eos-go"
	"github.com/siddontang/go/log"
)

func TestGetTokenBalanceByAddress(t *testing.T) {
	wm := NewWalletManager(nil)

	wm.Config.ServerAPI = "https://accchain.io"
	wm.Api = eos.New(wm.Config.ServerAPI)

	contract := openwallet.SmartContract{
		Address:  "accio.token:ACC",
		Symbol:   "ACC",
		Name:     "ACC",
		Decimals: 4,
	}
	balance, err := wm.ContractDecoder.GetTokenBalanceByAddress(contract, "accjiahua111")
	if err != nil {
		log.Error("GetTokenBalanceByAddress failed, unexpected error:", err)
		return
	}

	log.Info("balance:", balance)

}
