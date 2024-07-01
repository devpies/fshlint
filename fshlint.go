// Package fshlint provides a linter for FSH (FHIR Shorthand) files, leveraging
// the ANTLR language recognizer. Unlike SUSHI, which also uses ANTLR but compiles
// FSH files into FHIR artifacts, fshlint is a simple CLI tool that checks .fsh files
// for syntax and semantic issues without compilation. This makes it ideal for quick
// validations in CI pipelines or integration into a project's Makefile for automated
// checks alongside tests.
//
// ANTLR uses grammar definitions in *.g4 files: FSHLexer.g4 for the lexer and FSH.g4
// for the parser. These grammars generate the Go files utilized by fshlint. The lexer
// identifies tokens in the FSH input, while the parser builds an abstract syntax tree
// (AST) to enforce grammar rules. Both grammar files are sourced directly from the
// SUSHI project. To ensure compatibility with updates to the FSH language, the grammar
// files in fshlint's antlr directory must be updated whenever the FSH specification
// changes.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/devpies/fshlint/internal/parser"
)

const (
	colorReset    = "\033[0m"
	colorBold     = "\033[1m"
	colorLightRed = "\033[91m"
	singleSpace   = " "
)

// Custom logger without timestamp.
var customLogger = log.New(os.Stderr, "", 0)

// FshErrorListener is a custom error listener for ANTLR that collects syntax errors during parsing.
// It embeds the default ANTLR error listener and adds a slice to store error messages.
type FshErrorListener struct {
	*antlr.DefaultErrorListener
	errors []string
}

// NewFshErrorListener creates and returns a new instance of FshErrorListener.
// The returned listener will have an empty slice ready to collect error messages.
func NewFshErrorListener() *FshErrorListener {
	return &FshErrorListener{errors: []string{}}
}

// SyntaxError is a custom implementation of the ANTLR error listener's SyntaxError method.
// It captures syntax errors encountered during the parsing process and stores them in the errors slice.
// Parameters:
// - _ antlr.Recognizer: The recognizer where the error occurred (not used here).
// - _ any: The offending symbol that caused the error (not used here).
// - line: The line number where the error occurred.
// - column: The column number where the error occurred.
// - msg: The error message.
// - _ antlr.RecognitionException: The exception thrown by the parser (not used here).
func (el *FshErrorListener) SyntaxError(_ antlr.Recognizer, _ any,
	line, column int, msg string, _ antlr.RecognitionException) {
	err := fmt.Sprintf("line %d:%d %s", line, column, msg)
	el.errors = append(el.errors, err)
}

func lintFile(filePath string) ([]string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	input := antlr.NewInputStream(string(content))
	lexer := parser.NewFSHLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewFSHParser(stream)

	errorListener := NewFshErrorListener()
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)

	p.Doc()

	return errorListener.errors, nil
}

func getLineContent(filePath string, line int) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentLine := 1
	for scanner.Scan() {
		if currentLine == line {
			return scanner.Text(), nil
		}
		currentLine++
	}
	return "", fmt.Errorf("line %d not found in file %s", line, filePath)
}

func formatError(filePath string, line, column int, msg string) string {
	lineContent, err := getLineContent(filePath, line)
	if err != nil {
		return fmt.Sprintf("%s: %s\n", filePath, msg)
	}
	indicator := strings.Repeat(singleSpace, column) + "^"
	return fmt.Sprintf("%s%s%s:%d:%d: %s%s%s\n%s\n%s\n",
		colorBold, filePath, colorReset, line, column,
		colorLightRed, msg, colorReset,
		lineContent,
		indicator,
	)
}

func printHelp() {
	helpText := `Usage: fshlint [OPTIONS] <path-to-fsh-file-or-directory>

Options:
  --help, -h    Show this help message

Examples:
  fshlint ./path/to/files
  fshlint ./path/to/specific/file.fsh
`
	fmt.Println(helpText)
}

func main() {
	help := flag.Bool("help", false, "Show this help message")
	h := flag.Bool("h", false, "Show this help message")
	flag.Parse()

	if *help || *h {
		printHelp()
		return
	}

	path := os.Args[1]

	// Check if the provided path exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		suggestion := "Please provide a valid file or directory path."
		customLogger.Fatalf("Error: The specified path '%s' does not exist. %s\n", path, suggestion)
	}

	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".fsh") {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		customLogger.Fatalf("Error walking the path '%s': %v\n", path, err)
	}

	exitCode := 0

	for _, file := range files {
		errors, err := lintFile(file)
		if err != nil {
			customLogger.Printf("Error linting file %s: %v\n", file, err)
			continue
		}
		if len(errors) > 0 {
			exitCode = 1
		}
		for _, e := range errors {
			parts := strings.Split(e, singleSpace)
			lineCol := strings.Split(parts[1], ":")
			line := lineCol[0]
			column := lineCol[1]
			msg := strings.Join(parts[2:], singleSpace)
			lineNum := 0
			columnNum := 0
			fmt.Sscanf(line, "%d", &lineNum)
			fmt.Sscanf(column, "%d", &columnNum)
			formattedError := formatError(file, lineNum, columnNum, msg)
			fmt.Print(formattedError)
		}
	}

	os.Exit(exitCode)
}
