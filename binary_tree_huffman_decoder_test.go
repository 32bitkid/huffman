package huffman_test

import t "testing"
import "bytes"
import br "github.com/32bitkid/bitreader"
import "github.com/32bitkid/huffman"

func ExampleNewHuffmanDecoder_simple() {
	table := huffman.HuffmanTable{
		"1":  "True",
		"01": "False",
		"00": "Maybe",
	}
	huffman.NewBinaryTreeHuffmanDecoder(table)
}

func ExampleHuffmanTable() {
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
	huffman.NewBinaryTreeHuffmanDecoder(table)
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
