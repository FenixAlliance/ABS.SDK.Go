# WebPortalDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Root** | Pointer to **bool** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Domain** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**WebsiteThemeID** | Pointer to **NullableString** |  | [optional] 
**BusinessDomainID** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordID** | Pointer to **NullableString** |  | [optional] 
**BusinessPortalApplicationID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWebPortalDto

`func NewWebPortalDto() *WebPortalDto`

NewWebPortalDto instantiates a new WebPortalDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebPortalDtoWithDefaults

`func NewWebPortalDtoWithDefaults() *WebPortalDto`

NewWebPortalDtoWithDefaults instantiates a new WebPortalDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebPortalDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebPortalDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebPortalDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebPortalDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *WebPortalDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *WebPortalDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *WebPortalDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebPortalDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebPortalDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WebPortalDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *WebPortalDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *WebPortalDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetRoot

`func (o *WebPortalDto) GetRoot() bool`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *WebPortalDto) GetRootOk() (*bool, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *WebPortalDto) SetRoot(v bool)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *WebPortalDto) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### GetTitle

`func (o *WebPortalDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WebPortalDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WebPortalDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WebPortalDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *WebPortalDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *WebPortalDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDomain

`func (o *WebPortalDto) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *WebPortalDto) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *WebPortalDto) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *WebPortalDto) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *WebPortalDto) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *WebPortalDto) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetDisabled

`func (o *WebPortalDto) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *WebPortalDto) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *WebPortalDto) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *WebPortalDto) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetTenantId

`func (o *WebPortalDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *WebPortalDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *WebPortalDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *WebPortalDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *WebPortalDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *WebPortalDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetDescription

`func (o *WebPortalDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebPortalDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebPortalDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebPortalDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WebPortalDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WebPortalDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetWebsiteThemeID

`func (o *WebPortalDto) GetWebsiteThemeID() string`

GetWebsiteThemeID returns the WebsiteThemeID field if non-nil, zero value otherwise.

### GetWebsiteThemeIDOk

`func (o *WebPortalDto) GetWebsiteThemeIDOk() (*string, bool)`

GetWebsiteThemeIDOk returns a tuple with the WebsiteThemeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteThemeID

`func (o *WebPortalDto) SetWebsiteThemeID(v string)`

SetWebsiteThemeID sets WebsiteThemeID field to given value.

### HasWebsiteThemeID

`func (o *WebPortalDto) HasWebsiteThemeID() bool`

HasWebsiteThemeID returns a boolean if a field has been set.

### SetWebsiteThemeIDNil

`func (o *WebPortalDto) SetWebsiteThemeIDNil(b bool)`

 SetWebsiteThemeIDNil sets the value for WebsiteThemeID to be an explicit nil

### UnsetWebsiteThemeID
`func (o *WebPortalDto) UnsetWebsiteThemeID()`

UnsetWebsiteThemeID ensures that no value is present for WebsiteThemeID, not even an explicit nil
### GetBusinessDomainID

`func (o *WebPortalDto) GetBusinessDomainID() string`

GetBusinessDomainID returns the BusinessDomainID field if non-nil, zero value otherwise.

### GetBusinessDomainIDOk

`func (o *WebPortalDto) GetBusinessDomainIDOk() (*string, bool)`

GetBusinessDomainIDOk returns a tuple with the BusinessDomainID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessDomainID

`func (o *WebPortalDto) SetBusinessDomainID(v string)`

SetBusinessDomainID sets BusinessDomainID field to given value.

### HasBusinessDomainID

`func (o *WebPortalDto) HasBusinessDomainID() bool`

HasBusinessDomainID returns a boolean if a field has been set.

### SetBusinessDomainIDNil

`func (o *WebPortalDto) SetBusinessDomainIDNil(b bool)`

 SetBusinessDomainIDNil sets the value for BusinessDomainID to be an explicit nil

### UnsetBusinessDomainID
`func (o *WebPortalDto) UnsetBusinessDomainID()`

UnsetBusinessDomainID ensures that no value is present for BusinessDomainID, not even an explicit nil
### GetBusinessProfileRecordID

`func (o *WebPortalDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *WebPortalDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *WebPortalDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *WebPortalDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *WebPortalDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *WebPortalDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil
### GetBusinessPortalApplicationID

`func (o *WebPortalDto) GetBusinessPortalApplicationID() string`

GetBusinessPortalApplicationID returns the BusinessPortalApplicationID field if non-nil, zero value otherwise.

### GetBusinessPortalApplicationIDOk

`func (o *WebPortalDto) GetBusinessPortalApplicationIDOk() (*string, bool)`

GetBusinessPortalApplicationIDOk returns a tuple with the BusinessPortalApplicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPortalApplicationID

`func (o *WebPortalDto) SetBusinessPortalApplicationID(v string)`

SetBusinessPortalApplicationID sets BusinessPortalApplicationID field to given value.

### HasBusinessPortalApplicationID

`func (o *WebPortalDto) HasBusinessPortalApplicationID() bool`

HasBusinessPortalApplicationID returns a boolean if a field has been set.

### SetBusinessPortalApplicationIDNil

`func (o *WebPortalDto) SetBusinessPortalApplicationIDNil(b bool)`

 SetBusinessPortalApplicationIDNil sets the value for BusinessPortalApplicationID to be an explicit nil

### UnsetBusinessPortalApplicationID
`func (o *WebPortalDto) UnsetBusinessPortalApplicationID()`

UnsetBusinessPortalApplicationID ensures that no value is present for BusinessPortalApplicationID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


