/*
Pulp 3 API

Testing RepositoriesMavenApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package zest

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/release"
)

func Test_zest_RepositoriesMavenApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RepositoriesMavenApiService RepositoriesMavenMavenCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RepositoriesMavenApi.RepositoriesMavenMavenCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesMavenApiService RepositoriesMavenMavenDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mavenMavenRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesMavenApi.RepositoriesMavenMavenDelete(context.Background(), mavenMavenRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesMavenApiService RepositoriesMavenMavenList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RepositoriesMavenApi.RepositoriesMavenMavenList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesMavenApiService RepositoriesMavenMavenPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mavenMavenRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesMavenApi.RepositoriesMavenMavenPartialUpdate(context.Background(), mavenMavenRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesMavenApiService RepositoriesMavenMavenRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mavenMavenRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesMavenApi.RepositoriesMavenMavenRead(context.Background(), mavenMavenRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesMavenApiService RepositoriesMavenMavenUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mavenMavenRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesMavenApi.RepositoriesMavenMavenUpdate(context.Background(), mavenMavenRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}