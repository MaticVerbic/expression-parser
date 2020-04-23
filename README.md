# expression-parser

Tree based expression parser.

### Documentation

```go
expr := expression.New() // Create a new expression parser.

const, err := expr.Const(5) // Create a new constant.
variable, err := expr.Variable("x") // Create a new variable.
variable.SetVal(3.14) // Set a variable value.
variable.UnsetValue() // Unset a variable value.

err := expr.Add(left, right) // Add a new addition to the top of the stack.
err := expr.Sub(left, right) // Add a new substitution to the top of the stack.
err := expr.Mul(left, right) // Add a new multiplication to the top of the stack.
err := expr.Div(left, right) // Add a new division to the top of the stack.

expr.Eval() // Evaluate the expression at given parameters.
```

### Example

```go
linearFunc := expression.New()

k, err := linearFunc.Const(5)
if err != nil {
  panic(err)
}

n, err := linearFunc.Const(3)
if err != nil {
  panic(err)
}

x, err := linearFunc.Var("x")
if err != nil {
  panic(err)
}

err = linearFunc.Mul(k, x)
if err != nil {
  panic(err)
}

fmt.Println(linearFunc) // (5 * x)

err = linearFunc.Add(nil, n)
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