/*
Pulp 3 API

Testing ContentRepoMetadataFilesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package zest

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/release/v2025"
)

func Test_zest_ContentRepoMetadataFilesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentRepoMetadataFilesAPIService ContentRpmRepoMetadataFilesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.ContentRepoMetadataFilesAPI.ContentRpmRepoMetadataFilesList(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentRepoMetadataFilesAPIService ContentRpmRepoMetadataFilesRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRepoMetadataFileHref string

		resp, httpRes, err := apiClient.ContentRepoMetadataFilesAPI.ContentRpmRepoMetadataFilesRead(context.Background(), rpmRepoMetadataFileHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
