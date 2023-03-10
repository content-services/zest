/*
Pulp 3 API

Testing DistributionsPypiApiService

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

func Test_zest_DistributionsPypiApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DistributionsPypiApiService DistributionsPythonPypiCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DistributionsPypiApi.DistributionsPythonPypiCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsPypiApiService DistributionsPythonPypiDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pythonPythonDistributionHref string

		resp, httpRes, err := apiClient.DistributionsPypiApi.DistributionsPythonPypiDelete(context.Background(), pythonPythonDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsPypiApiService DistributionsPythonPypiList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DistributionsPypiApi.DistributionsPythonPypiList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsPypiApiService DistributionsPythonPypiPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pythonPythonDistributionHref string

		resp, httpRes, err := apiClient.DistributionsPypiApi.DistributionsPythonPypiPartialUpdate(context.Background(), pythonPythonDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsPypiApiService DistributionsPythonPypiRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pythonPythonDistributionHref string

		resp, httpRes, err := apiClient.DistributionsPypiApi.DistributionsPythonPypiRead(context.Background(), pythonPythonDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsPypiApiService DistributionsPythonPypiUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pythonPythonDistributionHref string

		resp, httpRes, err := apiClient.DistributionsPypiApi.DistributionsPythonPypiUpdate(context.Background(), pythonPythonDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
