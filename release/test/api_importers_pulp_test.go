/*
Pulp 3 API

Testing ImportersPulpAPIService

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

func Test_zest_ImportersPulpAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ImportersPulpAPIService ImportersCorePulpCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.ImportersPulpAPI.ImportersCorePulpCreate(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportersPulpAPIService ImportersCorePulpDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpImporterHref string

		httpRes, err := apiClient.ImportersPulpAPI.ImportersCorePulpDelete(context.Background(), pulpImporterHref).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportersPulpAPIService ImportersCorePulpList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.ImportersPulpAPI.ImportersCorePulpList(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportersPulpAPIService ImportersCorePulpPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpImporterHref string

		resp, httpRes, err := apiClient.ImportersPulpAPI.ImportersCorePulpPartialUpdate(context.Background(), pulpImporterHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportersPulpAPIService ImportersCorePulpRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpImporterHref string

		resp, httpRes, err := apiClient.ImportersPulpAPI.ImportersCorePulpRead(context.Background(), pulpImporterHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportersPulpAPIService ImportersCorePulpUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpImporterHref string

		resp, httpRes, err := apiClient.ImportersPulpAPI.ImportersCorePulpUpdate(context.Background(), pulpImporterHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
