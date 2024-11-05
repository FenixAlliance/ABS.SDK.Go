# SupportRequestAttachmentUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notes** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Author** | Pointer to **NullableString** |  | [optional] 
**IsFolder** | Pointer to **bool** |  | [optional] 
**FileName** | Pointer to **NullableString** |  | [optional] 
**Abstract** | Pointer to **NullableString** |  | [optional] 
**KeyWords** | Pointer to **NullableString** |  | [optional] 
**ValidResponse** | Pointer to **bool** |  | [optional] 
**ParentFileUploadID** | Pointer to **NullableString** |  | [optional] 
**FilePath** | Pointer to **NullableString** |  | [optional] 
**File** | Pointer to **Nullable*os.File** |  | [optional] 
**ContentType** | Pointer to **NullableString** |  | [optional] 
**FileLength** | Pointer to **int64** |  | [optional] 

## Methods

### NewSupportRequestAttachmentUpdateDto

`func NewSupportRequestAttachmentUpdateDto() *SupportRequestAttachmentUpdateDto`

NewSupportRequestAttachmentUpdateDto instantiates a new SupportRequestAttachmentUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportRequestAttachmentUpdateDtoWithDefaults

`func NewSupportRequestAttachmentUpdateDtoWithDefaults() *SupportRequestAttachmentUpdateDto`

NewSupportRequestAttachmentUpdateDtoWithDefaults instantiates a new SupportRequestAttachmentUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotes

`func (o *SupportRequestAttachmentUpdateDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *SupportRequestAttachmentUpdateDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *SupportRequestAttachmentUpdateDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *SupportRequestAttachmentUpdateDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *SupportRequestAttachmentUpdateDto) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *SupportRequestAttachmentUpdateDto) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetMetadata

`func (o *SupportRequestAttachmentUpdateDto) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SupportRequestAttachmentUpdateDto) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SupportRequestAttachmentUpdateDto) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SupportRequestAttachmentUpdateDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *SupportRequestAttachmentUpdateDto) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *SupportRequestAttachmentUpdateDto) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetTitle

`func (o *SupportRequestAttachmentUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SupportRequestAttachmentUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SupportRequestAttachmentUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SupportRequestAttachmentUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SupportRequestAttachmentUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SupportRequestAttachmentUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAuthor

`func (o *SupportRequestAttachmentUpdateDto) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *SupportRequestAttachmentUpdateDto) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *SupportRequestAttachmentUpdateDto) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *SupportRequestAttachmentUpdateDto) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *SupportRequestAttachmentUpdateDto) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *SupportRequestAttachmentUpdateDto) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetIsFolder

`func (o *SupportRequestAttachmentUpdateDto) GetIsFolder() bool`

GetIsFolder returns the IsFolder field if non-nil, zero value otherwise.

### GetIsFolderOk

`func (o *SupportRequestAttachmentUpdateDto) GetIsFolderOk() (*bool, bool)`

GetIsFolderOk returns a tuple with the IsFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFolder

`func (o *SupportRequestAttachmentUpdateDto) SetIsFolder(v bool)`

SetIsFolder sets IsFolder field to given value.

### HasIsFolder

`func (o *SupportRequestAttachmentUpdateDto) HasIsFolder() bool`

HasIsFolder returns a boolean if a field has been set.

### GetFileName

`func (o *SupportRequestAttachmentUpdateDto) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *SupportRequestAttachmentUpdateDto) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *SupportRequestAttachmentUpdateDto) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *SupportRequestAttachmentUpdateDto) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *SupportRequestAttachmentUpdateDto) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *SupportRequestAttachmentUpdateDto) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetAbstract

`func (o *SupportRequestAttachmentUpdateDto) GetAbstract() string`

GetAbstract returns the Abstract field if non-nil, zero value otherwise.

### GetAbstractOk

`func (o *SupportRequestAttachmentUpdateDto) GetAbstractOk() (*string, bool)`

GetAbstractOk returns a tuple with the Abstract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstract

`func (o *SupportRequestAttachmentUpdateDto) SetAbstract(v string)`

SetAbstract sets Abstract field to given value.

### HasAbstract

`func (o *SupportRequestAttachmentUpdateDto) HasAbstract() bool`

HasAbstract returns a boolean if a field has been set.

### SetAbstractNil

`func (o *SupportRequestAttachmentUpdateDto) SetAbstractNil(b bool)`

 SetAbstractNil sets the value for Abstract to be an explicit nil

### UnsetAbstract
`func (o *SupportRequestAttachmentUpdateDto) UnsetAbstract()`

UnsetAbstract ensures that no value is present for Abstract, not even an explicit nil
### GetKeyWords

`func (o *SupportRequestAttachmentUpdateDto) GetKeyWords() string`

GetKeyWords returns the KeyWords field if non-nil, zero value otherwise.

### GetKeyWordsOk

`func (o *SupportRequestAttachmentUpdateDto) GetKeyWordsOk() (*string, bool)`

GetKeyWordsOk returns a tuple with the KeyWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyWords

`func (o *SupportRequestAttachmentUpdateDto) SetKeyWords(v string)`

SetKeyWords sets KeyWords field to given value.

### HasKeyWords

`func (o *SupportRequestAttachmentUpdateDto) HasKeyWords() bool`

HasKeyWords returns a boolean if a field has been set.

### SetKeyWordsNil

`func (o *SupportRequestAttachmentUpdateDto) SetKeyWordsNil(b bool)`

 SetKeyWordsNil sets the value for KeyWords to be an explicit nil

### UnsetKeyWords
`func (o *SupportRequestAttachmentUpdateDto) UnsetKeyWords()`

UnsetKeyWords ensures that no value is present for KeyWords, not even an explicit nil
### GetValidResponse

`func (o *SupportRequestAttachmentUpdateDto) GetValidResponse() bool`

GetValidResponse returns the ValidResponse field if non-nil, zero value otherwise.

### GetValidResponseOk

`func (o *SupportRequestAttachmentUpdateDto) GetValidResponseOk() (*bool, bool)`

GetValidResponseOk returns a tuple with the ValidResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidResponse

`func (o *SupportRequestAttachmentUpdateDto) SetValidResponse(v bool)`

SetValidResponse sets ValidResponse field to given value.

### HasValidResponse

`func (o *SupportRequestAttachmentUpdateDto) HasValidResponse() bool`

HasValidResponse returns a boolean if a field has been set.

### GetParentFileUploadID

`func (o *SupportRequestAttachmentUpdateDto) GetParentFileUploadID() string`

GetParentFileUploadID returns the ParentFileUploadID field if non-nil, zero value otherwise.

### GetParentFileUploadIDOk

`func (o *SupportRequestAttachmentUpdateDto) GetParentFileUploadIDOk() (*string, bool)`

GetParentFileUploadIDOk returns a tuple with the ParentFileUploadID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFileUploadID

`func (o *SupportRequestAttachmentUpdateDto) SetParentFileUploadID(v string)`

SetParentFileUploadID sets ParentFileUploadID field to given value.

### HasParentFileUploadID

`func (o *SupportRequestAttachmentUpdateDto) HasParentFileUploadID() bool`

HasParentFileUploadID returns a boolean if a field has been set.

### SetParentFileUploadIDNil

`func (o *SupportRequestAttachmentUpdateDto) SetParentFileUploadIDNil(b bool)`

 SetParentFileUploadIDNil sets the value for ParentFileUploadID to be an explicit nil

### UnsetParentFileUploadID
`func (o *SupportRequestAttachmentUpdateDto) UnsetParentFileUploadID()`

UnsetParentFileUploadID ensures that no value is present for ParentFileUploadID, not even an explicit nil
### GetFilePath

`func (o *SupportRequestAttachmentUpdateDto) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *SupportRequestAttachmentUpdateDto) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *SupportRequestAttachmentUpdateDto) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *SupportRequestAttachmentUpdateDto) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### SetFilePathNil

`func (o *SupportRequestAttachmentUpdateDto) SetFilePathNil(b bool)`

 SetFilePathNil sets the value for FilePath to be an explicit nil

### UnsetFilePath
`func (o *SupportRequestAttachmentUpdateDto) UnsetFilePath()`

UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil
### GetFile

`func (o *SupportRequestAttachmentUpdateDto) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *SupportRequestAttachmentUpdateDto) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *SupportRequestAttachmentUpdateDto) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *SupportRequestAttachmentUpdateDto) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFileNil

`func (o *SupportRequestAttachmentUpdateDto) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *SupportRequestAttachmentUpdateDto) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil
### GetContentType

`func (o *SupportRequestAttachmentUpdateDto) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *SupportRequestAttachmentUpdateDto) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *SupportRequestAttachmentUpdateDto) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *SupportRequestAttachmentUpdateDto) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *SupportRequestAttachmentUpdateDto) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *SupportRequestAttachmentUpdateDto) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetFileLength

`func (o *SupportRequestAttachmentUpdateDto) GetFileLength() int64`

GetFileLength returns the FileLength field if non-nil, zero value otherwise.

### GetFileLengthOk

`func (o *SupportRequestAttachmentUpdateDto) GetFileLengthOk() (*int64, bool)`

GetFileLengthOk returns a tuple with the FileLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLength

`func (o *SupportRequestAttachmentUpdateDto) SetFileLength(v int64)`

SetFileLength sets FileLength field to given value.

### HasFileLength

`func (o *SupportRequestAttachmentUpdateDto) HasFileLength() bool`

HasFileLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


