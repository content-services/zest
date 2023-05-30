/*
Pulp 3 API

Testing RemotesRoleAPIService

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

func Test_zest_RemotesRoleAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RemotesRoleAPIService RemotesAnsibleRoleAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleRoleRemoteHref string

		resp, httpRes, err := apiClient.RemotesRoleAPI.RemotesAnsibleRoleAddRole(context.Background(), ansibleRoleRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesRoleAPIService RemotesAnsibleRoleCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RemotesRoleAPI.RemotesAnsibleRoleCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesRoleAPIService RemotesAnsibleRoleDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleRoleRemoteHref string

		resp, httpRes, err := apiClient.RemotesRoleAPI.RemotesAnsibleRoleDelete(context.Background(), ansibleRoleRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesRoleAPIService RemotesAnsibleRoleList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RemotesRoleAPI.RemotesAnsibleRoleList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesRoleAPIService RemotesAnsibleRoleListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleRoleRemoteHref string

		resp, httpRes, err := apiClient.RemotesRoleAPI.RemotesAnsibleRoleListRoles(context.Background(), ansibleRoleRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesRoleAPIService RemotesAnsibleRoleMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleRoleRemoteHref string

		resp, httpRes, err := apiClient.RemotesRoleAPI.RemotesAnsibleRoleMyPermissions(context.Background(), ansibleRoleRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesRoleAPIService RemotesAnsibleRolePartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleRoleRemoteHref string

		resp, httpRes, err := apiClient.RemotesRoleAPI.RemotesAnsibleRolePartialUpdate(context.Background(), ansibleRoleRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesRoleAPIService RemotesAnsibleRoleRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleRoleRemoteHref string

		resp, httpRes, err := apiClient.RemotesRoleAPI.RemotesAnsibleRoleRead(context.Background(), ansibleRoleRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesRoleAPIService RemotesAnsibleRoleRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleRoleRemoteHref string

		resp, httpRes, err := apiClient.RemotesRoleAPI.RemotesAnsibleRoleRemoveRole(context.Background(), ansibleRoleRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesRoleAPIService RemotesAnsibleRoleUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleRoleRemoteHref string

		resp, httpRes, err := apiClient.RemotesRoleAPI.RemotesAnsibleRoleUpdate(context.Background(), ansibleRoleRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
