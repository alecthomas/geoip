# A pure Go interface to the free [MaxMind GeoIP](http://dev.maxmind.com/geoip/legacy/downloadable) database

This implementation compiles an optimized version of the database into Go (using [gobundle](http://github.com/alecthomas/gobundle)), and so does not rely on external files.

### Caveats

Currently only IPv4->country lookups are supported.

### Installation

```bash
$ # Library
$ go get github.com/alecthomas/geoip
$ # Command
$ go get github.com/alecthomas/cmd/geoip
```

### Usage

Command line usage:

```bash
$ geoip 1.1.1.1
Australia (AU)
```

See [GoDoc](http://godoc.org/github.com/alecthomas/geoip) for API documentation.

```go
geo, err := geoip.New()
country := geo.Lookup(net.ParseIP("1.1.1.1"))
fmt.Printf("%s\n", country)
```


### Performance

A benchmark is included (`go test -bench='.*'`). On my 2GHz Core i7 each lookup takes around 550ns.

### Updating the database

The package can be rebuilt with an updated datbase with the following commands:

```bash
$ curl -O http://www.maxmind.com/download/geoip/database/GeoIPCountryCSV.zip
$ unzip GeoIPCountryCSV.zip
$ make
```

### Legal

This product includes GeoLite data created by MaxMind, available from [maxmind.com](http://www.maxmind.com).
