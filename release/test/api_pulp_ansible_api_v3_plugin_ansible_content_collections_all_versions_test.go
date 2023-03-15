/*
Pulp 3 API

Testing PulpAnsibleApiV3PluginAnsibleContentCollectionsAllVersionsApiService

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

func Test_zest/release/v3_PulpAnsibleApiV3PluginAnsibleContentCollectionsAllVersionsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PulpAnsibleApiV3PluginAnsibleContentCollectionsAllVersionsApiService PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsAllVersionsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string
		var path string

		resp, httpRes, err := apiClient.PulpAnsibleApiV3PluginAnsibleContentCollectionsAllVersionsApi.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsAllVersionsList(context.Background(), distroBasePath, path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
