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
**WebsiteThemeID** | Pointer to **NullableString** |  | [optional] 
**BusinessDomainID** | Pointer to **NullableString** |  | [optional] 
**BusinessPortalApplicationID** | Pointer to **NullableString** |  | [optional] 

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
### GetWebsiteThemeID

`func (o *WebPortalCreateDto) GetWebsiteThemeID() string`

GetWebsiteThemeID returns the WebsiteThemeID field if non-nil, zero value otherwise.

### GetWebsiteThemeIDOk

`func (o *WebPortalCreateDto) GetWebsiteThemeIDOk() (*string, bool)`

GetWebsiteThemeIDOk returns a tuple with the WebsiteThemeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteThemeID

`func (o *WebPortalCreateDto) SetWebsiteThemeID(v string)`

SetWebsiteThemeID sets WebsiteThemeID field to given value.

### HasWebsiteThemeID

`func (o *WebPortalCreateDto) HasWebsiteThemeID() bool`

HasWebsiteThemeID returns a boolean if a field has been set.

### SetWebsiteThemeIDNil

`func (o *WebPortalCreateDto) SetWebsiteThemeIDNil(b bool)`

 SetWebsiteThemeIDNil sets the value for WebsiteThemeID to be an explicit nil

### UnsetWebsiteThemeID
`func (o *WebPortalCreateDto) UnsetWebsiteThemeID()`

UnsetWebsiteThemeID ensures that no value is present for WebsiteThemeID, not even an explicit nil
### GetBusinessDomainID

`func (o *WebPortalCreateDto) GetBusinessDomainID() string`

GetBusinessDomainID returns the BusinessDomainID field if non-nil, zero value otherwise.

### GetBusinessDomainIDOk

`func (o *WebPortalCreateDto) GetBusinessDomainIDOk() (*string, bool)`

GetBusinessDomainIDOk returns a tuple with the BusinessDomainID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessDomainID

`func (o *WebPortalCreateDto) SetBusinessDomainID(v string)`

SetBusinessDomainID sets BusinessDomainID field to given value.

### HasBusinessDomainID

`func (o *WebPortalCreateDto) HasBusinessDomainID() bool`

HasBusinessDomainID returns a boolean if a field has been set.

### SetBusinessDomainIDNil

`func (o *WebPortalCreateDto) SetBusinessDomainIDNil(b bool)`

 SetBusinessDomainIDNil sets the value for BusinessDomainID to be an explicit nil

### UnsetBusinessDomainID
`func (o *WebPortalCreateDto) UnsetBusinessDomainID()`

UnsetBusinessDomainID ensures that no value is present for BusinessDomainID, not even an explicit nil
### GetBusinessPortalApplicationID

`func (o *WebPortalCreateDto) GetBusinessPortalApplicationID() string`

GetBusinessPortalApplicationID returns the BusinessPortalApplicationID field if non-nil, zero value otherwise.

### GetBusinessPortalApplicationIDOk

`func (o *WebPortalCreateDto) GetBusinessPortalApplicationIDOk() (*string, bool)`

GetBusinessPortalApplicationIDOk returns a tuple with the BusinessPortalApplicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPortalApplicationID

`func (o *WebPortalCreateDto) SetBusinessPortalApplicationID(v string)`

SetBusinessPortalApplicationID sets BusinessPortalApplicationID field to given value.

### HasBusinessPortalApplicationID

`func (o *WebPortalCreateDto) HasBusinessPortalApplicationID() bool`

HasBusinessPortalApplicationID returns a boolean if a field has been set.

### SetBusinessPortalApplicationIDNil

`func (o *WebPortalCreateDto) SetBusinessPortalApplicationIDNil(b bool)`

 SetBusinessPortalApplicationIDNil sets the value for BusinessPortalApplicationID to be an explicit nil

### UnsetBusinessPortalApplicationID
`func (o *WebPortalCreateDto) UnsetBusinessPortalApplicationID()`

UnsetBusinessPortalApplicationID ensures that no value is present for BusinessPortalApplicationID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


