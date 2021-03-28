# Pinger
Easy command line tool. Allows you to run regular ping tests. Based on [go-ping/ping](http://github.com/go-ping/ping)

### Flags
```bash
  -c int
        sets pings per try count (default 5)
  -d duration
        sets ping interval (default 5s)
  -h string
        sets host that will be pinged. Might be IP address (default "google.com")
  -i duration
        sets pings interval in ms (default 50ms)
  -o string
        sets file to output log. Might be relative or absolute (default "stdout")
  -t duration
        sets pings timeout in ms (default 250ms)
```

### Output
Outputs ping result to specified location as TSV time series. Display time in system local time

Command line output
``` 
pinger -o ./log.txt
host: google.com
output: stdout
count: 5
duration: 5s
timeout: 250ms
interval: 50ms
```
```
###./log.txt content additionally formatted for readme
2021/03/28 13:11:52     Error   PacketsSent     PacketsRecv     PacketLoss      MinRtt          AvgRtt          MaxRtt
2021/03/28 13:11:57             5               5               0.00            18.6672ms       19.1923ms       19.3919ms
2021/03/28 13:12:02             5               5               0.00            18.9973ms       19.411741ms     20.97ms
2021/03/28 13:12:07             5               5               0.00            18.9984ms       19.20164ms      19.9987ms
2021/03/28 13:12:12             5               5               0.00            19.0982ms       19.39886ms      19.9996ms
2021/03/28 13:12:17             5               5               0.00            19.0356ms       19.5405ms       19.9735ms
2021/03/28 13:12:22             5               5               0.00            19.0489ms       19.76836ms      20.118ms

```