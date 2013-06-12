package geoip

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"testing"
)

func BenchmarkGeoIpLookup(t *testing.B) {
	db, err := New()
	if err != nil {
		panic(err.Error())
	}
	ips := []net.IP{}
	for i := 0; i < 1000; i++ {
		ip := ""
		for j := 0; j < 4; j++ {
			ip += fmt.Sprintf("%d.", rand.Int()%255)
		}
		ip = strings.TrimRight(ip, ".")
		ips = append(ips, net.ParseIP(ip))
	}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		ip := ips[rand.Int()%len(ips)]
		db.Lookup(ip)
	}
}
