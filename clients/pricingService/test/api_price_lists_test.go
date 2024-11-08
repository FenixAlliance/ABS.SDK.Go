/*
PricingService

Testing PriceListsAPIService

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

func Test_openapi_PriceListsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PriceListsAPIService ApiV2PricingServicePriceListsCountGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PriceListsAPI.ApiV2PricingServicePriceListsCountGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PriceListsAPIService ApiV2PricingServicePriceListsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PriceListsAPI.ApiV2PricingServicePriceListsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PriceListsAPIService ApiV2PricingServicePriceListsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PriceListsAPI.ApiV2PricingServicePriceListsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PriceListsAPIService ApiV2PricingServicePriceListsPriceListIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var priceListId string

		resp, httpRes, err := apiClient.PriceListsAPI.ApiV2PricingServicePriceListsPriceListIdDelete(context.Background(), priceListId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PriceListsAPIService ApiV2PricingServicePriceListsPriceListIdPricesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var priceListId string

		resp, httpRes, err := apiClient.PriceListsAPI.ApiV2PricingServicePriceListsPriceListIdPricesPost(context.Background(), priceListId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PriceListsAPIService ApiV2PricingServicePriceListsPriceListIdPricesPriceIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var priceListId string
		var priceId string

		resp, httpRes, err := apiClient.PriceListsAPI.ApiV2PricingServicePriceListsPriceListIdPricesPriceIdDelete(context.Background(), priceListId, priceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PriceListsAPIService ApiV2PricingServicePriceListsPriceListIdPricesPriceIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var priceListId string
		var priceId string

		resp, httpRes, err := apiClient.PriceListsAPI.ApiV2PricingServicePriceListsPriceListIdPricesPriceIdPut(context.Background(), priceListId, priceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PriceListsAPIService ApiV2PricingServicePriceListsPriceListIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var priceListId string

		resp, httpRes, err := apiClient.PriceListsAPI.ApiV2PricingServicePriceListsPriceListIdPut(context.Background(), priceListId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PriceListsAPIService GetPriceListAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var priceListId string

		resp, httpRes, err := apiClient.PriceListsAPI.GetPriceListAsync(context.Background(), priceListId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PriceListsAPIService GetPriceListPriceAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var priceListId string
		var priceId string

		resp, httpRes, err := apiClient.PriceListsAPI.GetPriceListPriceAsync(context.Background(), priceListId, priceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PriceListsAPIService GetPriceListPricesAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var priceListId string

		resp, httpRes, err := apiClient.PriceListsAPI.GetPriceListPricesAsync(context.Background(), priceListId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
