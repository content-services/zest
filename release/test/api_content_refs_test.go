/*
Pulp 3 API

Testing ContentRefsAPIService

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

func Test_zest_ContentRefsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentRefsAPIService ContentOstreeRefsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.ContentRefsAPI.ContentOstreeRefsList(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentRefsAPIService ContentOstreeRefsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ostreeOstreeRefHref string

		resp, httpRes, err := apiClient.ContentRefsAPI.ContentOstreeRefsRead(context.Background(), ostreeOstreeRefHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
