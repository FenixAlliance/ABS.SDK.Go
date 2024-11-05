# SupportRequestAttachmentCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Author** | Pointer to **NullableString** |  | [optional] 
**IsFolder** | Pointer to **bool** |  | [optional] 
**FileName** | Pointer to **NullableString** |  | [optional] 
**Abstract** | Pointer to **NullableString** |  | [optional] 
**KeyWords** | Pointer to **NullableString** |  | [optional] 
**ValidResponse** | Pointer to **bool** |  | [optional] 
**ParentFileUploadId** | Pointer to **NullableString** |  | [optional] 
**FilePath** | Pointer to **NullableString** |  | [optional] 
**File** | Pointer to **Nullable*os.File** |  | [optional] 
**BusinessID** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordID** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to **NullableString** |  | [optional] 
**SupportRequestID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSupportRequestAttachmentCreateDto

`func NewSupportRequestAttachmentCreateDto() *SupportRequestAttachmentCreateDto`

NewSupportRequestAttachmentCreateDto instantiates a new SupportRequestAttachmentCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportRequestAttachmentCreateDtoWithDefaults

`func NewSupportRequestAttachmentCreateDtoWithDefaults() *SupportRequestAttachmentCreateDto`

NewSupportRequestAttachmentCreateDtoWithDefaults instantiates a new SupportRequestAttachmentCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SupportRequestAttachmentCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportRequestAttachmentCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportRequestAttachmentCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SupportRequestAttachmentCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *SupportRequestAttachmentCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SupportRequestAttachmentCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SupportRequestAttachmentCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SupportRequestAttachmentCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetNotes

`func (o *SupportRequestAttachmentCreateDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *SupportRequestAttachmentCreateDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *SupportRequestAttachmentCreateDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *SupportRequestAttachmentCreateDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *SupportRequestAttachmentCreateDto) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *SupportRequestAttachmentCreateDto) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetTitle

`func (o *SupportRequestAttachmentCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SupportRequestAttachmentCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SupportRequestAttachmentCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SupportRequestAttachmentCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SupportRequestAttachmentCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SupportRequestAttachmentCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAuthor

`func (o *SupportRequestAttachmentCreateDto) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *SupportRequestAttachmentCreateDto) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *SupportRequestAttachmentCreateDto) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *SupportRequestAttachmentCreateDto) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *SupportRequestAttachmentCreateDto) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *SupportRequestAttachmentCreateDto) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetIsFolder

`func (o *SupportRequestAttachmentCreateDto) GetIsFolder() bool`

GetIsFolder returns the IsFolder field if non-nil, zero value otherwise.

### GetIsFolderOk

`func (o *SupportRequestAttachmentCreateDto) GetIsFolderOk() (*bool, bool)`

GetIsFolderOk returns a tuple with the IsFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFolder

`func (o *SupportRequestAttachmentCreateDto) SetIsFolder(v bool)`

SetIsFolder sets IsFolder field to given value.

### HasIsFolder

`func (o *SupportRequestAttachmentCreateDto) HasIsFolder() bool`

HasIsFolder returns a boolean if a field has been set.

### GetFileName

`func (o *SupportRequestAttachmentCreateDto) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *SupportRequestAttachmentCreateDto) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *SupportRequestAttachmentCreateDto) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *SupportRequestAttachmentCreateDto) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *SupportRequestAttachmentCreateDto) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *SupportRequestAttachmentCreateDto) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetAbstract

`func (o *SupportRequestAttachmentCreateDto) GetAbstract() string`

GetAbstract returns the Abstract field if non-nil, zero value otherwise.

### GetAbstractOk

`func (o *SupportRequestAttachmentCreateDto) GetAbstractOk() (*string, bool)`

GetAbstractOk returns a tuple with the Abstract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstract

`func (o *SupportRequestAttachmentCreateDto) SetAbstract(v string)`

SetAbstract sets Abstract field to given value.

### HasAbstract

`func (o *SupportRequestAttachmentCreateDto) HasAbstract() bool`

HasAbstract returns a boolean if a field has been set.

### SetAbstractNil

`func (o *SupportRequestAttachmentCreateDto) SetAbstractNil(b bool)`

 SetAbstractNil sets the value for Abstract to be an explicit nil

### UnsetAbstract
`func (o *SupportRequestAttachmentCreateDto) UnsetAbstract()`

UnsetAbstract ensures that no value is present for Abstract, not even an explicit nil
### GetKeyWords

`func (o *SupportRequestAttachmentCreateDto) GetKeyWords() string`

GetKeyWords returns the KeyWords field if non-nil, zero value otherwise.

### GetKeyWordsOk

`func (o *SupportRequestAttachmentCreateDto) GetKeyWordsOk() (*string, bool)`

GetKeyWordsOk returns a tuple with the KeyWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyWords

`func (o *SupportRequestAttachmentCreateDto) SetKeyWords(v string)`

SetKeyWords sets KeyWords field to given value.

### HasKeyWords

`func (o *SupportRequestAttachmentCreateDto) HasKeyWords() bool`

HasKeyWords returns a boolean if a field has been set.

### SetKeyWordsNil

`func (o *SupportRequestAttachmentCreateDto) SetKeyWordsNil(b bool)`

 SetKeyWordsNil sets the value for KeyWords to be an explicit nil

### UnsetKeyWords
`func (o *SupportRequestAttachmentCreateDto) UnsetKeyWords()`

UnsetKeyWords ensures that no value is present for KeyWords, not even an explicit nil
### GetValidResponse

`func (o *SupportRequestAttachmentCreateDto) GetValidResponse() bool`

GetValidResponse returns the ValidResponse field if non-nil, zero value otherwise.

### GetValidResponseOk

`func (o *SupportRequestAttachmentCreateDto) GetValidResponseOk() (*bool, bool)`

GetValidResponseOk returns a tuple with the ValidResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidResponse

`func (o *SupportRequestAttachmentCreateDto) SetValidResponse(v bool)`

SetValidResponse sets ValidResponse field to given value.

### HasValidResponse

`func (o *SupportRequestAttachmentCreateDto) HasValidResponse() bool`

HasValidResponse returns a boolean if a field has been set.

### GetParentFileUploadId

`func (o *SupportRequestAttachmentCreateDto) GetParentFileUploadId() string`

GetParentFileUploadId returns the ParentFileUploadId field if non-nil, zero value otherwise.

### GetParentFileUploadIdOk

`func (o *SupportRequestAttachmentCreateDto) GetParentFileUploadIdOk() (*string, bool)`

GetParentFileUploadIdOk returns a tuple with the ParentFileUploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFileUploadId

`func (o *SupportRequestAttachmentCreateDto) SetParentFileUploadId(v string)`

SetParentFileUploadId sets ParentFileUploadId field to given value.

### HasParentFileUploadId

`func (o *SupportRequestAttachmentCreateDto) HasParentFileUploadId() bool`

HasParentFileUploadId returns a boolean if a field has been set.

### SetParentFileUploadIdNil

`func (o *SupportRequestAttachmentCreateDto) SetParentFileUploadIdNil(b bool)`

 SetParentFileUploadIdNil sets the value for ParentFileUploadId to be an explicit nil

### UnsetParentFileUploadId
`func (o *SupportRequestAttachmentCreateDto) UnsetParentFileUploadId()`

UnsetParentFileUploadId ensures that no value is present for ParentFileUploadId, not even an explicit nil
### GetFilePath

`func (o *SupportRequestAttachmentCreateDto) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *SupportRequestAttachmentCreateDto) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *SupportRequestAttachmentCreateDto) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *SupportRequestAttachmentCreateDto) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### SetFilePathNil

`func (o *SupportRequestAttachmentCreateDto) SetFilePathNil(b bool)`

 SetFilePathNil sets the value for FilePath to be an explicit nil

### UnsetFilePath
`func (o *SupportRequestAttachmentCreateDto) UnsetFilePath()`

UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil
### GetFile

`func (o *SupportRequestAttachmentCreateDto) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *SupportRequestAttachmentCreateDto) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *SupportRequestAttachmentCreateDto) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *SupportRequestAttachmentCreateDto) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFileNil

`func (o *SupportRequestAttachmentCreateDto) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *SupportRequestAttachmentCreateDto) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil
### GetBusinessID

`func (o *SupportRequestAttachmentCreateDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *SupportRequestAttachmentCreateDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *SupportRequestAttachmentCreateDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *SupportRequestAttachmentCreateDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *SupportRequestAttachmentCreateDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *SupportRequestAttachmentCreateDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil
### GetBusinessProfileRecordID

`func (o *SupportRequestAttachmentCreateDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *SupportRequestAttachmentCreateDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *SupportRequestAttachmentCreateDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *SupportRequestAttachmentCreateDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *SupportRequestAttachmentCreateDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *SupportRequestAttachmentCreateDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil
### GetMetadata

`func (o *SupportRequestAttachmentCreateDto) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SupportRequestAttachmentCreateDto) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SupportRequestAttachmentCreateDto) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SupportRequestAttachmentCreateDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *SupportRequestAttachmentCreateDto) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *SupportRequestAttachmentCreateDto) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetSupportRequestID

`func (o *SupportRequestAttachmentCreateDto) GetSupportRequestID() string`

GetSupportRequestID returns the SupportRequestID field if non-nil, zero value otherwise.

### GetSupportRequestIDOk

`func (o *SupportRequestAttachmentCreateDto) GetSupportRequestIDOk() (*string, bool)`

GetSupportRequestIDOk returns a tuple with the SupportRequestID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportRequestID

`func (o *SupportRequestAttachmentCreateDto) SetSupportRequestID(v string)`

SetSupportRequestID sets SupportRequestID field to given value.

### HasSupportRequestID

`func (o *SupportRequestAttachmentCreateDto) HasSupportRequestID() bool`

HasSupportRequestID returns a boolean if a field has been set.

### SetSupportRequestIDNil

`func (o *SupportRequestAttachmentCreateDto) SetSupportRequestIDNil(b bool)`

 SetSupportRequestIDNil sets the value for SupportRequestID to be an explicit nil

### UnsetSupportRequestID
`func (o *SupportRequestAttachmentCreateDto) UnsetSupportRequestID()`

UnsetSupportRequestID ensures that no value is present for SupportRequestID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


