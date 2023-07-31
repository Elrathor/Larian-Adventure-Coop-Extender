package main

import (
	"context"
	"flag"
	"log"
	"runtime"
	"sync"
	"time"

	pb "github.com/Elrathor/Larian-Adventure-Coop-Extender/lace"
	"github.com/micmonay/keybd_event"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewExchangeClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetCommand(ctx, &pb.GetCommandRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// Initialize Keyboard
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	// For linux, it is very important to wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	// Select keys to be pressed
	kb.SetKeys(keybd_event.VK_A)
	kb.HasSHIFT(false)

	ticker := time.NewTicker(time.Second)
	quit := make(chan struct{})
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		for {
			select {
			case <-ticker.C:
				command := r.GetCommand()
				if command == "SAVE" {
					// Press the selected keys
					err = kb.Launching()
					if err != nil {
						panic(err)
					}
				}

			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	wg.Wait()
}
