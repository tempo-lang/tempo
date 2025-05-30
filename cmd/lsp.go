package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tempo-lang/tempo/lsp"
)

var lspCmd = &cobra.Command{
	Use:   "lsp",
	Short: "Start the Tempo Language Server",
	Run: func(cmd *cobra.Command, args []string) {
		lsp.StartServer()
	},
}

func init() {
	rootCmd.AddCommand(lspCmd)
}
