# ContactRelationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**ContactId** | Pointer to **NullableString** |  | [optional] 
**ContactName** | Pointer to **NullableString** |  | [optional] 
**RelatedContactId** | Pointer to **NullableString** |  | [optional] 
**RelatedContactName** | Pointer to **NullableString** |  | [optional] 
**ContactRelationTypeId** | Pointer to **NullableString** |  | [optional] 
**ContactRelationTypeName** | Pointer to **NullableString** |  | [optional] 
**QualifiedName** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewContactRelationDto

`func NewContactRelationDto() *ContactRelationDto`

NewContactRelationDto instantiates a new ContactRelationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactRelationDtoWithDefaults

`func NewContactRelationDtoWithDefaults() *ContactRelationDto`

NewContactRelationDtoWithDefaults instantiates a new ContactRelationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactRelationDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactRelationDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactRelationDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ContactRelationDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ContactRelationDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ContactRelationDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ContactRelationDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ContactRelationDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ContactRelationDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ContactRelationDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ContactRelationDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ContactRelationDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetContactId

`func (o *ContactRelationDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *ContactRelationDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *ContactRelationDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *ContactRelationDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *ContactRelationDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *ContactRelationDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetContactName

`func (o *ContactRelationDto) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *ContactRelationDto) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *ContactRelationDto) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *ContactRelationDto) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### SetContactNameNil

`func (o *ContactRelationDto) SetContactNameNil(b bool)`

 SetContactNameNil sets the value for ContactName to be an explicit nil

### UnsetContactName
`func (o *ContactRelationDto) UnsetContactName()`

UnsetContactName ensures that no value is present for ContactName, not even an explicit nil
### GetRelatedContactId

`func (o *ContactRelationDto) GetRelatedContactId() string`

GetRelatedContactId returns the RelatedContactId field if non-nil, zero value otherwise.

### GetRelatedContactIdOk

`func (o *ContactRelationDto) GetRelatedContactIdOk() (*string, bool)`

GetRelatedContactIdOk returns a tuple with the RelatedContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedContactId

`func (o *ContactRelationDto) SetRelatedContactId(v string)`

SetRelatedContactId sets RelatedContactId field to given value.

### HasRelatedContactId

`func (o *ContactRelationDto) HasRelatedContactId() bool`

HasRelatedContactId returns a boolean if a field has been set.

### SetRelatedContactIdNil

`func (o *ContactRelationDto) SetRelatedContactIdNil(b bool)`

 SetRelatedContactIdNil sets the value for RelatedContactId to be an explicit nil

### UnsetRelatedContactId
`func (o *ContactRelationDto) UnsetRelatedContactId()`

UnsetRelatedContactId ensures that no value is present for RelatedContactId, not even an explicit nil
### GetRelatedContactName

`func (o *ContactRelationDto) GetRelatedContactName() string`

GetRelatedContactName returns the RelatedContactName field if non-nil, zero value otherwise.

### GetRelatedContactNameOk

`func (o *ContactRelationDto) GetRelatedContactNameOk() (*string, bool)`

GetRelatedContactNameOk returns a tuple with the RelatedContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedContactName

`func (o *ContactRelationDto) SetRelatedContactName(v string)`

SetRelatedContactName sets RelatedContactName field to given value.

### HasRelatedContactName

`func (o *ContactRelationDto) HasRelatedContactName() bool`

HasRelatedContactName returns a boolean if a field has been set.

### SetRelatedContactNameNil

`func (o *ContactRelationDto) SetRelatedContactNameNil(b bool)`

 SetRelatedContactNameNil sets the value for RelatedContactName to be an explicit nil

### UnsetRelatedContactName
`func (o *ContactRelationDto) UnsetRelatedContactName()`

UnsetRelatedContactName ensures that no value is present for RelatedContactName, not even an explicit nil
### GetContactRelationTypeId

`func (o *ContactRelationDto) GetContactRelationTypeId() string`

GetContactRelationTypeId returns the ContactRelationTypeId field if non-nil, zero value otherwise.

### GetContactRelationTypeIdOk

`func (o *ContactRelationDto) GetContactRelationTypeIdOk() (*string, bool)`

GetContactRelationTypeIdOk returns a tuple with the ContactRelationTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactRelationTypeId

`func (o *ContactRelationDto) SetContactRelationTypeId(v string)`

SetContactRelationTypeId sets ContactRelationTypeId field to given value.

### HasContactRelationTypeId

`func (o *ContactRelationDto) HasContactRelationTypeId() bool`

HasContactRelationTypeId returns a boolean if a field has been set.

### SetContactRelationTypeIdNil

`func (o *ContactRelationDto) SetContactRelationTypeIdNil(b bool)`

 SetContactRelationTypeIdNil sets the value for ContactRelationTypeId to be an explicit nil

### UnsetContactRelationTypeId
`func (o *ContactRelationDto) UnsetContactRelationTypeId()`

UnsetContactRelationTypeId ensures that no value is present for ContactRelationTypeId, not even an explicit nil
### GetContactRelationTypeName

`func (o *ContactRelationDto) GetContactRelationTypeName() string`

GetContactRelationTypeName returns the ContactRelationTypeName field if non-nil, zero value otherwise.

### GetContactRelationTypeNameOk

`func (o *ContactRelationDto) GetContactRelationTypeNameOk() (*string, bool)`

GetContactRelationTypeNameOk returns a tuple with the ContactRelationTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactRelationTypeName

`func (o *ContactRelationDto) SetContactRelationTypeName(v string)`

SetContactRelationTypeName sets ContactRelationTypeName field to given value.

### HasContactRelationTypeName

`func (o *ContactRelationDto) HasContactRelationTypeName() bool`

HasContactRelationTypeName returns a boolean if a field has been set.

### SetContactRelationTypeNameNil

`func (o *ContactRelationDto) SetContactRelationTypeNameNil(b bool)`

 SetContactRelationTypeNameNil sets the value for ContactRelationTypeName to be an explicit nil

### UnsetContactRelationTypeName
`func (o *ContactRelationDto) UnsetContactRelationTypeName()`

UnsetContactRelationTypeName ensures that no value is present for ContactRelationTypeName, not even an explicit nil
### GetQualifiedName

`func (o *ContactRelationDto) GetQualifiedName() string`

GetQualifiedName returns the QualifiedName field if non-nil, zero value otherwise.

### GetQualifiedNameOk

`func (o *ContactRelationDto) GetQualifiedNameOk() (*string, bool)`

GetQualifiedNameOk returns a tuple with the QualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedName

`func (o *ContactRelationDto) SetQualifiedName(v string)`

SetQualifiedName sets QualifiedName field to given value.

### HasQualifiedName

`func (o *ContactRelationDto) HasQualifiedName() bool`

HasQualifiedName returns a boolean if a field has been set.

### SetQualifiedNameNil

`func (o *ContactRelationDto) SetQualifiedNameNil(b bool)`

 SetQualifiedNameNil sets the value for QualifiedName to be an explicit nil

### UnsetQualifiedName
`func (o *ContactRelationDto) UnsetQualifiedName()`

UnsetQualifiedName ensures that no value is present for QualifiedName, not even an explicit nil
### GetTenantId

`func (o *ContactRelationDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ContactRelationDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ContactRelationDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ContactRelationDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ContactRelationDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ContactRelationDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


