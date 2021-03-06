package builder

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"sort"
	"strings"
)

type visitor struct {
	Path    string
	Package string
	Boxes   []string
	Errors  []error
}

func newVisitor(path string) *visitor {
	return &visitor{
		Path:   path,
		Boxes:  []string{},
		Errors: []error{},
	}
}

func (v *visitor) Run() error {
	b, err := ioutil.ReadFile(v.Path)
	if err != nil {
		return err
	}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, v.Path, string(b), parser.ParseComments)
	if err != nil {
		return err
	}

	v.Package = file.Name.Name
	ast.Walk(v, file)

	m := map[string]string{}
	for _, s := range v.Boxes {
		m[s] = s
	}
	v.Boxes = []string{}
	for k := range m {
		v.Boxes = append(v.Boxes, k)
	}

	sort.Strings(v.Boxes)

	if len(v.Errors) > 0 {
		s := make([]string, len(v.Errors))
		for i, e := range v.Errors {
			s[i] = e.Error()
		}
		return errors.New(strings.Join(s, "\n"))
	}
	return nil
}

func (v *visitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return v
	}
	if err := v.eval(node); err != nil {
		v.Errors = append(v.Errors, err)
	}
	return v
}

func (v *visitor) eval(node ast.Node) error {
	switch t := node.(type) {
	case *ast.CallExpr:
		return v.evalExpr(t)
	case *ast.Ident:
		return v.evalIdent(t)
	case *ast.GenDecl:
		for _, n := range t.Specs {
			if err := v.eval(n); err != nil {
				return err
			}
		}
	case *ast.FuncDecl:
		if t.Body == nil {
			return nil
		}
		for _, b := range t.Body.List {
			if err := v.evalStmt(b); err != nil {
				return err
			}
		}
		return nil
	case *ast.ValueSpec:
		for _, e := range t.Values {
			if err := v.evalExpr(e); err != nil {
				return err
			}
		}
	}
	return nil
}

func (v *visitor) evalStmt(stmt ast.Stmt) error {
	switch t := stmt.(type) {
	case *ast.ExprStmt:
		return v.evalExpr(t.X)
	case *ast.AssignStmt:
		for _, e := range t.Rhs {
			if err := v.evalArgs(e); err != nil {
				return err
			}
		}
	}
	return nil
}

func (v *visitor) evalExpr(expr ast.Expr) error {
	switch t := expr.(type) {
	case *ast.CallExpr:
		if t.Fun == nil {
			return nil
		}
		for _, a := range t.Args {
			switch at := a.(type) {
			case *ast.CallExpr:
				if sel, ok := t.Fun.(*ast.SelectorExpr); ok {
					return v.evalSelector(at, sel)
				}

				if err := v.evalArgs(at); err != nil {
					return err
				}
			case *ast.CompositeLit:
				for _, e := range at.Elts {
					if err := v.evalExpr(e); err != nil {
						return err
					}
				}
			}
		}
		if ft, ok := t.Fun.(*ast.SelectorExpr); ok {
			return v.evalSelector(t, ft)
		}
	case *ast.KeyValueExpr:
		return v.evalExpr(t.Value)
	}
	return nil
}

func (v *visitor) evalArgs(expr ast.Expr) error {
	switch at := expr.(type) {
	case *ast.CompositeLit:
		for _, e := range at.Elts {
			if err := v.evalExpr(e); err != nil {
				return err
			}
		}
	// case *ast.BasicLit:
	// fmt.Println("evalArgs", at.Value)
	// v.addBox(at.Value)
	case *ast.CallExpr:
		if at.Fun == nil {
			return nil
		}
		switch st := at.Fun.(type) {
		case *ast.SelectorExpr:
			if err := v.evalSelector(at, st); err != nil {
				return err
			}
		case *ast.Ident:
			return v.evalIdent(st)
		}
		for _, a := range at.Args {
			if err := v.evalArgs(a); err != nil {
				return err
			}
		}
	}
	return nil
}

func (v *visitor) evalSelector(expr *ast.CallExpr, sel *ast.SelectorExpr) error {
	x, ok := sel.X.(*ast.Ident)
	if !ok {
		return nil
	}
	if x.Name == "packr" && sel.Sel.Name == "NewBox" {
		for _, e := range expr.Args {
			switch at := e.(type) {
			case *ast.Ident:
				switch at.Obj.Kind {
				case ast.Var:
					if as, ok := at.Obj.Decl.(*ast.AssignStmt); ok {
						v.addVariable(as)
					}
				case ast.Con:
					if vs, ok := at.Obj.Decl.(*ast.ValueSpec); ok {
						v.addConstant(vs)
					}
				}
				return v.evalIdent(at)
			case *ast.BasicLit:
				v.addBox(at.Value)
			case *ast.CallExpr:
				return v.evalExpr(at)
			}
		}
	}

	return nil
}

func (v *visitor) evalIdent(i *ast.Ident) error {
	if i.Obj == nil {
		return nil
	}
	if s, ok := i.Obj.Decl.(*ast.AssignStmt); ok {
		return v.evalStmt(s)
	}
	return nil
}

func (v *visitor) addBox(b string) {
	b = strings.Replace(b, "\"", "", -1)
	v.Boxes = append(v.Boxes, b)
}

func (v *visitor) addVariable(as *ast.AssignStmt) error {
	if len(as.Rhs) == 1 {
		if bs, ok := as.Rhs[0].(*ast.BasicLit); ok {
			v.addBox(bs.Value)
		}
	}
	return nil
}

func (v *visitor) addConstant(vs *ast.ValueSpec) error {
	if len(vs.Values) == 1 {
		if bs, ok := vs.Values[0].(*ast.BasicLit); ok {
			v.addBox(bs.Value)
		}
	}
	return nil
}
