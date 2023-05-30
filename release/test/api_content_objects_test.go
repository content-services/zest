/*
Pulp 3 API

Testing ContentObjectsAPIService

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

func Test_zest_ContentObjectsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentObjectsAPIService ContentOstreeObjectsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentObjectsAPI.ContentOstreeObjectsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentObjectsAPIService ContentOstreeObjectsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ostreeOstreeObjectHref string

		resp, httpRes, err := apiClient.ContentObjectsAPI.ContentOstreeObjectsRead(context.Background(), ostreeOstreeObjectHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
