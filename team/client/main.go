package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"io"
	"log"
	"math"
	"sync"
	"team/api/proto/armydevice"
	"github.com/montanaflynn/stats"
	"gorm.io/gorm"
)

var (
	addr = flag.String("addr", "localhost:8080", "Address to connect to")
	k = flag.Float64("k", 0.999, "Anomaly coefficient")
)

var pool = sync.Pool{
	New: func() any {
		return []float64{}
	},
}

type Anomalies struct {
	SessionId string
	Frequency float64
	Timestamp string
}


func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	c := armydevice.NewArmyDeviceClient(conn)
	req := &armydevice.Empty{}

	resStream, err := c.GetArmyDeviceInfo(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
	}

	var samples []float64
	f := 0
	for i := 0; i < 100; i++ {
		samples = pool.Get().([]float64)
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if f == 0 {
			log.Println("Current session id:", msg.SessionId)
			log.Println("Start time:", msg.Timestamp.AsTime())
			f = 1
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		samples = append(samples, msg.Frequency)
		pool.Put(samples)
	}

	samples = pool.Get().([]float64)
	PredictedMean, err := stats.Mean(samples)
	if err != nil {
		log.Fatalln(err)
	}
	PredictedSTD, err := stats.StandardDeviation(samples)
	pool.Put(samples)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Mean:", PredictedMean, "STD:", PredictedSTD)

	dsn := "host=localhost user=hardella password=123 dbname=golang port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to database")
	}
	err = db.AutoMigrate(&Anomalies{})
	if err != nil {
		log.Fatalln(err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if math.Abs(msg.Frequency - PredictedMean) > PredictedSTD * *k {
			anomaly := Anomalies{
				SessionId: msg.GetSessionId(),
				Frequency: msg.Frequency,
				Timestamp: msg.GetTimestamp().AsTime().String(),
			}
			db.Create(&anomaly)
			log.Printf("Anomaly %v\n", msg.Frequency)
		} else {
			log.Println("Not anomaly:", msg.Frequency)
		}
	}
}