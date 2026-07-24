# ContactRelationCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**ContactId** | Pointer to **NullableString** |  | [optional] 
**RelatedContactId** | Pointer to **NullableString** |  | [optional] 
**ContactRelationTypeId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewContactRelationCreateDto

`func NewContactRelationCreateDto() *ContactRelationCreateDto`

NewContactRelationCreateDto instantiates a new ContactRelationCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactRelationCreateDtoWithDefaults

`func NewContactRelationCreateDtoWithDefaults() *ContactRelationCreateDto`

NewContactRelationCreateDtoWithDefaults instantiates a new ContactRelationCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactRelationCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactRelationCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactRelationCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ContactRelationCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ContactRelationCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ContactRelationCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ContactRelationCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ContactRelationCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetContactId

`func (o *ContactRelationCreateDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *ContactRelationCreateDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *ContactRelationCreateDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *ContactRelationCreateDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *ContactRelationCreateDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *ContactRelationCreateDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetRelatedContactId

`func (o *ContactRelationCreateDto) GetRelatedContactId() string`

GetRelatedContactId returns the RelatedContactId field if non-nil, zero value otherwise.

### GetRelatedContactIdOk

`func (o *ContactRelationCreateDto) GetRelatedContactIdOk() (*string, bool)`

GetRelatedContactIdOk returns a tuple with the RelatedContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedContactId

`func (o *ContactRelationCreateDto) SetRelatedContactId(v string)`

SetRelatedContactId sets RelatedContactId field to given value.

### HasRelatedContactId

`func (o *ContactRelationCreateDto) HasRelatedContactId() bool`

HasRelatedContactId returns a boolean if a field has been set.

### SetRelatedContactIdNil

`func (o *ContactRelationCreateDto) SetRelatedContactIdNil(b bool)`

 SetRelatedContactIdNil sets the value for RelatedContactId to be an explicit nil

### UnsetRelatedContactId
`func (o *ContactRelationCreateDto) UnsetRelatedContactId()`

UnsetRelatedContactId ensures that no value is present for RelatedContactId, not even an explicit nil
### GetContactRelationTypeId

`func (o *ContactRelationCreateDto) GetContactRelationTypeId() string`

GetContactRelationTypeId returns the ContactRelationTypeId field if non-nil, zero value otherwise.

### GetContactRelationTypeIdOk

`func (o *ContactRelationCreateDto) GetContactRelationTypeIdOk() (*string, bool)`

GetContactRelationTypeIdOk returns a tuple with the ContactRelationTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactRelationTypeId

`func (o *ContactRelationCreateDto) SetContactRelationTypeId(v string)`

SetContactRelationTypeId sets ContactRelationTypeId field to given value.

### HasContactRelationTypeId

`func (o *ContactRelationCreateDto) HasContactRelationTypeId() bool`

HasContactRelationTypeId returns a boolean if a field has been set.

### SetContactRelationTypeIdNil

`func (o *ContactRelationCreateDto) SetContactRelationTypeIdNil(b bool)`

 SetContactRelationTypeIdNil sets the value for ContactRelationTypeId to be an explicit nil

### UnsetContactRelationTypeId
`func (o *ContactRelationCreateDto) UnsetContactRelationTypeId()`

UnsetContactRelationTypeId ensures that no value is present for ContactRelationTypeId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


