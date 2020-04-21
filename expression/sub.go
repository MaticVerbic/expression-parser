package expression

import (
	"fmt"

	"github.com/pkg/errors"
)

// Sub operation.
type Sub struct {
	left  Expr
	right Expr
}

// String representation of expression.
func (s *Sub) String() string {
	return fmt.Sprintf("(%s - %s)", s.left.String(), s.right.String())
}

// Eval evalutes the expression.
func (s *Sub) Eval() (float64, error) {
	leftVal, err := s.left.Eval()
	if err != nil {
		return 0, errors.Wrap(err, "failed to evaluate the left term")
	}

	rightVal, err := s.right.Eval()
	if err != nil {
		return 0, errors.Wrap(err, "failed to evaluate the right term")
	}

	return leftVal - rightVal, nil
}

// SetLeft sets left field, represents operation.
func (s *Sub) SetLeft(expr Expr) {
	s.left = expr
}

// SetRight sets right field, represents operation.
func (s *Sub) SetRight(expr Expr) {
	s.right = expr
}
