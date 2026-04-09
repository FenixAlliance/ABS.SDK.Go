# SubscriptionUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**IndividualId** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **NullableString** |  | [optional] 
**SubscriptionPlanId** | Pointer to **NullableString** |  | [optional] 
**SubscriptionClass** | Pointer to **string** |  | [optional] 

## Methods

### NewSubscriptionUpdateDto

`func NewSubscriptionUpdateDto() *SubscriptionUpdateDto`

NewSubscriptionUpdateDto instantiates a new SubscriptionUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionUpdateDtoWithDefaults

`func NewSubscriptionUpdateDtoWithDefaults() *SubscriptionUpdateDto`

NewSubscriptionUpdateDtoWithDefaults instantiates a new SubscriptionUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionUpdateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionUpdateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionUpdateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionUpdateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *SubscriptionUpdateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SubscriptionUpdateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SubscriptionUpdateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SubscriptionUpdateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetIndividualId

`func (o *SubscriptionUpdateDto) GetIndividualId() string`

GetIndividualId returns the IndividualId field if non-nil, zero value otherwise.

### GetIndividualIdOk

`func (o *SubscriptionUpdateDto) GetIndividualIdOk() (*string, bool)`

GetIndividualIdOk returns a tuple with the IndividualId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualId

`func (o *SubscriptionUpdateDto) SetIndividualId(v string)`

SetIndividualId sets IndividualId field to given value.

### HasIndividualId

`func (o *SubscriptionUpdateDto) HasIndividualId() bool`

HasIndividualId returns a boolean if a field has been set.

### SetIndividualIdNil

`func (o *SubscriptionUpdateDto) SetIndividualIdNil(b bool)`

 SetIndividualIdNil sets the value for IndividualId to be an explicit nil

### UnsetIndividualId
`func (o *SubscriptionUpdateDto) UnsetIndividualId()`

UnsetIndividualId ensures that no value is present for IndividualId, not even an explicit nil
### GetOrganizationId

`func (o *SubscriptionUpdateDto) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *SubscriptionUpdateDto) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *SubscriptionUpdateDto) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *SubscriptionUpdateDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *SubscriptionUpdateDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *SubscriptionUpdateDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetSubscriptionPlanId

`func (o *SubscriptionUpdateDto) GetSubscriptionPlanId() string`

GetSubscriptionPlanId returns the SubscriptionPlanId field if non-nil, zero value otherwise.

### GetSubscriptionPlanIdOk

`func (o *SubscriptionUpdateDto) GetSubscriptionPlanIdOk() (*string, bool)`

GetSubscriptionPlanIdOk returns a tuple with the SubscriptionPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPlanId

`func (o *SubscriptionUpdateDto) SetSubscriptionPlanId(v string)`

SetSubscriptionPlanId sets SubscriptionPlanId field to given value.

### HasSubscriptionPlanId

`func (o *SubscriptionUpdateDto) HasSubscriptionPlanId() bool`

HasSubscriptionPlanId returns a boolean if a field has been set.

### SetSubscriptionPlanIdNil

`func (o *SubscriptionUpdateDto) SetSubscriptionPlanIdNil(b bool)`

 SetSubscriptionPlanIdNil sets the value for SubscriptionPlanId to be an explicit nil

### UnsetSubscriptionPlanId
`func (o *SubscriptionUpdateDto) UnsetSubscriptionPlanId()`

UnsetSubscriptionPlanId ensures that no value is present for SubscriptionPlanId, not even an explicit nil
### GetSubscriptionClass

`func (o *SubscriptionUpdateDto) GetSubscriptionClass() string`

GetSubscriptionClass returns the SubscriptionClass field if non-nil, zero value otherwise.

### GetSubscriptionClassOk

`func (o *SubscriptionUpdateDto) GetSubscriptionClassOk() (*string, bool)`

GetSubscriptionClassOk returns a tuple with the SubscriptionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionClass

`func (o *SubscriptionUpdateDto) SetSubscriptionClass(v string)`

SetSubscriptionClass sets SubscriptionClass field to given value.

### HasSubscriptionClass

`func (o *SubscriptionUpdateDto) HasSubscriptionClass() bool`

HasSubscriptionClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


