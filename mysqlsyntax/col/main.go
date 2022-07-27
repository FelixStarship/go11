package main

import (
	"fmt"

	"github.com/pingcap/tidb/parser"
	"github.com/pingcap/tidb/parser/ast"
	_ "github.com/pingcap/tidb/parser/test_driver"
)

type colX struct {
	colNames []string
}

func (v *colX) Enter(in ast.Node) (ast.Node, bool) {
	if name, ok := in.(*ast.DeleteStmt); ok {
		if name.Where != nil {
			fmt.Println("aaaaa", name.Where.GetType())
		}
		fmt.Println(name.Where)
	}
	return in, false
}

func (v *colX) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func extract(rootNode *ast.StmtNode) []string {
	v := &colX{}
	(*rootNode).Accept(v)
	return v.colNames
}

func parse(sql string) (*ast.StmtNode, error) {
	p := parser.New()

	stmtNodes, _, err := p.Parse(sql, ";", "")
	if err != nil {
		return nil, err
	}

	return &stmtNodes[0], nil
}

func main() {
	astNode, err := parse("delete from")
	if err != nil {
		fmt.Printf("parse error: %v\n", err.Error())
		return
	}
	fmt.Printf("%v\n", extract(astNode))
	fmt.Printf("%v\n", *astNode)
}
