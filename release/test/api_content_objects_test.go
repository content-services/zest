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
	openapiclient "github.com/content-services/zest/release/v2025"
)

func Test_zest_ContentObjectsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentObjectsAPIService ContentOstreeObjectsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.ContentObjectsAPI.ContentOstreeObjectsList(context.Background(), pulpDomain).Execute()

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

	t.Run("Test ContentObjectsAPIService ContentOstreeObjectsSetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ostreeOstreeObjectHref string

		resp, httpRes, err := apiClient.ContentObjectsAPI.ContentOstreeObjectsSetLabel(context.Background(), ostreeOstreeObjectHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentObjectsAPIService ContentOstreeObjectsUnsetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ostreeOstreeObjectHref string

		resp, httpRes, err := apiClient.ContentObjectsAPI.ContentOstreeObjectsUnsetLabel(context.Background(), ostreeOstreeObjectHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
