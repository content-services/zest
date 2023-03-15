/*
Pulp 3 API

Testing TaskSchedulesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package zest

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/v3/release"
)

func Test_zest_TaskSchedulesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TaskSchedulesApiService TaskSchedulesAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskScheduleHref string

		resp, httpRes, err := apiClient.TaskSchedulesApi.TaskSchedulesAddRole(context.Background(), taskScheduleHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskSchedulesApiService TaskSchedulesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TaskSchedulesApi.TaskSchedulesList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskSchedulesApiService TaskSchedulesListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskScheduleHref string

		resp, httpRes, err := apiClient.TaskSchedulesApi.TaskSchedulesListRoles(context.Background(), taskScheduleHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskSchedulesApiService TaskSchedulesMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskScheduleHref string

		resp, httpRes, err := apiClient.TaskSchedulesApi.TaskSchedulesMyPermissions(context.Background(), taskScheduleHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskSchedulesApiService TaskSchedulesRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskScheduleHref string

		resp, httpRes, err := apiClient.TaskSchedulesApi.TaskSchedulesRead(context.Background(), taskScheduleHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskSchedulesApiService TaskSchedulesRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskScheduleHref string

		resp, httpRes, err := apiClient.TaskSchedulesApi.TaskSchedulesRemoveRole(context.Background(), taskScheduleHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
