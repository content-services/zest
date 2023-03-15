/*
Pulp 3 API

Testing DistributionsAnsibleApiService

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

func Test_zest_DistributionsAnsibleApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DistributionsAnsibleApiService DistributionsAnsibleAnsibleCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DistributionsAnsibleApi.DistributionsAnsibleAnsibleCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsAnsibleApiService DistributionsAnsibleAnsibleDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleAnsibleDistributionHref string

		resp, httpRes, err := apiClient.DistributionsAnsibleApi.DistributionsAnsibleAnsibleDelete(context.Background(), ansibleAnsibleDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsAnsibleApiService DistributionsAnsibleAnsibleList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DistributionsAnsibleApi.DistributionsAnsibleAnsibleList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsAnsibleApiService DistributionsAnsibleAnsiblePartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleAnsibleDistributionHref string

		resp, httpRes, err := apiClient.DistributionsAnsibleApi.DistributionsAnsibleAnsiblePartialUpdate(context.Background(), ansibleAnsibleDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsAnsibleApiService DistributionsAnsibleAnsibleRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleAnsibleDistributionHref string

		resp, httpRes, err := apiClient.DistributionsAnsibleApi.DistributionsAnsibleAnsibleRead(context.Background(), ansibleAnsibleDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsAnsibleApiService DistributionsAnsibleAnsibleUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleAnsibleDistributionHref string

		resp, httpRes, err := apiClient.DistributionsAnsibleApi.DistributionsAnsibleAnsibleUpdate(context.Background(), ansibleAnsibleDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
