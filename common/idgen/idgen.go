package idgen

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

type IDGen struct {
	node *snowflake.Node
}

func New(nodeNum int64) (*IDGen, error) {
	node, err := snowflake.NewNode(nodeNum)
	if err != nil {
		return nil, fmt.Errorf("idgen.New: %w", err)
	}
	return &IDGen{
		node: node,
	}, nil
}

func (i *IDGen) NewID() string {
	return i.node.Generate().String()
}
