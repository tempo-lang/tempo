package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/spf13/cobra"
	"github.com/tempo-lang/tempo/compiler"
	"github.com/tempo-lang/tempo/type_check/type_error"
)

var packageName string
var targetLang TargetLangFlag
var runtimePath string

type TargetLangFlag string

const ()

// Set implements pflag.Value.
func (t *TargetLangFlag) Set(v string) error {
	v = strings.ToLower(v)
	switch v {
	case "go", "ts", "js", "java":
		*t = TargetLangFlag(v)
		return nil
	default:
		return errors.New(`must be one of "go", "ts", "js", "java"`)
	}
}

// String implements pflag.Value.
func (t *TargetLangFlag) String() string {
	return string(*t)
}

// Type implements pflag.Value.
func (t *TargetLangFlag) Type() string {
	return "language"
}

var buildCmd = &cobra.Command{
	Use:     "build [flags] file",
	Example: "  tempo build choreography.tempo",
	Short:   "Compile tempo source code",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("missing source file")
		}

		if len(args) > 1 {
			return fmt.Errorf("accepts 1 arg, received %d", len(args))
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := path.Clean(args[0])

		input, err := antlr.NewFileStream(inputFile)
		if err != nil {
			fmt.Printf("failed to get file stream: %v\n", err)
			os.Exit(1)
		}

		filename := path.Base(inputFile)
		if packageName == "" {
			packageName = filename[0 : len(filename)-len(path.Ext(filename))]
		}
		options := compiler.Options{
			PackageName: packageName,
			Language:    compiler.CompilerLanguage(targetLang),
			RuntimePath: runtimePath,
		}

		output, errors := compiler.Compile(input, &options)
		if errors != nil {
			for _, err := range errors {
				if typeErr, ok := err.(type_error.Error); ok {
					type_error.FormatError(os.Stdout, input, typeErr, !*disableTerminalColor)
				} else {
					fmt.Printf("%v\n", err)
				}
			}
			os.Exit(1)
		}

		fmt.Printf("%s", output)
	},
}

func init() {
	buildCmd.Flags().StringVarP(&packageName, "package", "p", "choreography", "The Go package name for the generated code")
	buildCmd.Flags().VarP(&targetLang, "lang", "l", `Target language, allowed: "go", "ts"`)
	buildCmd.Flags().StringVarP(&runtimePath, "runtime", "r", "@tempo-lang/tempo/runtime", "The path to the Typescript runtime module")
	rootCmd.AddCommand(buildCmd)
}
