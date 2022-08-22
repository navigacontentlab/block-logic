package blocklogic

import (
	"github.com/navigacontentlab/navigadoc/doc"
)

type Condition struct {
	In    string      `json:"in,omitempty"`
	And   []Condition `json:"and,omitempty"`
	Or    []Condition `json:"or,omitempty"`
	Rel   string      `json:"rel,omitempty"`
	UUID  string      `json:"uuid,omitempty"`
	Type  string      `json:"type,omitempty"`
	Title string      `json:"title,omitempty"`
	URI   string      `json:"uri,omitempty"`
}

func (c Condition) TestDocument(document doc.Document) bool {
	blocks := []doc.Block{}

	switch c.In {
	case "links":
		blocks = append(blocks, document.Links...)
	case "meta":
		blocks = append(blocks, document.Meta...)
	case "content":
		blocks = append(blocks, document.Content...)
	case "*":
		blocks = append(blocks, document.Links...)
		blocks = append(blocks, document.Meta...)
		blocks = append(blocks, document.Content...)
	}

	return c.testBlocks(true, blocks)
}

func (c Condition) testBlocks(toplevel bool, blocks []doc.Block) bool {
	if !toplevel && c.In != "" {
		childBlocks := []doc.Block{}

		for _, block := range blocks {
			switch c.In {
			case "links":
				childBlocks = append(childBlocks, block.Links...)
			case "meta":
				childBlocks = append(childBlocks, block.Meta...)
			case "content":
				childBlocks = append(childBlocks, block.Content...)
			case "*":
				childBlocks = append(childBlocks, block.Links...)
				childBlocks = append(childBlocks, block.Meta...)
				childBlocks = append(childBlocks, block.Content...)
			}
		}

		return c.test(childBlocks)
	}

	return c.test(blocks)
}

func (c Condition) test(blocks []doc.Block) bool {
	hasMatch := false

	for _, block := range blocks {
		if matchBlock(c, block) {
			hasMatch = true

			break
		}
	}

	if !hasMatch {
		return false
	}

	if c.Or != nil && !c.testOr(blocks) {
		return false
	}

	if c.And != nil && !c.testAnd(blocks) {
		return false
	}

	return true
}

func (c Condition) testOr(blocks []doc.Block) bool {
	for _, pattern := range c.Or {
		if pattern.testBlocks(false, blocks) {
			return true
		}
	}

	return false
}

func (c Condition) testAnd(blocks []doc.Block) bool {
	for _, pattern := range c.And {
		if !pattern.testBlocks(false, blocks) {
			return false
		}
	}

	return true
}

func matchBlock(pattern Condition, block doc.Block) bool {
	// UUID
	if pattern.UUID != "" && pattern.UUID != block.UUID {
		return false
	}

	// Type
	if pattern.Type != "" && pattern.Type != block.Type {
		return false
	}

	// Rel
	if pattern.Rel != "" && pattern.Rel != block.Rel {
		return false
	}

	// Title
	if pattern.Title != "" && pattern.Title != block.Title {
		return false
	}

	// URI
	if pattern.URI != "" && pattern.URI != block.URI {
		return false
	}

	return true
}
