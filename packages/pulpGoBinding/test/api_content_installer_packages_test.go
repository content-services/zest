/*
Pulp 3 API

Testing ContentInstallerPackagesApiService

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

func Test_pulpGoBinding_ContentInstallerPackagesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentInstallerPackagesApiService ContentDebInstallerPackagesCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentInstallerPackagesApi.ContentDebInstallerPackagesCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentInstallerPackagesApiService ContentDebInstallerPackagesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentInstallerPackagesApi.ContentDebInstallerPackagesList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentInstallerPackagesApiService ContentDebInstallerPackagesRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debInstallerPackageHref string

		resp, httpRes, err := apiClient.ContentInstallerPackagesApi.ContentDebInstallerPackagesRead(context.Background(), debInstallerPackageHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
