/*
Pulp 3 API

Testing PublicationsVerbatimApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package pulpGoBinding

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/pulpGoBinding/packages/pulpGoBinding"
)

func Test_pulpGoBinding_PublicationsVerbatimApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PublicationsVerbatimApiService PublicationsDebVerbatimCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PublicationsVerbatimApi.PublicationsDebVerbatimCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PublicationsVerbatimApiService PublicationsDebVerbatimDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debVerbatimPublicationHref string

		httpRes, err := apiClient.PublicationsVerbatimApi.PublicationsDebVerbatimDelete(context.Background(), debVerbatimPublicationHref).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PublicationsVerbatimApiService PublicationsDebVerbatimList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PublicationsVerbatimApi.PublicationsDebVerbatimList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PublicationsVerbatimApiService PublicationsDebVerbatimRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debVerbatimPublicationHref string

		resp, httpRes, err := apiClient.PublicationsVerbatimApi.PublicationsDebVerbatimRead(context.Background(), debVerbatimPublicationHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}