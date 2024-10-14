# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendSiteDevicesArbitratryBleBeacon**](UtilitiesLocationApi.md#SendSiteDevicesArbitratryBleBeacon) | **Post** /api/v1/sites/{site_id}/devices/send_ble_beacon | sendSiteDevicesArbitratryBleBeacon

# **SendSiteDevicesArbitratryBleBeacon**
> SendSiteDevicesArbitratryBleBeacon(ctx, siteId, optional)
sendSiteDevicesArbitratryBleBeacon

Send arbitrary BLE Beacon for a period of time  Note that only the devices that are connected will be restarted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesLocationApiSendSiteDevicesArbitratryBleBeaconOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesLocationApiSendSiteDevicesArbitratryBleBeaconOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DevicesSendBleBeaconBody**](DevicesSendBleBeaconBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

