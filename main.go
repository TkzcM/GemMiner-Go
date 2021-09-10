package main

import (
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/miguelmota/go-solidity-sha3"
	"math/big"
)
// no need to change these lines
var nonce = flag.Int("nonce",0,"your nonce")
var diffInt = flag.Int("diff",114514,"difficult now")
var gemKind = flag.Int("kind",1,"gem kind")
var address = flag.String("address","0x4E6FEC28f5316C2829D41Bc2187202c70EC75Bc7","fantom address")
var saltStart = flag.String("salt","2300000","salt pointer")

func main() {
	flag.Parse()
	realNonce := int64(*nonce)
	realDiffInt := int64(*diffInt)
	salt, _ := new(big.Int).SetString(*saltStart,10)
	plus := big.NewInt(1)
	realGemKind := int64(*gemKind)
	z := new(big.Int)
	uintMax, _ := z.SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935",10)
	diff := big.NewInt(realDiffInt)
	var result = new(big.Int).Div(uintMax,diff)
	for true{
		hash := solsha3.SoliditySHA3(
			// types
			[]string{"uint256", "bytes32","address","address","uint256","uint256","uint256"},

			// values
			[]interface{}{
				big.NewInt(250),
				"0x000080440000047163a56455ac4bc6b1f1b88efadf17db76e5c52c0ca594fd9b",
				common.HexToAddress("0x342EbF0A5ceC4404CcFF73a40f9c30288Fc72611"),
				common.HexToAddress(*address),
				big.NewInt(realGemKind),
				big.NewInt(realNonce),
				salt,
			},
		)
		var luck = new(big.Int).SetBytes(hash)
		if luck.Cmp(result) != 1{
			fmt.Println(salt)
			break
		}
		salt.Add(salt,plus)
	}
}