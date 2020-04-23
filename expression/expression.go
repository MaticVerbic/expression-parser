package expression

import (
	"fmt"

	"github.com/pkg/errors"
)

// Expr defines an expression interface.
type Expr interface {
	String() string
	eval() (float64, error)
}

// Op defines an operation interface.
type Op interface {
	setLeft(val Expr)
	setRight(val Expr)
}

// Expression ...
type Expression struct {
	expr Expr
	vars map[string]*Term
}

// New returns new Expression.
func New() *Expression {
	return &Expression{
		vars: map[string]*Term{},
	}
}

// Add adds a new addition to expression.
func (e *Expression) Add(left, right Expr) error {
	return e.newOp(left, right, &Add{})
}

// Sub adds a new substitution to expression.
func (e *Expression) Sub(left, right Expr) error {
	return e.newOp(left, right, &Sub{})
}

// Div adds a new division to expression.
func (e *Expression) Div(left, right Expr) error {
	return e.newOp(left, right, &Div{})
}

// Mul adds a new multiplication to expression.
func (e *Expression) Mul(left, right Expr) error {
	return e.newOp(left, right, &Mul{})
}

// Const creates a new constant.
func (e *Expression) Const(value interface{}) (*Term, error) {
	val, precision, err := validateNum(value)
	if err != nil {
		return nil, errors.Wrap(err, "failed to validate value")
	}

	p := 0

	if precision > 0 {
		p = precision
	}

	return &Term{
		value:           val,
		isSet:           true,
		stringPrecision: fmt.Sprintf("%d", p),
	}, nil
}

// Var returns a new variable.
func (e *Expression) Var(label string) (*Term, error) {
	if l := e.vars[label]; l != nil {
		return nil, errors.New("var already defined")
	}

	p := 0

	t := &Term{
		label:           label,
		stringPrecision: fmt.Sprintf("%d", p),
	}

	e.vars[label] = t
	return t, nil
}

// String representation of expression.
func (e *Expression) String() string {
	return fmt.Sprintf("%s", e.expr.String())
}

// Eval evalutes the expression.
func (e *Expression) Eval() (float64, error) {
	return e.eval()
}

// Eval evalutes the expression.
func (e *Expression) eval() (float64, error) {
	res, err := e.expr.eval()
	if err != nil {
		return res, errors.Wrap(err, e.String())
	}
	return res, nil
}

func validateNum(in interface{}) (float64, int, error) {
	switch in.(type) {
	case int, int8, int16, int32, int64:
		return float64(in.(int)), 0, nil
	case float32:
		return float64(in.(float32)), 2, nil
	case float64:
		return in.(float64), 2, nil
	default:
		return 0, 0, errors.New("invalid data type")
	}
}

func (e *Expression) joinVars(other *Expression) error {
	for key, value := range other.vars {
		if val := e.vars[key]; val != nil && val != value {
			return fmt.Errorf("variable douplicate: %s", key)
		}

		e.vars[key] = value
	}

	return nil
}

// NewOp adds a new operation to Expression
func (e *Expression) newOp(left, right Expr, op Op) error {
	if e.expr == nil && (left == nil || right == nil) {
		return errors.New("nil expression")
	}

	if left == nil {
		if right == nil {
			return fmt.Errorf("invalid input")
		}

		switch right.(type) {
		case *Expression:
			err := e.joinVars(right.(*Expression))
			if err != nil {
				return err
			}

			op.setLeft(e.expr)
			op.setRight(right.(*Expression).expr)
			e.expr = op.(Expr)

		default:
			op.setLeft(e.expr)
			op.setRight(right)
			e.expr = op.(Expr)
		}

		return nil
	}

	if right == nil {
		switch left.(type) {
		case *Expression:
			err := e.joinVars(left.(*Expression))
			if err != nil {
				return err
			}

			op.setLeft(left.(*Expression).expr)
			op.setRight(e.expr)
			e.expr = op.(Expr)

		default:

			op.setLeft(left)
			op.setRight(e.expr)
			e.expr = op.(Expr)
		}

		return nil
	}

	switch left.(type) {
	case *Expression:
		switch right.(type) {
		case *Expression:
			err := e.joinVars(left.(*Expression))
			if err != nil {
				return err
			}
			err = e.joinVars(right.(*Expression))
			if err != nil {
				return err
			}

			op.setLeft(left.(*Expression).expr)
			op.setRight(right.(*Expression).expr)
			e.expr = op.(Expr)

		default:
			err := e.joinVars(left.(*Expression))
			if err != nil {
				return err
			}
			op.setLeft(left.(*Expression).expr)
			op.setRight(right)
			e.expr = op.(Expr)
		}
	default:
		switch right.(type) {
		case *Expression:
			err := e.joinVars(right.(*Expression))
			if err != nil {
				return err
			}

			op.setLeft(left)
			op.setRight(right.(*Expression).expr)
			e.expr = op.(Expr)

		default:
			op.setLeft(left)
			op.setRight(right)
			e.expr = op.(Expr)
		}
	}

	return nil
}
