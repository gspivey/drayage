package cmd

import (
	"log"

	pb "github.com/gspivey/drayage/protocol"
	clientpackage "github.com/gspivey/drayage/protocol/client"
	dv "github.com/gspivey/drayage/volume"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var host string
var newvolume string

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.PersistentFlags().StringVarP(&newvolume, "newvolume", "n", "", "Name of docker volume to create")
	addCmd.PersistentFlags().StringVarP(&host, "host", "o", "", "Hostname of where volume will be created")
	err := viper.BindPFlag("newvolume", cpCmd.PersistentFlags().Lookup("newvolume"))
	if err != nil {
		panic(err)
	}
	err = viper.BindPFlag("host", cpCmd.PersistentFlags().Lookup("host"))
	if err != nil {
		panic(err)
	}
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a named docker volume to a configured node",
	Long:  `Add a named docker volume to a configured node`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Variables are %s %s", host, newvolume)
		if "" == host {
			// Add volume will create a volume when it is given an empty string
			// These volume's will have the default hash name
			v, err := dv.AddVolume(newvolume)
			if err != nil {
				panic(err)
			}
			// FIXME update log message to stop printing internal map info
			log.Printf("Returned volume %v", v)
		} else {
			log.Printf("Creating volume on host: %v", host)
			// FIXME Replace with efficient string concatenation`
			serverAddr := host + ":50051"
			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())
			conn, err := grpc.Dial(serverAddr, opts...)
			if err != nil {
				log.Fatalf("fail to dial: %v", err)
			}
			defer conn.Close()
			client := pb.NewCommsProtoClient(conn)
			v := &pb.Volume{newvolume}
			clientpackage.AddVolume(client, v)
		}
	},
}
