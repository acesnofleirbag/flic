package main

import (
	"bytes"
	"flag"
	"flic/guard"
	"fmt"
	"io"
	"os"
)

var (
    path string
    recursive bool
)

func Process(path string) {
    files, err := os.ReadDir(path)
    guard.Err(err)

    for _, item := range files {
        if item.IsDir() && !recursive {
            continue
        }

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
    flag.StringVar(&path, "path", ".", "path to execute line counting")
    flag.BoolVar(&recursive, "recursive", true, "recursive line counting")

    flag.Parse()

    Process(path)
}
