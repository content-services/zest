/*
Pulp 3 API

Testing DistributionsArtifactsAPIService

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

func Test_zest_DistributionsArtifactsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DistributionsArtifactsAPIService DistributionsCoreArtifactsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.DistributionsArtifactsAPI.DistributionsCoreArtifactsList(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DistributionsArtifactsAPIService DistributionsCoreArtifactsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var artifactDistributionHref string

		resp, httpRes, err := apiClient.DistributionsArtifactsAPI.DistributionsCoreArtifactsRead(context.Background(), artifactDistributionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
