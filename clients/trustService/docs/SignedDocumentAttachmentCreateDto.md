# SignedDocumentAttachmentCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**SignedDocumentId** | **string** |  | 
**Title** | **string** |  | 
**FileName** | Pointer to **NullableString** |  | [optional] 
**ContentType** | Pointer to **NullableString** |  | [optional] 
**FileLength** | Pointer to **int64** |  | [optional] 
**Hash** | Pointer to **NullableString** |  | [optional] 
**FileUploadUrl** | Pointer to **NullableString** |  | [optional] 
**StorageKey** | **string** |  | 
**StorageProviderKey** | Pointer to **NullableString** |  | [optional] 
**AttachmentRole** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSignedDocumentAttachmentCreateDto

`func NewSignedDocumentAttachmentCreateDto(signedDocumentId string, title string, storageKey string, ) *SignedDocumentAttachmentCreateDto`

NewSignedDocumentAttachmentCreateDto instantiates a new SignedDocumentAttachmentCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignedDocumentAttachmentCreateDtoWithDefaults

`func NewSignedDocumentAttachmentCreateDtoWithDefaults() *SignedDocumentAttachmentCreateDto`

NewSignedDocumentAttachmentCreateDtoWithDefaults instantiates a new SignedDocumentAttachmentCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SignedDocumentAttachmentCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SignedDocumentAttachmentCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SignedDocumentAttachmentCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SignedDocumentAttachmentCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *SignedDocumentAttachmentCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SignedDocumentAttachmentCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SignedDocumentAttachmentCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SignedDocumentAttachmentCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetSignedDocumentId

`func (o *SignedDocumentAttachmentCreateDto) GetSignedDocumentId() string`

GetSignedDocumentId returns the SignedDocumentId field if non-nil, zero value otherwise.

### GetSignedDocumentIdOk

`func (o *SignedDocumentAttachmentCreateDto) GetSignedDocumentIdOk() (*string, bool)`

GetSignedDocumentIdOk returns a tuple with the SignedDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedDocumentId

`func (o *SignedDocumentAttachmentCreateDto) SetSignedDocumentId(v string)`

SetSignedDocumentId sets SignedDocumentId field to given value.


### GetTitle

`func (o *SignedDocumentAttachmentCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SignedDocumentAttachmentCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SignedDocumentAttachmentCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetFileName

`func (o *SignedDocumentAttachmentCreateDto) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *SignedDocumentAttachmentCreateDto) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *SignedDocumentAttachmentCreateDto) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *SignedDocumentAttachmentCreateDto) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *SignedDocumentAttachmentCreateDto) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *SignedDocumentAttachmentCreateDto) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetContentType

`func (o *SignedDocumentAttachmentCreateDto) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *SignedDocumentAttachmentCreateDto) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *SignedDocumentAttachmentCreateDto) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *SignedDocumentAttachmentCreateDto) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *SignedDocumentAttachmentCreateDto) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *SignedDocumentAttachmentCreateDto) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetFileLength

`func (o *SignedDocumentAttachmentCreateDto) GetFileLength() int64`

GetFileLength returns the FileLength field if non-nil, zero value otherwise.

### GetFileLengthOk

`func (o *SignedDocumentAttachmentCreateDto) GetFileLengthOk() (*int64, bool)`

GetFileLengthOk returns a tuple with the FileLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLength

`func (o *SignedDocumentAttachmentCreateDto) SetFileLength(v int64)`

SetFileLength sets FileLength field to given value.

### HasFileLength

`func (o *SignedDocumentAttachmentCreateDto) HasFileLength() bool`

HasFileLength returns a boolean if a field has been set.

### GetHash

`func (o *SignedDocumentAttachmentCreateDto) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *SignedDocumentAttachmentCreateDto) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *SignedDocumentAttachmentCreateDto) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *SignedDocumentAttachmentCreateDto) HasHash() bool`

HasHash returns a boolean if a field has been set.

### SetHashNil

`func (o *SignedDocumentAttachmentCreateDto) SetHashNil(b bool)`

 SetHashNil sets the value for Hash to be an explicit nil

### UnsetHash
`func (o *SignedDocumentAttachmentCreateDto) UnsetHash()`

UnsetHash ensures that no value is present for Hash, not even an explicit nil
### GetFileUploadUrl

`func (o *SignedDocumentAttachmentCreateDto) GetFileUploadUrl() string`

GetFileUploadUrl returns the FileUploadUrl field if non-nil, zero value otherwise.

### GetFileUploadUrlOk

`func (o *SignedDocumentAttachmentCreateDto) GetFileUploadUrlOk() (*string, bool)`

GetFileUploadUrlOk returns a tuple with the FileUploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUploadUrl

`func (o *SignedDocumentAttachmentCreateDto) SetFileUploadUrl(v string)`

SetFileUploadUrl sets FileUploadUrl field to given value.

### HasFileUploadUrl

`func (o *SignedDocumentAttachmentCreateDto) HasFileUploadUrl() bool`

HasFileUploadUrl returns a boolean if a field has been set.

### SetFileUploadUrlNil

`func (o *SignedDocumentAttachmentCreateDto) SetFileUploadUrlNil(b bool)`

 SetFileUploadUrlNil sets the value for FileUploadUrl to be an explicit nil

### UnsetFileUploadUrl
`func (o *SignedDocumentAttachmentCreateDto) UnsetFileUploadUrl()`

UnsetFileUploadUrl ensures that no value is present for FileUploadUrl, not even an explicit nil
### GetStorageKey

`func (o *SignedDocumentAttachmentCreateDto) GetStorageKey() string`

GetStorageKey returns the StorageKey field if non-nil, zero value otherwise.

### GetStorageKeyOk

`func (o *SignedDocumentAttachmentCreateDto) GetStorageKeyOk() (*string, bool)`

GetStorageKeyOk returns a tuple with the StorageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageKey

`func (o *SignedDocumentAttachmentCreateDto) SetStorageKey(v string)`

SetStorageKey sets StorageKey field to given value.


### GetStorageProviderKey

`func (o *SignedDocumentAttachmentCreateDto) GetStorageProviderKey() string`

GetStorageProviderKey returns the StorageProviderKey field if non-nil, zero value otherwise.

### GetStorageProviderKeyOk

`func (o *SignedDocumentAttachmentCreateDto) GetStorageProviderKeyOk() (*string, bool)`

GetStorageProviderKeyOk returns a tuple with the StorageProviderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProviderKey

`func (o *SignedDocumentAttachmentCreateDto) SetStorageProviderKey(v string)`

SetStorageProviderKey sets StorageProviderKey field to given value.

### HasStorageProviderKey

`func (o *SignedDocumentAttachmentCreateDto) HasStorageProviderKey() bool`

HasStorageProviderKey returns a boolean if a field has been set.

### SetStorageProviderKeyNil

`func (o *SignedDocumentAttachmentCreateDto) SetStorageProviderKeyNil(b bool)`

 SetStorageProviderKeyNil sets the value for StorageProviderKey to be an explicit nil

### UnsetStorageProviderKey
`func (o *SignedDocumentAttachmentCreateDto) UnsetStorageProviderKey()`

UnsetStorageProviderKey ensures that no value is present for StorageProviderKey, not even an explicit nil
### GetAttachmentRole

`func (o *SignedDocumentAttachmentCreateDto) GetAttachmentRole() string`

GetAttachmentRole returns the AttachmentRole field if non-nil, zero value otherwise.

### GetAttachmentRoleOk

`func (o *SignedDocumentAttachmentCreateDto) GetAttachmentRoleOk() (*string, bool)`

GetAttachmentRoleOk returns a tuple with the AttachmentRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentRole

`func (o *SignedDocumentAttachmentCreateDto) SetAttachmentRole(v string)`

SetAttachmentRole sets AttachmentRole field to given value.

### HasAttachmentRole

`func (o *SignedDocumentAttachmentCreateDto) HasAttachmentRole() bool`

HasAttachmentRole returns a boolean if a field has been set.

### SetAttachmentRoleNil

`func (o *SignedDocumentAttachmentCreateDto) SetAttachmentRoleNil(b bool)`

 SetAttachmentRoleNil sets the value for AttachmentRole to be an explicit nil

### UnsetAttachmentRole
`func (o *SignedDocumentAttachmentCreateDto) UnsetAttachmentRole()`

UnsetAttachmentRole ensures that no value is present for AttachmentRole, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


