/*
Pulp 3 API

Testing DomainsApiService

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

func Test_zest_DomainsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DomainsApiService DomainsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DomainsApi.DomainsCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DomainsApiService DomainsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domainHref string

		resp, httpRes, err := apiClient.DomainsApi.DomainsDelete(context.Background(), domainHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DomainsApiService DomainsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DomainsApi.DomainsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DomainsApiService DomainsPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domainHref string

		resp, httpRes, err := apiClient.DomainsApi.DomainsPartialUpdate(context.Background(), domainHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DomainsApiService DomainsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domainHref string

		resp, httpRes, err := apiClient.DomainsApi.DomainsRead(context.Background(), domainHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DomainsApiService DomainsUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domainHref string

		resp, httpRes, err := apiClient.DomainsApi.DomainsUpdate(context.Background(), domainHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
