package twtxt

import (
	"io"
)

func copytime(writer io.Writer, reader io.Reader) error {
	return copyuntil(writer, reader, '\t', false)
}
