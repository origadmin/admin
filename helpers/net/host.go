/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package net implements the functions, types, and interfaces for the module.
package net

import (
	"net"
	"os"
)

func GetHostAddr(env string) string {
	// Gets the IP address from the environment variable
	ip := os.Getenv(env)
	if ip != "" {
		return ip
	}

	// Obtain the LAN IP address
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}

	return ""
}
