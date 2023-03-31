/*
Pulp 3 API

Testing PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsArtifactsApiService

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

func Test_zest_PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsArtifactsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsArtifactsApiService PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsDownload", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string
		var filename string

		httpRes, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsArtifactsApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsDownload(context.Background(), distroBasePath, filename).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
