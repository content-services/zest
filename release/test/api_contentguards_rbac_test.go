/*
Pulp 3 API

Testing ContentguardsRbacApiService

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

func Test_zest_ContentguardsRbacApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentguardsRbacApiService ContentguardsCoreRbacAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rBACContentGuardHref string

		resp, httpRes, err := apiClient.ContentguardsRbacApi.ContentguardsCoreRbacAddRole(context.Background(), rBACContentGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsRbacApiService ContentguardsCoreRbacCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentguardsRbacApi.ContentguardsCoreRbacCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsRbacApiService ContentguardsCoreRbacDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rBACContentGuardHref string

		httpRes, err := apiClient.ContentguardsRbacApi.ContentguardsCoreRbacDelete(context.Background(), rBACContentGuardHref).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsRbacApiService ContentguardsCoreRbacList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentguardsRbacApi.ContentguardsCoreRbacList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsRbacApiService ContentguardsCoreRbacListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rBACContentGuardHref string

		resp, httpRes, err := apiClient.ContentguardsRbacApi.ContentguardsCoreRbacListRoles(context.Background(), rBACContentGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsRbacApiService ContentguardsCoreRbacMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rBACContentGuardHref string

		resp, httpRes, err := apiClient.ContentguardsRbacApi.ContentguardsCoreRbacMyPermissions(context.Background(), rBACContentGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsRbacApiService ContentguardsCoreRbacPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rBACContentGuardHref string

		resp, httpRes, err := apiClient.ContentguardsRbacApi.ContentguardsCoreRbacPartialUpdate(context.Background(), rBACContentGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsRbacApiService ContentguardsCoreRbacRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rBACContentGuardHref string

		resp, httpRes, err := apiClient.ContentguardsRbacApi.ContentguardsCoreRbacRead(context.Background(), rBACContentGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsRbacApiService ContentguardsCoreRbacRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rBACContentGuardHref string

		resp, httpRes, err := apiClient.ContentguardsRbacApi.ContentguardsCoreRbacRemoveRole(context.Background(), rBACContentGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsRbacApiService ContentguardsCoreRbacUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rBACContentGuardHref string

		resp, httpRes, err := apiClient.ContentguardsRbacApi.ContentguardsCoreRbacUpdate(context.Background(), rBACContentGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
