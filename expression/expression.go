package expression

import (
	"fmt"

	"github.com/pkg/errors"
)

// Expr defines an expression interface.
type Expr interface {
	String() string
	Eval() (float64, error)
}

// Op defines an operation interface.
type Op interface {
	SetLeft(val Expr)
	SetRight(val Expr)
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

// NewAdd adds a new addition to expression.
func (e *Expression) NewAdd(left, right Expr) error {
	return e.newOp(left, right, &Add{})
}

// NewSub adds a new substitution to expression.
func (e *Expression) NewSub(left, right Expr) error {
	return e.newOp(left, right, &Sub{})
}

// NewDiv adds a new division to expression.
func (e *Expression) NewDiv(left, right Expr) error {
	return e.newOp(left, right, &Div{})
}

// NewMul adds a new multiplication to expression.
func (e *Expression) NewMul(left, right Expr) error {
	return e.newOp(left, right, &Mul{})
}

// NewConst creates a new constant.
func (e *Expression) NewConst(value interface{}, precision int) (*Term, error) {
	val, err := validateNum(value)
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

// NewVar returns a new variable.
func (e *Expression) NewVar(label string, precision int) (*Term, error) {
	if l := e.vars[label]; l != nil {
		return nil, errors.New("var already defined")
	}

	p := 0

	if precision > 0 {
		p = precision
	}

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
	res, err := e.expr.Eval()
	if err != nil {
		return res, errors.Wrap(err, e.String())
	}
	return res, nil
}

func validateNum(in interface{}) (float64, error) {
	switch in.(type) {
	case int, int8, int16, int32, int64:
		return float64(in.(int)), nil
	case float32:
		return float64(in.(float32)), nil
	case float64:
		return in.(float64), nil
	default:
		return 0, errors.New("invalid data type")
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

			op.SetLeft(e.expr)
			op.SetRight(right.(*Expression).expr)
			e.expr = op.(Expr)

		default:
			op.SetLeft(e.expr)
			op.SetRight(right)
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

			op.SetLeft(left.(*Expression).expr)
			op.SetRight(e.expr)
			e.expr = op.(Expr)

		default:

			op.SetLeft(left)
			op.SetRight(e.expr)
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

			op.SetLeft(left.(*Expression).expr)
			op.SetRight(right.(*Expression).expr)
			e.expr = op.(Expr)

		default:
			err := e.joinVars(left.(*Expression))
			if err != nil {
				return err
			}
			op.SetLeft(left.(*Expression).expr)
			op.SetRight(right)
			e.expr = op.(Expr)
		}
	default:
		switch right.(type) {
		case *Expression:
			err := e.joinVars(right.(*Expression))
			if err != nil {
				return err
			}

			op.SetLeft(left)
			op.SetRight(right.(*Expression).expr)
			e.expr = op.(Expr)

		default:
			op.SetLeft(left)
			op.SetRight(right)
			e.expr = op.(Expr)
		}
	}

	return nil
}
