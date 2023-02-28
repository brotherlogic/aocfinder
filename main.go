package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/brotherlogic/goserver/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	pb "github.com/brotherlogic/adventofcode/proto"
)

func getClient() (pb.AdventServerServiceClient, error) {
	clientCert, err := tls.LoadX509KeyPair("/home/brotherlogic/keys/client.pem", "/home/brotherlogic/keys/client.key")
	if err != nil {
		return nil, err
	}

	trustedCert, err := ioutil.ReadFile("/home/brotherlogic/keys/cacert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(trustedCert) {
		return nil, fmt.Errorf("Unable to append cert")
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
		MinVersion:   tls.VersionTLS13,
		MaxVersion:   tls.VersionTLS13,
	}

	// Create a new TLS credentials based on the TLS configuration
	cred := credentials.NewTLS(tlsConfig)

	conn, err := grpc.Dial("adventofcode.brotherlogic-backend.com:80", grpc.WithTransportCredentials(cred))
	if err != nil {
		return nil, err
	}

	return pb.NewAdventServerServiceClient(conn), nil
}

func assess(year, day, part int) (bool, error) {
	ctx, cancel := utils.ManualContext("aocfinder", time.Minute*5)
	defer cancel()

	client, err := getClient()
	if err != nil {
		return false, err
	}

	_, err = client.Solve(ctx, &pb.SolveRequest{Year: int32(year), Day: int32(day), Part: int32(part)})
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func main() {
	for year := 2015; year <= time.Now().Year(); year++ {
		for day := 1; day <= time.Now().Day(); day++ {
			for part := 1; part <= 2; part++ {
				fmt.Printf("Trying %v, Day %v, Part %v\n", year, day, part)
				val, err := assess(year, day, part)
				if err != nil {
					log.Fatalf("Bad run: %v", err)
				}

				if !val {
					fmt.Printf("FOUND\n")
					return
				}
			}
		}
	}
}
