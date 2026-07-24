# LicenseFeatureCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Key** | **string** |  | 
**Value** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**LicenseTypeId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLicenseFeatureCreateDto

`func NewLicenseFeatureCreateDto(key string, value string, name string, ) *LicenseFeatureCreateDto`

NewLicenseFeatureCreateDto instantiates a new LicenseFeatureCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseFeatureCreateDtoWithDefaults

`func NewLicenseFeatureCreateDtoWithDefaults() *LicenseFeatureCreateDto`

NewLicenseFeatureCreateDtoWithDefaults instantiates a new LicenseFeatureCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LicenseFeatureCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LicenseFeatureCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LicenseFeatureCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LicenseFeatureCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *LicenseFeatureCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LicenseFeatureCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LicenseFeatureCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LicenseFeatureCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetCode

`func (o *LicenseFeatureCreateDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *LicenseFeatureCreateDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *LicenseFeatureCreateDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *LicenseFeatureCreateDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *LicenseFeatureCreateDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *LicenseFeatureCreateDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetKey

`func (o *LicenseFeatureCreateDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *LicenseFeatureCreateDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *LicenseFeatureCreateDto) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *LicenseFeatureCreateDto) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LicenseFeatureCreateDto) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LicenseFeatureCreateDto) SetValue(v string)`

SetValue sets Value field to given value.


### GetName

`func (o *LicenseFeatureCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LicenseFeatureCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LicenseFeatureCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *LicenseFeatureCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LicenseFeatureCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LicenseFeatureCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LicenseFeatureCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LicenseFeatureCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LicenseFeatureCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLicenseTypeId

`func (o *LicenseFeatureCreateDto) GetLicenseTypeId() string`

GetLicenseTypeId returns the LicenseTypeId field if non-nil, zero value otherwise.

### GetLicenseTypeIdOk

`func (o *LicenseFeatureCreateDto) GetLicenseTypeIdOk() (*string, bool)`

GetLicenseTypeIdOk returns a tuple with the LicenseTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseTypeId

`func (o *LicenseFeatureCreateDto) SetLicenseTypeId(v string)`

SetLicenseTypeId sets LicenseTypeId field to given value.

### HasLicenseTypeId

`func (o *LicenseFeatureCreateDto) HasLicenseTypeId() bool`

HasLicenseTypeId returns a boolean if a field has been set.

### SetLicenseTypeIdNil

`func (o *LicenseFeatureCreateDto) SetLicenseTypeIdNil(b bool)`

 SetLicenseTypeIdNil sets the value for LicenseTypeId to be an explicit nil

### UnsetLicenseTypeId
`func (o *LicenseFeatureCreateDto) UnsetLicenseTypeId()`

UnsetLicenseTypeId ensures that no value is present for LicenseTypeId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


