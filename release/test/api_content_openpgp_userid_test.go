/*
Pulp 3 API

Testing ContentOpenpgpUseridAPIService

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

func Test_zest_ContentOpenpgpUseridAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentOpenpgpUseridAPIService ContentCoreOpenpgpUseridList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.ContentOpenpgpUseridAPI.ContentCoreOpenpgpUseridList(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentOpenpgpUseridAPIService ContentCoreOpenpgpUseridRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var openPGPUserIDHref string

		resp, httpRes, err := apiClient.ContentOpenpgpUseridAPI.ContentCoreOpenpgpUseridRead(context.Background(), openPGPUserIDHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentOpenpgpUseridAPIService ContentCoreOpenpgpUseridSetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var openPGPUserIDHref string

		resp, httpRes, err := apiClient.ContentOpenpgpUseridAPI.ContentCoreOpenpgpUseridSetLabel(context.Background(), openPGPUserIDHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentOpenpgpUseridAPIService ContentCoreOpenpgpUseridUnsetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var openPGPUserIDHref string

		resp, httpRes, err := apiClient.ContentOpenpgpUseridAPI.ContentCoreOpenpgpUseridUnsetLabel(context.Background(), openPGPUserIDHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
