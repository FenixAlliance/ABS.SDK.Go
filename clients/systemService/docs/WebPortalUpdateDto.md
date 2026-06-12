# WebPortalUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Root** | Pointer to **bool** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Domain** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**WebsiteThemeId** | Pointer to **NullableString** |  | [optional] 
**BusinessDomainId** | Pointer to **NullableString** |  | [optional] 
**BusinessPortalApplicationId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWebPortalUpdateDto

`func NewWebPortalUpdateDto() *WebPortalUpdateDto`

NewWebPortalUpdateDto instantiates a new WebPortalUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebPortalUpdateDtoWithDefaults

`func NewWebPortalUpdateDtoWithDefaults() *WebPortalUpdateDto`

NewWebPortalUpdateDtoWithDefaults instantiates a new WebPortalUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoot

`func (o *WebPortalUpdateDto) GetRoot() bool`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *WebPortalUpdateDto) GetRootOk() (*bool, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *WebPortalUpdateDto) SetRoot(v bool)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *WebPortalUpdateDto) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### GetTitle

`func (o *WebPortalUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WebPortalUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WebPortalUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WebPortalUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *WebPortalUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *WebPortalUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDomain

`func (o *WebPortalUpdateDto) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *WebPortalUpdateDto) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *WebPortalUpdateDto) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *WebPortalUpdateDto) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *WebPortalUpdateDto) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *WebPortalUpdateDto) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetDisabled

`func (o *WebPortalUpdateDto) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *WebPortalUpdateDto) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *WebPortalUpdateDto) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *WebPortalUpdateDto) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDescription

`func (o *WebPortalUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebPortalUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebPortalUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebPortalUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WebPortalUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WebPortalUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetWebsiteThemeId

`func (o *WebPortalUpdateDto) GetWebsiteThemeId() string`

GetWebsiteThemeId returns the WebsiteThemeId field if non-nil, zero value otherwise.

### GetWebsiteThemeIdOk

`func (o *WebPortalUpdateDto) GetWebsiteThemeIdOk() (*string, bool)`

GetWebsiteThemeIdOk returns a tuple with the WebsiteThemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteThemeId

`func (o *WebPortalUpdateDto) SetWebsiteThemeId(v string)`

SetWebsiteThemeId sets WebsiteThemeId field to given value.

### HasWebsiteThemeId

`func (o *WebPortalUpdateDto) HasWebsiteThemeId() bool`

HasWebsiteThemeId returns a boolean if a field has been set.

### SetWebsiteThemeIdNil

`func (o *WebPortalUpdateDto) SetWebsiteThemeIdNil(b bool)`

 SetWebsiteThemeIdNil sets the value for WebsiteThemeId to be an explicit nil

### UnsetWebsiteThemeId
`func (o *WebPortalUpdateDto) UnsetWebsiteThemeId()`

UnsetWebsiteThemeId ensures that no value is present for WebsiteThemeId, not even an explicit nil
### GetBusinessDomainId

`func (o *WebPortalUpdateDto) GetBusinessDomainId() string`

GetBusinessDomainId returns the BusinessDomainId field if non-nil, zero value otherwise.

### GetBusinessDomainIdOk

`func (o *WebPortalUpdateDto) GetBusinessDomainIdOk() (*string, bool)`

GetBusinessDomainIdOk returns a tuple with the BusinessDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessDomainId

`func (o *WebPortalUpdateDto) SetBusinessDomainId(v string)`

SetBusinessDomainId sets BusinessDomainId field to given value.

### HasBusinessDomainId

`func (o *WebPortalUpdateDto) HasBusinessDomainId() bool`

HasBusinessDomainId returns a boolean if a field has been set.

### SetBusinessDomainIdNil

`func (o *WebPortalUpdateDto) SetBusinessDomainIdNil(b bool)`

 SetBusinessDomainIdNil sets the value for BusinessDomainId to be an explicit nil

### UnsetBusinessDomainId
`func (o *WebPortalUpdateDto) UnsetBusinessDomainId()`

UnsetBusinessDomainId ensures that no value is present for BusinessDomainId, not even an explicit nil
### GetBusinessPortalApplicationId

`func (o *WebPortalUpdateDto) GetBusinessPortalApplicationId() string`

GetBusinessPortalApplicationId returns the BusinessPortalApplicationId field if non-nil, zero value otherwise.

### GetBusinessPortalApplicationIdOk

`func (o *WebPortalUpdateDto) GetBusinessPortalApplicationIdOk() (*string, bool)`

GetBusinessPortalApplicationIdOk returns a tuple with the BusinessPortalApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPortalApplicationId

`func (o *WebPortalUpdateDto) SetBusinessPortalApplicationId(v string)`

SetBusinessPortalApplicationId sets BusinessPortalApplicationId field to given value.

### HasBusinessPortalApplicationId

`func (o *WebPortalUpdateDto) HasBusinessPortalApplicationId() bool`

HasBusinessPortalApplicationId returns a boolean if a field has been set.

### SetBusinessPortalApplicationIdNil

`func (o *WebPortalUpdateDto) SetBusinessPortalApplicationIdNil(b bool)`

 SetBusinessPortalApplicationIdNil sets the value for BusinessPortalApplicationId to be an explicit nil

### UnsetBusinessPortalApplicationId
`func (o *WebPortalUpdateDto) UnsetBusinessPortalApplicationId()`

UnsetBusinessPortalApplicationId ensures that no value is present for BusinessPortalApplicationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


