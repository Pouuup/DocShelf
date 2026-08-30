package models

type BlockType string

const (
	BlockParagraph BlockType = "paragraph"
	BlockHeading   BlockType = "heading"
	BlockCode      BlockType = "code"
)

type Block struct {
	Type BlockType
	Text string
}
