/*
Pulp 3 API

Testing PypiAPIService

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

func Test_zest_PypiAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PypiAPIService PypiRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var path string
		var pulpDomain string

		resp, httpRes, err := apiClient.PypiAPI.PypiRead(context.Background(), path, pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}