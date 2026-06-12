# WebPortalCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Root** | Pointer to **bool** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Domain** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**WebsiteThemeId** | Pointer to **NullableString** |  | [optional] 
**BusinessDomainId** | Pointer to **NullableString** |  | [optional] 
**BusinessPortalApplicationId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWebPortalCreateDto

`func NewWebPortalCreateDto() *WebPortalCreateDto`

NewWebPortalCreateDto instantiates a new WebPortalCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebPortalCreateDtoWithDefaults

`func NewWebPortalCreateDtoWithDefaults() *WebPortalCreateDto`

NewWebPortalCreateDtoWithDefaults instantiates a new WebPortalCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebPortalCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebPortalCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebPortalCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebPortalCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *WebPortalCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebPortalCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebPortalCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WebPortalCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetRoot

`func (o *WebPortalCreateDto) GetRoot() bool`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *WebPortalCreateDto) GetRootOk() (*bool, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *WebPortalCreateDto) SetRoot(v bool)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *WebPortalCreateDto) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### GetTitle

`func (o *WebPortalCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WebPortalCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WebPortalCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WebPortalCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *WebPortalCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *WebPortalCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDomain

`func (o *WebPortalCreateDto) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *WebPortalCreateDto) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *WebPortalCreateDto) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *WebPortalCreateDto) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *WebPortalCreateDto) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *WebPortalCreateDto) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetDisabled

`func (o *WebPortalCreateDto) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *WebPortalCreateDto) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *WebPortalCreateDto) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *WebPortalCreateDto) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDescription

`func (o *WebPortalCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebPortalCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebPortalCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebPortalCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WebPortalCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WebPortalCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetWebsiteThemeId

`func (o *WebPortalCreateDto) GetWebsiteThemeId() string`

GetWebsiteThemeId returns the WebsiteThemeId field if non-nil, zero value otherwise.

### GetWebsiteThemeIdOk

`func (o *WebPortalCreateDto) GetWebsiteThemeIdOk() (*string, bool)`

GetWebsiteThemeIdOk returns a tuple with the WebsiteThemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteThemeId

`func (o *WebPortalCreateDto) SetWebsiteThemeId(v string)`

SetWebsiteThemeId sets WebsiteThemeId field to given value.

### HasWebsiteThemeId

`func (o *WebPortalCreateDto) HasWebsiteThemeId() bool`

HasWebsiteThemeId returns a boolean if a field has been set.

### SetWebsiteThemeIdNil

`func (o *WebPortalCreateDto) SetWebsiteThemeIdNil(b bool)`

 SetWebsiteThemeIdNil sets the value for WebsiteThemeId to be an explicit nil

### UnsetWebsiteThemeId
`func (o *WebPortalCreateDto) UnsetWebsiteThemeId()`

UnsetWebsiteThemeId ensures that no value is present for WebsiteThemeId, not even an explicit nil
### GetBusinessDomainId

`func (o *WebPortalCreateDto) GetBusinessDomainId() string`

GetBusinessDomainId returns the BusinessDomainId field if non-nil, zero value otherwise.

### GetBusinessDomainIdOk

`func (o *WebPortalCreateDto) GetBusinessDomainIdOk() (*string, bool)`

GetBusinessDomainIdOk returns a tuple with the BusinessDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessDomainId

`func (o *WebPortalCreateDto) SetBusinessDomainId(v string)`

SetBusinessDomainId sets BusinessDomainId field to given value.

### HasBusinessDomainId

`func (o *WebPortalCreateDto) HasBusinessDomainId() bool`

HasBusinessDomainId returns a boolean if a field has been set.

### SetBusinessDomainIdNil

`func (o *WebPortalCreateDto) SetBusinessDomainIdNil(b bool)`

 SetBusinessDomainIdNil sets the value for BusinessDomainId to be an explicit nil

### UnsetBusinessDomainId
`func (o *WebPortalCreateDto) UnsetBusinessDomainId()`

UnsetBusinessDomainId ensures that no value is present for BusinessDomainId, not even an explicit nil
### GetBusinessPortalApplicationId

`func (o *WebPortalCreateDto) GetBusinessPortalApplicationId() string`

GetBusinessPortalApplicationId returns the BusinessPortalApplicationId field if non-nil, zero value otherwise.

### GetBusinessPortalApplicationIdOk

`func (o *WebPortalCreateDto) GetBusinessPortalApplicationIdOk() (*string, bool)`

GetBusinessPortalApplicationIdOk returns a tuple with the BusinessPortalApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPortalApplicationId

`func (o *WebPortalCreateDto) SetBusinessPortalApplicationId(v string)`

SetBusinessPortalApplicationId sets BusinessPortalApplicationId field to given value.

### HasBusinessPortalApplicationId

`func (o *WebPortalCreateDto) HasBusinessPortalApplicationId() bool`

HasBusinessPortalApplicationId returns a boolean if a field has been set.

### SetBusinessPortalApplicationIdNil

`func (o *WebPortalCreateDto) SetBusinessPortalApplicationIdNil(b bool)`

 SetBusinessPortalApplicationIdNil sets the value for BusinessPortalApplicationId to be an explicit nil

### UnsetBusinessPortalApplicationId
`func (o *WebPortalCreateDto) UnsetBusinessPortalApplicationId()`

UnsetBusinessPortalApplicationId ensures that no value is present for BusinessPortalApplicationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


