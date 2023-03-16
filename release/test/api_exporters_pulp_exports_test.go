/*
Pulp 3 API

Testing ExportersPulpExportsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package zest/release/v3

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/release/v3"
)

func Test_zest/release/v3_ExportersPulpExportsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ExportersPulpExportsApiService ExportersCorePulpExportsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpExporterHref string

		resp, httpRes, err := apiClient.ExportersPulpExportsApi.ExportersCorePulpExportsCreate(context.Background(), pulpExporterHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportersPulpExportsApiService ExportersCorePulpExportsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpPulpExportHref string

		httpRes, err := apiClient.ExportersPulpExportsApi.ExportersCorePulpExportsDelete(context.Background(), pulpPulpExportHref).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportersPulpExportsApiService ExportersCorePulpExportsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpExporterHref string

		resp, httpRes, err := apiClient.ExportersPulpExportsApi.ExportersCorePulpExportsList(context.Background(), pulpExporterHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportersPulpExportsApiService ExportersCorePulpExportsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpPulpExportHref string

		resp, httpRes, err := apiClient.ExportersPulpExportsApi.ExportersCorePulpExportsRead(context.Background(), pulpPulpExportHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
