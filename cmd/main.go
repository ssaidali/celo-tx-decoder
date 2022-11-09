package main

import (
	"log"
	"math/big"

	"github.com/mstraka100/celo-blockchain-small/types"
	"github.com/mstraka100/celo-blockchain-small/rlp"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func main() {
	chainID := big.NewInt(44787)
	rawTxString := "0xf8691c841dcd6500826aa480808094aab3a017541d065e5bbeccd28b3ccbfd18559d61018083015e0aa0322a3595b0caf92f8e3021213a0dd91539cfedc23ba49706afac1d7f5ac333c3a04075543e78032e57d92c88d137e28fb683014363e159545c359158390e2af47a"

	log.Printf("rawTx: %v", rawTxString)

	rawTxBytes, _ := hexutil.Decode(rawTxString)

	tx, err := decode(rawTxBytes)

	if err != nil {
		log.Printf("err: %v", err)
		return
	}

	log.Printf("to: %s", tx.To().String())

	from, err := types.Sender(types.NewLondonSigner(chainID), tx)
	if err != nil {
		log.Printf("err: %v", err)
		return
	}

	log.Printf("from: %s", from.String())
}

func decode(data []byte) (*types.Transaction, error) {
	var tx types.Transaction
	t, err := &tx, rlp.DecodeBytes(data, &tx)
	return t, err
}
