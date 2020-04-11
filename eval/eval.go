package eval

import (
	"fmt"
	"math"
)

// A Var identifies a variable
type Var string

// A literal is a numeric constant
type literal float64

// Env is environment for variable resultion
type Env map[Var]float64

// Expr is generalized expression
type Expr interface {
	Eval(env Env) float64
}

type unary struct {
	op rune
	x  Expr
}

type binary struct {
	op   rune
	x, y Expr
}

type call struct {
	fn   string
	args []Expr
}

// Eval return the value of variable represented by Var
func (v Var) Eval(env Env) float64 {
	return env[v]
}

// Eval return the value literal
func (l literal) Eval(env Env) float64 {
	return float64(l)
}

// Eval returns unary operator result
func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("Unsupported unary operator: %q", u.op))
}

// Eval returns unary operator result
func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("Unsupported binary operator: %q", b.op))
}

// Eval returns unary operator result
func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("Unsupported function call: %q", c.fn))
}

/*func Parse(expr string) (Expr, error) {
	// TODO: fix here later
	return nil, nil
}*/
