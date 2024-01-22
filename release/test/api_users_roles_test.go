/*
Pulp 3 API

Testing UsersRolesAPIService

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

func Test_zest_UsersRolesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UsersRolesAPIService UsersRolesCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var authUserHref string

		resp, httpRes, err := apiClient.UsersRolesAPI.UsersRolesCreate(context.Background(), authUserHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersRolesAPIService UsersRolesDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var authUsersUserRoleHref string

		httpRes, err := apiClient.UsersRolesAPI.UsersRolesDelete(context.Background(), authUsersUserRoleHref).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersRolesAPIService UsersRolesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var authUserHref string

		resp, httpRes, err := apiClient.UsersRolesAPI.UsersRolesList(context.Background(), authUserHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersRolesAPIService UsersRolesRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var authUsersUserRoleHref string

		resp, httpRes, err := apiClient.UsersRolesAPI.UsersRolesRead(context.Background(), authUsersUserRoleHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
