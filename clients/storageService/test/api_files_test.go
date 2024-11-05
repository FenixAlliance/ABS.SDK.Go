/*
StorageService

Testing FilesAPIService

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

func Test_openapi_FilesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FilesAPIService CreateFileAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FilesAPI.CreateFileAsync(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FilesAPIService DeleteFileAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileId string

		resp, httpRes, err := apiClient.FilesAPI.DeleteFileAsync(context.Background(), fileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FilesAPIService DownloadFileAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileId string

		resp, httpRes, err := apiClient.FilesAPI.DownloadFileAsync(context.Background(), fileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FilesAPIService GetFileAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileId string

		resp, httpRes, err := apiClient.FilesAPI.GetFileAsync(context.Background(), fileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FilesAPIService GetFilesAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FilesAPI.GetFilesAsync(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FilesAPIService UpdateFileAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileId string

		resp, httpRes, err := apiClient.FilesAPI.UpdateFileAsync(context.Background(), fileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}