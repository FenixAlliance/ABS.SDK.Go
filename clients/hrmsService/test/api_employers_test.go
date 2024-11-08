/*
HrmsService

Testing EmployersAPIService

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

func Test_openapi_EmployersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EmployersAPIService CreateEmployerAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.EmployersAPI.CreateEmployerAsync(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmployersAPIService DeleteEmployerAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var employerId string

		httpRes, err := apiClient.EmployersAPI.DeleteEmployerAsync(context.Background(), employerId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmployersAPIService GetEmployerByIdAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var employerId string

		resp, httpRes, err := apiClient.EmployersAPI.GetEmployerByIdAsync(context.Background(), employerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmployersAPIService GetEmployersAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EmployersAPI.GetEmployersAsync(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmployersAPIService GetEmployersCountAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EmployersAPI.GetEmployersCountAsync(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmployersAPIService UpdateEmployerAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var employerId string

		httpRes, err := apiClient.EmployersAPI.UpdateEmployerAsync(context.Background(), employerId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
