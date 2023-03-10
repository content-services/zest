/*
Pulp 3 API

Testing PulpAnsibleApiV3PluginAnsibleContentCollectionsAllCollectionsApiService

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

func Test_zest_PulpAnsibleApiV3PluginAnsibleContentCollectionsAllCollectionsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PulpAnsibleApiV3PluginAnsibleContentCollectionsAllCollectionsApiService PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsAllCollectionsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string
		var path string

		resp, httpRes, err := apiClient.PulpAnsibleApiV3PluginAnsibleContentCollectionsAllCollectionsApi.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsAllCollectionsList(context.Background(), distroBasePath, path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
