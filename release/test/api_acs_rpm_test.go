/*
Pulp 3 API

Testing AcsRpmApiService

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

func Test_zest_AcsRpmApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AcsRpmApiService AcsRpmRpmAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmAlternateContentSourceHref string

		resp, httpRes, err := apiClient.AcsRpmApi.AcsRpmRpmAddRole(context.Background(), rpmRpmAlternateContentSourceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsRpmApiService AcsRpmRpmCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AcsRpmApi.AcsRpmRpmCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsRpmApiService AcsRpmRpmDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmAlternateContentSourceHref string

		resp, httpRes, err := apiClient.AcsRpmApi.AcsRpmRpmDelete(context.Background(), rpmRpmAlternateContentSourceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsRpmApiService AcsRpmRpmList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AcsRpmApi.AcsRpmRpmList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsRpmApiService AcsRpmRpmListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmAlternateContentSourceHref string

		resp, httpRes, err := apiClient.AcsRpmApi.AcsRpmRpmListRoles(context.Background(), rpmRpmAlternateContentSourceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsRpmApiService AcsRpmRpmMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmAlternateContentSourceHref string

		resp, httpRes, err := apiClient.AcsRpmApi.AcsRpmRpmMyPermissions(context.Background(), rpmRpmAlternateContentSourceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsRpmApiService AcsRpmRpmPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmAlternateContentSourceHref string

		resp, httpRes, err := apiClient.AcsRpmApi.AcsRpmRpmPartialUpdate(context.Background(), rpmRpmAlternateContentSourceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsRpmApiService AcsRpmRpmRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmAlternateContentSourceHref string

		resp, httpRes, err := apiClient.AcsRpmApi.AcsRpmRpmRead(context.Background(), rpmRpmAlternateContentSourceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsRpmApiService AcsRpmRpmRefresh", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmAlternateContentSourceHref string

		resp, httpRes, err := apiClient.AcsRpmApi.AcsRpmRpmRefresh(context.Background(), rpmRpmAlternateContentSourceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsRpmApiService AcsRpmRpmRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmAlternateContentSourceHref string

		resp, httpRes, err := apiClient.AcsRpmApi.AcsRpmRpmRemoveRole(context.Background(), rpmRpmAlternateContentSourceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsRpmApiService AcsRpmRpmUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmAlternateContentSourceHref string

		resp, httpRes, err := apiClient.AcsRpmApi.AcsRpmRpmUpdate(context.Background(), rpmRpmAlternateContentSourceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
