/*
Pulp 3 API

Testing RemotesAptApiService

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

func Test_zest_RemotesAptApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RemotesAptApiService RemotesDebAptCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RemotesAptApi.RemotesDebAptCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesAptApiService RemotesDebAptDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debAptRemoteHref string

		resp, httpRes, err := apiClient.RemotesAptApi.RemotesDebAptDelete(context.Background(), debAptRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesAptApiService RemotesDebAptList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RemotesAptApi.RemotesDebAptList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesAptApiService RemotesDebAptPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debAptRemoteHref string

		resp, httpRes, err := apiClient.RemotesAptApi.RemotesDebAptPartialUpdate(context.Background(), debAptRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesAptApiService RemotesDebAptRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debAptRemoteHref string

		resp, httpRes, err := apiClient.RemotesAptApi.RemotesDebAptRead(context.Background(), debAptRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesAptApiService RemotesDebAptUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debAptRemoteHref string

		resp, httpRes, err := apiClient.RemotesAptApi.RemotesDebAptUpdate(context.Background(), debAptRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
