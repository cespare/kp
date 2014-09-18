package main

import (
	"crypto/aes"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/howeyc/gopass"
)

// https://github.com/keepassx/keepassx/blob/master/src/format/KeePass2.h
// https://github.com/keepassx/keepassx/blob/master/src/format/KeePass2Reader.cpp
// http://keepass.info/help/base/importexport.html

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s FILENAME", os.Args[0])
	}
	encrypted, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Master pw: ")
	pw := gopass.GetPasswd()
	key := sha256.Sum256(pw)
	cipher, err := aes.NewCipher(key[:])
	if err != nil {
		log.Fatal(err)
	}
	decrypted := make([]byte, len(encrypted))
	cipher.Decrypt(decrypted, encrypted)
	os.Stdout.Write(decrypted)
}
