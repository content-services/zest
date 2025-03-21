/*
Pulp 3 API

Testing TaskSchedulesAPIService

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

func Test_zest_TaskSchedulesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TaskSchedulesAPIService TaskSchedulesAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskScheduleHref string

		resp, httpRes, err := apiClient.TaskSchedulesAPI.TaskSchedulesAddRole(context.Background(), taskScheduleHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskSchedulesAPIService TaskSchedulesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.TaskSchedulesAPI.TaskSchedulesList(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskSchedulesAPIService TaskSchedulesListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskScheduleHref string

		resp, httpRes, err := apiClient.TaskSchedulesAPI.TaskSchedulesListRoles(context.Background(), taskScheduleHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskSchedulesAPIService TaskSchedulesMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskScheduleHref string

		resp, httpRes, err := apiClient.TaskSchedulesAPI.TaskSchedulesMyPermissions(context.Background(), taskScheduleHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskSchedulesAPIService TaskSchedulesRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskScheduleHref string

		resp, httpRes, err := apiClient.TaskSchedulesAPI.TaskSchedulesRead(context.Background(), taskScheduleHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskSchedulesAPIService TaskSchedulesRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskScheduleHref string

		resp, httpRes, err := apiClient.TaskSchedulesAPI.TaskSchedulesRemoveRole(context.Background(), taskScheduleHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
