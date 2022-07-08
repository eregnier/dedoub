package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type FileDefinition struct {
	Size         int64
	LastModified time.Time
	Filename     string
	Path         string
	Checksum     string
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Missing parameters. expecting </folder/to/inspect> </tmp/result.json>")
	}
	db := fileList(strings.Split(os.Args[1], ";"))
	counts := findDuplicates(db)
	duplicates := filterDuplicates(db, counts)
	writeResult(duplicates)
}

func writeResult(duplicates map[string][]FileDefinition) {
	data, err := json.MarshalIndent(duplicates, "", "    ")
	if err != nil {
		log.Fatal("unable to write result")
	}
	err = ioutil.WriteFile(os.Args[2], data, 0644)
	if err != nil {
		log.Fatal("error while writting result file")
	}
}

func filterDuplicates(db []FileDefinition, counts map[string]int) map[string][]FileDefinition {
	var filteredDb = make(map[string][]FileDefinition)
	for _, item := range db {
		if counts[item.Checksum] > 1 {
			filteredDb[item.Checksum] = append(filteredDb[item.Checksum], item)
		}
	}
	return filteredDb
}

func findDuplicates(db []FileDefinition) map[string]int {
	counts := make(map[string]int)
	for _, item := range db {
		counts[item.Checksum] += 1
	}
	return counts
}

func fileList(folders []string) []FileDefinition {
	var db []FileDefinition
	for _, folder := range folders {
		err := filepath.Walk(folder,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() {
					db = append(db, FileDefinition{
						Filename:     filepath.Base(path),
						Path:         path,
						Checksum:     fileChecksum(path),
						Size:         info.Size(),
						LastModified: info.ModTime(),
					})
				}
				return nil
			})
		if err != nil {
			log.Println(err)
		}
	}
	return db
}

func fileChecksum(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal("unable to compute file checksum > ", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
