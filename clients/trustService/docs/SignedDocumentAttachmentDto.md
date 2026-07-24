# SignedDocumentAttachmentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**SignedDocumentId** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**FileName** | Pointer to **NullableString** |  | [optional] 
**ContentType** | Pointer to **NullableString** |  | [optional] 
**FileLength** | Pointer to **int64** |  | [optional] 
**Hash** | Pointer to **NullableString** |  | [optional] 
**FileUploadUrl** | Pointer to **NullableString** |  | [optional] 
**StorageKey** | Pointer to **NullableString** |  | [optional] 
**StorageProviderKey** | Pointer to **NullableString** |  | [optional] 
**ScanStatus** | Pointer to **string** |  | [optional] 
**Mutability** | Pointer to **string** |  | [optional] 
**AttachmentRole** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**SignedDocumentTitle** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSignedDocumentAttachmentDto

`func NewSignedDocumentAttachmentDto() *SignedDocumentAttachmentDto`

NewSignedDocumentAttachmentDto instantiates a new SignedDocumentAttachmentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignedDocumentAttachmentDtoWithDefaults

`func NewSignedDocumentAttachmentDtoWithDefaults() *SignedDocumentAttachmentDto`

NewSignedDocumentAttachmentDtoWithDefaults instantiates a new SignedDocumentAttachmentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SignedDocumentAttachmentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SignedDocumentAttachmentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SignedDocumentAttachmentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SignedDocumentAttachmentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SignedDocumentAttachmentDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SignedDocumentAttachmentDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SignedDocumentAttachmentDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SignedDocumentAttachmentDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SignedDocumentAttachmentDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SignedDocumentAttachmentDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SignedDocumentAttachmentDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SignedDocumentAttachmentDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetSignedDocumentId

`func (o *SignedDocumentAttachmentDto) GetSignedDocumentId() string`

GetSignedDocumentId returns the SignedDocumentId field if non-nil, zero value otherwise.

### GetSignedDocumentIdOk

`func (o *SignedDocumentAttachmentDto) GetSignedDocumentIdOk() (*string, bool)`

GetSignedDocumentIdOk returns a tuple with the SignedDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedDocumentId

`func (o *SignedDocumentAttachmentDto) SetSignedDocumentId(v string)`

SetSignedDocumentId sets SignedDocumentId field to given value.

### HasSignedDocumentId

`func (o *SignedDocumentAttachmentDto) HasSignedDocumentId() bool`

HasSignedDocumentId returns a boolean if a field has been set.

### SetSignedDocumentIdNil

`func (o *SignedDocumentAttachmentDto) SetSignedDocumentIdNil(b bool)`

 SetSignedDocumentIdNil sets the value for SignedDocumentId to be an explicit nil

### UnsetSignedDocumentId
`func (o *SignedDocumentAttachmentDto) UnsetSignedDocumentId()`

UnsetSignedDocumentId ensures that no value is present for SignedDocumentId, not even an explicit nil
### GetTitle

`func (o *SignedDocumentAttachmentDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SignedDocumentAttachmentDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SignedDocumentAttachmentDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SignedDocumentAttachmentDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SignedDocumentAttachmentDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SignedDocumentAttachmentDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetFileName

`func (o *SignedDocumentAttachmentDto) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *SignedDocumentAttachmentDto) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *SignedDocumentAttachmentDto) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *SignedDocumentAttachmentDto) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *SignedDocumentAttachmentDto) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *SignedDocumentAttachmentDto) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetContentType

`func (o *SignedDocumentAttachmentDto) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *SignedDocumentAttachmentDto) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *SignedDocumentAttachmentDto) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *SignedDocumentAttachmentDto) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *SignedDocumentAttachmentDto) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *SignedDocumentAttachmentDto) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetFileLength

`func (o *SignedDocumentAttachmentDto) GetFileLength() int64`

GetFileLength returns the FileLength field if non-nil, zero value otherwise.

### GetFileLengthOk

`func (o *SignedDocumentAttachmentDto) GetFileLengthOk() (*int64, bool)`

GetFileLengthOk returns a tuple with the FileLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLength

`func (o *SignedDocumentAttachmentDto) SetFileLength(v int64)`

SetFileLength sets FileLength field to given value.

### HasFileLength

`func (o *SignedDocumentAttachmentDto) HasFileLength() bool`

HasFileLength returns a boolean if a field has been set.

### GetHash

`func (o *SignedDocumentAttachmentDto) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *SignedDocumentAttachmentDto) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *SignedDocumentAttachmentDto) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *SignedDocumentAttachmentDto) HasHash() bool`

HasHash returns a boolean if a field has been set.

### SetHashNil

`func (o *SignedDocumentAttachmentDto) SetHashNil(b bool)`

 SetHashNil sets the value for Hash to be an explicit nil

### UnsetHash
`func (o *SignedDocumentAttachmentDto) UnsetHash()`

UnsetHash ensures that no value is present for Hash, not even an explicit nil
### GetFileUploadUrl

`func (o *SignedDocumentAttachmentDto) GetFileUploadUrl() string`

GetFileUploadUrl returns the FileUploadUrl field if non-nil, zero value otherwise.

### GetFileUploadUrlOk

`func (o *SignedDocumentAttachmentDto) GetFileUploadUrlOk() (*string, bool)`

GetFileUploadUrlOk returns a tuple with the FileUploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUploadUrl

`func (o *SignedDocumentAttachmentDto) SetFileUploadUrl(v string)`

SetFileUploadUrl sets FileUploadUrl field to given value.

### HasFileUploadUrl

`func (o *SignedDocumentAttachmentDto) HasFileUploadUrl() bool`

HasFileUploadUrl returns a boolean if a field has been set.

### SetFileUploadUrlNil

`func (o *SignedDocumentAttachmentDto) SetFileUploadUrlNil(b bool)`

 SetFileUploadUrlNil sets the value for FileUploadUrl to be an explicit nil

### UnsetFileUploadUrl
`func (o *SignedDocumentAttachmentDto) UnsetFileUploadUrl()`

UnsetFileUploadUrl ensures that no value is present for FileUploadUrl, not even an explicit nil
### GetStorageKey

`func (o *SignedDocumentAttachmentDto) GetStorageKey() string`

GetStorageKey returns the StorageKey field if non-nil, zero value otherwise.

### GetStorageKeyOk

`func (o *SignedDocumentAttachmentDto) GetStorageKeyOk() (*string, bool)`

GetStorageKeyOk returns a tuple with the StorageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageKey

`func (o *SignedDocumentAttachmentDto) SetStorageKey(v string)`

SetStorageKey sets StorageKey field to given value.

### HasStorageKey

`func (o *SignedDocumentAttachmentDto) HasStorageKey() bool`

HasStorageKey returns a boolean if a field has been set.

### SetStorageKeyNil

`func (o *SignedDocumentAttachmentDto) SetStorageKeyNil(b bool)`

 SetStorageKeyNil sets the value for StorageKey to be an explicit nil

### UnsetStorageKey
`func (o *SignedDocumentAttachmentDto) UnsetStorageKey()`

UnsetStorageKey ensures that no value is present for StorageKey, not even an explicit nil
### GetStorageProviderKey

`func (o *SignedDocumentAttachmentDto) GetStorageProviderKey() string`

GetStorageProviderKey returns the StorageProviderKey field if non-nil, zero value otherwise.

### GetStorageProviderKeyOk

`func (o *SignedDocumentAttachmentDto) GetStorageProviderKeyOk() (*string, bool)`

GetStorageProviderKeyOk returns a tuple with the StorageProviderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProviderKey

`func (o *SignedDocumentAttachmentDto) SetStorageProviderKey(v string)`

SetStorageProviderKey sets StorageProviderKey field to given value.

### HasStorageProviderKey

`func (o *SignedDocumentAttachmentDto) HasStorageProviderKey() bool`

HasStorageProviderKey returns a boolean if a field has been set.

### SetStorageProviderKeyNil

`func (o *SignedDocumentAttachmentDto) SetStorageProviderKeyNil(b bool)`

 SetStorageProviderKeyNil sets the value for StorageProviderKey to be an explicit nil

### UnsetStorageProviderKey
`func (o *SignedDocumentAttachmentDto) UnsetStorageProviderKey()`

UnsetStorageProviderKey ensures that no value is present for StorageProviderKey, not even an explicit nil
### GetScanStatus

`func (o *SignedDocumentAttachmentDto) GetScanStatus() string`

GetScanStatus returns the ScanStatus field if non-nil, zero value otherwise.

### GetScanStatusOk

`func (o *SignedDocumentAttachmentDto) GetScanStatusOk() (*string, bool)`

GetScanStatusOk returns a tuple with the ScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanStatus

`func (o *SignedDocumentAttachmentDto) SetScanStatus(v string)`

SetScanStatus sets ScanStatus field to given value.

### HasScanStatus

`func (o *SignedDocumentAttachmentDto) HasScanStatus() bool`

HasScanStatus returns a boolean if a field has been set.

### GetMutability

`func (o *SignedDocumentAttachmentDto) GetMutability() string`

GetMutability returns the Mutability field if non-nil, zero value otherwise.

### GetMutabilityOk

`func (o *SignedDocumentAttachmentDto) GetMutabilityOk() (*string, bool)`

GetMutabilityOk returns a tuple with the Mutability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutability

`func (o *SignedDocumentAttachmentDto) SetMutability(v string)`

SetMutability sets Mutability field to given value.

### HasMutability

`func (o *SignedDocumentAttachmentDto) HasMutability() bool`

HasMutability returns a boolean if a field has been set.

### GetAttachmentRole

`func (o *SignedDocumentAttachmentDto) GetAttachmentRole() string`

GetAttachmentRole returns the AttachmentRole field if non-nil, zero value otherwise.

### GetAttachmentRoleOk

`func (o *SignedDocumentAttachmentDto) GetAttachmentRoleOk() (*string, bool)`

GetAttachmentRoleOk returns a tuple with the AttachmentRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentRole

`func (o *SignedDocumentAttachmentDto) SetAttachmentRole(v string)`

SetAttachmentRole sets AttachmentRole field to given value.

### HasAttachmentRole

`func (o *SignedDocumentAttachmentDto) HasAttachmentRole() bool`

HasAttachmentRole returns a boolean if a field has been set.

### SetAttachmentRoleNil

`func (o *SignedDocumentAttachmentDto) SetAttachmentRoleNil(b bool)`

 SetAttachmentRoleNil sets the value for AttachmentRole to be an explicit nil

### UnsetAttachmentRole
`func (o *SignedDocumentAttachmentDto) UnsetAttachmentRole()`

UnsetAttachmentRole ensures that no value is present for AttachmentRole, not even an explicit nil
### GetTenantId

`func (o *SignedDocumentAttachmentDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SignedDocumentAttachmentDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SignedDocumentAttachmentDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SignedDocumentAttachmentDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *SignedDocumentAttachmentDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *SignedDocumentAttachmentDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetUserId

`func (o *SignedDocumentAttachmentDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SignedDocumentAttachmentDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SignedDocumentAttachmentDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SignedDocumentAttachmentDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *SignedDocumentAttachmentDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *SignedDocumentAttachmentDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetEnrollmentId

`func (o *SignedDocumentAttachmentDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *SignedDocumentAttachmentDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *SignedDocumentAttachmentDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *SignedDocumentAttachmentDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *SignedDocumentAttachmentDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *SignedDocumentAttachmentDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetSignedDocumentTitle

`func (o *SignedDocumentAttachmentDto) GetSignedDocumentTitle() string`

GetSignedDocumentTitle returns the SignedDocumentTitle field if non-nil, zero value otherwise.

### GetSignedDocumentTitleOk

`func (o *SignedDocumentAttachmentDto) GetSignedDocumentTitleOk() (*string, bool)`

GetSignedDocumentTitleOk returns a tuple with the SignedDocumentTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedDocumentTitle

`func (o *SignedDocumentAttachmentDto) SetSignedDocumentTitle(v string)`

SetSignedDocumentTitle sets SignedDocumentTitle field to given value.

### HasSignedDocumentTitle

`func (o *SignedDocumentAttachmentDto) HasSignedDocumentTitle() bool`

HasSignedDocumentTitle returns a boolean if a field has been set.

### SetSignedDocumentTitleNil

`func (o *SignedDocumentAttachmentDto) SetSignedDocumentTitleNil(b bool)`

 SetSignedDocumentTitleNil sets the value for SignedDocumentTitle to be an explicit nil

### UnsetSignedDocumentTitle
`func (o *SignedDocumentAttachmentDto) UnsetSignedDocumentTitle()`

UnsetSignedDocumentTitle ensures that no value is present for SignedDocumentTitle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


