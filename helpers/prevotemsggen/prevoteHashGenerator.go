package main

import (
	"fmt"
	"strings"

	"encoding/hex"

	"github.com/tendermint/tendermint/crypto/tmhash"
)

func main() {
	// codeHash provided is the complete hash(not truncated)
	// codeHash is also
	var salt, codeHash, operatorAddr string
	fmt.Println("Enter salt codehash and operatoraddress (space separated)")
	fmt.Scan(&salt, &codeHash, &operatorAddr)

	standardFormatInput := strings.Join([]string{salt, codeHash, operatorAddr}, ":")
	prevotehash := tmhash.SumTruncated([]byte(standardFormatInput))
	// prevotehash := sha256.Sum256([]byte(standardFormatInput))

	// fmt.Print("The prevote hash is : ", string(prevotehash[:20]))
	fmt.Println("The prevote hash is : ", hex.EncodeToString(prevotehash[:20]))
}

// 0c346f992ca2a96fcb6b1901d60ff6d65282f6fc
