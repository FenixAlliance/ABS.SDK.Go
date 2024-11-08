/*
MarketingService

Testing EmailSignaturesAPIService

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

func Test_openapi_EmailSignaturesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EmailSignaturesAPIService ApiV2MarketingServiceEmailSignaturesCountGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EmailSignaturesAPI.ApiV2MarketingServiceEmailSignaturesCountGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailSignaturesAPIService ApiV2MarketingServiceEmailSignaturesEmailsignatureIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var emailsignatureId string

		resp, httpRes, err := apiClient.EmailSignaturesAPI.ApiV2MarketingServiceEmailSignaturesEmailsignatureIdDelete(context.Background(), emailsignatureId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailSignaturesAPIService ApiV2MarketingServiceEmailSignaturesEmailsignatureIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var emailsignatureId string

		resp, httpRes, err := apiClient.EmailSignaturesAPI.ApiV2MarketingServiceEmailSignaturesEmailsignatureIdGet(context.Background(), emailsignatureId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailSignaturesAPIService ApiV2MarketingServiceEmailSignaturesEmailsignatureIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var emailsignatureId string

		resp, httpRes, err := apiClient.EmailSignaturesAPI.ApiV2MarketingServiceEmailSignaturesEmailsignatureIdPut(context.Background(), emailsignatureId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailSignaturesAPIService ApiV2MarketingServiceEmailSignaturesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EmailSignaturesAPI.ApiV2MarketingServiceEmailSignaturesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailSignaturesAPIService ApiV2MarketingServiceEmailSignaturesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EmailSignaturesAPI.ApiV2MarketingServiceEmailSignaturesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
