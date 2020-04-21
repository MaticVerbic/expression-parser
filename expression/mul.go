package expression

import (
	"fmt"

	"github.com/pkg/errors"
)

// Mul operation.
type Mul struct {
	left  Expr
	right Expr
}

// String representation of expression.
func (m *Mul) String() string {
	return fmt.Sprintf("(%s * %s)", m.left.String(), m.right.String())
}

// Eval evalutes the expression.
func (m *Mul) eval() (float64, error) {
	leftVal, err := m.left.eval()
	if err != nil {
		return 0, errors.Wrap(err, "failed to evaluate the left term")
	}

	rightVal, err := m.right.eval()
	if err != nil {
		return 0, errors.Wrap(err, "failed to evaluate the right term")
	}

	return leftVal * rightVal, nil
}

// setLeft sets left field, represents operation.
func (m *Mul) setLeft(expr Expr) {
	m.left = expr
}

// setRight sets right field, represents operation.
func (m *Mul) setRight(expr Expr) {
	m.right = expr
}
