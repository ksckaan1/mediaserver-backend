package idgen

import (
	"fmt"
	"mediaserver/internal/port"

	"github.com/bwmarrin/snowflake"
)

var _ port.IDGenerator = (*IDGen)(nil)

type IDGen struct {
	node *snowflake.Node
}

func New() (*IDGen, error) {
	node, err := snowflake.NewNode(1)
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
