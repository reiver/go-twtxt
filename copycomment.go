package twtxt

import (
	"io"
)

func copycomment(writer io.Writer, reader io.Reader) error {
	return copyuntil(writer, reader, '\n', true)
}

