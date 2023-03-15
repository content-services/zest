/*
Pulp 3 API

Testing PulpAnsibleApiV2CollectionsVersionsApiService

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

func Test_openapi_PulpAnsibleApiV2CollectionsVersionsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PulpAnsibleApiV2CollectionsVersionsApiService PulpAnsibleGalaxyApiV2CollectionsVersionsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string
		var path string
		var version string

		httpRes, err := apiClient.PulpAnsibleApiV2CollectionsVersionsApi.PulpAnsibleGalaxyApiV2CollectionsVersionsGet(context.Background(), name, namespace, path, version).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
