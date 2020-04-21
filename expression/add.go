package expression

import (
	"fmt"

	"github.com/pkg/errors"
)

// Add operation.
type Add struct {
	left  Expr
	right Expr
}

// String representation of expression.
func (a *Add) String() string {
	return fmt.Sprintf("(%s + %s)", a.left.String(), a.right.String())
}

// Eval evalutes the expression.
func (a *Add) Eval() (float64, error) {
	leftVal, err := a.left.Eval()
	if err != nil {
		return 0, errors.Wrap(err, "failed to evaluate the left term")
	}

	rightVal, err := a.right.Eval()
	if err != nil {
		return 0, errors.Wrap(err, "failed to evaluate the right term")
	}

	return leftVal + rightVal, nil
}

// SetLeft sets left field, represents operation.
func (a *Add) SetLeft(expr Expr) {
	a.left = expr
}

// SetRight sets right field, represents operation.
func (a *Add) SetRight(expr Expr) {
	a.right = expr
}
