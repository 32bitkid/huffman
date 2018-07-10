// Package huffman implements a naïve Huffman decoder.
package huffman

import "errors"
import "github.com/32bitkid/bitreader"

// HuffmanDecoder is the interface that wraps the basic Decode method.
//
// Upon a match, the proper number of bits will be advanced in the bitstream,
// and the corresponding value will be returned.
type HuffmanDecoder interface {
	Decode(bitreader.BitReader) (interface{}, error)
}

// HuffmanTable is the configuration state given to a decoder.
type HuffmanTable map[string]interface{}

// ErrMissingHuffmanValue indiciates the decoded value could
// not be found in the huffman tree.
var ErrMissingHuffmanValue = errors.New("huffman: value not found")

// NewHuffmanDecoder creates a new huffman decoder with the default implementation.
func NewHuffmanDecoder(t HuffmanTable) HuffmanDecoder {
	return NewBinaryTreeHuffmanDecoder(t)
}
