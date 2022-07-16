package twtxt

import (
	"github.com/reiver/go-cast"

	"bufio"
	"fmt"
	"io"
	"strings"
	"time"
	"unicode/utf8"
)


// Decoder reads and decodes twtxt values from an input stream.
//
// For example:
//
//	var decoder twtxt.Decoder
//	
//	err := decoder.SetReader(reader)
//	if nil != err {
//		return err
//	}
//	defer decoder.Close()
//	
//	for decoder.Next() {
//		
//		kind := decoder.Kind()
//		
//		switch kind {
//		case twtxt.KindStatus:
//			
//			var when time.Time
//			var content string
//			
//			err := decoder.Decode(&when, &content)
//			if nil != err {
//				return err
//			}
//			
//			//@TODO
//			
//		case twtxt.KindComment:
//			
//			var comment string
//			
//			err := decoder.Decode(&comment)
//			if nil != err {
//				return err
//			}
//			
//			//@TODO
//			
//		case twtxt.KindInvalid:
//			
//			var invalid string
//			
//			err := decoder.Decode(&invalid)
//			if nil != err {
//				return err
//			}
//			
//			//@TODO
//			
//		default:
//			//@TODO
//		}
//	}
//	if err := decoder.Err(); nil != err {
//		return err
//	}
type Decoder struct {
	err error
	reader io.Reader
	readbuffer *bufio.Reader
	nextkind Kind
}

// Close closes the decoder.
func (receiver *Decoder) Close() error {
	if nil == receiver {
		return nil
	}

	receiver.err = nil
	receiver.reader = nil
	receiver.readbuffer = nil
	receiver.nextkind = KindUndefined

	return nil
}

// Decode decodes the next twtxt line.
//
// You should call Kind() before calling Decode() to know what type and how many parameters to pass it.
func (receiver *Decoder) Decode(dst ...any) error {
	if nil == receiver {
		return errNilReceiver
	}
	if nil != receiver.err {
		return receiver.err
	}

	var err error
	{
		switch receiver.nextkind {
		case KindStatus:
			err = receiver.decodeStatus(dst...)
		case KindComment:
			err = receiver.decodeComment(dst...)
		case KindInvalid:
			err = receiver.decodeInvalid(dst...)
		default:
			err = errInvalidTwtxt
		}
	}

	{
		if nil != err {
			return receiver.errored(err)
		}

		return nil
	}
}

func (receiver *Decoder) decodeComment(dst ...any) error {
	if nil == receiver {
		return errNilReceiver
	}
	if nil != receiver.err {
		return receiver.err
	}

	{
		const expected = 1

		if length := len(dst); length < expected {
			return fmt.Errorf("actual number of destinations to decode twtxt comment into is %d but expected %d", length, expected)
		}
	}

	var readbuffer *bufio.Reader
	{
		readbuffer = receiver.readbuffer
		if nil == readbuffer {
			return errReaderNotFound
		}
	}

	{
		switch casted := dst[0].(type) {
		case io.Writer:
			return copycomment(casted, readbuffer)
		}
	}

	var comment string
	{
		var commentbuffer strings.Builder
		{
			err := copycomment(&commentbuffer, readbuffer)
			if nil != err && io.EOF != err {
				return err
			}
		}

		comment = commentbuffer.String()
	}

	{
		err := cast.SetString(dst[0], comment)
		if nil != err {
			return err
		}
	}

	return nil
}

func (receiver *Decoder) decodeInvalid(dst ...any) error {
	if nil == receiver {
		return errNilReceiver
	}
	if nil != receiver.err {
		return receiver.err
	}

	{
		const expected = 1

		if length := len(dst); length < expected {
			return fmt.Errorf("actual number of destinations to decode twtxt invalid into is %d but expected %d", length, expected)
		}
	}

	var readbuffer *bufio.Reader
	{
		readbuffer = receiver.readbuffer
		if nil == readbuffer {
			return errReaderNotFound
		}
	}

	{
		switch casted := dst[0].(type) {
		case io.Writer:
			return copyinvalid(casted, readbuffer)
		}
	}

	var data string
	{
		var databuffer strings.Builder
		{
			err := copyinvalid(&databuffer, readbuffer)
			if nil != err && io.EOF != err {
				return err
			}
		}

		data = databuffer.String()
	}

	{
		err := cast.SetString(dst[0], data)
		if nil != err {
			return err
		}
	}

	return nil
}

func (receiver *Decoder) decodeStatus(dst ...any) error {
	if nil == receiver {
		return errNilReceiver
	}
	if nil != receiver.err {
		return receiver.err
	}

	{
		const expected = 2

		if length := len(dst); length < expected {
			return fmt.Errorf("actual number of destinations to decode twtxt status into is %d but expected %d", length, expected)
		}
	}

	var readbuffer *bufio.Reader
	{
		readbuffer = receiver.readbuffer
		if nil == readbuffer {
			return errReaderNotFound
		}
	}


	var when time.Time
	{
		var whenBuffer strings.Builder
		{
			err := copytime(&whenBuffer, readbuffer)
			if nil != err {
				return err
			}
		}

		var whenString string = whenBuffer.String()

		{
			var err error

			when, err = time.Parse(time.RFC3339, whenString)
			if nil != err {
				return err
			}
		}
	}

	{
		switch casted := dst[0].(type) {
		case *time.Time:
			*casted = when
		case io.Writer:
			_, err := io.WriteString(casted, when.String())
			if nil != err {
				return err
			}
		case interface{Set(string)error}:
			err := casted.Set(when.String())
			if nil != err {
				return err
			}
		case interface{Scan(any)error}:
			err := casted.Scan(when)
			if nil != err {
				return err
			}
		case *string:
			if nil != casted {
				*casted = when.String()
			}
		case *[]byte:
			if nil != casted {
				*casted = []byte(when.String())
			}
		}
	}

	{
		switch casted := dst[1].(type) {
		case io.Writer:
			return copycontent(casted, readbuffer)
		}
	}

	var content string
	{
		var contentbuffer strings.Builder
		{
			err := copycontent(&contentbuffer, readbuffer)
			if nil != err && io.EOF != err {
				return err
			}
		}

		content = contentbuffer.String()
	}

	{
		err := cast.SetString(dst[1], content)
		if nil != err {
			return err
		}
	}

	return nil
}

func (receiver *Decoder) errored(err error) error {
	if nil == receiver {
		return err
	}

	receiver.err = err
	return err
}

// Err returns an errors that happened.
func (receiver *Decoder) Err() error {
	if nil == receiver {
		return nil
	}

	return receiver.err
}

// Kind returns the kind of the next twtxt line.
func (receiver *Decoder) Kind() Kind {
	if nil == receiver {
		return KindUndefined
	}

	return receiver.nextkind
}

// Next returns whether there is a next twtxt line or not.
func (receiver *Decoder) Next() bool {
	if nil == receiver {
		return false
	}

	if nil != receiver.err {
		return false
	}

	var readbuffer *bufio.Reader
	{
		readbuffer = receiver.readbuffer
		if nil == readbuffer {
			receiver.err = errReaderNotFound
			return false
		}
	}

	var r rune
	{
		var err error

		r, _, err = receiver.readbuffer.ReadRune()
		if io.EOF == err {
			return false
		}
		if nil != err {
			receiver.err = err
			return false
		}
		if utf8.RuneError == r {
			receiver.err = errRuneError
			return false
		}

		err = readbuffer.UnreadRune()
		if nil != err {
			receiver.err = err
			return false
		}
	}

	{
		switch r {
		case '0','1','2','3','4','5','6','7','8','9':
			receiver.nextkind = KindStatus
			return true
		case '#':
			receiver.nextkind = KindComment
			return true
		default:
			receiver.nextkind = KindInvalid
			return true
		}
	}
}

// SetReader provides the twtxt.Decoder with the io.Reader it will get input from.
//
// A twtxt.Decoder cannot have its reader reset to a new io.Reader.
// It will return an error if it is successfully called.
func (receiver *Decoder) SetReader(reader io.Reader) error {
	if nil == receiver {
		return errNilReceiver
	}
	if nil == reader {
		return errNilReader
	}

	if nil != receiver.reader {
		return errReaderFound
	}

	receiver.reader = reader

	receiver.readbuffer = bufio.NewReader(reader)

	return nil

}
