/*
Pulp 3 API

Testing ContentConfigsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package zest/release/v3

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/release/v3"
)

func Test_zest/release/v3_ContentConfigsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentConfigsApiService ContentOstreeConfigsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentConfigsApi.ContentOstreeConfigsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentConfigsApiService ContentOstreeConfigsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ostreeOstreeConfigHref string

		resp, httpRes, err := apiClient.ContentConfigsApi.ContentOstreeConfigsRead(context.Background(), ostreeOstreeConfigHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
