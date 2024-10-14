# AllOfsiteSettingAutoUpgrade

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomVersions** | **map[string]string** | custom versions for different models. Property key is the model name (e.g. \&quot;AP41\&quot;) | [optional] [default to null]
**DayOfWeek** | [***Object**](.md) |  | [optional] [default to null]
**Enabled** | **bool** | whether auto upgrade should happen (Note that Mist may auto-upgrade if the version is not supported) | [optional] [default to false]
**TimeOfDay** | **string** | any / HH:MM (24-hour format), upgrade will happen within up to 1-hour from this time | [optional] [default to null]
**Version** | [***Object**](.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

