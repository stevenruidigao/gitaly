package remote

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	"gitlab.com/gitlab-org/gitaly-proto/go/gitalypb"
	"gitlab.com/gitlab-org/gitaly/internal/git"
	"gitlab.com/gitlab-org/gitaly/internal/rubyserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AddRemote adds a remote to the repository
func (s *server) AddRemote(ctx context.Context, req *gitalypb.AddRemoteRequest) (*gitalypb.AddRemoteResponse, error) {
	if err := validateAddRemoteRequest(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "AddRemote: %v", err)
	}

	client, err := s.RemoteServiceClient(ctx)
	if err != nil {
		return nil, err
	}

	clientCtx, err := rubyserver.SetHeaders(ctx, req.GetRepository())
	if err != nil {
		return nil, err
	}

	return client.AddRemote(clientCtx, req)
}

func validateAddRemoteRequest(req *gitalypb.AddRemoteRequest) error {
	if strings.TrimSpace(req.GetName()) == "" {
		return fmt.Errorf("empty remote name")
	}
	if req.GetUrl() == "" {
		return fmt.Errorf("empty remote url")
	}

	return nil
}

// RemoveRemote removes the given remote
func (s *server) RemoveRemote(ctx context.Context, req *gitalypb.RemoveRemoteRequest) (*gitalypb.RemoveRemoteResponse, error) {
	if err := validateRemoveRemoteRequest(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "AddRemote: %v", err)
	}

	client, err := s.RemoteServiceClient(ctx)
	if err != nil {
		return nil, err
	}

	clientCtx, err := rubyserver.SetHeaders(ctx, req.GetRepository())
	if err != nil {
		return nil, err
	}

	return client.RemoveRemote(clientCtx, req)
}

func (s *server) FindRemoteRepository(ctx context.Context, req *gitalypb.FindRemoteRepositoryRequest) (*gitalypb.FindRemoteRepositoryResponse, error) {
	if req.GetRemote() == "" {
		return nil, status.Error(codes.InvalidArgument, "FindRemoteRepository: empty remote can't be checked.")
	}

	cmd, err := git.CommandWithoutRepo(ctx, "ls-remote", req.GetRemote(), "HEAD")

	if err != nil {
		return nil, status.Errorf(codes.Internal, "error executing git command: %s", err)
	}

	output, err := ioutil.ReadAll(cmd)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to read stdout: %s", err)
	}
	if err := cmd.Wait(); err != nil {
		return &gitalypb.FindRemoteRepositoryResponse{Exists: false}, nil
	}

	// The output of a successful command is structured like
	// Regexp would've read better, but this is faster
	// 58fbff2e0d3b620f591a748c158799ead87b51cd	HEAD
	fields := bytes.Fields(output)
	match := len(fields) == 2 && len(fields[0]) == 40 && string(fields[1]) == "HEAD"

	return &gitalypb.FindRemoteRepositoryResponse{Exists: match}, nil
}

func validateRemoveRemoteRequest(req *gitalypb.RemoveRemoteRequest) error {
	if req.GetName() == "" {
		return fmt.Errorf("empty remote name")
	}

	return nil
}
