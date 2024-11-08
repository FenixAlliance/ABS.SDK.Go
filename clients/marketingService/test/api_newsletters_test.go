/*
MarketingService

Testing NewslettersAPIService

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

func Test_openapi_NewslettersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NewslettersAPIService ApiV2MarketingServiceNewslettersCountGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NewslettersAPI.ApiV2MarketingServiceNewslettersCountGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NewslettersAPIService ApiV2MarketingServiceNewslettersGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.NewslettersAPI.ApiV2MarketingServiceNewslettersGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NewslettersAPIService ApiV2MarketingServiceNewslettersNewsletterIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var newsletterId string

		resp, httpRes, err := apiClient.NewslettersAPI.ApiV2MarketingServiceNewslettersNewsletterIdDelete(context.Background(), newsletterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NewslettersAPIService ApiV2MarketingServiceNewslettersNewsletterIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var newsletterId string

		resp, httpRes, err := apiClient.NewslettersAPI.ApiV2MarketingServiceNewslettersNewsletterIdGet(context.Background(), newsletterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NewslettersAPIService ApiV2MarketingServiceNewslettersNewsletterIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var newsletterId string

		resp, httpRes, err := apiClient.NewslettersAPI.ApiV2MarketingServiceNewslettersNewsletterIdPut(context.Background(), newsletterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NewslettersAPIService ApiV2MarketingServiceNewslettersPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NewslettersAPI.ApiV2MarketingServiceNewslettersPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
