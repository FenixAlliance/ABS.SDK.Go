# ISwaggerSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**TermsOfService** | Pointer to **NullableString** |  | [optional] 
**SwaggerEndpoint** | Pointer to [**ISwaggerEndpoint**](ISwaggerEndpoint.md) |  | [optional] 
**OpenApiContact** | Pointer to [**ISwaggerContact**](ISwaggerContact.md) |  | [optional] 
**License** | Pointer to [**ISwaggerLicense**](ISwaggerLicense.md) |  | [optional] 

## Methods

### NewISwaggerSpec

`func NewISwaggerSpec() *ISwaggerSpec`

NewISwaggerSpec instantiates a new ISwaggerSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewISwaggerSpecWithDefaults

`func NewISwaggerSpecWithDefaults() *ISwaggerSpec`

NewISwaggerSpecWithDefaults instantiates a new ISwaggerSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *ISwaggerSpec) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ISwaggerSpec) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ISwaggerSpec) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *ISwaggerSpec) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetName

`func (o *ISwaggerSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ISwaggerSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ISwaggerSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ISwaggerSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ISwaggerSpec) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ISwaggerSpec) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTitle

`func (o *ISwaggerSpec) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ISwaggerSpec) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ISwaggerSpec) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ISwaggerSpec) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ISwaggerSpec) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ISwaggerSpec) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetVersion

`func (o *ISwaggerSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ISwaggerSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ISwaggerSpec) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ISwaggerSpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ISwaggerSpec) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ISwaggerSpec) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetDescription

`func (o *ISwaggerSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ISwaggerSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ISwaggerSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ISwaggerSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ISwaggerSpec) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ISwaggerSpec) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTermsOfService

`func (o *ISwaggerSpec) GetTermsOfService() string`

GetTermsOfService returns the TermsOfService field if non-nil, zero value otherwise.

### GetTermsOfServiceOk

`func (o *ISwaggerSpec) GetTermsOfServiceOk() (*string, bool)`

GetTermsOfServiceOk returns a tuple with the TermsOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfService

`func (o *ISwaggerSpec) SetTermsOfService(v string)`

SetTermsOfService sets TermsOfService field to given value.

### HasTermsOfService

`func (o *ISwaggerSpec) HasTermsOfService() bool`

HasTermsOfService returns a boolean if a field has been set.

### SetTermsOfServiceNil

`func (o *ISwaggerSpec) SetTermsOfServiceNil(b bool)`

 SetTermsOfServiceNil sets the value for TermsOfService to be an explicit nil

### UnsetTermsOfService
`func (o *ISwaggerSpec) UnsetTermsOfService()`

UnsetTermsOfService ensures that no value is present for TermsOfService, not even an explicit nil
### GetSwaggerEndpoint

`func (o *ISwaggerSpec) GetSwaggerEndpoint() ISwaggerEndpoint`

GetSwaggerEndpoint returns the SwaggerEndpoint field if non-nil, zero value otherwise.

### GetSwaggerEndpointOk

`func (o *ISwaggerSpec) GetSwaggerEndpointOk() (*ISwaggerEndpoint, bool)`

GetSwaggerEndpointOk returns a tuple with the SwaggerEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwaggerEndpoint

`func (o *ISwaggerSpec) SetSwaggerEndpoint(v ISwaggerEndpoint)`

SetSwaggerEndpoint sets SwaggerEndpoint field to given value.

### HasSwaggerEndpoint

`func (o *ISwaggerSpec) HasSwaggerEndpoint() bool`

HasSwaggerEndpoint returns a boolean if a field has been set.

### GetOpenApiContact

`func (o *ISwaggerSpec) GetOpenApiContact() ISwaggerContact`

GetOpenApiContact returns the OpenApiContact field if non-nil, zero value otherwise.

### GetOpenApiContactOk

`func (o *ISwaggerSpec) GetOpenApiContactOk() (*ISwaggerContact, bool)`

GetOpenApiContactOk returns a tuple with the OpenApiContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenApiContact

`func (o *ISwaggerSpec) SetOpenApiContact(v ISwaggerContact)`

SetOpenApiContact sets OpenApiContact field to given value.

### HasOpenApiContact

`func (o *ISwaggerSpec) HasOpenApiContact() bool`

HasOpenApiContact returns a boolean if a field has been set.

### GetLicense

`func (o *ISwaggerSpec) GetLicense() ISwaggerLicense`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *ISwaggerSpec) GetLicenseOk() (*ISwaggerLicense, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *ISwaggerSpec) SetLicense(v ISwaggerLicense)`

SetLicense sets License field to given value.

### HasLicense

`func (o *ISwaggerSpec) HasLicense() bool`

HasLicense returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


