/*
Pulp 3 API

Testing ContentInstallerFileIndicesApiService

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

func Test_zest_ContentInstallerFileIndicesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentInstallerFileIndicesApiService ContentDebInstallerFileIndicesCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentInstallerFileIndicesApi.ContentDebInstallerFileIndicesCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentInstallerFileIndicesApiService ContentDebInstallerFileIndicesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentInstallerFileIndicesApi.ContentDebInstallerFileIndicesList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentInstallerFileIndicesApiService ContentDebInstallerFileIndicesRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debInstallerFileIndexHref string

		resp, httpRes, err := apiClient.ContentInstallerFileIndicesApi.ContentDebInstallerFileIndicesRead(context.Background(), debInstallerFileIndexHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
