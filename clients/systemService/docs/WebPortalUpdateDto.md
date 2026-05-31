# WebPortalUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Root** | Pointer to **bool** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Domain** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**WebsiteThemeID** | Pointer to **NullableString** |  | [optional] 
**BusinessDomainID** | Pointer to **NullableString** |  | [optional] 
**BusinessPortalApplicationID** | Pointer to **NullableString** |  | [optional] 

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
### GetWebsiteThemeID

`func (o *WebPortalUpdateDto) GetWebsiteThemeID() string`

GetWebsiteThemeID returns the WebsiteThemeID field if non-nil, zero value otherwise.

### GetWebsiteThemeIDOk

`func (o *WebPortalUpdateDto) GetWebsiteThemeIDOk() (*string, bool)`

GetWebsiteThemeIDOk returns a tuple with the WebsiteThemeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteThemeID

`func (o *WebPortalUpdateDto) SetWebsiteThemeID(v string)`

SetWebsiteThemeID sets WebsiteThemeID field to given value.

### HasWebsiteThemeID

`func (o *WebPortalUpdateDto) HasWebsiteThemeID() bool`

HasWebsiteThemeID returns a boolean if a field has been set.

### SetWebsiteThemeIDNil

`func (o *WebPortalUpdateDto) SetWebsiteThemeIDNil(b bool)`

 SetWebsiteThemeIDNil sets the value for WebsiteThemeID to be an explicit nil

### UnsetWebsiteThemeID
`func (o *WebPortalUpdateDto) UnsetWebsiteThemeID()`

UnsetWebsiteThemeID ensures that no value is present for WebsiteThemeID, not even an explicit nil
### GetBusinessDomainID

`func (o *WebPortalUpdateDto) GetBusinessDomainID() string`

GetBusinessDomainID returns the BusinessDomainID field if non-nil, zero value otherwise.

### GetBusinessDomainIDOk

`func (o *WebPortalUpdateDto) GetBusinessDomainIDOk() (*string, bool)`

GetBusinessDomainIDOk returns a tuple with the BusinessDomainID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessDomainID

`func (o *WebPortalUpdateDto) SetBusinessDomainID(v string)`

SetBusinessDomainID sets BusinessDomainID field to given value.

### HasBusinessDomainID

`func (o *WebPortalUpdateDto) HasBusinessDomainID() bool`

HasBusinessDomainID returns a boolean if a field has been set.

### SetBusinessDomainIDNil

`func (o *WebPortalUpdateDto) SetBusinessDomainIDNil(b bool)`

 SetBusinessDomainIDNil sets the value for BusinessDomainID to be an explicit nil

### UnsetBusinessDomainID
`func (o *WebPortalUpdateDto) UnsetBusinessDomainID()`

UnsetBusinessDomainID ensures that no value is present for BusinessDomainID, not even an explicit nil
### GetBusinessPortalApplicationID

`func (o *WebPortalUpdateDto) GetBusinessPortalApplicationID() string`

GetBusinessPortalApplicationID returns the BusinessPortalApplicationID field if non-nil, zero value otherwise.

### GetBusinessPortalApplicationIDOk

`func (o *WebPortalUpdateDto) GetBusinessPortalApplicationIDOk() (*string, bool)`

GetBusinessPortalApplicationIDOk returns a tuple with the BusinessPortalApplicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPortalApplicationID

`func (o *WebPortalUpdateDto) SetBusinessPortalApplicationID(v string)`

SetBusinessPortalApplicationID sets BusinessPortalApplicationID field to given value.

### HasBusinessPortalApplicationID

`func (o *WebPortalUpdateDto) HasBusinessPortalApplicationID() bool`

HasBusinessPortalApplicationID returns a boolean if a field has been set.

### SetBusinessPortalApplicationIDNil

`func (o *WebPortalUpdateDto) SetBusinessPortalApplicationIDNil(b bool)`

 SetBusinessPortalApplicationIDNil sets the value for BusinessPortalApplicationID to be an explicit nil

### UnsetBusinessPortalApplicationID
`func (o *WebPortalUpdateDto) UnsetBusinessPortalApplicationID()`

UnsetBusinessPortalApplicationID ensures that no value is present for BusinessPortalApplicationID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


