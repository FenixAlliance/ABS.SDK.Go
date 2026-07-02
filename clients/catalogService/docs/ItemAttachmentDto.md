# ItemAttachmentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Author** | Pointer to **NullableString** |  | [optional] 
**IsFolder** | Pointer to **bool** |  | [optional] 
**Hash** | Pointer to **NullableString** |  | [optional] 
**FileUrl** | Pointer to **NullableString** |  | [optional] 
**FilePath** | Pointer to **NullableString** |  | [optional] 
**StorageKey** | Pointer to **NullableString** |  | [optional] 
**StorageProviderKey** | Pointer to **NullableString** |  | [optional] 
**FileName** | Pointer to **NullableString** |  | [optional] 
**Abstract** | Pointer to **NullableString** |  | [optional] 
**KeyWords** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to **NullableString** |  | [optional] 
**FileLength** | Pointer to **int64** |  | [optional] 
**ContentType** | Pointer to **NullableString** |  | [optional] 
**ParentFileId** | Pointer to **NullableString** |  | [optional] 
**ValidResponse** | Pointer to **bool** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**SocialProfileId** | Pointer to **NullableString** |  | [optional] 
**FolderPath** | Pointer to **NullableString** |  | [optional] 
**ScanStatus** | Pointer to **string** |  | [optional] 
**ThumbnailStatus** | Pointer to **string** |  | [optional] 
**HasThumbnail** | Pointer to **bool** |  | [optional] [readonly] 
**ThumbnailStorageKey** | Pointer to **NullableString** |  | [optional] 
**ThumbnailContentType** | Pointer to **NullableString** |  | [optional] 
**ThumbnailWidth** | Pointer to **int32** |  | [optional] 
**ThumbnailHeight** | Pointer to **int32** |  | [optional] 
**PublicAccessType** | Pointer to **string** |  | [optional] 
**ItemId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemAttachmentDto

`func NewItemAttachmentDto() *ItemAttachmentDto`

NewItemAttachmentDto instantiates a new ItemAttachmentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemAttachmentDtoWithDefaults

`func NewItemAttachmentDtoWithDefaults() *ItemAttachmentDto`

NewItemAttachmentDtoWithDefaults instantiates a new ItemAttachmentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemAttachmentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemAttachmentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemAttachmentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemAttachmentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ItemAttachmentDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ItemAttachmentDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ItemAttachmentDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemAttachmentDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemAttachmentDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemAttachmentDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ItemAttachmentDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ItemAttachmentDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetNotes

`func (o *ItemAttachmentDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ItemAttachmentDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ItemAttachmentDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ItemAttachmentDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *ItemAttachmentDto) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *ItemAttachmentDto) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetTitle

`func (o *ItemAttachmentDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ItemAttachmentDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ItemAttachmentDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ItemAttachmentDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ItemAttachmentDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ItemAttachmentDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAuthor

`func (o *ItemAttachmentDto) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ItemAttachmentDto) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ItemAttachmentDto) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ItemAttachmentDto) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *ItemAttachmentDto) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *ItemAttachmentDto) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetIsFolder

`func (o *ItemAttachmentDto) GetIsFolder() bool`

GetIsFolder returns the IsFolder field if non-nil, zero value otherwise.

### GetIsFolderOk

`func (o *ItemAttachmentDto) GetIsFolderOk() (*bool, bool)`

GetIsFolderOk returns a tuple with the IsFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFolder

`func (o *ItemAttachmentDto) SetIsFolder(v bool)`

SetIsFolder sets IsFolder field to given value.

### HasIsFolder

`func (o *ItemAttachmentDto) HasIsFolder() bool`

HasIsFolder returns a boolean if a field has been set.

### GetHash

`func (o *ItemAttachmentDto) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ItemAttachmentDto) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ItemAttachmentDto) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ItemAttachmentDto) HasHash() bool`

HasHash returns a boolean if a field has been set.

### SetHashNil

`func (o *ItemAttachmentDto) SetHashNil(b bool)`

 SetHashNil sets the value for Hash to be an explicit nil

### UnsetHash
`func (o *ItemAttachmentDto) UnsetHash()`

UnsetHash ensures that no value is present for Hash, not even an explicit nil
### GetFileUrl

`func (o *ItemAttachmentDto) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *ItemAttachmentDto) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *ItemAttachmentDto) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *ItemAttachmentDto) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### SetFileUrlNil

`func (o *ItemAttachmentDto) SetFileUrlNil(b bool)`

 SetFileUrlNil sets the value for FileUrl to be an explicit nil

### UnsetFileUrl
`func (o *ItemAttachmentDto) UnsetFileUrl()`

UnsetFileUrl ensures that no value is present for FileUrl, not even an explicit nil
### GetFilePath

`func (o *ItemAttachmentDto) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *ItemAttachmentDto) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *ItemAttachmentDto) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *ItemAttachmentDto) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### SetFilePathNil

`func (o *ItemAttachmentDto) SetFilePathNil(b bool)`

 SetFilePathNil sets the value for FilePath to be an explicit nil

### UnsetFilePath
`func (o *ItemAttachmentDto) UnsetFilePath()`

UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil
### GetStorageKey

`func (o *ItemAttachmentDto) GetStorageKey() string`

GetStorageKey returns the StorageKey field if non-nil, zero value otherwise.

### GetStorageKeyOk

`func (o *ItemAttachmentDto) GetStorageKeyOk() (*string, bool)`

GetStorageKeyOk returns a tuple with the StorageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageKey

`func (o *ItemAttachmentDto) SetStorageKey(v string)`

SetStorageKey sets StorageKey field to given value.

### HasStorageKey

`func (o *ItemAttachmentDto) HasStorageKey() bool`

HasStorageKey returns a boolean if a field has been set.

### SetStorageKeyNil

`func (o *ItemAttachmentDto) SetStorageKeyNil(b bool)`

 SetStorageKeyNil sets the value for StorageKey to be an explicit nil

### UnsetStorageKey
`func (o *ItemAttachmentDto) UnsetStorageKey()`

UnsetStorageKey ensures that no value is present for StorageKey, not even an explicit nil
### GetStorageProviderKey

`func (o *ItemAttachmentDto) GetStorageProviderKey() string`

GetStorageProviderKey returns the StorageProviderKey field if non-nil, zero value otherwise.

### GetStorageProviderKeyOk

`func (o *ItemAttachmentDto) GetStorageProviderKeyOk() (*string, bool)`

GetStorageProviderKeyOk returns a tuple with the StorageProviderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProviderKey

`func (o *ItemAttachmentDto) SetStorageProviderKey(v string)`

SetStorageProviderKey sets StorageProviderKey field to given value.

### HasStorageProviderKey

`func (o *ItemAttachmentDto) HasStorageProviderKey() bool`

HasStorageProviderKey returns a boolean if a field has been set.

### SetStorageProviderKeyNil

`func (o *ItemAttachmentDto) SetStorageProviderKeyNil(b bool)`

 SetStorageProviderKeyNil sets the value for StorageProviderKey to be an explicit nil

### UnsetStorageProviderKey
`func (o *ItemAttachmentDto) UnsetStorageProviderKey()`

UnsetStorageProviderKey ensures that no value is present for StorageProviderKey, not even an explicit nil
### GetFileName

`func (o *ItemAttachmentDto) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ItemAttachmentDto) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ItemAttachmentDto) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ItemAttachmentDto) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *ItemAttachmentDto) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *ItemAttachmentDto) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetAbstract

`func (o *ItemAttachmentDto) GetAbstract() string`

GetAbstract returns the Abstract field if non-nil, zero value otherwise.

### GetAbstractOk

`func (o *ItemAttachmentDto) GetAbstractOk() (*string, bool)`

GetAbstractOk returns a tuple with the Abstract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstract

`func (o *ItemAttachmentDto) SetAbstract(v string)`

SetAbstract sets Abstract field to given value.

### HasAbstract

`func (o *ItemAttachmentDto) HasAbstract() bool`

HasAbstract returns a boolean if a field has been set.

### SetAbstractNil

`func (o *ItemAttachmentDto) SetAbstractNil(b bool)`

 SetAbstractNil sets the value for Abstract to be an explicit nil

### UnsetAbstract
`func (o *ItemAttachmentDto) UnsetAbstract()`

UnsetAbstract ensures that no value is present for Abstract, not even an explicit nil
### GetKeyWords

`func (o *ItemAttachmentDto) GetKeyWords() string`

GetKeyWords returns the KeyWords field if non-nil, zero value otherwise.

### GetKeyWordsOk

`func (o *ItemAttachmentDto) GetKeyWordsOk() (*string, bool)`

GetKeyWordsOk returns a tuple with the KeyWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyWords

`func (o *ItemAttachmentDto) SetKeyWords(v string)`

SetKeyWords sets KeyWords field to given value.

### HasKeyWords

`func (o *ItemAttachmentDto) HasKeyWords() bool`

HasKeyWords returns a boolean if a field has been set.

### SetKeyWordsNil

`func (o *ItemAttachmentDto) SetKeyWordsNil(b bool)`

 SetKeyWordsNil sets the value for KeyWords to be an explicit nil

### UnsetKeyWords
`func (o *ItemAttachmentDto) UnsetKeyWords()`

UnsetKeyWords ensures that no value is present for KeyWords, not even an explicit nil
### GetMetadata

`func (o *ItemAttachmentDto) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ItemAttachmentDto) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ItemAttachmentDto) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ItemAttachmentDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ItemAttachmentDto) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ItemAttachmentDto) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetFileLength

`func (o *ItemAttachmentDto) GetFileLength() int64`

GetFileLength returns the FileLength field if non-nil, zero value otherwise.

### GetFileLengthOk

`func (o *ItemAttachmentDto) GetFileLengthOk() (*int64, bool)`

GetFileLengthOk returns a tuple with the FileLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLength

`func (o *ItemAttachmentDto) SetFileLength(v int64)`

SetFileLength sets FileLength field to given value.

### HasFileLength

`func (o *ItemAttachmentDto) HasFileLength() bool`

HasFileLength returns a boolean if a field has been set.

### GetContentType

`func (o *ItemAttachmentDto) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ItemAttachmentDto) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ItemAttachmentDto) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ItemAttachmentDto) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *ItemAttachmentDto) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *ItemAttachmentDto) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetParentFileId

`func (o *ItemAttachmentDto) GetParentFileId() string`

GetParentFileId returns the ParentFileId field if non-nil, zero value otherwise.

### GetParentFileIdOk

`func (o *ItemAttachmentDto) GetParentFileIdOk() (*string, bool)`

GetParentFileIdOk returns a tuple with the ParentFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFileId

`func (o *ItemAttachmentDto) SetParentFileId(v string)`

SetParentFileId sets ParentFileId field to given value.

### HasParentFileId

`func (o *ItemAttachmentDto) HasParentFileId() bool`

HasParentFileId returns a boolean if a field has been set.

### SetParentFileIdNil

`func (o *ItemAttachmentDto) SetParentFileIdNil(b bool)`

 SetParentFileIdNil sets the value for ParentFileId to be an explicit nil

### UnsetParentFileId
`func (o *ItemAttachmentDto) UnsetParentFileId()`

UnsetParentFileId ensures that no value is present for ParentFileId, not even an explicit nil
### GetValidResponse

`func (o *ItemAttachmentDto) GetValidResponse() bool`

GetValidResponse returns the ValidResponse field if non-nil, zero value otherwise.

### GetValidResponseOk

`func (o *ItemAttachmentDto) GetValidResponseOk() (*bool, bool)`

GetValidResponseOk returns a tuple with the ValidResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidResponse

`func (o *ItemAttachmentDto) SetValidResponse(v bool)`

SetValidResponse sets ValidResponse field to given value.

### HasValidResponse

`func (o *ItemAttachmentDto) HasValidResponse() bool`

HasValidResponse returns a boolean if a field has been set.

### GetUserId

`func (o *ItemAttachmentDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ItemAttachmentDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ItemAttachmentDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ItemAttachmentDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *ItemAttachmentDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ItemAttachmentDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetTenantId

`func (o *ItemAttachmentDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ItemAttachmentDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ItemAttachmentDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ItemAttachmentDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ItemAttachmentDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ItemAttachmentDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *ItemAttachmentDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *ItemAttachmentDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *ItemAttachmentDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *ItemAttachmentDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *ItemAttachmentDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *ItemAttachmentDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetSocialProfileId

`func (o *ItemAttachmentDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *ItemAttachmentDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *ItemAttachmentDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *ItemAttachmentDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *ItemAttachmentDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *ItemAttachmentDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil
### GetFolderPath

`func (o *ItemAttachmentDto) GetFolderPath() string`

GetFolderPath returns the FolderPath field if non-nil, zero value otherwise.

### GetFolderPathOk

`func (o *ItemAttachmentDto) GetFolderPathOk() (*string, bool)`

GetFolderPathOk returns a tuple with the FolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderPath

`func (o *ItemAttachmentDto) SetFolderPath(v string)`

SetFolderPath sets FolderPath field to given value.

### HasFolderPath

`func (o *ItemAttachmentDto) HasFolderPath() bool`

HasFolderPath returns a boolean if a field has been set.

### SetFolderPathNil

`func (o *ItemAttachmentDto) SetFolderPathNil(b bool)`

 SetFolderPathNil sets the value for FolderPath to be an explicit nil

### UnsetFolderPath
`func (o *ItemAttachmentDto) UnsetFolderPath()`

UnsetFolderPath ensures that no value is present for FolderPath, not even an explicit nil
### GetScanStatus

`func (o *ItemAttachmentDto) GetScanStatus() string`

GetScanStatus returns the ScanStatus field if non-nil, zero value otherwise.

### GetScanStatusOk

`func (o *ItemAttachmentDto) GetScanStatusOk() (*string, bool)`

GetScanStatusOk returns a tuple with the ScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanStatus

`func (o *ItemAttachmentDto) SetScanStatus(v string)`

SetScanStatus sets ScanStatus field to given value.

### HasScanStatus

`func (o *ItemAttachmentDto) HasScanStatus() bool`

HasScanStatus returns a boolean if a field has been set.

### GetThumbnailStatus

`func (o *ItemAttachmentDto) GetThumbnailStatus() string`

GetThumbnailStatus returns the ThumbnailStatus field if non-nil, zero value otherwise.

### GetThumbnailStatusOk

`func (o *ItemAttachmentDto) GetThumbnailStatusOk() (*string, bool)`

GetThumbnailStatusOk returns a tuple with the ThumbnailStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailStatus

`func (o *ItemAttachmentDto) SetThumbnailStatus(v string)`

SetThumbnailStatus sets ThumbnailStatus field to given value.

### HasThumbnailStatus

`func (o *ItemAttachmentDto) HasThumbnailStatus() bool`

HasThumbnailStatus returns a boolean if a field has been set.

### GetHasThumbnail

`func (o *ItemAttachmentDto) GetHasThumbnail() bool`

GetHasThumbnail returns the HasThumbnail field if non-nil, zero value otherwise.

### GetHasThumbnailOk

`func (o *ItemAttachmentDto) GetHasThumbnailOk() (*bool, bool)`

GetHasThumbnailOk returns a tuple with the HasThumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasThumbnail

`func (o *ItemAttachmentDto) SetHasThumbnail(v bool)`

SetHasThumbnail sets HasThumbnail field to given value.

### HasHasThumbnail

`func (o *ItemAttachmentDto) HasHasThumbnail() bool`

HasHasThumbnail returns a boolean if a field has been set.

### GetThumbnailStorageKey

`func (o *ItemAttachmentDto) GetThumbnailStorageKey() string`

GetThumbnailStorageKey returns the ThumbnailStorageKey field if non-nil, zero value otherwise.

### GetThumbnailStorageKeyOk

`func (o *ItemAttachmentDto) GetThumbnailStorageKeyOk() (*string, bool)`

GetThumbnailStorageKeyOk returns a tuple with the ThumbnailStorageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailStorageKey

`func (o *ItemAttachmentDto) SetThumbnailStorageKey(v string)`

SetThumbnailStorageKey sets ThumbnailStorageKey field to given value.

### HasThumbnailStorageKey

`func (o *ItemAttachmentDto) HasThumbnailStorageKey() bool`

HasThumbnailStorageKey returns a boolean if a field has been set.

### SetThumbnailStorageKeyNil

`func (o *ItemAttachmentDto) SetThumbnailStorageKeyNil(b bool)`

 SetThumbnailStorageKeyNil sets the value for ThumbnailStorageKey to be an explicit nil

### UnsetThumbnailStorageKey
`func (o *ItemAttachmentDto) UnsetThumbnailStorageKey()`

UnsetThumbnailStorageKey ensures that no value is present for ThumbnailStorageKey, not even an explicit nil
### GetThumbnailContentType

`func (o *ItemAttachmentDto) GetThumbnailContentType() string`

GetThumbnailContentType returns the ThumbnailContentType field if non-nil, zero value otherwise.

### GetThumbnailContentTypeOk

`func (o *ItemAttachmentDto) GetThumbnailContentTypeOk() (*string, bool)`

GetThumbnailContentTypeOk returns a tuple with the ThumbnailContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailContentType

`func (o *ItemAttachmentDto) SetThumbnailContentType(v string)`

SetThumbnailContentType sets ThumbnailContentType field to given value.

### HasThumbnailContentType

`func (o *ItemAttachmentDto) HasThumbnailContentType() bool`

HasThumbnailContentType returns a boolean if a field has been set.

### SetThumbnailContentTypeNil

`func (o *ItemAttachmentDto) SetThumbnailContentTypeNil(b bool)`

 SetThumbnailContentTypeNil sets the value for ThumbnailContentType to be an explicit nil

### UnsetThumbnailContentType
`func (o *ItemAttachmentDto) UnsetThumbnailContentType()`

UnsetThumbnailContentType ensures that no value is present for ThumbnailContentType, not even an explicit nil
### GetThumbnailWidth

`func (o *ItemAttachmentDto) GetThumbnailWidth() int32`

GetThumbnailWidth returns the ThumbnailWidth field if non-nil, zero value otherwise.

### GetThumbnailWidthOk

`func (o *ItemAttachmentDto) GetThumbnailWidthOk() (*int32, bool)`

GetThumbnailWidthOk returns a tuple with the ThumbnailWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailWidth

`func (o *ItemAttachmentDto) SetThumbnailWidth(v int32)`

SetThumbnailWidth sets ThumbnailWidth field to given value.

### HasThumbnailWidth

`func (o *ItemAttachmentDto) HasThumbnailWidth() bool`

HasThumbnailWidth returns a boolean if a field has been set.

### GetThumbnailHeight

`func (o *ItemAttachmentDto) GetThumbnailHeight() int32`

GetThumbnailHeight returns the ThumbnailHeight field if non-nil, zero value otherwise.

### GetThumbnailHeightOk

`func (o *ItemAttachmentDto) GetThumbnailHeightOk() (*int32, bool)`

GetThumbnailHeightOk returns a tuple with the ThumbnailHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailHeight

`func (o *ItemAttachmentDto) SetThumbnailHeight(v int32)`

SetThumbnailHeight sets ThumbnailHeight field to given value.

### HasThumbnailHeight

`func (o *ItemAttachmentDto) HasThumbnailHeight() bool`

HasThumbnailHeight returns a boolean if a field has been set.

### GetPublicAccessType

`func (o *ItemAttachmentDto) GetPublicAccessType() string`

GetPublicAccessType returns the PublicAccessType field if non-nil, zero value otherwise.

### GetPublicAccessTypeOk

`func (o *ItemAttachmentDto) GetPublicAccessTypeOk() (*string, bool)`

GetPublicAccessTypeOk returns a tuple with the PublicAccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAccessType

`func (o *ItemAttachmentDto) SetPublicAccessType(v string)`

SetPublicAccessType sets PublicAccessType field to given value.

### HasPublicAccessType

`func (o *ItemAttachmentDto) HasPublicAccessType() bool`

HasPublicAccessType returns a boolean if a field has been set.

### GetItemId

`func (o *ItemAttachmentDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ItemAttachmentDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ItemAttachmentDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ItemAttachmentDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *ItemAttachmentDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *ItemAttachmentDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


