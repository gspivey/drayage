package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use: "start",
	Short: "Start drayage",
	Long: `start drayage with its appropriate configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("")
	},
}