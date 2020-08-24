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
	"github.com/astaxie/beego/config"
	"path/filepath"
	"testing"

	"github.com/blocktree/openwallet/v2/openwallet"
	"github.com/siddontang/go/log"
)

func testNewWalletManager() *WalletManager {
	wm := NewWalletManager(nil)
	//读取配置
	absFile := filepath.Join("conf", "ACC.ini")
	//log.Debug("absFile:", absFile)
	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		return nil
	}
	wm.LoadAssetsConfig(c)
	wm.Api.Debug = true
	return wm
}

func TestWalletManager_GetInfo(t *testing.T) {
	wm := testNewWalletManager()
	r, err := wm.Api.GetInfo()
	if err != nil {
		log.Errorf("unexpected error: %v", err)
		return
	}
	log.Infof("%+v", r)
}

func TestWalletManager_GetAccount(t *testing.T) {
	wm := testNewWalletManager()
	r, err := wm.Api.GetAccount("aaaaaaaaaaaa")
	if err != nil {
		log.Errorf("unexpected error: %v", err)
		return
	}
	log.Infof("%+v", r)
}

func TestGetTokenBalanceByAddress(t *testing.T) {
	wm := testNewWalletManager()

	contract := openwallet.SmartContract{
		Address:  "accio.token:ACC",
		Symbol:   "ACC",
		Name:     "ACC",
		Decimals: 4,
	}
	balance, err := wm.ContractDecoder.GetTokenBalanceByAddress(contract, "blockins.acc")
	if err != nil {
		log.Error("GetTokenBalanceByAddress failed, unexpected error:", err)
		return
	}

	log.Info("balance:", balance)

}
