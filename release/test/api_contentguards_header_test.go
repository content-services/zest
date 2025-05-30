/*
Pulp 3 API

Testing ContentguardsHeaderAPIService

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

func Test_zest_ContentguardsHeaderAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentguardsHeaderAPIService ContentguardsCoreHeaderAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var headerContentGuardHref string

		resp, httpRes, err := apiClient.ContentguardsHeaderAPI.ContentguardsCoreHeaderAddRole(context.Background(), headerContentGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsHeaderAPIService ContentguardsCoreHeaderCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.ContentguardsHeaderAPI.ContentguardsCoreHeaderCreate(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsHeaderAPIService ContentguardsCoreHeaderDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var headerContentGuardHref string

		httpRes, err := apiClient.ContentguardsHeaderAPI.ContentguardsCoreHeaderDelete(context.Background(), headerContentGuardHref).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsHeaderAPIService ContentguardsCoreHeaderList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.ContentguardsHeaderAPI.ContentguardsCoreHeaderList(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsHeaderAPIService ContentguardsCoreHeaderListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var headerContentGuardHref string

		resp, httpRes, err := apiClient.ContentguardsHeaderAPI.ContentguardsCoreHeaderListRoles(context.Background(), headerContentGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsHeaderAPIService ContentguardsCoreHeaderMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var headerContentGuardHref string

		resp, httpRes, err := apiClient.ContentguardsHeaderAPI.ContentguardsCoreHeaderMyPermissions(context.Background(), headerContentGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsHeaderAPIService ContentguardsCoreHeaderPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var headerContentGuardHref string

		resp, httpRes, err := apiClient.ContentguardsHeaderAPI.ContentguardsCoreHeaderPartialUpdate(context.Background(), headerContentGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsHeaderAPIService ContentguardsCoreHeaderRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var headerContentGuardHref string

		resp, httpRes, err := apiClient.ContentguardsHeaderAPI.ContentguardsCoreHeaderRead(context.Background(), headerContentGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsHeaderAPIService ContentguardsCoreHeaderRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var headerContentGuardHref string

		resp, httpRes, err := apiClient.ContentguardsHeaderAPI.ContentguardsCoreHeaderRemoveRole(context.Background(), headerContentGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsHeaderAPIService ContentguardsCoreHeaderUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var headerContentGuardHref string

		resp, httpRes, err := apiClient.ContentguardsHeaderAPI.ContentguardsCoreHeaderUpdate(context.Background(), headerContentGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
