package cmd

import (
	"github.com/gspivey/drayage/archive"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var pickvolume string

func init() {
	rootCmd.AddCommand(tarCmd)

	tarCmd.PersistentFlags().StringVarP(&pickvolume, "pickvolume", "p", "", "Name of docker volume to tar")
	err := viper.BindPFlag("pickvolume", tarCmd.PersistentFlags().Lookup("pickvolume"))
	if err != nil {
		panic(err)
	}
}

var tarCmd = &cobra.Command{
	Use:   "tar",
	Short: "Tar available docker volumes. Specify volume or tar all",
	Long:  `Tar available docker volumes. Run "drayage ls --all to see available nodes and volumes`,
	Run: func(cmd *cobra.Command, args []string) {
		archive.WriteToTar(pickvolume)
	},
}
