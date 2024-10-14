# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgCapturingStatus**](UtilitiesPCAPsApi.md#GetOrgCapturingStatus) | **Get** /api/v1/orgs/{org_id}/pcaps/capture | getOrgCapturingStatus
[**GetSiteCapturingStatus**](UtilitiesPCAPsApi.md#GetSiteCapturingStatus) | **Get** /api/v1/sites/{site_id}/pcaps/capture | getSiteCapturingStatus
[**ListOrgPacketCaptures**](UtilitiesPCAPsApi.md#ListOrgPacketCaptures) | **Get** /api/v1/orgs/{org_id}/pcaps | listOrgPacketCaptures
[**ListSitePacketCaptures**](UtilitiesPCAPsApi.md#ListSitePacketCaptures) | **Get** /api/v1/sites/{site_id}/pcaps | listSitePacketCaptures
[**StartOrgPacketCapture**](UtilitiesPCAPsApi.md#StartOrgPacketCapture) | **Post** /api/v1/orgs/{org_id}/pcaps/capture | startOrgPacketCapture
[**StartSitePacketCapture**](UtilitiesPCAPsApi.md#StartSitePacketCapture) | **Post** /api/v1/sites/{site_id}/pcaps/capture | startSitePacketCapture
[**StopOrgPacketCapture**](UtilitiesPCAPsApi.md#StopOrgPacketCapture) | **Delete** /api/v1/orgs/{org_id}/pcaps/capture | stopOrgPacketCapture
[**StopSitePacketCapture**](UtilitiesPCAPsApi.md#StopSitePacketCapture) | **Delete** /api/v1/sites/{site_id}/pcaps/capture | stopSitePacketCapture
[**UpdateSitePacketCapture**](UtilitiesPCAPsApi.md#UpdateSitePacketCapture) | **Put** /api/v1/sites/{site_id}/pcaps/{pcap_id} | updateSitePacketCapture

# **GetOrgCapturingStatus**
> InlineResponse200192 GetOrgCapturingStatus(ctx, orgId)
getOrgCapturingStatus

Get Org Capturing status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200192**](inline_response_200_192.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteCapturingStatus**
> InlineResponse200192 GetSiteCapturingStatus(ctx, siteId)
getSiteCapturingStatus

Get Capturing status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200192**](inline_response_200_192.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgPacketCaptures**
> InlineResponse200191 ListOrgPacketCaptures(ctx, orgId, optional)
listOrgPacketCaptures

Get List of Org  Packet Captures

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesPCAPsApiListOrgPacketCapturesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesPCAPsApiListOrgPacketCapturesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**InlineResponse200191**](inline_response_200_191.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSitePacketCaptures**
> InlineResponse200191 ListSitePacketCaptures(ctx, siteId, optional)
listSitePacketCaptures

Get List of Site Packet Captures

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesPCAPsApiListSitePacketCapturesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesPCAPsApiListSitePacketCapturesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientMac** | **optional.String**| optional client mac filter | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**InlineResponse200191**](inline_response_200_191.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartOrgPacketCapture**
> InlineResponse200193 StartOrgPacketCapture(ctx, orgId, optional)
startOrgPacketCapture

Initiate a Packet Capture  The output will be available through websocket. As there can be multiple command issued against the same AP at the same time and the output all goes through the same websocket stream, session is introduced for demux.  #### Subscribe to Device Command outputs `WS /api-ws/v1/stream`  ```json {     \"subscribe\": \"/sites/{site_id}/pcaps\" } ``` #### Response (Wireless/RadioTap) ```json {   \"event\": \"data\"   \"channel\": \"/orgs/67970e46-4e12-11e6-9188-0242ac110007/pcaps\"   \"data\": {       \"capture_id\": \"f039b1b4-a23e-48b2-906a-0da40524de73\",        \"pcap_dict\": {           \"dst_mac\": \"68:ec:c5:09:2e:87\",           \"src_mac\": \"8c:3b:ad:e0:47:40\",            \"vlan\": 1,            \"src_ip\": \"34.224.147.117\",            \"dst_ip\": \"192.168.1.55\",           \"dst_port\": 51635,            \"src_port\": 443,           \"protocol\": \"TCP\",            \"mxedge_id\": \"00000000-0000-0000-1000-001122334455\",           \"direction\": \"tx\",            \"timestamp\": 1652247615,            \"length\": 159.0,            \"interface\": \"port0\",           \"info\": \"1652247616.007409 IP ec2-34-224-147-117.compute-1.amazonaws.com.https > ip-192-168-1-55.ec2.internal.51635: Flags [P.], seq                      2192123968:2192124057, ack 4035166782, win 12, options [nop,nop,TS val 597467050 ecr 740580660], length 89\\\\n\",           },        \"pcap_raw\": \"1MOyoQIABAAAAAAAAAAAAP//AAABAAAAQEx7YhMzAACfAAAAnwAAAGjsxQkuh4w7reBHQIEAAAEIAEUAAI1bLEAAKAZ/CiLgk3XAqAE3AbvJs4KpKEDwg8I+gBgADFf9AAABAQgKI5yfqiwkXTQXAwMAVKY5JopoKQrVEn0/3ld4YntctGEH/rTZuwtCvzSncFw71QJveJi9uxHs57KC8w9Apph3YvXJrmWg7M37+o+YV0KH/xmr626s5Bkhb3QhKOu+NoNEmA==\\\"     } } ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesPCAPsApiStartOrgPacketCaptureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesPCAPsApiStartOrgPacketCaptureOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PcapsCaptureBody**](PcapsCaptureBody.md)| Request Body | 

### Return type

[**InlineResponse200193**](inline_response_200_193.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartSitePacketCapture**
> InlineResponse200193 StartSitePacketCapture(ctx, siteId, optional)
startSitePacketCapture

Initiate a Site Packet Capture  The output will be available through websocket. As there can be multiple command issued against the same AP at the same time and the output all goes through the same websocket stream, session is introduced for demux.  #### Subscribe to Device Command outputs `WS /api-ws/v1/stream`  ```json {     \"subscribe\": \"/sites/{site_id}/pcaps\" } ``` #### Response (MxEdge) ```json {     \"event\": \"data\"     \"channel\": \"/sites/:site_id/pcaps\"     \"data\": {          \"capture_id\": \"6b1be4fb-b239-44d9-9d3b-cb1ff3af1721\",      \"lost_messages\": 0          \"pcap_dict\": {              \"channel_frequency\": 2412,              \"channel\": \"1\",              \"datarate\": \"1.0 Mbps\",              \"rssi\": -75,               \"dst\": \"78:bd:bc:ca:0b:0a\",              \"src\": \"18:b8:1f:4c:91:c0\",              \"bssid\": \"18:b8:1f:4c:91:c0\",              \"frame_type\": \"Management\",               \"frame_subtype\": \"Probe Response\",           \"proto\": \"802.11\",               \"ap_mac\": \"d4:20:b0:81:99:2e\",               \"direction\": \"tx\",               \"timestamp\": 1652246543,               \"length\": 416.0,              \"interface\": \"radiotap\",              \"info\": \"1652246544.467733 1683216786us tsft 1.0 Mb/s 2412 MHz 11g -75dBm signal -82dBm noise antenna 0 Probe Response (ATTKmsWiVS) [1.0* 2.0* 5.5* 11.0* 18.0 24.0 36.0 54.0 Mbit] CH: 2, PRIVACY\\\\n\",          },          \"pcap_raw\": \"1MOyoQIABAAAAAAAAAAAAP//AAABAAAAEEh7Yh5VBwCgAQAAoAEAAAAAKwBvCADAAQAAAIw7reCS2VNkAAAAABACbAmABLWuAAEAEBgAAwACAABQADoBeL28ygsKGLgfTJHAGLgfTJHAcIZ2WDlBJQAAAGQAERUACkFUVEttc1dpVlMBCIKEi5YkMEhsAwECBwZVUyABCx4gAQAjAhkAKgEEMgQMEhhgMBQBAAAPrAQBAAAPrAQBAAAPrAIMAAsFAQAbAABGBTIIAQAALRqtCR////8AAAAAAAAAAAAAAAAAAAAAAAAAAD0WAggVAAAAAAAAAAAAAAAAAAAAAAAAAH8IBAAIAAAAAEDdkwBQ8gQQSgABEBBEAAECEDsAAQMQRwAQn2481frn3KT+uGod2ERx+RAhAAtBcnJpcywgSW5jLhAjAApCR1cyMTAtNzAwECQACkJHVzIxMC03MDAQQgAKQkdXMjEwLTcwMBBUAAgABgBQ8gQAARARAA5BcnJpcyBXaXJlbGVzcxAIAAIgCBA8AAEBEEkABgA3KgABIN0JABAYAgEQHAAA3RgAUPICAQGEAAOkAAAnpAAAQkNeAGIyLwAzjakr\" } ```  #### Response (Wired) ```json {     \"event\": \"data\"     \"channel\": \"/sites/67970e46-4e12-11e6-9188-0242ac110007/pcaps\"     \"data\": {         \"capture_id\": \"f039b1b4-a23e-48b2-906a-0da40524de73\",          \"pcap_dict\": {              \"dst_mac\": \"68:ec:c5:09:2e:87\",              \"src_mac\": \"8c:3b:ad:e0:47:40\",               \"vlan\": 1,               \"src_ip\": \"34.224.147.117\",               \"dst_ip\": \"192.168.1.55\",              \"dst_port\": 51635,               \"src_port\": 443,              \"proto\": \"TCP\",               \"ap_mac\": \"d4:20:b0:81:99:2e\",              \"direction\": \"tx\",               \"timestamp\": 1652247615,               \"length\": 159.0,               \"interface\": \"wired\",              \"info\": \"1652247616.007409 IP ec2-34-224-147-117.compute-1.amazonaws.com.https > ip-192-168-1-55.ec2.internal.51635: Flags [P.], seq 2192123968:2192124057, ack 4035166782, win 12, options [nop,nop,TS val 597467050 ecr 740580660], length 89\\\\n\",              },          \"pcap_raw\": \"1MOyoQIABAAAAAAAAAAAAP//AAABAAAAQEx7YhMzAACfAAAAnwAAAGjsxQkuh4w7reBHQIEAAAEIAEUAAI1bLEAAKAZ/CiLgk3XAqAE3AbvJs4KpKEDwg8I+gBgADFf9AAABAQgKI5yfqiwkXTQXAwMAVKY5JopoKQrVEn0/3ld4YntctGEH/rTZuwtCvzSncFw71QJveJi9uxHs57KC8w9Apph3YvXJrmWg7M37+o+YV0KH/xmr626s5Bkhb3QhKOu+NoNEmA==\"      } } ```  #### Stop Response (Wired/Wireless) ```json {     \"event\": \"data\"     \"channel\": \"/sites/67970e46-4e12-11e6-9188-0242ac110007/pcaps\"     \"data\": {       \"capture_id\": \"a2f7374d-6a70-41fd-8a3f-71e42573baaf\",        \"lost_messages\": 0,       \"pcap_dict\": null     } } ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesPCAPsApiStartSitePacketCaptureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesPCAPsApiStartSitePacketCaptureOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PcapsCaptureBody1**](PcapsCaptureBody1.md)| Request Body | 

### Return type

[**InlineResponse200193**](inline_response_200_193.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopOrgPacketCapture**
> StopOrgPacketCapture(ctx, orgId)
stopOrgPacketCapture

Stop current Org capture

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopSitePacketCapture**
> StopSitePacketCapture(ctx, siteId)
stopSitePacketCapture

Stop current capture

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSitePacketCapture**
> UpdateSitePacketCapture(ctx, siteId, pcapId, optional)
updateSitePacketCapture

Update or add notes to a completed packet capture

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **pcapId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesPCAPsApiUpdateSitePacketCaptureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesPCAPsApiUpdateSitePacketCaptureOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of PcapsPcapIdBody**](PcapsPcapIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

