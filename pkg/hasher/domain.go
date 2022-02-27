package hasher

import "fmt"

type Output struct {
	URL  string
	Sum  []byte
	Size int
}

func (out Output) String() string {
	return fmt.Sprintf("%s %x", out.URL, out.Sum[:out.Size])
}
