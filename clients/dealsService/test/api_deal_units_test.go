/*
DealsService

Testing DealUnitsAPIService

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

func Test_openapi_DealUnitsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DealUnitsAPIService ApiV2DealsServiceDealUnitsCountGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsCountGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DealUnitsAPIService ApiV2DealsServiceDealUnitsDealUnitIdCalculatePut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dealUnitId string

		resp, httpRes, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdCalculatePut(context.Background(), dealUnitId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DealUnitsAPIService ApiV2DealsServiceDealUnitsDealUnitIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dealUnitId string

		resp, httpRes, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdDelete(context.Background(), dealUnitId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DealUnitsAPIService ApiV2DealsServiceDealUnitsDealUnitIdExtendedGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dealUnitId string

		resp, httpRes, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdExtendedGet(context.Background(), dealUnitId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DealUnitsAPIService ApiV2DealsServiceDealUnitsDealUnitIdLinesCountGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dealUnitId string

		resp, httpRes, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdLinesCountGet(context.Background(), dealUnitId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DealUnitsAPIService ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdCalculatePut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dealUnitId string
		var dealUnitLineId string

		resp, httpRes, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdCalculatePut(context.Background(), dealUnitId, dealUnitLineId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DealUnitsAPIService ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dealUnitId string
		var dealUnitLineId string

		resp, httpRes, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdDelete(context.Background(), dealUnitId, dealUnitLineId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DealUnitsAPIService ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dealUnitId string
		var dealUnitLineId string

		resp, httpRes, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdGet(context.Background(), dealUnitId, dealUnitLineId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DealUnitsAPIService ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dealUnitId string
		var dealUnitLineId string

		resp, httpRes, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdPut(context.Background(), dealUnitId, dealUnitLineId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DealUnitsAPIService ApiV2DealsServiceDealUnitsDealUnitIdLinesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dealUnitId string

		resp, httpRes, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdLinesGet(context.Background(), dealUnitId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DealUnitsAPIService ApiV2DealsServiceDealUnitsDealUnitIdLinesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dealUnitId string

		resp, httpRes, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdLinesPost(context.Background(), dealUnitId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DealUnitsAPIService ApiV2DealsServiceDealUnitsDealUnitIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dealUnitId string

		resp, httpRes, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdPut(context.Background(), dealUnitId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DealUnitsAPIService ApiV2DealsServiceDealUnitsExtendedGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsExtendedGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DealUnitsAPIService ApiV2DealsServiceDealUnitsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DealUnitsAPIService ApiV2DealsServiceDealUnitsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DealUnitsAPIService GetDealUnitAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dealUnitId string

		resp, httpRes, err := apiClient.DealUnitsAPI.GetDealUnitAsync(context.Background(), dealUnitId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}