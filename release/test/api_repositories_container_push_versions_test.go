/*
Pulp 3 API

Testing RepositoriesContainerPushVersionsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package zest

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/release/v3"
)

func Test_zest_RepositoriesContainerPushVersionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RepositoriesContainerPushVersionsAPIService RepositoriesContainerContainerPushVersionsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerPushRepositoryVersionHref string

		resp, httpRes, err := apiClient.RepositoriesContainerPushVersionsAPI.RepositoriesContainerContainerPushVersionsDelete(context.Background(), containerContainerPushRepositoryVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerPushVersionsAPIService RepositoriesContainerContainerPushVersionsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerPushRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerPushVersionsAPI.RepositoriesContainerContainerPushVersionsList(context.Background(), containerContainerPushRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerPushVersionsAPIService RepositoriesContainerContainerPushVersionsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerPushRepositoryVersionHref string

		resp, httpRes, err := apiClient.RepositoriesContainerPushVersionsAPI.RepositoriesContainerContainerPushVersionsRead(context.Background(), containerContainerPushRepositoryVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerPushVersionsAPIService RepositoriesContainerContainerPushVersionsRepair", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerPushRepositoryVersionHref string

		resp, httpRes, err := apiClient.RepositoriesContainerPushVersionsAPI.RepositoriesContainerContainerPushVersionsRepair(context.Background(), containerContainerPushRepositoryVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
