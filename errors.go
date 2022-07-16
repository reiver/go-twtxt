package twtxt

import (
	"github.com/reiver/go-fck"
)

const (
	errInvalidTwtxt   = fck.Error("invalid twtxt")
	errNilReader      = fck.Error("nil reader")
	errNilReceiver    = fck.Error("nil receiver")
	errReaderFound    = fck.Error("reader found")
	errReaderNotFound = fck.Error("reader not found")
	errRuneError      = fck.Error("rune error")
)
