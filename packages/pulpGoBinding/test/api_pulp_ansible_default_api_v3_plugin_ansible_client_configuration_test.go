/*
Pulp 3 API

Testing PulpAnsibleDefaultApiV3PluginAnsibleClientConfigurationApiService

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

func Test_pulpGoBinding_PulpAnsibleDefaultApiV3PluginAnsibleClientConfigurationApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PulpAnsibleDefaultApiV3PluginAnsibleClientConfigurationApiService PulpAnsibleGalaxyDefaultApiV3PluginAnsibleClientConfigurationGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleClientConfigurationApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleClientConfigurationGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
