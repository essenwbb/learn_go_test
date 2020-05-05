```shell script
go test -bench . -cpuprofile cpu.out
go tool pprof cpu.out
```