# SignatureVerificationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsValid** | Pointer to **bool** |  | [optional] 
**Method** | Pointer to **NullableString** |  | [optional] 
**SignerSubject** | Pointer to **NullableString** |  | [optional] 
**SignerThumbprint** | Pointer to **NullableString** |  | [optional] 
**Issues** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSignatureVerificationDto

`func NewSignatureVerificationDto() *SignatureVerificationDto`

NewSignatureVerificationDto instantiates a new SignatureVerificationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureVerificationDtoWithDefaults

`func NewSignatureVerificationDtoWithDefaults() *SignatureVerificationDto`

NewSignatureVerificationDtoWithDefaults instantiates a new SignatureVerificationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsValid

`func (o *SignatureVerificationDto) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *SignatureVerificationDto) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *SignatureVerificationDto) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *SignatureVerificationDto) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### GetMethod

`func (o *SignatureVerificationDto) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *SignatureVerificationDto) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *SignatureVerificationDto) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *SignatureVerificationDto) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### SetMethodNil

`func (o *SignatureVerificationDto) SetMethodNil(b bool)`

 SetMethodNil sets the value for Method to be an explicit nil

### UnsetMethod
`func (o *SignatureVerificationDto) UnsetMethod()`

UnsetMethod ensures that no value is present for Method, not even an explicit nil
### GetSignerSubject

`func (o *SignatureVerificationDto) GetSignerSubject() string`

GetSignerSubject returns the SignerSubject field if non-nil, zero value otherwise.

### GetSignerSubjectOk

`func (o *SignatureVerificationDto) GetSignerSubjectOk() (*string, bool)`

GetSignerSubjectOk returns a tuple with the SignerSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerSubject

`func (o *SignatureVerificationDto) SetSignerSubject(v string)`

SetSignerSubject sets SignerSubject field to given value.

### HasSignerSubject

`func (o *SignatureVerificationDto) HasSignerSubject() bool`

HasSignerSubject returns a boolean if a field has been set.

### SetSignerSubjectNil

`func (o *SignatureVerificationDto) SetSignerSubjectNil(b bool)`

 SetSignerSubjectNil sets the value for SignerSubject to be an explicit nil

### UnsetSignerSubject
`func (o *SignatureVerificationDto) UnsetSignerSubject()`

UnsetSignerSubject ensures that no value is present for SignerSubject, not even an explicit nil
### GetSignerThumbprint

`func (o *SignatureVerificationDto) GetSignerThumbprint() string`

GetSignerThumbprint returns the SignerThumbprint field if non-nil, zero value otherwise.

### GetSignerThumbprintOk

`func (o *SignatureVerificationDto) GetSignerThumbprintOk() (*string, bool)`

GetSignerThumbprintOk returns a tuple with the SignerThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerThumbprint

`func (o *SignatureVerificationDto) SetSignerThumbprint(v string)`

SetSignerThumbprint sets SignerThumbprint field to given value.

### HasSignerThumbprint

`func (o *SignatureVerificationDto) HasSignerThumbprint() bool`

HasSignerThumbprint returns a boolean if a field has been set.

### SetSignerThumbprintNil

`func (o *SignatureVerificationDto) SetSignerThumbprintNil(b bool)`

 SetSignerThumbprintNil sets the value for SignerThumbprint to be an explicit nil

### UnsetSignerThumbprint
`func (o *SignatureVerificationDto) UnsetSignerThumbprint()`

UnsetSignerThumbprint ensures that no value is present for SignerThumbprint, not even an explicit nil
### GetIssues

`func (o *SignatureVerificationDto) GetIssues() []string`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *SignatureVerificationDto) GetIssuesOk() (*[]string, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *SignatureVerificationDto) SetIssues(v []string)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *SignatureVerificationDto) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### SetIssuesNil

`func (o *SignatureVerificationDto) SetIssuesNil(b bool)`

 SetIssuesNil sets the value for Issues to be an explicit nil

### UnsetIssues
`func (o *SignatureVerificationDto) UnsetIssues()`

UnsetIssues ensures that no value is present for Issues, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


