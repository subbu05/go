// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build typeparams
// +build typeparams

package typeparams

import (
	"fmt"
	"go/ast"
)

const Enabled = true

func PackExpr(list []ast.Expr) ast.Expr {
	switch len(list) {
	case 0:
		return nil
	case 1:
		return list[0]
	default:
		return &ast.ListExpr{ElemList: list}
	}
}

// TODO(gri) Should find a more efficient solution that doesn't
//           require introduction of a new slice for simple
//           expressions.
func UnpackExpr(x ast.Expr) []ast.Expr {
	if x, _ := x.(*ast.ListExpr); x != nil {
		return x.ElemList
	}
	if x != nil {
		return []ast.Expr{x}
	}
	return nil
}

func Get(n ast.Node) *ast.FieldList {
	switch n := n.(type) {
	case *ast.TypeSpec:
		return n.TParams
	case *ast.FuncType:
		return n.TParams
	default:
		panic(fmt.Sprintf("node type %T has no type parameters", n))
	}
}

func Set(n ast.Node, params *ast.FieldList) {
	switch n := n.(type) {
	case *ast.TypeSpec:
		n.TParams = params
	case *ast.FuncType:
		n.TParams = params
	default:
		panic(fmt.Sprintf("node type %T has no type parameters", n))
	}
}
