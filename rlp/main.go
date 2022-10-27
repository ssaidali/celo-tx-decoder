package main

import (
	"log"

	"github.com/celo-org/celo-blockchain/core/types"
	"github.com/celo-org/celo-blockchain/rlp"
)

func main() {
	data := "0xf8691c841dcd6500826aa480808094aab3a017541d065e5bbeccd28b3ccbfd18559d61018083015e0aa0322a3595b0caf92f8e3021213a0dd91539cfedc23ba49706afac1d7f5ac333c3a04075543e78032e57d92c88d137e28fb683014363e159545c359158390e2af47a"

	log.Printf("rawTx: %v", data)

	tx, err := decode(data)

	if err != nil {
		log.Printf("err: %v", err)
		return
	}

	log.Printf("tx: %v", tx)
}

func decode(data string) (*types.Transaction, error) {
	var tx types.Transaction
	t, err := &tx, rlp.DecodeBytes([]byte(data), &tx)
	return t, err
}
