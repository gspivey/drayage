package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(tlsCmd)
	tlsCmd.PersistentFlags().Bool("init", true, "use to generate required tls files")
	tlsCmd.PersistentFlags().Bool("insecure", false, "use to run drayage over HTTP without TLS")
	viper.BindPFlag("init", tlsCmd.PersistentFlags().Lookup("init"))
	viper.BindPFlag("insecure", tlsCmd.PersistentFlags().Lookup("insecure"))
}

var tlsCmd = &cobra.Command{
	Use: "tls",
	Short: "Create all files needed for each configured host for encryption and authentication",
	Long: `Run tls init to generate a CA and host certifactes for each listed host`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("")
	},
}
