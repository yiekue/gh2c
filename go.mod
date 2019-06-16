module github.com/yiekue/gh2c

go 1.12

require (
	golang.org/x/net v0.0.0
	golang.org/x/text v0.3.0 // indirect
)

replace (
	golang.org/x/net v0.0.0 => github.com/golang/net v0.0.0-20190206173232-65e2d4e15006
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
)
