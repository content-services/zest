/*
Pulp 3 API

Testing RemotesCollectionApiService

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

func Test_zest_RemotesCollectionApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RemotesCollectionApiService RemotesAnsibleCollectionAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleCollectionRemoteHref string

		resp, httpRes, err := apiClient.RemotesCollectionApi.RemotesAnsibleCollectionAddRole(context.Background(), ansibleCollectionRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesCollectionApiService RemotesAnsibleCollectionCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RemotesCollectionApi.RemotesAnsibleCollectionCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesCollectionApiService RemotesAnsibleCollectionDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleCollectionRemoteHref string

		resp, httpRes, err := apiClient.RemotesCollectionApi.RemotesAnsibleCollectionDelete(context.Background(), ansibleCollectionRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesCollectionApiService RemotesAnsibleCollectionList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RemotesCollectionApi.RemotesAnsibleCollectionList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesCollectionApiService RemotesAnsibleCollectionListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleCollectionRemoteHref string

		resp, httpRes, err := apiClient.RemotesCollectionApi.RemotesAnsibleCollectionListRoles(context.Background(), ansibleCollectionRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesCollectionApiService RemotesAnsibleCollectionMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleCollectionRemoteHref string

		resp, httpRes, err := apiClient.RemotesCollectionApi.RemotesAnsibleCollectionMyPermissions(context.Background(), ansibleCollectionRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesCollectionApiService RemotesAnsibleCollectionPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleCollectionRemoteHref string

		resp, httpRes, err := apiClient.RemotesCollectionApi.RemotesAnsibleCollectionPartialUpdate(context.Background(), ansibleCollectionRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesCollectionApiService RemotesAnsibleCollectionRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleCollectionRemoteHref string

		resp, httpRes, err := apiClient.RemotesCollectionApi.RemotesAnsibleCollectionRead(context.Background(), ansibleCollectionRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesCollectionApiService RemotesAnsibleCollectionRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleCollectionRemoteHref string

		resp, httpRes, err := apiClient.RemotesCollectionApi.RemotesAnsibleCollectionRemoveRole(context.Background(), ansibleCollectionRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesCollectionApiService RemotesAnsibleCollectionUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleCollectionRemoteHref string

		resp, httpRes, err := apiClient.RemotesCollectionApi.RemotesAnsibleCollectionUpdate(context.Background(), ansibleCollectionRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
