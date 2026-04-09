# ObjectEmailDispatchRequest

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
**Payload** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewObjectEmailDispatchRequest

`func NewObjectEmailDispatchRequest(title string, message string, culture string, uiCulture string, recipients []string, ) *ObjectEmailDispatchRequest`

NewObjectEmailDispatchRequest instantiates a new ObjectEmailDispatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectEmailDispatchRequestWithDefaults

`func NewObjectEmailDispatchRequestWithDefaults() *ObjectEmailDispatchRequest`

NewObjectEmailDispatchRequestWithDefaults instantiates a new ObjectEmailDispatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ObjectEmailDispatchRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ObjectEmailDispatchRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ObjectEmailDispatchRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetMessage

`func (o *ObjectEmailDispatchRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ObjectEmailDispatchRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ObjectEmailDispatchRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetButtonLink

`func (o *ObjectEmailDispatchRequest) GetButtonLink() string`

GetButtonLink returns the ButtonLink field if non-nil, zero value otherwise.

### GetButtonLinkOk

`func (o *ObjectEmailDispatchRequest) GetButtonLinkOk() (*string, bool)`

GetButtonLinkOk returns a tuple with the ButtonLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonLink

`func (o *ObjectEmailDispatchRequest) SetButtonLink(v string)`

SetButtonLink sets ButtonLink field to given value.

### HasButtonLink

`func (o *ObjectEmailDispatchRequest) HasButtonLink() bool`

HasButtonLink returns a boolean if a field has been set.

### SetButtonLinkNil

`func (o *ObjectEmailDispatchRequest) SetButtonLinkNil(b bool)`

 SetButtonLinkNil sets the value for ButtonLink to be an explicit nil

### UnsetButtonLink
`func (o *ObjectEmailDispatchRequest) UnsetButtonLink()`

UnsetButtonLink ensures that no value is present for ButtonLink, not even an explicit nil
### GetButtonText

`func (o *ObjectEmailDispatchRequest) GetButtonText() string`

GetButtonText returns the ButtonText field if non-nil, zero value otherwise.

### GetButtonTextOk

`func (o *ObjectEmailDispatchRequest) GetButtonTextOk() (*string, bool)`

GetButtonTextOk returns a tuple with the ButtonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonText

`func (o *ObjectEmailDispatchRequest) SetButtonText(v string)`

SetButtonText sets ButtonText field to given value.

### HasButtonText

`func (o *ObjectEmailDispatchRequest) HasButtonText() bool`

HasButtonText returns a boolean if a field has been set.

### SetButtonTextNil

`func (o *ObjectEmailDispatchRequest) SetButtonTextNil(b bool)`

 SetButtonTextNil sets the value for ButtonText to be an explicit nil

### UnsetButtonText
`func (o *ObjectEmailDispatchRequest) UnsetButtonText()`

UnsetButtonText ensures that no value is present for ButtonText, not even an explicit nil
### GetAlertMessage

`func (o *ObjectEmailDispatchRequest) GetAlertMessage() string`

GetAlertMessage returns the AlertMessage field if non-nil, zero value otherwise.

### GetAlertMessageOk

`func (o *ObjectEmailDispatchRequest) GetAlertMessageOk() (*string, bool)`

GetAlertMessageOk returns a tuple with the AlertMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertMessage

`func (o *ObjectEmailDispatchRequest) SetAlertMessage(v string)`

SetAlertMessage sets AlertMessage field to given value.

### HasAlertMessage

`func (o *ObjectEmailDispatchRequest) HasAlertMessage() bool`

HasAlertMessage returns a boolean if a field has been set.

### SetAlertMessageNil

`func (o *ObjectEmailDispatchRequest) SetAlertMessageNil(b bool)`

 SetAlertMessageNil sets the value for AlertMessage to be an explicit nil

### UnsetAlertMessage
`func (o *ObjectEmailDispatchRequest) UnsetAlertMessage()`

UnsetAlertMessage ensures that no value is present for AlertMessage, not even an explicit nil
### GetAlertType

`func (o *ObjectEmailDispatchRequest) GetAlertType() string`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *ObjectEmailDispatchRequest) GetAlertTypeOk() (*string, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *ObjectEmailDispatchRequest) SetAlertType(v string)`

SetAlertType sets AlertType field to given value.

### HasAlertType

`func (o *ObjectEmailDispatchRequest) HasAlertType() bool`

HasAlertType returns a boolean if a field has been set.

### GetCulture

`func (o *ObjectEmailDispatchRequest) GetCulture() string`

GetCulture returns the Culture field if non-nil, zero value otherwise.

### GetCultureOk

`func (o *ObjectEmailDispatchRequest) GetCultureOk() (*string, bool)`

GetCultureOk returns a tuple with the Culture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCulture

`func (o *ObjectEmailDispatchRequest) SetCulture(v string)`

SetCulture sets Culture field to given value.


### GetUiCulture

`func (o *ObjectEmailDispatchRequest) GetUiCulture() string`

GetUiCulture returns the UiCulture field if non-nil, zero value otherwise.

### GetUiCultureOk

`func (o *ObjectEmailDispatchRequest) GetUiCultureOk() (*string, bool)`

GetUiCultureOk returns a tuple with the UiCulture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiCulture

`func (o *ObjectEmailDispatchRequest) SetUiCulture(v string)`

SetUiCulture sets UiCulture field to given value.


### GetRecipients

`func (o *ObjectEmailDispatchRequest) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ObjectEmailDispatchRequest) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ObjectEmailDispatchRequest) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.


### GetContactIds

`func (o *ObjectEmailDispatchRequest) GetContactIds() []string`

GetContactIds returns the ContactIds field if non-nil, zero value otherwise.

### GetContactIdsOk

`func (o *ObjectEmailDispatchRequest) GetContactIdsOk() (*[]string, bool)`

GetContactIdsOk returns a tuple with the ContactIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactIds

`func (o *ObjectEmailDispatchRequest) SetContactIds(v []string)`

SetContactIds sets ContactIds field to given value.

### HasContactIds

`func (o *ObjectEmailDispatchRequest) HasContactIds() bool`

HasContactIds returns a boolean if a field has been set.

### SetContactIdsNil

`func (o *ObjectEmailDispatchRequest) SetContactIdsNil(b bool)`

 SetContactIdsNil sets the value for ContactIds to be an explicit nil

### UnsetContactIds
`func (o *ObjectEmailDispatchRequest) UnsetContactIds()`

UnsetContactIds ensures that no value is present for ContactIds, not even an explicit nil
### GetTenantIds

`func (o *ObjectEmailDispatchRequest) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *ObjectEmailDispatchRequest) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *ObjectEmailDispatchRequest) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *ObjectEmailDispatchRequest) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.

### SetTenantIdsNil

`func (o *ObjectEmailDispatchRequest) SetTenantIdsNil(b bool)`

 SetTenantIdsNil sets the value for TenantIds to be an explicit nil

### UnsetTenantIds
`func (o *ObjectEmailDispatchRequest) UnsetTenantIds()`

UnsetTenantIds ensures that no value is present for TenantIds, not even an explicit nil
### GetUserIds

`func (o *ObjectEmailDispatchRequest) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *ObjectEmailDispatchRequest) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *ObjectEmailDispatchRequest) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *ObjectEmailDispatchRequest) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### SetUserIdsNil

`func (o *ObjectEmailDispatchRequest) SetUserIdsNil(b bool)`

 SetUserIdsNil sets the value for UserIds to be an explicit nil

### UnsetUserIds
`func (o *ObjectEmailDispatchRequest) UnsetUserIds()`

UnsetUserIds ensures that no value is present for UserIds, not even an explicit nil
### GetTemplateUrl

`func (o *ObjectEmailDispatchRequest) GetTemplateUrl() string`

GetTemplateUrl returns the TemplateUrl field if non-nil, zero value otherwise.

### GetTemplateUrlOk

`func (o *ObjectEmailDispatchRequest) GetTemplateUrlOk() (*string, bool)`

GetTemplateUrlOk returns a tuple with the TemplateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateUrl

`func (o *ObjectEmailDispatchRequest) SetTemplateUrl(v string)`

SetTemplateUrl sets TemplateUrl field to given value.

### HasTemplateUrl

`func (o *ObjectEmailDispatchRequest) HasTemplateUrl() bool`

HasTemplateUrl returns a boolean if a field has been set.

### SetTemplateUrlNil

`func (o *ObjectEmailDispatchRequest) SetTemplateUrlNil(b bool)`

 SetTemplateUrlNil sets the value for TemplateUrl to be an explicit nil

### UnsetTemplateUrl
`func (o *ObjectEmailDispatchRequest) UnsetTemplateUrl()`

UnsetTemplateUrl ensures that no value is present for TemplateUrl, not even an explicit nil
### GetEmailTemplateId

`func (o *ObjectEmailDispatchRequest) GetEmailTemplateId() string`

GetEmailTemplateId returns the EmailTemplateId field if non-nil, zero value otherwise.

### GetEmailTemplateIdOk

`func (o *ObjectEmailDispatchRequest) GetEmailTemplateIdOk() (*string, bool)`

GetEmailTemplateIdOk returns a tuple with the EmailTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplateId

`func (o *ObjectEmailDispatchRequest) SetEmailTemplateId(v string)`

SetEmailTemplateId sets EmailTemplateId field to given value.

### HasEmailTemplateId

`func (o *ObjectEmailDispatchRequest) HasEmailTemplateId() bool`

HasEmailTemplateId returns a boolean if a field has been set.

### SetEmailTemplateIdNil

`func (o *ObjectEmailDispatchRequest) SetEmailTemplateIdNil(b bool)`

 SetEmailTemplateIdNil sets the value for EmailTemplateId to be an explicit nil

### UnsetEmailTemplateId
`func (o *ObjectEmailDispatchRequest) UnsetEmailTemplateId()`

UnsetEmailTemplateId ensures that no value is present for EmailTemplateId, not even an explicit nil
### GetPayload

`func (o *ObjectEmailDispatchRequest) GetPayload() interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ObjectEmailDispatchRequest) GetPayloadOk() (*interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ObjectEmailDispatchRequest) SetPayload(v interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *ObjectEmailDispatchRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *ObjectEmailDispatchRequest) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *ObjectEmailDispatchRequest) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


