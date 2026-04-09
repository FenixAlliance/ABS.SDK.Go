# ContactPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactPointType** | Pointer to **NullableString** |  | [optional] 
**Telephone** | Pointer to **NullableString** |  | [optional] 
**ContactType** | Pointer to **NullableString** |  | [optional] 
**ContactOption** | Pointer to **NullableString** |  | [optional] 
**AreaServed** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewContactPoint

`func NewContactPoint() *ContactPoint`

NewContactPoint instantiates a new ContactPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactPointWithDefaults

`func NewContactPointWithDefaults() *ContactPoint`

NewContactPointWithDefaults instantiates a new ContactPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactPointType

`func (o *ContactPoint) GetContactPointType() string`

GetContactPointType returns the ContactPointType field if non-nil, zero value otherwise.

### GetContactPointTypeOk

`func (o *ContactPoint) GetContactPointTypeOk() (*string, bool)`

GetContactPointTypeOk returns a tuple with the ContactPointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPointType

`func (o *ContactPoint) SetContactPointType(v string)`

SetContactPointType sets ContactPointType field to given value.

### HasContactPointType

`func (o *ContactPoint) HasContactPointType() bool`

HasContactPointType returns a boolean if a field has been set.

### SetContactPointTypeNil

`func (o *ContactPoint) SetContactPointTypeNil(b bool)`

 SetContactPointTypeNil sets the value for ContactPointType to be an explicit nil

### UnsetContactPointType
`func (o *ContactPoint) UnsetContactPointType()`

UnsetContactPointType ensures that no value is present for ContactPointType, not even an explicit nil
### GetTelephone

`func (o *ContactPoint) GetTelephone() string`

GetTelephone returns the Telephone field if non-nil, zero value otherwise.

### GetTelephoneOk

`func (o *ContactPoint) GetTelephoneOk() (*string, bool)`

GetTelephoneOk returns a tuple with the Telephone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephone

`func (o *ContactPoint) SetTelephone(v string)`

SetTelephone sets Telephone field to given value.

### HasTelephone

`func (o *ContactPoint) HasTelephone() bool`

HasTelephone returns a boolean if a field has been set.

### SetTelephoneNil

`func (o *ContactPoint) SetTelephoneNil(b bool)`

 SetTelephoneNil sets the value for Telephone to be an explicit nil

### UnsetTelephone
`func (o *ContactPoint) UnsetTelephone()`

UnsetTelephone ensures that no value is present for Telephone, not even an explicit nil
### GetContactType

`func (o *ContactPoint) GetContactType() string`

GetContactType returns the ContactType field if non-nil, zero value otherwise.

### GetContactTypeOk

`func (o *ContactPoint) GetContactTypeOk() (*string, bool)`

GetContactTypeOk returns a tuple with the ContactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactType

`func (o *ContactPoint) SetContactType(v string)`

SetContactType sets ContactType field to given value.

### HasContactType

`func (o *ContactPoint) HasContactType() bool`

HasContactType returns a boolean if a field has been set.

### SetContactTypeNil

`func (o *ContactPoint) SetContactTypeNil(b bool)`

 SetContactTypeNil sets the value for ContactType to be an explicit nil

### UnsetContactType
`func (o *ContactPoint) UnsetContactType()`

UnsetContactType ensures that no value is present for ContactType, not even an explicit nil
### GetContactOption

`func (o *ContactPoint) GetContactOption() string`

GetContactOption returns the ContactOption field if non-nil, zero value otherwise.

### GetContactOptionOk

`func (o *ContactPoint) GetContactOptionOk() (*string, bool)`

GetContactOptionOk returns a tuple with the ContactOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactOption

`func (o *ContactPoint) SetContactOption(v string)`

SetContactOption sets ContactOption field to given value.

### HasContactOption

`func (o *ContactPoint) HasContactOption() bool`

HasContactOption returns a boolean if a field has been set.

### SetContactOptionNil

`func (o *ContactPoint) SetContactOptionNil(b bool)`

 SetContactOptionNil sets the value for ContactOption to be an explicit nil

### UnsetContactOption
`func (o *ContactPoint) UnsetContactOption()`

UnsetContactOption ensures that no value is present for ContactOption, not even an explicit nil
### GetAreaServed

`func (o *ContactPoint) GetAreaServed() []string`

GetAreaServed returns the AreaServed field if non-nil, zero value otherwise.

### GetAreaServedOk

`func (o *ContactPoint) GetAreaServedOk() (*[]string, bool)`

GetAreaServedOk returns a tuple with the AreaServed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaServed

`func (o *ContactPoint) SetAreaServed(v []string)`

SetAreaServed sets AreaServed field to given value.

### HasAreaServed

`func (o *ContactPoint) HasAreaServed() bool`

HasAreaServed returns a boolean if a field has been set.

### SetAreaServedNil

`func (o *ContactPoint) SetAreaServedNil(b bool)`

 SetAreaServedNil sets the value for AreaServed to be an explicit nil

### UnsetAreaServed
`func (o *ContactPoint) UnsetAreaServed()`

UnsetAreaServed ensures that no value is present for AreaServed, not even an explicit nil
### GetType

`func (o *ContactPoint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContactPoint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContactPoint) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ContactPoint) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ContactPoint) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ContactPoint) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


