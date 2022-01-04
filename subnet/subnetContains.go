package subnet

import (
	"fmt"
	"net"
	"os"
)

func SubnetContains(args []string) bool {
	// args[0]: should be a network in a slash notation: A.B.C.D/Y
	// args[1]: should be an IP addr or a network in a slash notation

	_, ipv4Net, err := net.ParseCIDR(args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ipv4Addr := net.ParseIP(args[1])
	if ipv4Addr == nil {
		// maybe its a network?
		_, ipv4Net2, err := net.ParseCIDR(args[1])
		if err != nil {
			fmt.Println(string(args[1]) + " is not recognized as IPv4 address nor a network")
			os.Exit(1)
		}
		// checking if subnet contains another subnet:
		if subnetContainsSubnet(ipv4Net, ipv4Net2) {
			return true
		}
	} else {
		// checking if subnet contains an IP addr:
		if ipv4Net.Contains(ipv4Addr) {
			return true
		}
	}

	return false
}

func subnetContainsSubnet(net1 *net.IPNet, net2 *net.IPNet) bool {
	// function verifies whether net1 contains net2

	// for two networks A and B, if network A has smaller mask (lower number
	// in slash notation), and network A consists B's IP address,
	// than B is contained by A

	// example:
	//
	// network A: 192.168.0.0/20
	// network B: 192.168.12.0/30
	//
	// 1. network A has smaller mask (/20) than network's B mask (/30)
	// 2. network A consists B's IP address 192.168.12.0
	//
	// thus network A 192.168.0.0/20 contains network B 192.168.12.0/30

	net1Mask, _ := net1.Mask.Size()
    net2Mask, _ := net2.Mask.Size()

    if net1.Contains(net2.IP) && net2Mask >= net1Mask {
		return true
	}

	return false
}