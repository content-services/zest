/*
Pulp 3 API

Testing PulpAnsibleApiV3CollectionsVersionsDocsBlobAPIService

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

func Test_zest_PulpAnsibleApiV3CollectionsVersionsDocsBlobAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PulpAnsibleApiV3CollectionsVersionsDocsBlobAPIService PulpAnsibleGalaxyApiV3CollectionsVersionsDocsBlobRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string
		var path string
		var version string

		resp, httpRes, err := apiClient.PulpAnsibleApiV3CollectionsVersionsDocsBlobAPI.PulpAnsibleGalaxyApiV3CollectionsVersionsDocsBlobRead(context.Background(), name, namespace, path, version).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
