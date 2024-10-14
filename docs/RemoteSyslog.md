# RemoteSyslog

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archive** | [***RemoteSyslogArchive**](remote_syslog_archive.md) |  | [optional] [default to null]
**Console** | [***RemoteSyslogConsole**](remote_syslog_console.md) |  | [optional] [default to null]
**Enabled** | **bool** |  | [optional] [default to false]
**Files** | [**[]RemoteSyslogFileConfig**](remote_syslog_file_config.md) |  | [optional] [default to null]
**Network** | **string** | if source_address is configured, will use the vlan firstly otherwise use source_ip | [optional] [default to null]
**SendToAllServers** | **bool** |  | [optional] [default to false]
**Servers** | [**[]RemoteSyslogServer**](remote_syslog_server.md) |  | [optional] [default to null]
**TimeFormat** | [***AllOfremoteSyslogTimeFormat**](AllOfremoteSyslogTimeFormat.md) |  | [optional] [default to null]
**Users** | [**[]RemoteSyslogUser**](remote_syslog_user.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

