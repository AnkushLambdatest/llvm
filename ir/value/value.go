// Package value provides a definition of LLVM IR values.
package value

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
)

// Value is an LLVM IR value, which may be used as an operand of instructions
// and terminators.
//
// A Value has one of the following underlying types.
//
//    constant.Constant   // https://pkg.go.dev/github.com/llir/llvm/ir/constant#Constant
//    value.Named         // https://pkg.go.dev/github.com/llir/llvm/ir/value#Named
//    TODO: add literal metadata value?
type Value interface {
	// String returns the LLVM syntax representation of the value as a type-value
	// pair.
	fmt.Stringer
	// Type returns the type of the value.
	Type() types.Type
	// Ident returns the identifier associated with the value.
	Ident() string
}

// Named is a named LLVM IR value.
//
// A Named value has one of the following underlying types.
//
//    *ir.Global            // https://pkg.go.dev/github.com/llir/llvm/ir#Global
//    *ir.Func              // https://pkg.go.dev/github.com/llir/llvm/ir#Func
//    *ir.Param             // https://pkg.go.dev/github.com/llir/llvm/ir#Param
//    *ir.Block             // https://pkg.go.dev/github.com/llir/llvm/ir#Block
//    TODO: add named metadata value?
//    ir.Instruction        // https://pkg.go.dev/github.com/llir/llvm/ir#Instruction (except store and fence)
//    *ir.TermInvoke        // https://pkg.go.dev/github.com/llir/llvm/ir#TermInvoke
//    *ir.TermCatchSwitch   // https://pkg.go.dev/github.com/llir/llvm/ir#TermCatchSwitch (token result used by catchpad)
type Named interface {
	Value
	// Name returns the name of the value.
	Name() string
	// SetName sets the name of the value.
	SetName(name string)
}

// User is an instruction or terminator which uses values as operands.
type User interface {
	// Operands returns a mutable list of operands of the given value user
	// (instruction or terminator).
	Operands() []*Value
}
