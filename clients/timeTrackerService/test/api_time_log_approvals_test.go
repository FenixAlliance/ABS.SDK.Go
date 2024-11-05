/*
TimeTrackerService

Testing TimeLogApprovalsAPIService

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

func Test_openapi_TimeLogApprovalsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TimeLogApprovalsAPIService ApiV2TimeTrackerServiceTimeLogApprovalsApprovalIdApproverPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var approvalId string

		httpRes, err := apiClient.TimeLogApprovalsAPI.ApiV2TimeTrackerServiceTimeLogApprovalsApprovalIdApproverPut(context.Background(), approvalId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TimeLogApprovalsAPIService ApiV2TimeTrackerServiceTimeLogApprovalsApprovalIdStatusPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var approvalId string

		httpRes, err := apiClient.TimeLogApprovalsAPI.ApiV2TimeTrackerServiceTimeLogApprovalsApprovalIdStatusPut(context.Background(), approvalId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TimeLogApprovalsAPIService ApiV2TimeTrackerServiceTimeLogApprovalsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.TimeLogApprovalsAPI.ApiV2TimeTrackerServiceTimeLogApprovalsPost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
