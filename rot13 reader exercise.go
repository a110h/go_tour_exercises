package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (p rot13Reader) read_rot13(b byte) byte {
	
	alph := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	reversedAlph := "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"

	indx := strings.Index(alph, string(b))
	fmt.Sprintf("indx: %d", indx)
	
	if indx == -1 {
		return b	
	}
	
	return reversedAlph[indx]
}

func (p rot13Reader) Read(b []byte) (int, error) {		
	n, err := p.r.Read(b)
	if err != nil {
		return n, err
	}

	for i := 0; i < len(b); i++ {
		b[i] = p.read_rot13(b[i])
	}
	return n, err
}

func main() {
	
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
