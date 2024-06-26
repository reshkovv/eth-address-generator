package main

import (
	"crypto/ecdsa"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type ethAddress struct {
	Address    string
	PrivateKey string
}

func (addr *ethAddress) Print() {
	fmt.Printf("Address: %v  \nPrivate key: %v\n", addr.Address, addr.PrivateKey)
}

func main() {
	var quantity int
	var prefix string
	var suffix string
	var contains string

	flag.IntVar(&quantity, "n", 1, "Number of addresses generated")
	flag.StringVar(&prefix, "p", "", "Address prefix")
	flag.StringVar(&suffix, "s", "", "Address suffix")
	flag.StringVar(&contains, "q", "", "Address substring")

	flag.Parse()

	counter := 0
	generatedAddresses := make([]ethAddress, 0)
	for {
		if counter >= quantity {
			for _, address := range generatedAddresses {
				address.Print()
				fmt.Println()
			}
			return
		}
		addr := generateAddress()
		addrLower := strings.ToLower(addr.Address)

		if strings.HasPrefix(addrLower, "0x"+prefix) &&
			strings.HasSuffix(addrLower, suffix) &&
			strings.Contains(addrLower, contains) {
			generatedAddresses = append(generatedAddresses, addr)
			counter++
		}
	}
}

func generateAddress() ethAddress {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return ethAddress{address, hexutil.Encode(privateKeyBytes)[2:]}
}
