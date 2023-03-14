/*
Pulp 3 API

Testing PulpAnsibleDefaultApiV3ArtifactsCollectionsApiService

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

func Test_zest_PulpAnsibleDefaultApiV3ArtifactsCollectionsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PulpAnsibleDefaultApiV3ArtifactsCollectionsApiService PulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var filename string
		var path string

		httpRes, err := apiClient.PulpAnsibleDefaultApiV3ArtifactsCollectionsApi.PulpAnsibleGalaxyDefaultApiV3ArtifactsCollectionsGet(context.Background(), filename, path).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
