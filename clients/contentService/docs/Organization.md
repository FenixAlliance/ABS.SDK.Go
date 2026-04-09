# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LegalName** | Pointer to **NullableString** |  | [optional] 
**BusinessName** | Pointer to **NullableString** |  | [optional] 
**Slogan** | Pointer to **NullableString** |  | [optional] 
**Homepage** | Pointer to **NullableString** |  | [optional] 
**FacebookPageUsername** | Pointer to **NullableString** |  | [optional] 
**InstagramUsername** | Pointer to **NullableString** |  | [optional] 
**LinkedInUsername** | Pointer to **NullableString** |  | [optional] 
**TwitterHandler** | Pointer to **NullableString** |  | [optional] 
**GitHubUsername** | Pointer to **NullableString** |  | [optional] 
**ContactPoints** | Pointer to [**[]ContactPoint**](ContactPoint.md) |  | [optional] 

## Methods

### NewOrganization

`func NewOrganization() *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLegalName

`func (o *Organization) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *Organization) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *Organization) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *Organization) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### SetLegalNameNil

`func (o *Organization) SetLegalNameNil(b bool)`

 SetLegalNameNil sets the value for LegalName to be an explicit nil

### UnsetLegalName
`func (o *Organization) UnsetLegalName()`

UnsetLegalName ensures that no value is present for LegalName, not even an explicit nil
### GetBusinessName

`func (o *Organization) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *Organization) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *Organization) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.

### HasBusinessName

`func (o *Organization) HasBusinessName() bool`

HasBusinessName returns a boolean if a field has been set.

### SetBusinessNameNil

`func (o *Organization) SetBusinessNameNil(b bool)`

 SetBusinessNameNil sets the value for BusinessName to be an explicit nil

### UnsetBusinessName
`func (o *Organization) UnsetBusinessName()`

UnsetBusinessName ensures that no value is present for BusinessName, not even an explicit nil
### GetSlogan

`func (o *Organization) GetSlogan() string`

GetSlogan returns the Slogan field if non-nil, zero value otherwise.

### GetSloganOk

`func (o *Organization) GetSloganOk() (*string, bool)`

GetSloganOk returns a tuple with the Slogan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlogan

`func (o *Organization) SetSlogan(v string)`

SetSlogan sets Slogan field to given value.

### HasSlogan

`func (o *Organization) HasSlogan() bool`

HasSlogan returns a boolean if a field has been set.

### SetSloganNil

`func (o *Organization) SetSloganNil(b bool)`

 SetSloganNil sets the value for Slogan to be an explicit nil

### UnsetSlogan
`func (o *Organization) UnsetSlogan()`

UnsetSlogan ensures that no value is present for Slogan, not even an explicit nil
### GetHomepage

`func (o *Organization) GetHomepage() string`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *Organization) GetHomepageOk() (*string, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *Organization) SetHomepage(v string)`

SetHomepage sets Homepage field to given value.

### HasHomepage

`func (o *Organization) HasHomepage() bool`

HasHomepage returns a boolean if a field has been set.

### SetHomepageNil

`func (o *Organization) SetHomepageNil(b bool)`

 SetHomepageNil sets the value for Homepage to be an explicit nil

### UnsetHomepage
`func (o *Organization) UnsetHomepage()`

UnsetHomepage ensures that no value is present for Homepage, not even an explicit nil
### GetFacebookPageUsername

`func (o *Organization) GetFacebookPageUsername() string`

GetFacebookPageUsername returns the FacebookPageUsername field if non-nil, zero value otherwise.

### GetFacebookPageUsernameOk

`func (o *Organization) GetFacebookPageUsernameOk() (*string, bool)`

GetFacebookPageUsernameOk returns a tuple with the FacebookPageUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookPageUsername

`func (o *Organization) SetFacebookPageUsername(v string)`

SetFacebookPageUsername sets FacebookPageUsername field to given value.

### HasFacebookPageUsername

`func (o *Organization) HasFacebookPageUsername() bool`

HasFacebookPageUsername returns a boolean if a field has been set.

### SetFacebookPageUsernameNil

`func (o *Organization) SetFacebookPageUsernameNil(b bool)`

 SetFacebookPageUsernameNil sets the value for FacebookPageUsername to be an explicit nil

### UnsetFacebookPageUsername
`func (o *Organization) UnsetFacebookPageUsername()`

UnsetFacebookPageUsername ensures that no value is present for FacebookPageUsername, not even an explicit nil
### GetInstagramUsername

`func (o *Organization) GetInstagramUsername() string`

GetInstagramUsername returns the InstagramUsername field if non-nil, zero value otherwise.

### GetInstagramUsernameOk

`func (o *Organization) GetInstagramUsernameOk() (*string, bool)`

GetInstagramUsernameOk returns a tuple with the InstagramUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramUsername

`func (o *Organization) SetInstagramUsername(v string)`

SetInstagramUsername sets InstagramUsername field to given value.

### HasInstagramUsername

`func (o *Organization) HasInstagramUsername() bool`

HasInstagramUsername returns a boolean if a field has been set.

### SetInstagramUsernameNil

`func (o *Organization) SetInstagramUsernameNil(b bool)`

 SetInstagramUsernameNil sets the value for InstagramUsername to be an explicit nil

### UnsetInstagramUsername
`func (o *Organization) UnsetInstagramUsername()`

UnsetInstagramUsername ensures that no value is present for InstagramUsername, not even an explicit nil
### GetLinkedInUsername

`func (o *Organization) GetLinkedInUsername() string`

GetLinkedInUsername returns the LinkedInUsername field if non-nil, zero value otherwise.

### GetLinkedInUsernameOk

`func (o *Organization) GetLinkedInUsernameOk() (*string, bool)`

GetLinkedInUsernameOk returns a tuple with the LinkedInUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedInUsername

`func (o *Organization) SetLinkedInUsername(v string)`

SetLinkedInUsername sets LinkedInUsername field to given value.

### HasLinkedInUsername

`func (o *Organization) HasLinkedInUsername() bool`

HasLinkedInUsername returns a boolean if a field has been set.

### SetLinkedInUsernameNil

`func (o *Organization) SetLinkedInUsernameNil(b bool)`

 SetLinkedInUsernameNil sets the value for LinkedInUsername to be an explicit nil

### UnsetLinkedInUsername
`func (o *Organization) UnsetLinkedInUsername()`

UnsetLinkedInUsername ensures that no value is present for LinkedInUsername, not even an explicit nil
### GetTwitterHandler

`func (o *Organization) GetTwitterHandler() string`

GetTwitterHandler returns the TwitterHandler field if non-nil, zero value otherwise.

### GetTwitterHandlerOk

`func (o *Organization) GetTwitterHandlerOk() (*string, bool)`

GetTwitterHandlerOk returns a tuple with the TwitterHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterHandler

`func (o *Organization) SetTwitterHandler(v string)`

SetTwitterHandler sets TwitterHandler field to given value.

### HasTwitterHandler

`func (o *Organization) HasTwitterHandler() bool`

HasTwitterHandler returns a boolean if a field has been set.

### SetTwitterHandlerNil

`func (o *Organization) SetTwitterHandlerNil(b bool)`

 SetTwitterHandlerNil sets the value for TwitterHandler to be an explicit nil

### UnsetTwitterHandler
`func (o *Organization) UnsetTwitterHandler()`

UnsetTwitterHandler ensures that no value is present for TwitterHandler, not even an explicit nil
### GetGitHubUsername

`func (o *Organization) GetGitHubUsername() string`

GetGitHubUsername returns the GitHubUsername field if non-nil, zero value otherwise.

### GetGitHubUsernameOk

`func (o *Organization) GetGitHubUsernameOk() (*string, bool)`

GetGitHubUsernameOk returns a tuple with the GitHubUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitHubUsername

`func (o *Organization) SetGitHubUsername(v string)`

SetGitHubUsername sets GitHubUsername field to given value.

### HasGitHubUsername

`func (o *Organization) HasGitHubUsername() bool`

HasGitHubUsername returns a boolean if a field has been set.

### SetGitHubUsernameNil

`func (o *Organization) SetGitHubUsernameNil(b bool)`

 SetGitHubUsernameNil sets the value for GitHubUsername to be an explicit nil

### UnsetGitHubUsername
`func (o *Organization) UnsetGitHubUsername()`

UnsetGitHubUsername ensures that no value is present for GitHubUsername, not even an explicit nil
### GetContactPoints

`func (o *Organization) GetContactPoints() []ContactPoint`

GetContactPoints returns the ContactPoints field if non-nil, zero value otherwise.

### GetContactPointsOk

`func (o *Organization) GetContactPointsOk() (*[]ContactPoint, bool)`

GetContactPointsOk returns a tuple with the ContactPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPoints

`func (o *Organization) SetContactPoints(v []ContactPoint)`

SetContactPoints sets ContactPoints field to given value.

### HasContactPoints

`func (o *Organization) HasContactPoints() bool`

HasContactPoints returns a boolean if a field has been set.

### SetContactPointsNil

`func (o *Organization) SetContactPointsNil(b bool)`

 SetContactPointsNil sets the value for ContactPoints to be an explicit nil

### UnsetContactPoints
`func (o *Organization) UnsetContactPoints()`

UnsetContactPoints ensures that no value is present for ContactPoints, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


