package main

import (
	"bytes"
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

var (
	// TODO allow this to be overwritten by flags
	MaxInputParamsOnOneLine = 3
)

func main() {
	singlechecker.Main(Analyzer)
}

var Analyzer = &analysis.Analyzer{
	Name: "newlineInputParams",
	Doc:  "Reports when input parameters should be on newlines",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			fd, ok := n.(*ast.FuncDecl)
			if !ok {
				// its not a func, keep looking
				return true
			}

			params := fd.Type.Params
			expected := getExpectedInputs(params)

			if !shouldReplace(params, expected) {
				// continue on
				return true
			}

			pass.Report(analysis.Diagnostic{
				Pos:            params.Pos(),
				Message:        "function with input params that should be on separate lines",
				SuggestedFixes: convertToSuggestions(params, expected),
			})

			return true
		})
	}

	return nil, nil
}

type nameByType struct {
	Names []string
	Type  string
}

func getExpectedInputs(params *ast.FieldList) []nameByType {
	if len(params.List) == 0 {
		return nil
	}

	res := make([]nameByType, 0, len(params.List))
	res = append(res, nameByType{
		Type: fmt.Sprintf("%s", params.List[0].Type),
	})

	for pi := 0; pi < len(params.List); pi++ {
		piType := fmt.Sprintf("%s", params.List[pi].Type)
		if piType != res[len(res)-1].Type {
			res = append(res, nameByType{
				Type: piType,
			})
		}

		for _, n := range params.List[pi].Names {
			res[len(res)-1].Names = append(res[len(res)-1].Names, n.Name)
		}
	}

	return res
}

func shouldReplace(
	params *ast.FieldList,
	expected []nameByType,
) bool {
	if len(expected) == 0 {
		return false
	}

	numInputs := 0
	for _, p := range params.List {
		numInputs += len(p.Names)
	}
	// TODO if the number of newlines in the actual file is more than the len(expected),
	// then we probably don't need to re-address?
	return numInputs > MaxInputParamsOnOneLine
}

func convertToSuggestions(
	params *ast.FieldList,
	expected []nameByType,
) []analysis.SuggestedFix {
	var newExpr bytes.Buffer
	newExpr.WriteString("(\n")

	for _, e := range expected {
		newExpr.WriteString("\t")
		for i, n := range e.Names {
			if i > 0 {
				newExpr.WriteString(`, `)
			}
			newExpr.WriteString(n)
		}
		newExpr.WriteString(` `)
		newExpr.WriteString(e.Type)
		newExpr.WriteString(",\n")

	}

	return []analysis.SuggestedFix{{
		Message: "Place one input param (type) on each line",
		TextEdits: []analysis.TextEdit{{
			Pos:     params.Opening,
			End:     params.Closing,
			NewText: newExpr.Bytes(),
		}},
	}}
}
