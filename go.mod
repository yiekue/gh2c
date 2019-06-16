module github.com/yiekue/gh2c

go 1.12

require (
	golang.org/x/net v0.0.0-20180821023952-922f4815f713
	golang.org/x/text v0.3.0 // indirect
)

replace (
	golang.org/x/net v0.0.0-20180821023952-922f4815f713 => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
)
