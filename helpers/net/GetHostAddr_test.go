package net

import (
	"errors"
	"net"
	"os"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

// Return IP address from environment variable when it exists
func TestGetHostAddrReturnsEnvVar(t *testing.T) {
	// Arr matey! Let's check if we can find treasure in them env vars!
	os.Setenv("TEST_IP", "192.168.1.100")
	defer os.Unsetenv("TEST_IP")

	result := GetHostAddr("TEST_IP")

	assert.Equal(t, "192.168.1.100", result)
}

// Return first valid IPv4 LAN address when environment variable is empty
func TestGetHostAddrReturnsFirstValidLANIP(t *testing.T) {
	// Shiver me timbers! Time to find our local network booty!
	mockAddrs := []net.Addr{
		&net.IPNet{IP: net.ParseIP("192.168.1.5"), Mask: net.IPv4Mask(255, 255, 255, 0)},
	}

	patch := monkey.Patch(net.InterfaceAddrs, func() ([]net.Addr, error) {
		return mockAddrs, nil
	})
	defer patch.Unpatch()

	result := GetHostAddr("NONEXISTENT_ENV")

	assert.Equal(t, "192.168.1.5", result)
}

// Successfully handle multiple network interfaces and return correct IP
func TestGetHostAddrHandlesMultipleInterfaces(t *testing.T) {
	// Yo ho ho! Let's navigate through these multiple network seas!
	mockAddrs := []net.Addr{
		&net.IPNet{IP: net.ParseIP("::1"), Mask: net.IPMask(net.ParseIP("ffff:ffff:ffff:ffff::"))},
		&net.IPNet{IP: net.ParseIP("192.168.1.5"), Mask: net.IPv4Mask(255, 255, 255, 0)},
		&net.IPNet{IP: net.ParseIP("10.0.0.1"), Mask: net.IPv4Mask(255, 0, 0, 0)},
	}

	patch := monkey.Patch(net.InterfaceAddrs, func() ([]net.Addr, error) {
		return mockAddrs, nil
	})
	defer patch.Unpatch()

	result := GetHostAddr("NONEXISTENT_ENV")

	assert.Equal(t, "192.168.1.5", result)
}

// Return empty string when environment variable is empty and no network interfaces found
func TestGetHostAddrNoInterfacesFound(t *testing.T) {
	// Blimey! Our network treasure chest be empty!
	mockAddrs := []net.Addr{}

	patch := monkey.Patch(net.InterfaceAddrs, func() ([]net.Addr, error) {
		return mockAddrs, nil
	})
	defer patch.Unpatch()

	result := GetHostAddr("NONEXISTENT_ENV")

	assert.Equal(t, "", result)
}

// Return empty string when environment variable is empty and only loopback addresses exist
func TestGetHostAddrOnlyLoopbackExists(t *testing.T) {
	// Arrr! Nothing but localhost in these waters!
	mockAddrs := []net.Addr{
		&net.IPNet{IP: net.ParseIP("127.0.0.1"), Mask: net.IPv4Mask(255, 0, 0, 0)},
	}

	patch := monkey.Patch(net.InterfaceAddrs, func() ([]net.Addr, error) {
		return mockAddrs, nil
	})
	defer patch.Unpatch()

	result := GetHostAddr("NONEXISTENT_ENV")

	assert.Equal(t, "", result)
}

// Return empty string when net.InterfaceAddrs() returns error
func TestGetHostAddrInterfaceError(t *testing.T) {
	// Blast! Our network compass be broken!
	patch := monkey.Patch(net.InterfaceAddrs, func() ([]net.Addr, error) {
		return nil, errors.New("network error")
	})
	defer patch.Unpatch()

	result := GetHostAddr("NONEXISTENT_ENV")

	assert.Equal(t, "", result)
}

// Handle case when all IPs are IPv6 addresses
func TestGetHostAddrOnlyIPv6Exists(t *testing.T) {
	// Yarr! These newfangled IPv6 addresses be confusing the crew!
	mockAddrs := []net.Addr{
		&net.IPNet{IP: net.ParseIP("2001:db8::1"), Mask: net.IPMask(net.ParseIP("ffff:ffff:ffff:ffff::"))},
		&net.IPNet{IP: net.ParseIP("fe80::1"), Mask: net.IPMask(net.ParseIP("ffff:ffff:ffff:ffff::"))},
	}

	patch := monkey.Patch(net.InterfaceAddrs, func() ([]net.Addr, error) {
		return mockAddrs, nil
	})
	defer patch.Unpatch()

	result := GetHostAddr("NONEXISTENT_ENV")

	assert.Equal(t, "", result)
}
