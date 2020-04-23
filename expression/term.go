package expression

import (
	"fmt"

	"github.com/pkg/errors"
)

// Term ...
type Term struct {
	label           string
	value           float64
	isSet           bool
	stringPrecision string
}

// UnsetVal unsets a value of Term.
func (t *Term) UnsetVal() {
	t.isSet = false
}

// SetVal sets a value of Term.
func (t *Term) SetVal(value interface{}) error {
	val, precision, err := validateNum(value)
	if err != nil {
		return errors.Wrap(err, "failed to validate value")
	}

	t.value = val
	t.isSet = true
	t.stringPrecision = fmt.Sprintf("%d", precision)
	return nil
}

// Eval evalutes the expression.
func (t *Term) eval() (float64, error) {
	if !t.isSet {
		return 0, fmt.Errorf("term not set: %s", t.label)
	}

	return t.value, nil
}

// String representation of expression.
func (t *Term) String() string {
	if !t.isSet {
		return t.label
	}

	return fmt.Sprintf("%."+t.stringPrecision+"f", t.value)
}
