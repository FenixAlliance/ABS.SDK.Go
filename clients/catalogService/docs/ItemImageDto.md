# ItemImageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**BusinessID** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordID** | Pointer to **NullableString** |  | [optional] 
**ItemID** | Pointer to **NullableString** |  | [optional] 
**IsItemMozaicBG** | Pointer to **bool** |  | [optional] 
**MD5Hash** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to **NullableString** |  | [optional] 
**FileUploadURL** | Pointer to **NullableString** |  | [optional] 
**FileName** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Abstract** | Pointer to **NullableString** |  | [optional] 
**Author** | Pointer to **NullableString** |  | [optional] 
**KeyWords** | Pointer to **NullableString** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**ContentType** | Pointer to **NullableString** |  | [optional] 
**FileLength** | Pointer to **int64** |  | [optional] 
**ValidResponse** | Pointer to **bool** |  | [optional] 
**SocialProfileID** | Pointer to **NullableString** |  | [optional] 
**ParentFileUploadID** | Pointer to **NullableString** |  | [optional] 
**AccountHolderID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemImageDto

`func NewItemImageDto() *ItemImageDto`

NewItemImageDto instantiates a new ItemImageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemImageDtoWithDefaults

`func NewItemImageDtoWithDefaults() *ItemImageDto`

NewItemImageDtoWithDefaults instantiates a new ItemImageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemImageDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemImageDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemImageDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemImageDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ItemImageDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ItemImageDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ItemImageDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemImageDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemImageDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemImageDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ItemImageDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ItemImageDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetBusinessID

`func (o *ItemImageDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *ItemImageDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *ItemImageDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *ItemImageDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *ItemImageDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *ItemImageDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil
### GetBusinessProfileRecordID

`func (o *ItemImageDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *ItemImageDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *ItemImageDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *ItemImageDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *ItemImageDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *ItemImageDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil
### GetItemID

`func (o *ItemImageDto) GetItemID() string`

GetItemID returns the ItemID field if non-nil, zero value otherwise.

### GetItemIDOk

`func (o *ItemImageDto) GetItemIDOk() (*string, bool)`

GetItemIDOk returns a tuple with the ItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemID

`func (o *ItemImageDto) SetItemID(v string)`

SetItemID sets ItemID field to given value.

### HasItemID

`func (o *ItemImageDto) HasItemID() bool`

HasItemID returns a boolean if a field has been set.

### SetItemIDNil

`func (o *ItemImageDto) SetItemIDNil(b bool)`

 SetItemIDNil sets the value for ItemID to be an explicit nil

### UnsetItemID
`func (o *ItemImageDto) UnsetItemID()`

UnsetItemID ensures that no value is present for ItemID, not even an explicit nil
### GetIsItemMozaicBG

`func (o *ItemImageDto) GetIsItemMozaicBG() bool`

GetIsItemMozaicBG returns the IsItemMozaicBG field if non-nil, zero value otherwise.

### GetIsItemMozaicBGOk

`func (o *ItemImageDto) GetIsItemMozaicBGOk() (*bool, bool)`

GetIsItemMozaicBGOk returns a tuple with the IsItemMozaicBG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsItemMozaicBG

`func (o *ItemImageDto) SetIsItemMozaicBG(v bool)`

SetIsItemMozaicBG sets IsItemMozaicBG field to given value.

### HasIsItemMozaicBG

`func (o *ItemImageDto) HasIsItemMozaicBG() bool`

HasIsItemMozaicBG returns a boolean if a field has been set.

### GetMD5Hash

`func (o *ItemImageDto) GetMD5Hash() string`

GetMD5Hash returns the MD5Hash field if non-nil, zero value otherwise.

### GetMD5HashOk

`func (o *ItemImageDto) GetMD5HashOk() (*string, bool)`

GetMD5HashOk returns a tuple with the MD5Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMD5Hash

`func (o *ItemImageDto) SetMD5Hash(v string)`

SetMD5Hash sets MD5Hash field to given value.

### HasMD5Hash

`func (o *ItemImageDto) HasMD5Hash() bool`

HasMD5Hash returns a boolean if a field has been set.

### SetMD5HashNil

`func (o *ItemImageDto) SetMD5HashNil(b bool)`

 SetMD5HashNil sets the value for MD5Hash to be an explicit nil

### UnsetMD5Hash
`func (o *ItemImageDto) UnsetMD5Hash()`

UnsetMD5Hash ensures that no value is present for MD5Hash, not even an explicit nil
### GetMetadata

`func (o *ItemImageDto) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ItemImageDto) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ItemImageDto) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ItemImageDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ItemImageDto) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ItemImageDto) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetFileUploadURL

`func (o *ItemImageDto) GetFileUploadURL() string`

GetFileUploadURL returns the FileUploadURL field if non-nil, zero value otherwise.

### GetFileUploadURLOk

`func (o *ItemImageDto) GetFileUploadURLOk() (*string, bool)`

GetFileUploadURLOk returns a tuple with the FileUploadURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUploadURL

`func (o *ItemImageDto) SetFileUploadURL(v string)`

SetFileUploadURL sets FileUploadURL field to given value.

### HasFileUploadURL

`func (o *ItemImageDto) HasFileUploadURL() bool`

HasFileUploadURL returns a boolean if a field has been set.

### SetFileUploadURLNil

`func (o *ItemImageDto) SetFileUploadURLNil(b bool)`

 SetFileUploadURLNil sets the value for FileUploadURL to be an explicit nil

### UnsetFileUploadURL
`func (o *ItemImageDto) UnsetFileUploadURL()`

UnsetFileUploadURL ensures that no value is present for FileUploadURL, not even an explicit nil
### GetFileName

`func (o *ItemImageDto) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ItemImageDto) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ItemImageDto) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ItemImageDto) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *ItemImageDto) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *ItemImageDto) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetTitle

`func (o *ItemImageDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ItemImageDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ItemImageDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ItemImageDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ItemImageDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ItemImageDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAbstract

`func (o *ItemImageDto) GetAbstract() string`

GetAbstract returns the Abstract field if non-nil, zero value otherwise.

### GetAbstractOk

`func (o *ItemImageDto) GetAbstractOk() (*string, bool)`

GetAbstractOk returns a tuple with the Abstract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstract

`func (o *ItemImageDto) SetAbstract(v string)`

SetAbstract sets Abstract field to given value.

### HasAbstract

`func (o *ItemImageDto) HasAbstract() bool`

HasAbstract returns a boolean if a field has been set.

### SetAbstractNil

`func (o *ItemImageDto) SetAbstractNil(b bool)`

 SetAbstractNil sets the value for Abstract to be an explicit nil

### UnsetAbstract
`func (o *ItemImageDto) UnsetAbstract()`

UnsetAbstract ensures that no value is present for Abstract, not even an explicit nil
### GetAuthor

`func (o *ItemImageDto) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ItemImageDto) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ItemImageDto) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ItemImageDto) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *ItemImageDto) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *ItemImageDto) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetKeyWords

`func (o *ItemImageDto) GetKeyWords() string`

GetKeyWords returns the KeyWords field if non-nil, zero value otherwise.

### GetKeyWordsOk

`func (o *ItemImageDto) GetKeyWordsOk() (*string, bool)`

GetKeyWordsOk returns a tuple with the KeyWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyWords

`func (o *ItemImageDto) SetKeyWords(v string)`

SetKeyWords sets KeyWords field to given value.

### HasKeyWords

`func (o *ItemImageDto) HasKeyWords() bool`

HasKeyWords returns a boolean if a field has been set.

### SetKeyWordsNil

`func (o *ItemImageDto) SetKeyWordsNil(b bool)`

 SetKeyWordsNil sets the value for KeyWords to be an explicit nil

### UnsetKeyWords
`func (o *ItemImageDto) UnsetKeyWords()`

UnsetKeyWords ensures that no value is present for KeyWords, not even an explicit nil
### GetNotes

`func (o *ItemImageDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ItemImageDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ItemImageDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ItemImageDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *ItemImageDto) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *ItemImageDto) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetContentType

`func (o *ItemImageDto) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ItemImageDto) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ItemImageDto) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ItemImageDto) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *ItemImageDto) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *ItemImageDto) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetFileLength

`func (o *ItemImageDto) GetFileLength() int64`

GetFileLength returns the FileLength field if non-nil, zero value otherwise.

### GetFileLengthOk

`func (o *ItemImageDto) GetFileLengthOk() (*int64, bool)`

GetFileLengthOk returns a tuple with the FileLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLength

`func (o *ItemImageDto) SetFileLength(v int64)`

SetFileLength sets FileLength field to given value.

### HasFileLength

`func (o *ItemImageDto) HasFileLength() bool`

HasFileLength returns a boolean if a field has been set.

### GetValidResponse

`func (o *ItemImageDto) GetValidResponse() bool`

GetValidResponse returns the ValidResponse field if non-nil, zero value otherwise.

### GetValidResponseOk

`func (o *ItemImageDto) GetValidResponseOk() (*bool, bool)`

GetValidResponseOk returns a tuple with the ValidResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidResponse

`func (o *ItemImageDto) SetValidResponse(v bool)`

SetValidResponse sets ValidResponse field to given value.

### HasValidResponse

`func (o *ItemImageDto) HasValidResponse() bool`

HasValidResponse returns a boolean if a field has been set.

### GetSocialProfileID

`func (o *ItemImageDto) GetSocialProfileID() string`

GetSocialProfileID returns the SocialProfileID field if non-nil, zero value otherwise.

### GetSocialProfileIDOk

`func (o *ItemImageDto) GetSocialProfileIDOk() (*string, bool)`

GetSocialProfileIDOk returns a tuple with the SocialProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileID

`func (o *ItemImageDto) SetSocialProfileID(v string)`

SetSocialProfileID sets SocialProfileID field to given value.

### HasSocialProfileID

`func (o *ItemImageDto) HasSocialProfileID() bool`

HasSocialProfileID returns a boolean if a field has been set.

### SetSocialProfileIDNil

`func (o *ItemImageDto) SetSocialProfileIDNil(b bool)`

 SetSocialProfileIDNil sets the value for SocialProfileID to be an explicit nil

### UnsetSocialProfileID
`func (o *ItemImageDto) UnsetSocialProfileID()`

UnsetSocialProfileID ensures that no value is present for SocialProfileID, not even an explicit nil
### GetParentFileUploadID

`func (o *ItemImageDto) GetParentFileUploadID() string`

GetParentFileUploadID returns the ParentFileUploadID field if non-nil, zero value otherwise.

### GetParentFileUploadIDOk

`func (o *ItemImageDto) GetParentFileUploadIDOk() (*string, bool)`

GetParentFileUploadIDOk returns a tuple with the ParentFileUploadID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFileUploadID

`func (o *ItemImageDto) SetParentFileUploadID(v string)`

SetParentFileUploadID sets ParentFileUploadID field to given value.

### HasParentFileUploadID

`func (o *ItemImageDto) HasParentFileUploadID() bool`

HasParentFileUploadID returns a boolean if a field has been set.

### SetParentFileUploadIDNil

`func (o *ItemImageDto) SetParentFileUploadIDNil(b bool)`

 SetParentFileUploadIDNil sets the value for ParentFileUploadID to be an explicit nil

### UnsetParentFileUploadID
`func (o *ItemImageDto) UnsetParentFileUploadID()`

UnsetParentFileUploadID ensures that no value is present for ParentFileUploadID, not even an explicit nil
### GetAccountHolderID

`func (o *ItemImageDto) GetAccountHolderID() string`

GetAccountHolderID returns the AccountHolderID field if non-nil, zero value otherwise.

### GetAccountHolderIDOk

`func (o *ItemImageDto) GetAccountHolderIDOk() (*string, bool)`

GetAccountHolderIDOk returns a tuple with the AccountHolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderID

`func (o *ItemImageDto) SetAccountHolderID(v string)`

SetAccountHolderID sets AccountHolderID field to given value.

### HasAccountHolderID

`func (o *ItemImageDto) HasAccountHolderID() bool`

HasAccountHolderID returns a boolean if a field has been set.

### SetAccountHolderIDNil

`func (o *ItemImageDto) SetAccountHolderIDNil(b bool)`

 SetAccountHolderIDNil sets the value for AccountHolderID to be an explicit nil

### UnsetAccountHolderID
`func (o *ItemImageDto) UnsetAccountHolderID()`

UnsetAccountHolderID ensures that no value is present for AccountHolderID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


