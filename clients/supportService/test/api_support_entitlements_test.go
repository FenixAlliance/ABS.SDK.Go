/*
SupportService

Testing SupportEntitlementsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_SupportEntitlementsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SupportEntitlementsAPIService ApiV2SupportServiceSupportEntitlementsCountGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SupportEntitlementsAPI.ApiV2SupportServiceSupportEntitlementsCountGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SupportEntitlementsAPIService ApiV2SupportServiceSupportEntitlementsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SupportEntitlementsAPI.ApiV2SupportServiceSupportEntitlementsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SupportEntitlementsAPIService ApiV2SupportServiceSupportEntitlementsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SupportEntitlementsAPI.ApiV2SupportServiceSupportEntitlementsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SupportEntitlementsAPIService ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var supportEntitlementId string

		resp, httpRes, err := apiClient.SupportEntitlementsAPI.ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdDelete(context.Background(), supportEntitlementId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SupportEntitlementsAPIService ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var supportEntitlementId string

		resp, httpRes, err := apiClient.SupportEntitlementsAPI.ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdGet(context.Background(), supportEntitlementId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SupportEntitlementsAPIService ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var supportEntitlementId string

		resp, httpRes, err := apiClient.SupportEntitlementsAPI.ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdPut(context.Background(), supportEntitlementId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
