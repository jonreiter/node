/**
 * generates the special bridge addresses in both hex and wallet-address format
 */
package main

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/bech32"
)

func main() {
	for _, b := range [][]byte{[]byte("BinanceChainStakeDelegation"), []byte("BinanceChainDepositedCoins")} {
		fmt.Println("for wallet string:", string(b))
		DepositedCoinsAccAddr := types.AccAddress(crypto.AddressHash(b))
		encoded, _ := bech32.ConvertAndEncode("bnb", DepositedCoinsAccAddr)
		fmt.Println("as hex:", DepositedCoinsAccAddr)
		fmt.Println("as bnb address:", encoded)
	}
}
