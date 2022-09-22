package main

import (
	"fmt"
	"tests-intro/address"
)

func main() {
	addressType := address.AddressType("Avenue Paulista")
	fmt.Println(addressType)
}