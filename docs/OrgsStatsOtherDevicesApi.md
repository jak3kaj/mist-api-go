# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgOtherDeviceStats**](OrgsStatsOtherDevicesApi.md#GetOrgOtherDeviceStats) | **Get** /api/v1/orgs/{org_id}/stats/otherdevices/{device_mac} | getOrgOtherDeviceStats

# **GetOrgOtherDeviceStats**
> InlineResponse20078 GetOrgOtherDeviceStats(ctx, orgId, deviceMac)
getOrgOtherDeviceStats

Get Otherdevice Stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **deviceMac** | **string**|  | 

### Return type

[**InlineResponse20078**](inline_response_200_78.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

