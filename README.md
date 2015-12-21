# huffman

Provides basic huffman decoder implemenation.

See the [documentation](https://godoc.org/github.com/32bitkid/huffman).

## Installation

```bash
$ go get github.com/32bitkid/huffman
```

## Examples

```go
// 0xa5 => 0b 1 01 00 1 01 => True, False, Maybe, True, False

br := br.NewBitReader(bytes.NewReader([]byte{0xa5}))

decoder := huffman.NewBinaryTreeHuffmanDecoder(huffman.HuffmanTable{
    "1":  "True",
    "01": "False",
    "00": "Maybe",
})

for {
    val, err := decoder.Decode(br)
    if err != nil {
        break
    } 
    fmt.Println("%s", val)
}
```
