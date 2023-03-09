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
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	pbaoc "github.com/brotherlogic/adventofcode/proto"
	pb "github.com/brotherlogic/aocfinder/proto"
	atpb "github.com/brotherlogic/aoctracker/proto"
	pbghc "github.com/brotherlogic/githubcard/proto"
	rspb "github.com/brotherlogic/rstore/proto"
)

func getClient() (pbaoc.AdventServerServiceClient, error) {
	clientCert, err := tls.LoadX509KeyPair("/home/simon/keys/client.pem", "/home/simon/keys/client.key")
	if err != nil {
		return nil, err
	}

	trustedCert, err := ioutil.ReadFile("/home/simon/keys/cacert.pem")
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

	return pbaoc.NewAdventServerServiceClient(conn), nil
}

func assess(year, day, part int) (bool, error) {
	ctx, cancel := utils.ManualContext("aocfinder", time.Minute*5)
	defer cancel()

	client, err := getClient()
	if err != nil {
		return false, err
	}

	_, err = client.Solve(ctx, &pbaoc.SolveRequest{Year: int32(year), Day: int32(day), Part: int32(part)})
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func raiseIssue(year, day, part int) (int32, error) {
	ctx, cancel := utils.ManualContext("aocfinder-issue", time.Minute)
	defer cancel()

	conn, err := utils.LFDialServer(ctx, "githubcard")
	if err != nil {
		return -1, err
	}

	client := pbghc.NewGithubClient(conn)
	issue, err := client.AddIssue(ctx, &pbghc.Issue{
		Service: "adventofcode",
		Title:   fmt.Sprintf("Solve AOC Puzzle (%v, %v part %v)", year, day, part),
		Body:    "Solve it",
	})
	if err != nil {
		return -1, err
	}
	return issue.Number, nil
}

func main() {
	dctx, dcancel := utils.ManualContext("aocfinder", time.Minute)
	defer dcancel()

	wo := &pb.WorkingOn{}

	conn, err := utils.LFDialServer(dctx, "rstore")
	if err != nil {
		log.Fatalf("Bad dial: %v", err)
	}

	client := rspb.NewRStoreServiceClient(conn)
	rconf, err := client.Read(dctx, &rspb.ReadRequest{Key: "aocfinder/config"})
	if status.Code(err) != codes.NotFound {
		if err != nil {
			log.Fatalf("bad read: %v", err)
		}
		err = proto.Unmarshal(rconf.GetValue().GetValue(), wo)
		if err != nil {
			log.Fatalf("Bad unmarshal: %v", err)
		}
	}

	wo.LastRun = time.Now().Unix()

	conn, err = utils.LFDialServer(dctx, "aocfinder")
	if err != nil {
		log.Fatalf("Bad dial: %v", err)
	}
	aocclient := atpb.NewAOCTrackerServiceClient(conn)
	aocclient.Track(dctx, &atpb.TrackRequest{CurrentYear: wo.GetYear()})

	if wo.CorrespondingIssue > 0 {
		log.Fatalf("Already working on an issue: %v", wo)
	}

	for year := 2015; year <= time.Now().Year(); year++ {
		for day := 1; day <= time.Now().Day(); day++ {
			for part := 1; part <= 2; part++ {
				fmt.Printf("Trying %v, Day %v, Part %v\n", year, day, part)
				val, err := assess(year, day, part)
				if err != nil {
					log.Fatalf("Bad run: %v", err)
				}

				if !val {
					num, err := raiseIssue(year, day, part)
					if err != nil {
						log.Fatalf("Bad issue: %v", err)
					}

					wo.Day = int32(day)
					wo.Part = int32(part)
					wo.Year = int32(year)
					wo.CorrespondingIssue = num

					bytes, err := proto.Marshal(wo)
					if err != nil {
						log.Fatalf("Bad marshal: %v", err)
					}
					tctx, tcancel := utils.ManualContext("aocfinder", time.Minute)
					defer tcancel()
					_, err = client.Write(tctx, &rspb.WriteRequest{Key: "aocfinder/config", Value: &anypb.Any{Value: bytes}})
					if err != nil {
						log.Fatalf("Bad write: %v", err)
					}
				}
			}
		}
	}
}
