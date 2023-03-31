/*
Pulp 3 API

Testing PublicationsRpmApiService

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

func Test_zest_PublicationsRpmApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PublicationsRpmApiService PublicationsRpmRpmAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmPublicationHref string

		resp, httpRes, err := apiClient.PublicationsRpmApi.PublicationsRpmRpmAddRole(context.Background(), rpmRpmPublicationHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PublicationsRpmApiService PublicationsRpmRpmCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PublicationsRpmApi.PublicationsRpmRpmCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PublicationsRpmApiService PublicationsRpmRpmDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmPublicationHref string

		httpRes, err := apiClient.PublicationsRpmApi.PublicationsRpmRpmDelete(context.Background(), rpmRpmPublicationHref).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PublicationsRpmApiService PublicationsRpmRpmList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PublicationsRpmApi.PublicationsRpmRpmList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PublicationsRpmApiService PublicationsRpmRpmListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmPublicationHref string

		resp, httpRes, err := apiClient.PublicationsRpmApi.PublicationsRpmRpmListRoles(context.Background(), rpmRpmPublicationHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PublicationsRpmApiService PublicationsRpmRpmMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmPublicationHref string

		resp, httpRes, err := apiClient.PublicationsRpmApi.PublicationsRpmRpmMyPermissions(context.Background(), rpmRpmPublicationHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PublicationsRpmApiService PublicationsRpmRpmRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmPublicationHref string

		resp, httpRes, err := apiClient.PublicationsRpmApi.PublicationsRpmRpmRead(context.Background(), rpmRpmPublicationHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PublicationsRpmApiService PublicationsRpmRpmRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmPublicationHref string

		resp, httpRes, err := apiClient.PublicationsRpmApi.PublicationsRpmRpmRemoveRole(context.Background(), rpmRpmPublicationHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
