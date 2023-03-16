/*
Pulp 3 API

Testing ContentPackagelangpacksApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package zest/release/v3

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/release/v3"
)

func Test_zest/release/v3_ContentPackagelangpacksApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentPackagelangpacksApiService ContentRpmPackagelangpacksList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentPackagelangpacksApi.ContentRpmPackagelangpacksList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackagelangpacksApiService ContentRpmPackagelangpacksRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmPackageLangpacksHref string

		resp, httpRes, err := apiClient.ContentPackagelangpacksApi.ContentRpmPackagelangpacksRead(context.Background(), rpmPackageLangpacksHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
