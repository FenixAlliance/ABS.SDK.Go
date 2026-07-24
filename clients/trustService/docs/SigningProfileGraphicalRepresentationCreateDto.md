# SigningProfileGraphicalRepresentationCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**SigningProfileId** | **string** |  | 
**Kind** | **string** |  | 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**FileUploadId** | Pointer to **NullableString** |  | [optional] 
**Sha256** | Pointer to **NullableString** |  | [optional] 
**VectorDataJson** | Pointer to **NullableString** |  | [optional] 
**TextValue** | Pointer to **NullableString** |  | [optional] 
**FontFamily** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 

## Methods

### NewSigningProfileGraphicalRepresentationCreateDto

`func NewSigningProfileGraphicalRepresentationCreateDto(signingProfileId string, kind string, ) *SigningProfileGraphicalRepresentationCreateDto`

NewSigningProfileGraphicalRepresentationCreateDto instantiates a new SigningProfileGraphicalRepresentationCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSigningProfileGraphicalRepresentationCreateDtoWithDefaults

`func NewSigningProfileGraphicalRepresentationCreateDtoWithDefaults() *SigningProfileGraphicalRepresentationCreateDto`

NewSigningProfileGraphicalRepresentationCreateDtoWithDefaults instantiates a new SigningProfileGraphicalRepresentationCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SigningProfileGraphicalRepresentationCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SigningProfileGraphicalRepresentationCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SigningProfileGraphicalRepresentationCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SigningProfileGraphicalRepresentationCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *SigningProfileGraphicalRepresentationCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SigningProfileGraphicalRepresentationCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SigningProfileGraphicalRepresentationCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SigningProfileGraphicalRepresentationCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetSigningProfileId

`func (o *SigningProfileGraphicalRepresentationCreateDto) GetSigningProfileId() string`

GetSigningProfileId returns the SigningProfileId field if non-nil, zero value otherwise.

### GetSigningProfileIdOk

`func (o *SigningProfileGraphicalRepresentationCreateDto) GetSigningProfileIdOk() (*string, bool)`

GetSigningProfileIdOk returns a tuple with the SigningProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningProfileId

`func (o *SigningProfileGraphicalRepresentationCreateDto) SetSigningProfileId(v string)`

SetSigningProfileId sets SigningProfileId field to given value.


### GetKind

`func (o *SigningProfileGraphicalRepresentationCreateDto) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SigningProfileGraphicalRepresentationCreateDto) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SigningProfileGraphicalRepresentationCreateDto) SetKind(v string)`

SetKind sets Kind field to given value.


### GetDisplayName

`func (o *SigningProfileGraphicalRepresentationCreateDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SigningProfileGraphicalRepresentationCreateDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SigningProfileGraphicalRepresentationCreateDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SigningProfileGraphicalRepresentationCreateDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *SigningProfileGraphicalRepresentationCreateDto) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *SigningProfileGraphicalRepresentationCreateDto) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetFileUploadId

`func (o *SigningProfileGraphicalRepresentationCreateDto) GetFileUploadId() string`

GetFileUploadId returns the FileUploadId field if non-nil, zero value otherwise.

### GetFileUploadIdOk

`func (o *SigningProfileGraphicalRepresentationCreateDto) GetFileUploadIdOk() (*string, bool)`

GetFileUploadIdOk returns a tuple with the FileUploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUploadId

`func (o *SigningProfileGraphicalRepresentationCreateDto) SetFileUploadId(v string)`

SetFileUploadId sets FileUploadId field to given value.

### HasFileUploadId

`func (o *SigningProfileGraphicalRepresentationCreateDto) HasFileUploadId() bool`

HasFileUploadId returns a boolean if a field has been set.

### SetFileUploadIdNil

`func (o *SigningProfileGraphicalRepresentationCreateDto) SetFileUploadIdNil(b bool)`

 SetFileUploadIdNil sets the value for FileUploadId to be an explicit nil

### UnsetFileUploadId
`func (o *SigningProfileGraphicalRepresentationCreateDto) UnsetFileUploadId()`

UnsetFileUploadId ensures that no value is present for FileUploadId, not even an explicit nil
### GetSha256

`func (o *SigningProfileGraphicalRepresentationCreateDto) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *SigningProfileGraphicalRepresentationCreateDto) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *SigningProfileGraphicalRepresentationCreateDto) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *SigningProfileGraphicalRepresentationCreateDto) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### SetSha256Nil

`func (o *SigningProfileGraphicalRepresentationCreateDto) SetSha256Nil(b bool)`

 SetSha256Nil sets the value for Sha256 to be an explicit nil

### UnsetSha256
`func (o *SigningProfileGraphicalRepresentationCreateDto) UnsetSha256()`

UnsetSha256 ensures that no value is present for Sha256, not even an explicit nil
### GetVectorDataJson

`func (o *SigningProfileGraphicalRepresentationCreateDto) GetVectorDataJson() string`

GetVectorDataJson returns the VectorDataJson field if non-nil, zero value otherwise.

### GetVectorDataJsonOk

`func (o *SigningProfileGraphicalRepresentationCreateDto) GetVectorDataJsonOk() (*string, bool)`

GetVectorDataJsonOk returns a tuple with the VectorDataJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorDataJson

`func (o *SigningProfileGraphicalRepresentationCreateDto) SetVectorDataJson(v string)`

SetVectorDataJson sets VectorDataJson field to given value.

### HasVectorDataJson

`func (o *SigningProfileGraphicalRepresentationCreateDto) HasVectorDataJson() bool`

HasVectorDataJson returns a boolean if a field has been set.

### SetVectorDataJsonNil

`func (o *SigningProfileGraphicalRepresentationCreateDto) SetVectorDataJsonNil(b bool)`

 SetVectorDataJsonNil sets the value for VectorDataJson to be an explicit nil

### UnsetVectorDataJson
`func (o *SigningProfileGraphicalRepresentationCreateDto) UnsetVectorDataJson()`

UnsetVectorDataJson ensures that no value is present for VectorDataJson, not even an explicit nil
### GetTextValue

`func (o *SigningProfileGraphicalRepresentationCreateDto) GetTextValue() string`

GetTextValue returns the TextValue field if non-nil, zero value otherwise.

### GetTextValueOk

`func (o *SigningProfileGraphicalRepresentationCreateDto) GetTextValueOk() (*string, bool)`

GetTextValueOk returns a tuple with the TextValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextValue

`func (o *SigningProfileGraphicalRepresentationCreateDto) SetTextValue(v string)`

SetTextValue sets TextValue field to given value.

### HasTextValue

`func (o *SigningProfileGraphicalRepresentationCreateDto) HasTextValue() bool`

HasTextValue returns a boolean if a field has been set.

### SetTextValueNil

`func (o *SigningProfileGraphicalRepresentationCreateDto) SetTextValueNil(b bool)`

 SetTextValueNil sets the value for TextValue to be an explicit nil

### UnsetTextValue
`func (o *SigningProfileGraphicalRepresentationCreateDto) UnsetTextValue()`

UnsetTextValue ensures that no value is present for TextValue, not even an explicit nil
### GetFontFamily

`func (o *SigningProfileGraphicalRepresentationCreateDto) GetFontFamily() string`

GetFontFamily returns the FontFamily field if non-nil, zero value otherwise.

### GetFontFamilyOk

`func (o *SigningProfileGraphicalRepresentationCreateDto) GetFontFamilyOk() (*string, bool)`

GetFontFamilyOk returns a tuple with the FontFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFontFamily

`func (o *SigningProfileGraphicalRepresentationCreateDto) SetFontFamily(v string)`

SetFontFamily sets FontFamily field to given value.

### HasFontFamily

`func (o *SigningProfileGraphicalRepresentationCreateDto) HasFontFamily() bool`

HasFontFamily returns a boolean if a field has been set.

### SetFontFamilyNil

`func (o *SigningProfileGraphicalRepresentationCreateDto) SetFontFamilyNil(b bool)`

 SetFontFamilyNil sets the value for FontFamily to be an explicit nil

### UnsetFontFamily
`func (o *SigningProfileGraphicalRepresentationCreateDto) UnsetFontFamily()`

UnsetFontFamily ensures that no value is present for FontFamily, not even an explicit nil
### GetIsDefault

`func (o *SigningProfileGraphicalRepresentationCreateDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *SigningProfileGraphicalRepresentationCreateDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *SigningProfileGraphicalRepresentationCreateDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *SigningProfileGraphicalRepresentationCreateDto) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsActive

`func (o *SigningProfileGraphicalRepresentationCreateDto) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *SigningProfileGraphicalRepresentationCreateDto) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *SigningProfileGraphicalRepresentationCreateDto) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *SigningProfileGraphicalRepresentationCreateDto) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


