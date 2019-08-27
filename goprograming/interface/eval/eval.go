package eval

import (
	"fmt"
	"math"
)

type Env map[Var]float64

// Expr: expression.
type Expr interface {
	Eval(env Env) float64
}

// Var represent a variable like x.
type Var string

// literal is a numeric const, like 3.14
type literal float64

// unary represent unary operator expression like -x
type unary struct {
	op rune // '+' or '-'
	x  Expr
}

// binary represent binary operator expression like x+y
type binary struct {
	op   rune // '+', '-', '*' or '/'
	x, y Expr
}

// call represent function call like sin(x)
type call struct {
	fn   string // on of "pow", "sin", "sqrt"
	args []Expr
}

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

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
	panic(fmt.Sprintf("unsupported unary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", c.fn))
}
