package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println("Number of CPUs:", runtime.NumCPU())

    fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0)) // If GOMAXPROCS  variable is not set, Go populates this variable by determining how many CPUs your system has by querying the operating system
}
