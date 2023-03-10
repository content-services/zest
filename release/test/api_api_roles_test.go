/*
Pulp 3 API

Testing ApiRolesApiService

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

func Test_zest_ApiRolesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ApiRolesApiService ApiV1RolesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleRoleHref string

		resp, httpRes, err := apiClient.ApiRolesApi.ApiV1RolesGet(context.Background(), ansibleRoleHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
