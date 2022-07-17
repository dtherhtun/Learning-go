package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

const (
	inputFile  = "./testdata/test1.md"
	resultFile = "test1.md.html"
	goldenFile = "./testdata/test1.md.html"
)

func TestParseContent(t *testing.T) {
	var expected = new(bytes.Buffer)
	input, err := ioutil.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}
	result, err := paraseContent(input, "")
	if err != nil {
		t.Fatal(err)
	}

	tpl, err := template.ParseFiles(goldenFile)
	if err != nil {
		t.Fatal(err)
	}

	if err := tpl.Execute(expected, time.Now().Format("2006-01-02 15:04:05")); err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(expected.Bytes(), result) {
		t.Logf("golden:\n%s\n", expected.Bytes())
		t.Logf("result:\n%s\n", result)
		t.Error("Result content does not match golden file")
	}
}

func TestRun(t *testing.T) {
	var mockStdOut bytes.Buffer
	var expected = new(bytes.Buffer)
	if err := run(inputFile, "", &mockStdOut, true); err != nil {
		t.Fatal(err)
	}

	resultFile := strings.TrimSpace(mockStdOut.String())

	result, err := ioutil.ReadFile(resultFile)
	if err != nil {
		t.Fatal(err)
	}

	tpl, err := template.ParseFiles(goldenFile)
	if err != nil {
		t.Fatal(err)
	}

	if err := tpl.Execute(expected, time.Now().Format("2006-01-02 15:04:05")); err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(expected.Bytes(), result) {
		t.Logf("golden:\n%s\n", expected.Bytes())
		t.Logf("result:\n%s\n", result)
		t.Error("Result content does not match golden file")
	}

	os.Remove(resultFile)
}
