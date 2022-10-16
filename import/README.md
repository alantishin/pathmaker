# Osm data import utils

## Import roads
```
osm2pgrouting -f <IMPORT_FILE>.osm -h <HOST> -d <DATABASE> -W <PASSWORD> -U <USER> --chunk 5000
```