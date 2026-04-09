# SubscriptionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**IndividualId** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **NullableString** |  | [optional] 
**SubscriptionPlanId** | Pointer to **NullableString** |  | [optional] 
**SubscriptionClass** | Pointer to **string** |  | [optional] 

## Methods

### NewSubscriptionDto

`func NewSubscriptionDto() *SubscriptionDto`

NewSubscriptionDto instantiates a new SubscriptionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionDtoWithDefaults

`func NewSubscriptionDtoWithDefaults() *SubscriptionDto`

NewSubscriptionDtoWithDefaults instantiates a new SubscriptionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SubscriptionDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SubscriptionDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SubscriptionDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SubscriptionDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SubscriptionDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SubscriptionDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SubscriptionDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SubscriptionDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetIndividualId

`func (o *SubscriptionDto) GetIndividualId() string`

GetIndividualId returns the IndividualId field if non-nil, zero value otherwise.

### GetIndividualIdOk

`func (o *SubscriptionDto) GetIndividualIdOk() (*string, bool)`

GetIndividualIdOk returns a tuple with the IndividualId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualId

`func (o *SubscriptionDto) SetIndividualId(v string)`

SetIndividualId sets IndividualId field to given value.

### HasIndividualId

`func (o *SubscriptionDto) HasIndividualId() bool`

HasIndividualId returns a boolean if a field has been set.

### SetIndividualIdNil

`func (o *SubscriptionDto) SetIndividualIdNil(b bool)`

 SetIndividualIdNil sets the value for IndividualId to be an explicit nil

### UnsetIndividualId
`func (o *SubscriptionDto) UnsetIndividualId()`

UnsetIndividualId ensures that no value is present for IndividualId, not even an explicit nil
### GetOrganizationId

`func (o *SubscriptionDto) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *SubscriptionDto) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *SubscriptionDto) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *SubscriptionDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *SubscriptionDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *SubscriptionDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetSubscriptionPlanId

`func (o *SubscriptionDto) GetSubscriptionPlanId() string`

GetSubscriptionPlanId returns the SubscriptionPlanId field if non-nil, zero value otherwise.

### GetSubscriptionPlanIdOk

`func (o *SubscriptionDto) GetSubscriptionPlanIdOk() (*string, bool)`

GetSubscriptionPlanIdOk returns a tuple with the SubscriptionPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPlanId

`func (o *SubscriptionDto) SetSubscriptionPlanId(v string)`

SetSubscriptionPlanId sets SubscriptionPlanId field to given value.

### HasSubscriptionPlanId

`func (o *SubscriptionDto) HasSubscriptionPlanId() bool`

HasSubscriptionPlanId returns a boolean if a field has been set.

### SetSubscriptionPlanIdNil

`func (o *SubscriptionDto) SetSubscriptionPlanIdNil(b bool)`

 SetSubscriptionPlanIdNil sets the value for SubscriptionPlanId to be an explicit nil

### UnsetSubscriptionPlanId
`func (o *SubscriptionDto) UnsetSubscriptionPlanId()`

UnsetSubscriptionPlanId ensures that no value is present for SubscriptionPlanId, not even an explicit nil
### GetSubscriptionClass

`func (o *SubscriptionDto) GetSubscriptionClass() string`

GetSubscriptionClass returns the SubscriptionClass field if non-nil, zero value otherwise.

### GetSubscriptionClassOk

`func (o *SubscriptionDto) GetSubscriptionClassOk() (*string, bool)`

GetSubscriptionClassOk returns a tuple with the SubscriptionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionClass

`func (o *SubscriptionDto) SetSubscriptionClass(v string)`

SetSubscriptionClass sets SubscriptionClass field to given value.

### HasSubscriptionClass

`func (o *SubscriptionDto) HasSubscriptionClass() bool`

HasSubscriptionClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


