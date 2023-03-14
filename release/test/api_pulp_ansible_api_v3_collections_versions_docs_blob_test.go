/*
Pulp 3 API

Testing PulpAnsibleApiV3CollectionsVersionsDocsBlobApiService

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

func Test_zest_PulpAnsibleApiV3CollectionsVersionsDocsBlobApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PulpAnsibleApiV3CollectionsVersionsDocsBlobApiService PulpAnsibleGalaxyApiV3CollectionsVersionsDocsBlobRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string
		var path string
		var version string

		resp, httpRes, err := apiClient.PulpAnsibleApiV3CollectionsVersionsDocsBlobApi.PulpAnsibleGalaxyApiV3CollectionsVersionsDocsBlobRead(context.Background(), name, namespace, path, version).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
