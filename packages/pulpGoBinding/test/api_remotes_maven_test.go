/*
Pulp 3 API

Testing RemotesMavenApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package pulpGoBinding

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/pulpGoBinding/packages/pulpGoBinding"
)

func Test_pulpGoBinding_RemotesMavenApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RemotesMavenApiService RemotesMavenMavenCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RemotesMavenApi.RemotesMavenMavenCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesMavenApiService RemotesMavenMavenDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mavenMavenRemoteHref string

		resp, httpRes, err := apiClient.RemotesMavenApi.RemotesMavenMavenDelete(context.Background(), mavenMavenRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesMavenApiService RemotesMavenMavenList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RemotesMavenApi.RemotesMavenMavenList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesMavenApiService RemotesMavenMavenPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mavenMavenRemoteHref string

		resp, httpRes, err := apiClient.RemotesMavenApi.RemotesMavenMavenPartialUpdate(context.Background(), mavenMavenRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesMavenApiService RemotesMavenMavenRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mavenMavenRemoteHref string

		resp, httpRes, err := apiClient.RemotesMavenApi.RemotesMavenMavenRead(context.Background(), mavenMavenRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesMavenApiService RemotesMavenMavenUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mavenMavenRemoteHref string

		resp, httpRes, err := apiClient.RemotesMavenApi.RemotesMavenMavenUpdate(context.Background(), mavenMavenRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}