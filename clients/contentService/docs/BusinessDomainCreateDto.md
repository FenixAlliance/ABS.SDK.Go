# BusinessDomainCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Domain** | **string** |  | 

## Methods

### NewBusinessDomainCreateDto

`func NewBusinessDomainCreateDto(domain string, ) *BusinessDomainCreateDto`

NewBusinessDomainCreateDto instantiates a new BusinessDomainCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessDomainCreateDtoWithDefaults

`func NewBusinessDomainCreateDtoWithDefaults() *BusinessDomainCreateDto`

NewBusinessDomainCreateDtoWithDefaults instantiates a new BusinessDomainCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BusinessDomainCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BusinessDomainCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BusinessDomainCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BusinessDomainCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *BusinessDomainCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BusinessDomainCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BusinessDomainCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BusinessDomainCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetDomain

`func (o *BusinessDomainCreateDto) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *BusinessDomainCreateDto) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *BusinessDomainCreateDto) SetDomain(v string)`

SetDomain sets Domain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


