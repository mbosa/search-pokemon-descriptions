package main

import (
	"context"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/mbosa/search-pokemon-descriptions/db"
)

const defaultDst = "wordCloudData.csv"

func main() {
	dst := defaultDst
	if len(os.Args) > 1 {
		dst = os.Args[1]
	}

	GenerateWordCloudData(dst)
}

func GenerateWordCloudData(dst string) {
	const minWordCount = 10

	f, err := os.OpenFile(dst, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	db := db.NewDbPool()

	wordsMap := make(map[string]int)

	rows, _ := db.Query(context.Background(), "SELECT description_tsvector FROM pokemon")
	for rows.Next() {
		var description_tsvector string

		rows.Scan(&description_tsvector)

		lexemes := strings.Split(description_tsvector, " ")
		for _, l := range lexemes {
			quotedLexeme := strings.Split(l, ":")[0]
			lexeme := quotedLexeme[1 : len(quotedLexeme)-1]

			if _, ok := wordsMap[lexeme]; !ok {
				wordsMap[lexeme] = 0
			}
			wordsMap[lexeme] += 1
		}
	}

	// order by value
	type temp struct {
		v int
		k string
	}
	s := make([]temp, 0)
	for k, v := range wordsMap {
		s = append(s, temp{v, k})
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i].v > s[j].v
	})

	for _, el := range s {
		if el.v < minWordCount {
			continue
		}

		if _, err = f.WriteString(strconv.Itoa(el.v) + ";" + el.k + "\n"); err != nil {
			panic(err)
		}
	}
}
