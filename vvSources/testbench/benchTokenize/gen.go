package nlp

//go:generate go test -bench . -run NONE

// profiling
//go:generate go test -bench . -run NONE -cpuprofile=cpu.prof

//go:generate go tool pprof -http=:8080 cpu.prof

// Clean
//go:generate rm -rf cpu.prof benchTokenize.test
