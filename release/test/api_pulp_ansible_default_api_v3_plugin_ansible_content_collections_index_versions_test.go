/*
Pulp 3 API

Testing PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package zest

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/release/v3"
)

func Test_zest_PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsApiService PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string
		var name string
		var namespace string
		var version string

		resp, httpRes, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsDelete(context.Background(), distroBasePath, name, namespace, version).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsApiService PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string
		var name string
		var namespace string

		resp, httpRes, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsList(context.Background(), distroBasePath, name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsApiService PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string
		var name string
		var namespace string
		var version string

		resp, httpRes, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexVersionsRead(context.Background(), distroBasePath, name, namespace, version).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
