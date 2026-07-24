# SigningParticipantDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**SigningRequestId** | Pointer to **NullableString** |  | [optional] 
**ContactId** | Pointer to **NullableString** |  | [optional] 
**ContactName** | Pointer to **NullableString** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**RoutingOrder** | Pointer to **int32** |  | [optional] 
**SentAtUtc** | Pointer to **NullableTime** |  | [optional] 
**ViewedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**SignedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**ApprovedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**DeclinedAtUtc** | Pointer to **NullableTime** |  | [optional] 
**DeclineReason** | Pointer to **NullableString** |  | [optional] 
**SignatureId** | Pointer to **NullableString** |  | [optional] 
**AccessTokenExpiresAtUtc** | Pointer to **NullableTime** |  | [optional] 
**CorrelationId** | Pointer to **NullableString** |  | [optional] 
**ExternalReference** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSigningParticipantDto

`func NewSigningParticipantDto() *SigningParticipantDto`

NewSigningParticipantDto instantiates a new SigningParticipantDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSigningParticipantDtoWithDefaults

`func NewSigningParticipantDtoWithDefaults() *SigningParticipantDto`

NewSigningParticipantDtoWithDefaults instantiates a new SigningParticipantDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SigningParticipantDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SigningParticipantDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SigningParticipantDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SigningParticipantDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SigningParticipantDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SigningParticipantDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTenantId

`func (o *SigningParticipantDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SigningParticipantDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SigningParticipantDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SigningParticipantDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *SigningParticipantDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *SigningParticipantDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetSigningRequestId

`func (o *SigningParticipantDto) GetSigningRequestId() string`

GetSigningRequestId returns the SigningRequestId field if non-nil, zero value otherwise.

### GetSigningRequestIdOk

`func (o *SigningParticipantDto) GetSigningRequestIdOk() (*string, bool)`

GetSigningRequestIdOk returns a tuple with the SigningRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningRequestId

`func (o *SigningParticipantDto) SetSigningRequestId(v string)`

SetSigningRequestId sets SigningRequestId field to given value.

### HasSigningRequestId

`func (o *SigningParticipantDto) HasSigningRequestId() bool`

HasSigningRequestId returns a boolean if a field has been set.

### SetSigningRequestIdNil

`func (o *SigningParticipantDto) SetSigningRequestIdNil(b bool)`

 SetSigningRequestIdNil sets the value for SigningRequestId to be an explicit nil

### UnsetSigningRequestId
`func (o *SigningParticipantDto) UnsetSigningRequestId()`

UnsetSigningRequestId ensures that no value is present for SigningRequestId, not even an explicit nil
### GetContactId

`func (o *SigningParticipantDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *SigningParticipantDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *SigningParticipantDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *SigningParticipantDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *SigningParticipantDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *SigningParticipantDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetContactName

`func (o *SigningParticipantDto) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *SigningParticipantDto) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *SigningParticipantDto) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *SigningParticipantDto) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### SetContactNameNil

`func (o *SigningParticipantDto) SetContactNameNil(b bool)`

 SetContactNameNil sets the value for ContactName to be an explicit nil

### UnsetContactName
`func (o *SigningParticipantDto) UnsetContactName()`

UnsetContactName ensures that no value is present for ContactName, not even an explicit nil
### GetRole

`func (o *SigningParticipantDto) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SigningParticipantDto) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SigningParticipantDto) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *SigningParticipantDto) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *SigningParticipantDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SigningParticipantDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SigningParticipantDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SigningParticipantDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRoutingOrder

`func (o *SigningParticipantDto) GetRoutingOrder() int32`

GetRoutingOrder returns the RoutingOrder field if non-nil, zero value otherwise.

### GetRoutingOrderOk

`func (o *SigningParticipantDto) GetRoutingOrderOk() (*int32, bool)`

GetRoutingOrderOk returns a tuple with the RoutingOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingOrder

`func (o *SigningParticipantDto) SetRoutingOrder(v int32)`

SetRoutingOrder sets RoutingOrder field to given value.

### HasRoutingOrder

`func (o *SigningParticipantDto) HasRoutingOrder() bool`

HasRoutingOrder returns a boolean if a field has been set.

### GetSentAtUtc

`func (o *SigningParticipantDto) GetSentAtUtc() time.Time`

GetSentAtUtc returns the SentAtUtc field if non-nil, zero value otherwise.

### GetSentAtUtcOk

`func (o *SigningParticipantDto) GetSentAtUtcOk() (*time.Time, bool)`

GetSentAtUtcOk returns a tuple with the SentAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAtUtc

`func (o *SigningParticipantDto) SetSentAtUtc(v time.Time)`

SetSentAtUtc sets SentAtUtc field to given value.

### HasSentAtUtc

`func (o *SigningParticipantDto) HasSentAtUtc() bool`

HasSentAtUtc returns a boolean if a field has been set.

### SetSentAtUtcNil

`func (o *SigningParticipantDto) SetSentAtUtcNil(b bool)`

 SetSentAtUtcNil sets the value for SentAtUtc to be an explicit nil

### UnsetSentAtUtc
`func (o *SigningParticipantDto) UnsetSentAtUtc()`

UnsetSentAtUtc ensures that no value is present for SentAtUtc, not even an explicit nil
### GetViewedAtUtc

`func (o *SigningParticipantDto) GetViewedAtUtc() time.Time`

GetViewedAtUtc returns the ViewedAtUtc field if non-nil, zero value otherwise.

### GetViewedAtUtcOk

`func (o *SigningParticipantDto) GetViewedAtUtcOk() (*time.Time, bool)`

GetViewedAtUtcOk returns a tuple with the ViewedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewedAtUtc

`func (o *SigningParticipantDto) SetViewedAtUtc(v time.Time)`

SetViewedAtUtc sets ViewedAtUtc field to given value.

### HasViewedAtUtc

`func (o *SigningParticipantDto) HasViewedAtUtc() bool`

HasViewedAtUtc returns a boolean if a field has been set.

### SetViewedAtUtcNil

`func (o *SigningParticipantDto) SetViewedAtUtcNil(b bool)`

 SetViewedAtUtcNil sets the value for ViewedAtUtc to be an explicit nil

### UnsetViewedAtUtc
`func (o *SigningParticipantDto) UnsetViewedAtUtc()`

UnsetViewedAtUtc ensures that no value is present for ViewedAtUtc, not even an explicit nil
### GetSignedAtUtc

`func (o *SigningParticipantDto) GetSignedAtUtc() time.Time`

GetSignedAtUtc returns the SignedAtUtc field if non-nil, zero value otherwise.

### GetSignedAtUtcOk

`func (o *SigningParticipantDto) GetSignedAtUtcOk() (*time.Time, bool)`

GetSignedAtUtcOk returns a tuple with the SignedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedAtUtc

`func (o *SigningParticipantDto) SetSignedAtUtc(v time.Time)`

SetSignedAtUtc sets SignedAtUtc field to given value.

### HasSignedAtUtc

`func (o *SigningParticipantDto) HasSignedAtUtc() bool`

HasSignedAtUtc returns a boolean if a field has been set.

### SetSignedAtUtcNil

`func (o *SigningParticipantDto) SetSignedAtUtcNil(b bool)`

 SetSignedAtUtcNil sets the value for SignedAtUtc to be an explicit nil

### UnsetSignedAtUtc
`func (o *SigningParticipantDto) UnsetSignedAtUtc()`

UnsetSignedAtUtc ensures that no value is present for SignedAtUtc, not even an explicit nil
### GetApprovedAtUtc

`func (o *SigningParticipantDto) GetApprovedAtUtc() time.Time`

GetApprovedAtUtc returns the ApprovedAtUtc field if non-nil, zero value otherwise.

### GetApprovedAtUtcOk

`func (o *SigningParticipantDto) GetApprovedAtUtcOk() (*time.Time, bool)`

GetApprovedAtUtcOk returns a tuple with the ApprovedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAtUtc

`func (o *SigningParticipantDto) SetApprovedAtUtc(v time.Time)`

SetApprovedAtUtc sets ApprovedAtUtc field to given value.

### HasApprovedAtUtc

`func (o *SigningParticipantDto) HasApprovedAtUtc() bool`

HasApprovedAtUtc returns a boolean if a field has been set.

### SetApprovedAtUtcNil

`func (o *SigningParticipantDto) SetApprovedAtUtcNil(b bool)`

 SetApprovedAtUtcNil sets the value for ApprovedAtUtc to be an explicit nil

### UnsetApprovedAtUtc
`func (o *SigningParticipantDto) UnsetApprovedAtUtc()`

UnsetApprovedAtUtc ensures that no value is present for ApprovedAtUtc, not even an explicit nil
### GetDeclinedAtUtc

`func (o *SigningParticipantDto) GetDeclinedAtUtc() time.Time`

GetDeclinedAtUtc returns the DeclinedAtUtc field if non-nil, zero value otherwise.

### GetDeclinedAtUtcOk

`func (o *SigningParticipantDto) GetDeclinedAtUtcOk() (*time.Time, bool)`

GetDeclinedAtUtcOk returns a tuple with the DeclinedAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclinedAtUtc

`func (o *SigningParticipantDto) SetDeclinedAtUtc(v time.Time)`

SetDeclinedAtUtc sets DeclinedAtUtc field to given value.

### HasDeclinedAtUtc

`func (o *SigningParticipantDto) HasDeclinedAtUtc() bool`

HasDeclinedAtUtc returns a boolean if a field has been set.

### SetDeclinedAtUtcNil

`func (o *SigningParticipantDto) SetDeclinedAtUtcNil(b bool)`

 SetDeclinedAtUtcNil sets the value for DeclinedAtUtc to be an explicit nil

### UnsetDeclinedAtUtc
`func (o *SigningParticipantDto) UnsetDeclinedAtUtc()`

UnsetDeclinedAtUtc ensures that no value is present for DeclinedAtUtc, not even an explicit nil
### GetDeclineReason

`func (o *SigningParticipantDto) GetDeclineReason() string`

GetDeclineReason returns the DeclineReason field if non-nil, zero value otherwise.

### GetDeclineReasonOk

`func (o *SigningParticipantDto) GetDeclineReasonOk() (*string, bool)`

GetDeclineReasonOk returns a tuple with the DeclineReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineReason

`func (o *SigningParticipantDto) SetDeclineReason(v string)`

SetDeclineReason sets DeclineReason field to given value.

### HasDeclineReason

`func (o *SigningParticipantDto) HasDeclineReason() bool`

HasDeclineReason returns a boolean if a field has been set.

### SetDeclineReasonNil

`func (o *SigningParticipantDto) SetDeclineReasonNil(b bool)`

 SetDeclineReasonNil sets the value for DeclineReason to be an explicit nil

### UnsetDeclineReason
`func (o *SigningParticipantDto) UnsetDeclineReason()`

UnsetDeclineReason ensures that no value is present for DeclineReason, not even an explicit nil
### GetSignatureId

`func (o *SigningParticipantDto) GetSignatureId() string`

GetSignatureId returns the SignatureId field if non-nil, zero value otherwise.

### GetSignatureIdOk

`func (o *SigningParticipantDto) GetSignatureIdOk() (*string, bool)`

GetSignatureIdOk returns a tuple with the SignatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureId

`func (o *SigningParticipantDto) SetSignatureId(v string)`

SetSignatureId sets SignatureId field to given value.

### HasSignatureId

`func (o *SigningParticipantDto) HasSignatureId() bool`

HasSignatureId returns a boolean if a field has been set.

### SetSignatureIdNil

`func (o *SigningParticipantDto) SetSignatureIdNil(b bool)`

 SetSignatureIdNil sets the value for SignatureId to be an explicit nil

### UnsetSignatureId
`func (o *SigningParticipantDto) UnsetSignatureId()`

UnsetSignatureId ensures that no value is present for SignatureId, not even an explicit nil
### GetAccessTokenExpiresAtUtc

`func (o *SigningParticipantDto) GetAccessTokenExpiresAtUtc() time.Time`

GetAccessTokenExpiresAtUtc returns the AccessTokenExpiresAtUtc field if non-nil, zero value otherwise.

### GetAccessTokenExpiresAtUtcOk

`func (o *SigningParticipantDto) GetAccessTokenExpiresAtUtcOk() (*time.Time, bool)`

GetAccessTokenExpiresAtUtcOk returns a tuple with the AccessTokenExpiresAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenExpiresAtUtc

`func (o *SigningParticipantDto) SetAccessTokenExpiresAtUtc(v time.Time)`

SetAccessTokenExpiresAtUtc sets AccessTokenExpiresAtUtc field to given value.

### HasAccessTokenExpiresAtUtc

`func (o *SigningParticipantDto) HasAccessTokenExpiresAtUtc() bool`

HasAccessTokenExpiresAtUtc returns a boolean if a field has been set.

### SetAccessTokenExpiresAtUtcNil

`func (o *SigningParticipantDto) SetAccessTokenExpiresAtUtcNil(b bool)`

 SetAccessTokenExpiresAtUtcNil sets the value for AccessTokenExpiresAtUtc to be an explicit nil

### UnsetAccessTokenExpiresAtUtc
`func (o *SigningParticipantDto) UnsetAccessTokenExpiresAtUtc()`

UnsetAccessTokenExpiresAtUtc ensures that no value is present for AccessTokenExpiresAtUtc, not even an explicit nil
### GetCorrelationId

`func (o *SigningParticipantDto) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *SigningParticipantDto) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *SigningParticipantDto) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *SigningParticipantDto) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *SigningParticipantDto) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *SigningParticipantDto) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
### GetExternalReference

`func (o *SigningParticipantDto) GetExternalReference() string`

GetExternalReference returns the ExternalReference field if non-nil, zero value otherwise.

### GetExternalReferenceOk

`func (o *SigningParticipantDto) GetExternalReferenceOk() (*string, bool)`

GetExternalReferenceOk returns a tuple with the ExternalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReference

`func (o *SigningParticipantDto) SetExternalReference(v string)`

SetExternalReference sets ExternalReference field to given value.

### HasExternalReference

`func (o *SigningParticipantDto) HasExternalReference() bool`

HasExternalReference returns a boolean if a field has been set.

### SetExternalReferenceNil

`func (o *SigningParticipantDto) SetExternalReferenceNil(b bool)`

 SetExternalReferenceNil sets the value for ExternalReference to be an explicit nil

### UnsetExternalReference
`func (o *SigningParticipantDto) UnsetExternalReference()`

UnsetExternalReference ensures that no value is present for ExternalReference, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


