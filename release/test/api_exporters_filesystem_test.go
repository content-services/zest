/*
Pulp 3 API

Testing ExportersFilesystemApiService

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

func Test_openapi_ExportersFilesystemApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ExportersFilesystemApiService ExportersCoreFilesystemCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ExportersFilesystemApi.ExportersCoreFilesystemCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportersFilesystemApiService ExportersCoreFilesystemDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var filesystemExporterHref string

		resp, httpRes, err := apiClient.ExportersFilesystemApi.ExportersCoreFilesystemDelete(context.Background(), filesystemExporterHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportersFilesystemApiService ExportersCoreFilesystemList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ExportersFilesystemApi.ExportersCoreFilesystemList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportersFilesystemApiService ExportersCoreFilesystemPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var filesystemExporterHref string

		resp, httpRes, err := apiClient.ExportersFilesystemApi.ExportersCoreFilesystemPartialUpdate(context.Background(), filesystemExporterHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportersFilesystemApiService ExportersCoreFilesystemRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var filesystemExporterHref string

		resp, httpRes, err := apiClient.ExportersFilesystemApi.ExportersCoreFilesystemRead(context.Background(), filesystemExporterHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportersFilesystemApiService ExportersCoreFilesystemUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var filesystemExporterHref string

		resp, httpRes, err := apiClient.ExportersFilesystemApi.ExportersCoreFilesystemUpdate(context.Background(), filesystemExporterHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
