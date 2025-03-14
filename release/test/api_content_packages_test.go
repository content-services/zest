/*
Pulp 3 API

Testing ContentPackagesAPIService

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

func Test_zest_ContentPackagesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentPackagesAPIService ContentNpmPackagesCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.ContentPackagesAPI.ContentNpmPackagesCreate(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackagesAPIService ContentNpmPackagesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.ContentPackagesAPI.ContentNpmPackagesList(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackagesAPIService ContentNpmPackagesRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var npmPackageHref string

		resp, httpRes, err := apiClient.ContentPackagesAPI.ContentNpmPackagesRead(context.Background(), npmPackageHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackagesAPIService ContentNpmPackagesSetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var npmPackageHref string

		resp, httpRes, err := apiClient.ContentPackagesAPI.ContentNpmPackagesSetLabel(context.Background(), npmPackageHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackagesAPIService ContentNpmPackagesUnsetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var npmPackageHref string

		resp, httpRes, err := apiClient.ContentPackagesAPI.ContentNpmPackagesUnsetLabel(context.Background(), npmPackageHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackagesAPIService ContentPythonPackagesCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.ContentPackagesAPI.ContentPythonPackagesCreate(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackagesAPIService ContentPythonPackagesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.ContentPackagesAPI.ContentPythonPackagesList(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackagesAPIService ContentPythonPackagesRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pythonPythonPackageContentHref string

		resp, httpRes, err := apiClient.ContentPackagesAPI.ContentPythonPackagesRead(context.Background(), pythonPythonPackageContentHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackagesAPIService ContentPythonPackagesSetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pythonPythonPackageContentHref string

		resp, httpRes, err := apiClient.ContentPackagesAPI.ContentPythonPackagesSetLabel(context.Background(), pythonPythonPackageContentHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackagesAPIService ContentPythonPackagesUnsetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pythonPythonPackageContentHref string

		resp, httpRes, err := apiClient.ContentPackagesAPI.ContentPythonPackagesUnsetLabel(context.Background(), pythonPythonPackageContentHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackagesAPIService ContentRpmPackagesCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.ContentPackagesAPI.ContentRpmPackagesCreate(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackagesAPIService ContentRpmPackagesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.ContentPackagesAPI.ContentRpmPackagesList(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackagesAPIService ContentRpmPackagesRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmPackageHref string

		resp, httpRes, err := apiClient.ContentPackagesAPI.ContentRpmPackagesRead(context.Background(), rpmPackageHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackagesAPIService ContentRpmPackagesSetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmPackageHref string

		resp, httpRes, err := apiClient.ContentPackagesAPI.ContentRpmPackagesSetLabel(context.Background(), rpmPackageHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackagesAPIService ContentRpmPackagesUnsetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmPackageHref string

		resp, httpRes, err := apiClient.ContentPackagesAPI.ContentRpmPackagesUnsetLabel(context.Background(), rpmPackageHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
