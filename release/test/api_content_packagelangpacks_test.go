/*
Pulp 3 API

Testing ContentPackagelangpacksAPIService

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

func Test_zest_ContentPackagelangpacksAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentPackagelangpacksAPIService ContentRpmPackagelangpacksList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.ContentPackagelangpacksAPI.ContentRpmPackagelangpacksList(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackagelangpacksAPIService ContentRpmPackagelangpacksRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmPackageLangpacksHref string

		resp, httpRes, err := apiClient.ContentPackagelangpacksAPI.ContentRpmPackagelangpacksRead(context.Background(), rpmPackageLangpacksHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackagelangpacksAPIService ContentRpmPackagelangpacksSetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmPackageLangpacksHref string

		resp, httpRes, err := apiClient.ContentPackagelangpacksAPI.ContentRpmPackagelangpacksSetLabel(context.Background(), rpmPackageLangpacksHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackagelangpacksAPIService ContentRpmPackagelangpacksUnsetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmPackageLangpacksHref string

		resp, httpRes, err := apiClient.ContentPackagelangpacksAPI.ContentRpmPackagelangpacksUnsetLabel(context.Background(), rpmPackageLangpacksHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
