/*
StorageService

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.1.4089
Contact: support@fenix-alliance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the FileUploadDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileUploadDto{}

// FileUploadDto struct for FileUploadDto
type FileUploadDto struct {
	Id NullableString `json:"id,omitempty"`
	Timestamp NullableTime `json:"timestamp,omitempty"`
	Notes NullableString `json:"notes,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Author NullableString `json:"author,omitempty"`
	IsFolder *bool `json:"isFolder,omitempty"`
	Hash NullableString `json:"hash,omitempty"`
	FileUrl NullableString `json:"fileUrl,omitempty"`
	FilePath NullableString `json:"filePath,omitempty"`
	FileName NullableString `json:"fileName,omitempty"`
	Abstract NullableString `json:"abstract,omitempty"`
	KeyWords NullableString `json:"keyWords,omitempty"`
	Metadata NullableString `json:"metadata,omitempty"`
	FileLength *int64 `json:"fileLength,omitempty"`
	ContentType NullableString `json:"contentType,omitempty"`
	ParentFileId NullableString `json:"parentFileId,omitempty"`
	ValidResponse *bool `json:"validResponse,omitempty"`
	UserId NullableString `json:"userId,omitempty"`
	TenantId NullableString `json:"tenantId,omitempty"`
	EnrollmentId NullableString `json:"enrollmentId,omitempty"`
	SocialProfileId NullableString `json:"socialProfileId,omitempty"`
	FolderPath NullableString `json:"folderPath,omitempty"`
}

// NewFileUploadDto instantiates a new FileUploadDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileUploadDto() *FileUploadDto {
	this := FileUploadDto{}
	return &this
}

// NewFileUploadDtoWithDefaults instantiates a new FileUploadDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileUploadDtoWithDefaults() *FileUploadDto {
	this := FileUploadDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileUploadDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileUploadDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *FileUploadDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *FileUploadDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *FileUploadDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *FileUploadDto) UnsetId() {
	o.Id.Unset()
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileUploadDto) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileUploadDto) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *FileUploadDto) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableTime and assigns it to the Timestamp field.
func (o *FileUploadDto) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}
// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *FileUploadDto) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *FileUploadDto) UnsetTimestamp() {
	o.Timestamp.Unset()
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileUploadDto) GetNotes() string {
	if o == nil || IsNil(o.Notes.Get()) {
		var ret string
		return ret
	}
	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileUploadDto) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// HasNotes returns a boolean if a field has been set.
func (o *FileUploadDto) HasNotes() bool {
	if o != nil && o.Notes.IsSet() {
		return true
	}

	return false
}

// SetNotes gets a reference to the given NullableString and assigns it to the Notes field.
func (o *FileUploadDto) SetNotes(v string) {
	o.Notes.Set(&v)
}
// SetNotesNil sets the value for Notes to be an explicit nil
func (o *FileUploadDto) SetNotesNil() {
	o.Notes.Set(nil)
}

// UnsetNotes ensures that no value is present for Notes, not even an explicit nil
func (o *FileUploadDto) UnsetNotes() {
	o.Notes.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileUploadDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileUploadDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *FileUploadDto) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *FileUploadDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *FileUploadDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *FileUploadDto) UnsetTitle() {
	o.Title.Unset()
}

// GetAuthor returns the Author field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileUploadDto) GetAuthor() string {
	if o == nil || IsNil(o.Author.Get()) {
		var ret string
		return ret
	}
	return *o.Author.Get()
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileUploadDto) GetAuthorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Author.Get(), o.Author.IsSet()
}

// HasAuthor returns a boolean if a field has been set.
func (o *FileUploadDto) HasAuthor() bool {
	if o != nil && o.Author.IsSet() {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given NullableString and assigns it to the Author field.
func (o *FileUploadDto) SetAuthor(v string) {
	o.Author.Set(&v)
}
// SetAuthorNil sets the value for Author to be an explicit nil
func (o *FileUploadDto) SetAuthorNil() {
	o.Author.Set(nil)
}

// UnsetAuthor ensures that no value is present for Author, not even an explicit nil
func (o *FileUploadDto) UnsetAuthor() {
	o.Author.Unset()
}

// GetIsFolder returns the IsFolder field value if set, zero value otherwise.
func (o *FileUploadDto) GetIsFolder() bool {
	if o == nil || IsNil(o.IsFolder) {
		var ret bool
		return ret
	}
	return *o.IsFolder
}

// GetIsFolderOk returns a tuple with the IsFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileUploadDto) GetIsFolderOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFolder) {
		return nil, false
	}
	return o.IsFolder, true
}

// HasIsFolder returns a boolean if a field has been set.
func (o *FileUploadDto) HasIsFolder() bool {
	if o != nil && !IsNil(o.IsFolder) {
		return true
	}

	return false
}

// SetIsFolder gets a reference to the given bool and assigns it to the IsFolder field.
func (o *FileUploadDto) SetIsFolder(v bool) {
	o.IsFolder = &v
}

// GetHash returns the Hash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileUploadDto) GetHash() string {
	if o == nil || IsNil(o.Hash.Get()) {
		var ret string
		return ret
	}
	return *o.Hash.Get()
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileUploadDto) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hash.Get(), o.Hash.IsSet()
}

// HasHash returns a boolean if a field has been set.
func (o *FileUploadDto) HasHash() bool {
	if o != nil && o.Hash.IsSet() {
		return true
	}

	return false
}

// SetHash gets a reference to the given NullableString and assigns it to the Hash field.
func (o *FileUploadDto) SetHash(v string) {
	o.Hash.Set(&v)
}
// SetHashNil sets the value for Hash to be an explicit nil
func (o *FileUploadDto) SetHashNil() {
	o.Hash.Set(nil)
}

// UnsetHash ensures that no value is present for Hash, not even an explicit nil
func (o *FileUploadDto) UnsetHash() {
	o.Hash.Unset()
}

// GetFileUrl returns the FileUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileUploadDto) GetFileUrl() string {
	if o == nil || IsNil(o.FileUrl.Get()) {
		var ret string
		return ret
	}
	return *o.FileUrl.Get()
}

// GetFileUrlOk returns a tuple with the FileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileUploadDto) GetFileUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileUrl.Get(), o.FileUrl.IsSet()
}

// HasFileUrl returns a boolean if a field has been set.
func (o *FileUploadDto) HasFileUrl() bool {
	if o != nil && o.FileUrl.IsSet() {
		return true
	}

	return false
}

// SetFileUrl gets a reference to the given NullableString and assigns it to the FileUrl field.
func (o *FileUploadDto) SetFileUrl(v string) {
	o.FileUrl.Set(&v)
}
// SetFileUrlNil sets the value for FileUrl to be an explicit nil
func (o *FileUploadDto) SetFileUrlNil() {
	o.FileUrl.Set(nil)
}

// UnsetFileUrl ensures that no value is present for FileUrl, not even an explicit nil
func (o *FileUploadDto) UnsetFileUrl() {
	o.FileUrl.Unset()
}

// GetFilePath returns the FilePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileUploadDto) GetFilePath() string {
	if o == nil || IsNil(o.FilePath.Get()) {
		var ret string
		return ret
	}
	return *o.FilePath.Get()
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileUploadDto) GetFilePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FilePath.Get(), o.FilePath.IsSet()
}

// HasFilePath returns a boolean if a field has been set.
func (o *FileUploadDto) HasFilePath() bool {
	if o != nil && o.FilePath.IsSet() {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given NullableString and assigns it to the FilePath field.
func (o *FileUploadDto) SetFilePath(v string) {
	o.FilePath.Set(&v)
}
// SetFilePathNil sets the value for FilePath to be an explicit nil
func (o *FileUploadDto) SetFilePathNil() {
	o.FilePath.Set(nil)
}

// UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil
func (o *FileUploadDto) UnsetFilePath() {
	o.FilePath.Unset()
}

// GetFileName returns the FileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileUploadDto) GetFileName() string {
	if o == nil || IsNil(o.FileName.Get()) {
		var ret string
		return ret
	}
	return *o.FileName.Get()
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileUploadDto) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileName.Get(), o.FileName.IsSet()
}

// HasFileName returns a boolean if a field has been set.
func (o *FileUploadDto) HasFileName() bool {
	if o != nil && o.FileName.IsSet() {
		return true
	}

	return false
}

// SetFileName gets a reference to the given NullableString and assigns it to the FileName field.
func (o *FileUploadDto) SetFileName(v string) {
	o.FileName.Set(&v)
}
// SetFileNameNil sets the value for FileName to be an explicit nil
func (o *FileUploadDto) SetFileNameNil() {
	o.FileName.Set(nil)
}

// UnsetFileName ensures that no value is present for FileName, not even an explicit nil
func (o *FileUploadDto) UnsetFileName() {
	o.FileName.Unset()
}

// GetAbstract returns the Abstract field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileUploadDto) GetAbstract() string {
	if o == nil || IsNil(o.Abstract.Get()) {
		var ret string
		return ret
	}
	return *o.Abstract.Get()
}

// GetAbstractOk returns a tuple with the Abstract field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileUploadDto) GetAbstractOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Abstract.Get(), o.Abstract.IsSet()
}

// HasAbstract returns a boolean if a field has been set.
func (o *FileUploadDto) HasAbstract() bool {
	if o != nil && o.Abstract.IsSet() {
		return true
	}

	return false
}

// SetAbstract gets a reference to the given NullableString and assigns it to the Abstract field.
func (o *FileUploadDto) SetAbstract(v string) {
	o.Abstract.Set(&v)
}
// SetAbstractNil sets the value for Abstract to be an explicit nil
func (o *FileUploadDto) SetAbstractNil() {
	o.Abstract.Set(nil)
}

// UnsetAbstract ensures that no value is present for Abstract, not even an explicit nil
func (o *FileUploadDto) UnsetAbstract() {
	o.Abstract.Unset()
}

// GetKeyWords returns the KeyWords field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileUploadDto) GetKeyWords() string {
	if o == nil || IsNil(o.KeyWords.Get()) {
		var ret string
		return ret
	}
	return *o.KeyWords.Get()
}

// GetKeyWordsOk returns a tuple with the KeyWords field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileUploadDto) GetKeyWordsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeyWords.Get(), o.KeyWords.IsSet()
}

// HasKeyWords returns a boolean if a field has been set.
func (o *FileUploadDto) HasKeyWords() bool {
	if o != nil && o.KeyWords.IsSet() {
		return true
	}

	return false
}

// SetKeyWords gets a reference to the given NullableString and assigns it to the KeyWords field.
func (o *FileUploadDto) SetKeyWords(v string) {
	o.KeyWords.Set(&v)
}
// SetKeyWordsNil sets the value for KeyWords to be an explicit nil
func (o *FileUploadDto) SetKeyWordsNil() {
	o.KeyWords.Set(nil)
}

// UnsetKeyWords ensures that no value is present for KeyWords, not even an explicit nil
func (o *FileUploadDto) UnsetKeyWords() {
	o.KeyWords.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileUploadDto) GetMetadata() string {
	if o == nil || IsNil(o.Metadata.Get()) {
		var ret string
		return ret
	}
	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileUploadDto) GetMetadataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// HasMetadata returns a boolean if a field has been set.
func (o *FileUploadDto) HasMetadata() bool {
	if o != nil && o.Metadata.IsSet() {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given NullableString and assigns it to the Metadata field.
func (o *FileUploadDto) SetMetadata(v string) {
	o.Metadata.Set(&v)
}
// SetMetadataNil sets the value for Metadata to be an explicit nil
func (o *FileUploadDto) SetMetadataNil() {
	o.Metadata.Set(nil)
}

// UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
func (o *FileUploadDto) UnsetMetadata() {
	o.Metadata.Unset()
}

// GetFileLength returns the FileLength field value if set, zero value otherwise.
func (o *FileUploadDto) GetFileLength() int64 {
	if o == nil || IsNil(o.FileLength) {
		var ret int64
		return ret
	}
	return *o.FileLength
}

// GetFileLengthOk returns a tuple with the FileLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileUploadDto) GetFileLengthOk() (*int64, bool) {
	if o == nil || IsNil(o.FileLength) {
		return nil, false
	}
	return o.FileLength, true
}

// HasFileLength returns a boolean if a field has been set.
func (o *FileUploadDto) HasFileLength() bool {
	if o != nil && !IsNil(o.FileLength) {
		return true
	}

	return false
}

// SetFileLength gets a reference to the given int64 and assigns it to the FileLength field.
func (o *FileUploadDto) SetFileLength(v int64) {
	o.FileLength = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileUploadDto) GetContentType() string {
	if o == nil || IsNil(o.ContentType.Get()) {
		var ret string
		return ret
	}
	return *o.ContentType.Get()
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileUploadDto) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentType.Get(), o.ContentType.IsSet()
}

// HasContentType returns a boolean if a field has been set.
func (o *FileUploadDto) HasContentType() bool {
	if o != nil && o.ContentType.IsSet() {
		return true
	}

	return false
}

// SetContentType gets a reference to the given NullableString and assigns it to the ContentType field.
func (o *FileUploadDto) SetContentType(v string) {
	o.ContentType.Set(&v)
}
// SetContentTypeNil sets the value for ContentType to be an explicit nil
func (o *FileUploadDto) SetContentTypeNil() {
	o.ContentType.Set(nil)
}

// UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
func (o *FileUploadDto) UnsetContentType() {
	o.ContentType.Unset()
}

// GetParentFileId returns the ParentFileId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileUploadDto) GetParentFileId() string {
	if o == nil || IsNil(o.ParentFileId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentFileId.Get()
}

// GetParentFileIdOk returns a tuple with the ParentFileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileUploadDto) GetParentFileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentFileId.Get(), o.ParentFileId.IsSet()
}

// HasParentFileId returns a boolean if a field has been set.
func (o *FileUploadDto) HasParentFileId() bool {
	if o != nil && o.ParentFileId.IsSet() {
		return true
	}

	return false
}

// SetParentFileId gets a reference to the given NullableString and assigns it to the ParentFileId field.
func (o *FileUploadDto) SetParentFileId(v string) {
	o.ParentFileId.Set(&v)
}
// SetParentFileIdNil sets the value for ParentFileId to be an explicit nil
func (o *FileUploadDto) SetParentFileIdNil() {
	o.ParentFileId.Set(nil)
}

// UnsetParentFileId ensures that no value is present for ParentFileId, not even an explicit nil
func (o *FileUploadDto) UnsetParentFileId() {
	o.ParentFileId.Unset()
}

// GetValidResponse returns the ValidResponse field value if set, zero value otherwise.
func (o *FileUploadDto) GetValidResponse() bool {
	if o == nil || IsNil(o.ValidResponse) {
		var ret bool
		return ret
	}
	return *o.ValidResponse
}

// GetValidResponseOk returns a tuple with the ValidResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileUploadDto) GetValidResponseOk() (*bool, bool) {
	if o == nil || IsNil(o.ValidResponse) {
		return nil, false
	}
	return o.ValidResponse, true
}

// HasValidResponse returns a boolean if a field has been set.
func (o *FileUploadDto) HasValidResponse() bool {
	if o != nil && !IsNil(o.ValidResponse) {
		return true
	}

	return false
}

// SetValidResponse gets a reference to the given bool and assigns it to the ValidResponse field.
func (o *FileUploadDto) SetValidResponse(v bool) {
	o.ValidResponse = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileUploadDto) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileUploadDto) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *FileUploadDto) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *FileUploadDto) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *FileUploadDto) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *FileUploadDto) UnsetUserId() {
	o.UserId.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileUploadDto) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileUploadDto) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *FileUploadDto) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *FileUploadDto) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *FileUploadDto) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *FileUploadDto) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetEnrollmentId returns the EnrollmentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileUploadDto) GetEnrollmentId() string {
	if o == nil || IsNil(o.EnrollmentId.Get()) {
		var ret string
		return ret
	}
	return *o.EnrollmentId.Get()
}

// GetEnrollmentIdOk returns a tuple with the EnrollmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileUploadDto) GetEnrollmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnrollmentId.Get(), o.EnrollmentId.IsSet()
}

// HasEnrollmentId returns a boolean if a field has been set.
func (o *FileUploadDto) HasEnrollmentId() bool {
	if o != nil && o.EnrollmentId.IsSet() {
		return true
	}

	return false
}

// SetEnrollmentId gets a reference to the given NullableString and assigns it to the EnrollmentId field.
func (o *FileUploadDto) SetEnrollmentId(v string) {
	o.EnrollmentId.Set(&v)
}
// SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil
func (o *FileUploadDto) SetEnrollmentIdNil() {
	o.EnrollmentId.Set(nil)
}

// UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
func (o *FileUploadDto) UnsetEnrollmentId() {
	o.EnrollmentId.Unset()
}

// GetSocialProfileId returns the SocialProfileId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileUploadDto) GetSocialProfileId() string {
	if o == nil || IsNil(o.SocialProfileId.Get()) {
		var ret string
		return ret
	}
	return *o.SocialProfileId.Get()
}

// GetSocialProfileIdOk returns a tuple with the SocialProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileUploadDto) GetSocialProfileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SocialProfileId.Get(), o.SocialProfileId.IsSet()
}

// HasSocialProfileId returns a boolean if a field has been set.
func (o *FileUploadDto) HasSocialProfileId() bool {
	if o != nil && o.SocialProfileId.IsSet() {
		return true
	}

	return false
}

// SetSocialProfileId gets a reference to the given NullableString and assigns it to the SocialProfileId field.
func (o *FileUploadDto) SetSocialProfileId(v string) {
	o.SocialProfileId.Set(&v)
}
// SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil
func (o *FileUploadDto) SetSocialProfileIdNil() {
	o.SocialProfileId.Set(nil)
}

// UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil
func (o *FileUploadDto) UnsetSocialProfileId() {
	o.SocialProfileId.Unset()
}

// GetFolderPath returns the FolderPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileUploadDto) GetFolderPath() string {
	if o == nil || IsNil(o.FolderPath.Get()) {
		var ret string
		return ret
	}
	return *o.FolderPath.Get()
}

// GetFolderPathOk returns a tuple with the FolderPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileUploadDto) GetFolderPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FolderPath.Get(), o.FolderPath.IsSet()
}

// HasFolderPath returns a boolean if a field has been set.
func (o *FileUploadDto) HasFolderPath() bool {
	if o != nil && o.FolderPath.IsSet() {
		return true
	}

	return false
}

// SetFolderPath gets a reference to the given NullableString and assigns it to the FolderPath field.
func (o *FileUploadDto) SetFolderPath(v string) {
	o.FolderPath.Set(&v)
}
// SetFolderPathNil sets the value for FolderPath to be an explicit nil
func (o *FileUploadDto) SetFolderPathNil() {
	o.FolderPath.Set(nil)
}

// UnsetFolderPath ensures that no value is present for FolderPath, not even an explicit nil
func (o *FileUploadDto) UnsetFolderPath() {
	o.FolderPath.Unset()
}

func (o FileUploadDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileUploadDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	if o.Notes.IsSet() {
		toSerialize["notes"] = o.Notes.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Author.IsSet() {
		toSerialize["author"] = o.Author.Get()
	}
	if !IsNil(o.IsFolder) {
		toSerialize["isFolder"] = o.IsFolder
	}
	if o.Hash.IsSet() {
		toSerialize["hash"] = o.Hash.Get()
	}
	if o.FileUrl.IsSet() {
		toSerialize["fileUrl"] = o.FileUrl.Get()
	}
	if o.FilePath.IsSet() {
		toSerialize["filePath"] = o.FilePath.Get()
	}
	if o.FileName.IsSet() {
		toSerialize["fileName"] = o.FileName.Get()
	}
	if o.Abstract.IsSet() {
		toSerialize["abstract"] = o.Abstract.Get()
	}
	if o.KeyWords.IsSet() {
		toSerialize["keyWords"] = o.KeyWords.Get()
	}
	if o.Metadata.IsSet() {
		toSerialize["metadata"] = o.Metadata.Get()
	}
	if !IsNil(o.FileLength) {
		toSerialize["fileLength"] = o.FileLength
	}
	if o.ContentType.IsSet() {
		toSerialize["contentType"] = o.ContentType.Get()
	}
	if o.ParentFileId.IsSet() {
		toSerialize["parentFileId"] = o.ParentFileId.Get()
	}
	if !IsNil(o.ValidResponse) {
		toSerialize["validResponse"] = o.ValidResponse
	}
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.EnrollmentId.IsSet() {
		toSerialize["enrollmentId"] = o.EnrollmentId.Get()
	}
	if o.SocialProfileId.IsSet() {
		toSerialize["socialProfileId"] = o.SocialProfileId.Get()
	}
	if o.FolderPath.IsSet() {
		toSerialize["folderPath"] = o.FolderPath.Get()
	}
	return toSerialize, nil
}

type NullableFileUploadDto struct {
	value *FileUploadDto
	isSet bool
}

func (v NullableFileUploadDto) Get() *FileUploadDto {
	return v.value
}

func (v *NullableFileUploadDto) Set(val *FileUploadDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFileUploadDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFileUploadDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileUploadDto(val *FileUploadDto) *NullableFileUploadDto {
	return &NullableFileUploadDto{value: val, isSet: true}
}

func (v NullableFileUploadDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileUploadDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


