/*
Pulp 3 API

Testing ContentModulemdObsoletesApiService

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

func Test_zest_ContentModulemdObsoletesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentModulemdObsoletesApiService ContentRpmModulemdObsoletesCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentModulemdObsoletesApi.ContentRpmModulemdObsoletesCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentModulemdObsoletesApiService ContentRpmModulemdObsoletesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentModulemdObsoletesApi.ContentRpmModulemdObsoletesList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentModulemdObsoletesApiService ContentRpmModulemdObsoletesRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmModulemdObsoleteHref string

		resp, httpRes, err := apiClient.ContentModulemdObsoletesApi.ContentRpmModulemdObsoletesRead(context.Background(), rpmModulemdObsoleteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
