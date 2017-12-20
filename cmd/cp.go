package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/spf13/viper"
)

var src string
var dst string
var volume string

func init() {
	rootCmd.AddCommand(cpCmd)
	cpCmd.PersistentFlags().StringVarP(&volume, "volume", "v", "", "Name of docker volume to copy")
	cpCmd.PersistentFlags().StringVarP(&dst, "dst", "d", "", "Hostname of destination node")
	cpCmd.PersistentFlags().StringVarP(&src, "src", "s", "", "Hostname of source node")
	viper.BindPFlag("volume", cpCmd.PersistentFlags().Lookup("volume"))
	viper.BindPFlag("dst", cpCmd.PersistentFlags().Lookup("dst"))
	viper.BindPFlag("src", cpCmd.PersistentFlags().Lookup("src"))
}

var cpCmd = &cobra.Command{
	Use: "cp",
	Short: "Copy volume to and from any two configured nodes",
	Long: `Copy volumes. Run "drayage ls --all to see available nodes and volumes"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("")
	},
}