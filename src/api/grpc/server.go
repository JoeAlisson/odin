package grpc

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/joealisson/odin/api/grpc/ptypes"
	"github.com/joealisson/odin/ingestion"
)

type server struct {
	ptypes.UnimplementedOdinServer
	ingestor ingestion.Ingestor
}

func (s *server) Audit(ctx context.Context, req *ptypes.AuditRequest) (*emptypb.Empty, error) {
	if err := s.ingestor.Ingest(ctx, convert(req)); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func NewServer(ingestor ingestion.Ingestor) *server {
	return &server{
		ingestor: ingestor,
	}
}

func convert(req *ptypes.AuditRequest) *ingestion.AuditLog {
	return &ingestion.AuditLog{
		System: req.System,
		User: ingestion.User{
			UID:      req.User.Uid,
			Username: req.User.Username,
			IP:       req.User.Ip,
		},
		Changes: ingestion.Changes{
			Entity:      req.Changes.Entity,
			Description: req.Changes.Description,
			Kind:        ingestion.Kind(req.Changes.Kind),
			Values:      req.Changes.Values,
		},
	}
}
