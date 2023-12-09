//go:build !go1.21

package tfo

import "net"

func DialerMultipathTCP(d *net.Dialer) bool {
	return false
}

func DialerSetMultipathTCP(d *net.Dialer, use bool) {
}
