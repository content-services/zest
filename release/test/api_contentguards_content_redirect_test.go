/*
Pulp 3 API

Testing ContentguardsContentRedirectAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package zest

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/release/v2023"
)

func Test_zest_ContentguardsContentRedirectAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentguardsContentRedirectAPIService ContentguardsCoreContentRedirectAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var contentRedirectContentGuardHref string

		resp, httpRes, err := apiClient.ContentguardsContentRedirectAPI.ContentguardsCoreContentRedirectAddRole(context.Background(), contentRedirectContentGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsContentRedirectAPIService ContentguardsCoreContentRedirectCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.ContentguardsContentRedirectAPI.ContentguardsCoreContentRedirectCreate(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsContentRedirectAPIService ContentguardsCoreContentRedirectDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var contentRedirectContentGuardHref string

		httpRes, err := apiClient.ContentguardsContentRedirectAPI.ContentguardsCoreContentRedirectDelete(context.Background(), contentRedirectContentGuardHref).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsContentRedirectAPIService ContentguardsCoreContentRedirectList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.ContentguardsContentRedirectAPI.ContentguardsCoreContentRedirectList(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsContentRedirectAPIService ContentguardsCoreContentRedirectListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var contentRedirectContentGuardHref string

		resp, httpRes, err := apiClient.ContentguardsContentRedirectAPI.ContentguardsCoreContentRedirectListRoles(context.Background(), contentRedirectContentGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsContentRedirectAPIService ContentguardsCoreContentRedirectMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var contentRedirectContentGuardHref string

		resp, httpRes, err := apiClient.ContentguardsContentRedirectAPI.ContentguardsCoreContentRedirectMyPermissions(context.Background(), contentRedirectContentGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsContentRedirectAPIService ContentguardsCoreContentRedirectPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var contentRedirectContentGuardHref string

		resp, httpRes, err := apiClient.ContentguardsContentRedirectAPI.ContentguardsCoreContentRedirectPartialUpdate(context.Background(), contentRedirectContentGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsContentRedirectAPIService ContentguardsCoreContentRedirectRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var contentRedirectContentGuardHref string

		resp, httpRes, err := apiClient.ContentguardsContentRedirectAPI.ContentguardsCoreContentRedirectRead(context.Background(), contentRedirectContentGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsContentRedirectAPIService ContentguardsCoreContentRedirectRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var contentRedirectContentGuardHref string

		resp, httpRes, err := apiClient.ContentguardsContentRedirectAPI.ContentguardsCoreContentRedirectRemoveRole(context.Background(), contentRedirectContentGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsContentRedirectAPIService ContentguardsCoreContentRedirectUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var contentRedirectContentGuardHref string

		resp, httpRes, err := apiClient.ContentguardsContentRedirectAPI.ContentguardsCoreContentRedirectUpdate(context.Background(), contentRedirectContentGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
