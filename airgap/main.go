package main

import (
	"log"
	"math/big"

	"github.com/celo-org/rosetta/airgap"
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

func decode(data string) (airgap.Transaction, error) {
	var tx airgap.Transaction
	err := tx.Deserialize([]byte(data), big.NewInt(44787))
	return tx, err
}
