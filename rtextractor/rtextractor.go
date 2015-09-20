package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var inputPath = flag.String("i", "", "Input file")
var outputPath = flag.String("o", "", "Output file")
var digitRegex = regexp.MustCompile("\\d+")
var rtrwSeparatorRegex = regexp.MustCompile("[//\\\\-]")
var rtRegex = regexp.MustCompile("(?i)RT(\\W*\\d+)*")

const IndexLokasiRW = 9
const IndexLokasi = 10
const IndexMasalah = 11
const IndexSolusi = 12

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	flag.Parse()

	inputFile, err := os.Open(*inputPath)
	check(err)
	defer inputFile.Close()

	csvReader := csv.NewReader(inputFile)
	headers, _ := csvReader.Read()
	records, _ := csvReader.ReadAll()

	outputFile, err := os.Create(*outputPath)
	check(err)
	defer outputFile.Close()

	csvWriter := csv.NewWriter(outputFile)
	headers = append(headers[:IndexLokasiRW], append([]string{"lokasi_rt"}, headers[IndexLokasiRW:]...)...)
	csvWriter.Write(headers)

	rtCounter := 0

	for _, record := range records {
		alamatLokasi := record[IndexLokasi]
		masalah := record[IndexMasalah]
		usulanSolusi := record[IndexSolusi]

		rtCandidates := extractFromMultipleCandidates(rtRegex, alamatLokasi, masalah, usulanSolusi)

		rtCounter += len(rtCandidates)

		if len(rtCandidates) > 0 {
			fmt.Println("\nAlamat:", alamatLokasi)
			fmt.Println("Masalah:", masalah)
			fmt.Println("Solusi:", usulanSolusi)
			fmt.Println("Found RT:", rtCandidates)
		}

		rtString := strings.Join(rtCandidates, ", ")

		record = append(record[:IndexLokasiRW], append([]string{rtString}, record[IndexLokasiRW:]...)...)
		csvWriter.Write(record)
	}

	fmt.Println("Finished processing", len(records), "records, found", rtCounter, "RT")

	csvWriter.Flush()
}

func extractFromMultipleCandidates(regex *regexp.Regexp, candidateSources ...string) []string {
	for _, candidateSource := range candidateSources {
		candidates := extractCandidates(regex, candidateSource)
		if len(candidates) > 0 {
			return candidates
		}
	}

	return []string{""}
}

func extractCandidates(regex *regexp.Regexp, s string) []string {
	matches := regex.FindAllString(s, -1)
	candidates := extractDigits(matches)

	return candidates
}

func extractDigits(strings []string) []string {
	digits := make([]string, 0)

	for _, s := range strings {
		foundDigits := digitRegex.FindAllString(extractBeforeSeparator(s), -1)

		for _, foundDigit := range foundDigits {
			if !contains(digits, foundDigit) {
				digits = append(digits, foundDigit)
			}
		}
	}
	return digits
}

// we need to filter out cases where the RW is implicitly appended on the RT (e.g. RT 001/02)
func extractBeforeSeparator(s string) string {
	loc := rtrwSeparatorRegex.FindStringIndex(s)
	if loc == nil {
		return s
	} else {
		return s[0:loc[0]]
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
