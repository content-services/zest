/*
Pulp 3 API

Testing ContentPackagegroupsAPIService

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

func Test_zest_ContentPackagegroupsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentPackagegroupsAPIService ContentRpmPackagegroupsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.ContentPackagegroupsAPI.ContentRpmPackagegroupsList(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackagegroupsAPIService ContentRpmPackagegroupsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmPackageGroupHref string

		resp, httpRes, err := apiClient.ContentPackagegroupsAPI.ContentRpmPackagegroupsRead(context.Background(), rpmPackageGroupHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackagegroupsAPIService ContentRpmPackagegroupsSetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmPackageGroupHref string

		resp, httpRes, err := apiClient.ContentPackagegroupsAPI.ContentRpmPackagegroupsSetLabel(context.Background(), rpmPackageGroupHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackagegroupsAPIService ContentRpmPackagegroupsUnsetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmPackageGroupHref string

		resp, httpRes, err := apiClient.ContentPackagegroupsAPI.ContentRpmPackagegroupsUnsetLabel(context.Background(), rpmPackageGroupHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
