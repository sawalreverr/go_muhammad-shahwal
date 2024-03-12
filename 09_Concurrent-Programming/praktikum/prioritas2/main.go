package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wordCh := make(chan string)
	result := make(map[string]int)
	slc := [][]string{{"word", "frequencies"}}

	f, err := os.Open("readme.txt")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	r := bufio.NewScanner(f)
	r.Split(bufio.ScanWords)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for r.Scan() {
			word := strings.Join(SplitWord(r.Text(), ":,."), "")
			wordCh <- word
		}

		close(wordCh)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		csvFile, err := os.Create("result.csv")
		if err != nil {
			log.Println(err)
		}
		defer csvFile.Close()

		wr := csv.NewWriter(csvFile)

		for word := range wordCh {
			result[word] += 1
		}

		for k, v := range result {
			slc = append(slc, []string{k, strconv.Itoa(v)})
		}

		for _, s := range slc {
			_ = wr.Write(s)
		}
		wr.Flush()
	}()

	wg.Wait()
	fmt.Println("done!")
}

func SplitWord(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}
