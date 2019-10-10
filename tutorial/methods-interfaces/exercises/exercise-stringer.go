package main

import (
  "fmt"
  "strings"
  "strconv"
)

type IPAddr [4]byte

func (p IPAddr) String() string {
    var ipParts []string
    for _, item := range p {
    // can not use string(item) or string(int(item))
    // string will convert int or byte to their ASCII presentation
    // not literal numeric string like "0"
	ipParts = append(ipParts, strconv.Itoa(int(item)))
    }
    return strings.Join(ipParts, ".")
}

func main() {
    hosts := map[string]IPAddr{
	"loopback":  {127, 0, 0, 1},
	"googleDNS": {8, 8, 8, 8},
    }
    for name, ip := range hosts {
	fmt.Printf("%v: %v\n", name, ip)
    }
}
