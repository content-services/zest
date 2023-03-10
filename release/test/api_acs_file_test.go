/*
Pulp 3 API

Testing AcsFileApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package zest

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/release"
)

func Test_zest_AcsFileApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AcsFileApiService AcsFileFileAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileAlternateContentSourceHref string

		resp, httpRes, err := apiClient.AcsFileApi.AcsFileFileAddRole(context.Background(), fileFileAlternateContentSourceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsFileApiService AcsFileFileCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AcsFileApi.AcsFileFileCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsFileApiService AcsFileFileDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileAlternateContentSourceHref string

		resp, httpRes, err := apiClient.AcsFileApi.AcsFileFileDelete(context.Background(), fileFileAlternateContentSourceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsFileApiService AcsFileFileList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AcsFileApi.AcsFileFileList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsFileApiService AcsFileFileListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileAlternateContentSourceHref string

		resp, httpRes, err := apiClient.AcsFileApi.AcsFileFileListRoles(context.Background(), fileFileAlternateContentSourceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsFileApiService AcsFileFileMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileAlternateContentSourceHref string

		resp, httpRes, err := apiClient.AcsFileApi.AcsFileFileMyPermissions(context.Background(), fileFileAlternateContentSourceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsFileApiService AcsFileFilePartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileAlternateContentSourceHref string

		resp, httpRes, err := apiClient.AcsFileApi.AcsFileFilePartialUpdate(context.Background(), fileFileAlternateContentSourceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsFileApiService AcsFileFileRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileAlternateContentSourceHref string

		resp, httpRes, err := apiClient.AcsFileApi.AcsFileFileRead(context.Background(), fileFileAlternateContentSourceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsFileApiService AcsFileFileRefresh", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileAlternateContentSourceHref string

		resp, httpRes, err := apiClient.AcsFileApi.AcsFileFileRefresh(context.Background(), fileFileAlternateContentSourceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsFileApiService AcsFileFileRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileAlternateContentSourceHref string

		resp, httpRes, err := apiClient.AcsFileApi.AcsFileFileRemoveRole(context.Background(), fileFileAlternateContentSourceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsFileApiService AcsFileFileUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileAlternateContentSourceHref string

		resp, httpRes, err := apiClient.AcsFileApi.AcsFileFileUpdate(context.Background(), fileFileAlternateContentSourceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
