# SubscriptionCreateDto

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

### NewSubscriptionCreateDto

`func NewSubscriptionCreateDto() *SubscriptionCreateDto`

NewSubscriptionCreateDto instantiates a new SubscriptionCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionCreateDtoWithDefaults

`func NewSubscriptionCreateDtoWithDefaults() *SubscriptionCreateDto`

NewSubscriptionCreateDtoWithDefaults instantiates a new SubscriptionCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *SubscriptionCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SubscriptionCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SubscriptionCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SubscriptionCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetIndividualId

`func (o *SubscriptionCreateDto) GetIndividualId() string`

GetIndividualId returns the IndividualId field if non-nil, zero value otherwise.

### GetIndividualIdOk

`func (o *SubscriptionCreateDto) GetIndividualIdOk() (*string, bool)`

GetIndividualIdOk returns a tuple with the IndividualId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualId

`func (o *SubscriptionCreateDto) SetIndividualId(v string)`

SetIndividualId sets IndividualId field to given value.

### HasIndividualId

`func (o *SubscriptionCreateDto) HasIndividualId() bool`

HasIndividualId returns a boolean if a field has been set.

### SetIndividualIdNil

`func (o *SubscriptionCreateDto) SetIndividualIdNil(b bool)`

 SetIndividualIdNil sets the value for IndividualId to be an explicit nil

### UnsetIndividualId
`func (o *SubscriptionCreateDto) UnsetIndividualId()`

UnsetIndividualId ensures that no value is present for IndividualId, not even an explicit nil
### GetOrganizationId

`func (o *SubscriptionCreateDto) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *SubscriptionCreateDto) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *SubscriptionCreateDto) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *SubscriptionCreateDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *SubscriptionCreateDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *SubscriptionCreateDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetSubscriptionPlanId

`func (o *SubscriptionCreateDto) GetSubscriptionPlanId() string`

GetSubscriptionPlanId returns the SubscriptionPlanId field if non-nil, zero value otherwise.

### GetSubscriptionPlanIdOk

`func (o *SubscriptionCreateDto) GetSubscriptionPlanIdOk() (*string, bool)`

GetSubscriptionPlanIdOk returns a tuple with the SubscriptionPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPlanId

`func (o *SubscriptionCreateDto) SetSubscriptionPlanId(v string)`

SetSubscriptionPlanId sets SubscriptionPlanId field to given value.

### HasSubscriptionPlanId

`func (o *SubscriptionCreateDto) HasSubscriptionPlanId() bool`

HasSubscriptionPlanId returns a boolean if a field has been set.

### SetSubscriptionPlanIdNil

`func (o *SubscriptionCreateDto) SetSubscriptionPlanIdNil(b bool)`

 SetSubscriptionPlanIdNil sets the value for SubscriptionPlanId to be an explicit nil

### UnsetSubscriptionPlanId
`func (o *SubscriptionCreateDto) UnsetSubscriptionPlanId()`

UnsetSubscriptionPlanId ensures that no value is present for SubscriptionPlanId, not even an explicit nil
### GetSubscriptionClass

`func (o *SubscriptionCreateDto) GetSubscriptionClass() string`

GetSubscriptionClass returns the SubscriptionClass field if non-nil, zero value otherwise.

### GetSubscriptionClassOk

`func (o *SubscriptionCreateDto) GetSubscriptionClassOk() (*string, bool)`

GetSubscriptionClassOk returns a tuple with the SubscriptionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionClass

`func (o *SubscriptionCreateDto) SetSubscriptionClass(v string)`

SetSubscriptionClass sets SubscriptionClass field to given value.

### HasSubscriptionClass

`func (o *SubscriptionCreateDto) HasSubscriptionClass() bool`

HasSubscriptionClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


