package rewrite

import (
	"context"

	"google.golang.org/grpc"

	pb "backend/gen"
)

type Service struct {
	grpcClient pb.RewriterClient
}

func NewService(grpcConn *grpc.ClientConn) *Service {
	return &Service{
		grpcClient: pb.NewRewriterClient(grpcConn),
	}
}

func (s *Service) RewriteText(ctx context.Context, text, style string) (*RewriteResponse, error) {
	resp, err := s.grpcClient.Rewrite(ctx, &pb.RewriteRequest{
		Text:  text,
		Style: style,
	})
	if err != nil {
		return nil, err
	}

	return &RewriteResponse{
		OriginalText:  resp.OriginalText,
		RewrittenText: resp.RewrittenText,
		AppliedStyle:  resp.AppliedStyle,
	}, nil
}
