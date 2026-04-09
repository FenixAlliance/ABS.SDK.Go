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
**ItemID** | Pointer to **NullableString** |  | [optional] 

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
### GetItemID

`func (o *ItemAttachmentDto) GetItemID() string`

GetItemID returns the ItemID field if non-nil, zero value otherwise.

### GetItemIDOk

`func (o *ItemAttachmentDto) GetItemIDOk() (*string, bool)`

GetItemIDOk returns a tuple with the ItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemID

`func (o *ItemAttachmentDto) SetItemID(v string)`

SetItemID sets ItemID field to given value.

### HasItemID

`func (o *ItemAttachmentDto) HasItemID() bool`

HasItemID returns a boolean if a field has been set.

### SetItemIDNil

`func (o *ItemAttachmentDto) SetItemIDNil(b bool)`

 SetItemIDNil sets the value for ItemID to be an explicit nil

### UnsetItemID
`func (o *ItemAttachmentDto) UnsetItemID()`

UnsetItemID ensures that no value is present for ItemID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


