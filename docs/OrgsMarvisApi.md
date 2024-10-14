# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TroubleshootOrg**](OrgsMarvisApi.md#TroubleshootOrg) | **Get** /api/v1/orgs/{org_id}/troubleshoot | troubleshootOrg

# **TroubleshootOrg**
> InlineResponse20065 TroubleshootOrg(ctx, orgId, optional)
troubleshootOrg

Troubleshoot sites, devices, clients, and wired clientsfor maximum of last 7 days from current time. See search APIs for device information: - [search Device]($e/Orgs%20Devices/searchOrgDevices) - [search Wireless Client]($e/Orgs%20Clients%20-%20Wireless/searchOrgWirelessClients) - [search Wired Client]($e/Orgs%20Clients%20-%20Wired/searchOrgWiredClients) - [search Wan Client]($e/Orgs%20Clients%20-%20Wan/searchOrgWanClients)  **NOTE**: requires Marvis subscription license

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsMarvisApiTroubleshootOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsMarvisApiTroubleshootOrgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mac** | **optional.String**| **required** when troubleshooting device or a client | 
 **siteId** | [**optional.Interface of string**](.md)| **required** when troubleshooting site | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **type_** | [**optional.Interface of Type1**](.md)| when troubleshooting site, type of network to troubleshoot | 

### Return type

[**InlineResponse20065**](inline_response_200_65.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

