# MxclusterRadsec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcctServers** | [**[]MxclusterRadsecAcctServer**](mxcluster_radsec_acct_server.md) | list of RADIUS accounting servers, optional, order matters where the first one is treated as primary | [optional] [default to null]
**AuthServers** | [**[]MxclusterRadsecAuthServer**](mxcluster_radsec_auth_server.md) | list of RADIUS authentication servers, order matters where the first one is treated as primary | [optional] [default to null]
**Enabled** | **bool** | whether to enable service on Mist Edge i.e. RADIUS proxy over TLS | [optional] [default to null]
**MatchSsid** | **bool** | whether to match ssid in request message to select from a subset of RADIUS servers | [optional] [default to null]
**NasIpSource** | [***AllOfmxclusterRadsecNasIpSource**](AllOfmxclusterRadsecNasIpSource.md) |  | [optional] [default to null]
**ProxyHosts** | **[]string** | hostnames or IPs for Mist AP to use as the TLS Server (i.e. they are reachable from AP) in addition to &#x60;tunterm_hosts&#x60; | [optional] [default to null]
**ServerSelection** | [***AllOfmxclusterRadsecServerSelection**](AllOfmxclusterRadsecServerSelection.md) |  | [optional] [default to null]
**SrcIpSource** | [***AllOfmxclusterRadsecSrcIpSource**](AllOfmxclusterRadsecSrcIpSource.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

