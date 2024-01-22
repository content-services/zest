/*
Pulp 3 API

Testing PublicationsGemAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package zest

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/release/v2024"
)

func Test_zest_PublicationsGemAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PublicationsGemAPIService PublicationsGemGemAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gemGemPublicationHref string

		resp, httpRes, err := apiClient.PublicationsGemAPI.PublicationsGemGemAddRole(context.Background(), gemGemPublicationHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PublicationsGemAPIService PublicationsGemGemCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.PublicationsGemAPI.PublicationsGemGemCreate(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PublicationsGemAPIService PublicationsGemGemDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gemGemPublicationHref string

		httpRes, err := apiClient.PublicationsGemAPI.PublicationsGemGemDelete(context.Background(), gemGemPublicationHref).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PublicationsGemAPIService PublicationsGemGemList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.PublicationsGemAPI.PublicationsGemGemList(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PublicationsGemAPIService PublicationsGemGemListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gemGemPublicationHref string

		resp, httpRes, err := apiClient.PublicationsGemAPI.PublicationsGemGemListRoles(context.Background(), gemGemPublicationHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PublicationsGemAPIService PublicationsGemGemMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gemGemPublicationHref string

		resp, httpRes, err := apiClient.PublicationsGemAPI.PublicationsGemGemMyPermissions(context.Background(), gemGemPublicationHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PublicationsGemAPIService PublicationsGemGemRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gemGemPublicationHref string

		resp, httpRes, err := apiClient.PublicationsGemAPI.PublicationsGemGemRead(context.Background(), gemGemPublicationHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PublicationsGemAPIService PublicationsGemGemRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var gemGemPublicationHref string

		resp, httpRes, err := apiClient.PublicationsGemAPI.PublicationsGemGemRemoveRole(context.Background(), gemGemPublicationHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
