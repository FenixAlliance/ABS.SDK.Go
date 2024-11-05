# Blob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **int32** |  | [optional] 
**IsFolder** | Pointer to **bool** |  | [optional] [readonly] 
**IsFile** | Pointer to **bool** |  | [optional] [readonly] 
**FolderPath** | Pointer to **NullableString** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] [readonly] 
**Size** | Pointer to **NullableInt64** |  | [optional] 
**Md5** | Pointer to **NullableString** |  | [optional] 
**CreatedTime** | Pointer to **NullableTime** |  | [optional] 
**LastModificationTime** | Pointer to **NullableTime** |  | [optional] 
**FullPath** | Pointer to **NullableString** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Metadata** | Pointer to **map[string]string** |  | [optional] [readonly] 
**IsRootFolder** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewBlob

`func NewBlob() *Blob`

NewBlob instantiates a new Blob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlobWithDefaults

`func NewBlobWithDefaults() *Blob`

NewBlobWithDefaults instantiates a new Blob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *Blob) GetKind() int32`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Blob) GetKindOk() (*int32, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Blob) SetKind(v int32)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Blob) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetIsFolder

`func (o *Blob) GetIsFolder() bool`

GetIsFolder returns the IsFolder field if non-nil, zero value otherwise.

### GetIsFolderOk

`func (o *Blob) GetIsFolderOk() (*bool, bool)`

GetIsFolderOk returns a tuple with the IsFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFolder

`func (o *Blob) SetIsFolder(v bool)`

SetIsFolder sets IsFolder field to given value.

### HasIsFolder

`func (o *Blob) HasIsFolder() bool`

HasIsFolder returns a boolean if a field has been set.

### GetIsFile

`func (o *Blob) GetIsFile() bool`

GetIsFile returns the IsFile field if non-nil, zero value otherwise.

### GetIsFileOk

`func (o *Blob) GetIsFileOk() (*bool, bool)`

GetIsFileOk returns a tuple with the IsFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFile

`func (o *Blob) SetIsFile(v bool)`

SetIsFile sets IsFile field to given value.

### HasIsFile

`func (o *Blob) HasIsFile() bool`

HasIsFile returns a boolean if a field has been set.

### GetFolderPath

`func (o *Blob) GetFolderPath() string`

GetFolderPath returns the FolderPath field if non-nil, zero value otherwise.

### GetFolderPathOk

`func (o *Blob) GetFolderPathOk() (*string, bool)`

GetFolderPathOk returns a tuple with the FolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderPath

`func (o *Blob) SetFolderPath(v string)`

SetFolderPath sets FolderPath field to given value.

### HasFolderPath

`func (o *Blob) HasFolderPath() bool`

HasFolderPath returns a boolean if a field has been set.

### SetFolderPathNil

`func (o *Blob) SetFolderPathNil(b bool)`

 SetFolderPathNil sets the value for FolderPath to be an explicit nil

### UnsetFolderPath
`func (o *Blob) UnsetFolderPath()`

UnsetFolderPath ensures that no value is present for FolderPath, not even an explicit nil
### GetName

`func (o *Blob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Blob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Blob) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Blob) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Blob) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Blob) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSize

`func (o *Blob) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Blob) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Blob) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *Blob) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *Blob) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *Blob) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetMd5

`func (o *Blob) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *Blob) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *Blob) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *Blob) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### SetMd5Nil

`func (o *Blob) SetMd5Nil(b bool)`

 SetMd5Nil sets the value for Md5 to be an explicit nil

### UnsetMd5
`func (o *Blob) UnsetMd5()`

UnsetMd5 ensures that no value is present for Md5, not even an explicit nil
### GetCreatedTime

`func (o *Blob) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Blob) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Blob) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Blob) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### SetCreatedTimeNil

`func (o *Blob) SetCreatedTimeNil(b bool)`

 SetCreatedTimeNil sets the value for CreatedTime to be an explicit nil

### UnsetCreatedTime
`func (o *Blob) UnsetCreatedTime()`

UnsetCreatedTime ensures that no value is present for CreatedTime, not even an explicit nil
### GetLastModificationTime

`func (o *Blob) GetLastModificationTime() time.Time`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *Blob) GetLastModificationTimeOk() (*time.Time, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *Blob) SetLastModificationTime(v time.Time)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *Blob) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.

### SetLastModificationTimeNil

`func (o *Blob) SetLastModificationTimeNil(b bool)`

 SetLastModificationTimeNil sets the value for LastModificationTime to be an explicit nil

### UnsetLastModificationTime
`func (o *Blob) UnsetLastModificationTime()`

UnsetLastModificationTime ensures that no value is present for LastModificationTime, not even an explicit nil
### GetFullPath

`func (o *Blob) GetFullPath() string`

GetFullPath returns the FullPath field if non-nil, zero value otherwise.

### GetFullPathOk

`func (o *Blob) GetFullPathOk() (*string, bool)`

GetFullPathOk returns a tuple with the FullPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPath

`func (o *Blob) SetFullPath(v string)`

SetFullPath sets FullPath field to given value.

### HasFullPath

`func (o *Blob) HasFullPath() bool`

HasFullPath returns a boolean if a field has been set.

### SetFullPathNil

`func (o *Blob) SetFullPathNil(b bool)`

 SetFullPathNil sets the value for FullPath to be an explicit nil

### UnsetFullPath
`func (o *Blob) UnsetFullPath()`

UnsetFullPath ensures that no value is present for FullPath, not even an explicit nil
### GetProperties

`func (o *Blob) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Blob) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Blob) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Blob) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *Blob) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *Blob) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetMetadata

`func (o *Blob) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Blob) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Blob) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Blob) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *Blob) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *Blob) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetIsRootFolder

`func (o *Blob) GetIsRootFolder() bool`

GetIsRootFolder returns the IsRootFolder field if non-nil, zero value otherwise.

### GetIsRootFolderOk

`func (o *Blob) GetIsRootFolderOk() (*bool, bool)`

GetIsRootFolderOk returns a tuple with the IsRootFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRootFolder

`func (o *Blob) SetIsRootFolder(v bool)`

SetIsRootFolder sets IsRootFolder field to given value.

### HasIsRootFolder

`func (o *Blob) HasIsRootFolder() bool`

HasIsRootFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


