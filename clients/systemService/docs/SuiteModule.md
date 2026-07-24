# SuiteModule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**MarkedForDeletion** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] [readonly] 
**Order** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**FullName** | Pointer to **NullableString** |  | [optional] [readonly] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Configuration** | Pointer to **NullableString** |  | [optional] 
**Author** | Pointer to **NullableString** |  | [optional] 
**AuthorUrl** | Pointer to **NullableString** |  | [optional] 
**License** | Pointer to **NullableString** |  | [optional] 
**RequireLicenseAcceptance** | Pointer to **NullableBool** |  | [optional] 
**Repository** | Pointer to **NullableString** |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**Image** | Pointer to **NullableString** |  | [optional] 
**NuSpecPath** | Pointer to **NullableString** |  | [optional] 
**Manifest** | Pointer to **NullableString** |  | [optional] 
**Logo** | Pointer to **NullableString** |  | [optional] 
**Website** | Pointer to **NullableString** |  | [optional] 
**Documentation** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**OpenApiDefinitionSpec** | Pointer to [**IOpenApiDefinitionSpec**](IOpenApiDefinitionSpec.md) |  | [optional] 
**SwaggerSpecs** | Pointer to [**[]IOpenApiDefinitionSpec**](IOpenApiDefinitionSpec.md) |  | [optional] 
**AssemblyPaths** | Pointer to **[]string** |  | [optional] 
**RequiredPermissions** | Pointer to **[]string** |  | [optional] [readonly] 
**Version** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSuiteModule

`func NewSuiteModule() *SuiteModule`

NewSuiteModule instantiates a new SuiteModule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuiteModuleWithDefaults

`func NewSuiteModuleWithDefaults() *SuiteModule`

NewSuiteModuleWithDefaults instantiates a new SuiteModule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *SuiteModule) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *SuiteModule) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *SuiteModule) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *SuiteModule) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetMarkedForDeletion

`func (o *SuiteModule) GetMarkedForDeletion() bool`

GetMarkedForDeletion returns the MarkedForDeletion field if non-nil, zero value otherwise.

### GetMarkedForDeletionOk

`func (o *SuiteModule) GetMarkedForDeletionOk() (*bool, bool)`

GetMarkedForDeletionOk returns a tuple with the MarkedForDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedForDeletion

`func (o *SuiteModule) SetMarkedForDeletion(v bool)`

SetMarkedForDeletion sets MarkedForDeletion field to given value.

### HasMarkedForDeletion

`func (o *SuiteModule) HasMarkedForDeletion() bool`

HasMarkedForDeletion returns a boolean if a field has been set.

### GetActive

`func (o *SuiteModule) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SuiteModule) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SuiteModule) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SuiteModule) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetOrder

`func (o *SuiteModule) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *SuiteModule) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *SuiteModule) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *SuiteModule) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetId

`func (o *SuiteModule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SuiteModule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SuiteModule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SuiteModule) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SuiteModule) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SuiteModule) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *SuiteModule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SuiteModule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SuiteModule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SuiteModule) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SuiteModule) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SuiteModule) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFullName

`func (o *SuiteModule) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *SuiteModule) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *SuiteModule) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *SuiteModule) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *SuiteModule) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *SuiteModule) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetDescription

`func (o *SuiteModule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SuiteModule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SuiteModule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SuiteModule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SuiteModule) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SuiteModule) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *SuiteModule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SuiteModule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SuiteModule) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SuiteModule) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCategory

`func (o *SuiteModule) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SuiteModule) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SuiteModule) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SuiteModule) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetConfiguration

`func (o *SuiteModule) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *SuiteModule) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *SuiteModule) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *SuiteModule) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *SuiteModule) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *SuiteModule) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetAuthor

`func (o *SuiteModule) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *SuiteModule) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *SuiteModule) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *SuiteModule) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *SuiteModule) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *SuiteModule) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetAuthorUrl

`func (o *SuiteModule) GetAuthorUrl() string`

GetAuthorUrl returns the AuthorUrl field if non-nil, zero value otherwise.

### GetAuthorUrlOk

`func (o *SuiteModule) GetAuthorUrlOk() (*string, bool)`

GetAuthorUrlOk returns a tuple with the AuthorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorUrl

`func (o *SuiteModule) SetAuthorUrl(v string)`

SetAuthorUrl sets AuthorUrl field to given value.

### HasAuthorUrl

`func (o *SuiteModule) HasAuthorUrl() bool`

HasAuthorUrl returns a boolean if a field has been set.

### SetAuthorUrlNil

`func (o *SuiteModule) SetAuthorUrlNil(b bool)`

 SetAuthorUrlNil sets the value for AuthorUrl to be an explicit nil

### UnsetAuthorUrl
`func (o *SuiteModule) UnsetAuthorUrl()`

UnsetAuthorUrl ensures that no value is present for AuthorUrl, not even an explicit nil
### GetLicense

`func (o *SuiteModule) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *SuiteModule) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *SuiteModule) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *SuiteModule) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### SetLicenseNil

`func (o *SuiteModule) SetLicenseNil(b bool)`

 SetLicenseNil sets the value for License to be an explicit nil

### UnsetLicense
`func (o *SuiteModule) UnsetLicense()`

UnsetLicense ensures that no value is present for License, not even an explicit nil
### GetRequireLicenseAcceptance

`func (o *SuiteModule) GetRequireLicenseAcceptance() bool`

GetRequireLicenseAcceptance returns the RequireLicenseAcceptance field if non-nil, zero value otherwise.

### GetRequireLicenseAcceptanceOk

`func (o *SuiteModule) GetRequireLicenseAcceptanceOk() (*bool, bool)`

GetRequireLicenseAcceptanceOk returns a tuple with the RequireLicenseAcceptance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireLicenseAcceptance

`func (o *SuiteModule) SetRequireLicenseAcceptance(v bool)`

SetRequireLicenseAcceptance sets RequireLicenseAcceptance field to given value.

### HasRequireLicenseAcceptance

`func (o *SuiteModule) HasRequireLicenseAcceptance() bool`

HasRequireLicenseAcceptance returns a boolean if a field has been set.

### SetRequireLicenseAcceptanceNil

`func (o *SuiteModule) SetRequireLicenseAcceptanceNil(b bool)`

 SetRequireLicenseAcceptanceNil sets the value for RequireLicenseAcceptance to be an explicit nil

### UnsetRequireLicenseAcceptance
`func (o *SuiteModule) UnsetRequireLicenseAcceptance()`

UnsetRequireLicenseAcceptance ensures that no value is present for RequireLicenseAcceptance, not even an explicit nil
### GetRepository

`func (o *SuiteModule) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *SuiteModule) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *SuiteModule) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *SuiteModule) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *SuiteModule) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *SuiteModule) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetIcon

`func (o *SuiteModule) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *SuiteModule) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *SuiteModule) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *SuiteModule) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *SuiteModule) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *SuiteModule) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetImage

`func (o *SuiteModule) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *SuiteModule) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *SuiteModule) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *SuiteModule) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *SuiteModule) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *SuiteModule) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetNuSpecPath

`func (o *SuiteModule) GetNuSpecPath() string`

GetNuSpecPath returns the NuSpecPath field if non-nil, zero value otherwise.

### GetNuSpecPathOk

`func (o *SuiteModule) GetNuSpecPathOk() (*string, bool)`

GetNuSpecPathOk returns a tuple with the NuSpecPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuSpecPath

`func (o *SuiteModule) SetNuSpecPath(v string)`

SetNuSpecPath sets NuSpecPath field to given value.

### HasNuSpecPath

`func (o *SuiteModule) HasNuSpecPath() bool`

HasNuSpecPath returns a boolean if a field has been set.

### SetNuSpecPathNil

`func (o *SuiteModule) SetNuSpecPathNil(b bool)`

 SetNuSpecPathNil sets the value for NuSpecPath to be an explicit nil

### UnsetNuSpecPath
`func (o *SuiteModule) UnsetNuSpecPath()`

UnsetNuSpecPath ensures that no value is present for NuSpecPath, not even an explicit nil
### GetManifest

`func (o *SuiteModule) GetManifest() string`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *SuiteModule) GetManifestOk() (*string, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *SuiteModule) SetManifest(v string)`

SetManifest sets Manifest field to given value.

### HasManifest

`func (o *SuiteModule) HasManifest() bool`

HasManifest returns a boolean if a field has been set.

### SetManifestNil

`func (o *SuiteModule) SetManifestNil(b bool)`

 SetManifestNil sets the value for Manifest to be an explicit nil

### UnsetManifest
`func (o *SuiteModule) UnsetManifest()`

UnsetManifest ensures that no value is present for Manifest, not even an explicit nil
### GetLogo

`func (o *SuiteModule) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *SuiteModule) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *SuiteModule) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *SuiteModule) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### SetLogoNil

`func (o *SuiteModule) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *SuiteModule) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil
### GetWebsite

`func (o *SuiteModule) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *SuiteModule) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *SuiteModule) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *SuiteModule) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### SetWebsiteNil

`func (o *SuiteModule) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *SuiteModule) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil
### GetDocumentation

`func (o *SuiteModule) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *SuiteModule) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *SuiteModule) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *SuiteModule) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### SetDocumentationNil

`func (o *SuiteModule) SetDocumentationNil(b bool)`

 SetDocumentationNil sets the value for Documentation to be an explicit nil

### UnsetDocumentation
`func (o *SuiteModule) UnsetDocumentation()`

UnsetDocumentation ensures that no value is present for Documentation, not even an explicit nil
### GetUrl

`func (o *SuiteModule) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SuiteModule) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SuiteModule) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SuiteModule) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *SuiteModule) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *SuiteModule) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetPath

`func (o *SuiteModule) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SuiteModule) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SuiteModule) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SuiteModule) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *SuiteModule) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *SuiteModule) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetOpenApiDefinitionSpec

`func (o *SuiteModule) GetOpenApiDefinitionSpec() IOpenApiDefinitionSpec`

GetOpenApiDefinitionSpec returns the OpenApiDefinitionSpec field if non-nil, zero value otherwise.

### GetOpenApiDefinitionSpecOk

`func (o *SuiteModule) GetOpenApiDefinitionSpecOk() (*IOpenApiDefinitionSpec, bool)`

GetOpenApiDefinitionSpecOk returns a tuple with the OpenApiDefinitionSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenApiDefinitionSpec

`func (o *SuiteModule) SetOpenApiDefinitionSpec(v IOpenApiDefinitionSpec)`

SetOpenApiDefinitionSpec sets OpenApiDefinitionSpec field to given value.

### HasOpenApiDefinitionSpec

`func (o *SuiteModule) HasOpenApiDefinitionSpec() bool`

HasOpenApiDefinitionSpec returns a boolean if a field has been set.

### GetSwaggerSpecs

`func (o *SuiteModule) GetSwaggerSpecs() []IOpenApiDefinitionSpec`

GetSwaggerSpecs returns the SwaggerSpecs field if non-nil, zero value otherwise.

### GetSwaggerSpecsOk

`func (o *SuiteModule) GetSwaggerSpecsOk() (*[]IOpenApiDefinitionSpec, bool)`

GetSwaggerSpecsOk returns a tuple with the SwaggerSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwaggerSpecs

`func (o *SuiteModule) SetSwaggerSpecs(v []IOpenApiDefinitionSpec)`

SetSwaggerSpecs sets SwaggerSpecs field to given value.

### HasSwaggerSpecs

`func (o *SuiteModule) HasSwaggerSpecs() bool`

HasSwaggerSpecs returns a boolean if a field has been set.

### SetSwaggerSpecsNil

`func (o *SuiteModule) SetSwaggerSpecsNil(b bool)`

 SetSwaggerSpecsNil sets the value for SwaggerSpecs to be an explicit nil

### UnsetSwaggerSpecs
`func (o *SuiteModule) UnsetSwaggerSpecs()`

UnsetSwaggerSpecs ensures that no value is present for SwaggerSpecs, not even an explicit nil
### GetAssemblyPaths

`func (o *SuiteModule) GetAssemblyPaths() []string`

GetAssemblyPaths returns the AssemblyPaths field if non-nil, zero value otherwise.

### GetAssemblyPathsOk

`func (o *SuiteModule) GetAssemblyPathsOk() (*[]string, bool)`

GetAssemblyPathsOk returns a tuple with the AssemblyPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyPaths

`func (o *SuiteModule) SetAssemblyPaths(v []string)`

SetAssemblyPaths sets AssemblyPaths field to given value.

### HasAssemblyPaths

`func (o *SuiteModule) HasAssemblyPaths() bool`

HasAssemblyPaths returns a boolean if a field has been set.

### SetAssemblyPathsNil

`func (o *SuiteModule) SetAssemblyPathsNil(b bool)`

 SetAssemblyPathsNil sets the value for AssemblyPaths to be an explicit nil

### UnsetAssemblyPaths
`func (o *SuiteModule) UnsetAssemblyPaths()`

UnsetAssemblyPaths ensures that no value is present for AssemblyPaths, not even an explicit nil
### GetRequiredPermissions

`func (o *SuiteModule) GetRequiredPermissions() []string`

GetRequiredPermissions returns the RequiredPermissions field if non-nil, zero value otherwise.

### GetRequiredPermissionsOk

`func (o *SuiteModule) GetRequiredPermissionsOk() (*[]string, bool)`

GetRequiredPermissionsOk returns a tuple with the RequiredPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredPermissions

`func (o *SuiteModule) SetRequiredPermissions(v []string)`

SetRequiredPermissions sets RequiredPermissions field to given value.

### HasRequiredPermissions

`func (o *SuiteModule) HasRequiredPermissions() bool`

HasRequiredPermissions returns a boolean if a field has been set.

### SetRequiredPermissionsNil

`func (o *SuiteModule) SetRequiredPermissionsNil(b bool)`

 SetRequiredPermissionsNil sets the value for RequiredPermissions to be an explicit nil

### UnsetRequiredPermissions
`func (o *SuiteModule) UnsetRequiredPermissions()`

UnsetRequiredPermissions ensures that no value is present for RequiredPermissions, not even an explicit nil
### GetVersion

`func (o *SuiteModule) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SuiteModule) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SuiteModule) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SuiteModule) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *SuiteModule) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *SuiteModule) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


