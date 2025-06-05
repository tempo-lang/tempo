package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var disableTerminalColor *bool

var rootCmd = &cobra.Command{
	Use:   "tempo",
	Short: "Tempo choreographic programming language",
	Long:  `The compiler for Tempo, a practical choreographic programming language.`,
}

func init() {
	disableTerminalColor = rootCmd.Flags().Bool("nocolor", false, "Disable colors in the terminal")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
