# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelOrgSsrUpgrade**](UtilitiesUpgradeApi.md#CancelOrgSsrUpgrade) | **Post** /api/v1/orgs/{org_id}/ssr/upgrade/{upgrade_id}/cancel | cancelOrgSsrUpgrade
[**CancelSiteDeviceUpgrade**](UtilitiesUpgradeApi.md#CancelSiteDeviceUpgrade) | **Post** /api/v1/sites/{site_id}/devices/upgrade/{upgrade_id}/cancel | cancelSiteDeviceUpgrade
[**GetOrgDeviceUpgrade**](UtilitiesUpgradeApi.md#GetOrgDeviceUpgrade) | **Get** /api/v1/orgs/{org_id}/devices/upgrade/{upgrade_id} | getOrgDeviceUpgrade
[**GetOrgMxEdgeUpgrade**](UtilitiesUpgradeApi.md#GetOrgMxEdgeUpgrade) | **Get** /api/v1/orgs/{org_id}/mxedges/upgrade/{upgrade_id} | getOrgMxEdgeUpgrade
[**GetSiteDeviceUpgrade**](UtilitiesUpgradeApi.md#GetSiteDeviceUpgrade) | **Get** /api/v1/sites/{site_id}/devices/upgrade/{upgrade_id} | getSiteDeviceUpgrade
[**GetSiteSsrUpgrade**](UtilitiesUpgradeApi.md#GetSiteSsrUpgrade) | **Get** /api/v1/sites/{site_id}/ssr/upgrade/{upgrade_id} | getSiteSsrUpgrade
[**ListOrgAvailableDeviceVersions**](UtilitiesUpgradeApi.md#ListOrgAvailableDeviceVersions) | **Get** /api/v1/orgs/{org_id}/devices/versions | listOrgAvailableDeviceVersions
[**ListOrgAvailableSsrVersions**](UtilitiesUpgradeApi.md#ListOrgAvailableSsrVersions) | **Get** /api/v1/orgs/{org_id}/ssr/versions | listOrgAvailableSsrVersions
[**ListOrgDeviceUpgrades**](UtilitiesUpgradeApi.md#ListOrgDeviceUpgrades) | **Get** /api/v1/orgs/{org_id}/devices/upgrade | listOrgDeviceUpgrades
[**ListOrgMxEdgeUpgrades**](UtilitiesUpgradeApi.md#ListOrgMxEdgeUpgrades) | **Get** /api/v1/orgs/{org_id}/mxedges/upgrade | listOrgMxEdgeUpgrades
[**ListOrgSsrUpgrades**](UtilitiesUpgradeApi.md#ListOrgSsrUpgrades) | **Get** /api/v1/orgs/{org_id}/ssr/upgrade | listOrgSsrUpgrades
[**ListSiteAvailableDeviceVersions**](UtilitiesUpgradeApi.md#ListSiteAvailableDeviceVersions) | **Get** /api/v1/sites/{site_id}/devices/versions | listSiteAvailableDeviceVersions
[**ListSiteDeviceUpgrades**](UtilitiesUpgradeApi.md#ListSiteDeviceUpgrades) | **Get** /api/v1/sites/{site_id}/devices/upgrade | listSiteDeviceUpgrades
[**UpgradeDevice**](UtilitiesUpgradeApi.md#UpgradeDevice) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/upgrade | upgradeDevice
[**UpgradeOrgDevices**](UtilitiesUpgradeApi.md#UpgradeOrgDevices) | **Post** /api/v1/orgs/{org_id}/devices/upgrade | upgradeOrgDevices
[**UpgradeOrgJsiDevice**](UtilitiesUpgradeApi.md#UpgradeOrgJsiDevice) | **Post** /api/v1/orgs/{org_id}/jsi/devices/{device_mac}/upgrade | upgradeOrgJsiDevice
[**UpgradeOrgMxEdges**](UtilitiesUpgradeApi.md#UpgradeOrgMxEdges) | **Post** /api/v1/orgs/{org_id}/mxedges/upgrade | upgradeOrgMxEdges
[**UpgradeOrgSsrs**](UtilitiesUpgradeApi.md#UpgradeOrgSsrs) | **Post** /api/v1/orgs/{org_id}/ssr/upgrade | upgradeOrgSsrs
[**UpgradeSiteDevices**](UtilitiesUpgradeApi.md#UpgradeSiteDevices) | **Post** /api/v1/sites/{site_id}/devices/upgrade | upgradeSiteDevices
[**UpgradeSsr**](UtilitiesUpgradeApi.md#UpgradeSsr) | **Post** /api/v1/sites/{site_id}/ssr/{device_id}/upgrade | upgradeSsr

# **CancelOrgSsrUpgrade**
> CancelOrgSsrUpgrade(ctx, orgId, upgradeId)
cancelOrgSsrUpgrade

Best effort to cancel an upgrade. Devices which are already upgraded wont be touched↵ 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **upgradeId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelSiteDeviceUpgrade**
> CancelSiteDeviceUpgrade(ctx, siteId, upgradeId)
cancelSiteDeviceUpgrade

Best effort to cancel an upgrade. Devices which are already upgraded wont be touched

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **upgradeId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgDeviceUpgrade**
> InlineResponse200195 GetOrgDeviceUpgrade(ctx, orgId, upgradeId)
getOrgDeviceUpgrade

Get Multiple Devices Upgrade

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **upgradeId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200195**](inline_response_200_195.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgMxEdgeUpgrade**
> ResponseMxedgeUpgrade GetOrgMxEdgeUpgrade(ctx, orgId, upgradeId)
getOrgMxEdgeUpgrade

Get Mist Edge Upgrade

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **upgradeId** | [**string**](.md)|  | 

### Return type

[**ResponseMxedgeUpgrade**](response_mxedge_upgrade.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteDeviceUpgrade**
> InlineResponse200198 GetSiteDeviceUpgrade(ctx, siteId, upgradeId)
getSiteDeviceUpgrade

Get Site Device Upgrade

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **upgradeId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200198**](inline_response_200_198.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteSsrUpgrade**
> InlineResponse200200 GetSiteSsrUpgrade(ctx, siteId, upgradeId)
getSiteSsrUpgrade

Get Specific Site SSR Upgrade

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **upgradeId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200200**](inline_response_200_200.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgAvailableDeviceVersions**
> []DeviceVersionItem ListOrgAvailableDeviceVersions(ctx, orgId, optional)
listOrgAvailableDeviceVersions

Get List of Available Device Versions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesUpgradeApiListOrgAvailableDeviceVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesUpgradeApiListOrgAvailableDeviceVersionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**optional.Interface of DeviceType**](.md)|  | 
 **model** | **optional.String**| fetch version for device model, use/combine with &#x60;type&#x60; as needed (for switch and gateway devices) | 

### Return type

[**[]DeviceVersionItem**](device_version_item.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgAvailableSsrVersions**
> []SsrVersion ListOrgAvailableSsrVersions(ctx, orgId, optional)
listOrgAvailableSsrVersions

Get available version for SSR

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesUpgradeApiListOrgAvailableSsrVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesUpgradeApiListOrgAvailableSsrVersionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **channel** | **optional.String**|  | 

### Return type

[**[]SsrVersion**](ssr_version.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgDeviceUpgrades**
> []OrgDeviceUpgrade ListOrgDeviceUpgrades(ctx, orgId)
listOrgDeviceUpgrades

Get List of Org multiple devces upgrades

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**[]OrgDeviceUpgrade**](org_device_upgrade.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgMxEdgeUpgrades**
> []ResponseMxedgeUpgrade ListOrgMxEdgeUpgrades(ctx, orgId)
listOrgMxEdgeUpgrades

Get List of Org Mist Edge Upgrades

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**[]ResponseMxedgeUpgrade**](response_mxedge_upgrade.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgSsrUpgrades**
> []SsrUpgradeResponse ListOrgSsrUpgrades(ctx, orgId)
listOrgSsrUpgrades

Get List of Org SSR Upgrades

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**[]SsrUpgradeResponse**](ssr_upgrade_response.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteAvailableDeviceVersions**
> []DeviceVersionItem ListSiteAvailableDeviceVersions(ctx, siteId, optional)
listSiteAvailableDeviceVersions

Get List of Available Device Versions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesUpgradeApiListSiteAvailableDeviceVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesUpgradeApiListSiteAvailableDeviceVersionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**optional.Interface of DeviceType**](.md)|  | 
 **model** | **optional.String**| fetch version for device model, use/combine with &#x60;type&#x60; as needed (for switch and gateway devices) | 

### Return type

[**[]DeviceVersionItem**](device_version_item.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteDeviceUpgrades**
> []ResponseSiteDeviceUpgrade ListSiteDeviceUpgrades(ctx, siteId, optional)
listSiteDeviceUpgrades

Get all upgrades for site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesUpgradeApiListSiteDeviceUpgradesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesUpgradeApiListSiteDeviceUpgradesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | [**optional.Interface of DeviceUpgradeStatus**](.md)|  | 

### Return type

[**[]ResponseSiteDeviceUpgrade**](response_site_device_upgrade.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpgradeDevice**
> InlineResponse200199 UpgradeDevice(ctx, siteId, deviceId, optional)
upgradeDevice

Device Upgrade

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesUpgradeApiUpgradeDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesUpgradeApiUpgradeDeviceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdUpgradeBody**](DeviceIdUpgradeBody.md)|  | 

### Return type

[**InlineResponse200199**](inline_response_200_199.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpgradeOrgDevices**
> InlineResponse200194 UpgradeOrgDevices(ctx, orgId, optional)
upgradeOrgDevices

Upgrade Multiple Sites (Only supported for Access Points ugprades)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesUpgradeApiUpgradeOrgDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesUpgradeApiUpgradeOrgDevicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpgradeOrgDevices**](UpgradeOrgDevices.md)|  | 

### Return type

[**InlineResponse200194**](inline_response_200_194.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpgradeOrgJsiDevice**
> UpgradeOrgJsiDevice(ctx, orgId, deviceMac, optional)
upgradeOrgJsiDevice

Upgrade

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **deviceMac** | **string**|  | 
 **optional** | ***UtilitiesUpgradeApiUpgradeOrgJsiDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesUpgradeApiUpgradeOrgJsiDeviceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceMacUpgradeBody**](DeviceMacUpgradeBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpgradeOrgMxEdges**
> UpgradeOrgMxEdges(ctx, orgId, optional)
upgradeOrgMxEdges

Upgrade Mist Edges

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesUpgradeApiUpgradeOrgMxEdgesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesUpgradeApiUpgradeOrgMxEdgesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MxedgesUpgradeBody**](MxedgesUpgradeBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpgradeOrgSsrs**
> InlineResponse200196 UpgradeOrgSsrs(ctx, orgId, optional)
upgradeOrgSsrs

Upgrade Org SSRs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesUpgradeApiUpgradeOrgSsrsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesUpgradeApiUpgradeOrgSsrsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SsrUpgradeBody**](SsrUpgradeBody.md)|  | 

### Return type

[**InlineResponse200196**](inline_response_200_196.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpgradeSiteDevices**
> InlineResponse200197 UpgradeSiteDevices(ctx, siteId, optional)
upgradeSiteDevices

Upgrade Site Device  **Note**: this call doesn’t guarantee the devices to be upgraded right away (they may be offline)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesUpgradeApiUpgradeSiteDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesUpgradeApiUpgradeSiteDevicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DevicesUpgradeBody**](DevicesUpgradeBody.md)| Request Body | 

### Return type

[**InlineResponse200197**](inline_response_200_197.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpgradeSsr**
> InlineResponse200196 UpgradeSsr(ctx, siteId, deviceId, optional)
upgradeSsr

Upgrade Site SSR device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesUpgradeApiUpgradeSsrOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesUpgradeApiUpgradeSsrOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdUpgradeBody1**](DeviceIdUpgradeBody1.md)|  | 

### Return type

[**InlineResponse200196**](inline_response_200_196.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

