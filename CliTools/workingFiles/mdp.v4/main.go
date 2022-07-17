package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	defaultTemplate = `<!DOCTYPE html>
	<html>
		<head>
			<meta http-equiv="constent-type" content="text/html; charset=utf-8">
			<title>{{ .Title }} </title>
		</head>
		<body>
		<h1>{{ .CreatedAt }}</h1>
		{{ .Body }}
		</body>
	</html>
	`
)

type content struct {
	Title     string
	CreatedAt string
	Body      template.HTML
}

func main() {
	ch := make(chan string)

	file := flag.Bool("file", false, "Markdown file to preview")
	skipPreview := flag.Bool("s", false, "Skip auto-preview")
	tFname := flag.String("t", "", "Alternate template name")
	flag.Parse()
	var input []byte

	if os.Getenv("MD_FILE") != "" {
		file, err := ioutil.ReadFile(os.Getenv("MD_FILE"))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		input = file
	}

	if *file {
		if file := flag.Arg(0); file != "" {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			input = data
		} else {
			go getMD(os.Stdin, ch)
			data := <-ch
			if data == "" && len(flag.Args()) < 1 {
				fmt.Fprintln(os.Stderr, "file name can not be blank")
				os.Exit(1)
			}
			input = []byte(data)
		}
	}
	if len(input) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(input, *tFname, os.Stdout, *skipPreview); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(input []byte, tFname string, out io.Writer, skipPreview bool) error {

	htmlData, err := paraseContent(input, tFname)
	if err != nil {
		return err
	}

	temp, err := ioutil.TempFile("", "mdp*.html")
	if err != nil {
		return err
	}

	if err := temp.Close(); err != nil {
		return err
	}
	outName := temp.Name()
	fmt.Fprintln(out, outName)

	if err := saveHTML(outName, htmlData); err != nil {
		return err
	}

	if skipPreview {
		return nil
	}
	defer os.Remove(outName)
	return preview(outName)
}

func paraseContent(input []byte, tFname string) ([]byte, error) {
	var buffer bytes.Buffer

	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	t, err := template.New("mdp").Parse(defaultTemplate)
	if err != nil {
		return nil, err
	}

	if tFname != "" {
		t, err = template.ParseFiles(tFname)
		if err != nil {
			return nil, err
		}
	}

	c := content{
		Title:     "Markdown Preview Tool",
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		Body:      template.HTML(body),
	}

	if err := t.Execute(&buffer, c); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func saveHTML(outFname string, data []byte) error {
	return ioutil.WriteFile(outFname, data, 0644)
}

func preview(fname string) error {
	cName := ""
	cParams := []string{}

	switch runtime.GOOS {
	case "linux":
		cName = "xdg-open"
	case "windows":
		cName = "cmd.exe"
		cParams = []string{"/C", "start"}
	case "darwin":
		cName = "open"
	default:
		return fmt.Errorf("OS not supported")
	}

	cParams = append(cParams, fname)
	cPath, err := exec.LookPath(cName)
	if err != nil {
		return err
	}

	err = exec.Command(cPath, cParams...).Run()

	time.Sleep(2 * time.Second)
	return err
}

func getMD(r io.Reader, ch chan string) {

	var data string
	/*buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	fmt.Println(buf.Bytes())*/

	s := bufio.NewReader(r)

	fmt.Println(s.ReadBytes('\n'))

	for {
		fmt.Println("Before readstring")
		input, err := s.ReadString('\n')
		fmt.Println("After readstring")
		if err == io.EOF {
			break
		}
		data += input
	}

	data = strings.TrimSuffix(data, "\n")

	ch <- data
}
