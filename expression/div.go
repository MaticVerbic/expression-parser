package expression

import (
	"fmt"

	"github.com/pkg/errors"
)

// Div operation.
type Div struct {
	left  Expr
	right Expr
}

// String representation of expression.
func (d *Div) String() string {
	return fmt.Sprintf("(%s / %s)", d.left.String(), d.right.String())
}

// Eval evalutes the expression.
func (d *Div) eval() (float64, error) {
	leftVal, err := d.left.eval()
	if err != nil {
		return 0, errors.Wrap(err, "failed to evaluate the left term")
	}

	rightVal, err := d.right.eval()
	if err != nil {
		return 0, errors.Wrap(err, "failed to evaluate the right term")
	}

	return leftVal / rightVal, nil
}

// setLeft sets left field, represents operation.
func (d *Div) setLeft(expr Expr) {
	d.left = expr
}

// setRight sets right field, represents operation.
func (d *Div) setRight(expr Expr) {
	d.right = expr
}
