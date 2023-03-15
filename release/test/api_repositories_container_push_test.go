/*
Pulp 3 API

Testing RepositoriesContainerPushApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package zest

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/v3/release"
)

func Test_zest_RepositoriesContainerPushApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RepositoriesContainerPushApiService RepositoriesContainerContainerPushAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerPushRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerPushApi.RepositoriesContainerContainerPushAddRole(context.Background(), containerContainerPushRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerPushApiService RepositoriesContainerContainerPushList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RepositoriesContainerPushApi.RepositoriesContainerContainerPushList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerPushApiService RepositoriesContainerContainerPushListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerPushRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerPushApi.RepositoriesContainerContainerPushListRoles(context.Background(), containerContainerPushRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerPushApiService RepositoriesContainerContainerPushMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerPushRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerPushApi.RepositoriesContainerContainerPushMyPermissions(context.Background(), containerContainerPushRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerPushApiService RepositoriesContainerContainerPushPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerPushRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerPushApi.RepositoriesContainerContainerPushPartialUpdate(context.Background(), containerContainerPushRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerPushApiService RepositoriesContainerContainerPushRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerPushRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerPushApi.RepositoriesContainerContainerPushRead(context.Background(), containerContainerPushRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerPushApiService RepositoriesContainerContainerPushRemoveImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerPushRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerPushApi.RepositoriesContainerContainerPushRemoveImage(context.Background(), containerContainerPushRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerPushApiService RepositoriesContainerContainerPushRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerPushRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerPushApi.RepositoriesContainerContainerPushRemoveRole(context.Background(), containerContainerPushRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerPushApiService RepositoriesContainerContainerPushRemoveSignatures", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerPushRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerPushApi.RepositoriesContainerContainerPushRemoveSignatures(context.Background(), containerContainerPushRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerPushApiService RepositoriesContainerContainerPushSign", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerPushRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerPushApi.RepositoriesContainerContainerPushSign(context.Background(), containerContainerPushRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerPushApiService RepositoriesContainerContainerPushTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerPushRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerPushApi.RepositoriesContainerContainerPushTag(context.Background(), containerContainerPushRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerPushApiService RepositoriesContainerContainerPushUntag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerPushRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerPushApi.RepositoriesContainerContainerPushUntag(context.Background(), containerContainerPushRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerPushApiService RepositoriesContainerContainerPushUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerPushRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerPushApi.RepositoriesContainerContainerPushUpdate(context.Background(), containerContainerPushRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
