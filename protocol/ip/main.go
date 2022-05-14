package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	checkIP("127.0.0.1")
	checkIP("127.0-1.90.100")
	checkIP("237.500.900")
	checkIP("0:0:0:0:0:0:0:1")
	checkIP("0:0:0:0:0:0:0:19999")
	checkIP("fda3:97c:1eb:fff0:5444:903a:33f0:3a6b")

	resolveIP("ip", "https://google.com")
	resolveIP("ip", "https://www.google.com")
	resolveIP("ip6", "www.google.com")

}

func checkPort(port string) {

}

func resolveIP(netname, name string) {
	addr, err := net.ResolveIPAddr(netname, name)
	if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}
	fmt.Println("(", netname, ")Resolved address for ", name, " is ", addr.String())

}

func checkIP(name string) {
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println(name, " : Invalid address")
	} else {
		fmt.Println("The address is ", addr.String(), " and ", addr)

		// subnet mask itu misahin network address sama host address, tiap ip itu punya 2 component itu
		// misal ip 192.168.3.22
		// maka networknya : 192.168.3.0 ini tergantung settingnya sebenernya
		// maka hostnya : 192.168.1.0 host itu selalu bit pertama (right most bit) first host (00000001)
		// last host address nya : 192.168.254.0 -> 254 dari 1111110

		ones := 24
		bits := 32

		mask := net.CIDRMask(ones, bits)
		network := addr.Mask(mask)
		fmt.Println("Address is ", addr.String(),
			"\nMask length is ", bits,
			"\nLeading ones count is ", ones,
			"\nMask is (hex) ", mask.String(),
			"\nNetwork is ", network.String(),
			"\n----\n")
	}
}
