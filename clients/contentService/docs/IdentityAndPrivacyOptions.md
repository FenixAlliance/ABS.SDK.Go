# IdentityAndPrivacyOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowGuestOrders** | Pointer to **bool** |  | [optional] 
**AllowGuestCartRecognition** | Pointer to **bool** |  | [optional] 
**AllowRemoveDownloadAccessOnRequest** | Pointer to **bool** |  | [optional] 
**AllowRemovePersonalDataFromOrdersOnRequest** | Pointer to **bool** |  | [optional] 
**AllowRemovePersonalDataFromSubscriptionsOnRequest** | Pointer to **bool** |  | [optional] 
**StoreCheckoutPrivacyPolicyNotice** | Pointer to **NullableString** |  | [optional] 
**StoreRegistrationPrivacyPolicyNotice** | Pointer to **NullableString** |  | [optional] 
**DefaultCustomerLocation** | Pointer to **string** |  | [optional] 
**InactiveCartsRetentionPolicy** | Pointer to [**StoreDataRetentionPolicy**](StoreDataRetentionPolicy.md) |  | [optional] 
**PendingOrdersRetentionPolicy** | Pointer to [**StoreDataRetentionPolicy**](StoreDataRetentionPolicy.md) |  | [optional] 
**FailedOrdersRetentionPolicy** | Pointer to [**StoreDataRetentionPolicy**](StoreDataRetentionPolicy.md) |  | [optional] 
**CancelledOrdersRetentionPolicy** | Pointer to [**StoreDataRetentionPolicy**](StoreDataRetentionPolicy.md) |  | [optional] 
**CompletedOrdersRetentionPolicy** | Pointer to [**StoreDataRetentionPolicy**](StoreDataRetentionPolicy.md) |  | [optional] 

## Methods

### NewIdentityAndPrivacyOptions

`func NewIdentityAndPrivacyOptions() *IdentityAndPrivacyOptions`

NewIdentityAndPrivacyOptions instantiates a new IdentityAndPrivacyOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityAndPrivacyOptionsWithDefaults

`func NewIdentityAndPrivacyOptionsWithDefaults() *IdentityAndPrivacyOptions`

NewIdentityAndPrivacyOptionsWithDefaults instantiates a new IdentityAndPrivacyOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowGuestOrders

`func (o *IdentityAndPrivacyOptions) GetAllowGuestOrders() bool`

GetAllowGuestOrders returns the AllowGuestOrders field if non-nil, zero value otherwise.

### GetAllowGuestOrdersOk

`func (o *IdentityAndPrivacyOptions) GetAllowGuestOrdersOk() (*bool, bool)`

GetAllowGuestOrdersOk returns a tuple with the AllowGuestOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGuestOrders

`func (o *IdentityAndPrivacyOptions) SetAllowGuestOrders(v bool)`

SetAllowGuestOrders sets AllowGuestOrders field to given value.

### HasAllowGuestOrders

`func (o *IdentityAndPrivacyOptions) HasAllowGuestOrders() bool`

HasAllowGuestOrders returns a boolean if a field has been set.

### GetAllowGuestCartRecognition

`func (o *IdentityAndPrivacyOptions) GetAllowGuestCartRecognition() bool`

GetAllowGuestCartRecognition returns the AllowGuestCartRecognition field if non-nil, zero value otherwise.

### GetAllowGuestCartRecognitionOk

`func (o *IdentityAndPrivacyOptions) GetAllowGuestCartRecognitionOk() (*bool, bool)`

GetAllowGuestCartRecognitionOk returns a tuple with the AllowGuestCartRecognition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGuestCartRecognition

`func (o *IdentityAndPrivacyOptions) SetAllowGuestCartRecognition(v bool)`

SetAllowGuestCartRecognition sets AllowGuestCartRecognition field to given value.

### HasAllowGuestCartRecognition

`func (o *IdentityAndPrivacyOptions) HasAllowGuestCartRecognition() bool`

HasAllowGuestCartRecognition returns a boolean if a field has been set.

### GetAllowRemoveDownloadAccessOnRequest

`func (o *IdentityAndPrivacyOptions) GetAllowRemoveDownloadAccessOnRequest() bool`

GetAllowRemoveDownloadAccessOnRequest returns the AllowRemoveDownloadAccessOnRequest field if non-nil, zero value otherwise.

### GetAllowRemoveDownloadAccessOnRequestOk

`func (o *IdentityAndPrivacyOptions) GetAllowRemoveDownloadAccessOnRequestOk() (*bool, bool)`

GetAllowRemoveDownloadAccessOnRequestOk returns a tuple with the AllowRemoveDownloadAccessOnRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRemoveDownloadAccessOnRequest

`func (o *IdentityAndPrivacyOptions) SetAllowRemoveDownloadAccessOnRequest(v bool)`

SetAllowRemoveDownloadAccessOnRequest sets AllowRemoveDownloadAccessOnRequest field to given value.

### HasAllowRemoveDownloadAccessOnRequest

`func (o *IdentityAndPrivacyOptions) HasAllowRemoveDownloadAccessOnRequest() bool`

HasAllowRemoveDownloadAccessOnRequest returns a boolean if a field has been set.

### GetAllowRemovePersonalDataFromOrdersOnRequest

`func (o *IdentityAndPrivacyOptions) GetAllowRemovePersonalDataFromOrdersOnRequest() bool`

GetAllowRemovePersonalDataFromOrdersOnRequest returns the AllowRemovePersonalDataFromOrdersOnRequest field if non-nil, zero value otherwise.

### GetAllowRemovePersonalDataFromOrdersOnRequestOk

`func (o *IdentityAndPrivacyOptions) GetAllowRemovePersonalDataFromOrdersOnRequestOk() (*bool, bool)`

GetAllowRemovePersonalDataFromOrdersOnRequestOk returns a tuple with the AllowRemovePersonalDataFromOrdersOnRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRemovePersonalDataFromOrdersOnRequest

`func (o *IdentityAndPrivacyOptions) SetAllowRemovePersonalDataFromOrdersOnRequest(v bool)`

SetAllowRemovePersonalDataFromOrdersOnRequest sets AllowRemovePersonalDataFromOrdersOnRequest field to given value.

### HasAllowRemovePersonalDataFromOrdersOnRequest

`func (o *IdentityAndPrivacyOptions) HasAllowRemovePersonalDataFromOrdersOnRequest() bool`

HasAllowRemovePersonalDataFromOrdersOnRequest returns a boolean if a field has been set.

### GetAllowRemovePersonalDataFromSubscriptionsOnRequest

`func (o *IdentityAndPrivacyOptions) GetAllowRemovePersonalDataFromSubscriptionsOnRequest() bool`

GetAllowRemovePersonalDataFromSubscriptionsOnRequest returns the AllowRemovePersonalDataFromSubscriptionsOnRequest field if non-nil, zero value otherwise.

### GetAllowRemovePersonalDataFromSubscriptionsOnRequestOk

`func (o *IdentityAndPrivacyOptions) GetAllowRemovePersonalDataFromSubscriptionsOnRequestOk() (*bool, bool)`

GetAllowRemovePersonalDataFromSubscriptionsOnRequestOk returns a tuple with the AllowRemovePersonalDataFromSubscriptionsOnRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRemovePersonalDataFromSubscriptionsOnRequest

`func (o *IdentityAndPrivacyOptions) SetAllowRemovePersonalDataFromSubscriptionsOnRequest(v bool)`

SetAllowRemovePersonalDataFromSubscriptionsOnRequest sets AllowRemovePersonalDataFromSubscriptionsOnRequest field to given value.

### HasAllowRemovePersonalDataFromSubscriptionsOnRequest

`func (o *IdentityAndPrivacyOptions) HasAllowRemovePersonalDataFromSubscriptionsOnRequest() bool`

HasAllowRemovePersonalDataFromSubscriptionsOnRequest returns a boolean if a field has been set.

### GetStoreCheckoutPrivacyPolicyNotice

`func (o *IdentityAndPrivacyOptions) GetStoreCheckoutPrivacyPolicyNotice() string`

GetStoreCheckoutPrivacyPolicyNotice returns the StoreCheckoutPrivacyPolicyNotice field if non-nil, zero value otherwise.

### GetStoreCheckoutPrivacyPolicyNoticeOk

`func (o *IdentityAndPrivacyOptions) GetStoreCheckoutPrivacyPolicyNoticeOk() (*string, bool)`

GetStoreCheckoutPrivacyPolicyNoticeOk returns a tuple with the StoreCheckoutPrivacyPolicyNotice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreCheckoutPrivacyPolicyNotice

`func (o *IdentityAndPrivacyOptions) SetStoreCheckoutPrivacyPolicyNotice(v string)`

SetStoreCheckoutPrivacyPolicyNotice sets StoreCheckoutPrivacyPolicyNotice field to given value.

### HasStoreCheckoutPrivacyPolicyNotice

`func (o *IdentityAndPrivacyOptions) HasStoreCheckoutPrivacyPolicyNotice() bool`

HasStoreCheckoutPrivacyPolicyNotice returns a boolean if a field has been set.

### SetStoreCheckoutPrivacyPolicyNoticeNil

`func (o *IdentityAndPrivacyOptions) SetStoreCheckoutPrivacyPolicyNoticeNil(b bool)`

 SetStoreCheckoutPrivacyPolicyNoticeNil sets the value for StoreCheckoutPrivacyPolicyNotice to be an explicit nil

### UnsetStoreCheckoutPrivacyPolicyNotice
`func (o *IdentityAndPrivacyOptions) UnsetStoreCheckoutPrivacyPolicyNotice()`

UnsetStoreCheckoutPrivacyPolicyNotice ensures that no value is present for StoreCheckoutPrivacyPolicyNotice, not even an explicit nil
### GetStoreRegistrationPrivacyPolicyNotice

`func (o *IdentityAndPrivacyOptions) GetStoreRegistrationPrivacyPolicyNotice() string`

GetStoreRegistrationPrivacyPolicyNotice returns the StoreRegistrationPrivacyPolicyNotice field if non-nil, zero value otherwise.

### GetStoreRegistrationPrivacyPolicyNoticeOk

`func (o *IdentityAndPrivacyOptions) GetStoreRegistrationPrivacyPolicyNoticeOk() (*string, bool)`

GetStoreRegistrationPrivacyPolicyNoticeOk returns a tuple with the StoreRegistrationPrivacyPolicyNotice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreRegistrationPrivacyPolicyNotice

`func (o *IdentityAndPrivacyOptions) SetStoreRegistrationPrivacyPolicyNotice(v string)`

SetStoreRegistrationPrivacyPolicyNotice sets StoreRegistrationPrivacyPolicyNotice field to given value.

### HasStoreRegistrationPrivacyPolicyNotice

`func (o *IdentityAndPrivacyOptions) HasStoreRegistrationPrivacyPolicyNotice() bool`

HasStoreRegistrationPrivacyPolicyNotice returns a boolean if a field has been set.

### SetStoreRegistrationPrivacyPolicyNoticeNil

`func (o *IdentityAndPrivacyOptions) SetStoreRegistrationPrivacyPolicyNoticeNil(b bool)`

 SetStoreRegistrationPrivacyPolicyNoticeNil sets the value for StoreRegistrationPrivacyPolicyNotice to be an explicit nil

### UnsetStoreRegistrationPrivacyPolicyNotice
`func (o *IdentityAndPrivacyOptions) UnsetStoreRegistrationPrivacyPolicyNotice()`

UnsetStoreRegistrationPrivacyPolicyNotice ensures that no value is present for StoreRegistrationPrivacyPolicyNotice, not even an explicit nil
### GetDefaultCustomerLocation

`func (o *IdentityAndPrivacyOptions) GetDefaultCustomerLocation() string`

GetDefaultCustomerLocation returns the DefaultCustomerLocation field if non-nil, zero value otherwise.

### GetDefaultCustomerLocationOk

`func (o *IdentityAndPrivacyOptions) GetDefaultCustomerLocationOk() (*string, bool)`

GetDefaultCustomerLocationOk returns a tuple with the DefaultCustomerLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCustomerLocation

`func (o *IdentityAndPrivacyOptions) SetDefaultCustomerLocation(v string)`

SetDefaultCustomerLocation sets DefaultCustomerLocation field to given value.

### HasDefaultCustomerLocation

`func (o *IdentityAndPrivacyOptions) HasDefaultCustomerLocation() bool`

HasDefaultCustomerLocation returns a boolean if a field has been set.

### GetInactiveCartsRetentionPolicy

`func (o *IdentityAndPrivacyOptions) GetInactiveCartsRetentionPolicy() StoreDataRetentionPolicy`

GetInactiveCartsRetentionPolicy returns the InactiveCartsRetentionPolicy field if non-nil, zero value otherwise.

### GetInactiveCartsRetentionPolicyOk

`func (o *IdentityAndPrivacyOptions) GetInactiveCartsRetentionPolicyOk() (*StoreDataRetentionPolicy, bool)`

GetInactiveCartsRetentionPolicyOk returns a tuple with the InactiveCartsRetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveCartsRetentionPolicy

`func (o *IdentityAndPrivacyOptions) SetInactiveCartsRetentionPolicy(v StoreDataRetentionPolicy)`

SetInactiveCartsRetentionPolicy sets InactiveCartsRetentionPolicy field to given value.

### HasInactiveCartsRetentionPolicy

`func (o *IdentityAndPrivacyOptions) HasInactiveCartsRetentionPolicy() bool`

HasInactiveCartsRetentionPolicy returns a boolean if a field has been set.

### GetPendingOrdersRetentionPolicy

`func (o *IdentityAndPrivacyOptions) GetPendingOrdersRetentionPolicy() StoreDataRetentionPolicy`

GetPendingOrdersRetentionPolicy returns the PendingOrdersRetentionPolicy field if non-nil, zero value otherwise.

### GetPendingOrdersRetentionPolicyOk

`func (o *IdentityAndPrivacyOptions) GetPendingOrdersRetentionPolicyOk() (*StoreDataRetentionPolicy, bool)`

GetPendingOrdersRetentionPolicyOk returns a tuple with the PendingOrdersRetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingOrdersRetentionPolicy

`func (o *IdentityAndPrivacyOptions) SetPendingOrdersRetentionPolicy(v StoreDataRetentionPolicy)`

SetPendingOrdersRetentionPolicy sets PendingOrdersRetentionPolicy field to given value.

### HasPendingOrdersRetentionPolicy

`func (o *IdentityAndPrivacyOptions) HasPendingOrdersRetentionPolicy() bool`

HasPendingOrdersRetentionPolicy returns a boolean if a field has been set.

### GetFailedOrdersRetentionPolicy

`func (o *IdentityAndPrivacyOptions) GetFailedOrdersRetentionPolicy() StoreDataRetentionPolicy`

GetFailedOrdersRetentionPolicy returns the FailedOrdersRetentionPolicy field if non-nil, zero value otherwise.

### GetFailedOrdersRetentionPolicyOk

`func (o *IdentityAndPrivacyOptions) GetFailedOrdersRetentionPolicyOk() (*StoreDataRetentionPolicy, bool)`

GetFailedOrdersRetentionPolicyOk returns a tuple with the FailedOrdersRetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedOrdersRetentionPolicy

`func (o *IdentityAndPrivacyOptions) SetFailedOrdersRetentionPolicy(v StoreDataRetentionPolicy)`

SetFailedOrdersRetentionPolicy sets FailedOrdersRetentionPolicy field to given value.

### HasFailedOrdersRetentionPolicy

`func (o *IdentityAndPrivacyOptions) HasFailedOrdersRetentionPolicy() bool`

HasFailedOrdersRetentionPolicy returns a boolean if a field has been set.

### GetCancelledOrdersRetentionPolicy

`func (o *IdentityAndPrivacyOptions) GetCancelledOrdersRetentionPolicy() StoreDataRetentionPolicy`

GetCancelledOrdersRetentionPolicy returns the CancelledOrdersRetentionPolicy field if non-nil, zero value otherwise.

### GetCancelledOrdersRetentionPolicyOk

`func (o *IdentityAndPrivacyOptions) GetCancelledOrdersRetentionPolicyOk() (*StoreDataRetentionPolicy, bool)`

GetCancelledOrdersRetentionPolicyOk returns a tuple with the CancelledOrdersRetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledOrdersRetentionPolicy

`func (o *IdentityAndPrivacyOptions) SetCancelledOrdersRetentionPolicy(v StoreDataRetentionPolicy)`

SetCancelledOrdersRetentionPolicy sets CancelledOrdersRetentionPolicy field to given value.

### HasCancelledOrdersRetentionPolicy

`func (o *IdentityAndPrivacyOptions) HasCancelledOrdersRetentionPolicy() bool`

HasCancelledOrdersRetentionPolicy returns a boolean if a field has been set.

### GetCompletedOrdersRetentionPolicy

`func (o *IdentityAndPrivacyOptions) GetCompletedOrdersRetentionPolicy() StoreDataRetentionPolicy`

GetCompletedOrdersRetentionPolicy returns the CompletedOrdersRetentionPolicy field if non-nil, zero value otherwise.

### GetCompletedOrdersRetentionPolicyOk

`func (o *IdentityAndPrivacyOptions) GetCompletedOrdersRetentionPolicyOk() (*StoreDataRetentionPolicy, bool)`

GetCompletedOrdersRetentionPolicyOk returns a tuple with the CompletedOrdersRetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedOrdersRetentionPolicy

`func (o *IdentityAndPrivacyOptions) SetCompletedOrdersRetentionPolicy(v StoreDataRetentionPolicy)`

SetCompletedOrdersRetentionPolicy sets CompletedOrdersRetentionPolicy field to given value.

### HasCompletedOrdersRetentionPolicy

`func (o *IdentityAndPrivacyOptions) HasCompletedOrdersRetentionPolicy() bool`

HasCompletedOrdersRetentionPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


