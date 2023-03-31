/*
Pulp 3 API

Testing PulpAnsibleApiApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package zest

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/release/v3"
)

func Test_zest_PulpAnsibleApiApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PulpAnsibleApiApiService PulpAnsibleGalaxyApiGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var path string

		httpRes, err := apiClient.PulpAnsibleApiApi.PulpAnsibleGalaxyApiGet(context.Background(), path).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
