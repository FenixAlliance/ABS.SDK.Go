# EmailOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromName** | Pointer to **NullableString** |  | [optional] 
**FromEmailAddress** | Pointer to **NullableString** |  | [optional] 
**HeaderImage** | Pointer to **NullableString** |  | [optional] 
**FooterText** | Pointer to **NullableString** |  | [optional] 
**BaseColor** | Pointer to **NullableString** |  | [optional] 
**BackgroundColor** | Pointer to **NullableString** |  | [optional] 
**BodyBackgroundColor** | Pointer to **NullableString** |  | [optional] 
**BodyTextColor** | Pointer to **NullableString** |  | [optional] 
**NewOrderEmailID** | Pointer to **NullableString** |  | [optional] 
**CancelledOrderEmailID** | Pointer to **NullableString** |  | [optional] 
**FailedOrderEmailID** | Pointer to **NullableString** |  | [optional] 
**OnHoldOrderEmailID** | Pointer to **NullableString** |  | [optional] 
**ProcessingOrderEmailID** | Pointer to **NullableString** |  | [optional] 
**CompletedOrderEmailID** | Pointer to **NullableString** |  | [optional] 
**RefundedOrderEmailID** | Pointer to **NullableString** |  | [optional] 
**OrderDetailsEmailID** | Pointer to **NullableString** |  | [optional] 
**CustomerInvoiceEmailID** | Pointer to **NullableString** |  | [optional] 
**CustomerNoteEmailID** | Pointer to **NullableString** |  | [optional] 
**PasswordResetEmailID** | Pointer to **NullableString** |  | [optional] 
**NewRenewalOrderEmailID** | Pointer to **NullableString** |  | [optional] 
**NewSubscriptionEmailID** | Pointer to **NullableString** |  | [optional] 
**SubscriptionWelcomeEmailID** | Pointer to **NullableString** |  | [optional] 
**SuspendedSubscriptionEmailID** | Pointer to **NullableString** |  | [optional] 
**OverdueSubscriptionEmailID** | Pointer to **NullableString** |  | [optional] 
**ExpiredSubscriptionEmailID** | Pointer to **NullableString** |  | [optional] 
**SwitchedSubscriptionEmailID** | Pointer to **NullableString** |  | [optional] 
**NewAccountEmailID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEmailOptions

`func NewEmailOptions() *EmailOptions`

NewEmailOptions instantiates a new EmailOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailOptionsWithDefaults

`func NewEmailOptionsWithDefaults() *EmailOptions`

NewEmailOptionsWithDefaults instantiates a new EmailOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromName

`func (o *EmailOptions) GetFromName() string`

GetFromName returns the FromName field if non-nil, zero value otherwise.

### GetFromNameOk

`func (o *EmailOptions) GetFromNameOk() (*string, bool)`

GetFromNameOk returns a tuple with the FromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromName

`func (o *EmailOptions) SetFromName(v string)`

SetFromName sets FromName field to given value.

### HasFromName

`func (o *EmailOptions) HasFromName() bool`

HasFromName returns a boolean if a field has been set.

### SetFromNameNil

`func (o *EmailOptions) SetFromNameNil(b bool)`

 SetFromNameNil sets the value for FromName to be an explicit nil

### UnsetFromName
`func (o *EmailOptions) UnsetFromName()`

UnsetFromName ensures that no value is present for FromName, not even an explicit nil
### GetFromEmailAddress

`func (o *EmailOptions) GetFromEmailAddress() string`

GetFromEmailAddress returns the FromEmailAddress field if non-nil, zero value otherwise.

### GetFromEmailAddressOk

`func (o *EmailOptions) GetFromEmailAddressOk() (*string, bool)`

GetFromEmailAddressOk returns a tuple with the FromEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmailAddress

`func (o *EmailOptions) SetFromEmailAddress(v string)`

SetFromEmailAddress sets FromEmailAddress field to given value.

### HasFromEmailAddress

`func (o *EmailOptions) HasFromEmailAddress() bool`

HasFromEmailAddress returns a boolean if a field has been set.

### SetFromEmailAddressNil

`func (o *EmailOptions) SetFromEmailAddressNil(b bool)`

 SetFromEmailAddressNil sets the value for FromEmailAddress to be an explicit nil

### UnsetFromEmailAddress
`func (o *EmailOptions) UnsetFromEmailAddress()`

UnsetFromEmailAddress ensures that no value is present for FromEmailAddress, not even an explicit nil
### GetHeaderImage

`func (o *EmailOptions) GetHeaderImage() string`

GetHeaderImage returns the HeaderImage field if non-nil, zero value otherwise.

### GetHeaderImageOk

`func (o *EmailOptions) GetHeaderImageOk() (*string, bool)`

GetHeaderImageOk returns a tuple with the HeaderImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderImage

`func (o *EmailOptions) SetHeaderImage(v string)`

SetHeaderImage sets HeaderImage field to given value.

### HasHeaderImage

`func (o *EmailOptions) HasHeaderImage() bool`

HasHeaderImage returns a boolean if a field has been set.

### SetHeaderImageNil

`func (o *EmailOptions) SetHeaderImageNil(b bool)`

 SetHeaderImageNil sets the value for HeaderImage to be an explicit nil

### UnsetHeaderImage
`func (o *EmailOptions) UnsetHeaderImage()`

UnsetHeaderImage ensures that no value is present for HeaderImage, not even an explicit nil
### GetFooterText

`func (o *EmailOptions) GetFooterText() string`

GetFooterText returns the FooterText field if non-nil, zero value otherwise.

### GetFooterTextOk

`func (o *EmailOptions) GetFooterTextOk() (*string, bool)`

GetFooterTextOk returns a tuple with the FooterText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterText

`func (o *EmailOptions) SetFooterText(v string)`

SetFooterText sets FooterText field to given value.

### HasFooterText

`func (o *EmailOptions) HasFooterText() bool`

HasFooterText returns a boolean if a field has been set.

### SetFooterTextNil

`func (o *EmailOptions) SetFooterTextNil(b bool)`

 SetFooterTextNil sets the value for FooterText to be an explicit nil

### UnsetFooterText
`func (o *EmailOptions) UnsetFooterText()`

UnsetFooterText ensures that no value is present for FooterText, not even an explicit nil
### GetBaseColor

`func (o *EmailOptions) GetBaseColor() string`

GetBaseColor returns the BaseColor field if non-nil, zero value otherwise.

### GetBaseColorOk

`func (o *EmailOptions) GetBaseColorOk() (*string, bool)`

GetBaseColorOk returns a tuple with the BaseColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseColor

`func (o *EmailOptions) SetBaseColor(v string)`

SetBaseColor sets BaseColor field to given value.

### HasBaseColor

`func (o *EmailOptions) HasBaseColor() bool`

HasBaseColor returns a boolean if a field has been set.

### SetBaseColorNil

`func (o *EmailOptions) SetBaseColorNil(b bool)`

 SetBaseColorNil sets the value for BaseColor to be an explicit nil

### UnsetBaseColor
`func (o *EmailOptions) UnsetBaseColor()`

UnsetBaseColor ensures that no value is present for BaseColor, not even an explicit nil
### GetBackgroundColor

`func (o *EmailOptions) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *EmailOptions) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *EmailOptions) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *EmailOptions) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### SetBackgroundColorNil

`func (o *EmailOptions) SetBackgroundColorNil(b bool)`

 SetBackgroundColorNil sets the value for BackgroundColor to be an explicit nil

### UnsetBackgroundColor
`func (o *EmailOptions) UnsetBackgroundColor()`

UnsetBackgroundColor ensures that no value is present for BackgroundColor, not even an explicit nil
### GetBodyBackgroundColor

`func (o *EmailOptions) GetBodyBackgroundColor() string`

GetBodyBackgroundColor returns the BodyBackgroundColor field if non-nil, zero value otherwise.

### GetBodyBackgroundColorOk

`func (o *EmailOptions) GetBodyBackgroundColorOk() (*string, bool)`

GetBodyBackgroundColorOk returns a tuple with the BodyBackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyBackgroundColor

`func (o *EmailOptions) SetBodyBackgroundColor(v string)`

SetBodyBackgroundColor sets BodyBackgroundColor field to given value.

### HasBodyBackgroundColor

`func (o *EmailOptions) HasBodyBackgroundColor() bool`

HasBodyBackgroundColor returns a boolean if a field has been set.

### SetBodyBackgroundColorNil

`func (o *EmailOptions) SetBodyBackgroundColorNil(b bool)`

 SetBodyBackgroundColorNil sets the value for BodyBackgroundColor to be an explicit nil

### UnsetBodyBackgroundColor
`func (o *EmailOptions) UnsetBodyBackgroundColor()`

UnsetBodyBackgroundColor ensures that no value is present for BodyBackgroundColor, not even an explicit nil
### GetBodyTextColor

`func (o *EmailOptions) GetBodyTextColor() string`

GetBodyTextColor returns the BodyTextColor field if non-nil, zero value otherwise.

### GetBodyTextColorOk

`func (o *EmailOptions) GetBodyTextColorOk() (*string, bool)`

GetBodyTextColorOk returns a tuple with the BodyTextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyTextColor

`func (o *EmailOptions) SetBodyTextColor(v string)`

SetBodyTextColor sets BodyTextColor field to given value.

### HasBodyTextColor

`func (o *EmailOptions) HasBodyTextColor() bool`

HasBodyTextColor returns a boolean if a field has been set.

### SetBodyTextColorNil

`func (o *EmailOptions) SetBodyTextColorNil(b bool)`

 SetBodyTextColorNil sets the value for BodyTextColor to be an explicit nil

### UnsetBodyTextColor
`func (o *EmailOptions) UnsetBodyTextColor()`

UnsetBodyTextColor ensures that no value is present for BodyTextColor, not even an explicit nil
### GetNewOrderEmailID

`func (o *EmailOptions) GetNewOrderEmailID() string`

GetNewOrderEmailID returns the NewOrderEmailID field if non-nil, zero value otherwise.

### GetNewOrderEmailIDOk

`func (o *EmailOptions) GetNewOrderEmailIDOk() (*string, bool)`

GetNewOrderEmailIDOk returns a tuple with the NewOrderEmailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOrderEmailID

`func (o *EmailOptions) SetNewOrderEmailID(v string)`

SetNewOrderEmailID sets NewOrderEmailID field to given value.

### HasNewOrderEmailID

`func (o *EmailOptions) HasNewOrderEmailID() bool`

HasNewOrderEmailID returns a boolean if a field has been set.

### SetNewOrderEmailIDNil

`func (o *EmailOptions) SetNewOrderEmailIDNil(b bool)`

 SetNewOrderEmailIDNil sets the value for NewOrderEmailID to be an explicit nil

### UnsetNewOrderEmailID
`func (o *EmailOptions) UnsetNewOrderEmailID()`

UnsetNewOrderEmailID ensures that no value is present for NewOrderEmailID, not even an explicit nil
### GetCancelledOrderEmailID

`func (o *EmailOptions) GetCancelledOrderEmailID() string`

GetCancelledOrderEmailID returns the CancelledOrderEmailID field if non-nil, zero value otherwise.

### GetCancelledOrderEmailIDOk

`func (o *EmailOptions) GetCancelledOrderEmailIDOk() (*string, bool)`

GetCancelledOrderEmailIDOk returns a tuple with the CancelledOrderEmailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledOrderEmailID

`func (o *EmailOptions) SetCancelledOrderEmailID(v string)`

SetCancelledOrderEmailID sets CancelledOrderEmailID field to given value.

### HasCancelledOrderEmailID

`func (o *EmailOptions) HasCancelledOrderEmailID() bool`

HasCancelledOrderEmailID returns a boolean if a field has been set.

### SetCancelledOrderEmailIDNil

`func (o *EmailOptions) SetCancelledOrderEmailIDNil(b bool)`

 SetCancelledOrderEmailIDNil sets the value for CancelledOrderEmailID to be an explicit nil

### UnsetCancelledOrderEmailID
`func (o *EmailOptions) UnsetCancelledOrderEmailID()`

UnsetCancelledOrderEmailID ensures that no value is present for CancelledOrderEmailID, not even an explicit nil
### GetFailedOrderEmailID

`func (o *EmailOptions) GetFailedOrderEmailID() string`

GetFailedOrderEmailID returns the FailedOrderEmailID field if non-nil, zero value otherwise.

### GetFailedOrderEmailIDOk

`func (o *EmailOptions) GetFailedOrderEmailIDOk() (*string, bool)`

GetFailedOrderEmailIDOk returns a tuple with the FailedOrderEmailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedOrderEmailID

`func (o *EmailOptions) SetFailedOrderEmailID(v string)`

SetFailedOrderEmailID sets FailedOrderEmailID field to given value.

### HasFailedOrderEmailID

`func (o *EmailOptions) HasFailedOrderEmailID() bool`

HasFailedOrderEmailID returns a boolean if a field has been set.

### SetFailedOrderEmailIDNil

`func (o *EmailOptions) SetFailedOrderEmailIDNil(b bool)`

 SetFailedOrderEmailIDNil sets the value for FailedOrderEmailID to be an explicit nil

### UnsetFailedOrderEmailID
`func (o *EmailOptions) UnsetFailedOrderEmailID()`

UnsetFailedOrderEmailID ensures that no value is present for FailedOrderEmailID, not even an explicit nil
### GetOnHoldOrderEmailID

`func (o *EmailOptions) GetOnHoldOrderEmailID() string`

GetOnHoldOrderEmailID returns the OnHoldOrderEmailID field if non-nil, zero value otherwise.

### GetOnHoldOrderEmailIDOk

`func (o *EmailOptions) GetOnHoldOrderEmailIDOk() (*string, bool)`

GetOnHoldOrderEmailIDOk returns a tuple with the OnHoldOrderEmailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnHoldOrderEmailID

`func (o *EmailOptions) SetOnHoldOrderEmailID(v string)`

SetOnHoldOrderEmailID sets OnHoldOrderEmailID field to given value.

### HasOnHoldOrderEmailID

`func (o *EmailOptions) HasOnHoldOrderEmailID() bool`

HasOnHoldOrderEmailID returns a boolean if a field has been set.

### SetOnHoldOrderEmailIDNil

`func (o *EmailOptions) SetOnHoldOrderEmailIDNil(b bool)`

 SetOnHoldOrderEmailIDNil sets the value for OnHoldOrderEmailID to be an explicit nil

### UnsetOnHoldOrderEmailID
`func (o *EmailOptions) UnsetOnHoldOrderEmailID()`

UnsetOnHoldOrderEmailID ensures that no value is present for OnHoldOrderEmailID, not even an explicit nil
### GetProcessingOrderEmailID

`func (o *EmailOptions) GetProcessingOrderEmailID() string`

GetProcessingOrderEmailID returns the ProcessingOrderEmailID field if non-nil, zero value otherwise.

### GetProcessingOrderEmailIDOk

`func (o *EmailOptions) GetProcessingOrderEmailIDOk() (*string, bool)`

GetProcessingOrderEmailIDOk returns a tuple with the ProcessingOrderEmailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingOrderEmailID

`func (o *EmailOptions) SetProcessingOrderEmailID(v string)`

SetProcessingOrderEmailID sets ProcessingOrderEmailID field to given value.

### HasProcessingOrderEmailID

`func (o *EmailOptions) HasProcessingOrderEmailID() bool`

HasProcessingOrderEmailID returns a boolean if a field has been set.

### SetProcessingOrderEmailIDNil

`func (o *EmailOptions) SetProcessingOrderEmailIDNil(b bool)`

 SetProcessingOrderEmailIDNil sets the value for ProcessingOrderEmailID to be an explicit nil

### UnsetProcessingOrderEmailID
`func (o *EmailOptions) UnsetProcessingOrderEmailID()`

UnsetProcessingOrderEmailID ensures that no value is present for ProcessingOrderEmailID, not even an explicit nil
### GetCompletedOrderEmailID

`func (o *EmailOptions) GetCompletedOrderEmailID() string`

GetCompletedOrderEmailID returns the CompletedOrderEmailID field if non-nil, zero value otherwise.

### GetCompletedOrderEmailIDOk

`func (o *EmailOptions) GetCompletedOrderEmailIDOk() (*string, bool)`

GetCompletedOrderEmailIDOk returns a tuple with the CompletedOrderEmailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedOrderEmailID

`func (o *EmailOptions) SetCompletedOrderEmailID(v string)`

SetCompletedOrderEmailID sets CompletedOrderEmailID field to given value.

### HasCompletedOrderEmailID

`func (o *EmailOptions) HasCompletedOrderEmailID() bool`

HasCompletedOrderEmailID returns a boolean if a field has been set.

### SetCompletedOrderEmailIDNil

`func (o *EmailOptions) SetCompletedOrderEmailIDNil(b bool)`

 SetCompletedOrderEmailIDNil sets the value for CompletedOrderEmailID to be an explicit nil

### UnsetCompletedOrderEmailID
`func (o *EmailOptions) UnsetCompletedOrderEmailID()`

UnsetCompletedOrderEmailID ensures that no value is present for CompletedOrderEmailID, not even an explicit nil
### GetRefundedOrderEmailID

`func (o *EmailOptions) GetRefundedOrderEmailID() string`

GetRefundedOrderEmailID returns the RefundedOrderEmailID field if non-nil, zero value otherwise.

### GetRefundedOrderEmailIDOk

`func (o *EmailOptions) GetRefundedOrderEmailIDOk() (*string, bool)`

GetRefundedOrderEmailIDOk returns a tuple with the RefundedOrderEmailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedOrderEmailID

`func (o *EmailOptions) SetRefundedOrderEmailID(v string)`

SetRefundedOrderEmailID sets RefundedOrderEmailID field to given value.

### HasRefundedOrderEmailID

`func (o *EmailOptions) HasRefundedOrderEmailID() bool`

HasRefundedOrderEmailID returns a boolean if a field has been set.

### SetRefundedOrderEmailIDNil

`func (o *EmailOptions) SetRefundedOrderEmailIDNil(b bool)`

 SetRefundedOrderEmailIDNil sets the value for RefundedOrderEmailID to be an explicit nil

### UnsetRefundedOrderEmailID
`func (o *EmailOptions) UnsetRefundedOrderEmailID()`

UnsetRefundedOrderEmailID ensures that no value is present for RefundedOrderEmailID, not even an explicit nil
### GetOrderDetailsEmailID

`func (o *EmailOptions) GetOrderDetailsEmailID() string`

GetOrderDetailsEmailID returns the OrderDetailsEmailID field if non-nil, zero value otherwise.

### GetOrderDetailsEmailIDOk

`func (o *EmailOptions) GetOrderDetailsEmailIDOk() (*string, bool)`

GetOrderDetailsEmailIDOk returns a tuple with the OrderDetailsEmailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDetailsEmailID

`func (o *EmailOptions) SetOrderDetailsEmailID(v string)`

SetOrderDetailsEmailID sets OrderDetailsEmailID field to given value.

### HasOrderDetailsEmailID

`func (o *EmailOptions) HasOrderDetailsEmailID() bool`

HasOrderDetailsEmailID returns a boolean if a field has been set.

### SetOrderDetailsEmailIDNil

`func (o *EmailOptions) SetOrderDetailsEmailIDNil(b bool)`

 SetOrderDetailsEmailIDNil sets the value for OrderDetailsEmailID to be an explicit nil

### UnsetOrderDetailsEmailID
`func (o *EmailOptions) UnsetOrderDetailsEmailID()`

UnsetOrderDetailsEmailID ensures that no value is present for OrderDetailsEmailID, not even an explicit nil
### GetCustomerInvoiceEmailID

`func (o *EmailOptions) GetCustomerInvoiceEmailID() string`

GetCustomerInvoiceEmailID returns the CustomerInvoiceEmailID field if non-nil, zero value otherwise.

### GetCustomerInvoiceEmailIDOk

`func (o *EmailOptions) GetCustomerInvoiceEmailIDOk() (*string, bool)`

GetCustomerInvoiceEmailIDOk returns a tuple with the CustomerInvoiceEmailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerInvoiceEmailID

`func (o *EmailOptions) SetCustomerInvoiceEmailID(v string)`

SetCustomerInvoiceEmailID sets CustomerInvoiceEmailID field to given value.

### HasCustomerInvoiceEmailID

`func (o *EmailOptions) HasCustomerInvoiceEmailID() bool`

HasCustomerInvoiceEmailID returns a boolean if a field has been set.

### SetCustomerInvoiceEmailIDNil

`func (o *EmailOptions) SetCustomerInvoiceEmailIDNil(b bool)`

 SetCustomerInvoiceEmailIDNil sets the value for CustomerInvoiceEmailID to be an explicit nil

### UnsetCustomerInvoiceEmailID
`func (o *EmailOptions) UnsetCustomerInvoiceEmailID()`

UnsetCustomerInvoiceEmailID ensures that no value is present for CustomerInvoiceEmailID, not even an explicit nil
### GetCustomerNoteEmailID

`func (o *EmailOptions) GetCustomerNoteEmailID() string`

GetCustomerNoteEmailID returns the CustomerNoteEmailID field if non-nil, zero value otherwise.

### GetCustomerNoteEmailIDOk

`func (o *EmailOptions) GetCustomerNoteEmailIDOk() (*string, bool)`

GetCustomerNoteEmailIDOk returns a tuple with the CustomerNoteEmailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNoteEmailID

`func (o *EmailOptions) SetCustomerNoteEmailID(v string)`

SetCustomerNoteEmailID sets CustomerNoteEmailID field to given value.

### HasCustomerNoteEmailID

`func (o *EmailOptions) HasCustomerNoteEmailID() bool`

HasCustomerNoteEmailID returns a boolean if a field has been set.

### SetCustomerNoteEmailIDNil

`func (o *EmailOptions) SetCustomerNoteEmailIDNil(b bool)`

 SetCustomerNoteEmailIDNil sets the value for CustomerNoteEmailID to be an explicit nil

### UnsetCustomerNoteEmailID
`func (o *EmailOptions) UnsetCustomerNoteEmailID()`

UnsetCustomerNoteEmailID ensures that no value is present for CustomerNoteEmailID, not even an explicit nil
### GetPasswordResetEmailID

`func (o *EmailOptions) GetPasswordResetEmailID() string`

GetPasswordResetEmailID returns the PasswordResetEmailID field if non-nil, zero value otherwise.

### GetPasswordResetEmailIDOk

`func (o *EmailOptions) GetPasswordResetEmailIDOk() (*string, bool)`

GetPasswordResetEmailIDOk returns a tuple with the PasswordResetEmailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordResetEmailID

`func (o *EmailOptions) SetPasswordResetEmailID(v string)`

SetPasswordResetEmailID sets PasswordResetEmailID field to given value.

### HasPasswordResetEmailID

`func (o *EmailOptions) HasPasswordResetEmailID() bool`

HasPasswordResetEmailID returns a boolean if a field has been set.

### SetPasswordResetEmailIDNil

`func (o *EmailOptions) SetPasswordResetEmailIDNil(b bool)`

 SetPasswordResetEmailIDNil sets the value for PasswordResetEmailID to be an explicit nil

### UnsetPasswordResetEmailID
`func (o *EmailOptions) UnsetPasswordResetEmailID()`

UnsetPasswordResetEmailID ensures that no value is present for PasswordResetEmailID, not even an explicit nil
### GetNewRenewalOrderEmailID

`func (o *EmailOptions) GetNewRenewalOrderEmailID() string`

GetNewRenewalOrderEmailID returns the NewRenewalOrderEmailID field if non-nil, zero value otherwise.

### GetNewRenewalOrderEmailIDOk

`func (o *EmailOptions) GetNewRenewalOrderEmailIDOk() (*string, bool)`

GetNewRenewalOrderEmailIDOk returns a tuple with the NewRenewalOrderEmailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRenewalOrderEmailID

`func (o *EmailOptions) SetNewRenewalOrderEmailID(v string)`

SetNewRenewalOrderEmailID sets NewRenewalOrderEmailID field to given value.

### HasNewRenewalOrderEmailID

`func (o *EmailOptions) HasNewRenewalOrderEmailID() bool`

HasNewRenewalOrderEmailID returns a boolean if a field has been set.

### SetNewRenewalOrderEmailIDNil

`func (o *EmailOptions) SetNewRenewalOrderEmailIDNil(b bool)`

 SetNewRenewalOrderEmailIDNil sets the value for NewRenewalOrderEmailID to be an explicit nil

### UnsetNewRenewalOrderEmailID
`func (o *EmailOptions) UnsetNewRenewalOrderEmailID()`

UnsetNewRenewalOrderEmailID ensures that no value is present for NewRenewalOrderEmailID, not even an explicit nil
### GetNewSubscriptionEmailID

`func (o *EmailOptions) GetNewSubscriptionEmailID() string`

GetNewSubscriptionEmailID returns the NewSubscriptionEmailID field if non-nil, zero value otherwise.

### GetNewSubscriptionEmailIDOk

`func (o *EmailOptions) GetNewSubscriptionEmailIDOk() (*string, bool)`

GetNewSubscriptionEmailIDOk returns a tuple with the NewSubscriptionEmailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSubscriptionEmailID

`func (o *EmailOptions) SetNewSubscriptionEmailID(v string)`

SetNewSubscriptionEmailID sets NewSubscriptionEmailID field to given value.

### HasNewSubscriptionEmailID

`func (o *EmailOptions) HasNewSubscriptionEmailID() bool`

HasNewSubscriptionEmailID returns a boolean if a field has been set.

### SetNewSubscriptionEmailIDNil

`func (o *EmailOptions) SetNewSubscriptionEmailIDNil(b bool)`

 SetNewSubscriptionEmailIDNil sets the value for NewSubscriptionEmailID to be an explicit nil

### UnsetNewSubscriptionEmailID
`func (o *EmailOptions) UnsetNewSubscriptionEmailID()`

UnsetNewSubscriptionEmailID ensures that no value is present for NewSubscriptionEmailID, not even an explicit nil
### GetSubscriptionWelcomeEmailID

`func (o *EmailOptions) GetSubscriptionWelcomeEmailID() string`

GetSubscriptionWelcomeEmailID returns the SubscriptionWelcomeEmailID field if non-nil, zero value otherwise.

### GetSubscriptionWelcomeEmailIDOk

`func (o *EmailOptions) GetSubscriptionWelcomeEmailIDOk() (*string, bool)`

GetSubscriptionWelcomeEmailIDOk returns a tuple with the SubscriptionWelcomeEmailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionWelcomeEmailID

`func (o *EmailOptions) SetSubscriptionWelcomeEmailID(v string)`

SetSubscriptionWelcomeEmailID sets SubscriptionWelcomeEmailID field to given value.

### HasSubscriptionWelcomeEmailID

`func (o *EmailOptions) HasSubscriptionWelcomeEmailID() bool`

HasSubscriptionWelcomeEmailID returns a boolean if a field has been set.

### SetSubscriptionWelcomeEmailIDNil

`func (o *EmailOptions) SetSubscriptionWelcomeEmailIDNil(b bool)`

 SetSubscriptionWelcomeEmailIDNil sets the value for SubscriptionWelcomeEmailID to be an explicit nil

### UnsetSubscriptionWelcomeEmailID
`func (o *EmailOptions) UnsetSubscriptionWelcomeEmailID()`

UnsetSubscriptionWelcomeEmailID ensures that no value is present for SubscriptionWelcomeEmailID, not even an explicit nil
### GetSuspendedSubscriptionEmailID

`func (o *EmailOptions) GetSuspendedSubscriptionEmailID() string`

GetSuspendedSubscriptionEmailID returns the SuspendedSubscriptionEmailID field if non-nil, zero value otherwise.

### GetSuspendedSubscriptionEmailIDOk

`func (o *EmailOptions) GetSuspendedSubscriptionEmailIDOk() (*string, bool)`

GetSuspendedSubscriptionEmailIDOk returns a tuple with the SuspendedSubscriptionEmailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendedSubscriptionEmailID

`func (o *EmailOptions) SetSuspendedSubscriptionEmailID(v string)`

SetSuspendedSubscriptionEmailID sets SuspendedSubscriptionEmailID field to given value.

### HasSuspendedSubscriptionEmailID

`func (o *EmailOptions) HasSuspendedSubscriptionEmailID() bool`

HasSuspendedSubscriptionEmailID returns a boolean if a field has been set.

### SetSuspendedSubscriptionEmailIDNil

`func (o *EmailOptions) SetSuspendedSubscriptionEmailIDNil(b bool)`

 SetSuspendedSubscriptionEmailIDNil sets the value for SuspendedSubscriptionEmailID to be an explicit nil

### UnsetSuspendedSubscriptionEmailID
`func (o *EmailOptions) UnsetSuspendedSubscriptionEmailID()`

UnsetSuspendedSubscriptionEmailID ensures that no value is present for SuspendedSubscriptionEmailID, not even an explicit nil
### GetOverdueSubscriptionEmailID

`func (o *EmailOptions) GetOverdueSubscriptionEmailID() string`

GetOverdueSubscriptionEmailID returns the OverdueSubscriptionEmailID field if non-nil, zero value otherwise.

### GetOverdueSubscriptionEmailIDOk

`func (o *EmailOptions) GetOverdueSubscriptionEmailIDOk() (*string, bool)`

GetOverdueSubscriptionEmailIDOk returns a tuple with the OverdueSubscriptionEmailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdueSubscriptionEmailID

`func (o *EmailOptions) SetOverdueSubscriptionEmailID(v string)`

SetOverdueSubscriptionEmailID sets OverdueSubscriptionEmailID field to given value.

### HasOverdueSubscriptionEmailID

`func (o *EmailOptions) HasOverdueSubscriptionEmailID() bool`

HasOverdueSubscriptionEmailID returns a boolean if a field has been set.

### SetOverdueSubscriptionEmailIDNil

`func (o *EmailOptions) SetOverdueSubscriptionEmailIDNil(b bool)`

 SetOverdueSubscriptionEmailIDNil sets the value for OverdueSubscriptionEmailID to be an explicit nil

### UnsetOverdueSubscriptionEmailID
`func (o *EmailOptions) UnsetOverdueSubscriptionEmailID()`

UnsetOverdueSubscriptionEmailID ensures that no value is present for OverdueSubscriptionEmailID, not even an explicit nil
### GetExpiredSubscriptionEmailID

`func (o *EmailOptions) GetExpiredSubscriptionEmailID() string`

GetExpiredSubscriptionEmailID returns the ExpiredSubscriptionEmailID field if non-nil, zero value otherwise.

### GetExpiredSubscriptionEmailIDOk

`func (o *EmailOptions) GetExpiredSubscriptionEmailIDOk() (*string, bool)`

GetExpiredSubscriptionEmailIDOk returns a tuple with the ExpiredSubscriptionEmailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredSubscriptionEmailID

`func (o *EmailOptions) SetExpiredSubscriptionEmailID(v string)`

SetExpiredSubscriptionEmailID sets ExpiredSubscriptionEmailID field to given value.

### HasExpiredSubscriptionEmailID

`func (o *EmailOptions) HasExpiredSubscriptionEmailID() bool`

HasExpiredSubscriptionEmailID returns a boolean if a field has been set.

### SetExpiredSubscriptionEmailIDNil

`func (o *EmailOptions) SetExpiredSubscriptionEmailIDNil(b bool)`

 SetExpiredSubscriptionEmailIDNil sets the value for ExpiredSubscriptionEmailID to be an explicit nil

### UnsetExpiredSubscriptionEmailID
`func (o *EmailOptions) UnsetExpiredSubscriptionEmailID()`

UnsetExpiredSubscriptionEmailID ensures that no value is present for ExpiredSubscriptionEmailID, not even an explicit nil
### GetSwitchedSubscriptionEmailID

`func (o *EmailOptions) GetSwitchedSubscriptionEmailID() string`

GetSwitchedSubscriptionEmailID returns the SwitchedSubscriptionEmailID field if non-nil, zero value otherwise.

### GetSwitchedSubscriptionEmailIDOk

`func (o *EmailOptions) GetSwitchedSubscriptionEmailIDOk() (*string, bool)`

GetSwitchedSubscriptionEmailIDOk returns a tuple with the SwitchedSubscriptionEmailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchedSubscriptionEmailID

`func (o *EmailOptions) SetSwitchedSubscriptionEmailID(v string)`

SetSwitchedSubscriptionEmailID sets SwitchedSubscriptionEmailID field to given value.

### HasSwitchedSubscriptionEmailID

`func (o *EmailOptions) HasSwitchedSubscriptionEmailID() bool`

HasSwitchedSubscriptionEmailID returns a boolean if a field has been set.

### SetSwitchedSubscriptionEmailIDNil

`func (o *EmailOptions) SetSwitchedSubscriptionEmailIDNil(b bool)`

 SetSwitchedSubscriptionEmailIDNil sets the value for SwitchedSubscriptionEmailID to be an explicit nil

### UnsetSwitchedSubscriptionEmailID
`func (o *EmailOptions) UnsetSwitchedSubscriptionEmailID()`

UnsetSwitchedSubscriptionEmailID ensures that no value is present for SwitchedSubscriptionEmailID, not even an explicit nil
### GetNewAccountEmailID

`func (o *EmailOptions) GetNewAccountEmailID() string`

GetNewAccountEmailID returns the NewAccountEmailID field if non-nil, zero value otherwise.

### GetNewAccountEmailIDOk

`func (o *EmailOptions) GetNewAccountEmailIDOk() (*string, bool)`

GetNewAccountEmailIDOk returns a tuple with the NewAccountEmailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAccountEmailID

`func (o *EmailOptions) SetNewAccountEmailID(v string)`

SetNewAccountEmailID sets NewAccountEmailID field to given value.

### HasNewAccountEmailID

`func (o *EmailOptions) HasNewAccountEmailID() bool`

HasNewAccountEmailID returns a boolean if a field has been set.

### SetNewAccountEmailIDNil

`func (o *EmailOptions) SetNewAccountEmailIDNil(b bool)`

 SetNewAccountEmailIDNil sets the value for NewAccountEmailID to be an explicit nil

### UnsetNewAccountEmailID
`func (o *EmailOptions) UnsetNewAccountEmailID()`

UnsetNewAccountEmailID ensures that no value is present for NewAccountEmailID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


