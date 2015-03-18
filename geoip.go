// A pure Go interface to the free MaxMind GeoIP database.
//
// eg.
//
//      geo, err := geoip.New()
//      country := geo.Lookup(net.ParseIP("1.1.1.1"))
//      fmt.Printf("%s\n", country)
package geoip

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"sort"
	"unsafe"

	"github.com/alecthomas/geoip/db"
)

type ipRange struct {
	start, end [4]byte
	country    [2]byte
}

type Country struct {
	// ISO 3166-1 short country code.
	Short string
	// Full country name.
	Long string
}

func (c *Country) String() string {
	return fmt.Sprintf("%s (%s)", c.Long, c.Short)
}

type GeoIP struct {
	ranges    []*ipRange
	countries map[string]*Country
}

func New() (*GeoIP, error) {
	r, err := db.DbBundle.Open("ranges.db")
	if err != nil {
		return nil, err
	}
	rb, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	rangesn := len(rb) / 10
	ranges := make([]*ipRange, rangesn)

	for i := 0; i < rangesn; i++ {
		ranges[i] = (*ipRange)(unsafe.Pointer(&rb[i*10]))
	}

	// Load countries
	c, err := db.DbBundle.Open("countries.csv")
	if err != nil {
		return nil, err
	}
	countries := map[string]*Country{}
	cc := csv.NewReader(c)
	for {
		row, err := cc.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		countries[row[0]] = &Country{Short: row[0], Long: row[1]}
	}

	return &GeoIP{
		countries: countries,
		ranges:    ranges,
	}, nil
}

// Find country of IP.
func (g *GeoIP) Lookup(ip net.IP) *Country {
	bip := []byte(ip.To4())
	i := sort.Search(len(g.ranges), func(i int) bool {
		return bytes.Compare(g.ranges[i].start[:], bip) > 0
	})
	if i > 0 {
		i--
	}
	r := g.ranges[i]
	if bytes.Compare(bip, r.start[:]) >= 0 && bytes.Compare(bip, r.end[:]) <= 0 {
		return g.countries[string(r.country[:])]
	}
	return nil
}
