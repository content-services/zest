/*
Pulp 3 API

Testing ContentReleaseArchitecturesApiService

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

func Test_pulpGoBinding_ContentReleaseArchitecturesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentReleaseArchitecturesApiService ContentDebReleaseArchitecturesCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentReleaseArchitecturesApi.ContentDebReleaseArchitecturesCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentReleaseArchitecturesApiService ContentDebReleaseArchitecturesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentReleaseArchitecturesApi.ContentDebReleaseArchitecturesList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentReleaseArchitecturesApiService ContentDebReleaseArchitecturesRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debReleaseArchitectureHref string

		resp, httpRes, err := apiClient.ContentReleaseArchitecturesApi.ContentDebReleaseArchitecturesRead(context.Background(), debReleaseArchitectureHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
