/*
Pulp 3 API

Testing GroupsRolesAPIService

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

func Test_zest_GroupsRolesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GroupsRolesAPIService GroupsRolesCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupHref string

		resp, httpRes, err := apiClient.GroupsRolesAPI.GroupsRolesCreate(context.Background(), groupHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsRolesAPIService GroupsRolesDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupsGroupRoleHref string

		httpRes, err := apiClient.GroupsRolesAPI.GroupsRolesDelete(context.Background(), groupsGroupRoleHref).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsRolesAPIService GroupsRolesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupHref string

		resp, httpRes, err := apiClient.GroupsRolesAPI.GroupsRolesList(context.Background(), groupHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsRolesAPIService GroupsRolesRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupsGroupRoleHref string

		resp, httpRes, err := apiClient.GroupsRolesAPI.GroupsRolesRead(context.Background(), groupsGroupRoleHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
