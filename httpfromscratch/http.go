package httpfromscratch

import (
    "errors"
    "fmt"
    "io"
    "os"
    "strings"
)

// ReadFile reads and processes the content of a file line by line.
func ReadFile(inputFile string) error {
    f, err := os.Open(inputFile)
    if err != nil {
        return fmt.Errorf("could not open %s: %w", inputFile, err)
    }
    defer f.Close()

    currentLine := ""
    buffer := make([]byte, 8)

    for {
        n, err := f.Read(buffer)
        if err != nil {
            if currentLine != "" {
                fmt.Printf("read: %s\n", currentLine)
                currentLine = ""
            }
            if errors.Is(err, io.EOF) {
                break
            }
            return fmt.Errorf("error reading file: %w", err)
        }

        str := string(buffer[:n])
        clParts := strings.Split(str, "\n")

        // Process all parts except the last one
        for i := range(len(clParts)-1) {
            currentLine += clParts[i]
            fmt.Printf("read: %s\n", currentLine)
            currentLine = "" // Reset for the next line
        }

        // Handle the last part (incomplete line)
        currentLine += clParts[len(clParts)-1]
    }

    return nil
}