# FinancialBookCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**TenantID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFinancialBookCreateDto

`func NewFinancialBookCreateDto(name string, ) *FinancialBookCreateDto`

NewFinancialBookCreateDto instantiates a new FinancialBookCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialBookCreateDtoWithDefaults

`func NewFinancialBookCreateDtoWithDefaults() *FinancialBookCreateDto`

NewFinancialBookCreateDtoWithDefaults instantiates a new FinancialBookCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FinancialBookCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FinancialBookCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FinancialBookCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FinancialBookCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *FinancialBookCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *FinancialBookCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *FinancialBookCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *FinancialBookCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *FinancialBookCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FinancialBookCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FinancialBookCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FinancialBookCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FinancialBookCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FinancialBookCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FinancialBookCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FinancialBookCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FinancialBookCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTenantID

`func (o *FinancialBookCreateDto) GetTenantID() string`

GetTenantID returns the TenantID field if non-nil, zero value otherwise.

### GetTenantIDOk

`func (o *FinancialBookCreateDto) GetTenantIDOk() (*string, bool)`

GetTenantIDOk returns a tuple with the TenantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantID

`func (o *FinancialBookCreateDto) SetTenantID(v string)`

SetTenantID sets TenantID field to given value.

### HasTenantID

`func (o *FinancialBookCreateDto) HasTenantID() bool`

HasTenantID returns a boolean if a field has been set.

### SetTenantIDNil

`func (o *FinancialBookCreateDto) SetTenantIDNil(b bool)`

 SetTenantIDNil sets the value for TenantID to be an explicit nil

### UnsetTenantID
`func (o *FinancialBookCreateDto) UnsetTenantID()`

UnsetTenantID ensures that no value is present for TenantID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


