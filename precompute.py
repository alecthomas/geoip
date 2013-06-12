import csv
import struct

"""Precompute more efficient data structures for lookups."""

Ranges = struct.Struct('!LL2s')


all_countries = {}

with open('GeoIPCountryWhois.csv') as input, open('ranges.db', 'w') as ranges:
    for row in csv.reader(input):
        all_countries[row[-2]] = row[-1]
        start, end = map(long, row[2:4])
        range = Ranges.pack(start, end, row[4])
        ranges.write(range)

with open('countries.csv', 'w') as countries:
    countries_writer = csv.writer(countries, quoting=csv.QUOTE_ALL)
    for c in sorted(all_countries.items()):
        countries_writer.writerow(c)
