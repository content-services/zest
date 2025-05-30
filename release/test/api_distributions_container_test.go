/*
Pulp 3 API

Testing DistributionsContainerAPIService

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

func Test_zest_DistributionsContainerAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DistributionsContainerAPIService DistributionsContainerContainerAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerDistributionHref string

		resp, httpRes, err := apiClient.DistributionsContainerAPI.DistributionsContainerContainerAddRole(context.Background(), containerContainerDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsContainerAPIService DistributionsContainerContainerCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.DistributionsContainerAPI.DistributionsContainerContainerCreate(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsContainerAPIService DistributionsContainerContainerDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerDistributionHref string

		resp, httpRes, err := apiClient.DistributionsContainerAPI.DistributionsContainerContainerDelete(context.Background(), containerContainerDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsContainerAPIService DistributionsContainerContainerList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.DistributionsContainerAPI.DistributionsContainerContainerList(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsContainerAPIService DistributionsContainerContainerListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerDistributionHref string

		resp, httpRes, err := apiClient.DistributionsContainerAPI.DistributionsContainerContainerListRoles(context.Background(), containerContainerDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsContainerAPIService DistributionsContainerContainerMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerDistributionHref string

		resp, httpRes, err := apiClient.DistributionsContainerAPI.DistributionsContainerContainerMyPermissions(context.Background(), containerContainerDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsContainerAPIService DistributionsContainerContainerPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerDistributionHref string

		resp, httpRes, err := apiClient.DistributionsContainerAPI.DistributionsContainerContainerPartialUpdate(context.Background(), containerContainerDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsContainerAPIService DistributionsContainerContainerRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerDistributionHref string

		resp, httpRes, err := apiClient.DistributionsContainerAPI.DistributionsContainerContainerRead(context.Background(), containerContainerDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsContainerAPIService DistributionsContainerContainerRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerDistributionHref string

		resp, httpRes, err := apiClient.DistributionsContainerAPI.DistributionsContainerContainerRemoveRole(context.Background(), containerContainerDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsContainerAPIService DistributionsContainerContainerSetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerDistributionHref string

		resp, httpRes, err := apiClient.DistributionsContainerAPI.DistributionsContainerContainerSetLabel(context.Background(), containerContainerDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsContainerAPIService DistributionsContainerContainerUnsetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerDistributionHref string

		resp, httpRes, err := apiClient.DistributionsContainerAPI.DistributionsContainerContainerUnsetLabel(context.Background(), containerContainerDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsContainerAPIService DistributionsContainerContainerUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerDistributionHref string

		resp, httpRes, err := apiClient.DistributionsContainerAPI.DistributionsContainerContainerUpdate(context.Background(), containerContainerDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
