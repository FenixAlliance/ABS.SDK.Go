# TenantTerritoryCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**BusinessID** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordID** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ParentTerritoryID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTenantTerritoryCreateDto

`func NewTenantTerritoryCreateDto() *TenantTerritoryCreateDto`

NewTenantTerritoryCreateDto instantiates a new TenantTerritoryCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantTerritoryCreateDtoWithDefaults

`func NewTenantTerritoryCreateDtoWithDefaults() *TenantTerritoryCreateDto`

NewTenantTerritoryCreateDtoWithDefaults instantiates a new TenantTerritoryCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantTerritoryCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantTerritoryCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantTerritoryCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantTerritoryCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *TenantTerritoryCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TenantTerritoryCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TenantTerritoryCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TenantTerritoryCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetBusinessID

`func (o *TenantTerritoryCreateDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *TenantTerritoryCreateDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *TenantTerritoryCreateDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *TenantTerritoryCreateDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *TenantTerritoryCreateDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *TenantTerritoryCreateDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil
### GetBusinessProfileRecordID

`func (o *TenantTerritoryCreateDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *TenantTerritoryCreateDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *TenantTerritoryCreateDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *TenantTerritoryCreateDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *TenantTerritoryCreateDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *TenantTerritoryCreateDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil
### GetName

`func (o *TenantTerritoryCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantTerritoryCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantTerritoryCreateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TenantTerritoryCreateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TenantTerritoryCreateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TenantTerritoryCreateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *TenantTerritoryCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TenantTerritoryCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TenantTerritoryCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TenantTerritoryCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TenantTerritoryCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TenantTerritoryCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetParentTerritoryID

`func (o *TenantTerritoryCreateDto) GetParentTerritoryID() string`

GetParentTerritoryID returns the ParentTerritoryID field if non-nil, zero value otherwise.

### GetParentTerritoryIDOk

`func (o *TenantTerritoryCreateDto) GetParentTerritoryIDOk() (*string, bool)`

GetParentTerritoryIDOk returns a tuple with the ParentTerritoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTerritoryID

`func (o *TenantTerritoryCreateDto) SetParentTerritoryID(v string)`

SetParentTerritoryID sets ParentTerritoryID field to given value.

### HasParentTerritoryID

`func (o *TenantTerritoryCreateDto) HasParentTerritoryID() bool`

HasParentTerritoryID returns a boolean if a field has been set.

### SetParentTerritoryIDNil

`func (o *TenantTerritoryCreateDto) SetParentTerritoryIDNil(b bool)`

 SetParentTerritoryIDNil sets the value for ParentTerritoryID to be an explicit nil

### UnsetParentTerritoryID
`func (o *TenantTerritoryCreateDto) UnsetParentTerritoryID()`

UnsetParentTerritoryID ensures that no value is present for ParentTerritoryID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


