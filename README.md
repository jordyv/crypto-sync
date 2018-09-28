# Crypto Sync #

## Usage ##

```bash
Usage: crypto-sync [<flags>] <coin code> <coin code>
  -host string
        Host for InfluxDB (default "localhost")
  -port int
        Port for InfluxDB (default 6086)
  -time duration
        Time between syncs, eg 10s or 5m (default 10s)
```

```bash
./crypto-sync -time 5m lsk btk
```

## Run as background process ##
```bash
# Run every 5 minutes and output to syslog
./crypto-sync -time 5m lsk btk | logger -t crypto-sync &
```