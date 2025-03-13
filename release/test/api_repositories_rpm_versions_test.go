/*
Pulp 3 API

Testing RepositoriesRpmVersionsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package zest

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/release/v2025"
)

func Test_zest_RepositoriesRpmVersionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RepositoriesRpmVersionsAPIService RepositoriesRpmRpmVersionsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmRepositoryVersionHref string

		resp, httpRes, err := apiClient.RepositoriesRpmVersionsAPI.RepositoriesRpmRpmVersionsDelete(context.Background(), rpmRpmRepositoryVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesRpmVersionsAPIService RepositoriesRpmRpmVersionsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesRpmVersionsAPI.RepositoriesRpmRpmVersionsList(context.Background(), rpmRpmRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesRpmVersionsAPIService RepositoriesRpmRpmVersionsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmRepositoryVersionHref string

		resp, httpRes, err := apiClient.RepositoriesRpmVersionsAPI.RepositoriesRpmRpmVersionsRead(context.Background(), rpmRpmRepositoryVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesRpmVersionsAPIService RepositoriesRpmRpmVersionsRepair", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmRepositoryVersionHref string

		resp, httpRes, err := apiClient.RepositoriesRpmVersionsAPI.RepositoriesRpmRpmVersionsRepair(context.Background(), rpmRpmRepositoryVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
