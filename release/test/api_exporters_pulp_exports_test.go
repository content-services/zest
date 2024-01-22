/*
Pulp 3 API

Testing ExportersPulpExportsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package zest

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/release/v2024"
)

func Test_zest_ExportersPulpExportsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ExportersPulpExportsAPIService ExportersCorePulpExportsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpExporterHref string

		resp, httpRes, err := apiClient.ExportersPulpExportsAPI.ExportersCorePulpExportsCreate(context.Background(), pulpExporterHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportersPulpExportsAPIService ExportersCorePulpExportsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpPulpExportHref string

		httpRes, err := apiClient.ExportersPulpExportsAPI.ExportersCorePulpExportsDelete(context.Background(), pulpPulpExportHref).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportersPulpExportsAPIService ExportersCorePulpExportsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpExporterHref string

		resp, httpRes, err := apiClient.ExportersPulpExportsAPI.ExportersCorePulpExportsList(context.Background(), pulpExporterHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportersPulpExportsAPIService ExportersCorePulpExportsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpPulpExportHref string

		resp, httpRes, err := apiClient.ExportersPulpExportsAPI.ExportersCorePulpExportsRead(context.Background(), pulpPulpExportHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
