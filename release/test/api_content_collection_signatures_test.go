/*
Pulp 3 API

Testing ContentCollectionSignaturesApiService

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

func Test_zest_ContentCollectionSignaturesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentCollectionSignaturesApiService ContentAnsibleCollectionSignaturesCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentCollectionSignaturesApi.ContentAnsibleCollectionSignaturesCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentCollectionSignaturesApiService ContentAnsibleCollectionSignaturesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentCollectionSignaturesApi.ContentAnsibleCollectionSignaturesList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentCollectionSignaturesApiService ContentAnsibleCollectionSignaturesRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleCollectionVersionSignatureHref string

		resp, httpRes, err := apiClient.ContentCollectionSignaturesApi.ContentAnsibleCollectionSignaturesRead(context.Background(), ansibleCollectionVersionSignatureHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
