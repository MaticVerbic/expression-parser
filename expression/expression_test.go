package expression

import (
	"fmt"
	"testing"
)

func TestExpression(t *testing.T) {
	expr1 := New()
	expr2 := New()

	c1, err := expr1.NewConst(1, 0) // represents an integer in string form
	if err != nil {
		t.Errorf("%v", err)
	}

	c2, err := expr1.NewConst(2, 1) // represents an float in string form
	if err != nil {
		t.Errorf("%v", err)
	}

	c3, err := expr1.NewConst(3.14, 2)
	if err != nil {
		t.Errorf("%v", err)
	}

	v1, err := expr1.NewVar("x", 0)
	if err != nil {
		t.Errorf("%v", err)
	}

	v2, err := expr1.NewVar("y", 1)
	if err != nil {
		t.Errorf("%v", err)
	}

	v3, err := expr1.NewVar("z", 2)
	if err != nil {
		t.Errorf("%v", err)
	}

	t.Run("expr1", func(t *testing.T) {
		err = expr1.NewAdd(c1, c3)
		if err != nil {
			t.Errorf("%v", err)
		}

		if expr1.String() != "(1 + 3.14)" {
			t.Logf("Expected: %s Got: %s", expr1.String(), "(1 + 3.14)")
			t.FailNow()
		}

		res, err := expr1.Eval()
		if err != nil {
			t.Errorf("%v", err)
		}

		if fmt.Sprintf("%.2f", res) != "4.14" {
			t.Logf("Expected: %s Got: %s", "4.14", fmt.Sprintf("%.2f", res))
			t.FailNow()
		}

		err = expr1.NewSub(nil, c2)
		if err != nil {
			t.Errorf("%v", err)
		}

		if expr1.String() != "((1 + 3.14) - 2.0)" {
			t.Logf("Expected: %s Got: %s", "((1 + 3.14) - 2.0)", expr1.String())
			t.FailNow()
		}

		res, err = expr1.Eval()
		if err != nil {
			t.Errorf("%v", err)
		}

		if fmt.Sprintf("%.2f", res) != "2.14" {
			t.Logf("Expected: %s Got: %s", "2.14", fmt.Sprintf("%.2f", res))
			t.FailNow()
		}
	})

	t.Run("expr2", func(t *testing.T) {
		err = expr2.NewMul(v1, v2)
		if err != nil {
			t.Errorf("%v", err)
		}

		if expr2.String() != "(x * y)" {
			t.Logf("Expected: %s Got: %s", expr1.String(), "(x * y)")
			t.FailNow()
		}

		_, err = expr2.Eval()
		if err == nil {
			t.Log("expected error but got none")
			t.FailNow()
		}

		err = expr2.NewDiv(v3, nil)
		if err != nil {
			t.Errorf("%v", err)
		}

		if expr2.String() != "(z / (x * y))" {
			t.Logf("Expected: %s Got: %s", expr1.String(), "(z / (x * y))")
			t.FailNow()
		}

		_, err = expr2.Eval()
		if err == nil {
			t.Log("expected error but got none")
			t.FailNow()
		}

		err = v1.SetVal(1)
		if err != nil {
			t.Errorf("%v", err)
		}

		err = v2.SetVal(2)
		if err != nil {
			t.Errorf("%v", err)
		}

		err = v3.SetVal(10)
		if err != nil {
			t.Errorf("%v", err)
		}

		res, err := expr2.Eval()
		if err != nil {
			t.Errorf("%v", err)
		}

		if fmt.Sprintf("%.2f", res) != "5.00" {
			t.Logf("Expected: %s Got: %s", "5.00", fmt.Sprintf("%.2f", res))
			t.FailNow()
		}
	})
}
