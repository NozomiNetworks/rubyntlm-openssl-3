## scramble_md4_rc4

scramble_md4_rc4 feeds md4 and rc4 with data and keys derived sequentially from file. It outputs the resulting hash followed by the file name for each file specified on the command line.

It is implemented in Ruby using the custom implementations included in this gem, as well as in Go using its built in implementations. This can be used to verify that the custom Ruby implementations give the same result as the Go ones, which are assumed correct.

### Usage

```
./scramble_md4_rc4.rb <file1> [file2] ... > ruby.res
go run scramble_md4_rc4.go <file1> [file2] ... > go.res
diff ruby.res go.res
```
`diff` should result in no output, showing that the implementations are equivalent.
