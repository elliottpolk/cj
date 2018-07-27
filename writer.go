package main

import (
	"fmt"
	"os"
)

type writer struct {
	*os.File
}

func (w *writer) write(data string) error {
	if _, err := fmt.Fprint(w.File, data); err != nil {
		return err
	}

	return nil
}
