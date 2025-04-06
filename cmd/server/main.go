package main

import (
	"context"
	"flag"
	"log"
	"os/signal"
	"sync"
	"syscall"

	"github.com/mohammadne/zanbil/cmd"
	"github.com/mohammadne/zanbil/internal/api/http"
	"github.com/mohammadne/zanbil/internal/api/http/i18n"
	"github.com/mohammadne/zanbil/internal/config"
	"github.com/mohammadne/zanbil/internal/repositories/cache"
	"github.com/mohammadne/zanbil/internal/repositories/storage"
	"github.com/mohammadne/zanbil/internal/usecases"
	"github.com/mohammadne/zanbil/pkg/databases/postgres"
	"github.com/mohammadne/zanbil/pkg/databases/redis"
	"github.com/mohammadne/zanbil/pkg/observability/logger"
	"go.uber.org/zap"
)

func main() {
	monitorPort := flag.Int("monitor-port", 8087, "The server port which handles monitoring endpoints (default: 8087)")
	requestPort := flag.Int("request-port", 8088, "The server port which handles http requests (default: 8088)")
	environmentRaw := flag.String("environment", "", "The environment (default: local)")
	flag.Parse() // Parse the command-line flags

	cfg, err := config.Load(config.ToEnvironment(*environmentRaw))
	if err != nil {
		log.Fatalf("failed to load config: \n%v", err)
	}

	logger, err := logger.New(cfg.Logger)
	if err != nil {
		log.Fatalf("failed to initialize logger: \n%v", err)
	}

	logger.Warn("Build Information", cmd.BuildInfo()...)

	postgres, err := postgres.Open(cfg.Postgres, config.Namespace, config.System)
	if err != nil {
		logger.Fatal("error connecting to postgres database", zap.Error(err))
	}

	// storages
	storageCategories := storage.NewCategories(logger, postgres)

	redis, err := redis.Open(cfg.Redis)
	if err != nil {
		logger.Fatal("error connecting to redis database", zap.Error(err))
	}

	// caches
	cacheCategories := cache.NewCategories(logger, redis)

	// usecases
	usecasesCategories := usecases.NewCategories(logger, cacheCategories, storageCategories)

	i18n, err := i18n.New(logger)
	if err != nil {
		logger.Fatal("failed to load i18n", zap.Error(err))
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	var wg sync.WaitGroup

	wg.Add(1)
	go http.New(logger, i18n, usecasesCategories).
		Serve(ctx, &wg, *monitorPort, *requestPort)

	<-ctx.Done()
	wg.Wait()
	logger.Warn("interruption signal recieved, gracefully shutdown the server")
}
