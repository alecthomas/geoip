db/db.go: GeoIPCountryWhois.csv Makefile precompute.py
	python precompute.py
	gobundle --package=db --target=db/db.go ranges.db countries.csv
	rm -f ranges.db countries.csv
