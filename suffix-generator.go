package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

const fileName = "keys_.csv"

func genKeyPair() (string, string) {
	key, _ := crypto.GenerateKey()
	addr := crypto.PubkeyToAddress(key.PublicKey)
	addrStr := "0x" + hex.EncodeToString(addr[:])
	keyStr := hex.EncodeToString(crypto.FromECDSA(key))
	return addrStr, keyStr
}

func worker(suffix string) {
	for i := 0; i != -1; i++ {
		addr, priv := genKeyPair()
		if strings.Contains(strings.ToLower(addr), strings.ToLower(suffix)) {
			fmt.Printf("Found address with suffix %v --> %v\n", suffix, addr)
			// create file
			keyFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal("Can't create file:", err)
			}

			// write to file
			dataWriter := bufio.NewWriter(keyFile)
			dataWriter.WriteString(fmt.Sprintf("%v;%v\n", addr, priv))
			dataWriter.Flush()
			keyFile.Close()
			os.Exit(0)
		}
	}
}

func main() {
	var threadsFlag = flag.Int("threads", 5, "The number of threads the script will work with")
	var suffix = flag.String("suffix", "12345", "Address suffix, for example 0x12345 or 0xfffff")
	flag.Parse()

	log.Printf("Number of threads %v\nGenerating address with suffix %v", threadsFlag, suffix)

	for n := 0; n != *threadsFlag; n++ {
		go worker(*suffix)
	}

	var wait1 string
	fmt.Scan(&wait1)
}
