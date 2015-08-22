package main

import (
    "os"
    "bufio"
    "bytes"
    "flag"
)

var inputFlag = flag.String("i", "", "Input file")
var outputFlag = flag.String("o", "", "Output file")

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    var buffer bytes.Buffer
    var i int = 0

    flag.Parse()

    file, err := os.Open(*inputFlag)
    check(err)
    defer file.Close()

    output, err := os.Create(*outputFlag)
    check(err)
    defer output.Close()

    scanner := bufio.NewScanner(file)
    writer := bufio.NewWriter(output)

    for scanner.Scan() {
        if buffer.Len() > 0 {
            buffer.WriteString(",")
        }

        buffer.WriteString(scanner.Text())
        i++

        if i % 20 == 0 {
            writer.WriteString(buffer.String() + "\n")
            buffer.Reset()
        }
    }

    writer.WriteString(buffer.String())
    writer.Flush()
}