// SPDX-FileCopyrightText: 2024 Josef Andersson
//
// SPDX-License-Identifier: EUPL-1.2

// Package provider handles operations related to Git providers and repositories.
package provider

import (
	"context"
	"fmt"

	"itiquette/git-provider-sync/internal/interfaces"
	"itiquette/git-provider-sync/internal/log"
	"itiquette/git-provider-sync/internal/model"
	config "itiquette/git-provider-sync/internal/model/configuration"
)

// Clone clones multiple repositories based on their metadata.
// It takes a context, a SourceReader interface for cloning operations,
// and a slice of RepositoryMetainfo containing information about the repositories to clone.
// It returns a slice of GitRepository interfaces representing the cloned repositories and any error encountered.
func Clone(ctx context.Context, reader interfaces.SourceReader, syncCfg config.SyncConfig, projectinfos []model.ProjectInfo) ([]interfaces.GitRepository, error) {
	logger := log.Logger(ctx)
	logger.Trace().Msg("Entering Clone")

	repositories := make([]interfaces.GitRepository, 0, len(projectinfos))

	for _, projectInfo := range projectinfos {
		opt := model.NewCloneOption(ctx, projectInfo, true, syncCfg)

		resultRepo, err := reader.Clone(ctx, opt)
		if err != nil {
			return nil, fmt.Errorf("failed to clone repository %s: %w", projectInfo.OriginalName, err)
		}

		resultRepo.ProjectMetaInfo = &projectInfo

		repositories = append(repositories, resultRepo)
	}

	return repositories, nil
}

// FetchProjectInfo retrieves metadata information for repositories from a Git provider.
// It takes a context, provider configuration, and a GitProvider interface.
// It returns a slice of RepositoryMetainfo containing the fetched metadata and any error encountered.
func FetchProjectInfo(ctx context.Context, syncCfg config.SyncConfig, gitProvider interfaces.GitProvider) ([]model.ProjectInfo, error) {
	logger := log.Logger(ctx)

	// Log the metadata fetching operation
	logger.Info().
		Str("Domain", syncCfg.GetDomain()).
		Str("Name", gitProvider.Name()).
		Str("Owner", syncCfg.Owner).
		Str("OwnerType", syncCfg.OwnerType).
		Msg("Fetching repository projectinfo/s from:")

	// Fetch the metadata from the Git provider
	// The 'true' parameter likely indicates that all available metadata should be fetched
	providerOption := model.NewProviderOption(
		syncCfg.IncludeForks,
		syncCfg.Owner,
		syncCfg.OwnerType,
		syncCfg.Repositories.IncludedRepositories(),
		syncCfg.Repositories.ExcludedRepositories(),
	)

	metainfo, err := gitProvider.ProjectInfos(ctx, providerOption, true)
	if err != nil {
		return nil, fmt.Errorf("failed to get repository meta information: %w", err)
	}

	return metainfo, nil
}
