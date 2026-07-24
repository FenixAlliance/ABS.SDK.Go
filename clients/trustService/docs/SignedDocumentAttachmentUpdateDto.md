# SignedDocumentAttachmentUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**AttachmentRole** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSignedDocumentAttachmentUpdateDto

`func NewSignedDocumentAttachmentUpdateDto(title string, ) *SignedDocumentAttachmentUpdateDto`

NewSignedDocumentAttachmentUpdateDto instantiates a new SignedDocumentAttachmentUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignedDocumentAttachmentUpdateDtoWithDefaults

`func NewSignedDocumentAttachmentUpdateDtoWithDefaults() *SignedDocumentAttachmentUpdateDto`

NewSignedDocumentAttachmentUpdateDtoWithDefaults instantiates a new SignedDocumentAttachmentUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *SignedDocumentAttachmentUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SignedDocumentAttachmentUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SignedDocumentAttachmentUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetAttachmentRole

`func (o *SignedDocumentAttachmentUpdateDto) GetAttachmentRole() string`

GetAttachmentRole returns the AttachmentRole field if non-nil, zero value otherwise.

### GetAttachmentRoleOk

`func (o *SignedDocumentAttachmentUpdateDto) GetAttachmentRoleOk() (*string, bool)`

GetAttachmentRoleOk returns a tuple with the AttachmentRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentRole

`func (o *SignedDocumentAttachmentUpdateDto) SetAttachmentRole(v string)`

SetAttachmentRole sets AttachmentRole field to given value.

### HasAttachmentRole

`func (o *SignedDocumentAttachmentUpdateDto) HasAttachmentRole() bool`

HasAttachmentRole returns a boolean if a field has been set.

### SetAttachmentRoleNil

`func (o *SignedDocumentAttachmentUpdateDto) SetAttachmentRoleNil(b bool)`

 SetAttachmentRoleNil sets the value for AttachmentRole to be an explicit nil

### UnsetAttachmentRole
`func (o *SignedDocumentAttachmentUpdateDto) UnsetAttachmentRole()`

UnsetAttachmentRole ensures that no value is present for AttachmentRole, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


