/*
Pulp 3 API

Testing GroupsAPIService

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

func Test_zest_GroupsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GroupsAPIService GroupsAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupHref string

		resp, httpRes, err := apiClient.GroupsAPI.GroupsAddRole(context.Background(), groupHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService GroupsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.GroupsAPI.GroupsCreate(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService GroupsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupHref string

		httpRes, err := apiClient.GroupsAPI.GroupsDelete(context.Background(), groupHref).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService GroupsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.GroupsAPI.GroupsList(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService GroupsListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupHref string

		resp, httpRes, err := apiClient.GroupsAPI.GroupsListRoles(context.Background(), groupHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService GroupsMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupHref string

		resp, httpRes, err := apiClient.GroupsAPI.GroupsMyPermissions(context.Background(), groupHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService GroupsPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupHref string

		resp, httpRes, err := apiClient.GroupsAPI.GroupsPartialUpdate(context.Background(), groupHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService GroupsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupHref string

		resp, httpRes, err := apiClient.GroupsAPI.GroupsRead(context.Background(), groupHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService GroupsRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupHref string

		resp, httpRes, err := apiClient.GroupsAPI.GroupsRemoveRole(context.Background(), groupHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService GroupsUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupHref string

		resp, httpRes, err := apiClient.GroupsAPI.GroupsUpdate(context.Background(), groupHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
