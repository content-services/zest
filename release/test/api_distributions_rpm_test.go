/*
Pulp 3 API

Testing DistributionsRpmApiService

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

func Test_zest_DistributionsRpmApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DistributionsRpmApiService DistributionsRpmRpmAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmDistributionHref string

		resp, httpRes, err := apiClient.DistributionsRpmApi.DistributionsRpmRpmAddRole(context.Background(), rpmRpmDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsRpmApiService DistributionsRpmRpmCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DistributionsRpmApi.DistributionsRpmRpmCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsRpmApiService DistributionsRpmRpmDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmDistributionHref string

		resp, httpRes, err := apiClient.DistributionsRpmApi.DistributionsRpmRpmDelete(context.Background(), rpmRpmDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsRpmApiService DistributionsRpmRpmList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DistributionsRpmApi.DistributionsRpmRpmList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsRpmApiService DistributionsRpmRpmListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmDistributionHref string

		resp, httpRes, err := apiClient.DistributionsRpmApi.DistributionsRpmRpmListRoles(context.Background(), rpmRpmDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsRpmApiService DistributionsRpmRpmMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmDistributionHref string

		resp, httpRes, err := apiClient.DistributionsRpmApi.DistributionsRpmRpmMyPermissions(context.Background(), rpmRpmDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsRpmApiService DistributionsRpmRpmPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmDistributionHref string

		resp, httpRes, err := apiClient.DistributionsRpmApi.DistributionsRpmRpmPartialUpdate(context.Background(), rpmRpmDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsRpmApiService DistributionsRpmRpmRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmDistributionHref string

		resp, httpRes, err := apiClient.DistributionsRpmApi.DistributionsRpmRpmRead(context.Background(), rpmRpmDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsRpmApiService DistributionsRpmRpmRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmDistributionHref string

		resp, httpRes, err := apiClient.DistributionsRpmApi.DistributionsRpmRpmRemoveRole(context.Background(), rpmRpmDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsRpmApiService DistributionsRpmRpmUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmDistributionHref string

		resp, httpRes, err := apiClient.DistributionsRpmApi.DistributionsRpmRpmUpdate(context.Background(), rpmRpmDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
