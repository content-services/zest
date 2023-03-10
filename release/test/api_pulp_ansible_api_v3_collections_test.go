/*
Pulp 3 API

Testing PulpAnsibleApiV3CollectionsApiService

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

func Test_zest_PulpAnsibleApiV3CollectionsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PulpAnsibleApiV3CollectionsApiService PulpAnsibleGalaxyApiV3CollectionsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string
		var path string

		resp, httpRes, err := apiClient.PulpAnsibleApiV3CollectionsApi.PulpAnsibleGalaxyApiV3CollectionsDelete(context.Background(), name, namespace, path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleApiV3CollectionsApiService PulpAnsibleGalaxyApiV3CollectionsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var path string

		resp, httpRes, err := apiClient.PulpAnsibleApiV3CollectionsApi.PulpAnsibleGalaxyApiV3CollectionsList(context.Background(), path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleApiV3CollectionsApiService PulpAnsibleGalaxyApiV3CollectionsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string
		var path string

		resp, httpRes, err := apiClient.PulpAnsibleApiV3CollectionsApi.PulpAnsibleGalaxyApiV3CollectionsRead(context.Background(), name, namespace, path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleApiV3CollectionsApiService PulpAnsibleGalaxyApiV3CollectionsUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string
		var path string

		resp, httpRes, err := apiClient.PulpAnsibleApiV3CollectionsApi.PulpAnsibleGalaxyApiV3CollectionsUpdate(context.Background(), name, namespace, path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
