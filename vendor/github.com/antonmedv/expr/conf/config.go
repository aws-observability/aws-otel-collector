package conf

import (
	"fmt"
	"reflect"

	"github.com/antonmedv/expr/ast"
	"github.com/antonmedv/expr/builtin"
	"github.com/antonmedv/expr/vm/runtime"
)

type Config struct {
<<<<<<< HEAD
	Env         interface{}
=======
	Env         any
>>>>>>> main
	Types       TypesTable
	MapEnv      bool
	DefaultType reflect.Type
	Operators   OperatorsTable
	Expect      reflect.Kind
<<<<<<< HEAD
=======
	ExpectAny   bool
>>>>>>> main
	Optimize    bool
	Strict      bool
	ConstFns    map[string]reflect.Value
	Visitors    []ast.Visitor
<<<<<<< HEAD
	Functions   map[string]*builtin.Function
=======
	Functions   map[string]*ast.Function
	Builtins    map[string]*ast.Function
	Disabled    map[string]bool // disabled builtins
>>>>>>> main
}

// CreateNew creates new config with default values.
func CreateNew() *Config {
	c := &Config{
<<<<<<< HEAD
		Operators: make(map[string][]string),
		ConstFns:  make(map[string]reflect.Value),
		Functions: make(map[string]*builtin.Function),
		Optimize:  true,
	}
	for _, f := range builtin.Builtins {
		c.Functions[f.Name] = f
=======
		Optimize:  true,
		Operators: make(map[string][]string),
		ConstFns:  make(map[string]reflect.Value),
		Functions: make(map[string]*ast.Function),
		Builtins:  make(map[string]*ast.Function),
		Disabled:  make(map[string]bool),
	}
	for _, f := range builtin.Builtins {
		c.Builtins[f.Name] = f
>>>>>>> main
	}
	return c
}

// New creates new config with environment.
<<<<<<< HEAD
func New(env interface{}) *Config {
=======
func New(env any) *Config {
>>>>>>> main
	c := CreateNew()
	c.WithEnv(env)
	return c
}

<<<<<<< HEAD
func (c *Config) WithEnv(env interface{}) {
	var mapEnv bool
	var mapValueType reflect.Type
	if _, ok := env.(map[string]interface{}); ok {
=======
func (c *Config) WithEnv(env any) {
	var mapEnv bool
	var mapValueType reflect.Type
	if _, ok := env.(map[string]any); ok {
>>>>>>> main
		mapEnv = true
	} else {
		if reflect.ValueOf(env).Kind() == reflect.Map {
			mapValueType = reflect.TypeOf(env).Elem()
		}
	}

	c.Env = env
	c.Types = CreateTypesTable(env)
	c.MapEnv = mapEnv
	c.DefaultType = mapValueType
	c.Strict = true
}

func (c *Config) Operator(operator string, fns ...string) {
	c.Operators[operator] = append(c.Operators[operator], fns...)
}

func (c *Config) ConstExpr(name string) {
	if c.Env == nil {
		panic("no environment is specified for ConstExpr()")
	}
	fn := reflect.ValueOf(runtime.Fetch(c.Env, name))
	if fn.Kind() != reflect.Func {
		panic(fmt.Errorf("const expression %q must be a function", name))
	}
	c.ConstFns[name] = fn
}

func (c *Config) Check() {
	for operator, fns := range c.Operators {
		for _, fn := range fns {
			fnType, ok := c.Types[fn]
			if !ok || fnType.Type.Kind() != reflect.Func {
				panic(fmt.Errorf("function %s for %s operator does not exist in the environment", fn, operator))
			}
			requiredNumIn := 2
			if fnType.Method {
				requiredNumIn = 3 // As first argument of method is receiver.
			}
			if fnType.Type.NumIn() != requiredNumIn || fnType.Type.NumOut() != 1 {
				panic(fmt.Errorf("function %s for %s operator does not have a correct signature", fn, operator))
			}
		}
	}
<<<<<<< HEAD
=======
	for fnName, t := range c.Types {
		if kind(t.Type) == reflect.Func {
			for _, b := range c.Builtins {
				if b.Name == fnName {
					panic(fmt.Errorf(`cannot override builtin %s(): use expr.DisableBuiltin("%s") to override`, b.Name, b.Name))
				}
			}
		}
	}
	for _, f := range c.Functions {
		for _, b := range c.Builtins {
			if b.Name == f.Name {
				panic(fmt.Errorf(`cannot override builtin %s(); use expr.DisableBuiltin("%s") to override`, f.Name, f.Name))
			}
		}
	}
>>>>>>> main
}
