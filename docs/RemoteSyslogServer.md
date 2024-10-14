# RemoteSyslogServer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | [**[]RemoteSyslogContent**](remote_syslog_content.md) |  | [optional] [default to null]
**ExplicitPriority** | **bool** |  | [optional] [default to null]
**Facility** | [***AllOfremoteSyslogServerFacility**](AllOfremoteSyslogServerFacility.md) |  | [optional] [default to null]
**Host** | **string** |  | [optional] [default to null]
**Match** | **string** |  | [optional] [default to null]
**Port** | **int32** |  | [optional] [default to 514]
**Protocol** | [***AllOfremoteSyslogServerProtocol**](AllOfremoteSyslogServerProtocol.md) |  | [optional] [default to null]
**RoutingInstance** | **string** |  | [optional] [default to null]
**Severity** | [***AllOfremoteSyslogServerSeverity**](AllOfremoteSyslogServerSeverity.md) |  | [optional] [default to null]
**SourceAddress** | **string** | if source_address is configured, will use the vlan firstly otherwise use source_ip | [optional] [default to null]
**StructuredData** | **bool** |  | [optional] [default to null]
**Tag** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

