/*
Pulp 3 API

Testing DomainsAPIService

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

func Test_zest_DomainsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DomainsAPIService DomainsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.DomainsAPI.DomainsCreate(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DomainsAPIService DomainsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domainHref string

		resp, httpRes, err := apiClient.DomainsAPI.DomainsDelete(context.Background(), domainHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DomainsAPIService DomainsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.DomainsAPI.DomainsList(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DomainsAPIService DomainsMigrate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.DomainsAPI.DomainsMigrate(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DomainsAPIService DomainsPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domainHref string

		resp, httpRes, err := apiClient.DomainsAPI.DomainsPartialUpdate(context.Background(), domainHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DomainsAPIService DomainsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domainHref string

		resp, httpRes, err := apiClient.DomainsAPI.DomainsRead(context.Background(), domainHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DomainsAPIService DomainsSetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domainHref string

		resp, httpRes, err := apiClient.DomainsAPI.DomainsSetLabel(context.Background(), domainHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DomainsAPIService DomainsUnsetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domainHref string

		resp, httpRes, err := apiClient.DomainsAPI.DomainsUnsetLabel(context.Background(), domainHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DomainsAPIService DomainsUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domainHref string

		resp, httpRes, err := apiClient.DomainsAPI.DomainsUpdate(context.Background(), domainHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
