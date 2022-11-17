package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"math/rand"
	"net"
	"team/api/proto/armydevice"
	"github.com/google/uuid"
	"time"
)

var (
	port = flag.Int("port", 8080, "The server port")
	minMean = -10
	maxMean = 10

	minSTD = 0.3
	maxSTD = 1.5
)

type Server struct {
	armydevice.UnimplementedArmyDeviceServer
}

func (s *Server) GetArmyDeviceInfo (req  *armydevice.Empty, stream armydevice.ArmyDevice_GetArmyDeviceInfoServer) error {
	randMean := float64(minMean) + rand.Float64() * (float64(maxMean) - float64(minMean))
	randStd := minSTD + rand.Float64() * (maxSTD - minSTD)

	log.Println("Random mean value:", randMean, "Random std value: ", randStd)

	uuid := uuid.New()
	log.Println("UUID generated", uuid)
	for {
		t := time.Now().UTC()
		timeUTC := timestamppb.New(t)
		log.Println("Current timestamp in UTC", timeUTC)
		log.Println("Current session", uuid)

		sample := rand.NormFloat64() * randStd + randMean
		log.Println("Sample:", sample)
		err := stream.Send(&armydevice.AddResponse{
			SessionId: uuid.String(),
			Frequency: sample,
			Timestamp: timeUTC,
		})
		if err != nil {
			return err
		}
		time.Sleep(100 * time.Millisecond)
	}
	return nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	armydevice.RegisterArmyDeviceServer(s, &Server{})
	log.Printf("Server is listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}