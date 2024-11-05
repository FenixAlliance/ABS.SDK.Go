/*
WalletsService

Testing WalletsAPIService

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

func Test_openapi_WalletsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test WalletsAPIService CreateWalletLocationAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.CreateWalletLocationAsync(context.Background(), walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService DeleteWalletLocationAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string
		var locationId string

		resp, httpRes, err := apiClient.WalletsAPI.DeleteWalletLocationAsync(context.Background(), walletId, locationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService GetIncomingPaymentsAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.GetIncomingPaymentsAsync(context.Background(), walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService GetIncomingPaymentsCountAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.GetIncomingPaymentsCountAsync(context.Background(), walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService GetIncomingWalletInvoicesAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.GetIncomingWalletInvoicesAsync(context.Background(), walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService GetIncomingWalletInvoicesCountAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.GetIncomingWalletInvoicesCountAsync(context.Background(), walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService GetOutgoingPaymentsAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.GetOutgoingPaymentsAsync(context.Background(), walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService GetOutgoingPaymentsCountAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.GetOutgoingPaymentsCountAsync(context.Background(), walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService GetOutgoingWalletInvoicesAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.GetOutgoingWalletInvoicesAsync(context.Background(), walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService GetOutgoingWalletInvoicesCountAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.GetOutgoingWalletInvoicesCountAsync(context.Background(), walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService GetWalletDetailsAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.GetWalletDetailsAsync(context.Background(), walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService GetWalletExtendedOrdersAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.GetWalletExtendedOrdersAsync(context.Background(), walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService GetWalletInvoicesAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.GetWalletInvoicesAsync(context.Background(), walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService GetWalletInvoicesCountAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.GetWalletInvoicesCountAsync(context.Background(), walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService GetWalletLocationAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string
		var locationId string

		resp, httpRes, err := apiClient.WalletsAPI.GetWalletLocationAsync(context.Background(), walletId, locationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService GetWalletLocationsAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.GetWalletLocationsAsync(context.Background(), walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService GetWalletLocationsCountAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.GetWalletLocationsCountAsync(context.Background(), walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService GetWalletOrdersAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.GetWalletOrdersAsync(context.Background(), walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService GetWalletOrdersCountAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.GetWalletOrdersCountAsync(context.Background(), walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService GetWalletPaymentsAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.GetWalletPaymentsAsync(context.Background(), walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService GetWalletPaymentsCountAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsAPI.GetWalletPaymentsCountAsync(context.Background(), walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsAPIService UpdateWalletLocationAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string
		var locationId string

		resp, httpRes, err := apiClient.WalletsAPI.UpdateWalletLocationAsync(context.Background(), walletId, locationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}