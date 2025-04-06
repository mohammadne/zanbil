package functional_test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/mohammadne/zanbil/internal/config"
	"github.com/mohammadne/zanbil/internal/repositories/storage"
	"github.com/mohammadne/zanbil/pkg/databases/postgres"
	"github.com/mohammadne/zanbil/pkg/observability/logger"
)

type StorageTestSuite struct {
	suite.Suite

	categories storage.Categories
}

func TestStorage(t *testing.T) {
	suite.Run(t, new(StorageTestSuite))
}

func (suite *StorageTestSuite) SetupSuite() {
	require := suite.Require()

	cfg, err := config.Load(config.EnvironmentLocal)
	require.Equal(nil, err)

	logger, err := logger.New(cfg.Logger)
	if err != nil {
		log.Fatalf("failed to initialize logger: \n%v", err)
	}

	postgres, err := postgres.Open(cfg.Postgres, config.Namespace, config.System)
	require.Equal(nil, err)

	suite.categories = storage.NewCategories(logger, postgres)
}

func (suite *StorageTestSuite) TestCategories() {
	require := suite.Require()

	suite.Run("all_categories", func() {
		storageCategories, failure := suite.categories.AllCategories(context.TODO())
		require.Equal(nil, failure)

		object, _ := json.MarshalIndent(storageCategories, "", "  ")
		fmt.Println(string(object))
	})
}
