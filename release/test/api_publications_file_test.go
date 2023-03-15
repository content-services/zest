/*
Pulp 3 API

Testing PublicationsFileApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package zest/release/v3

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/release/v3"
)

func Test_zest/release/v3_PublicationsFileApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PublicationsFileApiService PublicationsFileFileAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFilePublicationHref string

		resp, httpRes, err := apiClient.PublicationsFileApi.PublicationsFileFileAddRole(context.Background(), fileFilePublicationHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PublicationsFileApiService PublicationsFileFileCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PublicationsFileApi.PublicationsFileFileCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PublicationsFileApiService PublicationsFileFileDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFilePublicationHref string

		httpRes, err := apiClient.PublicationsFileApi.PublicationsFileFileDelete(context.Background(), fileFilePublicationHref).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PublicationsFileApiService PublicationsFileFileList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PublicationsFileApi.PublicationsFileFileList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PublicationsFileApiService PublicationsFileFileListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFilePublicationHref string

		resp, httpRes, err := apiClient.PublicationsFileApi.PublicationsFileFileListRoles(context.Background(), fileFilePublicationHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PublicationsFileApiService PublicationsFileFileMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFilePublicationHref string

		resp, httpRes, err := apiClient.PublicationsFileApi.PublicationsFileFileMyPermissions(context.Background(), fileFilePublicationHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PublicationsFileApiService PublicationsFileFileRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFilePublicationHref string

		resp, httpRes, err := apiClient.PublicationsFileApi.PublicationsFileFileRead(context.Background(), fileFilePublicationHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PublicationsFileApiService PublicationsFileFileRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFilePublicationHref string

		resp, httpRes, err := apiClient.PublicationsFileApi.PublicationsFileFileRemoveRole(context.Background(), fileFilePublicationHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
