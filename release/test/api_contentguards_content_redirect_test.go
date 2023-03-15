/*
Pulp 3 API

Testing ContentguardsContentRedirectApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/v3/release--package-name&#x3D;zest"
)

func Test_openapi_ContentguardsContentRedirectApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentguardsContentRedirectApiService ContentguardsCoreContentRedirectAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var contentRedirectContentGuardHref string

		resp, httpRes, err := apiClient.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectAddRole(context.Background(), contentRedirectContentGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsContentRedirectApiService ContentguardsCoreContentRedirectCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsContentRedirectApiService ContentguardsCoreContentRedirectDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var contentRedirectContentGuardHref string

		httpRes, err := apiClient.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectDelete(context.Background(), contentRedirectContentGuardHref).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsContentRedirectApiService ContentguardsCoreContentRedirectList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsContentRedirectApiService ContentguardsCoreContentRedirectListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var contentRedirectContentGuardHref string

		resp, httpRes, err := apiClient.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectListRoles(context.Background(), contentRedirectContentGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsContentRedirectApiService ContentguardsCoreContentRedirectMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var contentRedirectContentGuardHref string

		resp, httpRes, err := apiClient.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectMyPermissions(context.Background(), contentRedirectContentGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsContentRedirectApiService ContentguardsCoreContentRedirectPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var contentRedirectContentGuardHref string

		resp, httpRes, err := apiClient.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectPartialUpdate(context.Background(), contentRedirectContentGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsContentRedirectApiService ContentguardsCoreContentRedirectRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var contentRedirectContentGuardHref string

		resp, httpRes, err := apiClient.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectRead(context.Background(), contentRedirectContentGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsContentRedirectApiService ContentguardsCoreContentRedirectRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var contentRedirectContentGuardHref string

		resp, httpRes, err := apiClient.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectRemoveRole(context.Background(), contentRedirectContentGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsContentRedirectApiService ContentguardsCoreContentRedirectUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var contentRedirectContentGuardHref string

		resp, httpRes, err := apiClient.ContentguardsContentRedirectApi.ContentguardsCoreContentRedirectUpdate(context.Background(), contentRedirectContentGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
