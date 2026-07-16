// pipeline.go

package main

import (
	"context"
	"time"

	"github.com/Harsh-Sonker/security-news-scraper/internal/cluster"
	"github.com/Harsh-Sonker/security-news-scraper/internal/config"
	"github.com/Harsh-Sonker/security-news-scraper/internal/fetch"
	"github.com/Harsh-Sonker/security-news-scraper/internal/ingest"
	"github.com/Harsh-Sonker/security-news-scraper/internal/source"
	"github.com/Harsh-Sonker/security-news-scraper/internal/store"
)

func ingestAndCluster(ctx context.Context, fc *fetch.Client, st *store.Store, cfg config.Config, targets []source.Source, start time.Time) (ingest.Summary, cluster.Stats, error) {
	summary, err := ingest.Run(ctx, fc, st, cfg, targets, start)
	if err != nil {
		return ingest.Summary{}, cluster.Stats{}, err
	}
	sinceUnix := start.Unix() - int64(cfg.Cluster.LookbackHours)*secondsPerHour
	windowSeconds := int64(cfg.Cluster.WindowHours) * secondsPerHour
	stats, err := cluster.Rebuild(st, cfg.Cluster.TitleJaccard, windowSeconds, sinceUnix)
	if err != nil {
		return summary, cluster.Stats{}, err
	}
	return summary, stats, nil
}
