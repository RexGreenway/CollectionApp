package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/RexGreenway/CollectionApp/internal/genproto"
)

type collectionService struct {
	genproto.UnimplementedCollectionServiceServer
}

func (s *collectionService) ListCollections(
	req *genproto.ListCollectionsRequest,
	stream genproto.CollectionService_ListCollectionsServer,
) error {
	return nil
}

func (s *collectionService) GetCollection(
	ctx context.Context,
	req *genproto.GetCollectionRequest,
) (*genproto.Collection, error) {
	return &genproto.Collection{
		Items: []*genproto.Item{
			{
				Id:   "tid-1",
				Name: "item-1",
				Info: map[string]string{"dir": "nolan", "year": "2020"},
			},
			{
				Id:   "tid-2",
				Name: "item-2",
				Info: map[string]string{"dir": "tara", "year": "111§"},
			},
		},
	}, nil
}

func (s *collectionService) AddItem(
	ctx context.Context,
	item *genproto.Item,
) (*genproto.Item, error) {
	fmt.Println("hit AddItem method")

	return &genproto.Item{}, nil
}

func StartServer() error {
	lis, err := net.Listen("tcp", "localhost:50100")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	genproto.RegisterCollectionServiceServer(
		grpcServer,
		&collectionService{},
	)

	fmt.Println("starting grpc server on 50100")

	return grpcServer.Serve(lis)
}
