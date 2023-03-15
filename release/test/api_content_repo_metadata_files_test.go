/*
Pulp 3 API

Testing ContentRepoMetadataFilesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package zest/release/v3

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/release/v3"
)

func Test_zest/release/v3_ContentRepoMetadataFilesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentRepoMetadataFilesApiService ContentRpmRepoMetadataFilesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentRepoMetadataFilesApi.ContentRpmRepoMetadataFilesList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentRepoMetadataFilesApiService ContentRpmRepoMetadataFilesRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRepoMetadataFileHref string

		resp, httpRes, err := apiClient.ContentRepoMetadataFilesApi.ContentRpmRepoMetadataFilesRead(context.Background(), rpmRepoMetadataFileHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
