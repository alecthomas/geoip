db.go: GeoIPCountryWhois.csv precompute.py
	python precompute.py
	gobundle --compress --package=db --target=db/db.go ranges.db countries.csv
	rm -f ranges.db countries.csv
