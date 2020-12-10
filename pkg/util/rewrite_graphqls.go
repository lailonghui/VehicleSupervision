package util

import (
	"bytes"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/formatter"
	"github.com/vektah/gqlparser/v2/parser"
	"io/ioutil"
	"log"
	"os"
)

// 按格式，重新生成graphqls文件
func RewriteGraphqls(filename string) {
	name := "test"
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	doc, gqlErr := parser.ParseSchema(&ast.Source{
		Name:  name,
		Input: string(bs),
	})
	if gqlErr != nil {
		log.Fatal(gqlErr)
	}
	for _, define := range doc.Definitions {

		switch define.Kind {
		case ast.Object, ast.InputObject:
			define.Name = strcase.ToCamel(define.Name)
			for _, field := range define.Fields {
				//field.Name = strcase.ToLowerCamel(field.Name)
				if field.Arguments != nil {
					for _, argument := range field.Arguments {
						if argument.Type.Elem != nil {
							argument.Type.Elem.NamedType = strcase.ToCamel(argument.Type.Elem.NamedType)
						} else {
							argument.Type.NamedType = strcase.ToCamel(argument.Type.NamedType)

						}
					}
				}
				if field.Type != nil {
					if field.Type.Elem != nil {
						field.Type.Elem.NamedType = strcase.ToCamel(field.Type.Elem.NamedType)
					} else {
						field.Type.NamedType = strcase.ToCamel(field.Type.NamedType)
					}
				}

			}
			//define.Name = "a_" + define.Name

		case ast.Enum:
			define.Name = strcase.ToCamel(define.Name)
		case ast.Scalar:
			define.Name = strcase.ToCamel(define.Name)
		}
	}
	for _, define := range doc.Extensions {

		switch define.Kind {
		case ast.Object, ast.InputObject:
			define.Name = strcase.ToCamel(define.Name)
			for _, field := range define.Fields {
				//field.Name = strcase.ToLowerCamel(field.Name)
				if field.Arguments != nil {
					for _, argument := range field.Arguments {
						if argument.Type.Elem != nil {
							argument.Type.Elem.NamedType = strcase.ToCamel(argument.Type.Elem.NamedType)
						} else {
							argument.Type.NamedType = strcase.ToCamel(argument.Type.NamedType)

						}
					}
				}
				if field.Type != nil {
					if field.Type.Elem != nil {
						field.Type.Elem.NamedType = strcase.ToCamel(field.Type.Elem.NamedType)
					} else {
						field.Type.NamedType = strcase.ToCamel(field.Type.NamedType)
					}
				}

			}
			//define.Name = "a_" + define.Name

		case ast.Enum:
			define.Name = strcase.ToCamel(define.Name)
		case ast.Scalar:
			define.Name = strcase.ToCamel(define.Name)
		}
	}
	// exec format
	var buf bytes.Buffer
	f := formatter.NewFormatter(&buf)
	f.FormatSchemaDocument(doc)

	// validity check
	doc, gqlErr = parser.ParseSchema(&ast.Source{
		Name:  name,
		Input: buf.String(),
	})
	fmt.Println(buf.String())
	if gqlErr != nil {
		log.Fatal(gqlErr)
	}
	ioutil.WriteFile(filename, buf.Bytes(), os.ModePerm)
}
