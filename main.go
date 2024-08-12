package main

import (
	"bytes"
	"flic/guard"
	"fmt"
	"io"
	"os"
)

func Process(path string) {
    files, err := os.ReadDir(path)
    guard.Err(err)

    for _, item := range files {
        if item.IsDir() {
            Process(path + item.Name() + "/")
            continue
        }

        fileName := path + item.Name()

        file, err := os.Open(fileName)
        defer file.Close()
        guard.Err(err)

        count := 0
        buffer := make([]byte, 32 * 1024)

        for {
            c, err := file.Read(buffer)

            if err == io.EOF {
                fmt.Printf("%v: %v\n", count, fileName)
                break
            }

            guard.Err(err)

            count += bytes.Count(buffer[:c], []byte{'\n'})
        }
    }
}

func main() {
    args := os.Args

    if len(args) != 2 {
        fmt.Println("Usage: flic -path=<path> | sort -n > /tmp/flic.out")
        os.Exit(0)
    }

    path := args[1]

    Process(path)
}
