# TenantInvitationCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**UserEmail** | **string** |  | 

## Methods

### NewTenantInvitationCreateDto

`func NewTenantInvitationCreateDto(userEmail string, ) *TenantInvitationCreateDto`

NewTenantInvitationCreateDto instantiates a new TenantInvitationCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantInvitationCreateDtoWithDefaults

`func NewTenantInvitationCreateDtoWithDefaults() *TenantInvitationCreateDto`

NewTenantInvitationCreateDtoWithDefaults instantiates a new TenantInvitationCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantInvitationCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantInvitationCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantInvitationCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantInvitationCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *TenantInvitationCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TenantInvitationCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TenantInvitationCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TenantInvitationCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUserEmail

`func (o *TenantInvitationCreateDto) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *TenantInvitationCreateDto) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *TenantInvitationCreateDto) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


