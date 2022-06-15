package main

import (
	"fmt"
)

type OSError int

func (e *OSError) Error() string {
	return fmt.Sprintf("error #%d", *e)
}

func FileExists(path string) (bool, error) {
	var err error // is not initialized, meaning it has the zero value for a pointer - which is nil.
	return false, err
}

func main() {
	if _, err := FileExists("/no/such/file"); err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		fmt.Println("OK")
	}
}
