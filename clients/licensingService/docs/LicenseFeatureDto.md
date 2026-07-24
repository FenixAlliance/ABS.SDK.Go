# LicenseFeatureDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Key** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**LicenseTypeId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLicenseFeatureDto

`func NewLicenseFeatureDto() *LicenseFeatureDto`

NewLicenseFeatureDto instantiates a new LicenseFeatureDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseFeatureDtoWithDefaults

`func NewLicenseFeatureDtoWithDefaults() *LicenseFeatureDto`

NewLicenseFeatureDtoWithDefaults instantiates a new LicenseFeatureDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LicenseFeatureDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LicenseFeatureDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LicenseFeatureDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LicenseFeatureDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *LicenseFeatureDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *LicenseFeatureDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *LicenseFeatureDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LicenseFeatureDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LicenseFeatureDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LicenseFeatureDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *LicenseFeatureDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *LicenseFeatureDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetCode

`func (o *LicenseFeatureDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *LicenseFeatureDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *LicenseFeatureDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *LicenseFeatureDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *LicenseFeatureDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *LicenseFeatureDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetKey

`func (o *LicenseFeatureDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *LicenseFeatureDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *LicenseFeatureDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *LicenseFeatureDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *LicenseFeatureDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *LicenseFeatureDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetValue

`func (o *LicenseFeatureDto) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LicenseFeatureDto) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LicenseFeatureDto) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *LicenseFeatureDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *LicenseFeatureDto) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *LicenseFeatureDto) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetName

`func (o *LicenseFeatureDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LicenseFeatureDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LicenseFeatureDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LicenseFeatureDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *LicenseFeatureDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *LicenseFeatureDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *LicenseFeatureDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LicenseFeatureDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LicenseFeatureDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LicenseFeatureDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LicenseFeatureDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LicenseFeatureDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLicenseTypeId

`func (o *LicenseFeatureDto) GetLicenseTypeId() string`

GetLicenseTypeId returns the LicenseTypeId field if non-nil, zero value otherwise.

### GetLicenseTypeIdOk

`func (o *LicenseFeatureDto) GetLicenseTypeIdOk() (*string, bool)`

GetLicenseTypeIdOk returns a tuple with the LicenseTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseTypeId

`func (o *LicenseFeatureDto) SetLicenseTypeId(v string)`

SetLicenseTypeId sets LicenseTypeId field to given value.

### HasLicenseTypeId

`func (o *LicenseFeatureDto) HasLicenseTypeId() bool`

HasLicenseTypeId returns a boolean if a field has been set.

### SetLicenseTypeIdNil

`func (o *LicenseFeatureDto) SetLicenseTypeIdNil(b bool)`

 SetLicenseTypeIdNil sets the value for LicenseTypeId to be an explicit nil

### UnsetLicenseTypeId
`func (o *LicenseFeatureDto) UnsetLicenseTypeId()`

UnsetLicenseTypeId ensures that no value is present for LicenseTypeId, not even an explicit nil
### GetTenantId

`func (o *LicenseFeatureDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *LicenseFeatureDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *LicenseFeatureDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *LicenseFeatureDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *LicenseFeatureDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *LicenseFeatureDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *LicenseFeatureDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *LicenseFeatureDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *LicenseFeatureDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *LicenseFeatureDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *LicenseFeatureDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *LicenseFeatureDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


