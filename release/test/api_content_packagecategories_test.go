/*
Pulp 3 API

Testing ContentPackagecategoriesAPIService

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

func Test_zest_ContentPackagecategoriesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentPackagecategoriesAPIService ContentRpmPackagecategoriesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.ContentPackagecategoriesAPI.ContentRpmPackagecategoriesList(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackagecategoriesAPIService ContentRpmPackagecategoriesRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmPackageCategoryHref string

		resp, httpRes, err := apiClient.ContentPackagecategoriesAPI.ContentRpmPackagecategoriesRead(context.Background(), rpmPackageCategoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackagecategoriesAPIService ContentRpmPackagecategoriesSetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmPackageCategoryHref string

		resp, httpRes, err := apiClient.ContentPackagecategoriesAPI.ContentRpmPackagecategoriesSetLabel(context.Background(), rpmPackageCategoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackagecategoriesAPIService ContentRpmPackagecategoriesUnsetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmPackageCategoryHref string

		resp, httpRes, err := apiClient.ContentPackagecategoriesAPI.ContentRpmPackagecategoriesUnsetLabel(context.Background(), rpmPackageCategoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
