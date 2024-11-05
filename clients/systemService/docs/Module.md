# Module

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] [readonly] 
**Order** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**FullName** | Pointer to **NullableString** |  | [optional] [readonly] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 
**Configuration** | Pointer to **NullableString** |  | [optional] 
**Author** | Pointer to **NullableString** |  | [optional] 
**AuthorUrl** | Pointer to **NullableString** |  | [optional] 
**License** | Pointer to **NullableString** |  | [optional] 
**RequireLicenseAcceptance** | Pointer to **NullableBool** |  | [optional] 
**Repository** | Pointer to **NullableString** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**Image** | Pointer to **NullableString** |  | [optional] 
**NuSpecPath** | Pointer to **NullableString** |  | [optional] 
**Manifest** | Pointer to **NullableString** |  | [optional] 
**Documentation** | Pointer to **NullableString** |  | [optional] 
**Website** | Pointer to **NullableString** |  | [optional] 
**Logo** | Pointer to **NullableString** |  | [optional] 
**SwaggerSpec** | Pointer to [**ISwaggerSpec**](ISwaggerSpec.md) |  | [optional] 
**SwaggerSpecs** | Pointer to [**[]ISwaggerSpec**](ISwaggerSpec.md) |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**AssemblyPaths** | Pointer to **[]string** |  | [optional] 
**MarkedForDeletion** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewModule

`func NewModule() *Module`

NewModule instantiates a new Module object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleWithDefaults

`func NewModuleWithDefaults() *Module`

NewModuleWithDefaults instantiates a new Module object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *Module) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *Module) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *Module) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *Module) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetActive

`func (o *Module) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Module) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Module) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Module) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetOrder

`func (o *Module) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Module) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Module) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Module) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetId

`func (o *Module) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Module) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Module) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Module) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Module) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Module) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *Module) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Module) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Module) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Module) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Module) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Module) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFullName

`func (o *Module) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *Module) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *Module) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *Module) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *Module) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *Module) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetDescription

`func (o *Module) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Module) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Module) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Module) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Module) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Module) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *Module) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Module) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Module) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *Module) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConfiguration

`func (o *Module) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *Module) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *Module) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *Module) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *Module) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *Module) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetAuthor

`func (o *Module) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Module) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Module) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *Module) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *Module) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *Module) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetAuthorUrl

`func (o *Module) GetAuthorUrl() string`

GetAuthorUrl returns the AuthorUrl field if non-nil, zero value otherwise.

### GetAuthorUrlOk

`func (o *Module) GetAuthorUrlOk() (*string, bool)`

GetAuthorUrlOk returns a tuple with the AuthorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorUrl

`func (o *Module) SetAuthorUrl(v string)`

SetAuthorUrl sets AuthorUrl field to given value.

### HasAuthorUrl

`func (o *Module) HasAuthorUrl() bool`

HasAuthorUrl returns a boolean if a field has been set.

### SetAuthorUrlNil

`func (o *Module) SetAuthorUrlNil(b bool)`

 SetAuthorUrlNil sets the value for AuthorUrl to be an explicit nil

### UnsetAuthorUrl
`func (o *Module) UnsetAuthorUrl()`

UnsetAuthorUrl ensures that no value is present for AuthorUrl, not even an explicit nil
### GetLicense

`func (o *Module) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *Module) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *Module) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *Module) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### SetLicenseNil

`func (o *Module) SetLicenseNil(b bool)`

 SetLicenseNil sets the value for License to be an explicit nil

### UnsetLicense
`func (o *Module) UnsetLicense()`

UnsetLicense ensures that no value is present for License, not even an explicit nil
### GetRequireLicenseAcceptance

`func (o *Module) GetRequireLicenseAcceptance() bool`

GetRequireLicenseAcceptance returns the RequireLicenseAcceptance field if non-nil, zero value otherwise.

### GetRequireLicenseAcceptanceOk

`func (o *Module) GetRequireLicenseAcceptanceOk() (*bool, bool)`

GetRequireLicenseAcceptanceOk returns a tuple with the RequireLicenseAcceptance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireLicenseAcceptance

`func (o *Module) SetRequireLicenseAcceptance(v bool)`

SetRequireLicenseAcceptance sets RequireLicenseAcceptance field to given value.

### HasRequireLicenseAcceptance

`func (o *Module) HasRequireLicenseAcceptance() bool`

HasRequireLicenseAcceptance returns a boolean if a field has been set.

### SetRequireLicenseAcceptanceNil

`func (o *Module) SetRequireLicenseAcceptanceNil(b bool)`

 SetRequireLicenseAcceptanceNil sets the value for RequireLicenseAcceptance to be an explicit nil

### UnsetRequireLicenseAcceptance
`func (o *Module) UnsetRequireLicenseAcceptance()`

UnsetRequireLicenseAcceptance ensures that no value is present for RequireLicenseAcceptance, not even an explicit nil
### GetRepository

`func (o *Module) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *Module) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *Module) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *Module) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *Module) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *Module) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetPath

`func (o *Module) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Module) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Module) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Module) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *Module) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *Module) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetIcon

`func (o *Module) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Module) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Module) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *Module) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *Module) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *Module) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetImage

`func (o *Module) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Module) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Module) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *Module) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *Module) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *Module) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetNuSpecPath

`func (o *Module) GetNuSpecPath() string`

GetNuSpecPath returns the NuSpecPath field if non-nil, zero value otherwise.

### GetNuSpecPathOk

`func (o *Module) GetNuSpecPathOk() (*string, bool)`

GetNuSpecPathOk returns a tuple with the NuSpecPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuSpecPath

`func (o *Module) SetNuSpecPath(v string)`

SetNuSpecPath sets NuSpecPath field to given value.

### HasNuSpecPath

`func (o *Module) HasNuSpecPath() bool`

HasNuSpecPath returns a boolean if a field has been set.

### SetNuSpecPathNil

`func (o *Module) SetNuSpecPathNil(b bool)`

 SetNuSpecPathNil sets the value for NuSpecPath to be an explicit nil

### UnsetNuSpecPath
`func (o *Module) UnsetNuSpecPath()`

UnsetNuSpecPath ensures that no value is present for NuSpecPath, not even an explicit nil
### GetManifest

`func (o *Module) GetManifest() string`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *Module) GetManifestOk() (*string, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *Module) SetManifest(v string)`

SetManifest sets Manifest field to given value.

### HasManifest

`func (o *Module) HasManifest() bool`

HasManifest returns a boolean if a field has been set.

### SetManifestNil

`func (o *Module) SetManifestNil(b bool)`

 SetManifestNil sets the value for Manifest to be an explicit nil

### UnsetManifest
`func (o *Module) UnsetManifest()`

UnsetManifest ensures that no value is present for Manifest, not even an explicit nil
### GetDocumentation

`func (o *Module) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *Module) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *Module) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *Module) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### SetDocumentationNil

`func (o *Module) SetDocumentationNil(b bool)`

 SetDocumentationNil sets the value for Documentation to be an explicit nil

### UnsetDocumentation
`func (o *Module) UnsetDocumentation()`

UnsetDocumentation ensures that no value is present for Documentation, not even an explicit nil
### GetWebsite

`func (o *Module) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *Module) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *Module) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *Module) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### SetWebsiteNil

`func (o *Module) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *Module) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil
### GetLogo

`func (o *Module) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *Module) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *Module) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *Module) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### SetLogoNil

`func (o *Module) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *Module) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil
### GetSwaggerSpec

`func (o *Module) GetSwaggerSpec() ISwaggerSpec`

GetSwaggerSpec returns the SwaggerSpec field if non-nil, zero value otherwise.

### GetSwaggerSpecOk

`func (o *Module) GetSwaggerSpecOk() (*ISwaggerSpec, bool)`

GetSwaggerSpecOk returns a tuple with the SwaggerSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwaggerSpec

`func (o *Module) SetSwaggerSpec(v ISwaggerSpec)`

SetSwaggerSpec sets SwaggerSpec field to given value.

### HasSwaggerSpec

`func (o *Module) HasSwaggerSpec() bool`

HasSwaggerSpec returns a boolean if a field has been set.

### GetSwaggerSpecs

`func (o *Module) GetSwaggerSpecs() []ISwaggerSpec`

GetSwaggerSpecs returns the SwaggerSpecs field if non-nil, zero value otherwise.

### GetSwaggerSpecsOk

`func (o *Module) GetSwaggerSpecsOk() (*[]ISwaggerSpec, bool)`

GetSwaggerSpecsOk returns a tuple with the SwaggerSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwaggerSpecs

`func (o *Module) SetSwaggerSpecs(v []ISwaggerSpec)`

SetSwaggerSpecs sets SwaggerSpecs field to given value.

### HasSwaggerSpecs

`func (o *Module) HasSwaggerSpecs() bool`

HasSwaggerSpecs returns a boolean if a field has been set.

### SetSwaggerSpecsNil

`func (o *Module) SetSwaggerSpecsNil(b bool)`

 SetSwaggerSpecsNil sets the value for SwaggerSpecs to be an explicit nil

### UnsetSwaggerSpecs
`func (o *Module) UnsetSwaggerSpecs()`

UnsetSwaggerSpecs ensures that no value is present for SwaggerSpecs, not even an explicit nil
### GetUrl

`func (o *Module) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Module) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Module) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Module) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *Module) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *Module) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetAssemblyPaths

`func (o *Module) GetAssemblyPaths() []string`

GetAssemblyPaths returns the AssemblyPaths field if non-nil, zero value otherwise.

### GetAssemblyPathsOk

`func (o *Module) GetAssemblyPathsOk() (*[]string, bool)`

GetAssemblyPathsOk returns a tuple with the AssemblyPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyPaths

`func (o *Module) SetAssemblyPaths(v []string)`

SetAssemblyPaths sets AssemblyPaths field to given value.

### HasAssemblyPaths

`func (o *Module) HasAssemblyPaths() bool`

HasAssemblyPaths returns a boolean if a field has been set.

### SetAssemblyPathsNil

`func (o *Module) SetAssemblyPathsNil(b bool)`

 SetAssemblyPathsNil sets the value for AssemblyPaths to be an explicit nil

### UnsetAssemblyPaths
`func (o *Module) UnsetAssemblyPaths()`

UnsetAssemblyPaths ensures that no value is present for AssemblyPaths, not even an explicit nil
### GetMarkedForDeletion

`func (o *Module) GetMarkedForDeletion() bool`

GetMarkedForDeletion returns the MarkedForDeletion field if non-nil, zero value otherwise.

### GetMarkedForDeletionOk

`func (o *Module) GetMarkedForDeletionOk() (*bool, bool)`

GetMarkedForDeletionOk returns a tuple with the MarkedForDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedForDeletion

`func (o *Module) SetMarkedForDeletion(v bool)`

SetMarkedForDeletion sets MarkedForDeletion field to given value.

### HasMarkedForDeletion

`func (o *Module) HasMarkedForDeletion() bool`

HasMarkedForDeletion returns a boolean if a field has been set.

### GetVersion

`func (o *Module) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Module) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Module) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Module) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *Module) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *Module) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


