package objectpool

import (
	"context"

	"gitlab.com/gitlab-org/gitaly-proto/go/gitalypb"
	"gitlab.com/gitlab-org/gitaly/internal/helper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) LinkRepositoryToObjectPool(ctx context.Context, req *gitalypb.LinkRepositoryToObjectPoolRequest) (*gitalypb.LinkRepositoryToObjectPoolResponse, error) {
	if req.GetRepository() == nil {
		return nil, status.Error(codes.InvalidArgument, "no repository")
	}

	pool, err := poolForRequest(req)
	if err != nil {
		return nil, err
	}

	if err := pool.Link(ctx, req.GetRepository()); err != nil {
		return nil, status.Error(codes.Internal, helper.SanitizeString(err.Error()))
	}

	return &gitalypb.LinkRepositoryToObjectPoolResponse{}, nil
}

func (s *server) UnlinkRepositoryFromObjectPool(ctx context.Context, req *gitalypb.UnlinkRepositoryFromObjectPoolRequest) (*gitalypb.UnlinkRepositoryFromObjectPoolResponse, error) {
	if req.GetRepository() == nil {
		return nil, status.Error(codes.InvalidArgument, "no repository")
	}

	pool, err := poolForRequest(req)
	if err != nil {
		return nil, err
	}

	if err := pool.Unlink(ctx, req.GetRepository()); err != nil {
		return nil, err
	}

	return &gitalypb.UnlinkRepositoryFromObjectPoolResponse{}, nil
}
