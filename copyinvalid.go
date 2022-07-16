package twtxt

import (
	"io"
)

func copyinvalid(writer io.Writer, reader io.Reader) error {
	return copyuntil(writer, reader, '\n', true)
}
