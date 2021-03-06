package objectpool

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gitlab-org/gitaly-proto/go/gitalypb"
	"gitlab.com/gitlab-org/gitaly/internal/git"
	"gitlab.com/gitlab-org/gitaly/internal/git/objectpool"
	"gitlab.com/gitlab-org/gitaly/internal/testhelper"
)

func TestReduplicate(t *testing.T) {
	server, serverSocketPath := runObjectPoolServer(t)
	defer server.Stop()

	client, conn := newObjectPoolClient(t, serverSocketPath)
	defer conn.Close()

	ctx, cancel := testhelper.Context()
	defer cancel()

	testRepo, testRepoPath, cleanupFn := testhelper.NewTestRepo(t)
	defer cleanupFn()

	pool, err := objectpool.NewObjectPool(testRepo.GetStorageName(), t.Name())
	require.NoError(t, err)
	defer pool.Remove(ctx)
	require.NoError(t, pool.Create(ctx, testRepo))
	require.NoError(t, pool.Link(ctx, testRepo))
	testhelper.MustRunCommand(t, nil, "git", "-C", testRepoPath, "gc")

	existingObjectID := "55bc176024cfa3baaceb71db584c7e5df900ea65"
	// Corrupt the repository to check if the object can't be found
	require.NoError(t, pool.Unlink(ctx, testRepo))
	cmd, err := git.Command(ctx, testRepo, "cat-file", "-e", existingObjectID)
	require.NoError(t, err)
	require.Error(t, cmd.Wait())

	// Reduplicate and check if the objects appear again
	require.NoError(t, pool.Link(ctx, testRepo))
	_, err = client.ReduplicateRepository(ctx, &gitalypb.ReduplicateRepositoryRequest{Repository: testRepo})
	require.NoError(t, err)

	require.NoError(t, pool.Unlink(ctx, testRepo))
	testhelper.MustRunCommand(t, nil, "git", "-C", testRepoPath, "cat-file", "-e", existingObjectID)
}
