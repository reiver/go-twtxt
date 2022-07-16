package twtxt

import (
	"io"
	"strings"

	"testing"
)

func TestCopyuntil(t *testing.T) {

	tests := []struct{
		Input string
		Stop rune
		EOFToo bool
		Expected string
	}{
		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:     'A',
			EOFToo: false,
			Expected: "",
		},
		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:     'A',
			EOFToo: true,
			Expected: "",
		},



		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:      'B',
			EOFToo: false,
			Expected: "A",
		},
		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:      'B',
			EOFToo: true,
			Expected: "A",
		},



		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:       'C',
			EOFToo: false,
			Expected: "AB",
		},
		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:       'C',
			EOFToo: true,
			Expected: "AB",
		},



		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:        'D',
			EOFToo: false,
			Expected: "ABC",
		},
		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:        'D',
			EOFToo: true,
			Expected: "ABC",
		},



		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:         'E',
			EOFToo: false,
			Expected: "ABCD",
		},
		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:         'E',
			EOFToo: true,
			Expected: "ABCD",
		},



		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:          'F',
			EOFToo: false,
			Expected: "ABCDE",
		},
		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:          'F',
			EOFToo: true,
			Expected: "ABCDE",
		},



		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:           'G',
			EOFToo: false,
			Expected: "ABCDEF",
		},
		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:           'G',
			EOFToo: true,
			Expected: "ABCDEF",
		},



		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:            'H',
			EOFToo: false,
			Expected: "ABCDEFG",
		},
		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:            'H',
			EOFToo: true,
			Expected: "ABCDEFG",
		},



		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:             'I',
			EOFToo: false,
			Expected: "ABCDEFGH",
		},
		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:             'I',
			EOFToo: true,
			Expected: "ABCDEFGH",
		},



		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:              'J',
			EOFToo: false,
			Expected: "ABCDEFGHI",
		},
		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:              'J',
			EOFToo: true,
			Expected: "ABCDEFGHI",
		},



		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:               'K',
			EOFToo: false,
			Expected: "ABCDEFGHIJ",
		},
		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:               'K',
			EOFToo: true,
			Expected: "ABCDEFGHIJ",
		},



		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                'L',
			EOFToo: false,
			Expected: "ABCDEFGHIJK",
		},
		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                'L',
			EOFToo: true,
			Expected: "ABCDEFGHIJK",
		},



		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                 'M',
			EOFToo: false,
			Expected: "ABCDEFGHIJKL",
		},
		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                 'M',
			EOFToo: true,
			Expected: "ABCDEFGHIJKL",
		},



		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                  'N',
			EOFToo: false,
			Expected: "ABCDEFGHIJKLM",
		},
		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                  'N',
			EOFToo: true,
			Expected: "ABCDEFGHIJKLM",
		},



		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                   'O',
			EOFToo: false,
			Expected: "ABCDEFGHIJKLMN",
		},
		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                   'O',
			EOFToo: true,
			Expected: "ABCDEFGHIJKLMN",
		},



		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                    'P',
			EOFToo: false,
			Expected: "ABCDEFGHIJKLMNO",
		},
		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                    'P',
			EOFToo: true,
			Expected: "ABCDEFGHIJKLMNO",
		},



		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                     'Q',
			EOFToo: false,
			Expected: "ABCDEFGHIJKLMNOP",
		},
		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                     'Q',
			EOFToo: true,
			Expected: "ABCDEFGHIJKLMNOP",
		},



		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                      'R',
			EOFToo: false,
			Expected: "ABCDEFGHIJKLMNOPQ",
		},
		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                      'R',
			EOFToo: true,
			Expected: "ABCDEFGHIJKLMNOPQ",
		},



		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                       'S',
			EOFToo: false,
			Expected: "ABCDEFGHIJKLMNOPQR",
		},
		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                       'S',
			EOFToo: true,
			Expected: "ABCDEFGHIJKLMNOPQR",
		},



		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                        'T',
			EOFToo: false,
			Expected: "ABCDEFGHIJKLMNOPQRS",
		},
		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                        'T',
			EOFToo: true,
			Expected: "ABCDEFGHIJKLMNOPQRS",
		},



		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                         'U',
			EOFToo: false,
			Expected: "ABCDEFGHIJKLMNOPQRST",
		},
		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                         'U',
			EOFToo: true,
			Expected: "ABCDEFGHIJKLMNOPQRST",
		},



		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                          'V',
			EOFToo: false,
			Expected: "ABCDEFGHIJKLMNOPQRSTU",
		},
		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                          'V',
			EOFToo: true,
			Expected: "ABCDEFGHIJKLMNOPQRSTU",
		},



		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                           'W',
			EOFToo: false,
			Expected: "ABCDEFGHIJKLMNOPQRSTUV",
		},
		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                           'W',
			EOFToo: true,
			Expected: "ABCDEFGHIJKLMNOPQRSTUV",
		},



		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                            'X',
			EOFToo: false,
			Expected: "ABCDEFGHIJKLMNOPQRSTUVW",
		},
		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                            'X',
			EOFToo: true,
			Expected: "ABCDEFGHIJKLMNOPQRSTUVW",
		},



		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                             'Y',
			EOFToo: false,
			Expected: "ABCDEFGHIJKLMNOPQRSTUVWX",
		},
		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                             'Y',
			EOFToo: true,
			Expected: "ABCDEFGHIJKLMNOPQRSTUVWX",
		},



		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                              'Z',
			EOFToo: false,
			Expected: "ABCDEFGHIJKLMNOPQRSTUVWXY",
		},
		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                              'Z',
			EOFToo: true,
			Expected: "ABCDEFGHIJKLMNOPQRSTUVWXY",
		},



		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                               '[',
			EOFToo: true,
			Expected: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		},









		{
			Input:
				"# nickname = joeblow"                                                                         +"\n"+
				"# key = value"                                                                                +"\n"+
				""                                                                                             +"\n"+
				"2016-02-04T13:30:00+01:00" +"\t"+ "You can really go crazy here! ┐(ﾟ∀ﾟ)┌"                     +"\n"+
				"2016-02-03T23:05:00+01:00" +"\t"+ "@<example http://example.org/twtxt.txt> welcome to twtxt!",
			Stop: '\n',
			EOFToo: false,
			Expected: "# nickname = joeblow",
		},
		{
			Input:
				"# nickname = joeblow"                                                                         +"\n"+
				"# key = value"                                                                                +"\n"+
				""                                                                                             +"\n"+
				"2016-02-04T13:30:00+01:00" +"\t"+ "You can really go crazy here! ┐(ﾟ∀ﾟ)┌"                     +"\n"+
				"2016-02-03T23:05:00+01:00" +"\t"+ "@<example http://example.org/twtxt.txt> welcome to twtxt!",
			Stop: '\n',
			EOFToo: true,
			Expected: "# nickname = joeblow",
		},
	}

	for testNumber, test := range tests {
		var buffer strings.Builder
		var writer io.Writer = &buffer

		var reader io.Reader = strings.NewReader(test.Input)

		err := copyuntil(writer, reader, test.Stop, test.EOFToo)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("INPUT: %q", test.Input)
			t.Logf("STOP: %q", string(test.Stop))
			t.Logf("EOF-TOO: %t", test.EOFToo)
			t.Logf("EXPECTED: %q", test.Expected)
			continue
		}

		{
			expected := test.Expected
			actual   := buffer.String()

			if expected != actual {
				t.Errorf("For test #%d, the actual value that was written is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("INPUT: %q", test.Input)
				t.Logf("STOP: %q", string(test.Stop))
				t.Logf("EOF-TOO: %t", test.EOFToo)
				continue
			}
		}
	}
}

func TestCopyuntil_error(t *testing.T) {

	tests := []struct{
		Input string
		Stop rune
		EOFToo bool
		Expected string
	}{
		{
			Input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Stop:                               '[',
			EOFToo: false,
			Expected: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		},


		{
			Input:    "۱۲۳ 123 ABC 🙂 💀👿 \t :-)",
			Stop:     'x',
			EOFToo: false,
			Expected: "۱۲۳ 123 ABC 🙂 💀👿 \t :-)",
		},
	}

	for testNumber, test := range tests {
		var buffer strings.Builder
		var writer io.Writer = &buffer

		var reader io.Reader = strings.NewReader(test.Input)

		err := copyuntil(writer, reader, test.Stop, test.EOFToo)
		if nil == err {
			t.Errorf("For test #%d, expected an error, but did not actually get one.", testNumber)
			t.Logf("INPUT: %q", test.Input)
			t.Logf("STOP: %q", string(test.Stop))
			t.Logf("EOF-TOO: %t", test.EOFToo)
			t.Logf("EXPECTED: %q", test.Expected)
			continue
		}

		{
			expected := test.Expected
			actual   := buffer.String()

			if expected != actual {
				t.Errorf("For test #%d, the actual value that was written is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("INPUT: %q", test.Input)
				t.Logf("STOP: %q", string(test.Stop))
				t.Logf("EOF-TOO: %t", test.EOFToo)
			}
		}
	}
}
