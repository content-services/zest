/*
Pulp 3 API

Testing RemotesGitApiService

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

func Test_pulpGoBinding_RemotesGitApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RemotesGitApiService RemotesAnsibleGitCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RemotesGitApi.RemotesAnsibleGitCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesGitApiService RemotesAnsibleGitDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleGitRemoteHref string

		resp, httpRes, err := apiClient.RemotesGitApi.RemotesAnsibleGitDelete(context.Background(), ansibleGitRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesGitApiService RemotesAnsibleGitList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RemotesGitApi.RemotesAnsibleGitList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesGitApiService RemotesAnsibleGitPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleGitRemoteHref string

		resp, httpRes, err := apiClient.RemotesGitApi.RemotesAnsibleGitPartialUpdate(context.Background(), ansibleGitRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesGitApiService RemotesAnsibleGitRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleGitRemoteHref string

		resp, httpRes, err := apiClient.RemotesGitApi.RemotesAnsibleGitRead(context.Background(), ansibleGitRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RemotesGitApiService RemotesAnsibleGitUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleGitRemoteHref string

		resp, httpRes, err := apiClient.RemotesGitApi.RemotesAnsibleGitUpdate(context.Background(), ansibleGitRemoteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}