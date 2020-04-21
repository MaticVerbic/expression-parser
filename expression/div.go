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
func (d *Div) Eval() (float64, error) {
	leftVal, err := d.left.Eval()
	if err != nil {
		return 0, errors.Wrap(err, "failed to evaluate the left term")
	}

	rightVal, err := d.right.Eval()
	if err != nil {
		return 0, errors.Wrap(err, "failed to evaluate the right term")
	}

	return leftVal / rightVal, nil
}

// SetLeft sets left field, represents operation.
func (d *Div) SetLeft(expr Expr) {
	d.left = expr
}

// SetRight sets right field, represents operation.
func (d *Div) SetRight(expr Expr) {
	d.right = expr
}
