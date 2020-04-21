package expression

import (
	"fmt"
	"testing"
)

func TestSub(t *testing.T) {
	tests := []struct {
		op  *Sub
		err bool
		str string
		res float64
	}{
		{
			op: &Sub{
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
			str: "(2 - 1)",
			res: 1,
		},
		{
			op: &Sub{
				left: &Term{
					isSet: true,
					value: -3,
				},
				right: &Term{
					isSet: true,
					value: -5,
				},
			},
			err: false,
			str: "(-3 - -5)",
			res: 2,
		},
		{
			op: &Sub{
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
			str: "(-4 - 4)",
			res: -8,
		},
		{
			op: &Sub{
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
			str: "(x - 0)",
			res: 0,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if test.op.String() != test.str {
				t.Logf("Expected: %s, Got: %s", test.str, test.op.String())
				t.Fail()
			}

			res, err := test.op.eval()
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
