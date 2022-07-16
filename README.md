# go-twtxt

Package **twtxt** implements encoding and decoding of the twtxt data format.

A **twtxt** file looks something like this:
```
2016-02-04T13:30:00+01:00	You can really go crazy here! ┐(ﾟ∀ﾟ)┌
2016-02-03T23:05:00+01:00	@<example http://example.org/twtxt.txt> welcome to twtxt!
2016-02-01T11:00:00+01:00	This is just another example.
2015-12-12T12:00:00+01:00	Fiat lux!
```

And this:
```
# | |___      _| |___  _| |_
# | __\ \ /\ / / __\ \/ / __|
# | |_ \ V  V /| |_ >  <| |_
#  \__| \_/\_/  \__/_/\_\\__|
#
# This is a comment.

2022-07-16T14:15:55-07:00	Dorud
2021-01-14T08:43:00+01:00	I can count! — ۰ ۱ ۲ ۳ ۴ ۵ ۶ ۷ ۸ ۹
2020-12-30T23:06:32+01:00	This is an example post 😈
2020-12-23T19:54:53+01:00	Hello world!
```

**twtxt** is a data format used for creating feeds.

**twtxt** feeds are used to form a decentralized micro-blogging social network.

**twtxt** is somewhat similar to the historic feed data formats: CDF, RSS, and Atom;
in that all 4 are programmer-readable & machine-readable feed formats based on text.

All 4 are also similar to how the even older human-readable finger .plan files were used.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-twtxt

[![GoDoc](https://godoc.org/github.com/reiver/go-twtxt?status.svg)](https://godoc.org/github.com/reiver/go-twtxt)
