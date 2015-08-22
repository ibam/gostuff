package main

import (
    "os"
    "bufio"
    "bytes"
    "flag"
)

var inputPath = flag.String("i", "", "Input file")
var outputPath = flag.String("o", "", "Output file")
var separator = flag.String("s", ",", "Separator between elements in a row")
var groupedLine = flag.Int("l", 20, "Number of lines being grouped together")

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    var buffer bytes.Buffer
    var i int = 0

    flag.Parse()

    inputFile, err := os.Open(*inputPath)
    check(err)
    defer inputFile.Close()

    outputFile, err := os.Create(*outputPath)
    check(err)
    defer outputFile.Close()

    scanner := bufio.NewScanner(inputFile)
    writer := bufio.NewWriter(outputFile)

    for scanner.Scan() {
        if buffer.Len() > 0 {
            buffer.WriteString(*separator)
        }

        buffer.WriteString(scanner.Text())
        i++

        if i % *groupedLine == 0 {
            writer.WriteString(buffer.String() + "\n")
            buffer.Reset()
        }
    }

    writer.WriteString(buffer.String())
    writer.Flush()
}