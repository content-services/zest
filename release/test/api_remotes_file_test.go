/*
Pulp 3 API

Testing RemotesFileAPIService

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

func Test_zest_RemotesFileAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RemotesFileAPIService RemotesFileFileAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRemoteHref string

		resp, httpRes, err := apiClient.RemotesFileAPI.RemotesFileFileAddRole(context.Background(), fileFileRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesFileAPIService RemotesFileFileCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.RemotesFileAPI.RemotesFileFileCreate(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesFileAPIService RemotesFileFileDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRemoteHref string

		resp, httpRes, err := apiClient.RemotesFileAPI.RemotesFileFileDelete(context.Background(), fileFileRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesFileAPIService RemotesFileFileList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.RemotesFileAPI.RemotesFileFileList(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesFileAPIService RemotesFileFileListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRemoteHref string

		resp, httpRes, err := apiClient.RemotesFileAPI.RemotesFileFileListRoles(context.Background(), fileFileRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesFileAPIService RemotesFileFileMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRemoteHref string

		resp, httpRes, err := apiClient.RemotesFileAPI.RemotesFileFileMyPermissions(context.Background(), fileFileRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesFileAPIService RemotesFileFilePartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRemoteHref string

		resp, httpRes, err := apiClient.RemotesFileAPI.RemotesFileFilePartialUpdate(context.Background(), fileFileRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesFileAPIService RemotesFileFileRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRemoteHref string

		resp, httpRes, err := apiClient.RemotesFileAPI.RemotesFileFileRead(context.Background(), fileFileRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesFileAPIService RemotesFileFileRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRemoteHref string

		resp, httpRes, err := apiClient.RemotesFileAPI.RemotesFileFileRemoveRole(context.Background(), fileFileRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesFileAPIService RemotesFileFileSetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRemoteHref string

		resp, httpRes, err := apiClient.RemotesFileAPI.RemotesFileFileSetLabel(context.Background(), fileFileRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesFileAPIService RemotesFileFileUnsetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRemoteHref string

		resp, httpRes, err := apiClient.RemotesFileAPI.RemotesFileFileUnsetLabel(context.Background(), fileFileRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesFileAPIService RemotesFileFileUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRemoteHref string

		resp, httpRes, err := apiClient.RemotesFileAPI.RemotesFileFileUpdate(context.Background(), fileFileRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
