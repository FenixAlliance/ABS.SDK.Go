# InquiryRequestCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**LastName** | Pointer to **NullableString** |  | [optional] 
**Email** | **string** |  | 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**JobRole** | Pointer to **NullableString** |  | [optional] 
**OrganizationDomain** | Pointer to **NullableString** |  | [optional] 
**CountryId** | Pointer to **NullableString** |  | [optional] 
**Phone** | Pointer to **NullableString** |  | [optional] 
**Message** | **string** |  | 
**SocialProfileId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInquiryRequestCreateDto

`func NewInquiryRequestCreateDto(name string, email string, message string, ) *InquiryRequestCreateDto`

NewInquiryRequestCreateDto instantiates a new InquiryRequestCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInquiryRequestCreateDtoWithDefaults

`func NewInquiryRequestCreateDtoWithDefaults() *InquiryRequestCreateDto`

NewInquiryRequestCreateDtoWithDefaults instantiates a new InquiryRequestCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InquiryRequestCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InquiryRequestCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InquiryRequestCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InquiryRequestCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *InquiryRequestCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InquiryRequestCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InquiryRequestCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *InquiryRequestCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetType

`func (o *InquiryRequestCreateDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InquiryRequestCreateDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InquiryRequestCreateDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InquiryRequestCreateDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *InquiryRequestCreateDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *InquiryRequestCreateDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetName

`func (o *InquiryRequestCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InquiryRequestCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InquiryRequestCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetLastName

`func (o *InquiryRequestCreateDto) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *InquiryRequestCreateDto) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *InquiryRequestCreateDto) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *InquiryRequestCreateDto) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *InquiryRequestCreateDto) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *InquiryRequestCreateDto) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetEmail

`func (o *InquiryRequestCreateDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InquiryRequestCreateDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InquiryRequestCreateDto) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetOrganizationName

`func (o *InquiryRequestCreateDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *InquiryRequestCreateDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *InquiryRequestCreateDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *InquiryRequestCreateDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *InquiryRequestCreateDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *InquiryRequestCreateDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetJobRole

`func (o *InquiryRequestCreateDto) GetJobRole() string`

GetJobRole returns the JobRole field if non-nil, zero value otherwise.

### GetJobRoleOk

`func (o *InquiryRequestCreateDto) GetJobRoleOk() (*string, bool)`

GetJobRoleOk returns a tuple with the JobRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRole

`func (o *InquiryRequestCreateDto) SetJobRole(v string)`

SetJobRole sets JobRole field to given value.

### HasJobRole

`func (o *InquiryRequestCreateDto) HasJobRole() bool`

HasJobRole returns a boolean if a field has been set.

### SetJobRoleNil

`func (o *InquiryRequestCreateDto) SetJobRoleNil(b bool)`

 SetJobRoleNil sets the value for JobRole to be an explicit nil

### UnsetJobRole
`func (o *InquiryRequestCreateDto) UnsetJobRole()`

UnsetJobRole ensures that no value is present for JobRole, not even an explicit nil
### GetOrganizationDomain

`func (o *InquiryRequestCreateDto) GetOrganizationDomain() string`

GetOrganizationDomain returns the OrganizationDomain field if non-nil, zero value otherwise.

### GetOrganizationDomainOk

`func (o *InquiryRequestCreateDto) GetOrganizationDomainOk() (*string, bool)`

GetOrganizationDomainOk returns a tuple with the OrganizationDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationDomain

`func (o *InquiryRequestCreateDto) SetOrganizationDomain(v string)`

SetOrganizationDomain sets OrganizationDomain field to given value.

### HasOrganizationDomain

`func (o *InquiryRequestCreateDto) HasOrganizationDomain() bool`

HasOrganizationDomain returns a boolean if a field has been set.

### SetOrganizationDomainNil

`func (o *InquiryRequestCreateDto) SetOrganizationDomainNil(b bool)`

 SetOrganizationDomainNil sets the value for OrganizationDomain to be an explicit nil

### UnsetOrganizationDomain
`func (o *InquiryRequestCreateDto) UnsetOrganizationDomain()`

UnsetOrganizationDomain ensures that no value is present for OrganizationDomain, not even an explicit nil
### GetCountryId

`func (o *InquiryRequestCreateDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *InquiryRequestCreateDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *InquiryRequestCreateDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *InquiryRequestCreateDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *InquiryRequestCreateDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *InquiryRequestCreateDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetPhone

`func (o *InquiryRequestCreateDto) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *InquiryRequestCreateDto) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *InquiryRequestCreateDto) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *InquiryRequestCreateDto) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *InquiryRequestCreateDto) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *InquiryRequestCreateDto) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetMessage

`func (o *InquiryRequestCreateDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InquiryRequestCreateDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InquiryRequestCreateDto) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSocialProfileId

`func (o *InquiryRequestCreateDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *InquiryRequestCreateDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *InquiryRequestCreateDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *InquiryRequestCreateDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *InquiryRequestCreateDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *InquiryRequestCreateDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


