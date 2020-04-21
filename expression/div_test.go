package expression

import (
	"fmt"
	"testing"
)

func TestDiv(t *testing.T) {
	tests := []struct {
		op  *Div
		err bool
		str string
		res float64
	}{
		{
			op: &Div{
				left: &Term{
					isSet: true,
					value: 2,
				},
				right: &Term{
					isSet: true,
					value: 1,
				},
			},
			err: false,
			str: "(2 / 1)",
			res: 2,
		},
		{
			op: &Div{
				left: &Term{
					isSet: true,
					value: -10,
				},
				right: &Term{
					isSet: true,
					value: -2,
				},
			},
			err: false,
			str: "(-10 / -2)",
			res: 5,
		},
		{
			op: &Div{
				left: &Term{
					isSet: true,
					value: -4,
				},
				right: &Term{
					isSet: true,
					value: 4,
				},
			},
			err: false,
			str: "(-4 / 4)",
			res: -1,
		},
		{
			op: &Div{
				left: &Term{
					isSet: false,
					label: "x",
				},
				right: &Term{
					isSet: true,
					value: 0,
				},
			},
			err: true,
			str: "(x / 0)",
			res: 0,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if test.op.String() != test.str {
				t.Logf("Expected: %s, Got: %s", test.str, test.op.String())
				t.Fail()
			}

			res, err := test.op.Eval()
			if err != nil {
				if test.err {
					return
				}

				t.Errorf("%v", err)
			}
			if res != test.res {
				t.Logf("Expected: %4.f, Got: %.4f", test.res, res)
				t.Fail()
			}
		})
	}
}
