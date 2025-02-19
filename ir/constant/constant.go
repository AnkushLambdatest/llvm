// Package constant implements values representing immutable LLVM IR constants.
package constant

import (
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// === [ Constants ] ===========================================================

// Convenience constants.
var (
	// None token constant.
	None = &NoneToken{} // none
	// Boolean constants.
	True  = NewInt(types.I1, 1) // true
	False = NewInt(types.I1, 0) // false
)

// Constant is an LLVM IR constant; a value that is immutable at runtime, such
// as an integer or floating-point literal, or the address of a function or
// global variable.
//
// A Constant has one of the following underlying types.
//
// Simple constants
//
// https://llvm.org/docs/LangRef.html#simple-constants
//
//    *constant.Int         // https://pkg.go.dev/github.com/llir/llvm/ir/constant#Int
//    *constant.Float       // https://pkg.go.dev/github.com/llir/llvm/ir/constant#Float
//    *constant.Null        // https://pkg.go.dev/github.com/llir/llvm/ir/constant#Null
//    *constant.NoneToken   // https://pkg.go.dev/github.com/llir/llvm/ir/constant#NoneToken
//
// Complex constants
//
// https://llvm.org/docs/LangRef.html#complex-constants
//
//    *constant.Struct            // https://pkg.go.dev/github.com/llir/llvm/ir/constant#Struct
//    *constant.Array             // https://pkg.go.dev/github.com/llir/llvm/ir/constant#Array
//    *constant.CharArray         // https://pkg.go.dev/github.com/llir/llvm/ir/constant#CharArray
//    *constant.Vector            // https://pkg.go.dev/github.com/llir/llvm/ir/constant#Vector
//    *constant.ZeroInitializer   // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ZeroInitializer
//    TODO: include metadata node?
//
// Global variable and function addresses
//
// https://llvm.org/docs/LangRef.html#global-variable-and-function-addresses
//
//    *ir.Global   // https://pkg.go.dev/github.com/llir/llvm/ir#Global
//    *ir.Func     // https://pkg.go.dev/github.com/llir/llvm/ir#Func
//    *ir.Alias    // https://pkg.go.dev/github.com/llir/llvm/ir#Alias
//    *ir.IFunc    // https://pkg.go.dev/github.com/llir/llvm/ir#IFunc
//
// Undefined values
//
// https://llvm.org/docs/LangRef.html#undefined-values
//
//    *constant.Undef   // https://pkg.go.dev/github.com/llir/llvm/ir/constant#Undef
//
// Poison values
//
// https://llvm.org/docs/LangRef.html#poison-values
//
//    *constant.Poison   // https://pkg.go.dev/github.com/llir/llvm/ir/constant#Poison
//
// Addresses of basic blocks
//
// https://llvm.org/docs/LangRef.html#addresses-of-basic-blocks
//
//    *constant.BlockAddress   // https://pkg.go.dev/github.com/llir/llvm/ir/constant#BlockAddress
//
// Constant expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    constant.Expression   // https://pkg.go.dev/github.com/llir/llvm/ir/constant#Expression
type Constant interface {
	value.Value
	// IsConstant ensures that only constants can be assigned to the
	// constant.Constant interface.
	IsConstant()
}
