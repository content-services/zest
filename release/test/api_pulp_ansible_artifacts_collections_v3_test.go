/*
Pulp 3 API

Testing PulpAnsibleArtifactsCollectionsV3ApiService

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

func Test_zest_PulpAnsibleArtifactsCollectionsV3ApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PulpAnsibleArtifactsCollectionsV3ApiService PulpAnsibleGalaxyApiV3ArtifactsCollectionsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var path string

		resp, httpRes, err := apiClient.PulpAnsibleArtifactsCollectionsV3Api.PulpAnsibleGalaxyApiV3ArtifactsCollectionsCreate(context.Background(), path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleArtifactsCollectionsV3ApiService PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string
		var path string

		resp, httpRes, err := apiClient.PulpAnsibleArtifactsCollectionsV3Api.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsArtifactsCreate(context.Background(), distroBasePath, path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleArtifactsCollectionsV3ApiService PulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PulpAnsibleArtifactsCollectionsV3Api.PulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleArtifactsCollectionsV3ApiService PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string

		resp, httpRes, err := apiClient.PulpAnsibleArtifactsCollectionsV3Api.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsCreate(context.Background(), distroBasePath).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
