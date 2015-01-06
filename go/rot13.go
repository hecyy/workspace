package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(p []byte) (i int, err error) {
	i, err = r.r.Read(p)
	if err != nil {
		return -1, err
	}
	for i := range p {
		p[i] = rot13(p[i])
	}
	return
}

func rot13(c byte) byte {
	switch {
	case c >= byte('A') && c < byte('N'), c >= byte('a') && c < byte('n'):
		return c + 13
	case c >= byte('N') && c <= byte('Z'), c >= byte('n') && c <= byte('z'):
		return c - 13
	}
	// Default.
	return c
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
