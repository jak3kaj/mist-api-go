# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ValidateOrgIdpCredential**](OrgsNACIDPApi.md#ValidateOrgIdpCredential) | **Post** /api/v1/orgs/{org_id}/mist_nac/test_idp | validateOrgIdpCredential

# **ValidateOrgIdpCredential**
> WebsocketSession ValidateOrgIdpCredential(ctx, orgId, optional)
validateOrgIdpCredential

IDP Credential Validation. The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.  #### Subscribe to Device Command outputs `WS /api-ws/v1/stream`  ``` json {     \"subscribe\": \"orgs/:org_id/mist_nac/test_idp\" }   ```  ### Response (no idp can be found)  ``` json {     \"event\": \"data\",     \"channel\": \"/orgs/:org_id/mist_nac/test_idp\",     \"status\":      \"data\": {         \"status\": \"failure\",         \"error\": \"No matching IDP found\"     } }   ```  ### Response OK  ``` json {     \"event\": \"data\",     \"channel\": \"/orgs/:org_id/mist_nac/test_idp\",     \"status\":      \"data\": {         \"status\": \"success\",         \"idp_id\": \"915793c0-1355-4e98-b1c0-23df2227b357\",         \"idp_type\": \"ldap\",         // more attributes will be added later     } }   ```  ### Response Invalid Credentials  ``` json {     \"event\": \"data\",     \"channel\": \"/orgs/:org_id/mist_nac/test_idp\",     \"status\":      \"data\": {         \"status\": \"failure\",         \"error\": \"Invalid Credentials\",         \"idp_id\": \"915793c0-1355-4e98-b1c0-23df2227b357\",         \"idp_type\": \"ldap\",     } }   ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsNACIDPApiValidateOrgIdpCredentialOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsNACIDPApiValidateOrgIdpCredentialOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MistNacTestIdpBody**](MistNacTestIdpBody.md)|  | 

### Return type

[**WebsocketSession**](websocket_session.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

