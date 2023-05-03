/*
Pulp 3 API

Testing RemotesOstreeApiService

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

func Test_zest_RemotesOstreeApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RemotesOstreeApiService RemotesOstreeOstreeCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RemotesOstreeApi.RemotesOstreeOstreeCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesOstreeApiService RemotesOstreeOstreeDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ostreeOstreeRemoteHref string

		resp, httpRes, err := apiClient.RemotesOstreeApi.RemotesOstreeOstreeDelete(context.Background(), ostreeOstreeRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesOstreeApiService RemotesOstreeOstreeList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RemotesOstreeApi.RemotesOstreeOstreeList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesOstreeApiService RemotesOstreeOstreePartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ostreeOstreeRemoteHref string

		resp, httpRes, err := apiClient.RemotesOstreeApi.RemotesOstreeOstreePartialUpdate(context.Background(), ostreeOstreeRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesOstreeApiService RemotesOstreeOstreeRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ostreeOstreeRemoteHref string

		resp, httpRes, err := apiClient.RemotesOstreeApi.RemotesOstreeOstreeRead(context.Background(), ostreeOstreeRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesOstreeApiService RemotesOstreeOstreeUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ostreeOstreeRemoteHref string

		resp, httpRes, err := apiClient.RemotesOstreeApi.RemotesOstreeOstreeUpdate(context.Background(), ostreeOstreeRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}