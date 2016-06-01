// +build ignore

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
)

// use the tcell db
const dbURL = "https://raw.githubusercontent.com/gdamore/tcell/master/database.json"

type entry struct {
	Name   string
	Colors int
}

type byName []entry

func (s byName) Len() int           { return len(s) }
func (s byName) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s byName) Less(i, j int) bool { return s[i].Name < s[j].Name }

func getEntries() ([]entry, error) {
	resp, err := http.Get(dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %v", err)
	}
	defer resp.Body.Close()
	s := bufio.NewScanner(resp.Body)
	var entries []entry
	for s.Scan() {
		var e entry
		err := json.Unmarshal(s.Bytes(), &e)
		if err != nil {
			return nil, fmt.Errorf("error decoding: %v", err)
		}
		entries = append(entries, e)
	}
	return entries, s.Err()
}

func do() error {
	destFile := os.Args[1] + ".go"
	entries, err := getEntries()
	if err != nil {
		return err
	}

	sort.Sort(byName(entries))
	f, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()

	const header = `// this file was autogenerated by mkdb.go -- DO NOT EDIT

package tcolor

var db = []string{
`

	if _, err := w.Write([]byte(header)); err != nil {
		return err
	}
	for _, e := range entries {
		if e.Colors > 0 {
			if _, err := fmt.Fprintf(w, "\t%q,\n", e.Name); err != nil {
				return err
			}
		}
	}
	_, err = fmt.Fprint(w, "}\n")
	return err
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("mkdb: ")
	if err := do(); err != nil {
		log.Fatal(err)
	}
}
