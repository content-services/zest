/*
Pulp 3 API

Testing PulpAnsibleDefaultApiV3NamespacesAPIService

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

func Test_zest_PulpAnsibleDefaultApiV3NamespacesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PulpAnsibleDefaultApiV3NamespacesAPIService PulpAnsibleGalaxyDefaultApiV3NamespacesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PulpAnsibleDefaultApiV3NamespacesAPI.PulpAnsibleGalaxyDefaultApiV3NamespacesList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleDefaultApiV3NamespacesAPIService PulpAnsibleGalaxyDefaultApiV3NamespacesRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.PulpAnsibleDefaultApiV3NamespacesAPI.PulpAnsibleGalaxyDefaultApiV3NamespacesRead(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
