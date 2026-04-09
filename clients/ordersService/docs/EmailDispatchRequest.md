# EmailDispatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Message** | **string** |  | 
**ButtonLink** | Pointer to **NullableString** |  | [optional] 
**ButtonText** | Pointer to **NullableString** |  | [optional] 
**AlertMessage** | Pointer to **NullableString** |  | [optional] 
**AlertType** | Pointer to **string** |  | [optional] 
**Culture** | **string** |  | 
**UiCulture** | **string** |  | 
**Recipients** | **[]string** |  | 
**ContactIds** | Pointer to **[]string** |  | [optional] 
**TenantIds** | Pointer to **[]string** |  | [optional] 
**UserIds** | Pointer to **[]string** |  | [optional] 
**TemplateUrl** | Pointer to **NullableString** |  | [optional] 
**EmailTemplateId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEmailDispatchRequest

`func NewEmailDispatchRequest(title string, message string, culture string, uiCulture string, recipients []string, ) *EmailDispatchRequest`

NewEmailDispatchRequest instantiates a new EmailDispatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailDispatchRequestWithDefaults

`func NewEmailDispatchRequestWithDefaults() *EmailDispatchRequest`

NewEmailDispatchRequestWithDefaults instantiates a new EmailDispatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *EmailDispatchRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EmailDispatchRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EmailDispatchRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetMessage

`func (o *EmailDispatchRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EmailDispatchRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EmailDispatchRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetButtonLink

`func (o *EmailDispatchRequest) GetButtonLink() string`

GetButtonLink returns the ButtonLink field if non-nil, zero value otherwise.

### GetButtonLinkOk

`func (o *EmailDispatchRequest) GetButtonLinkOk() (*string, bool)`

GetButtonLinkOk returns a tuple with the ButtonLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonLink

`func (o *EmailDispatchRequest) SetButtonLink(v string)`

SetButtonLink sets ButtonLink field to given value.

### HasButtonLink

`func (o *EmailDispatchRequest) HasButtonLink() bool`

HasButtonLink returns a boolean if a field has been set.

### SetButtonLinkNil

`func (o *EmailDispatchRequest) SetButtonLinkNil(b bool)`

 SetButtonLinkNil sets the value for ButtonLink to be an explicit nil

### UnsetButtonLink
`func (o *EmailDispatchRequest) UnsetButtonLink()`

UnsetButtonLink ensures that no value is present for ButtonLink, not even an explicit nil
### GetButtonText

`func (o *EmailDispatchRequest) GetButtonText() string`

GetButtonText returns the ButtonText field if non-nil, zero value otherwise.

### GetButtonTextOk

`func (o *EmailDispatchRequest) GetButtonTextOk() (*string, bool)`

GetButtonTextOk returns a tuple with the ButtonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonText

`func (o *EmailDispatchRequest) SetButtonText(v string)`

SetButtonText sets ButtonText field to given value.

### HasButtonText

`func (o *EmailDispatchRequest) HasButtonText() bool`

HasButtonText returns a boolean if a field has been set.

### SetButtonTextNil

`func (o *EmailDispatchRequest) SetButtonTextNil(b bool)`

 SetButtonTextNil sets the value for ButtonText to be an explicit nil

### UnsetButtonText
`func (o *EmailDispatchRequest) UnsetButtonText()`

UnsetButtonText ensures that no value is present for ButtonText, not even an explicit nil
### GetAlertMessage

`func (o *EmailDispatchRequest) GetAlertMessage() string`

GetAlertMessage returns the AlertMessage field if non-nil, zero value otherwise.

### GetAlertMessageOk

`func (o *EmailDispatchRequest) GetAlertMessageOk() (*string, bool)`

GetAlertMessageOk returns a tuple with the AlertMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertMessage

`func (o *EmailDispatchRequest) SetAlertMessage(v string)`

SetAlertMessage sets AlertMessage field to given value.

### HasAlertMessage

`func (o *EmailDispatchRequest) HasAlertMessage() bool`

HasAlertMessage returns a boolean if a field has been set.

### SetAlertMessageNil

`func (o *EmailDispatchRequest) SetAlertMessageNil(b bool)`

 SetAlertMessageNil sets the value for AlertMessage to be an explicit nil

### UnsetAlertMessage
`func (o *EmailDispatchRequest) UnsetAlertMessage()`

UnsetAlertMessage ensures that no value is present for AlertMessage, not even an explicit nil
### GetAlertType

`func (o *EmailDispatchRequest) GetAlertType() string`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *EmailDispatchRequest) GetAlertTypeOk() (*string, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *EmailDispatchRequest) SetAlertType(v string)`

SetAlertType sets AlertType field to given value.

### HasAlertType

`func (o *EmailDispatchRequest) HasAlertType() bool`

HasAlertType returns a boolean if a field has been set.

### GetCulture

`func (o *EmailDispatchRequest) GetCulture() string`

GetCulture returns the Culture field if non-nil, zero value otherwise.

### GetCultureOk

`func (o *EmailDispatchRequest) GetCultureOk() (*string, bool)`

GetCultureOk returns a tuple with the Culture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCulture

`func (o *EmailDispatchRequest) SetCulture(v string)`

SetCulture sets Culture field to given value.


### GetUiCulture

`func (o *EmailDispatchRequest) GetUiCulture() string`

GetUiCulture returns the UiCulture field if non-nil, zero value otherwise.

### GetUiCultureOk

`func (o *EmailDispatchRequest) GetUiCultureOk() (*string, bool)`

GetUiCultureOk returns a tuple with the UiCulture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiCulture

`func (o *EmailDispatchRequest) SetUiCulture(v string)`

SetUiCulture sets UiCulture field to given value.


### GetRecipients

`func (o *EmailDispatchRequest) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *EmailDispatchRequest) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *EmailDispatchRequest) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.


### GetContactIds

`func (o *EmailDispatchRequest) GetContactIds() []string`

GetContactIds returns the ContactIds field if non-nil, zero value otherwise.

### GetContactIdsOk

`func (o *EmailDispatchRequest) GetContactIdsOk() (*[]string, bool)`

GetContactIdsOk returns a tuple with the ContactIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactIds

`func (o *EmailDispatchRequest) SetContactIds(v []string)`

SetContactIds sets ContactIds field to given value.

### HasContactIds

`func (o *EmailDispatchRequest) HasContactIds() bool`

HasContactIds returns a boolean if a field has been set.

### SetContactIdsNil

`func (o *EmailDispatchRequest) SetContactIdsNil(b bool)`

 SetContactIdsNil sets the value for ContactIds to be an explicit nil

### UnsetContactIds
`func (o *EmailDispatchRequest) UnsetContactIds()`

UnsetContactIds ensures that no value is present for ContactIds, not even an explicit nil
### GetTenantIds

`func (o *EmailDispatchRequest) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *EmailDispatchRequest) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *EmailDispatchRequest) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *EmailDispatchRequest) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.

### SetTenantIdsNil

`func (o *EmailDispatchRequest) SetTenantIdsNil(b bool)`

 SetTenantIdsNil sets the value for TenantIds to be an explicit nil

### UnsetTenantIds
`func (o *EmailDispatchRequest) UnsetTenantIds()`

UnsetTenantIds ensures that no value is present for TenantIds, not even an explicit nil
### GetUserIds

`func (o *EmailDispatchRequest) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *EmailDispatchRequest) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *EmailDispatchRequest) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *EmailDispatchRequest) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### SetUserIdsNil

`func (o *EmailDispatchRequest) SetUserIdsNil(b bool)`

 SetUserIdsNil sets the value for UserIds to be an explicit nil

### UnsetUserIds
`func (o *EmailDispatchRequest) UnsetUserIds()`

UnsetUserIds ensures that no value is present for UserIds, not even an explicit nil
### GetTemplateUrl

`func (o *EmailDispatchRequest) GetTemplateUrl() string`

GetTemplateUrl returns the TemplateUrl field if non-nil, zero value otherwise.

### GetTemplateUrlOk

`func (o *EmailDispatchRequest) GetTemplateUrlOk() (*string, bool)`

GetTemplateUrlOk returns a tuple with the TemplateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateUrl

`func (o *EmailDispatchRequest) SetTemplateUrl(v string)`

SetTemplateUrl sets TemplateUrl field to given value.

### HasTemplateUrl

`func (o *EmailDispatchRequest) HasTemplateUrl() bool`

HasTemplateUrl returns a boolean if a field has been set.

### SetTemplateUrlNil

`func (o *EmailDispatchRequest) SetTemplateUrlNil(b bool)`

 SetTemplateUrlNil sets the value for TemplateUrl to be an explicit nil

### UnsetTemplateUrl
`func (o *EmailDispatchRequest) UnsetTemplateUrl()`

UnsetTemplateUrl ensures that no value is present for TemplateUrl, not even an explicit nil
### GetEmailTemplateId

`func (o *EmailDispatchRequest) GetEmailTemplateId() string`

GetEmailTemplateId returns the EmailTemplateId field if non-nil, zero value otherwise.

### GetEmailTemplateIdOk

`func (o *EmailDispatchRequest) GetEmailTemplateIdOk() (*string, bool)`

GetEmailTemplateIdOk returns a tuple with the EmailTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplateId

`func (o *EmailDispatchRequest) SetEmailTemplateId(v string)`

SetEmailTemplateId sets EmailTemplateId field to given value.

### HasEmailTemplateId

`func (o *EmailDispatchRequest) HasEmailTemplateId() bool`

HasEmailTemplateId returns a boolean if a field has been set.

### SetEmailTemplateIdNil

`func (o *EmailDispatchRequest) SetEmailTemplateIdNil(b bool)`

 SetEmailTemplateIdNil sets the value for EmailTemplateId to be an explicit nil

### UnsetEmailTemplateId
`func (o *EmailDispatchRequest) UnsetEmailTemplateId()`

UnsetEmailTemplateId ensures that no value is present for EmailTemplateId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


