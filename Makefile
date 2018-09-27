all: generate
	go build
	go test -timeout 5s
	cd rdt && go test -timeout 5s

generate:
	node ./transclude.js sep2bits.txt *.go
	ragel -G2 -e -L -Z dfa.rl -o parser.go
	go fmt