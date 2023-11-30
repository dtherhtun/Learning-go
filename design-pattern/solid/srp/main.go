package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

/*
	A type should have one primary responsibility and
	as a result, it should have one reason to change that reason being
	somehow related to its primary responsibility
*/

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++ // keeping entry count
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry) // managing entries
	return entryCount
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// separation of concern

func (j *Journal) Save(filename string) {
	_ = os.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {

}

func (j *Journal) LoadFromWeb(url *url.URL) {

}

var LineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename, []byte(strings.Join(j.entries, LineSeparator)), 0644)
}

type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("I ate a bug")
	j.AddEntry("I play ingress")
	fmt.Println(j.String())

	SaveToFile(&j, "journal.txt")

	p := Persistence{"\r\n"}
	p.SaveToFile(&j, "journal.txt")
}
