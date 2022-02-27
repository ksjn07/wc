# wc
command line application for word count and line count for a given file or stdin input from command line
install by using `go build -o <YOUR_COMMAND_NAME>` or just by running `go install` 
make sure bunary generated from go build is mvoed to $PATH directories and $GOPATH/bin/ is also
added to $PATH.

## word count
`echo "some text with words" | ./wc `

above command should give number of words in the stdin

## line count
`echo "cehcking number of lines \n in this page"` | ./wc -l

above command should give number of lines in the given text from stdin

## character count
`echo main.go | ./wc -c \{`

above command should give total number { present in main.go

## input from a file
`./wc -c \{ -l -f main.go`

above command should take text content from a file instead of stdin
