/*
Pulp 3 API

Testing ContentReleaseComponentsApiService

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

func Test_openapi_ContentReleaseComponentsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentReleaseComponentsApiService ContentDebReleaseComponentsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentReleaseComponentsApi.ContentDebReleaseComponentsCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentReleaseComponentsApiService ContentDebReleaseComponentsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentReleaseComponentsApi.ContentDebReleaseComponentsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentReleaseComponentsApiService ContentDebReleaseComponentsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debReleaseComponentHref string

		resp, httpRes, err := apiClient.ContentReleaseComponentsApi.ContentDebReleaseComponentsRead(context.Background(), debReleaseComponentHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
