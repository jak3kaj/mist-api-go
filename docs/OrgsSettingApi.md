# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgWirelessClientsBlocklist**](OrgsSettingApi.md#CreateOrgWirelessClientsBlocklist) | **Post** /api/v1/orgs/{org_id}/setting/blacklist | createOrgWirelessClientsBlocklist
[**DeleteOrgWirelessClientsBlocklist**](OrgsSettingApi.md#DeleteOrgWirelessClientsBlocklist) | **Delete** /api/v1/orgs/{org_id}/setting/blacklist | deleteOrgWirelessClientsBlocklist
[**GetOrgSettings**](OrgsSettingApi.md#GetOrgSettings) | **Get** /api/v1/orgs/{org_id}/setting | getOrgSettings
[**SetOrgCustomBucket**](OrgsSettingApi.md#SetOrgCustomBucket) | **Post** /api/v1/orgs/{org_id}/setting/pcap_bucket/setup | setOrgCustomBucket
[**UpdateOrgSettings**](OrgsSettingApi.md#UpdateOrgSettings) | **Put** /api/v1/orgs/{org_id}/setting | updateOrgSettings
[**VerifyOrgCustomBucket**](OrgsSettingApi.md#VerifyOrgCustomBucket) | **Post** /api/v1/orgs/{org_id}/setting/pcap_bucket/verify | verifyOrgCustomBucket

# **CreateOrgWirelessClientsBlocklist**
> InlineResponse20095 CreateOrgWirelessClientsBlocklist(ctx, orgId, optional)
createOrgWirelessClientsBlocklist

Create Org Blacklist Client List.   If there is already a blacklist, this API will replace it with the new one.   Max number of blacklist clients is 1000.   Retrieve the current blacklisted clients from `blacklist_url` under Org:Setting 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsSettingApiCreateOrgWirelessClientsBlocklistOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSettingApiCreateOrgWirelessClientsBlocklistOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SettingBlacklistBody**](SettingBlacklistBody.md)| Request Body | 

### Return type

[**InlineResponse20095**](inline_response_200_95.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgWirelessClientsBlocklist**
> DeleteOrgWirelessClientsBlocklist(ctx, orgId)
deleteOrgWirelessClientsBlocklist

Delete Org Blacklist Station Clients

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

# **GetOrgSettings**
> InlineResponse20093 GetOrgSettings(ctx, orgId)
getOrgSettings

Get Org Settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20093**](inline_response_200_93.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetOrgCustomBucket**
> InlineResponse20096 SetOrgCustomBucket(ctx, orgId, optional)
setOrgCustomBucket

Provide Customer Bucket Name  Setting up Custom PCAP Bucket Involves the following: * provide the bucket name * we’ll attempt to write a file MIST_TOKEN * you have to verify the ownership of the bucket by providing the content of the MIST_TOKEN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsSettingApiSetOrgCustomBucketOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSettingApiSetOrgCustomBucketOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PcapBucketSetupBody**](PcapBucketSetupBody.md)| Request Body | 

### Return type

[**InlineResponse20096**](inline_response_200_96.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgSettings**
> InlineResponse20094 UpdateOrgSettings(ctx, orgId, optional)
updateOrgSettings

Update Org Settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsSettingApiUpdateOrgSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSettingApiUpdateOrgSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdSettingBody**](OrgIdSettingBody.md)| Request Body | 

### Return type

[**InlineResponse20094**](inline_response_200_94.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyOrgCustomBucket**
> VerifyOrgCustomBucket(ctx, orgId, optional)
verifyOrgCustomBucket

Verify Customer PCAP Bucket  **Note**: If successful, a “VERIFIED” file will be created in the bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsSettingApiVerifyOrgCustomBucketOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSettingApiVerifyOrgCustomBucketOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PcapBucketVerifyBody**](PcapBucketVerifyBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

