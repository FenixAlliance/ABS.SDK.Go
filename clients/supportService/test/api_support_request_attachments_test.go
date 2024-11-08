/*
SupportService

Testing SupportRequestAttachmentsAPIService

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

func Test_openapi_SupportRequestAttachmentsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SupportRequestAttachmentsAPIService ApiV2SupportServiceSupportRequestAttachmentsCountGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SupportRequestAttachmentsAPI.ApiV2SupportServiceSupportRequestAttachmentsCountGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SupportRequestAttachmentsAPIService ApiV2SupportServiceSupportRequestAttachmentsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SupportRequestAttachmentsAPI.ApiV2SupportServiceSupportRequestAttachmentsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SupportRequestAttachmentsAPIService ApiV2SupportServiceSupportRequestAttachmentsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SupportRequestAttachmentsAPI.ApiV2SupportServiceSupportRequestAttachmentsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SupportRequestAttachmentsAPIService ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var supportRequestAttachmentId string

		resp, httpRes, err := apiClient.SupportRequestAttachmentsAPI.ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdDelete(context.Background(), supportRequestAttachmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SupportRequestAttachmentsAPIService ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var supportRequestAttachmentId string

		resp, httpRes, err := apiClient.SupportRequestAttachmentsAPI.ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdGet(context.Background(), supportRequestAttachmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SupportRequestAttachmentsAPIService ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var supportRequestAttachmentId string

		resp, httpRes, err := apiClient.SupportRequestAttachmentsAPI.ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdPut(context.Background(), supportRequestAttachmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
