package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(lsCmd)
	lsCmd.PersistentFlags().Bool("all", false, "Show volume's on all configured nodes")
	viper.BindPFlag("all", tlsCmd.PersistentFlags().Lookup("ls"))
}

var lsCmd = &cobra.Command{
	Use: "ls",
	Short: "Show available docker volumes",
	Long: `Show volumes on localhost and/or all configured nodes. Also shows size of volume's'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("")
	},
}