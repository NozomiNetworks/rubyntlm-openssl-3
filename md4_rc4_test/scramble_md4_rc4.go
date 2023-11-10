package main

import (
	"bufio"
	rc4_impl "crypto/rc4"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	md4_impl "golang.org/x/crypto/md4"
)

func main() {
	for _, name := range os.Args[1:] {
		fmt.Printf("%s  %s\n", hex.EncodeToString(process(name)), name)
	}
}

func process(name string) (result []byte) {
	defer func() {
		if e := recover(); e != nil && e != io.EOF {
			panic(e)
		}
	}()
	file, err := os.Open(name)
	check(err)
	defer file.Close()
	r := bufio.NewReader(file)
	result = md4(nil)
	for {
		sizes := readAll(r, 2)
		key := readAll(r, int(sizes[0])+1)
		value := readAll(r, int(sizes[1])+1)
		result = md4(append(rc4(key, value), result...))
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readAll(r io.Reader, size int) []byte {
	data, err := io.ReadAll(io.LimitReader(r, int64(size)))
	check(err)
	if len(data) < size {
		panic(io.EOF)
	}
	return data
}

func md4(data []byte) []byte {
	hash := md4_impl.New()
	hash.Write(data)
	return hash.Sum(nil)
}

func rc4(key, data []byte) []byte {
	cypher, err := rc4_impl.NewCipher(key)
	check(err)
	result := make([]byte, len(data))
	cypher.XORKeyStream(result, data)
	return result
}
