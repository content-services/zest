/*
Pulp 3 API

Testing ContentModulemdObsoletesAPIService

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

func Test_zest_ContentModulemdObsoletesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentModulemdObsoletesAPIService ContentRpmModulemdObsoletesCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.ContentModulemdObsoletesAPI.ContentRpmModulemdObsoletesCreate(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentModulemdObsoletesAPIService ContentRpmModulemdObsoletesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.ContentModulemdObsoletesAPI.ContentRpmModulemdObsoletesList(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentModulemdObsoletesAPIService ContentRpmModulemdObsoletesRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmModulemdObsoleteHref string

		resp, httpRes, err := apiClient.ContentModulemdObsoletesAPI.ContentRpmModulemdObsoletesRead(context.Background(), rpmModulemdObsoleteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentModulemdObsoletesAPIService ContentRpmModulemdObsoletesSetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmModulemdObsoleteHref string

		resp, httpRes, err := apiClient.ContentModulemdObsoletesAPI.ContentRpmModulemdObsoletesSetLabel(context.Background(), rpmModulemdObsoleteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentModulemdObsoletesAPIService ContentRpmModulemdObsoletesUnsetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmModulemdObsoleteHref string

		resp, httpRes, err := apiClient.ContentModulemdObsoletesAPI.ContentRpmModulemdObsoletesUnsetLabel(context.Background(), rpmModulemdObsoleteHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
