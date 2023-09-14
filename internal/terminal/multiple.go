package terminal

import "fmt"

type Multiple struct {
	*Streams
}

func NewMultiple(streams *Streams) *Multiple {
	return &Multiple{
		Streams: streams,
	}
}

// Print is a helper for conveniently calling fmt.Fprint on the Stdout stream.
func (s *Multiple) Print(a ...interface{}) (n int, err error) {
	fmt.Println(a...)
	return fmt.Fprint(s.Stdout.File, a...)
}

// Printf is a helper for conveniently calling fmt.Fprintf on the Stdout stream.
func (s *Multiple) Printf(format string, a ...interface{}) (n int, err error) {
	fmt.Println(a...)
	return fmt.Fprintf(s.Stdout.File, format, a...)
}

// Println is a helper for conveniently calling fmt.Fprintln on the Stdout stream.
func (s *Multiple) Println(a ...interface{}) (n int, err error) {
	fmt.Println(a...)
	return fmt.Fprintln(s.Stdout.File, a...)
}

// Eprint is a helper for conveniently calling fmt.Fprint on the Stderr stream.
func (s *Multiple) Eprint(a ...interface{}) (n int, err error) {
	fmt.Println(a...)
	return fmt.Fprint(s.Stderr.File, a...)
}

// Eprintf is a helper for conveniently calling fmt.Fprintf on the Stderr stream.
func (s *Multiple) Eprintf(format string, a ...interface{}) (n int, err error) {
	fmt.Println(a...)
	return fmt.Fprintf(s.Stderr.File, format, a...)
}

// Eprintln is a helper for conveniently calling fmt.Fprintln on the Stderr stream.
func (s *Multiple) Eprintln(a ...interface{}) (n int, err error) {
	fmt.Println(a...)
	return fmt.Fprintln(s.Stderr.File, a...)
}
