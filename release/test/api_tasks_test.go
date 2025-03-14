/*
Pulp 3 API

Testing TasksAPIService

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

func Test_zest_TasksAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TasksAPIService TasksAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskHref string

		resp, httpRes, err := apiClient.TasksAPI.TasksAddRole(context.Background(), taskHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksAPIService TasksCancel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskHref string

		resp, httpRes, err := apiClient.TasksAPI.TasksCancel(context.Background(), taskHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksAPIService TasksDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskHref string

		httpRes, err := apiClient.TasksAPI.TasksDelete(context.Background(), taskHref).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksAPIService TasksList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TasksAPI.TasksList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksAPIService TasksList2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.TasksAPI.TasksList2(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksAPIService TasksListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskHref string

		resp, httpRes, err := apiClient.TasksAPI.TasksListRoles(context.Background(), taskHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksAPIService TasksMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskHref string

		resp, httpRes, err := apiClient.TasksAPI.TasksMyPermissions(context.Background(), taskHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksAPIService TasksProfileArtifacts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskHref string

		resp, httpRes, err := apiClient.TasksAPI.TasksProfileArtifacts(context.Background(), taskHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksAPIService TasksPurge", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.TasksAPI.TasksPurge(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksAPIService TasksRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskHref string

		resp, httpRes, err := apiClient.TasksAPI.TasksRead(context.Background(), taskHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksAPIService TasksRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskHref string

		resp, httpRes, err := apiClient.TasksAPI.TasksRemoveRole(context.Background(), taskHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
