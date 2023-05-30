/*
Pulp 3 API

Testing RepositoriesPythonAPIService

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

func Test_zest_RepositoriesPythonAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RepositoriesPythonAPIService RepositoriesPythonPythonCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RepositoriesPythonAPI.RepositoriesPythonPythonCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesPythonAPIService RepositoriesPythonPythonDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pythonPythonRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesPythonAPI.RepositoriesPythonPythonDelete(context.Background(), pythonPythonRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesPythonAPIService RepositoriesPythonPythonList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RepositoriesPythonAPI.RepositoriesPythonPythonList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesPythonAPIService RepositoriesPythonPythonModify", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pythonPythonRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesPythonAPI.RepositoriesPythonPythonModify(context.Background(), pythonPythonRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesPythonAPIService RepositoriesPythonPythonPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pythonPythonRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesPythonAPI.RepositoriesPythonPythonPartialUpdate(context.Background(), pythonPythonRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesPythonAPIService RepositoriesPythonPythonRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pythonPythonRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesPythonAPI.RepositoriesPythonPythonRead(context.Background(), pythonPythonRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesPythonAPIService RepositoriesPythonPythonSync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pythonPythonRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesPythonAPI.RepositoriesPythonPythonSync(context.Background(), pythonPythonRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesPythonAPIService RepositoriesPythonPythonUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pythonPythonRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesPythonAPI.RepositoriesPythonPythonUpdate(context.Background(), pythonPythonRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
