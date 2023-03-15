/*
Pulp 3 API

Testing ContentArtifactApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/v3/release--package-name&#x3D;zest"
)

func Test_openapi_ContentArtifactApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentArtifactApiService ContentMavenArtifactCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentArtifactApi.ContentMavenArtifactCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentArtifactApiService ContentMavenArtifactList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentArtifactApi.ContentMavenArtifactList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentArtifactApiService ContentMavenArtifactRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mavenMavenArtifactHref string

		resp, httpRes, err := apiClient.ContentArtifactApi.ContentMavenArtifactRead(context.Background(), mavenMavenArtifactHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
