package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func checkExt(path string, exts Exts) bool {

	for _, ext := range exts {

		if ext != "" && ext == path {
			return false
		}
	}
	return true
}

func filterOut(path string, exts Exts, minSize int64, info os.FileInfo) bool {

	if info.IsDir() || info.Size() < minSize {
		return true
	}

	if len(exts) > 0 {
		if checkExt(filepath.Ext(path), exts) {
			return true
		}
	}

	return false
}

func filterOutWithDate(path string, exts Exts, minSize int64, dateBefore string, info os.FileInfo) bool {

	if info.IsDir() || info.Size() < minSize || info.ModTime().Format("2006-01-02") != dateBefore || dateBefore == "" {
		return true
	}

	if len(exts) > 0 {
		if checkExt(filepath.Ext(path), exts) {
			return true
		}
	}

	return false
}

func listFile(path string, out io.Writer) error {
	_, err := fmt.Fprintln(out, path)
	return err
}

func delFile(path string, delLogger *log.Logger) error {
	if err := os.Remove(path); err != nil {
		return err
	}
	delLogger.Println(path)
	return nil
}

func archiveFile(distDir, root, path string) error {
	info, err := os.Stat(distDir)
	if err != nil {
		return err
	}

	if !info.IsDir() {
		return fmt.Errorf("%s is not a directory", distDir)
	}

	relDir, err := filepath.Rel(root, filepath.Dir(path))
	if err != nil {
		return err
	}

	dest := fmt.Sprintf("%s.gz", filepath.Base(path))
	targetPath := filepath.Join(distDir, relDir, dest)
	if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
		return err
	}

	out, err := os.OpenFile(targetPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer out.Close()

	in, err := os.Open(path)
	if err != nil {
		return err
	}
	defer in.Close()

	zw := gzip.NewWriter(out)
	zw.Name = filepath.Base(path)

	if _, err := io.Copy(zw, in); err != nil {
		return err
	}

	if err := zw.Close(); err != nil {
		return err
	}

	return out.Close()
}
