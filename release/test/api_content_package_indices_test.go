/*
Pulp 3 API

Testing ContentPackageIndicesApiService

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

func Test_zest_ContentPackageIndicesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentPackageIndicesApiService ContentDebPackageIndicesCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentPackageIndicesApi.ContentDebPackageIndicesCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackageIndicesApiService ContentDebPackageIndicesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentPackageIndicesApi.ContentDebPackageIndicesList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackageIndicesApiService ContentDebPackageIndicesRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debPackageIndexHref string

		resp, httpRes, err := apiClient.ContentPackageIndicesApi.ContentDebPackageIndicesRead(context.Background(), debPackageIndexHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
