/*
Pulp 3 API

Testing AcsRpmAPIService

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

func Test_zest_AcsRpmAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AcsRpmAPIService AcsRpmRpmAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmAlternateContentSourceHref string

		resp, httpRes, err := apiClient.AcsRpmAPI.AcsRpmRpmAddRole(context.Background(), rpmRpmAlternateContentSourceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsRpmAPIService AcsRpmRpmCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.AcsRpmAPI.AcsRpmRpmCreate(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsRpmAPIService AcsRpmRpmDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmAlternateContentSourceHref string

		resp, httpRes, err := apiClient.AcsRpmAPI.AcsRpmRpmDelete(context.Background(), rpmRpmAlternateContentSourceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsRpmAPIService AcsRpmRpmList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.AcsRpmAPI.AcsRpmRpmList(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsRpmAPIService AcsRpmRpmListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmAlternateContentSourceHref string

		resp, httpRes, err := apiClient.AcsRpmAPI.AcsRpmRpmListRoles(context.Background(), rpmRpmAlternateContentSourceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsRpmAPIService AcsRpmRpmMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmAlternateContentSourceHref string

		resp, httpRes, err := apiClient.AcsRpmAPI.AcsRpmRpmMyPermissions(context.Background(), rpmRpmAlternateContentSourceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsRpmAPIService AcsRpmRpmPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmAlternateContentSourceHref string

		resp, httpRes, err := apiClient.AcsRpmAPI.AcsRpmRpmPartialUpdate(context.Background(), rpmRpmAlternateContentSourceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsRpmAPIService AcsRpmRpmRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmAlternateContentSourceHref string

		resp, httpRes, err := apiClient.AcsRpmAPI.AcsRpmRpmRead(context.Background(), rpmRpmAlternateContentSourceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsRpmAPIService AcsRpmRpmRefresh", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmAlternateContentSourceHref string

		resp, httpRes, err := apiClient.AcsRpmAPI.AcsRpmRpmRefresh(context.Background(), rpmRpmAlternateContentSourceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsRpmAPIService AcsRpmRpmRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmAlternateContentSourceHref string

		resp, httpRes, err := apiClient.AcsRpmAPI.AcsRpmRpmRemoveRole(context.Background(), rpmRpmAlternateContentSourceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AcsRpmAPIService AcsRpmRpmUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmAlternateContentSourceHref string

		resp, httpRes, err := apiClient.AcsRpmAPI.AcsRpmRpmUpdate(context.Background(), rpmRpmAlternateContentSourceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
