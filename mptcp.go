//go:build go1.21

package tfo

import (
	"net"
	_ "unsafe"
)

//go:linkname DialerMultipathTCP net.(*Dialer).MultipathTCP
func DialerMultipathTCP(d *net.Dialer) bool

//go:linkname DialerSetMultipathTCP net.(*Dialer).SetMultipathTCP
func DialerSetMultipathTCP(d *net.Dialer, use bool)
