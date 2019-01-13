package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	if err := Main(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Main(args []string) error {
	switch len(args) {
	case 1:
		return Convert(os.Stdin, os.Stdout)
	case 2:
		from, err := os.Open(args[1])
		if err != nil {
			return err
		}
		defer from.Close()
		return Convert(from, os.Stdout)
	case 3:
		from, err := os.Open(args[1])
		if err != nil {
			return err
		}
		defer from.Close()
		to, err := os.Open(args[2])
		if err != nil {
			return err
		}
		defer to.Close()
		return Convert(from, to)
	default:
		return errors.New("Invalid arguments")
	}
}

var re = regexp.MustCompile(`\(注:[^)]+\)`)

func Convert(r io.Reader, w io.Writer) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	indexes := re.FindAllIndex(b, -1)
	if indexes == nil {
		_, err := w.Write(b)
		return err
	}

	beg := 0
	notes := make([][]byte, 0, len(indexes))

	for i, idx := range indexes {
		w.Write(b[beg:idx[0]])
		fmt.Fprintf(w, "[^%d]", i+1)
		notes = append(notes, b[idx[0]:idx[1]])
		beg = idx[1]
	}
	w.Write(b[beg:])
	w.Write([]byte("\n"))
	for i, note := range notes {
		fmt.Fprintf(w, "[^%d]: %s\n", i+1, note[5:len(note)-1])
	}

	return nil
}
