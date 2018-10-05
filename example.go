package main

import (
	"bytes"
	"fmt"
	"go/printer"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/tools/godoc"
)

func exampleCodeFunc(info *godoc.PageInfo, funcName string) string {
	var buf bytes.Buffer
	for _, eg := range info.Examples {
		name := stripExampleSuffix(eg.Name)

		if name != funcName {
			continue
		}

		// print code
		cnode := &printer.CommentedNode{Node: eg.Code, Comments: eg.Comments}
		config := &printer.Config{Mode: printer.UseSpaces, Tabwidth: 4}
		var buf1 bytes.Buffer
		config.Fprint(&buf1, info.FSet, cnode)
		code := buf1.String()

		// Additional formatting if this is a function body. Unfortunately, we
		// can't print statements individually because we would lose comments
		// on later statements.
		if n := len(code); n >= 2 && code[0] == '{' && code[n-1] == '}' {
			code = code[1 : n-1]
		}
		code = strings.Trim(code, "\n")
		code = stripLeadingSpaces(code)
		buf.WriteString(code)
	}
	return buf.String()
}

func exampleMdFunc(info *godoc.PageInfo, funcName string, prefix string) string {
	s := exampleCodeFunc(info, funcName)
	if s == "" {
		return s
	}
	s = fmt.Sprintf("```GO\n%s\n```\n", s)
	if prefix != "" {
		s = prefix + "\n" + s
	}
	return s
}

func startsWithUppercase(s string) bool {
	r, _ := utf8.DecodeRuneInString(s)
	return unicode.IsUpper(r)
}

func stripExampleSuffix(name string) string {
	if i := strings.LastIndex(name, "_"); i != -1 {
		if i < len(name)-1 && !startsWithUppercase(name[i+1:]) {
			name = name[:i]
		}
	}
	return name
}

func stripLeadingSpaces(s string) string {
	lines := strings.Split(s, "\n")
	nmax := 1000
	for _, ln := range lines {
		n := firstNonSpace(ln)
		if n < 0 {
			continue
		}
		if n == 0 {
			return s
		}
		if n < nmax {
			nmax = n
		}
	}
	for i := range lines {
		if len(lines[i]) >= nmax {
			lines[i] = lines[i][nmax:]
		}
	}
	return strings.Join(lines, "\n")
}

func firstNonSpace(s string) int {
	for i, c := range s {
		if c != ' ' {
			return i
		}
	}
	return -1
}
