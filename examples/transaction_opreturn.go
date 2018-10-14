// tokucore
//
// Copyright (c) 2018 TokuBlock
// BSD License

package main

import (
	"fmt"

	"github.com/tokublock/tokucore/xcore"
)

func assertNil(err error) {
	if err != nil {
		panic(err)
	}
}

// Demo for OP_RETURN transaction.
func main() {
	msg := []byte("666...satoshi")

	seed := []byte("this.is.bohu.seed.")
	bohuHDKey := xcore.NewHDKey(seed)
	bohuPrv := bohuHDKey.PrivateKey()
	bohu := bohuHDKey.GetAddress()

	// Satoshi.
	seed = []byte("this.is.satoshi.seed.")
	satoshiHDKey := xcore.NewHDKey(seed)
	satoshi := satoshiHDKey.GetAddress()

	// Output:
	bohuCoin := xcore.NewCoinBuilder().AddOutput(
		"faf1520f1d3e818fca695c2a903baa4a7eec4954f0b35aa01be1f2c1d2cfd802",
		0,
		1*xcore.Unit,
		"76a9145a927ddadc0ef3ae4501d0d9872b57c9584b9d8888ac",
	).ToCoins()[0]

	tx, err := xcore.NewTransactionBuilder().
		AddCoins(bohuCoin).
		AddKeys(bohuPrv).
		To(satoshi, 666666).
		Then().
		SetChange(bohu).
		SendFees(10000).
		Then().
		AddPushData(msg).
		Sign().
		BuildTransaction()
	assertNil(err)

	fmt.Printf("opreturn:%v\n", tx.ToString())
	fmt.Printf("opreturn.txid:%x\n", tx.ID())
	fmt.Printf("opreturn.tx:%x\n", tx.Serialize())
}
