# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInstallerDeviceImage**](InstallerApi.md#AddInstallerDeviceImage) | **Post** /api/v1/installer/orgs/{org_id}/devices/{device_mac}/{image_name} | addInstallerDeviceImage
[**ClaimInstallerDevices**](InstallerApi.md#ClaimInstallerDevices) | **Post** /api/v1/installer/orgs/{org_id}/devices | claimInstallerDevices
[**CreateInstallerMap**](InstallerApi.md#CreateInstallerMap) | **Post** /api/v1/installer/orgs/{org_id}/sites/{site_name}/maps/{map_id} | createInstallerMap
[**CreateInstallerVirtualChassis**](InstallerApi.md#CreateInstallerVirtualChassis) | **Post** /api/v1/installer/orgs/{org_id}/devices/{fpc0_mac}/vc | createInstallerVirtualChassis
[**CreateOrUpdateInstallerSites**](InstallerApi.md#CreateOrUpdateInstallerSites) | **Put** /api/v1/installer/orgs/{org_id}/sites/{site_name} | createOrUpdateInstallerSites
[**DeleteInstallerDeviceImage**](InstallerApi.md#DeleteInstallerDeviceImage) | **Delete** /api/v1/installer/orgs/{org_id}/devices/{device_mac}/{image_name} | deleteInstallerDeviceImage
[**DeleteInstallerMap**](InstallerApi.md#DeleteInstallerMap) | **Delete** /api/v1/installer/orgs/{org_id}/sites/{site_name}/maps/{map_id} | deleteInstallerMap
[**GetInstallerDeviceVirtualChassis**](InstallerApi.md#GetInstallerDeviceVirtualChassis) | **Get** /api/v1/installer/orgs/{org_id}/devices/{fpc0_mac}/vc | getInstallerDeviceVirtualChassis
[**ImportInstallerMap**](InstallerApi.md#ImportInstallerMap) | **Post** /api/v1/installer/orgs/{org_id}/sites/{site_name}/maps/import | importInstallerMap
[**ListInstallerAlarmTemplates**](InstallerApi.md#ListInstallerAlarmTemplates) | **Get** /api/v1/installer/orgs/{org_id}/alarmtemplates | listInstallerAlarmTemplates
[**ListInstallerDeviceProfiles**](InstallerApi.md#ListInstallerDeviceProfiles) | **Get** /api/v1/installer/orgs/{org_id}/deviceprofiles | listInstallerDeviceProfiles
[**ListInstallerListOfRenctlyClaimedDevices**](InstallerApi.md#ListInstallerListOfRenctlyClaimedDevices) | **Get** /api/v1/installer/orgs/{org_id}/devices | listInstallerListOfRenctlyClaimedDevices
[**ListInstallerMaps**](InstallerApi.md#ListInstallerMaps) | **Get** /api/v1/installer/orgs/{org_id}/sites/{site_name}/maps | listInstallerMaps
[**ListInstallerRfTemplatesNames**](InstallerApi.md#ListInstallerRfTemplatesNames) | **Get** /api/v1/installer/orgs/{org_id}/rftemplates | listInstallerRfTemplatesNames
[**ListInstallerSiteGroups**](InstallerApi.md#ListInstallerSiteGroups) | **Get** /api/v1/installer/orgs/{org_id}/sitegroups | listInstallerSiteGroups
[**ListInstallerSites**](InstallerApi.md#ListInstallerSites) | **Get** /api/v1/installer/orgs/{org_id}/sites | listInstallerSites
[**OptimizeInstallerRrm**](InstallerApi.md#OptimizeInstallerRrm) | **Get** /api/v1/installer/sites/{site_name}/optimize | optimizeInstallerRrm
[**ProvisionInstallerDevices**](InstallerApi.md#ProvisionInstallerDevices) | **Put** /api/v1/installer/orgs/{org_id}/devices/{device_mac} | provisionInstallerDevices
[**StartInstallerLocateDevice**](InstallerApi.md#StartInstallerLocateDevice) | **Post** /api/v1/installer/orgs/{org_id}/devices/{device_mac}/locate | startInstallerLocateDevice
[**StopInstallerLocateDevice**](InstallerApi.md#StopInstallerLocateDevice) | **Post** /api/v1/installer/orgs/{org_id}/devices/{device_mac}/unlocate | stopInstallerLocateDevice
[**UnassignInstallerRecentlyClaimedDevice**](InstallerApi.md#UnassignInstallerRecentlyClaimedDevice) | **Delete** /api/v1/installer/orgs/{org_id}/devices/{device_mac} | unassignInstallerRecentlyClaimedDevice
[**UpdateInstallerMap**](InstallerApi.md#UpdateInstallerMap) | **Put** /api/v1/installer/orgs/{org_id}/sites/{site_name}/maps/{map_id} | updateInstallerMap
[**UpdateInstallerVirtualChassisMember**](InstallerApi.md#UpdateInstallerVirtualChassisMember) | **Put** /api/v1/installer/orgs/{org_id}/devices/{fpc0_mac}/vc | updateInstallerVirtualChassisMember

# **AddInstallerDeviceImage**
> AddInstallerDeviceImage(ctx, orgId, imageName, deviceMac, optional)
addInstallerDeviceImage

Add image

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **imageName** | **string**|  | 
  **deviceMac** | **string**|  | 
 **optional** | ***InstallerApiAddInstallerDeviceImageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstallerApiAddInstallerDeviceImageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **autoDeviceprofileAssignment** | **optional.**|  | 
 **csv** | **optional.Interface of *os.File****optional.**|  | 
 **file** | **optional.Interface of *os.File****optional.**|  | 
 **json** | [**optional.Interface of MapImportJson**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClaimInstallerDevices**
> InlineResponse2006 ClaimInstallerDevices(ctx, orgId, optional)
claimInstallerDevices

This mirrors `POST /api/v1/orgs/{org_id}/inventory` (see Inventory API)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***InstallerApiClaimInstallerDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstallerApiClaimInstallerDevicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of []string**](string.md)| Request Body | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateInstallerMap**
> InlineResponse2009 CreateInstallerMap(ctx, orgId, siteName, mapId, optional)
createInstallerMap

Create a MAP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **siteName** | **string**|  | 
  **mapId** | [**string**](.md)|  | 
 **optional** | ***InstallerApiCreateInstallerMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstallerApiCreateInstallerMapOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of MapsMapIdBody1**](MapsMapIdBody1.md)| Request Body | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateInstallerVirtualChassis**
> InlineResponse2007 CreateInstallerVirtualChassis(ctx, orgId, fpc0Mac, optional)
createInstallerVirtualChassis

For models (e.g. EX3400 and up) having dedicated VC ports, it is easier to form a VC by just connecting cables with the dedicated VC ports. Cloud will detect the new VC and update the inventory.  In case that the user would like to choose the dedicated switch as a VC master. Or for EX2300-C-12P and EX2300-C-12T which doesn’t have dedicated VC ports, below are procedures to automate the VC creation:  1. Power on the switch that is choosen as the VC master first. And the powering on the other member switches. 2. Claim or adopt all these switches under the same organization’s Inventory 3. Assign these switches into the same Site 4. Invoke vc command on the switch choosen to be the VC master. For EX2300-C-12P, VC ports will be created automatically. 5. Connect the cables to the VC ports for these switches 6. Wait for the VC to be formed. The Org’s inventory will be updated for the new VC.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **fpc0Mac** | **string**| FPC0 MAC Address | 
 **optional** | ***InstallerApiCreateInstallerVirtualChassisOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstallerApiCreateInstallerVirtualChassisOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of Fpc0MacVcBody1**](Fpc0MacVcBody1.md)| Request Body | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateInstallerSites**
> CreateOrUpdateInstallerSites(ctx, orgId, siteName, optional)
createOrUpdateInstallerSites

Often the Installers are asked to assign Devices to Sites. The Sites can either be pre-created or created/modified by the Installer. If this is an update, the same grace period also applies.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **siteName** | **string**|  | 
 **optional** | ***InstallerApiCreateOrUpdateInstallerSitesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstallerApiCreateOrUpdateInstallerSitesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SitesSiteNameBody**](SitesSiteNameBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInstallerDeviceImage**
> DeleteInstallerDeviceImage(ctx, orgId, imageName, deviceMac)
deleteInstallerDeviceImage

delete image

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **imageName** | **string**|  | 
  **deviceMac** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInstallerMap**
> DeleteInstallerMap(ctx, orgId, siteName, mapId)
deleteInstallerMap

Delete Map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **siteName** | **string**|  | 
  **mapId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInstallerDeviceVirtualChassis**
> InlineResponse2007 GetInstallerDeviceVirtualChassis(ctx, orgId, fpc0Mac)
getInstallerDeviceVirtualChassis

Get VC Status  The API returns a combined view of the VC status which includes topology and stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **fpc0Mac** | **string**| FPC0 MAC Address | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportInstallerMap**
> InlineResponse2008 ImportInstallerMap(ctx, orgId, siteName, optional)
importInstallerMap

Import data from files is a multipart POST which has an file, an optional json, and an optional csv, to create floorplan, assign & place ap if name or mac matches

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **siteName** | **string**|  | 
 **optional** | ***InstallerApiImportInstallerMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstallerApiImportInstallerMapOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **autoDeviceprofileAssignment** | **optional.**|  | 
 **csv** | **optional.Interface of *os.File****optional.**|  | 
 **file** | **optional.Interface of *os.File****optional.**|  | 
 **json** | [**optional.Interface of MapImportJson**](.md)|  | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInstallerAlarmTemplates**
> []InstallersItem ListInstallerAlarmTemplates(ctx, orgId)
listInstallerAlarmTemplates

Get List of alarm templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**[]InstallersItem**](installers_item.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInstallerDeviceProfiles**
> []InstallersItem ListInstallerDeviceProfiles(ctx, orgId, optional)
listInstallerDeviceProfiles

Get List of Device Profiles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***InstallerApiListInstallerDeviceProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstallerApiListInstallerDeviceProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**optional.Interface of DeviceType**](.md)|  | 

### Return type

[**[]InstallersItem**](installers_item.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInstallerListOfRenctlyClaimedDevices**
> []InstallerDevice ListInstallerListOfRenctlyClaimedDevices(ctx, orgId, optional)
listInstallerListOfRenctlyClaimedDevices

Get List of recently claimed devices

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***InstallerApiListInstallerListOfRenctlyClaimedDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstallerApiListInstallerListOfRenctlyClaimedDevicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | **optional.String**| Device Model | 
 **siteName** | **optional.String**| Site Name | 
 **siteId** | [**optional.Interface of string**](.md)| Site ID | 
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]InstallerDevice**](installer_device.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInstallerMaps**
> []ModelMap ListInstallerMaps(ctx, orgId, siteName)
listInstallerMaps

Get List of Maps

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **siteName** | **string**|  | 

### Return type

[**[]ModelMap**](map.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInstallerRfTemplatesNames**
> []InstallersItem ListInstallerRfTemplatesNames(ctx, orgId)
listInstallerRfTemplatesNames

Get List of RF Templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**[]InstallersItem**](installers_item.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInstallerSiteGroups**
> []InstallersItem ListInstallerSiteGroups(ctx, orgId)
listInstallerSiteGroups

Get List of Site Groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**[]InstallersItem**](installers_item.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInstallerSites**
> []InstallerSite ListInstallerSites(ctx, orgId)
listInstallerSites

Get List of Sites

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**[]InstallerSite**](installer_site.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OptimizeInstallerRrm**
> OptimizeInstallerRrm(ctx, siteName)
optimizeInstallerRrm

After installation is considered complete (APs are placed on maps, all powered up), you can trigger an optimize operation where RRM will kick in (and maybe other things in the future) before it’s automatically scheduled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProvisionInstallerDevices**
> ProvisionInstallerDevices(ctx, orgId, deviceMac, optional)
provisionInstallerDevices

Provision or Replace a device   If replacing_mac is in the request payload, other attributes are ignored, we attempt to replace existing device (with mac replacing_mac) with the inventory device being configured. The replacement device must be in the inventory but not assigned, and the replacing_mac device must be assigned to a site, and satisfy grace period requirements. The Device replaced will become unassigned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **deviceMac** | **string**|  | 
 **optional** | ***InstallerApiProvisionInstallerDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstallerApiProvisionInstallerDevicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DevicesDeviceMacBody**](DevicesDeviceMacBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartInstallerLocateDevice**
> StartInstallerLocateDevice(ctx, orgId, deviceMac)
startInstallerLocateDevice

Locate a Device by blinking it’s LED, it’s a persisted state that has to be stopped by calling Stop Locating API

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **deviceMac** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopInstallerLocateDevice**
> StopInstallerLocateDevice(ctx, orgId, deviceMac)
stopInstallerLocateDevice

Stop it

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **deviceMac** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnassignInstallerRecentlyClaimedDevice**
> UnassignInstallerRecentlyClaimedDevice(ctx, orgId, deviceMac)
unassignInstallerRecentlyClaimedDevice

Unassign recently claimed devices

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **deviceMac** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateInstallerMap**
> InlineResponse2009 UpdateInstallerMap(ctx, orgId, siteName, mapId, optional)
updateInstallerMap

Update map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **siteName** | **string**|  | 
  **mapId** | [**string**](.md)|  | 
 **optional** | ***InstallerApiUpdateInstallerMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstallerApiUpdateInstallerMapOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of MapsMapIdBody**](MapsMapIdBody.md)| Request Body | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateInstallerVirtualChassisMember**
> InlineResponse2007 UpdateInstallerVirtualChassisMember(ctx, orgId, fpc0Mac, optional)
updateInstallerVirtualChassisMember

The VC creation and adding member switch API will update the device’ s virtual chassis config which is applied after VC is formed to create JUNOS pre-provisioned virtual chassis configuration.  ## Change to use preprovisioned VC To switch the VC to use preprovisioned VC, enable preprovisioned in virtual_chassis config. Both vc_role master and backup will be matched to routing-engine role in Junos preprovisioned VC config.  In this config, fpc0 has to be the same as the mac of device_id. Use renumber if you want to replace fpc0 which involves device_id change.  Notice: to configure preprovisioned VC, every member of the VC must be in the inventory.  ## Add new members For models (e.g. EX4300 and up) having dedicated VC ports, it is easier to add new member switches into a VC by just connecting cables with the dedicated VC ports. Cloud will detect the new members and update the inventory.  For EX2300 VC, adding new members requires to follow the procedures below: 1. Powering on the new member switches and ensuring cables are not connected to any VC ports. 2. Claim or adopt all new member switches under the VC’s organization Inventory 3. Assign all new member switches to the same Site as the VC 4. Invoke vc command to add switches to the VC. 5. Connect the cables to the VC ports for these switches 6. After a while, the Org’s Inventory shows this new switches has been added into the VC.  ## Removing member switch To remove a member switch from the VC, following the procedures below:  1. Ensuring the VC is connected to the cloud first 2. Unplug the cable from the VC port of the switch 3. Waiting for the VC state (vc_state) of this switch is changed to not-present 4. Invoke update_vc with remove to remove this switch from the VC 5. The Org’s Inventory shows the switch is removed.  Please notice that member ID 0 (fpc0) cannot be removed. When a VC has two switches left, unpluging the cable may result in the situation that fpc0 becomes a line card (LC). When this situation is happened, please re-plug in the cable, wait for both switches becoming present (show virtual-chassis) and then removing the cable again.  ## Renumber a member switch When a member switch doesn’ t work properly and needed to be replaced, the renumber API could be used. The following two types of renumber are supported:  1. Replace a non-fpc0 member switch 2. Replace fpc0. When fpc0 is relaced, PAPI device config and JUNOS config will be both updated.  For renumber to work, the following procedures are needed:  1. Ensuring the VC is connected to the cloud and the state of the member switch to be replaced must be non present.  2. Adding the new member switch to the VC  3. Waiting for the VC state (vc_state) of this VC to be updated to API server  4. Invoke vc with renumber to replace the new member switch from fpc X to  ## Perprovision VC members By specifying “preprovision” op, you can convert the current VC to pre-provisioned mode, update VC members as well as specify vc_ports when adding new members for device models without dedicated vc ports. Use renumber for fpc0 replacement which involves device_id change.  Note:  1. vc_ports is used for adding new members and not needed if * the device model has dedicated vc ports, or * no new member is added  2. New VC members to be added should exist in the same Site as the VC  Update Device’s VC config can achieve similar purpose by directly modifying current virtual_chassis config. However, it cannot fulfill requests to enabling vc_ports on new members that are yet to belong to current VC.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **fpc0Mac** | **string**| FPC0 MAC Address | 
 **optional** | ***InstallerApiUpdateInstallerVirtualChassisMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstallerApiUpdateInstallerVirtualChassisMemberOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of Fpc0MacVcBody**](Fpc0MacVcBody.md)| Request Body | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

