package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"path/filepath"
)

func main() {
	var folderName string
	flag.StringVar(&folderName, "f", "", "command flag for specified folder")
	flag.Parse()

	if folderName == "" {
		folderName, _ = os.Getwd()
		folderName += string(filepath.Separator)
	}
	paths, _ := getFileNames(folderName)

	for _, v := range paths {
		fmt.Println(v)
		f := getAst(v)
		traverseAst(f)
	}
}

func getFileNames(folderName string) ([]string, error) {
	var filePaths []string
	err := filepath.Walk(folderName, func(filePath string, f os.FileInfo, err error) error {
		ext := path.Ext(filePath)
		if ext == ".go" {
			filePaths = append(filePaths, filePath)
		}
		return nil
	})
	return filePaths, err
}

func getAst(fullPath string) *ast.File {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, fullPath, nil, 0)
	if err != nil {
		panic(err)
	}
	return file

}

func traverseAst(file *ast.File) {
	ast.Inspect(file, func(node ast.Node) bool {
		switch x := node.(type) {
		case *ast.FuncDecl:
			fmt.Println("### function Decl")
			fmt.Println("NAME: ", x.Name)
			if x.Recv != nil {
				fmt.Println(x.Recv.List[0].Type)
			}
		case *ast.CallExpr:
			fmt.Println("### function Call")
			fmt.Println(x.Fun)
		}
		return true
	})
}
