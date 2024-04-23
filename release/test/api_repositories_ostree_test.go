/*
Pulp 3 API

Testing RepositoriesOstreeAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package zest

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/release/v2024"
)

func Test_zest_RepositoriesOstreeAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RepositoriesOstreeAPIService RepositoriesOstreeOstreeAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ostreeOstreeRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeAddRole(context.Background(), ostreeOstreeRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesOstreeAPIService RepositoriesOstreeOstreeCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeCreate(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesOstreeAPIService RepositoriesOstreeOstreeDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ostreeOstreeRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeDelete(context.Background(), ostreeOstreeRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesOstreeAPIService RepositoriesOstreeOstreeImportAll", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ostreeOstreeRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeImportAll(context.Background(), ostreeOstreeRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesOstreeAPIService RepositoriesOstreeOstreeImportCommits", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ostreeOstreeRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeImportCommits(context.Background(), ostreeOstreeRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesOstreeAPIService RepositoriesOstreeOstreeList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeList(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesOstreeAPIService RepositoriesOstreeOstreeListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ostreeOstreeRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeListRoles(context.Background(), ostreeOstreeRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesOstreeAPIService RepositoriesOstreeOstreeModify", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ostreeOstreeRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeModify(context.Background(), ostreeOstreeRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesOstreeAPIService RepositoriesOstreeOstreeMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ostreeOstreeRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeMyPermissions(context.Background(), ostreeOstreeRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesOstreeAPIService RepositoriesOstreeOstreePartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ostreeOstreeRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreePartialUpdate(context.Background(), ostreeOstreeRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesOstreeAPIService RepositoriesOstreeOstreeRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ostreeOstreeRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeRead(context.Background(), ostreeOstreeRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesOstreeAPIService RepositoriesOstreeOstreeRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ostreeOstreeRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeRemoveRole(context.Background(), ostreeOstreeRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesOstreeAPIService RepositoriesOstreeOstreeSetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ostreeOstreeRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeSetLabel(context.Background(), ostreeOstreeRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesOstreeAPIService RepositoriesOstreeOstreeSync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ostreeOstreeRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeSync(context.Background(), ostreeOstreeRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesOstreeAPIService RepositoriesOstreeOstreeUnsetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ostreeOstreeRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeUnsetLabel(context.Background(), ostreeOstreeRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesOstreeAPIService RepositoriesOstreeOstreeUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ostreeOstreeRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesOstreeAPI.RepositoriesOstreeOstreeUpdate(context.Background(), ostreeOstreeRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
