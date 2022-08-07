package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type config struct {
	exts    Exts
	size    int64
	list    bool
	del     bool
	wlog    io.Writer
	archive string
	date    string
}

type Exts []string

func (arr *Exts) String() string { return "" }
func (arr *Exts) Set(value string) error {
	*arr = append(*arr, strings.TrimSpace(value))
	return nil
}

func main() {
	var exts Exts
	root := flag.String("root", ".", "Root directory to Start")
	logFile := flag.String("log", "", "Log deletes to this file")
	list := flag.Bool("list", false, "List files only")
	archive := flag.String("archive", "", "Achive directory")
	del := flag.Bool("del", false, "Delete files")
	flag.Var(&exts, "ext", "-ext .log -ext .sh ...")
	size := flag.Int64("size", 0, "Minimum file size")
	date := flag.String("date", "", "Search with Date")
	flag.Parse()

	var (
		f   = os.Stdout
		err error
	)

	if *logFile != "" {
		f, err = os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()
	}
	c := config{
		exts:    exts,
		size:    *size,
		list:    *list,
		del:     *del,
		wlog:    f,
		archive: *archive,
		date:    *date,
	}
	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(root string, out io.Writer, cfg config) error {
	delLogger := log.New(cfg.wlog, "DELETE FILE: ", log.LstdFlags)
	return filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if cfg.date != "" {
				if filterOutWithDate(path, cfg.exts, cfg.size, cfg.date, info) {
					return nil
				}
			} else {
				if filterOut(path, cfg.exts, cfg.size, info) {
					return nil
				}
			}

			if cfg.list {
				return listFile(path, out)
			}
			if cfg.archive != "" {
				if err := archiveFile(cfg.archive, root, path); err != nil {
					return err
				}
			}
			if cfg.del {
				return delFile(path, delLogger)
			}
			return listFile(path, out)
		})
}
