# IOpenApiDefinitionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**TermsOfService** | Pointer to **NullableString** |  | [optional] 
**OpenApiEndpoint** | Pointer to [**IOpenApiEndpoint**](IOpenApiEndpoint.md) |  | [optional] 
**OpenApiContact** | Pointer to [**IOpenApiContact**](IOpenApiContact.md) |  | [optional] 
**License** | Pointer to [**IOpenApiLicense**](IOpenApiLicense.md) |  | [optional] 

## Methods

### NewIOpenApiDefinitionSpec

`func NewIOpenApiDefinitionSpec() *IOpenApiDefinitionSpec`

NewIOpenApiDefinitionSpec instantiates a new IOpenApiDefinitionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIOpenApiDefinitionSpecWithDefaults

`func NewIOpenApiDefinitionSpecWithDefaults() *IOpenApiDefinitionSpec`

NewIOpenApiDefinitionSpecWithDefaults instantiates a new IOpenApiDefinitionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *IOpenApiDefinitionSpec) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *IOpenApiDefinitionSpec) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *IOpenApiDefinitionSpec) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *IOpenApiDefinitionSpec) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetName

`func (o *IOpenApiDefinitionSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IOpenApiDefinitionSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IOpenApiDefinitionSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IOpenApiDefinitionSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IOpenApiDefinitionSpec) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IOpenApiDefinitionSpec) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTitle

`func (o *IOpenApiDefinitionSpec) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IOpenApiDefinitionSpec) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IOpenApiDefinitionSpec) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IOpenApiDefinitionSpec) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *IOpenApiDefinitionSpec) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *IOpenApiDefinitionSpec) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetVersion

`func (o *IOpenApiDefinitionSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IOpenApiDefinitionSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IOpenApiDefinitionSpec) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IOpenApiDefinitionSpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *IOpenApiDefinitionSpec) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *IOpenApiDefinitionSpec) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetDescription

`func (o *IOpenApiDefinitionSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IOpenApiDefinitionSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IOpenApiDefinitionSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IOpenApiDefinitionSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *IOpenApiDefinitionSpec) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *IOpenApiDefinitionSpec) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTermsOfService

`func (o *IOpenApiDefinitionSpec) GetTermsOfService() string`

GetTermsOfService returns the TermsOfService field if non-nil, zero value otherwise.

### GetTermsOfServiceOk

`func (o *IOpenApiDefinitionSpec) GetTermsOfServiceOk() (*string, bool)`

GetTermsOfServiceOk returns a tuple with the TermsOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfService

`func (o *IOpenApiDefinitionSpec) SetTermsOfService(v string)`

SetTermsOfService sets TermsOfService field to given value.

### HasTermsOfService

`func (o *IOpenApiDefinitionSpec) HasTermsOfService() bool`

HasTermsOfService returns a boolean if a field has been set.

### SetTermsOfServiceNil

`func (o *IOpenApiDefinitionSpec) SetTermsOfServiceNil(b bool)`

 SetTermsOfServiceNil sets the value for TermsOfService to be an explicit nil

### UnsetTermsOfService
`func (o *IOpenApiDefinitionSpec) UnsetTermsOfService()`

UnsetTermsOfService ensures that no value is present for TermsOfService, not even an explicit nil
### GetOpenApiEndpoint

`func (o *IOpenApiDefinitionSpec) GetOpenApiEndpoint() IOpenApiEndpoint`

GetOpenApiEndpoint returns the OpenApiEndpoint field if non-nil, zero value otherwise.

### GetOpenApiEndpointOk

`func (o *IOpenApiDefinitionSpec) GetOpenApiEndpointOk() (*IOpenApiEndpoint, bool)`

GetOpenApiEndpointOk returns a tuple with the OpenApiEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenApiEndpoint

`func (o *IOpenApiDefinitionSpec) SetOpenApiEndpoint(v IOpenApiEndpoint)`

SetOpenApiEndpoint sets OpenApiEndpoint field to given value.

### HasOpenApiEndpoint

`func (o *IOpenApiDefinitionSpec) HasOpenApiEndpoint() bool`

HasOpenApiEndpoint returns a boolean if a field has been set.

### GetOpenApiContact

`func (o *IOpenApiDefinitionSpec) GetOpenApiContact() IOpenApiContact`

GetOpenApiContact returns the OpenApiContact field if non-nil, zero value otherwise.

### GetOpenApiContactOk

`func (o *IOpenApiDefinitionSpec) GetOpenApiContactOk() (*IOpenApiContact, bool)`

GetOpenApiContactOk returns a tuple with the OpenApiContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenApiContact

`func (o *IOpenApiDefinitionSpec) SetOpenApiContact(v IOpenApiContact)`

SetOpenApiContact sets OpenApiContact field to given value.

### HasOpenApiContact

`func (o *IOpenApiDefinitionSpec) HasOpenApiContact() bool`

HasOpenApiContact returns a boolean if a field has been set.

### GetLicense

`func (o *IOpenApiDefinitionSpec) GetLicense() IOpenApiLicense`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *IOpenApiDefinitionSpec) GetLicenseOk() (*IOpenApiLicense, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *IOpenApiDefinitionSpec) SetLicense(v IOpenApiLicense)`

SetLicense sets License field to given value.

### HasLicense

`func (o *IOpenApiDefinitionSpec) HasLicense() bool`

HasLicense returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


