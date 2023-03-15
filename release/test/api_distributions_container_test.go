/*
Pulp 3 API

Testing DistributionsContainerApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package zest

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/v3/release"
)

func Test_zest_DistributionsContainerApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DistributionsContainerApiService DistributionsContainerContainerAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerDistributionHref string

		resp, httpRes, err := apiClient.DistributionsContainerApi.DistributionsContainerContainerAddRole(context.Background(), containerContainerDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsContainerApiService DistributionsContainerContainerCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DistributionsContainerApi.DistributionsContainerContainerCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsContainerApiService DistributionsContainerContainerDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerDistributionHref string

		resp, httpRes, err := apiClient.DistributionsContainerApi.DistributionsContainerContainerDelete(context.Background(), containerContainerDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsContainerApiService DistributionsContainerContainerList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DistributionsContainerApi.DistributionsContainerContainerList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsContainerApiService DistributionsContainerContainerListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerDistributionHref string

		resp, httpRes, err := apiClient.DistributionsContainerApi.DistributionsContainerContainerListRoles(context.Background(), containerContainerDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsContainerApiService DistributionsContainerContainerMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerDistributionHref string

		resp, httpRes, err := apiClient.DistributionsContainerApi.DistributionsContainerContainerMyPermissions(context.Background(), containerContainerDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsContainerApiService DistributionsContainerContainerPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerDistributionHref string

		resp, httpRes, err := apiClient.DistributionsContainerApi.DistributionsContainerContainerPartialUpdate(context.Background(), containerContainerDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsContainerApiService DistributionsContainerContainerRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerDistributionHref string

		resp, httpRes, err := apiClient.DistributionsContainerApi.DistributionsContainerContainerRead(context.Background(), containerContainerDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsContainerApiService DistributionsContainerContainerRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerDistributionHref string

		resp, httpRes, err := apiClient.DistributionsContainerApi.DistributionsContainerContainerRemoveRole(context.Background(), containerContainerDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsContainerApiService DistributionsContainerContainerUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerDistributionHref string

		resp, httpRes, err := apiClient.DistributionsContainerApi.DistributionsContainerContainerUpdate(context.Background(), containerContainerDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
