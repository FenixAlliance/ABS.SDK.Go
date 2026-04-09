# TenantIndustryCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ParentBusinessIndustryID** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTenantIndustryCreateDto

`func NewTenantIndustryCreateDto() *TenantIndustryCreateDto`

NewTenantIndustryCreateDto instantiates a new TenantIndustryCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantIndustryCreateDtoWithDefaults

`func NewTenantIndustryCreateDtoWithDefaults() *TenantIndustryCreateDto`

NewTenantIndustryCreateDtoWithDefaults instantiates a new TenantIndustryCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantIndustryCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantIndustryCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantIndustryCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantIndustryCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *TenantIndustryCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TenantIndustryCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TenantIndustryCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TenantIndustryCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *TenantIndustryCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantIndustryCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantIndustryCreateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TenantIndustryCreateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TenantIndustryCreateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TenantIndustryCreateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParentBusinessIndustryID

`func (o *TenantIndustryCreateDto) GetParentBusinessIndustryID() string`

GetParentBusinessIndustryID returns the ParentBusinessIndustryID field if non-nil, zero value otherwise.

### GetParentBusinessIndustryIDOk

`func (o *TenantIndustryCreateDto) GetParentBusinessIndustryIDOk() (*string, bool)`

GetParentBusinessIndustryIDOk returns a tuple with the ParentBusinessIndustryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBusinessIndustryID

`func (o *TenantIndustryCreateDto) SetParentBusinessIndustryID(v string)`

SetParentBusinessIndustryID sets ParentBusinessIndustryID field to given value.

### HasParentBusinessIndustryID

`func (o *TenantIndustryCreateDto) HasParentBusinessIndustryID() bool`

HasParentBusinessIndustryID returns a boolean if a field has been set.

### SetParentBusinessIndustryIDNil

`func (o *TenantIndustryCreateDto) SetParentBusinessIndustryIDNil(b bool)`

 SetParentBusinessIndustryIDNil sets the value for ParentBusinessIndustryID to be an explicit nil

### UnsetParentBusinessIndustryID
`func (o *TenantIndustryCreateDto) UnsetParentBusinessIndustryID()`

UnsetParentBusinessIndustryID ensures that no value is present for ParentBusinessIndustryID, not even an explicit nil
### GetBusinessProfileRecordID

`func (o *TenantIndustryCreateDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *TenantIndustryCreateDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *TenantIndustryCreateDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *TenantIndustryCreateDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *TenantIndustryCreateDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *TenantIndustryCreateDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


