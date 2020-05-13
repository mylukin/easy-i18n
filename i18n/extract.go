package i18n

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Extract messages
func Extract(paths []string, outFile string) error {
	if len(paths) == 0 {
		paths = []string{"."}
	}
	messages := map[string]string{}
	for _, path := range paths {
		if err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			if filepath.Ext(path) != ".go" {
				return nil
			}

			// Don't extract from test files.
			if strings.HasSuffix(path, "_test.go") {
				return nil
			}
			buf, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			fset := token.NewFileSet()
			file, err := parser.ParseFile(fset, path, buf, parser.AllErrors)
			if err != nil {
				return err
			}

			// fmt.Printf("Extract %+v ...\n", path)
			i18NPackName := i18nPackageName(file)
			// ast.Print(fset, file)
			ast.Inspect(file, func(n ast.Node) bool {
				switch v := n.(type) {
				case *ast.CallExpr:
					if fn, ok := v.Fun.(*ast.SelectorExpr); ok {
						var packName string
						if pack, ok := fn.X.(*ast.Ident); ok {
							packName = pack.Name
						}
						funcName := fn.Sel.Name
						namePos := fset.Position(fn.Sel.NamePos)
						// Package name must be equal
						if len(packName) > 0 && i18NPackName == packName {
							// Function name must be equal
							if funcName == "Printf" || funcName == "Sprintf" || funcName == "Fprintf" {
								fmt.Printf("Extract %+v %v.%v ...\n", namePos, packName, funcName)
								// Find the string to be translated
								if str, ok := v.Args[1].(*ast.BasicLit); ok {
									id := strings.Trim(str.Value, "\"`")
									if _, ok := messages[id]; !ok {
										messages[id] = id
									}
								}
							}
							if funcName == "Plural" {
								fmt.Printf("Extract %+v %v.%v ...\n", namePos, packName, funcName)
								// Find the string to be translated
								for i := 0; i < len(v.Args); {
									if i++; i >= len(v.Args) {
										break
									}
									if str, ok := v.Args[i].(*ast.BasicLit); ok {
										id := strings.Trim(str.Value, "\"`")
										if _, ok := messages[id]; !ok {
											messages[id] = id
										}
									}
									i++
								}
							}

						}
					}
				}
				return true
			})
			return nil
		}); err != nil {
			return err
		}
	}

	var content []byte
	var err error
	of := strings.ToLower(outFile)
	if strings.HasSuffix(of, ".json") {
		content, err = marshal(messages, "json")
	}
	if strings.HasSuffix(of, ".toml") {
		content, err = marshal(messages, "toml")
	}
	if strings.HasSuffix(of, ".yaml") {
		content, err = marshal(messages, "yaml")
	}
	if err != nil {
		return err
	}
	err = os.MkdirAll(path.Dir(outFile), os.ModePerm)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(outFile, content, os.ModePerm)
	if err != nil {
		return err
	}
	fmt.Printf("Extract to %v ...\n", outFile)
	return nil
}

func i18nPackageName(file *ast.File) string {
	for _, i := range file.Imports {
		if i.Path.Kind == token.STRING && i.Path.Value == `"github.com/rain-31/easy-i18n/i18n"` {
			if i.Name == nil {
				return "i18n"
			}
			return i.Name.Name
		}
	}
	return ""
}
