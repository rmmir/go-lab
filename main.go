package main

import (
    "fmt"
    "log"

    http "github.com/rmmir/go-lab/httpfromscratch"
)

const inputFile = "messages.txt"

func main() {
    fmt.Printf("Reading data from %s\n", inputFile)
    fmt.Printf("===============================\n")

    err := http.ReadFile(inputFile)
    if err != nil {
        log.Fatalf("Error: %s\n", err)
    }
}