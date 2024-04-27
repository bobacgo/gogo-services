package network

import (
	"errors"
	"net"
	"strings"
)

func OutBoundIPV1() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip := strings.Split(localAddr.String(), ":")[0]
	return ip, err
}

func IsValidIP(addr string) bool {
	ip := net.ParseIP(addr)
	return ip.IsGlobalUnicast() && !ip.IsInterfaceLocalMulticast()
}

func OutBoundIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	lowest := int(^uint(0) >> 1)
	var result net.IP
	for _, iface := range ifaces {
		if (iface.Flags & net.FlagUp) == 0 {
			continue
		}
		if iface.Index < lowest || result == nil {
			lowest = iface.Index
		}
		if result != nil {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, rawAddr := range addrs {
			var ip net.IP
			switch addr := rawAddr.(type) {
			case *net.IPAddr:
				ip = addr.IP
			case *net.IPNet:
				ip = addr.IP
			default:
				continue
			}
			if IsValidIP(ip.String()) {
				result = ip
			}
		}
	}
	if result == nil {
		return "", errors.New("no valid ip found")
	}
	return result.String(), nil
}
