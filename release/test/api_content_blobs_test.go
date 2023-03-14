/*
Pulp 3 API

Testing ContentBlobsApiService

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

func Test_zest_ContentBlobsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentBlobsApiService ContentContainerBlobsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentBlobsApi.ContentContainerBlobsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentBlobsApiService ContentContainerBlobsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerBlobHref string

		resp, httpRes, err := apiClient.ContentBlobsApi.ContentContainerBlobsRead(context.Background(), containerBlobHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
