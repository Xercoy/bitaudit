package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/xercoy/bitaudit"
)

func main() {
	var auditFile, hash string

	flag.StringVar(&auditFile, "audit-file", "", "File to audit.")
	flag.StringVar(&hash, "hash", "", "name of hash to use.")

	flag.Parse()

	hash, err := bitaudit.CreateHash(auditFile, hash)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("%x", hash)
}
