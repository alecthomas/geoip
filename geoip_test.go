package geoip

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"testing"
)

func TestGeoIpLookup(t *testing.T) {
	db, err := New()
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		ip      net.IP
		country Country
	}{
		{net.ParseIP("8.8.8.8"), Country{"US", "United States"}},
		{net.ParseIP("8.8.8.8").To4(), Country{"US", "United States"}},
		{net.ParseIP("8.8.8.8").To4()[:3], Country{}},
		{net.ParseIP("8.8.8.8")[:15], Country{}},
		{nil, Country{}},
	}
	for _, c := range cases {
		var country Country
		if ct := db.Lookup(c.ip); ct != nil {
			country = *ct
		}
		if country != c.country {
			t.Errorf("%s (%#v) expected %v, got %v", c.ip, c.ip, c.country, country)
		}
	}
}

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
