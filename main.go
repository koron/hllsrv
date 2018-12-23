package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/axiomhq/hyperloglog"
	"github.com/koron/hllsrv/hlltc"
	"google.golang.org/grpc"
)

var (
	grpcAddr string
)

func main() {
	flag.StringVar(&grpcAddr, "grpc_addr", ":18080",
		`gRPC listening address: [{HOST}]:{PORT}`)
	flag.Parse()
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	srv := &server{}
	s := grpc.NewServer()
	hlltc.RegisterEstimatorServer(s, srv)
	hlltc.RegisterMergerServer(s, srv)
	ln, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return err
	}
	return s.Serve(ln)
}

type server struct{}

func unmarshal(b []byte) (*hyperloglog.Sketch, error) {
	sk := &hyperloglog.Sketch{}
	err := sk.UnmarshalBinary(b)
	if err != nil {
		return nil, err
	}
	return sk, nil
}

func merge(base []byte, others [][]byte) ([]byte, error) {
	acc, err := unmarshal(base)
	if err != nil {
		return nil, err
	}

	for _, o := range others {
		op, err := unmarshal(o)
		if err != nil {
			return nil, err
		}
		err = acc.Merge(op)
		if err != nil {
			return nil, err
		}
	}

	res, err := acc.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *server) Estimate(ctx context.Context, req *hlltc.EstimateRequest) (*hlltc.EstimateReply, error) {
	sk, err := unmarshal(req.GetSketch())
	if err != nil {
		return nil, err
	}
	sum := sk.Estimate()
	return &hlltc.EstimateReply{Sum: sum}, nil
}

func (s *server) Merge(ctx context.Context, req *hlltc.MergeRequest) (*hlltc.MergeReply, error) {
	b, err := merge(req.GetSketch(), req.GetSketches())
	if err != nil {
		return nil, err
	}
	return &hlltc.MergeReply{MergedSketch: b}, nil
}
