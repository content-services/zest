/*
Pulp 3 API

Testing ContentCollectionVersionsApiService

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

func Test_zest_ContentCollectionVersionsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentCollectionVersionsApiService ContentAnsibleCollectionVersionsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentCollectionVersionsApi.ContentAnsibleCollectionVersionsCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentCollectionVersionsApiService ContentAnsibleCollectionVersionsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentCollectionVersionsApi.ContentAnsibleCollectionVersionsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentCollectionVersionsApiService ContentAnsibleCollectionVersionsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleCollectionVersionHref string

		resp, httpRes, err := apiClient.ContentCollectionVersionsApi.ContentAnsibleCollectionVersionsRead(context.Background(), ansibleCollectionVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
