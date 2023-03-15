/*
Pulp 3 API

Testing RepositoriesRpmVersionsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package zest/release/v3

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/release/v3"
)

func Test_zest/release/v3_RepositoriesRpmVersionsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RepositoriesRpmVersionsApiService RepositoriesRpmRpmVersionsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmRepositoryVersionHref string

		resp, httpRes, err := apiClient.RepositoriesRpmVersionsApi.RepositoriesRpmRpmVersionsDelete(context.Background(), rpmRpmRepositoryVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesRpmVersionsApiService RepositoriesRpmRpmVersionsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesRpmVersionsApi.RepositoriesRpmRpmVersionsList(context.Background(), rpmRpmRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesRpmVersionsApiService RepositoriesRpmRpmVersionsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmRepositoryVersionHref string

		resp, httpRes, err := apiClient.RepositoriesRpmVersionsApi.RepositoriesRpmRpmVersionsRead(context.Background(), rpmRpmRepositoryVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesRpmVersionsApiService RepositoriesRpmRpmVersionsRepair", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmRepositoryVersionHref string

		resp, httpRes, err := apiClient.RepositoriesRpmVersionsApi.RepositoriesRpmRpmVersionsRepair(context.Background(), rpmRpmRepositoryVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
