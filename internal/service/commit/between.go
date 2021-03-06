package commit

import (
	"fmt"

	"gitlab.com/gitlab-org/gitaly-proto/go/gitalypb"
	"gitlab.com/gitlab-org/gitaly/internal/git"
	"gitlab.com/gitlab-org/gitaly/internal/helper"
	"gitlab.com/gitlab-org/gitaly/internal/helper/chunk"
)

type commitsBetweenSender struct {
	stream  gitalypb.CommitService_CommitsBetweenServer
	commits []*gitalypb.GitCommit
}

func (sender *commitsBetweenSender) Reset() { sender.commits = nil }
func (sender *commitsBetweenSender) Append(it chunk.Item) {
	sender.commits = append(sender.commits, it.(*gitalypb.GitCommit))
}
func (sender *commitsBetweenSender) Send() error {
	return sender.stream.Send(&gitalypb.CommitsBetweenResponse{Commits: sender.commits})
}

func (s *server) CommitsBetween(in *gitalypb.CommitsBetweenRequest, stream gitalypb.CommitService_CommitsBetweenServer) error {
	if err := validateCommitsBetween(in); err != nil {
		return helper.ErrInvalidArgument(err)
	}

	sender := &commitsBetweenSender{stream: stream}
	revisionRange := fmt.Sprintf("%s..%s", in.GetFrom(), in.GetTo())

	if err := sendCommits(stream.Context(), sender, in.GetRepository(), []string{revisionRange}, nil, "--reverse"); err != nil {
		return helper.ErrInternal(err)
	}

	return nil
}

func validateCommitsBetween(in *gitalypb.CommitsBetweenRequest) error {
	if err := git.ValidateRevision(in.GetFrom()); err != nil {
		return fmt.Errorf("from: %v", err)
	}

	if err := git.ValidateRevision(in.GetTo()); err != nil {
		return fmt.Errorf("to: %v", err)
	}

	return nil
}
