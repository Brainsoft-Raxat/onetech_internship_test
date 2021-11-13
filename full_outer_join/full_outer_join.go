package full_outer_join

import (
	"bufio"
	"log"
	"os"
	"sort"
)

func FullOuterJoin(f1Path, f2Path, resultPath string) {
	f1, err := os.Open(f1Path)
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()

	f2, err := os.Open(f2Path)
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	mp1 := Scan(f1)
	mp2 := Scan(f2)

	resultLines := make([]string, 0)

	for key, _ := range mp1 {
		if _, ok := mp2[key]; ok {
			delete(mp2, key)
		} else {
			resultLines = append(resultLines, key)
		}
	}

	for key, _ := range mp2 {
		resultLines = append(resultLines, key)
	}

	sort.Strings(resultLines)

	resFile, err := os.Create(resultPath)
	if err != nil {
		log.Fatal(err)
	}
	defer resFile.Close()

	for i, str := range resultLines {
		b := []byte(str)
		if i == len(resultLines)-1 {
			_, err = resFile.Write(b)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			b = append(b, byte('\n'))
			_, err = resFile.Write(b)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func Scan(file *os.File) map[string]int {
	mp := make(map[string]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		mp[scanner.Text()]++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil
	}

	return mp
}
