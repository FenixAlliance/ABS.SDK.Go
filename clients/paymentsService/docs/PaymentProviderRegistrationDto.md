# PaymentProviderRegistrationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**CreatedAtUtc** | Pointer to **time.Time** |  | [optional] 
**LastModifiedUtc** | Pointer to **NullableTime** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**ProviderCode** | Pointer to **NullableString** |  | [optional] 
**CredentialSetReference** | Pointer to **NullableString** |  | [optional] 
**HasCredential** | Pointer to **bool** |  | [optional] 
**CredentialMode** | Pointer to **string** |  | [optional] 
**ExternalAccountId** | Pointer to **NullableString** |  | [optional] 
**EnabledCapabilities** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewPaymentProviderRegistrationDto

`func NewPaymentProviderRegistrationDto() *PaymentProviderRegistrationDto`

NewPaymentProviderRegistrationDto instantiates a new PaymentProviderRegistrationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentProviderRegistrationDtoWithDefaults

`func NewPaymentProviderRegistrationDtoWithDefaults() *PaymentProviderRegistrationDto`

NewPaymentProviderRegistrationDtoWithDefaults instantiates a new PaymentProviderRegistrationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentProviderRegistrationDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentProviderRegistrationDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentProviderRegistrationDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentProviderRegistrationDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PaymentProviderRegistrationDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PaymentProviderRegistrationDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCreatedAtUtc

`func (o *PaymentProviderRegistrationDto) GetCreatedAtUtc() time.Time`

GetCreatedAtUtc returns the CreatedAtUtc field if non-nil, zero value otherwise.

### GetCreatedAtUtcOk

`func (o *PaymentProviderRegistrationDto) GetCreatedAtUtcOk() (*time.Time, bool)`

GetCreatedAtUtcOk returns a tuple with the CreatedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtUtc

`func (o *PaymentProviderRegistrationDto) SetCreatedAtUtc(v time.Time)`

SetCreatedAtUtc sets CreatedAtUtc field to given value.

### HasCreatedAtUtc

`func (o *PaymentProviderRegistrationDto) HasCreatedAtUtc() bool`

HasCreatedAtUtc returns a boolean if a field has been set.

### GetLastModifiedUtc

`func (o *PaymentProviderRegistrationDto) GetLastModifiedUtc() time.Time`

GetLastModifiedUtc returns the LastModifiedUtc field if non-nil, zero value otherwise.

### GetLastModifiedUtcOk

`func (o *PaymentProviderRegistrationDto) GetLastModifiedUtcOk() (*time.Time, bool)`

GetLastModifiedUtcOk returns a tuple with the LastModifiedUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedUtc

`func (o *PaymentProviderRegistrationDto) SetLastModifiedUtc(v time.Time)`

SetLastModifiedUtc sets LastModifiedUtc field to given value.

### HasLastModifiedUtc

`func (o *PaymentProviderRegistrationDto) HasLastModifiedUtc() bool`

HasLastModifiedUtc returns a boolean if a field has been set.

### SetLastModifiedUtcNil

`func (o *PaymentProviderRegistrationDto) SetLastModifiedUtcNil(b bool)`

 SetLastModifiedUtcNil sets the value for LastModifiedUtc to be an explicit nil

### UnsetLastModifiedUtc
`func (o *PaymentProviderRegistrationDto) UnsetLastModifiedUtc()`

UnsetLastModifiedUtc ensures that no value is present for LastModifiedUtc, not even an explicit nil
### GetTenantId

`func (o *PaymentProviderRegistrationDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PaymentProviderRegistrationDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PaymentProviderRegistrationDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *PaymentProviderRegistrationDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *PaymentProviderRegistrationDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *PaymentProviderRegistrationDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *PaymentProviderRegistrationDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *PaymentProviderRegistrationDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *PaymentProviderRegistrationDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *PaymentProviderRegistrationDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *PaymentProviderRegistrationDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *PaymentProviderRegistrationDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetProviderCode

`func (o *PaymentProviderRegistrationDto) GetProviderCode() string`

GetProviderCode returns the ProviderCode field if non-nil, zero value otherwise.

### GetProviderCodeOk

`func (o *PaymentProviderRegistrationDto) GetProviderCodeOk() (*string, bool)`

GetProviderCodeOk returns a tuple with the ProviderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderCode

`func (o *PaymentProviderRegistrationDto) SetProviderCode(v string)`

SetProviderCode sets ProviderCode field to given value.

### HasProviderCode

`func (o *PaymentProviderRegistrationDto) HasProviderCode() bool`

HasProviderCode returns a boolean if a field has been set.

### SetProviderCodeNil

`func (o *PaymentProviderRegistrationDto) SetProviderCodeNil(b bool)`

 SetProviderCodeNil sets the value for ProviderCode to be an explicit nil

### UnsetProviderCode
`func (o *PaymentProviderRegistrationDto) UnsetProviderCode()`

UnsetProviderCode ensures that no value is present for ProviderCode, not even an explicit nil
### GetCredentialSetReference

`func (o *PaymentProviderRegistrationDto) GetCredentialSetReference() string`

GetCredentialSetReference returns the CredentialSetReference field if non-nil, zero value otherwise.

### GetCredentialSetReferenceOk

`func (o *PaymentProviderRegistrationDto) GetCredentialSetReferenceOk() (*string, bool)`

GetCredentialSetReferenceOk returns a tuple with the CredentialSetReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialSetReference

`func (o *PaymentProviderRegistrationDto) SetCredentialSetReference(v string)`

SetCredentialSetReference sets CredentialSetReference field to given value.

### HasCredentialSetReference

`func (o *PaymentProviderRegistrationDto) HasCredentialSetReference() bool`

HasCredentialSetReference returns a boolean if a field has been set.

### SetCredentialSetReferenceNil

`func (o *PaymentProviderRegistrationDto) SetCredentialSetReferenceNil(b bool)`

 SetCredentialSetReferenceNil sets the value for CredentialSetReference to be an explicit nil

### UnsetCredentialSetReference
`func (o *PaymentProviderRegistrationDto) UnsetCredentialSetReference()`

UnsetCredentialSetReference ensures that no value is present for CredentialSetReference, not even an explicit nil
### GetHasCredential

`func (o *PaymentProviderRegistrationDto) GetHasCredential() bool`

GetHasCredential returns the HasCredential field if non-nil, zero value otherwise.

### GetHasCredentialOk

`func (o *PaymentProviderRegistrationDto) GetHasCredentialOk() (*bool, bool)`

GetHasCredentialOk returns a tuple with the HasCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCredential

`func (o *PaymentProviderRegistrationDto) SetHasCredential(v bool)`

SetHasCredential sets HasCredential field to given value.

### HasHasCredential

`func (o *PaymentProviderRegistrationDto) HasHasCredential() bool`

HasHasCredential returns a boolean if a field has been set.

### GetCredentialMode

`func (o *PaymentProviderRegistrationDto) GetCredentialMode() string`

GetCredentialMode returns the CredentialMode field if non-nil, zero value otherwise.

### GetCredentialModeOk

`func (o *PaymentProviderRegistrationDto) GetCredentialModeOk() (*string, bool)`

GetCredentialModeOk returns a tuple with the CredentialMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialMode

`func (o *PaymentProviderRegistrationDto) SetCredentialMode(v string)`

SetCredentialMode sets CredentialMode field to given value.

### HasCredentialMode

`func (o *PaymentProviderRegistrationDto) HasCredentialMode() bool`

HasCredentialMode returns a boolean if a field has been set.

### GetExternalAccountId

`func (o *PaymentProviderRegistrationDto) GetExternalAccountId() string`

GetExternalAccountId returns the ExternalAccountId field if non-nil, zero value otherwise.

### GetExternalAccountIdOk

`func (o *PaymentProviderRegistrationDto) GetExternalAccountIdOk() (*string, bool)`

GetExternalAccountIdOk returns a tuple with the ExternalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAccountId

`func (o *PaymentProviderRegistrationDto) SetExternalAccountId(v string)`

SetExternalAccountId sets ExternalAccountId field to given value.

### HasExternalAccountId

`func (o *PaymentProviderRegistrationDto) HasExternalAccountId() bool`

HasExternalAccountId returns a boolean if a field has been set.

### SetExternalAccountIdNil

`func (o *PaymentProviderRegistrationDto) SetExternalAccountIdNil(b bool)`

 SetExternalAccountIdNil sets the value for ExternalAccountId to be an explicit nil

### UnsetExternalAccountId
`func (o *PaymentProviderRegistrationDto) UnsetExternalAccountId()`

UnsetExternalAccountId ensures that no value is present for ExternalAccountId, not even an explicit nil
### GetEnabledCapabilities

`func (o *PaymentProviderRegistrationDto) GetEnabledCapabilities() string`

GetEnabledCapabilities returns the EnabledCapabilities field if non-nil, zero value otherwise.

### GetEnabledCapabilitiesOk

`func (o *PaymentProviderRegistrationDto) GetEnabledCapabilitiesOk() (*string, bool)`

GetEnabledCapabilitiesOk returns a tuple with the EnabledCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledCapabilities

`func (o *PaymentProviderRegistrationDto) SetEnabledCapabilities(v string)`

SetEnabledCapabilities sets EnabledCapabilities field to given value.

### HasEnabledCapabilities

`func (o *PaymentProviderRegistrationDto) HasEnabledCapabilities() bool`

HasEnabledCapabilities returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentProviderRegistrationDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentProviderRegistrationDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentProviderRegistrationDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentProviderRegistrationDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


