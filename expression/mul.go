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
func (m *Mul) Eval() (float64, error) {
	leftVal, err := m.left.Eval()
	if err != nil {
		return 0, errors.Wrap(err, "failed to evaluate the left term")
	}

	rightVal, err := m.right.Eval()
	if err != nil {
		return 0, errors.Wrap(err, "failed to evaluate the right term")
	}

	return leftVal * rightVal, nil
}

// SetLeft sets left field, represents operation.
func (m *Mul) SetLeft(expr Expr) {
	m.left = expr
}

// SetRight sets right field, represents operation.
func (m *Mul) SetRight(expr Expr) {
	m.right = expr
}
