/*
Pulp 3 API

Testing RemotesUlnAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package zest

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/release/v2024"
)

func Test_zest_RemotesUlnAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RemotesUlnAPIService RemotesRpmUlnAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmUlnRemoteHref string

		resp, httpRes, err := apiClient.RemotesUlnAPI.RemotesRpmUlnAddRole(context.Background(), rpmUlnRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesUlnAPIService RemotesRpmUlnCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.RemotesUlnAPI.RemotesRpmUlnCreate(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesUlnAPIService RemotesRpmUlnDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmUlnRemoteHref string

		resp, httpRes, err := apiClient.RemotesUlnAPI.RemotesRpmUlnDelete(context.Background(), rpmUlnRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesUlnAPIService RemotesRpmUlnList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.RemotesUlnAPI.RemotesRpmUlnList(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesUlnAPIService RemotesRpmUlnListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmUlnRemoteHref string

		resp, httpRes, err := apiClient.RemotesUlnAPI.RemotesRpmUlnListRoles(context.Background(), rpmUlnRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesUlnAPIService RemotesRpmUlnMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmUlnRemoteHref string

		resp, httpRes, err := apiClient.RemotesUlnAPI.RemotesRpmUlnMyPermissions(context.Background(), rpmUlnRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesUlnAPIService RemotesRpmUlnPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmUlnRemoteHref string

		resp, httpRes, err := apiClient.RemotesUlnAPI.RemotesRpmUlnPartialUpdate(context.Background(), rpmUlnRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesUlnAPIService RemotesRpmUlnRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmUlnRemoteHref string

		resp, httpRes, err := apiClient.RemotesUlnAPI.RemotesRpmUlnRead(context.Background(), rpmUlnRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesUlnAPIService RemotesRpmUlnRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmUlnRemoteHref string

		resp, httpRes, err := apiClient.RemotesUlnAPI.RemotesRpmUlnRemoveRole(context.Background(), rpmUlnRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesUlnAPIService RemotesRpmUlnSetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmUlnRemoteHref string

		resp, httpRes, err := apiClient.RemotesUlnAPI.RemotesRpmUlnSetLabel(context.Background(), rpmUlnRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesUlnAPIService RemotesRpmUlnUnsetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmUlnRemoteHref string

		resp, httpRes, err := apiClient.RemotesUlnAPI.RemotesRpmUlnUnsetLabel(context.Background(), rpmUlnRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesUlnAPIService RemotesRpmUlnUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmUlnRemoteHref string

		resp, httpRes, err := apiClient.RemotesUlnAPI.RemotesRpmUlnUpdate(context.Background(), rpmUlnRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
