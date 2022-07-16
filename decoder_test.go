package twtxt_test

import (
	"github.com/reiver/go-twtxt"

	"io"
	"strings"
	"time"

	"testing"
)

func TestDecoder_timeAndString(t *testing.T) {

	parseRFC3339Time := func(value string) time.Time {
		t, err := time.Parse(time.RFC3339, value)
		if nil != err {
			panic(err)
		}
		return t
	}

	tests := []struct {
		Input string
		Expected []struct{
			Kind twtxt.Kind
			When time.Time
			Content string
		}
	}{
		{
			Input:
				"2016-02-04T13:30:00+01:00" +"\t"+ "You can really go crazy here! ┐(ﾟ∀ﾟ)┌"                     +"\n"+
				"2016-02-03T23:05:00+01:00" +"\t"+ "@<example http://example.org/twtxt.txt> welcome to twtxt!" +"\n"+
				"2016-02-01T11:00:00+01:00" +"\t"+ "This is just another example."                             +"\n"+
				"2015-12-12T12:00:00+01:00" +"\t"+ "Fiat lux!"                                                 +"\n",
			Expected: []struct{
				Kind twtxt.Kind
				When time.Time
				Content string
			}{
				{
					Kind: twtxt.KindStatus,
					When: parseRFC3339Time("2016-02-04T13:30:00+01:00"),
					Content: "You can really go crazy here! ┐(ﾟ∀ﾟ)┌",
				},
				{
					Kind: twtxt.KindStatus,
					When: parseRFC3339Time("2016-02-03T23:05:00+01:00"),
					Content: "@<example http://example.org/twtxt.txt> welcome to twtxt!",
				},
				{
					Kind: twtxt.KindStatus,
					When: parseRFC3339Time("2016-02-01T11:00:00+01:00"),
					Content: "This is just another example.",
				},
				{
					Kind: twtxt.KindStatus,
					When: parseRFC3339Time("2015-12-12T12:00:00+01:00"),
					Content: "Fiat lux!",
				},
			},
		},



		{
			Input:
				`# | |___      _| |___  _| |_`                     +"\n"+
				`# | __\ \ /\ / / __\ \/ / __|`                    +"\n"+
				`# | |_ \ V  V /| |_ >  <| |_`                     +"\n"+
				`#  \__| \_/\_/  \__/_/\_\\__|`                    +"\n"+
				"#"                                                +"\n"+
				"# Twtxt is an open, distributed"                  +"\n"+
				"# microblogging platform that"                    +"\n"+
				"# uses human-readable text files,"                +"\n"+
				"# common transport protocols, and"                +"\n"+
				"# free software."                                 +"\n"+
				"#"                                                +"\n"+
				"# Learn more about twtxt at"                      +"\n"+
				"#   https://github.com/buckket/twtxt"             +"\n"+
				"#"                                                +"\n"+
				"# ------------------------------------------"     +"\n"+
				"#"                                                +"\n"+
				"# nick = joeblow  "                               +"\n"+
				"# url = https://example.com/feed/twtxt"           +"\n"+
				"# lang = en_CA"                                   +"\n"+
				""                                                 +"\n"+
				"2022-05-29T12:00:00+02:00" +"\t"+ "You can really go crazy here! ┐(ﾟ∀ﾟ)┌"                     +"\n"+
				"2022-05-15T18:36:34+02:00" +"\t"+ "@<example http://example.org/twtxt.txt> welcome to twtxt!" +"\n"+
				"2022-04-27T20:31:16+02:00" +"\t"+ "This is just another example."                             +"\n"+
				"2022-04-25T16:38:18+02:00" +"\t"+ "Fiat lux!"                                                 +"\n"+
				"2022-04-10T21:58:49+02:00" +"\t"+ "Fiat lux!"                                                 +"\n",
			Expected: []struct{
				Kind twtxt.Kind
				When time.Time
				Content string
			}{
				{
					Kind: twtxt.KindComment,
					Content: `# | |___      _| |___  _| |_`,
				},
				{
					Kind: twtxt.KindComment,
					Content: `# | __\ \ /\ / / __\ \/ / __|`,
				},
				{
					Kind: twtxt.KindComment,
					Content: `# | |_ \ V  V /| |_ >  <| |_`,
				},
				{
					Kind: twtxt.KindComment,
					Content: `#  \__| \_/\_/  \__/_/\_\\__|`,
				},
				{
					Kind: twtxt.KindComment,
					Content: "#",
				},
				{
					Kind: twtxt.KindComment,
					Content: "# Twtxt is an open, distributed",
				},
				{
					Kind: twtxt.KindComment,
					Content: "# microblogging platform that",
				},
				{
					Kind: twtxt.KindComment,
					Content: "# uses human-readable text files,",
				},
				{
					Kind: twtxt.KindComment,
					Content: "# common transport protocols, and",
				},
				{
					Kind: twtxt.KindComment,
					Content: "# free software.",
				},
				{
					Kind: twtxt.KindComment,
					Content: "#",
				},
				{
					Kind: twtxt.KindComment,
					Content: "# Learn more about twtxt at",
				},
				{
					Kind: twtxt.KindComment,
					Content: "#   https://github.com/buckket/twtxt",
				},
				{
					Kind: twtxt.KindComment,
					Content: "#",
				},
				{
					Kind: twtxt.KindComment,
					Content: "# ------------------------------------------",
				},
				{
					Kind: twtxt.KindComment,
					Content: "#",
				},
				{
					Kind: twtxt.KindComment,
					Content: "# nick = joeblow  ",
				},
				{
					Kind: twtxt.KindComment,
					Content: "# url = https://example.com/feed/twtxt",
				},
				{
					Kind: twtxt.KindComment,
					Content: "# lang = en_CA",
				},
				{
					Kind: twtxt.KindInvalid,
					Content: "",
				},
				{
					Kind: twtxt.KindStatus,
					When: parseRFC3339Time("2022-05-29T12:00:00+02:00"),
					Content: "You can really go crazy here! ┐(ﾟ∀ﾟ)┌",
				},
				{
					Kind: twtxt.KindStatus,
					When: parseRFC3339Time("2022-05-15T18:36:34+02:00"),
					Content: "@<example http://example.org/twtxt.txt> welcome to twtxt!",
				},
				{
					Kind: twtxt.KindStatus,
					When: parseRFC3339Time("2022-04-27T20:31:16+02:00"),
					Content: "This is just another example.",
				},
				{
					Kind: twtxt.KindStatus,
					When: parseRFC3339Time("2022-04-25T16:38:18+02:00"),
					Content: "Fiat lux!",
				},
				{
					Kind: twtxt.KindStatus,
					When: parseRFC3339Time("2022-04-10T21:58:49+02:00"),
					Content: "Fiat lux!",
				},
			},
		},
	}

	TestLoop: for testNumber, test := range tests {

		var decoder twtxt.Decoder

		{
			var reader io.Reader = strings.NewReader(test.Input)

			err := decoder.SetReader(reader)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %q", err, err)
				t.Logf("INPUT: %q", test.Input)
				continue
			}
		}

		var numIter int = 0
		for decoder.Next() {
			numIter++

			var kind twtxt.Kind
			{
				kind = decoder.Kind()

				expected := test.Expected[numIter-1].Kind
				actual   := kind

				if expected != actual {
					t.Errorf("For test #%d and iteration #%d,", testNumber, numIter-1)
					t.Logf("EXPECTED KIND: %q", expected)
					t.Logf("ACTUAL KIND: %q", actual)
					t.Errorf("EXPECTED CONTENT: %q", test.Expected[numIter-1].Content)
					continue TestLoop
				}
			}

			switch kind {
			case twtxt.KindStatus:


				var when time.Time
				var content string

				err := decoder.Decode(&when, &content)
				if nil != err {
					t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one.", testNumber, numIter-1)
					t.Logf("ERROR: (%T) %q", err, err)
					t.Logf("INPUT: %q", test.Input)
					continue TestLoop
				}

				{
					expected := test.Expected[numIter-1].When
					actual   := when

					if !expected.Equal(actual) {
						t.Errorf("For test #%d and iteration #%d, the actual 'when' is not what was expected.", testNumber, numIter-1)
						t.Logf("EXPECTED WHEN [%d]: %v", numIter-1, expected)
						t.Logf("ACTUAL   WHEN [%d]: %v", numIter-1, actual)
						continue TestLoop
					}
				}

				{
					expected := test.Expected[numIter-1].Content
					actual   := content

					if expected != actual {
						t.Errorf("For test #%d and iteration #%d, the actual 'content' is not what was expected.", testNumber, numIter-1)
						t.Logf("EXPECTED CONTENT [%d]: %q", numIter-1, expected)
						t.Logf("ACTUAL   CONTENT [%d]: %q", numIter-1, actual)
						continue TestLoop
					}
				}


			case twtxt.KindComment:


				var content string

				err := decoder.Decode(&content)
				if nil != err {
					t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one.", testNumber, numIter-1)
					t.Logf("ERROR: (%T) %q", err, err)
					t.Logf("INPUT: %q", test.Input)
					continue TestLoop
				}

				{
					expected := test.Expected[numIter-1].Content
					actual   := content

					if expected != actual {
						t.Errorf("For test #%d and iteration #%d, the actual 'comment-content' is not what was expected.", testNumber, numIter-1)
						t.Logf("EXPECTED CONTENT [%d]: %q", numIter-1, expected)
						t.Logf("ACTUAL   CONTENT [%d]: %q", numIter-1, actual)
						continue TestLoop
					}
				}


			case twtxt.KindInvalid:


				var content string

				err := decoder.Decode(&content)
				if nil != err {
					t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one.", testNumber, numIter-1)
					t.Logf("ERROR: (%T) %q", err, err)
					t.Logf("INPUT: %q", test.Input)
					continue TestLoop
				}

				{
					expected := test.Expected[numIter-1].Content
					actual   := content

					if expected != actual {
						t.Errorf("For test #%d and iteration #%d, the actual 'invalid-content' is not what was expected.", testNumber, numIter-1)
						t.Logf("EXPECTED CONTENT [%d]: %q", numIter-1, expected)
						t.Logf("ACTUAL   CONTENT [%d]: %q", numIter-1, actual)
						continue TestLoop
					}
				}


			default:
				panic(kind.String())
			}

		}
		if err := decoder.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %q", err, err)
			t.Logf("INPUT: %q", test.Input)
			continue
		}

		{
			expected := len(test.Expected)
			actual   := numIter

			if expected != actual {
				t.Errorf("For test #%d, actual number of iterations is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("VALUE:    %d", actual)
				continue
			}
		}
	}
}
