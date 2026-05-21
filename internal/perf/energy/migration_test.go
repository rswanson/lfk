package energy_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNoBareTeaTickInApp guards against regressions: every tea.Tick callsite
// in internal/app must go through energy.Tick so the probe can attribute it.
// New ticks without labels would silently make the harness blind.
func TestNoBareTeaTickInApp(t *testing.T) {
	root := filepath.Join("..", "..", "app")
	fset := token.NewFileSet()
	var offenders []string
	err := filepath.Walk(root, func(p string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(p, ".go") || strings.HasSuffix(p, "_test.go") {
			return err
		}
		f, perr := parser.ParseFile(fset, p, nil, 0)
		if perr != nil {
			return perr
		}
		ast.Inspect(f, func(n ast.Node) bool {
			ce, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			se, ok := ce.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}
			id, ok := se.X.(*ast.Ident)
			if !ok {
				return true
			}
			if id.Name == "tea" && se.Sel.Name == "Tick" {
				offenders = append(offenders, fset.Position(ce.Pos()).String())
			}
			return true
		})
		return nil
	})
	assert.NoError(t, err)
	assert.Empty(t, offenders, "bare tea.Tick callsites found in internal/app; use energy.Tick instead")
}
