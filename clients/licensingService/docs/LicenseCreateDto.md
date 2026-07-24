# LicenseCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**LicenseTypeId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLicenseCreateDto

`func NewLicenseCreateDto(title string, ) *LicenseCreateDto`

NewLicenseCreateDto instantiates a new LicenseCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseCreateDtoWithDefaults

`func NewLicenseCreateDtoWithDefaults() *LicenseCreateDto`

NewLicenseCreateDtoWithDefaults instantiates a new LicenseCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LicenseCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LicenseCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LicenseCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LicenseCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *LicenseCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LicenseCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LicenseCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LicenseCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *LicenseCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LicenseCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LicenseCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *LicenseCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LicenseCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LicenseCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LicenseCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LicenseCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LicenseCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCode

`func (o *LicenseCreateDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *LicenseCreateDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *LicenseCreateDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *LicenseCreateDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *LicenseCreateDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *LicenseCreateDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetLicenseTypeId

`func (o *LicenseCreateDto) GetLicenseTypeId() string`

GetLicenseTypeId returns the LicenseTypeId field if non-nil, zero value otherwise.

### GetLicenseTypeIdOk

`func (o *LicenseCreateDto) GetLicenseTypeIdOk() (*string, bool)`

GetLicenseTypeIdOk returns a tuple with the LicenseTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseTypeId

`func (o *LicenseCreateDto) SetLicenseTypeId(v string)`

SetLicenseTypeId sets LicenseTypeId field to given value.

### HasLicenseTypeId

`func (o *LicenseCreateDto) HasLicenseTypeId() bool`

HasLicenseTypeId returns a boolean if a field has been set.

### SetLicenseTypeIdNil

`func (o *LicenseCreateDto) SetLicenseTypeIdNil(b bool)`

 SetLicenseTypeIdNil sets the value for LicenseTypeId to be an explicit nil

### UnsetLicenseTypeId
`func (o *LicenseCreateDto) UnsetLicenseTypeId()`

UnsetLicenseTypeId ensures that no value is present for LicenseTypeId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


