package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var Ast string = "*"
var AstAst string = "**"
var Negation string = "!"

func CopyFile(src, dest string) error {
	srcFile, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	_, err = destFile.Write(srcFile)
	return err
}

// ** - all nested folders and their content
// * - part of filename
// TODO rewrite it using https://golang.org/pkg/path/filepath/#Glob
func FindAll(patterns ...string) ([]string, error) {
	res := make([]string, 0)
	regs := make([]string, 0)

	for _, p := range patterns {
		regs = append(regs, getRegExp(p))
	}

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		include := false
		for _, reg := range regs {
			if strings.HasPrefix(reg, Negation) {
				var exclude bool
				exclude, err = regexp.MatchString(reg[1:], path)
				if exclude {
					include = false
				}
			} else {
				include, err = regexp.MatchString(reg, path)
			}
			if err != nil {
				return err
			}
		}
		if include {
			res = append(res, path)
		}
		return nil
	})

	return res, err
}

func getRegExp(pattern string) string {
	ast := `[^\/]{0,}`
	astast := ".{0,}"
	exclude := strings.HasPrefix(pattern, Negation)

	if exclude {
		pattern = pattern[1:]
	}

	r := strings.Replace(pattern, "\\", `\\`, -1)
	r = strings.Replace(pattern, "/", `\/`, -1)
	r = strings.Replace(pattern, AstAst, astast, -1)
	r = strings.Replace(r, Ast, ast, -1)

	if exclude {
		return fmt.Sprintf("%s^%s$", Negation, r)
	}
	return fmt.Sprintf("^%s$", r)
}
