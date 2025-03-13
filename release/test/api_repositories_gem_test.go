/*
Pulp 3 API

Testing RepositoriesGemAPIService

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

func Test_zest_RepositoriesGemAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RepositoriesGemAPIService RepositoriesGemGemAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gemGemRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesGemAPI.RepositoriesGemGemAddRole(context.Background(), gemGemRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesGemAPIService RepositoriesGemGemCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.RepositoriesGemAPI.RepositoriesGemGemCreate(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesGemAPIService RepositoriesGemGemDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gemGemRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesGemAPI.RepositoriesGemGemDelete(context.Background(), gemGemRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesGemAPIService RepositoriesGemGemList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.RepositoriesGemAPI.RepositoriesGemGemList(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesGemAPIService RepositoriesGemGemListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gemGemRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesGemAPI.RepositoriesGemGemListRoles(context.Background(), gemGemRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesGemAPIService RepositoriesGemGemModify", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gemGemRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesGemAPI.RepositoriesGemGemModify(context.Background(), gemGemRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesGemAPIService RepositoriesGemGemMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gemGemRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesGemAPI.RepositoriesGemGemMyPermissions(context.Background(), gemGemRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesGemAPIService RepositoriesGemGemPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gemGemRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesGemAPI.RepositoriesGemGemPartialUpdate(context.Background(), gemGemRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesGemAPIService RepositoriesGemGemRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gemGemRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesGemAPI.RepositoriesGemGemRead(context.Background(), gemGemRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesGemAPIService RepositoriesGemGemRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gemGemRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesGemAPI.RepositoriesGemGemRemoveRole(context.Background(), gemGemRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesGemAPIService RepositoriesGemGemSetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gemGemRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesGemAPI.RepositoriesGemGemSetLabel(context.Background(), gemGemRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesGemAPIService RepositoriesGemGemSync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gemGemRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesGemAPI.RepositoriesGemGemSync(context.Background(), gemGemRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesGemAPIService RepositoriesGemGemUnsetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gemGemRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesGemAPI.RepositoriesGemGemUnsetLabel(context.Background(), gemGemRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesGemAPIService RepositoriesGemGemUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gemGemRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesGemAPI.RepositoriesGemGemUpdate(context.Background(), gemGemRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
