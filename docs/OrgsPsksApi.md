# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgPsk**](OrgsPsksApi.md#CreateOrgPsk) | **Post** /api/v1/orgs/{org_id}/psks | createOrgPsk
[**DeleteOrgPsk**](OrgsPsksApi.md#DeleteOrgPsk) | **Delete** /api/v1/orgs/{org_id}/psks/{psk_id} | deleteOrgPsk
[**DeleteOrgPskList**](OrgsPsksApi.md#DeleteOrgPskList) | **Post** /api/v1/orgs/{org_id}/psks/delete | deleteOrgPskList
[**DeleteOrgPskOldPassphrase**](OrgsPsksApi.md#DeleteOrgPskOldPassphrase) | **Post** /api/v1/orgs/{org_id}/psks/{psk_id}/delete_old_passphrase | deleteOrgPskOldPassphrase
[**GetOrgPsk**](OrgsPsksApi.md#GetOrgPsk) | **Get** /api/v1/orgs/{org_id}/psks/{psk_id} | getOrgPsk
[**ImportOrgPsks**](OrgsPsksApi.md#ImportOrgPsks) | **Post** /api/v1/orgs/{org_id}/psks/import | importOrgPsks
[**ListOrgPsks**](OrgsPsksApi.md#ListOrgPsks) | **Get** /api/v1/orgs/{org_id}/psks | listOrgPsks
[**UpdateOrgMultiplePsks**](OrgsPsksApi.md#UpdateOrgMultiplePsks) | **Put** /api/v1/orgs/{org_id}/psks | updateOrgMultiplePsks
[**UpdateOrgPsk**](OrgsPsksApi.md#UpdateOrgPsk) | **Put** /api/v1/orgs/{org_id}/psks/{psk_id} | updateOrgPsk

# **CreateOrgPsk**
> InlineResponse20085 CreateOrgPsk(ctx, orgId, optional)
createOrgPsk

Create Org PSK   When `usage`==`macs`, corresponding \"macs\" field will hold a list consisting of client mac addresses ([\"xx:xx:xx:xx:xx\",...]) or mac patterns([\"xx:xx:*\",\"xx*\",...]) or both ([\"xx:xx:xx:xx:xx:xx\", \"xx:*\", ...]). This list is capped at 5000

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsPsksApiCreateOrgPskOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsPsksApiCreateOrgPskOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdPsksBody**](OrgIdPsksBody.md)| Request Body | 
 **upsert** | **optional.**| if a key exists with the same &#x60;name&#x60;, replace it with the new one | 

### Return type

[**InlineResponse20085**](inline_response_200_85.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgPsk**
> DeleteOrgPsk(ctx, orgId, pskId)
deleteOrgPsk

Delete Org PSK

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **pskId** | [**string**](.md)| PSK ID | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgPskList**
> DeleteOrgPskList(ctx, orgId, optional)
deleteOrgPskList

Delete Org PSK List  Delete list of psks on the org. This API accepts single string or list of strings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsPsksApiDeleteOrgPskListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsPsksApiDeleteOrgPskListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PskIdList**](PskIdList.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgPskOldPassphrase**
> InlineResponse20085 DeleteOrgPskOldPassphrase(ctx, orgId, pskId)
deleteOrgPskOldPassphrase

Delete `old_passphrase` from PSK.  If successful, response is same as GET, returns the PSK with `old_passphrase` removed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **pskId** | [**string**](.md)| PSK ID | 

### Return type

[**InlineResponse20085**](inline_response_200_85.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgPsk**
> InlineResponse20085 GetOrgPsk(ctx, orgId, pskId)
getOrgPsk

Get Org PSK Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **pskId** | [**string**](.md)| PSK ID | 

### Return type

[**InlineResponse20085**](inline_response_200_85.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportOrgPsks**
> []Psk ImportOrgPsks(ctx, orgId, optional)
importOrgPsks

Import PSK from CSV file or JSON  ## CSV File Format ``` PSK Import CSV File Format: name,ssid,passphrase,usage,vlan_id,mac,max_usage,role,expire_time,notify_expiry,expiry_notification_time,notify_on_create_or_edit,email Common,warehouse,foryoureyesonly,single,35,a31425f31278,0,student,1618594236 Justin,reception,visible,multi,1002,200,teacher,1618594236 Common2,ssid,1245678-xx,single,35,a31425f31278,0,student,1618594236,true,7,true,admin@test.com ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsPsksApiImportOrgPsksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsPsksApiImportOrgPsksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | **optional.Interface of *os.File****optional.**|  | 

### Return type

[**[]Psk**](psk.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgPsks**
> []Psk ListOrgPsks(ctx, orgId, optional)
listOrgPsks

Get List of Org Psks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsPsksApiListOrgPsksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsPsksApiListOrgPsksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**|  | 
 **ssid** | **optional.String**|  | 
 **role** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]Psk**](psk.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgMultiplePsks**
> []Psk UpdateOrgMultiplePsks(ctx, orgId, optional)
updateOrgMultiplePsks

Update Multiple PSKs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsPsksApiUpdateOrgMultiplePsksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsPsksApiUpdateOrgMultiplePsksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of []Psk**](psk.md)|  | 

### Return type

[**[]Psk**](psk.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgPsk**
> InlineResponse20085 UpdateOrgPsk(ctx, orgId, pskId, optional)
updateOrgPsk

Update Org PSK

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **pskId** | [**string**](.md)| PSK ID | 
 **optional** | ***OrgsPsksApiUpdateOrgPskOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsPsksApiUpdateOrgPskOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of PsksPskIdBody**](PsksPskIdBody.md)| Request Body | 

### Return type

[**InlineResponse20085**](inline_response_200_85.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

