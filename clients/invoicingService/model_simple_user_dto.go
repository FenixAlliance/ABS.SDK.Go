/*
InvoicingService

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.1.4089
Contact: support@fenix-alliance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the SimpleUserDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimpleUserDto{}

// SimpleUserDto struct for SimpleUserDto
type SimpleUserDto struct {
	Id NullableString `json:"id,omitempty"`
	Timestamp NullableTime `json:"timestamp,omitempty"`
	FullName NullableString `json:"fullName,omitempty"`
	QualifiedName NullableString `json:"qualifiedName,omitempty"`
	PublicName NullableString `json:"publicName,omitempty"`
	LastName NullableString `json:"lastName,omitempty"`
	FirstName NullableString `json:"firstName,omitempty"`
	CoverUrl NullableString `json:"coverUrl,omitempty"`
	AvatarUrl NullableString `json:"avatarUrl,omitempty"`
	GitHubUrl NullableString `json:"gitHubUrl,omitempty"`
	CountryId NullableString `json:"countryId,omitempty"`
	TimezoneId NullableString `json:"timezoneId,omitempty"`
	WebsiteUrl NullableString `json:"websiteUrl,omitempty"`
	TwitterUrl NullableString `json:"twitterUrl,omitempty"`
	YouTubeUrl NullableString `json:"youTubeUrl,omitempty"`
	LinkedInUrl NullableString `json:"linkedInUrl,omitempty"`
	FacebookUrl NullableString `json:"facebookUrl,omitempty"`
	InstagramUrl NullableString `json:"instagramUrl,omitempty"`
	SocialProfileId NullableString `json:"socialProfileId,omitempty"`
}

// NewSimpleUserDto instantiates a new SimpleUserDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimpleUserDto() *SimpleUserDto {
	this := SimpleUserDto{}
	return &this
}

// NewSimpleUserDtoWithDefaults instantiates a new SimpleUserDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimpleUserDtoWithDefaults() *SimpleUserDto {
	this := SimpleUserDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimpleUserDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimpleUserDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *SimpleUserDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *SimpleUserDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *SimpleUserDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *SimpleUserDto) UnsetId() {
	o.Id.Unset()
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimpleUserDto) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimpleUserDto) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *SimpleUserDto) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableTime and assigns it to the Timestamp field.
func (o *SimpleUserDto) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}
// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *SimpleUserDto) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *SimpleUserDto) UnsetTimestamp() {
	o.Timestamp.Unset()
}

// GetFullName returns the FullName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimpleUserDto) GetFullName() string {
	if o == nil || IsNil(o.FullName.Get()) {
		var ret string
		return ret
	}
	return *o.FullName.Get()
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimpleUserDto) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FullName.Get(), o.FullName.IsSet()
}

// HasFullName returns a boolean if a field has been set.
func (o *SimpleUserDto) HasFullName() bool {
	if o != nil && o.FullName.IsSet() {
		return true
	}

	return false
}

// SetFullName gets a reference to the given NullableString and assigns it to the FullName field.
func (o *SimpleUserDto) SetFullName(v string) {
	o.FullName.Set(&v)
}
// SetFullNameNil sets the value for FullName to be an explicit nil
func (o *SimpleUserDto) SetFullNameNil() {
	o.FullName.Set(nil)
}

// UnsetFullName ensures that no value is present for FullName, not even an explicit nil
func (o *SimpleUserDto) UnsetFullName() {
	o.FullName.Unset()
}

// GetQualifiedName returns the QualifiedName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimpleUserDto) GetQualifiedName() string {
	if o == nil || IsNil(o.QualifiedName.Get()) {
		var ret string
		return ret
	}
	return *o.QualifiedName.Get()
}

// GetQualifiedNameOk returns a tuple with the QualifiedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimpleUserDto) GetQualifiedNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.QualifiedName.Get(), o.QualifiedName.IsSet()
}

// HasQualifiedName returns a boolean if a field has been set.
func (o *SimpleUserDto) HasQualifiedName() bool {
	if o != nil && o.QualifiedName.IsSet() {
		return true
	}

	return false
}

// SetQualifiedName gets a reference to the given NullableString and assigns it to the QualifiedName field.
func (o *SimpleUserDto) SetQualifiedName(v string) {
	o.QualifiedName.Set(&v)
}
// SetQualifiedNameNil sets the value for QualifiedName to be an explicit nil
func (o *SimpleUserDto) SetQualifiedNameNil() {
	o.QualifiedName.Set(nil)
}

// UnsetQualifiedName ensures that no value is present for QualifiedName, not even an explicit nil
func (o *SimpleUserDto) UnsetQualifiedName() {
	o.QualifiedName.Unset()
}

// GetPublicName returns the PublicName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimpleUserDto) GetPublicName() string {
	if o == nil || IsNil(o.PublicName.Get()) {
		var ret string
		return ret
	}
	return *o.PublicName.Get()
}

// GetPublicNameOk returns a tuple with the PublicName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimpleUserDto) GetPublicNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicName.Get(), o.PublicName.IsSet()
}

// HasPublicName returns a boolean if a field has been set.
func (o *SimpleUserDto) HasPublicName() bool {
	if o != nil && o.PublicName.IsSet() {
		return true
	}

	return false
}

// SetPublicName gets a reference to the given NullableString and assigns it to the PublicName field.
func (o *SimpleUserDto) SetPublicName(v string) {
	o.PublicName.Set(&v)
}
// SetPublicNameNil sets the value for PublicName to be an explicit nil
func (o *SimpleUserDto) SetPublicNameNil() {
	o.PublicName.Set(nil)
}

// UnsetPublicName ensures that no value is present for PublicName, not even an explicit nil
func (o *SimpleUserDto) UnsetPublicName() {
	o.PublicName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimpleUserDto) GetLastName() string {
	if o == nil || IsNil(o.LastName.Get()) {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimpleUserDto) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *SimpleUserDto) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *SimpleUserDto) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *SimpleUserDto) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *SimpleUserDto) UnsetLastName() {
	o.LastName.Unset()
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimpleUserDto) GetFirstName() string {
	if o == nil || IsNil(o.FirstName.Get()) {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimpleUserDto) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *SimpleUserDto) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *SimpleUserDto) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *SimpleUserDto) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *SimpleUserDto) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetCoverUrl returns the CoverUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimpleUserDto) GetCoverUrl() string {
	if o == nil || IsNil(o.CoverUrl.Get()) {
		var ret string
		return ret
	}
	return *o.CoverUrl.Get()
}

// GetCoverUrlOk returns a tuple with the CoverUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimpleUserDto) GetCoverUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CoverUrl.Get(), o.CoverUrl.IsSet()
}

// HasCoverUrl returns a boolean if a field has been set.
func (o *SimpleUserDto) HasCoverUrl() bool {
	if o != nil && o.CoverUrl.IsSet() {
		return true
	}

	return false
}

// SetCoverUrl gets a reference to the given NullableString and assigns it to the CoverUrl field.
func (o *SimpleUserDto) SetCoverUrl(v string) {
	o.CoverUrl.Set(&v)
}
// SetCoverUrlNil sets the value for CoverUrl to be an explicit nil
func (o *SimpleUserDto) SetCoverUrlNil() {
	o.CoverUrl.Set(nil)
}

// UnsetCoverUrl ensures that no value is present for CoverUrl, not even an explicit nil
func (o *SimpleUserDto) UnsetCoverUrl() {
	o.CoverUrl.Unset()
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimpleUserDto) GetAvatarUrl() string {
	if o == nil || IsNil(o.AvatarUrl.Get()) {
		var ret string
		return ret
	}
	return *o.AvatarUrl.Get()
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimpleUserDto) GetAvatarUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvatarUrl.Get(), o.AvatarUrl.IsSet()
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *SimpleUserDto) HasAvatarUrl() bool {
	if o != nil && o.AvatarUrl.IsSet() {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given NullableString and assigns it to the AvatarUrl field.
func (o *SimpleUserDto) SetAvatarUrl(v string) {
	o.AvatarUrl.Set(&v)
}
// SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil
func (o *SimpleUserDto) SetAvatarUrlNil() {
	o.AvatarUrl.Set(nil)
}

// UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
func (o *SimpleUserDto) UnsetAvatarUrl() {
	o.AvatarUrl.Unset()
}

// GetGitHubUrl returns the GitHubUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimpleUserDto) GetGitHubUrl() string {
	if o == nil || IsNil(o.GitHubUrl.Get()) {
		var ret string
		return ret
	}
	return *o.GitHubUrl.Get()
}

// GetGitHubUrlOk returns a tuple with the GitHubUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimpleUserDto) GetGitHubUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GitHubUrl.Get(), o.GitHubUrl.IsSet()
}

// HasGitHubUrl returns a boolean if a field has been set.
func (o *SimpleUserDto) HasGitHubUrl() bool {
	if o != nil && o.GitHubUrl.IsSet() {
		return true
	}

	return false
}

// SetGitHubUrl gets a reference to the given NullableString and assigns it to the GitHubUrl field.
func (o *SimpleUserDto) SetGitHubUrl(v string) {
	o.GitHubUrl.Set(&v)
}
// SetGitHubUrlNil sets the value for GitHubUrl to be an explicit nil
func (o *SimpleUserDto) SetGitHubUrlNil() {
	o.GitHubUrl.Set(nil)
}

// UnsetGitHubUrl ensures that no value is present for GitHubUrl, not even an explicit nil
func (o *SimpleUserDto) UnsetGitHubUrl() {
	o.GitHubUrl.Unset()
}

// GetCountryId returns the CountryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimpleUserDto) GetCountryId() string {
	if o == nil || IsNil(o.CountryId.Get()) {
		var ret string
		return ret
	}
	return *o.CountryId.Get()
}

// GetCountryIdOk returns a tuple with the CountryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimpleUserDto) GetCountryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountryId.Get(), o.CountryId.IsSet()
}

// HasCountryId returns a boolean if a field has been set.
func (o *SimpleUserDto) HasCountryId() bool {
	if o != nil && o.CountryId.IsSet() {
		return true
	}

	return false
}

// SetCountryId gets a reference to the given NullableString and assigns it to the CountryId field.
func (o *SimpleUserDto) SetCountryId(v string) {
	o.CountryId.Set(&v)
}
// SetCountryIdNil sets the value for CountryId to be an explicit nil
func (o *SimpleUserDto) SetCountryIdNil() {
	o.CountryId.Set(nil)
}

// UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
func (o *SimpleUserDto) UnsetCountryId() {
	o.CountryId.Unset()
}

// GetTimezoneId returns the TimezoneId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimpleUserDto) GetTimezoneId() string {
	if o == nil || IsNil(o.TimezoneId.Get()) {
		var ret string
		return ret
	}
	return *o.TimezoneId.Get()
}

// GetTimezoneIdOk returns a tuple with the TimezoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimpleUserDto) GetTimezoneIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimezoneId.Get(), o.TimezoneId.IsSet()
}

// HasTimezoneId returns a boolean if a field has been set.
func (o *SimpleUserDto) HasTimezoneId() bool {
	if o != nil && o.TimezoneId.IsSet() {
		return true
	}

	return false
}

// SetTimezoneId gets a reference to the given NullableString and assigns it to the TimezoneId field.
func (o *SimpleUserDto) SetTimezoneId(v string) {
	o.TimezoneId.Set(&v)
}
// SetTimezoneIdNil sets the value for TimezoneId to be an explicit nil
func (o *SimpleUserDto) SetTimezoneIdNil() {
	o.TimezoneId.Set(nil)
}

// UnsetTimezoneId ensures that no value is present for TimezoneId, not even an explicit nil
func (o *SimpleUserDto) UnsetTimezoneId() {
	o.TimezoneId.Unset()
}

// GetWebsiteUrl returns the WebsiteUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimpleUserDto) GetWebsiteUrl() string {
	if o == nil || IsNil(o.WebsiteUrl.Get()) {
		var ret string
		return ret
	}
	return *o.WebsiteUrl.Get()
}

// GetWebsiteUrlOk returns a tuple with the WebsiteUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimpleUserDto) GetWebsiteUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebsiteUrl.Get(), o.WebsiteUrl.IsSet()
}

// HasWebsiteUrl returns a boolean if a field has been set.
func (o *SimpleUserDto) HasWebsiteUrl() bool {
	if o != nil && o.WebsiteUrl.IsSet() {
		return true
	}

	return false
}

// SetWebsiteUrl gets a reference to the given NullableString and assigns it to the WebsiteUrl field.
func (o *SimpleUserDto) SetWebsiteUrl(v string) {
	o.WebsiteUrl.Set(&v)
}
// SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil
func (o *SimpleUserDto) SetWebsiteUrlNil() {
	o.WebsiteUrl.Set(nil)
}

// UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
func (o *SimpleUserDto) UnsetWebsiteUrl() {
	o.WebsiteUrl.Unset()
}

// GetTwitterUrl returns the TwitterUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimpleUserDto) GetTwitterUrl() string {
	if o == nil || IsNil(o.TwitterUrl.Get()) {
		var ret string
		return ret
	}
	return *o.TwitterUrl.Get()
}

// GetTwitterUrlOk returns a tuple with the TwitterUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimpleUserDto) GetTwitterUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TwitterUrl.Get(), o.TwitterUrl.IsSet()
}

// HasTwitterUrl returns a boolean if a field has been set.
func (o *SimpleUserDto) HasTwitterUrl() bool {
	if o != nil && o.TwitterUrl.IsSet() {
		return true
	}

	return false
}

// SetTwitterUrl gets a reference to the given NullableString and assigns it to the TwitterUrl field.
func (o *SimpleUserDto) SetTwitterUrl(v string) {
	o.TwitterUrl.Set(&v)
}
// SetTwitterUrlNil sets the value for TwitterUrl to be an explicit nil
func (o *SimpleUserDto) SetTwitterUrlNil() {
	o.TwitterUrl.Set(nil)
}

// UnsetTwitterUrl ensures that no value is present for TwitterUrl, not even an explicit nil
func (o *SimpleUserDto) UnsetTwitterUrl() {
	o.TwitterUrl.Unset()
}

// GetYouTubeUrl returns the YouTubeUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimpleUserDto) GetYouTubeUrl() string {
	if o == nil || IsNil(o.YouTubeUrl.Get()) {
		var ret string
		return ret
	}
	return *o.YouTubeUrl.Get()
}

// GetYouTubeUrlOk returns a tuple with the YouTubeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimpleUserDto) GetYouTubeUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.YouTubeUrl.Get(), o.YouTubeUrl.IsSet()
}

// HasYouTubeUrl returns a boolean if a field has been set.
func (o *SimpleUserDto) HasYouTubeUrl() bool {
	if o != nil && o.YouTubeUrl.IsSet() {
		return true
	}

	return false
}

// SetYouTubeUrl gets a reference to the given NullableString and assigns it to the YouTubeUrl field.
func (o *SimpleUserDto) SetYouTubeUrl(v string) {
	o.YouTubeUrl.Set(&v)
}
// SetYouTubeUrlNil sets the value for YouTubeUrl to be an explicit nil
func (o *SimpleUserDto) SetYouTubeUrlNil() {
	o.YouTubeUrl.Set(nil)
}

// UnsetYouTubeUrl ensures that no value is present for YouTubeUrl, not even an explicit nil
func (o *SimpleUserDto) UnsetYouTubeUrl() {
	o.YouTubeUrl.Unset()
}

// GetLinkedInUrl returns the LinkedInUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimpleUserDto) GetLinkedInUrl() string {
	if o == nil || IsNil(o.LinkedInUrl.Get()) {
		var ret string
		return ret
	}
	return *o.LinkedInUrl.Get()
}

// GetLinkedInUrlOk returns a tuple with the LinkedInUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimpleUserDto) GetLinkedInUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinkedInUrl.Get(), o.LinkedInUrl.IsSet()
}

// HasLinkedInUrl returns a boolean if a field has been set.
func (o *SimpleUserDto) HasLinkedInUrl() bool {
	if o != nil && o.LinkedInUrl.IsSet() {
		return true
	}

	return false
}

// SetLinkedInUrl gets a reference to the given NullableString and assigns it to the LinkedInUrl field.
func (o *SimpleUserDto) SetLinkedInUrl(v string) {
	o.LinkedInUrl.Set(&v)
}
// SetLinkedInUrlNil sets the value for LinkedInUrl to be an explicit nil
func (o *SimpleUserDto) SetLinkedInUrlNil() {
	o.LinkedInUrl.Set(nil)
}

// UnsetLinkedInUrl ensures that no value is present for LinkedInUrl, not even an explicit nil
func (o *SimpleUserDto) UnsetLinkedInUrl() {
	o.LinkedInUrl.Unset()
}

// GetFacebookUrl returns the FacebookUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimpleUserDto) GetFacebookUrl() string {
	if o == nil || IsNil(o.FacebookUrl.Get()) {
		var ret string
		return ret
	}
	return *o.FacebookUrl.Get()
}

// GetFacebookUrlOk returns a tuple with the FacebookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimpleUserDto) GetFacebookUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FacebookUrl.Get(), o.FacebookUrl.IsSet()
}

// HasFacebookUrl returns a boolean if a field has been set.
func (o *SimpleUserDto) HasFacebookUrl() bool {
	if o != nil && o.FacebookUrl.IsSet() {
		return true
	}

	return false
}

// SetFacebookUrl gets a reference to the given NullableString and assigns it to the FacebookUrl field.
func (o *SimpleUserDto) SetFacebookUrl(v string) {
	o.FacebookUrl.Set(&v)
}
// SetFacebookUrlNil sets the value for FacebookUrl to be an explicit nil
func (o *SimpleUserDto) SetFacebookUrlNil() {
	o.FacebookUrl.Set(nil)
}

// UnsetFacebookUrl ensures that no value is present for FacebookUrl, not even an explicit nil
func (o *SimpleUserDto) UnsetFacebookUrl() {
	o.FacebookUrl.Unset()
}

// GetInstagramUrl returns the InstagramUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimpleUserDto) GetInstagramUrl() string {
	if o == nil || IsNil(o.InstagramUrl.Get()) {
		var ret string
		return ret
	}
	return *o.InstagramUrl.Get()
}

// GetInstagramUrlOk returns a tuple with the InstagramUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimpleUserDto) GetInstagramUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstagramUrl.Get(), o.InstagramUrl.IsSet()
}

// HasInstagramUrl returns a boolean if a field has been set.
func (o *SimpleUserDto) HasInstagramUrl() bool {
	if o != nil && o.InstagramUrl.IsSet() {
		return true
	}

	return false
}

// SetInstagramUrl gets a reference to the given NullableString and assigns it to the InstagramUrl field.
func (o *SimpleUserDto) SetInstagramUrl(v string) {
	o.InstagramUrl.Set(&v)
}
// SetInstagramUrlNil sets the value for InstagramUrl to be an explicit nil
func (o *SimpleUserDto) SetInstagramUrlNil() {
	o.InstagramUrl.Set(nil)
}

// UnsetInstagramUrl ensures that no value is present for InstagramUrl, not even an explicit nil
func (o *SimpleUserDto) UnsetInstagramUrl() {
	o.InstagramUrl.Unset()
}

// GetSocialProfileId returns the SocialProfileId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimpleUserDto) GetSocialProfileId() string {
	if o == nil || IsNil(o.SocialProfileId.Get()) {
		var ret string
		return ret
	}
	return *o.SocialProfileId.Get()
}

// GetSocialProfileIdOk returns a tuple with the SocialProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimpleUserDto) GetSocialProfileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SocialProfileId.Get(), o.SocialProfileId.IsSet()
}

// HasSocialProfileId returns a boolean if a field has been set.
func (o *SimpleUserDto) HasSocialProfileId() bool {
	if o != nil && o.SocialProfileId.IsSet() {
		return true
	}

	return false
}

// SetSocialProfileId gets a reference to the given NullableString and assigns it to the SocialProfileId field.
func (o *SimpleUserDto) SetSocialProfileId(v string) {
	o.SocialProfileId.Set(&v)
}
// SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil
func (o *SimpleUserDto) SetSocialProfileIdNil() {
	o.SocialProfileId.Set(nil)
}

// UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil
func (o *SimpleUserDto) UnsetSocialProfileId() {
	o.SocialProfileId.Unset()
}

func (o SimpleUserDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimpleUserDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	if o.FullName.IsSet() {
		toSerialize["fullName"] = o.FullName.Get()
	}
	if o.QualifiedName.IsSet() {
		toSerialize["qualifiedName"] = o.QualifiedName.Get()
	}
	if o.PublicName.IsSet() {
		toSerialize["publicName"] = o.PublicName.Get()
	}
	if o.LastName.IsSet() {
		toSerialize["lastName"] = o.LastName.Get()
	}
	if o.FirstName.IsSet() {
		toSerialize["firstName"] = o.FirstName.Get()
	}
	if o.CoverUrl.IsSet() {
		toSerialize["coverUrl"] = o.CoverUrl.Get()
	}
	if o.AvatarUrl.IsSet() {
		toSerialize["avatarUrl"] = o.AvatarUrl.Get()
	}
	if o.GitHubUrl.IsSet() {
		toSerialize["gitHubUrl"] = o.GitHubUrl.Get()
	}
	if o.CountryId.IsSet() {
		toSerialize["countryId"] = o.CountryId.Get()
	}
	if o.TimezoneId.IsSet() {
		toSerialize["timezoneId"] = o.TimezoneId.Get()
	}
	if o.WebsiteUrl.IsSet() {
		toSerialize["websiteUrl"] = o.WebsiteUrl.Get()
	}
	if o.TwitterUrl.IsSet() {
		toSerialize["twitterUrl"] = o.TwitterUrl.Get()
	}
	if o.YouTubeUrl.IsSet() {
		toSerialize["youTubeUrl"] = o.YouTubeUrl.Get()
	}
	if o.LinkedInUrl.IsSet() {
		toSerialize["linkedInUrl"] = o.LinkedInUrl.Get()
	}
	if o.FacebookUrl.IsSet() {
		toSerialize["facebookUrl"] = o.FacebookUrl.Get()
	}
	if o.InstagramUrl.IsSet() {
		toSerialize["instagramUrl"] = o.InstagramUrl.Get()
	}
	if o.SocialProfileId.IsSet() {
		toSerialize["socialProfileId"] = o.SocialProfileId.Get()
	}
	return toSerialize, nil
}

type NullableSimpleUserDto struct {
	value *SimpleUserDto
	isSet bool
}

func (v NullableSimpleUserDto) Get() *SimpleUserDto {
	return v.value
}

func (v *NullableSimpleUserDto) Set(val *SimpleUserDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSimpleUserDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSimpleUserDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimpleUserDto(val *SimpleUserDto) *NullableSimpleUserDto {
	return &NullableSimpleUserDto{value: val, isSet: true}
}

func (v NullableSimpleUserDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimpleUserDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


