# Thorney Distance

This is a CLI program computing the distance between two given UK postcodes.
The actual data is provided by a sqlite database and comes from [Gibbs's github repository](https://github.com/Gibbs/uk-postcodes).
The distance is computed using the [Haversine formula](https://en.wikipedia.org/wiki/Haversine_formula).

## Why The Name?

Just a play-on-words between the [Manhattan distance](https://en.wikipedia.org/wiki/Taxicab_geometry) and [Thorney Island](https://en.wikipedia.org/wiki/Thorney_Island_(London)).

## Why SQlite?

The original data is stored in a CSV file. I chose to use SQlite because it's a
very lightweight database using a B-tree. Hence, a whole lot faster than a
simple file.

## Why the Haversine?

I think it's a good balance between precision and complexity. Even though the
Earth is not an exact sphere, this approximation is good enough distances within
the UK (and outside, but no overseas territory postcode is supported).

I could have used Google's API, but:
- I don't want to pay
- Calling Google from a CLI program is a bit silly

## How to Build the Database from the CSV File?

```
CREATE TABLE postcode(
  code TEXT PRIMARY KEY,
  lat REAL,
  long REAL
);
CREATE INDEX postcode_code_idx on postcode(code);
.mode csv
.import postcodes.csv postcode
```
