/*
Pulp 3 API

Testing ContentAdvisoriesAPIService

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

func Test_zest_ContentAdvisoriesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentAdvisoriesAPIService ContentRpmAdvisoriesCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.ContentAdvisoriesAPI.ContentRpmAdvisoriesCreate(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentAdvisoriesAPIService ContentRpmAdvisoriesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.ContentAdvisoriesAPI.ContentRpmAdvisoriesList(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentAdvisoriesAPIService ContentRpmAdvisoriesRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmUpdateRecordHref string

		resp, httpRes, err := apiClient.ContentAdvisoriesAPI.ContentRpmAdvisoriesRead(context.Background(), rpmUpdateRecordHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentAdvisoriesAPIService ContentRpmAdvisoriesSetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmUpdateRecordHref string

		resp, httpRes, err := apiClient.ContentAdvisoriesAPI.ContentRpmAdvisoriesSetLabel(context.Background(), rpmUpdateRecordHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentAdvisoriesAPIService ContentRpmAdvisoriesUnsetLabel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmUpdateRecordHref string

		resp, httpRes, err := apiClient.ContentAdvisoriesAPI.ContentRpmAdvisoriesUnsetLabel(context.Background(), rpmUpdateRecordHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
