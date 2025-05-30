/*
Pulp 3 API

Testing RemotesMavenAPIService

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

func Test_zest_RemotesMavenAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RemotesMavenAPIService RemotesMavenMavenCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.RemotesMavenAPI.RemotesMavenMavenCreate(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesMavenAPIService RemotesMavenMavenDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mavenMavenRemoteHref string

		resp, httpRes, err := apiClient.RemotesMavenAPI.RemotesMavenMavenDelete(context.Background(), mavenMavenRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesMavenAPIService RemotesMavenMavenList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.RemotesMavenAPI.RemotesMavenMavenList(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesMavenAPIService RemotesMavenMavenPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mavenMavenRemoteHref string

		resp, httpRes, err := apiClient.RemotesMavenAPI.RemotesMavenMavenPartialUpdate(context.Background(), mavenMavenRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesMavenAPIService RemotesMavenMavenRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mavenMavenRemoteHref string

		resp, httpRes, err := apiClient.RemotesMavenAPI.RemotesMavenMavenRead(context.Background(), mavenMavenRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesMavenAPIService RemotesMavenMavenSetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mavenMavenRemoteHref string

		resp, httpRes, err := apiClient.RemotesMavenAPI.RemotesMavenMavenSetLabel(context.Background(), mavenMavenRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesMavenAPIService RemotesMavenMavenUnsetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mavenMavenRemoteHref string

		resp, httpRes, err := apiClient.RemotesMavenAPI.RemotesMavenMavenUnsetLabel(context.Background(), mavenMavenRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesMavenAPIService RemotesMavenMavenUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mavenMavenRemoteHref string

		resp, httpRes, err := apiClient.RemotesMavenAPI.RemotesMavenMavenUpdate(context.Background(), mavenMavenRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
