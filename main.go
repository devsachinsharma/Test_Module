package main

import (
	"fmt"

	"github.com/devsachinsharma/Test_Module/decrypt"
	"github.com/devsachinsharma/Test_Module/encrypt"
)

func main() {
	v := encrypt.Encrypted("Kodekloud")
	fmt.Println(encrypt.Encrypted("Kodekloud"))
	fmt.Println(decrypt.Decrypted(v))
}
