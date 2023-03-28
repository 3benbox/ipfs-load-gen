package main

import (
	// "context"

	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	shell "github.com/ipfs/go-ipfs-api"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	payloadLengthMax = 2048
	payloadLengthMin = 1024
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func main() {
	ipfsApiAddress := os.Getenv("IPFS_API_URL")
	if ipfsApiAddress == "" {
		ipfsApiAddress = "localhost:5001"
	}

	sh := shell.NewShell(ipfsApiAddress)
	for {
		payloadLength := rand.Intn(payloadLengthMax-payloadLengthMin+1) + payloadLengthMin //(rand.Intn(max - min + 1) + min)
		payloadString := StringWithCharset(payloadLength, charset)
		cid, err := sh.Add(strings.NewReader(payloadString))
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s", err)
			os.Exit(1)
		}
		fmt.Printf("added %s\n", cid)
	}

}
