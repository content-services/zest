/*
Pulp 3 API

Testing RemotesContainerAPIService

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

func Test_zest_RemotesContainerAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RemotesContainerAPIService RemotesContainerContainerAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerRemoteHref string

		resp, httpRes, err := apiClient.RemotesContainerAPI.RemotesContainerContainerAddRole(context.Background(), containerContainerRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesContainerAPIService RemotesContainerContainerCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.RemotesContainerAPI.RemotesContainerContainerCreate(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesContainerAPIService RemotesContainerContainerDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerRemoteHref string

		resp, httpRes, err := apiClient.RemotesContainerAPI.RemotesContainerContainerDelete(context.Background(), containerContainerRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesContainerAPIService RemotesContainerContainerList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.RemotesContainerAPI.RemotesContainerContainerList(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesContainerAPIService RemotesContainerContainerListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerRemoteHref string

		resp, httpRes, err := apiClient.RemotesContainerAPI.RemotesContainerContainerListRoles(context.Background(), containerContainerRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesContainerAPIService RemotesContainerContainerMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerRemoteHref string

		resp, httpRes, err := apiClient.RemotesContainerAPI.RemotesContainerContainerMyPermissions(context.Background(), containerContainerRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesContainerAPIService RemotesContainerContainerPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerRemoteHref string

		resp, httpRes, err := apiClient.RemotesContainerAPI.RemotesContainerContainerPartialUpdate(context.Background(), containerContainerRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesContainerAPIService RemotesContainerContainerRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerRemoteHref string

		resp, httpRes, err := apiClient.RemotesContainerAPI.RemotesContainerContainerRead(context.Background(), containerContainerRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesContainerAPIService RemotesContainerContainerRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerRemoteHref string

		resp, httpRes, err := apiClient.RemotesContainerAPI.RemotesContainerContainerRemoveRole(context.Background(), containerContainerRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesContainerAPIService RemotesContainerContainerSetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerRemoteHref string

		resp, httpRes, err := apiClient.RemotesContainerAPI.RemotesContainerContainerSetLabel(context.Background(), containerContainerRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesContainerAPIService RemotesContainerContainerUnsetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerRemoteHref string

		resp, httpRes, err := apiClient.RemotesContainerAPI.RemotesContainerContainerUnsetLabel(context.Background(), containerContainerRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesContainerAPIService RemotesContainerContainerUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerRemoteHref string

		resp, httpRes, err := apiClient.RemotesContainerAPI.RemotesContainerContainerUpdate(context.Background(), containerContainerRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
