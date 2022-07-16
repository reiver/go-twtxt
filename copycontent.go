package twtxt

import (
	"io"
)

func copycontent(writer io.Writer, reader io.Reader) error {
	return copyuntil(writer, reader, '\n', true)
}

