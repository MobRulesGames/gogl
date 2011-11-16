// OpenGL Binding generator

package main

import (
	"fmt"
	"flag"
)

var (
	download    *bool   = flag.Bool("download", false, "Download spec files from Khronos registry first.")
	version     *string = flag.String("version", "2.1", "OpenGL version")
	deprecation *bool   = flag.Bool("deprecate", false, "Exclude all deprecated features.")
	outGLFile   *string = flag.String("outgl", "gl.go", "Output file name")
	// TODO: add flags for additional extensions
)

func enumFilter(category string, enum *Enum) bool {
	// TODO: filter by version/extension use flags
	return true
}

func functionFilter(category string, function *Function) bool {
	// TODO: filter by version/extension use flags
	return true
}

func main() {
	fmt.Printf("OpenGL binding generator for Go. Copyright (c) 2011 by Christoph Schunk\n")
	flag.Parse()

	if *download {
		DownloadOpenGLSpecs()
	}

	fmt.Printf("Parsing enumext.spec file...\n")
	enumCategories, err := ReadEnumsFromFile(OpenGLEnumExtSpecFile)
	if err != nil {
		panic(err.String())
	}

	fmt.Printf("Parsing gl.tm file ...\n")
	typeMap, err := ReadTypeMapFromFile(OpenGLTypeMapFile)
	if err != nil {
		panic(err.String())
	}

	fmt.Printf("Parsing gl.spec file ...\n")
	functions, supportedVersions, err := ReadFunctionsFromFile(OpenGLSpecFile)
	if err != nil {
		panic(err.String())
	}

	// TODO: This output is temporary for debugging
	fmt.Println("Supported versions:")
	fmt.Println(supportedVersions)

	fmt.Println("Enums:")
	for category, enums := range enumCategories {
		fmt.Printf("  %v\n", category)
		for _, enum := range enums {
			fmt.Printf("    %v = %v\n", enum.Name, enum.Value)
		}
	}
	fmt.Println("Types:")
	for abstractType, cType := range typeMap {
		fmt.Printf("  %v -> %v\n", abstractType, cType)
	}
	fmt.Println("Functions:")
	for category, functions := range functions {
		fmt.Printf("  %v\n", category)
		for _, function := range functions {
			fmt.Printf("    %v\n", function.Name)
			if function.Version.Valid() {
				fmt.Printf("      Version: %v\n", function.Version)
			}
			if function.DeprecatedVersion.Valid() {
				fmt.Printf("      Deprecated Version: %v\n", function.DeprecatedVersion)
			}
			fmt.Printf("      Return Type: %v\n", function.Return)
			if len(function.Parameters) > 0 {
				fmt.Printf("      Parameters:\n")
				for _, param := range function.Parameters {
					if param.InArray {
						fmt.Printf("        %v %v in array\n", param.Name, param.Type)
					} else {
						fmt.Printf("        %v %v\n", param.Name, param.Type)
					}
				}
			} else {
				fmt.Printf("      0 Parameters\n")
			}
		}
	}

	// TODO: just a test, do real unit testing
	//functions["Cat1"] = []Function{
	//Function{Name: "Foo1", Parameters: []Parameter{Parameter{"p1", "int"}, Parameter{"p2", "int"}}, Return: "void"},
	//Function{Name: "Foo2", Parameters: []Parameter{Parameter{"p1", "int"}, Parameter{"p2", "int"}, Parameter{"p3", "float"}}, Return: "void"},
	//Function{Name: "Foo3", Parameters: []Parameter{Parameter{"p1", "int"}, Parameter{"p2", "int"}, Parameter{"p3", "float"}}, Return: "void"},
	//}

	//Generate(*outGLFile, enumCategories, functions, typeMap, enumFilter, functionFilter)
}
