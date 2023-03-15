/*
Pulp 3 API

Testing RepositoriesAptApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/v3/release--package-name&#x3D;zest"
)

func Test_openapi_RepositoriesAptApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RepositoriesAptApiService RepositoriesDebAptCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RepositoriesAptApi.RepositoriesDebAptCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesAptApiService RepositoriesDebAptDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debAptRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesAptApi.RepositoriesDebAptDelete(context.Background(), debAptRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesAptApiService RepositoriesDebAptList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RepositoriesAptApi.RepositoriesDebAptList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesAptApiService RepositoriesDebAptModify", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debAptRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesAptApi.RepositoriesDebAptModify(context.Background(), debAptRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesAptApiService RepositoriesDebAptPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debAptRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesAptApi.RepositoriesDebAptPartialUpdate(context.Background(), debAptRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesAptApiService RepositoriesDebAptRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debAptRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesAptApi.RepositoriesDebAptRead(context.Background(), debAptRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesAptApiService RepositoriesDebAptSync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debAptRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesAptApi.RepositoriesDebAptSync(context.Background(), debAptRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesAptApiService RepositoriesDebAptUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debAptRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesAptApi.RepositoriesDebAptUpdate(context.Background(), debAptRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
