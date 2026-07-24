# LicenseAttributeCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**LicenseTypeId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLicenseAttributeCreateDto

`func NewLicenseAttributeCreateDto(name string, ) *LicenseAttributeCreateDto`

NewLicenseAttributeCreateDto instantiates a new LicenseAttributeCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseAttributeCreateDtoWithDefaults

`func NewLicenseAttributeCreateDtoWithDefaults() *LicenseAttributeCreateDto`

NewLicenseAttributeCreateDtoWithDefaults instantiates a new LicenseAttributeCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LicenseAttributeCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LicenseAttributeCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LicenseAttributeCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LicenseAttributeCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *LicenseAttributeCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LicenseAttributeCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LicenseAttributeCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LicenseAttributeCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetCode

`func (o *LicenseAttributeCreateDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *LicenseAttributeCreateDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *LicenseAttributeCreateDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *LicenseAttributeCreateDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *LicenseAttributeCreateDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *LicenseAttributeCreateDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetName

`func (o *LicenseAttributeCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LicenseAttributeCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LicenseAttributeCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *LicenseAttributeCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LicenseAttributeCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LicenseAttributeCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LicenseAttributeCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LicenseAttributeCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LicenseAttributeCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLicenseTypeId

`func (o *LicenseAttributeCreateDto) GetLicenseTypeId() string`

GetLicenseTypeId returns the LicenseTypeId field if non-nil, zero value otherwise.

### GetLicenseTypeIdOk

`func (o *LicenseAttributeCreateDto) GetLicenseTypeIdOk() (*string, bool)`

GetLicenseTypeIdOk returns a tuple with the LicenseTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseTypeId

`func (o *LicenseAttributeCreateDto) SetLicenseTypeId(v string)`

SetLicenseTypeId sets LicenseTypeId field to given value.

### HasLicenseTypeId

`func (o *LicenseAttributeCreateDto) HasLicenseTypeId() bool`

HasLicenseTypeId returns a boolean if a field has been set.

### SetLicenseTypeIdNil

`func (o *LicenseAttributeCreateDto) SetLicenseTypeIdNil(b bool)`

 SetLicenseTypeIdNil sets the value for LicenseTypeId to be an explicit nil

### UnsetLicenseTypeId
`func (o *LicenseAttributeCreateDto) UnsetLicenseTypeId()`

UnsetLicenseTypeId ensures that no value is present for LicenseTypeId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


