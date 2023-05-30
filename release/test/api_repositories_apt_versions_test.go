/*
Pulp 3 API

Testing RepositoriesAptVersionsAPIService

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

func Test_zest_RepositoriesAptVersionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RepositoriesAptVersionsAPIService RepositoriesDebAptVersionsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debAptRepositoryVersionHref string

		resp, httpRes, err := apiClient.RepositoriesAptVersionsAPI.RepositoriesDebAptVersionsDelete(context.Background(), debAptRepositoryVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesAptVersionsAPIService RepositoriesDebAptVersionsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debAptRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesAptVersionsAPI.RepositoriesDebAptVersionsList(context.Background(), debAptRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesAptVersionsAPIService RepositoriesDebAptVersionsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debAptRepositoryVersionHref string

		resp, httpRes, err := apiClient.RepositoriesAptVersionsAPI.RepositoriesDebAptVersionsRead(context.Background(), debAptRepositoryVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesAptVersionsAPIService RepositoriesDebAptVersionsRepair", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debAptRepositoryVersionHref string

		resp, httpRes, err := apiClient.RepositoriesAptVersionsAPI.RepositoriesDebAptVersionsRepair(context.Background(), debAptRepositoryVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
