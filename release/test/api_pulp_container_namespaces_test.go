/*
Pulp 3 API

Testing PulpContainerNamespacesApiService

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

func Test_zest_PulpContainerNamespacesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PulpContainerNamespacesApiService PulpContainerNamespacesAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerNamespaceHref string

		resp, httpRes, err := apiClient.PulpContainerNamespacesApi.PulpContainerNamespacesAddRole(context.Background(), containerContainerNamespaceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpContainerNamespacesApiService PulpContainerNamespacesCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PulpContainerNamespacesApi.PulpContainerNamespacesCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpContainerNamespacesApiService PulpContainerNamespacesDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerNamespaceHref string

		resp, httpRes, err := apiClient.PulpContainerNamespacesApi.PulpContainerNamespacesDelete(context.Background(), containerContainerNamespaceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpContainerNamespacesApiService PulpContainerNamespacesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PulpContainerNamespacesApi.PulpContainerNamespacesList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpContainerNamespacesApiService PulpContainerNamespacesListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerNamespaceHref string

		resp, httpRes, err := apiClient.PulpContainerNamespacesApi.PulpContainerNamespacesListRoles(context.Background(), containerContainerNamespaceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpContainerNamespacesApiService PulpContainerNamespacesMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerNamespaceHref string

		resp, httpRes, err := apiClient.PulpContainerNamespacesApi.PulpContainerNamespacesMyPermissions(context.Background(), containerContainerNamespaceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpContainerNamespacesApiService PulpContainerNamespacesRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerNamespaceHref string

		resp, httpRes, err := apiClient.PulpContainerNamespacesApi.PulpContainerNamespacesRead(context.Background(), containerContainerNamespaceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpContainerNamespacesApiService PulpContainerNamespacesRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerNamespaceHref string

		resp, httpRes, err := apiClient.PulpContainerNamespacesApi.PulpContainerNamespacesRemoveRole(context.Background(), containerContainerNamespaceHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
