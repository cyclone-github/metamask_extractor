package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

// program to extract metamask hash from .ldb file
// ldb files can be found at ..Google\Chrome\User Data\Default\Local Extension Settings\nkbihfbeogaeaoehlefnkodbefgpgknn
// https://support.metamask.io/hc/en-us/articles/360018766351-How-to-use-the-Vault-Decryptor-with-the-MetaMask-Vault-Data
// inspiration from 'extract-metamask-vaults.py' and my own shell commands for extracting metamask hashes
// coded by cyclone in Go

// version 0.1.1; initial github release

func printVersion() {
	fmt.Println("Program version: 0.1.1")
}

func printCyclone() {
	encoded := "Q29kZWQgYnkgQ3ljbG9uZSA6KQ=="
	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	fmt.Println(string(decoded))
}

func printHelp() {
	fmt.Println("Usage: metamask_extractor [-version] [-help] [-input <ldb file or directory containing ldb files>]")
}

func main() {
	versionFlag := flag.Bool("version", false, "Print program version")
	cycloneFlag := flag.Bool("cyclone", false, "")
	helpFlag := flag.Bool("help", false, "Print usage information")
	inputFlag := flag.String("input", "", "ldb file or directory containing ldb files")

	flag.Parse()

	if *versionFlag {
		printVersion()
		return
	}

	if *cycloneFlag {
		printCyclone()
		return
	}

	if *helpFlag {
		printHelp()
		return
	}

	if *inputFlag == "" {
		fmt.Fprintln(os.Stderr, "Error: -input flag is required")
		printHelp()
		os.Exit(0)
	}

	ldbFile := *inputFlag
	ldbDir := filepath.Dir(ldbFile)

	db, err := leveldb.OpenFile(ldbDir, &opt.Options{})
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to open ldb file:", err)
		os.Exit(0)
	}
	defer db.Close()

	var currentData bool
	var data json.RawMessage

	iter := db.NewIterator(nil, nil)
	defer iter.Release()
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()

		if bytes.Contains(key, []byte("data")) {
			if err := json.Unmarshal(value, &data); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to unmarshal json record value: %v\n", err)
				continue
			}

			dataStr := string(data)
			dataStr = strings.ReplaceAll(dataStr, "\\", "")
			if strings.Contains(dataStr, "salt") {
				if currentData {
					fmt.Println()
				} else {
					currentData = true
					fmt.Println("===== Current Vault Data =====")
				}
				walletStartText := "vault"
				walletDataStart := strings.Index(strings.ToLower(dataStr), walletStartText)
				walletDataTrimmed := dataStr[walletDataStart:]
				walletDataStart = strings.Index(walletDataTrimmed, "data")
				walletDataTrimmed = walletDataTrimmed[walletDataStart-2:]
				walletDataEnd := strings.Index(walletDataTrimmed, "}")
				walletData := walletDataTrimmed[:walletDataEnd+1]
				fmt.Println(walletData)

				// Extract salt, IV, and data values from json
				type Hash struct {
					Salt      string `json:"salt"`
					Iv        string `json:"iv"`
					VaultData string `json:"data"`
				}

				var h Hash
				err = json.Unmarshal([]byte(walletData), &h)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Failed to unmarshal salt, iv, and data values: %v\n", err)
					continue
				}

				// Print Metamask hash string
				metamaskHash := fmt.Sprintf("$metamask$%s$%s$%s", h.Salt, h.Iv, h.VaultData)
				fmt.Println("\n===== Metamask -m 26600 hash =====")
				fmt.Println(metamaskHash)
			}
		}
	}
}

// end code
