# SubscriptionsOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DripDownloadableContent** | Pointer to **bool** |  | [optional] 
**RetryFailedPayments** | Pointer to **bool** |  | [optional] 
**Allow0InitialCheckout** | Pointer to **bool** |  | [optional] 
**AllowMixedCheckout** | Pointer to **bool** |  | [optional] 
**SynchroniseRenewals** | Pointer to **bool** |  | [optional] 
**AddToCartButtonText** | Pointer to **NullableString** |  | [optional] 
**PlaceOrderButtonText** | Pointer to **NullableString** |  | [optional] 
**NewSubscriberRole** | Pointer to **NullableString** |  | [optional] 
**InactiveSubscriberRole** | Pointer to **NullableString** |  | [optional] 
**EnableAutomaticPayments** | Pointer to **int32** |  | [optional] 
**EnableManualRenewals** | Pointer to **int32** |  | [optional] 
**DisplayAutoRenewalToggle** | Pointer to **int32** |  | [optional] 
**AcceptEarlyRenewals** | Pointer to **int32** |  | [optional] 
**CustomerSuspensions** | Pointer to **int32** |  | [optional] 
**EnableSubscriptionSwitchingBetweenGroups** | Pointer to **int32** |  | [optional] 
**EnableSubscriptionSwitchingBetweenVariations** | Pointer to **int32** |  | [optional] 
**ProrateFirstRenewal** | Pointer to **string** |  | [optional] 

## Methods

### NewSubscriptionsOptions

`func NewSubscriptionsOptions() *SubscriptionsOptions`

NewSubscriptionsOptions instantiates a new SubscriptionsOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionsOptionsWithDefaults

`func NewSubscriptionsOptionsWithDefaults() *SubscriptionsOptions`

NewSubscriptionsOptionsWithDefaults instantiates a new SubscriptionsOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDripDownloadableContent

`func (o *SubscriptionsOptions) GetDripDownloadableContent() bool`

GetDripDownloadableContent returns the DripDownloadableContent field if non-nil, zero value otherwise.

### GetDripDownloadableContentOk

`func (o *SubscriptionsOptions) GetDripDownloadableContentOk() (*bool, bool)`

GetDripDownloadableContentOk returns a tuple with the DripDownloadableContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDripDownloadableContent

`func (o *SubscriptionsOptions) SetDripDownloadableContent(v bool)`

SetDripDownloadableContent sets DripDownloadableContent field to given value.

### HasDripDownloadableContent

`func (o *SubscriptionsOptions) HasDripDownloadableContent() bool`

HasDripDownloadableContent returns a boolean if a field has been set.

### GetRetryFailedPayments

`func (o *SubscriptionsOptions) GetRetryFailedPayments() bool`

GetRetryFailedPayments returns the RetryFailedPayments field if non-nil, zero value otherwise.

### GetRetryFailedPaymentsOk

`func (o *SubscriptionsOptions) GetRetryFailedPaymentsOk() (*bool, bool)`

GetRetryFailedPaymentsOk returns a tuple with the RetryFailedPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryFailedPayments

`func (o *SubscriptionsOptions) SetRetryFailedPayments(v bool)`

SetRetryFailedPayments sets RetryFailedPayments field to given value.

### HasRetryFailedPayments

`func (o *SubscriptionsOptions) HasRetryFailedPayments() bool`

HasRetryFailedPayments returns a boolean if a field has been set.

### GetAllow0InitialCheckout

`func (o *SubscriptionsOptions) GetAllow0InitialCheckout() bool`

GetAllow0InitialCheckout returns the Allow0InitialCheckout field if non-nil, zero value otherwise.

### GetAllow0InitialCheckoutOk

`func (o *SubscriptionsOptions) GetAllow0InitialCheckoutOk() (*bool, bool)`

GetAllow0InitialCheckoutOk returns a tuple with the Allow0InitialCheckout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllow0InitialCheckout

`func (o *SubscriptionsOptions) SetAllow0InitialCheckout(v bool)`

SetAllow0InitialCheckout sets Allow0InitialCheckout field to given value.

### HasAllow0InitialCheckout

`func (o *SubscriptionsOptions) HasAllow0InitialCheckout() bool`

HasAllow0InitialCheckout returns a boolean if a field has been set.

### GetAllowMixedCheckout

`func (o *SubscriptionsOptions) GetAllowMixedCheckout() bool`

GetAllowMixedCheckout returns the AllowMixedCheckout field if non-nil, zero value otherwise.

### GetAllowMixedCheckoutOk

`func (o *SubscriptionsOptions) GetAllowMixedCheckoutOk() (*bool, bool)`

GetAllowMixedCheckoutOk returns a tuple with the AllowMixedCheckout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMixedCheckout

`func (o *SubscriptionsOptions) SetAllowMixedCheckout(v bool)`

SetAllowMixedCheckout sets AllowMixedCheckout field to given value.

### HasAllowMixedCheckout

`func (o *SubscriptionsOptions) HasAllowMixedCheckout() bool`

HasAllowMixedCheckout returns a boolean if a field has been set.

### GetSynchroniseRenewals

`func (o *SubscriptionsOptions) GetSynchroniseRenewals() bool`

GetSynchroniseRenewals returns the SynchroniseRenewals field if non-nil, zero value otherwise.

### GetSynchroniseRenewalsOk

`func (o *SubscriptionsOptions) GetSynchroniseRenewalsOk() (*bool, bool)`

GetSynchroniseRenewalsOk returns a tuple with the SynchroniseRenewals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchroniseRenewals

`func (o *SubscriptionsOptions) SetSynchroniseRenewals(v bool)`

SetSynchroniseRenewals sets SynchroniseRenewals field to given value.

### HasSynchroniseRenewals

`func (o *SubscriptionsOptions) HasSynchroniseRenewals() bool`

HasSynchroniseRenewals returns a boolean if a field has been set.

### GetAddToCartButtonText

`func (o *SubscriptionsOptions) GetAddToCartButtonText() string`

GetAddToCartButtonText returns the AddToCartButtonText field if non-nil, zero value otherwise.

### GetAddToCartButtonTextOk

`func (o *SubscriptionsOptions) GetAddToCartButtonTextOk() (*string, bool)`

GetAddToCartButtonTextOk returns a tuple with the AddToCartButtonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddToCartButtonText

`func (o *SubscriptionsOptions) SetAddToCartButtonText(v string)`

SetAddToCartButtonText sets AddToCartButtonText field to given value.

### HasAddToCartButtonText

`func (o *SubscriptionsOptions) HasAddToCartButtonText() bool`

HasAddToCartButtonText returns a boolean if a field has been set.

### SetAddToCartButtonTextNil

`func (o *SubscriptionsOptions) SetAddToCartButtonTextNil(b bool)`

 SetAddToCartButtonTextNil sets the value for AddToCartButtonText to be an explicit nil

### UnsetAddToCartButtonText
`func (o *SubscriptionsOptions) UnsetAddToCartButtonText()`

UnsetAddToCartButtonText ensures that no value is present for AddToCartButtonText, not even an explicit nil
### GetPlaceOrderButtonText

`func (o *SubscriptionsOptions) GetPlaceOrderButtonText() string`

GetPlaceOrderButtonText returns the PlaceOrderButtonText field if non-nil, zero value otherwise.

### GetPlaceOrderButtonTextOk

`func (o *SubscriptionsOptions) GetPlaceOrderButtonTextOk() (*string, bool)`

GetPlaceOrderButtonTextOk returns a tuple with the PlaceOrderButtonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOrderButtonText

`func (o *SubscriptionsOptions) SetPlaceOrderButtonText(v string)`

SetPlaceOrderButtonText sets PlaceOrderButtonText field to given value.

### HasPlaceOrderButtonText

`func (o *SubscriptionsOptions) HasPlaceOrderButtonText() bool`

HasPlaceOrderButtonText returns a boolean if a field has been set.

### SetPlaceOrderButtonTextNil

`func (o *SubscriptionsOptions) SetPlaceOrderButtonTextNil(b bool)`

 SetPlaceOrderButtonTextNil sets the value for PlaceOrderButtonText to be an explicit nil

### UnsetPlaceOrderButtonText
`func (o *SubscriptionsOptions) UnsetPlaceOrderButtonText()`

UnsetPlaceOrderButtonText ensures that no value is present for PlaceOrderButtonText, not even an explicit nil
### GetNewSubscriberRole

`func (o *SubscriptionsOptions) GetNewSubscriberRole() string`

GetNewSubscriberRole returns the NewSubscriberRole field if non-nil, zero value otherwise.

### GetNewSubscriberRoleOk

`func (o *SubscriptionsOptions) GetNewSubscriberRoleOk() (*string, bool)`

GetNewSubscriberRoleOk returns a tuple with the NewSubscriberRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSubscriberRole

`func (o *SubscriptionsOptions) SetNewSubscriberRole(v string)`

SetNewSubscriberRole sets NewSubscriberRole field to given value.

### HasNewSubscriberRole

`func (o *SubscriptionsOptions) HasNewSubscriberRole() bool`

HasNewSubscriberRole returns a boolean if a field has been set.

### SetNewSubscriberRoleNil

`func (o *SubscriptionsOptions) SetNewSubscriberRoleNil(b bool)`

 SetNewSubscriberRoleNil sets the value for NewSubscriberRole to be an explicit nil

### UnsetNewSubscriberRole
`func (o *SubscriptionsOptions) UnsetNewSubscriberRole()`

UnsetNewSubscriberRole ensures that no value is present for NewSubscriberRole, not even an explicit nil
### GetInactiveSubscriberRole

`func (o *SubscriptionsOptions) GetInactiveSubscriberRole() string`

GetInactiveSubscriberRole returns the InactiveSubscriberRole field if non-nil, zero value otherwise.

### GetInactiveSubscriberRoleOk

`func (o *SubscriptionsOptions) GetInactiveSubscriberRoleOk() (*string, bool)`

GetInactiveSubscriberRoleOk returns a tuple with the InactiveSubscriberRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveSubscriberRole

`func (o *SubscriptionsOptions) SetInactiveSubscriberRole(v string)`

SetInactiveSubscriberRole sets InactiveSubscriberRole field to given value.

### HasInactiveSubscriberRole

`func (o *SubscriptionsOptions) HasInactiveSubscriberRole() bool`

HasInactiveSubscriberRole returns a boolean if a field has been set.

### SetInactiveSubscriberRoleNil

`func (o *SubscriptionsOptions) SetInactiveSubscriberRoleNil(b bool)`

 SetInactiveSubscriberRoleNil sets the value for InactiveSubscriberRole to be an explicit nil

### UnsetInactiveSubscriberRole
`func (o *SubscriptionsOptions) UnsetInactiveSubscriberRole()`

UnsetInactiveSubscriberRole ensures that no value is present for InactiveSubscriberRole, not even an explicit nil
### GetEnableAutomaticPayments

`func (o *SubscriptionsOptions) GetEnableAutomaticPayments() int32`

GetEnableAutomaticPayments returns the EnableAutomaticPayments field if non-nil, zero value otherwise.

### GetEnableAutomaticPaymentsOk

`func (o *SubscriptionsOptions) GetEnableAutomaticPaymentsOk() (*int32, bool)`

GetEnableAutomaticPaymentsOk returns a tuple with the EnableAutomaticPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticPayments

`func (o *SubscriptionsOptions) SetEnableAutomaticPayments(v int32)`

SetEnableAutomaticPayments sets EnableAutomaticPayments field to given value.

### HasEnableAutomaticPayments

`func (o *SubscriptionsOptions) HasEnableAutomaticPayments() bool`

HasEnableAutomaticPayments returns a boolean if a field has been set.

### GetEnableManualRenewals

`func (o *SubscriptionsOptions) GetEnableManualRenewals() int32`

GetEnableManualRenewals returns the EnableManualRenewals field if non-nil, zero value otherwise.

### GetEnableManualRenewalsOk

`func (o *SubscriptionsOptions) GetEnableManualRenewalsOk() (*int32, bool)`

GetEnableManualRenewalsOk returns a tuple with the EnableManualRenewals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableManualRenewals

`func (o *SubscriptionsOptions) SetEnableManualRenewals(v int32)`

SetEnableManualRenewals sets EnableManualRenewals field to given value.

### HasEnableManualRenewals

`func (o *SubscriptionsOptions) HasEnableManualRenewals() bool`

HasEnableManualRenewals returns a boolean if a field has been set.

### GetDisplayAutoRenewalToggle

`func (o *SubscriptionsOptions) GetDisplayAutoRenewalToggle() int32`

GetDisplayAutoRenewalToggle returns the DisplayAutoRenewalToggle field if non-nil, zero value otherwise.

### GetDisplayAutoRenewalToggleOk

`func (o *SubscriptionsOptions) GetDisplayAutoRenewalToggleOk() (*int32, bool)`

GetDisplayAutoRenewalToggleOk returns a tuple with the DisplayAutoRenewalToggle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayAutoRenewalToggle

`func (o *SubscriptionsOptions) SetDisplayAutoRenewalToggle(v int32)`

SetDisplayAutoRenewalToggle sets DisplayAutoRenewalToggle field to given value.

### HasDisplayAutoRenewalToggle

`func (o *SubscriptionsOptions) HasDisplayAutoRenewalToggle() bool`

HasDisplayAutoRenewalToggle returns a boolean if a field has been set.

### GetAcceptEarlyRenewals

`func (o *SubscriptionsOptions) GetAcceptEarlyRenewals() int32`

GetAcceptEarlyRenewals returns the AcceptEarlyRenewals field if non-nil, zero value otherwise.

### GetAcceptEarlyRenewalsOk

`func (o *SubscriptionsOptions) GetAcceptEarlyRenewalsOk() (*int32, bool)`

GetAcceptEarlyRenewalsOk returns a tuple with the AcceptEarlyRenewals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptEarlyRenewals

`func (o *SubscriptionsOptions) SetAcceptEarlyRenewals(v int32)`

SetAcceptEarlyRenewals sets AcceptEarlyRenewals field to given value.

### HasAcceptEarlyRenewals

`func (o *SubscriptionsOptions) HasAcceptEarlyRenewals() bool`

HasAcceptEarlyRenewals returns a boolean if a field has been set.

### GetCustomerSuspensions

`func (o *SubscriptionsOptions) GetCustomerSuspensions() int32`

GetCustomerSuspensions returns the CustomerSuspensions field if non-nil, zero value otherwise.

### GetCustomerSuspensionsOk

`func (o *SubscriptionsOptions) GetCustomerSuspensionsOk() (*int32, bool)`

GetCustomerSuspensionsOk returns a tuple with the CustomerSuspensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerSuspensions

`func (o *SubscriptionsOptions) SetCustomerSuspensions(v int32)`

SetCustomerSuspensions sets CustomerSuspensions field to given value.

### HasCustomerSuspensions

`func (o *SubscriptionsOptions) HasCustomerSuspensions() bool`

HasCustomerSuspensions returns a boolean if a field has been set.

### GetEnableSubscriptionSwitchingBetweenGroups

`func (o *SubscriptionsOptions) GetEnableSubscriptionSwitchingBetweenGroups() int32`

GetEnableSubscriptionSwitchingBetweenGroups returns the EnableSubscriptionSwitchingBetweenGroups field if non-nil, zero value otherwise.

### GetEnableSubscriptionSwitchingBetweenGroupsOk

`func (o *SubscriptionsOptions) GetEnableSubscriptionSwitchingBetweenGroupsOk() (*int32, bool)`

GetEnableSubscriptionSwitchingBetweenGroupsOk returns a tuple with the EnableSubscriptionSwitchingBetweenGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSubscriptionSwitchingBetweenGroups

`func (o *SubscriptionsOptions) SetEnableSubscriptionSwitchingBetweenGroups(v int32)`

SetEnableSubscriptionSwitchingBetweenGroups sets EnableSubscriptionSwitchingBetweenGroups field to given value.

### HasEnableSubscriptionSwitchingBetweenGroups

`func (o *SubscriptionsOptions) HasEnableSubscriptionSwitchingBetweenGroups() bool`

HasEnableSubscriptionSwitchingBetweenGroups returns a boolean if a field has been set.

### GetEnableSubscriptionSwitchingBetweenVariations

`func (o *SubscriptionsOptions) GetEnableSubscriptionSwitchingBetweenVariations() int32`

GetEnableSubscriptionSwitchingBetweenVariations returns the EnableSubscriptionSwitchingBetweenVariations field if non-nil, zero value otherwise.

### GetEnableSubscriptionSwitchingBetweenVariationsOk

`func (o *SubscriptionsOptions) GetEnableSubscriptionSwitchingBetweenVariationsOk() (*int32, bool)`

GetEnableSubscriptionSwitchingBetweenVariationsOk returns a tuple with the EnableSubscriptionSwitchingBetweenVariations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSubscriptionSwitchingBetweenVariations

`func (o *SubscriptionsOptions) SetEnableSubscriptionSwitchingBetweenVariations(v int32)`

SetEnableSubscriptionSwitchingBetweenVariations sets EnableSubscriptionSwitchingBetweenVariations field to given value.

### HasEnableSubscriptionSwitchingBetweenVariations

`func (o *SubscriptionsOptions) HasEnableSubscriptionSwitchingBetweenVariations() bool`

HasEnableSubscriptionSwitchingBetweenVariations returns a boolean if a field has been set.

### GetProrateFirstRenewal

`func (o *SubscriptionsOptions) GetProrateFirstRenewal() string`

GetProrateFirstRenewal returns the ProrateFirstRenewal field if non-nil, zero value otherwise.

### GetProrateFirstRenewalOk

`func (o *SubscriptionsOptions) GetProrateFirstRenewalOk() (*string, bool)`

GetProrateFirstRenewalOk returns a tuple with the ProrateFirstRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrateFirstRenewal

`func (o *SubscriptionsOptions) SetProrateFirstRenewal(v string)`

SetProrateFirstRenewal sets ProrateFirstRenewal field to given value.

### HasProrateFirstRenewal

`func (o *SubscriptionsOptions) HasProrateFirstRenewal() bool`

HasProrateFirstRenewal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


