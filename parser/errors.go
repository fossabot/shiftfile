package parser

import (
	"fmt"

	"github.com/elasticshift/shiftfile/token"
)

type PositionErr struct {
	Position token.Position
	Err      error
}

func (pe *PositionErr) Error() string {
	return fmt.Sprintf("At %s:%s", pe.Position, pe.Err)
}
