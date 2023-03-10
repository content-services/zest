/*
Pulp 3 API

Testing PypiMetadataApiService

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

func Test_zest_PypiMetadataApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PypiMetadataApiService PypiPypiRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var meta string
		var path string

		resp, httpRes, err := apiClient.PypiMetadataApi.PypiPypiRead(context.Background(), meta, path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
