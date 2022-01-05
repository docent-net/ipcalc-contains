package tests

import (
	"testing"
	"github.com/docent-net/ipcalc-contains/subnet"
)

func TestSubnetContainsShouldFailUncrecognizedIPv4Network(t *testing.T) {
	args := []string{"aaa", "192.168.123.0/30"}

	if subnet.SubnetContains(args) != false {
		t.Errorf("Expected %s to fail CIDR verification, but it did succeed", args[0])
	}
}

func TestSubnetContainsShouldFailUncrecognizedIPv4(t *testing.T) {
	args := []string{"192.168.1.0/20", "abc"}

	if subnet.SubnetContains(args) != false {
		t.Errorf("Expected %s to fail CIDR verification, but it did succeed", args[1])
	}
}

func TestSubnetContainsNetworkShouldContainNetwork(t *testing.T) {
	args := []string{"192.168.0.0/20", "192.168.12.0/30"}

	if subnet.SubnetContains(args) != true {
		t.Errorf("Expected %s to contain %s, but it failed", args[0], args[1])
	}
}

func TestSubnetContainsNetworkShouldNotContainNetwork(t *testing.T) {
	args := []string{"192.168.0.0/20", "192.168.128.0/30"}

	if subnet.SubnetContains(args) != false {
		t.Errorf("Expected %s to NOT contain %s, but it did", args[0], args[1])
	}
}

func TestSubnetContainsNetworkShouldContainIPv4(t *testing.T) {
	args := []string{"192.168.0.0/20", "192.168.12.123"}

	if subnet.SubnetContains(args) != true {
		t.Errorf("Expected %s to contain %s, but it failed", args[0], args[1])
	}
}

func TestSubnetContainsNetworkShouldNotContainIPv4(t *testing.T) {
	args := []string{"192.168.0.0/20", "10.0.0.12"}

	if subnet.SubnetContains(args) != false {
		t.Errorf("Expected %s to NOT contain %s, but it did", args[0], args[1])
	}
}