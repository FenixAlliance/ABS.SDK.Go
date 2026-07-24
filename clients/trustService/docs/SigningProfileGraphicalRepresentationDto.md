# SigningProfileGraphicalRepresentationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**SigningProfileId** | Pointer to **NullableString** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**FileUploadId** | Pointer to **NullableString** |  | [optional] 
**Sha256** | Pointer to **NullableString** |  | [optional] 
**VectorDataJson** | Pointer to **NullableString** |  | [optional] 
**TextValue** | Pointer to **NullableString** |  | [optional] 
**FontFamily** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**CreatedById** | Pointer to **NullableString** |  | [optional] 
**CreatedAtUtc** | Pointer to **time.Time** |  | [optional] 
**SigningProfileName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSigningProfileGraphicalRepresentationDto

`func NewSigningProfileGraphicalRepresentationDto() *SigningProfileGraphicalRepresentationDto`

NewSigningProfileGraphicalRepresentationDto instantiates a new SigningProfileGraphicalRepresentationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSigningProfileGraphicalRepresentationDtoWithDefaults

`func NewSigningProfileGraphicalRepresentationDtoWithDefaults() *SigningProfileGraphicalRepresentationDto`

NewSigningProfileGraphicalRepresentationDtoWithDefaults instantiates a new SigningProfileGraphicalRepresentationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SigningProfileGraphicalRepresentationDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SigningProfileGraphicalRepresentationDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SigningProfileGraphicalRepresentationDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SigningProfileGraphicalRepresentationDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SigningProfileGraphicalRepresentationDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SigningProfileGraphicalRepresentationDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SigningProfileGraphicalRepresentationDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SigningProfileGraphicalRepresentationDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SigningProfileGraphicalRepresentationDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SigningProfileGraphicalRepresentationDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SigningProfileGraphicalRepresentationDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SigningProfileGraphicalRepresentationDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetSigningProfileId

`func (o *SigningProfileGraphicalRepresentationDto) GetSigningProfileId() string`

GetSigningProfileId returns the SigningProfileId field if non-nil, zero value otherwise.

### GetSigningProfileIdOk

`func (o *SigningProfileGraphicalRepresentationDto) GetSigningProfileIdOk() (*string, bool)`

GetSigningProfileIdOk returns a tuple with the SigningProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningProfileId

`func (o *SigningProfileGraphicalRepresentationDto) SetSigningProfileId(v string)`

SetSigningProfileId sets SigningProfileId field to given value.

### HasSigningProfileId

`func (o *SigningProfileGraphicalRepresentationDto) HasSigningProfileId() bool`

HasSigningProfileId returns a boolean if a field has been set.

### SetSigningProfileIdNil

`func (o *SigningProfileGraphicalRepresentationDto) SetSigningProfileIdNil(b bool)`

 SetSigningProfileIdNil sets the value for SigningProfileId to be an explicit nil

### UnsetSigningProfileId
`func (o *SigningProfileGraphicalRepresentationDto) UnsetSigningProfileId()`

UnsetSigningProfileId ensures that no value is present for SigningProfileId, not even an explicit nil
### GetKind

`func (o *SigningProfileGraphicalRepresentationDto) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SigningProfileGraphicalRepresentationDto) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SigningProfileGraphicalRepresentationDto) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SigningProfileGraphicalRepresentationDto) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetDisplayName

`func (o *SigningProfileGraphicalRepresentationDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SigningProfileGraphicalRepresentationDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SigningProfileGraphicalRepresentationDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SigningProfileGraphicalRepresentationDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *SigningProfileGraphicalRepresentationDto) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *SigningProfileGraphicalRepresentationDto) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetFileUploadId

`func (o *SigningProfileGraphicalRepresentationDto) GetFileUploadId() string`

GetFileUploadId returns the FileUploadId field if non-nil, zero value otherwise.

### GetFileUploadIdOk

`func (o *SigningProfileGraphicalRepresentationDto) GetFileUploadIdOk() (*string, bool)`

GetFileUploadIdOk returns a tuple with the FileUploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUploadId

`func (o *SigningProfileGraphicalRepresentationDto) SetFileUploadId(v string)`

SetFileUploadId sets FileUploadId field to given value.

### HasFileUploadId

`func (o *SigningProfileGraphicalRepresentationDto) HasFileUploadId() bool`

HasFileUploadId returns a boolean if a field has been set.

### SetFileUploadIdNil

`func (o *SigningProfileGraphicalRepresentationDto) SetFileUploadIdNil(b bool)`

 SetFileUploadIdNil sets the value for FileUploadId to be an explicit nil

### UnsetFileUploadId
`func (o *SigningProfileGraphicalRepresentationDto) UnsetFileUploadId()`

UnsetFileUploadId ensures that no value is present for FileUploadId, not even an explicit nil
### GetSha256

`func (o *SigningProfileGraphicalRepresentationDto) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *SigningProfileGraphicalRepresentationDto) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *SigningProfileGraphicalRepresentationDto) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *SigningProfileGraphicalRepresentationDto) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### SetSha256Nil

`func (o *SigningProfileGraphicalRepresentationDto) SetSha256Nil(b bool)`

 SetSha256Nil sets the value for Sha256 to be an explicit nil

### UnsetSha256
`func (o *SigningProfileGraphicalRepresentationDto) UnsetSha256()`

UnsetSha256 ensures that no value is present for Sha256, not even an explicit nil
### GetVectorDataJson

`func (o *SigningProfileGraphicalRepresentationDto) GetVectorDataJson() string`

GetVectorDataJson returns the VectorDataJson field if non-nil, zero value otherwise.

### GetVectorDataJsonOk

`func (o *SigningProfileGraphicalRepresentationDto) GetVectorDataJsonOk() (*string, bool)`

GetVectorDataJsonOk returns a tuple with the VectorDataJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorDataJson

`func (o *SigningProfileGraphicalRepresentationDto) SetVectorDataJson(v string)`

SetVectorDataJson sets VectorDataJson field to given value.

### HasVectorDataJson

`func (o *SigningProfileGraphicalRepresentationDto) HasVectorDataJson() bool`

HasVectorDataJson returns a boolean if a field has been set.

### SetVectorDataJsonNil

`func (o *SigningProfileGraphicalRepresentationDto) SetVectorDataJsonNil(b bool)`

 SetVectorDataJsonNil sets the value for VectorDataJson to be an explicit nil

### UnsetVectorDataJson
`func (o *SigningProfileGraphicalRepresentationDto) UnsetVectorDataJson()`

UnsetVectorDataJson ensures that no value is present for VectorDataJson, not even an explicit nil
### GetTextValue

`func (o *SigningProfileGraphicalRepresentationDto) GetTextValue() string`

GetTextValue returns the TextValue field if non-nil, zero value otherwise.

### GetTextValueOk

`func (o *SigningProfileGraphicalRepresentationDto) GetTextValueOk() (*string, bool)`

GetTextValueOk returns a tuple with the TextValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextValue

`func (o *SigningProfileGraphicalRepresentationDto) SetTextValue(v string)`

SetTextValue sets TextValue field to given value.

### HasTextValue

`func (o *SigningProfileGraphicalRepresentationDto) HasTextValue() bool`

HasTextValue returns a boolean if a field has been set.

### SetTextValueNil

`func (o *SigningProfileGraphicalRepresentationDto) SetTextValueNil(b bool)`

 SetTextValueNil sets the value for TextValue to be an explicit nil

### UnsetTextValue
`func (o *SigningProfileGraphicalRepresentationDto) UnsetTextValue()`

UnsetTextValue ensures that no value is present for TextValue, not even an explicit nil
### GetFontFamily

`func (o *SigningProfileGraphicalRepresentationDto) GetFontFamily() string`

GetFontFamily returns the FontFamily field if non-nil, zero value otherwise.

### GetFontFamilyOk

`func (o *SigningProfileGraphicalRepresentationDto) GetFontFamilyOk() (*string, bool)`

GetFontFamilyOk returns a tuple with the FontFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFontFamily

`func (o *SigningProfileGraphicalRepresentationDto) SetFontFamily(v string)`

SetFontFamily sets FontFamily field to given value.

### HasFontFamily

`func (o *SigningProfileGraphicalRepresentationDto) HasFontFamily() bool`

HasFontFamily returns a boolean if a field has been set.

### SetFontFamilyNil

`func (o *SigningProfileGraphicalRepresentationDto) SetFontFamilyNil(b bool)`

 SetFontFamilyNil sets the value for FontFamily to be an explicit nil

### UnsetFontFamily
`func (o *SigningProfileGraphicalRepresentationDto) UnsetFontFamily()`

UnsetFontFamily ensures that no value is present for FontFamily, not even an explicit nil
### GetIsDefault

`func (o *SigningProfileGraphicalRepresentationDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *SigningProfileGraphicalRepresentationDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *SigningProfileGraphicalRepresentationDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *SigningProfileGraphicalRepresentationDto) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsActive

`func (o *SigningProfileGraphicalRepresentationDto) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *SigningProfileGraphicalRepresentationDto) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *SigningProfileGraphicalRepresentationDto) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *SigningProfileGraphicalRepresentationDto) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetTenantId

`func (o *SigningProfileGraphicalRepresentationDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SigningProfileGraphicalRepresentationDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SigningProfileGraphicalRepresentationDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SigningProfileGraphicalRepresentationDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *SigningProfileGraphicalRepresentationDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *SigningProfileGraphicalRepresentationDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *SigningProfileGraphicalRepresentationDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *SigningProfileGraphicalRepresentationDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *SigningProfileGraphicalRepresentationDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *SigningProfileGraphicalRepresentationDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *SigningProfileGraphicalRepresentationDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *SigningProfileGraphicalRepresentationDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetCreatedById

`func (o *SigningProfileGraphicalRepresentationDto) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *SigningProfileGraphicalRepresentationDto) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *SigningProfileGraphicalRepresentationDto) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *SigningProfileGraphicalRepresentationDto) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### SetCreatedByIdNil

`func (o *SigningProfileGraphicalRepresentationDto) SetCreatedByIdNil(b bool)`

 SetCreatedByIdNil sets the value for CreatedById to be an explicit nil

### UnsetCreatedById
`func (o *SigningProfileGraphicalRepresentationDto) UnsetCreatedById()`

UnsetCreatedById ensures that no value is present for CreatedById, not even an explicit nil
### GetCreatedAtUtc

`func (o *SigningProfileGraphicalRepresentationDto) GetCreatedAtUtc() time.Time`

GetCreatedAtUtc returns the CreatedAtUtc field if non-nil, zero value otherwise.

### GetCreatedAtUtcOk

`func (o *SigningProfileGraphicalRepresentationDto) GetCreatedAtUtcOk() (*time.Time, bool)`

GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtUtc

`func (o *SigningProfileGraphicalRepresentationDto) SetCreatedAtUtc(v time.Time)`

SetCreatedAtUtc sets CreatedAtUtc field to given value.

### HasCreatedAtUtc

`func (o *SigningProfileGraphicalRepresentationDto) HasCreatedAtUtc() bool`

HasCreatedAtUtc returns a boolean if a field has been set.

### GetSigningProfileName

`func (o *SigningProfileGraphicalRepresentationDto) GetSigningProfileName() string`

GetSigningProfileName returns the SigningProfileName field if non-nil, zero value otherwise.

### GetSigningProfileNameOk

`func (o *SigningProfileGraphicalRepresentationDto) GetSigningProfileNameOk() (*string, bool)`

GetSigningProfileNameOk returns a tuple with the SigningProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningProfileName

`func (o *SigningProfileGraphicalRepresentationDto) SetSigningProfileName(v string)`

SetSigningProfileName sets SigningProfileName field to given value.

### HasSigningProfileName

`func (o *SigningProfileGraphicalRepresentationDto) HasSigningProfileName() bool`

HasSigningProfileName returns a boolean if a field has been set.

### SetSigningProfileNameNil

`func (o *SigningProfileGraphicalRepresentationDto) SetSigningProfileNameNil(b bool)`

 SetSigningProfileNameNil sets the value for SigningProfileName to be an explicit nil

### UnsetSigningProfileName
`func (o *SigningProfileGraphicalRepresentationDto) UnsetSigningProfileName()`

UnsetSigningProfileName ensures that no value is present for SigningProfileName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


