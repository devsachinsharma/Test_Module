package main

import (
	"fmt"

	"github.com/devsachinsharma/Test_Module/decrypt"
	"github.com/devsachinsharma/Test_Module/encrypt"
)

func main() {
	v := encrypt.Encrypted("Kodeklouds")
	fmt.Println(encrypt.Encrypted("Kodeklouds"))
	fmt.Println(decrypt.Decrypted(v))
}
