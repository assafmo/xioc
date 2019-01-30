package xioc

import (
	"io"
	"unicode/utf8"
)

// AccumReader Thanks to https://play.golang.org/p/1pFFrhYKXW !!
// https://groups.google.com/d/msg/golang-nuts/msMHvOmmzOY/LbGWaEqLD-YJ
type AccumReader struct {
	initial []rune
	stored  []rune
	R       io.RuneReader
}

// ReadRune to implement the io.RuneReader interface
func (a *AccumReader) ReadRune() (r rune, size int, err error) {
	if len(a.initial) > 0 {
		r = a.initial[0]
		a.initial = a.initial[1:]
		size = utf8.RuneLen(r)
	} else {
		r, size, err = a.R.ReadRune()
	}
	if err == nil {
		a.stored = append(a.stored, r)
	}
	return
}

// Slide "slides" the reader so that it will start reading
// at index i1 from the last time Slide was called,
// and returns as a string the data from i0 to i1.
func (a *AccumReader) Slide(i0, i1 int) string {
	all := string(a.initial) + string(a.stored)
	s := all[i0:i1]
	a.initial = []rune(all[i1:])
	a.stored = nil
	return s
}
