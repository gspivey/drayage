package cmd

import (
	"fmt"
	"log"

	pb "github.com/drayage/protocol"
	clientpackage "github.com/drayage/protocol/client"
	dv "github.com/drayage/volume"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var all bool

func init() {
	rootCmd.AddCommand(lsCmd)
	lsCmd.PersistentFlags().BoolVarP(&all, "all", "a", false, "Show volume's on all configured nodes")
	viper.BindPFlag("all", tlsCmd.PersistentFlags().Lookup("ls"))
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Show available docker volumes",
	Long:  `Show volumes on localhost and/or all configured nodes. Also shows size of volume's'`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("all value is %v", all)
		if all {
			// 1 Iterate through configured hosts
			//
			serverAddr := "localhost:50051"
			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())
			conn, err := grpc.Dial(serverAddr, opts...)
			if err != nil {
				log.Fatalf("fail to dial: %v", err)
			}
			defer conn.Close()
			client := pb.NewCommsProtoClient(conn)
			v := &pb.Volume{"DoesNotMatter"}
			clientpackage.LSVolume(client, v)

		} else {
			//dv.VolumeExpiriment()
			v, _ := dv.ListVolumes()
			log.Printf("\n result \n %v", v)
			// TODO add filter on driver
			for _, vol := range v {
				fmt.Printf("Name: %s\n MountPath: %s\n Driver: %s\n\n", vol.Name, vol.Mountpoint, vol.Driver)
				vn, sz, up, _ := dv.AnalyzeVolume(vol)
				szh := dv.SIByteFormat(uint64(sz))
				fmt.Printf("Analyzed Volume: %s %s, %v\n", vn, szh, up)
			}
		}

	},
}
