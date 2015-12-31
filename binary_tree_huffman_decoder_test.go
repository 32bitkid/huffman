package huffman_test

import t "testing"
import "bytes"
import br "github.com/32bitkid/bitreader"
import "github.com/32bitkid/huffman"
import "fmt"

func ExampleHuffmanDecoder_simple() {
	table := huffman.HuffmanTable{
		"1":  true,
		"01": false,
		"00": "maybe",
	}
	ternaryDecoder := huffman.NewBinaryTreeHuffmanDecoder(table)

	r := br.NewBitReader(bytes.NewReader([]byte{0xa0}))

	val, _ := ternaryDecoder.Decode(r)
	fmt.Println(val)
	val, _ = ternaryDecoder.Decode(r)
	fmt.Println(val)
	val, _ = ternaryDecoder.Decode(r)
	fmt.Println(val)

	// output:
	// true
	// false
	// maybe
}

func Example() {
	// representation huffman tree generated from the
	// text "this is an example of a huffman tree"

	table := huffman.HuffmanTable{
		"111":   ' ',
		"010":   'a',
		"000":   'e',
		"1101":  'f',
		"1010":  'h',
		"1000":  'i',
		"0111":  'm',
		"0010":  'n',
		"1011":  's',
		"0110":  't',
		"11001": 'l',
		"00110": 'o',
		"10011": 'p',
		"11000": 'r',
		"00111": 'u',
		"10010": 'x',
	}

	decoder := huffman.NewBinaryTreeHuffmanDecoder(table)

	// the text "hello huffman" encoded using the table
	//
	// Text:       h    e   l     l     o     _   h    u     f    f    m    a   n
	// encoded: 0b 1010 000 11001 11001 00110 111 1010 00111 1101 1101 0111 010 0010
	// aligned: 0b 1010 0001 1001 1100 1001 1011 1101 0001 1111 0111 0101 1101 0001 0000
	// hex:     0x    a    1    9    c    9    b    d    1    f    7    5    d    1    0

	data := []byte{0xa1, 0x9c, 0x9b, 0xd1, 0xf7, 0x5d, 0x10}
	r := br.NewBitReader(bytes.NewReader(data))

	for i := 0; i < 13; i++ {
		val, _ := decoder.Decode(r)
		fmt.Printf("%c", val)
	}

	fmt.Println()

	// Output:
	// hello huffman
}

func TestBTHD_Simple(t *t.T) {

	// 0xa5 => 0b 1 01 00 1 01 => True, False, Maybe, True, False

	br := br.NewBitReader(bytes.NewReader([]byte{0xa5}))

	init := huffman.HuffmanTable{
		"1":  "True",
		"01": "False",
		"00": "Maybe",
	}

	hd := huffman.NewBinaryTreeHuffmanDecoder(init)

	actual, err := hd.Decode(br)
	if err != nil {
		t.Error(err)
	} else if expected := "True"; actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	actual, err = hd.Decode(br)
	if err != nil {
		t.Error(err)
	} else if expected := "False"; actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	actual, err = hd.Decode(br)
	if err != nil {
		t.Error(err)
	} else if expected := "Maybe"; actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	actual, err = hd.Decode(br)
	if err != nil {
		t.Error(err)
	} else if expected := "True"; actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	actual, err = hd.Decode(br)
	if err != nil {
		t.Error(err)
	} else if expected := "False"; actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

}

func TestBTHD_Formatted(t *t.T) {

	// 0xa5 => 0b 1 01 00 1 01 => True, False, Maybe, True, False

	br := br.NewBitReader(bytes.NewReader([]byte{0xa5}))

	init := huffman.HuffmanTable{
		"1":   "True",
		"0 1": "False",
		"0 0": "Maybe",
	}

	hd := huffman.NewBinaryTreeHuffmanDecoder(init)

	actual, err := hd.Decode(br)
	if err != nil {
		t.Error(err)
	} else if expected := "True"; actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	actual, err = hd.Decode(br)
	if err != nil {
		t.Error(err)
	} else if expected := "False"; actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	actual, err = hd.Decode(br)
	if err != nil {
		t.Error(err)
	} else if expected := "Maybe"; actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	actual, err = hd.Decode(br)
	if err != nil {
		t.Error(err)
	} else if expected := "True"; actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	actual, err = hd.Decode(br)
	if err != nil {
		t.Error(err)
	} else if expected := "False"; actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

}
