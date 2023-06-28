/*
Pulp 3 API

Testing ContentPackageenvironmentsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package zest

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/release/v2023"
)

func Test_zest_ContentPackageenvironmentsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentPackageenvironmentsAPIService ContentRpmPackageenvironmentsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentPackageenvironmentsAPI.ContentRpmPackageenvironmentsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackageenvironmentsAPIService ContentRpmPackageenvironmentsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmPackageEnvironmentHref string

		resp, httpRes, err := apiClient.ContentPackageenvironmentsAPI.ContentRpmPackageenvironmentsRead(context.Background(), rpmPackageEnvironmentHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
