package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ipA IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ipA[0], ipA[1], ipA[2], ipA[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for _, ip := range hosts {
		fmt.Println(ip)
	}
}

