// SPDX-FileCopyrightText: 2024 Josef Andersson
//
// SPDX-License-Identifier: EUPL-1.2

package target

import (
	"context"
	"errors"
	"fmt"
	"itiquette/git-provider-sync/internal/interfaces"
	"itiquette/git-provider-sync/internal/log"
	"itiquette/git-provider-sync/internal/model"
	gpsconfig "itiquette/git-provider-sync/internal/model/configuration"

	"github.com/go-git/go-git/v5"
	gogitconfig "github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
)

type GitLibOperation interface {
	Open(ctx context.Context, path string) (*git.Repository, error)
	GetWorktree(ctx context.Context, repo *git.Repository) (*git.Worktree, error)
	WorktreeStatus(ctx context.Context, wt *git.Worktree) error
	FetchBranches(ctx context.Context, repo *git.Repository, auth transport.AuthMethod, name string) error
	SetDefaultBranch(ctx context.Context, repo *git.Repository, branchName string) error
	CreateRemote(ctx context.Context, repo *git.Repository, name string, urls []string) error
	SetRemoteAndBranch(ctx context.Context, repository interfaces.GitRepository, targetDirPath string) error
}

type gitLibOperation struct {
}

func newGitLibOperation() GitLibOperation {
	return &gitLibOperation{}
}

func (r *gitLibOperation) Open(ctx context.Context, path string) (*git.Repository, error) {
	logger := log.Logger(ctx)
	logger.Trace().Msg("Entering open")
	logger.Debug().Str("path", path).Msg("open")

	repo, err := git.PlainOpen(path)
	if err != nil {
		return nil, fmt.Errorf("%w: %s: %w", ErrOpenRepository, path, err)
	}

	return repo, nil
}

func (r *gitLibOperation) GetWorktree(ctx context.Context, repo *git.Repository) (*git.Worktree, error) {
	logger := log.Logger(ctx)
	logger.Trace().Msg("Entering getWorktree")

	wt, err := repo.Worktree()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrWorktree, err)
	}

	return wt, nil
}

func (r *gitLibOperation) WorktreeStatus(ctx context.Context, wt *git.Worktree) error {
	logger := log.Logger(ctx)
	logger.Trace().Msg("Entering worktreeStatus")

	status, err := wt.Status()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrWorktree, err)
	}

	if !status.IsClean() {
		return ErrUncleanWorkspace
	}

	return nil
}

func (r *gitLibOperation) SetRemoteAndBranch(ctx context.Context, repository interfaces.GitRepository, targetDirPath string) error {
	logger := log.Logger(ctx)
	logger.Trace().Msg("Entering setRemoteAndBranch")
	logger.Debug().Str("targetDirPath", targetDirPath).Msg("setRemoteAndBranch")

	if model.CLIOptions(ctx).DryRun {
		logger.Info().Str("pushOpts", targetDirPath).Msg("dryRun. Ignored SetRemoteAndBranch.")

		return nil
	}

	repo, err := git.PlainOpen(targetDirPath)
	if err != nil {
		return fmt.Errorf("%w: %s: %w", ErrOpenRepository, targetDirPath, err)
	}

	remote, err := repository.GoGitRepository().Remote(gpsconfig.ORIGIN)
	if err == nil {
		if _, err := repo.CreateRemote(&gogitconfig.RemoteConfig{
			Name: gpsconfig.ORIGIN,
			URLs: remote.Config().URLs,
		}); err != nil {
			return fmt.Errorf("%w: %w", ErrRemoteCreation, err)
		}
	} else {
		logger.Warn().Str("targetDirPath", targetDirPath).Msg("Remote origin not found in repository")
	}

	return nil
}
func (r *gitLibOperation) FetchBranches(ctx context.Context, repo *git.Repository, auth transport.AuthMethod, name string) error {
	logger := log.Logger(ctx)
	logger.Trace().Msg("Entering fetchBranches")
	logger.Debug().Str("name", name).Msg("fetchBranches")

	options := &git.FetchOptions{
		RefSpecs: []gogitconfig.RefSpec{
			"refs/*:refs/*",
			"^refs/pull/*:refs/pull/*",
		},
		Auth: auth,
	}

	if err := repo.Fetch(options); err != nil {
		if errors.Is(err, git.NoErrAlreadyUpToDate) {
			logger.Debug().Str("name", name).Msg("repository already up-to-date")

			return nil
		}

		return fmt.Errorf("%w: %w", ErrFetchBranches, err)
	}

	return nil
}

func (r *gitLibOperation) SetDefaultBranch(ctx context.Context, repo *git.Repository, branchName string) error {
	logger := log.Logger(ctx)
	logger.Trace().Msg("Entering setDefaultBranch")
	logger.Debug().Str("branchName", branchName).Msg("setDefaultBranch")

	branchRef := plumbing.NewBranchReferenceName(branchName)

	worktree, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrWorktree, err)
	}

	if err = worktree.Checkout(&git.CheckoutOptions{
		Branch: branchRef,
		Force:  true,
	}); err != nil {
		return fmt.Errorf("%w: %s: %w", ErrBranchCheckout, branchName, err)
	}

	headRef := plumbing.NewSymbolicReference(plumbing.HEAD, branchRef)

	if err := repo.Storer.SetReference(headRef); err != nil {
		return fmt.Errorf("%w: %w", ErrHeadSet, err)
	}

	return nil
}

func (r *gitLibOperation) CreateRemote(ctx context.Context, repo *git.Repository, name string, urls []string) error {
	logger := log.Logger(ctx)
	logger.Trace().Msg("Entering createRemote")
	logger.Debug().Str("name", name).Msg("createRemote")

	_, err := repo.CreateRemote(&gogitconfig.RemoteConfig{
		Name: name,
		URLs: urls,
	})
	if err != nil {
		return fmt.Errorf("%w: %w", ErrRemoteCreation, err)
	}

	return nil
}
