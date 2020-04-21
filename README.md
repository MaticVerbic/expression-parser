# expression-parser

Tree based expression parser.

### Documentation

```go
expr := expression.New() // Create a new expression parser.

const, err := expr.NewConst(5, 0) // Create a new constant.
variable, err := expr.NewVariable("x", 0) // Create a new variable.
variable.SetVal(3.14) // Set a variable value.
variable.UnsetValue() // Unset a variable value.

err := expr.NewAdd(left, right) // Add a new addition to the top of the stack.
err := expr.NewSub(left, right) // Add a new substitution to the top of the stack.
err := expr.NewMul(left, right) // Add a new multiplication to the top of the stack.
err := expr.NewDiv(left, right) // Add a new division to the top of the stack.

expr.Eval() // Evaluate the expression at given parameters.
```

### Example

```go
linearFunc := expression.New()

k, err := linearFunc.NewConst(5, 0)
if err != nil {
  panic(err)
}

n, err := linearFunc.NewConst(3, 0)
if err != nil {
  panic(err)
}

x, err := linearFunc.NewVar("x", 2)
if err != nil {
  panic(err)
}

err = linearFunc.NewMul(k, x)
if err != nil {
  panic(err)
}

fmt.Println(linearFunc) // (5 * x)

err = linearFunc.NewAdd(nil, n)
if err != nil {
  panic(err)
}

fmt.Println(linearFunc) // ((5 * x) + 3)

x.SetVal(2)
result, err := linearFunc.Eval()
if err != nil {
  panic(err)
}

fmt.Println(result) // 13
```