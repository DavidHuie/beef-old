# beef

A bloom filter server in go.

## Commands
Beef defines the following four commands:

```
create KEY
# Returns 1

delete KEY
# Returns 1

insert KEY VALUE
# Returns 1

check KEY VALUE
# Returns 1 or 0

write KEY
# Persists the bloom filter to disk
```