/*
Pulp 3 API

Testing ContentPackageenvironmentsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package zest

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/release"
)

func Test_zest_ContentPackageenvironmentsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentPackageenvironmentsApiService ContentRpmPackageenvironmentsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentPackageenvironmentsApi.ContentRpmPackageenvironmentsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackageenvironmentsApiService ContentRpmPackageenvironmentsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmPackageEnvironmentHref string

		resp, httpRes, err := apiClient.ContentPackageenvironmentsApi.ContentRpmPackageenvironmentsRead(context.Background(), rpmPackageEnvironmentHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
