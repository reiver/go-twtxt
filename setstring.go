package twtxt

import (
	"fmt"
)

func setstring(dst any, src string) error {

	switch casted := dst.(type) {
	case interface{Set(string)error}:
		err := casted.Set(src)
		if nil != err {
			return err
		}
	case interface{Scan(any)error}:
		err := casted.Scan(src)
		if nil != err {
			return err
		}
	case *string:
		if nil != casted {
			*casted = src
		}
	case *[]byte:
		if nil != casted {
			*casted = []byte(src)
		}
	default:
		return fmt.Errorf("cannot set-string %q into something of type %T", src, dst)
	}

	return nil
}
