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
func (a *Add) eval() (float64, error) {
	leftVal, err := a.left.eval()
	if err != nil {
		return 0, errors.Wrap(err, "failed to evaluate the left term")
	}

	rightVal, err := a.right.eval()
	if err != nil {
		return 0, errors.Wrap(err, "failed to evaluate the right term")
	}

	return leftVal + rightVal, nil
}

// setLeft sets left field, represents operation.
func (a *Add) setLeft(expr Expr) {
	a.left = expr
}

// setRight sets right field, represents operation.
func (a *Add) setRight(expr Expr) {
	a.right = expr
}
