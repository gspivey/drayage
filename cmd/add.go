package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/spf13/viper"
)

var host string
var newvolume string

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.PersistentFlags().StringVarP(&newvolume, "newvolume", "nv", "", "Name of docker volume to copy")
	addCmd.PersistentFlags().StringVarP(&host, "host", "h", "", "Hostname of where vo;lume will be created")
	viper.BindPFlag("newvolume", cpCmd.PersistentFlags().Lookup("newvolume"))
	viper.BindPFlag("host", cpCmd.PersistentFlags().Lookup("host"))
}

var addCmd = &cobra.Command{
	Use: "host",
	Short: "Add a named docker volume to a configured node",
	Long: `Add a named docker volume to a configured node`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("")
	},
}