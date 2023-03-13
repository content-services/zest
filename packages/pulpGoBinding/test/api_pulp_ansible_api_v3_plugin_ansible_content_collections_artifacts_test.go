/*
Pulp 3 API

Testing PulpAnsibleApiV3PluginAnsibleContentCollectionsArtifactsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package pulpGoBinding

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/pulpGoBinding/packages/pulpGoBinding"
)

func Test_pulpGoBinding_PulpAnsibleApiV3PluginAnsibleContentCollectionsArtifactsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PulpAnsibleApiV3PluginAnsibleContentCollectionsArtifactsApiService PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsDownload", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string
		var filename string
		var path string

		httpRes, err := apiClient.PulpAnsibleApiV3PluginAnsibleContentCollectionsArtifactsApi.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsDownload(context.Background(), distroBasePath, filename, path).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}