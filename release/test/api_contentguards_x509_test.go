/*
Pulp 3 API

Testing ContentguardsX509ApiService

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

func Test_zest/release/v3_ContentguardsX509ApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentguardsX509ApiService ContentguardsCertguardX509Create", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentguardsX509Api.ContentguardsCertguardX509Create(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsX509ApiService ContentguardsCertguardX509Delete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var certguardX509CertGuardHref string

		httpRes, err := apiClient.ContentguardsX509Api.ContentguardsCertguardX509Delete(context.Background(), certguardX509CertGuardHref).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsX509ApiService ContentguardsCertguardX509List", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentguardsX509Api.ContentguardsCertguardX509List(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsX509ApiService ContentguardsCertguardX509PartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var certguardX509CertGuardHref string

		resp, httpRes, err := apiClient.ContentguardsX509Api.ContentguardsCertguardX509PartialUpdate(context.Background(), certguardX509CertGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsX509ApiService ContentguardsCertguardX509Read", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var certguardX509CertGuardHref string

		resp, httpRes, err := apiClient.ContentguardsX509Api.ContentguardsCertguardX509Read(context.Background(), certguardX509CertGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsX509ApiService ContentguardsCertguardX509Update", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var certguardX509CertGuardHref string

		resp, httpRes, err := apiClient.ContentguardsX509Api.ContentguardsCertguardX509Update(context.Background(), certguardX509CertGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
