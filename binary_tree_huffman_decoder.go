package huffman

import "regexp"
import "github.com/32bitkid/bitreader"

var cleaner = regexp.MustCompile(`[^01]`)

// NewBinaryTreeHuffmanDecoder creates a Huffamn decoder that will walk
// the huffman tree one bit at a time until a match is found.
func NewBinaryTreeHuffmanDecoder(t HuffmanTable) *binaryTreeHuffmanDecoder {
	root, depth := parseInitIntoTree(t)
	return &binaryTreeHuffmanDecoder{root, depth}
}

type binaryTreeHuffmanDecoder struct {
	root  *binaryHuffmanNode
	depth uint
}

type binaryHuffmanNode struct {
	left  interface{}
	right interface{}
}

func parseInitIntoTree(init HuffmanTable) (*binaryHuffmanNode, uint) {
	root := &binaryHuffmanNode{}
	var depth uint = 0

	for bitString, value := range init {
		currentNode := root

		bitString = cleaner.ReplaceAllString(bitString, "")

		bitStringLength := len(bitString)
		if uint(bitStringLength) > depth {
			depth = uint(bitStringLength)
		}

		for i := 0; i < bitStringLength; i++ {

			bit := bitString[i : i+1]

			if i < bitStringLength-1 {
				// Descending
				if bit == "1" {
					if currentNode.left == nil {
						nextNode := &binaryHuffmanNode{}
						currentNode.left = nextNode
						currentNode = nextNode
					} else if nextNode, ok := currentNode.left.(*binaryHuffmanNode); ok {
						currentNode = nextNode
					} else {
						panic("Invalid huffman tree")
					}
				} else {
					if currentNode.right == nil {
						nextNode := &binaryHuffmanNode{}
						currentNode.right = nextNode
						currentNode = nextNode
					} else if nextNode, ok := currentNode.right.(*binaryHuffmanNode); ok {
						currentNode = nextNode
					} else {
						panic("Invalid huffman tree")
					}
				}
			} else {
				// Ending

				if bit == "1" {
					currentNode.left = value
				} else {
					currentNode.right = value
				}
			}
		}
	}

	return root, depth
}

func (self *binaryTreeHuffmanDecoder) Decode(br bitreader.BitReader) (interface{}, error) {

	nextbits, err := br.Peek32(self.depth)

	if err != nil {
		return nil, err
	}

	currentNode := self.root

	for i := uint(1); i <= self.depth; i++ {
		var val interface{}

		mask := uint32(1) << (self.depth - i)
		bit := nextbits&mask == mask

		if bit {
			val = currentNode.left
		} else {
			val = currentNode.right
		}

		if val == nil {
			break
		} else if nextNode, ok := val.(*binaryHuffmanNode); ok {
			currentNode = nextNode
		} else {
			br.Trash(i)
			return val, nil
		}
	}

	return nil, ErrMissingHuffmanValue
}
