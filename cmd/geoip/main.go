package main

import (
	"fmt"
	"github.com/alecthomas/geoip"
	"net"
	"os"
)

func main() {
	db, err := geoip.New()
	if err != nil {
		panic(err.Error())
	}
	country := db.Lookup(net.ParseIP(os.Args[1]))
	fmt.Printf("%s\n", country)
}
