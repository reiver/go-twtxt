package twtxt

import (
	"github.com/reiver/go-utf8s"

	"fmt"
	"io"
)

func copyuntil(writer io.Writer, reader io.Reader, stop rune, eoftoo bool) error {
	if nil == reader {
		return errNilReader
	}

	if nil == writer {
		writer = io.Discard
	}

	Loop: for {

		var r rune
		{
			var err error
			var n int

			r, n, err = utf8s.ReadRune(reader)

			if 1 <= n && utf8s.RuneError != r && stop != r {
				m, err := utf8s.WriteRune(writer, r)
				if nil != err {
					return err
				}
				if m != n {
					return fmt.Errorf("expected to write %d but actually wrote %d", n, m)
				}
			}

			switch {
			case eoftoo && io.EOF == err:
				break Loop
			case nil != err:
				return err
			case utf8s.RuneError == r:
				return errRuneError
			}
		}

		if stop == r {
			break Loop
		}
	}

	return nil
}
