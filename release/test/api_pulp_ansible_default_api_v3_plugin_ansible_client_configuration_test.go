/*
Pulp 3 API

Testing PulpAnsibleDefaultApiV3PluginAnsibleClientConfigurationAPIService

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

func Test_zest_PulpAnsibleDefaultApiV3PluginAnsibleClientConfigurationAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PulpAnsibleDefaultApiV3PluginAnsibleClientConfigurationAPIService PulpAnsibleGalaxyDefaultApiV3PluginAnsibleClientConfigurationRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleClientConfigurationAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleClientConfigurationRead(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
