/*
Pulp 3 API

Testing PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllVersionsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/v3/release--package-name&#x3D;zest"
)

func Test_openapi_PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllVersionsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllVersionsApiService PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllVersionsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string

		resp, httpRes, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllVersionsApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllVersionsList(context.Background(), distroBasePath).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
