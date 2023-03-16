/*
Pulp 3 API

Testing RemotesFileApiService

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

func Test_zest/release/v3_RemotesFileApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RemotesFileApiService RemotesFileFileAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRemoteHref string

		resp, httpRes, err := apiClient.RemotesFileApi.RemotesFileFileAddRole(context.Background(), fileFileRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesFileApiService RemotesFileFileCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RemotesFileApi.RemotesFileFileCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesFileApiService RemotesFileFileDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRemoteHref string

		resp, httpRes, err := apiClient.RemotesFileApi.RemotesFileFileDelete(context.Background(), fileFileRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesFileApiService RemotesFileFileList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RemotesFileApi.RemotesFileFileList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesFileApiService RemotesFileFileListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRemoteHref string

		resp, httpRes, err := apiClient.RemotesFileApi.RemotesFileFileListRoles(context.Background(), fileFileRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesFileApiService RemotesFileFileMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRemoteHref string

		resp, httpRes, err := apiClient.RemotesFileApi.RemotesFileFileMyPermissions(context.Background(), fileFileRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesFileApiService RemotesFileFilePartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRemoteHref string

		resp, httpRes, err := apiClient.RemotesFileApi.RemotesFileFilePartialUpdate(context.Background(), fileFileRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesFileApiService RemotesFileFileRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRemoteHref string

		resp, httpRes, err := apiClient.RemotesFileApi.RemotesFileFileRead(context.Background(), fileFileRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesFileApiService RemotesFileFileRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRemoteHref string

		resp, httpRes, err := apiClient.RemotesFileApi.RemotesFileFileRemoveRole(context.Background(), fileFileRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesFileApiService RemotesFileFileUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRemoteHref string

		resp, httpRes, err := apiClient.RemotesFileApi.RemotesFileFileUpdate(context.Background(), fileFileRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
