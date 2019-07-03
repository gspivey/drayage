package cmd

import (
	"fmt"
	"log"
	"net"

	pb "github.com/gspivey/drayage/protocol"
	"github.com/gspivey/drayage/protocol/server"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// At some point make this configurable
const (
	port = ":50051"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start drayage",
	Long:  `start drayage with its appropriate configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("")
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Printf("Server started sucessfully on port: %s", port)
		commServer := grpc.NewServer()
		pb.RegisterCommsProtoServer(commServer, &server.Server{})
		// Register reflection service on gRPC server.
		reflection.Register(commServer)
		if err := commServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

	},
}
