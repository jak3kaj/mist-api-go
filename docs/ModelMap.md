# ModelMap

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | **float64** |  | [optional] [default to null]
**Flags** | **map[string]int32** | name/val pair objects for location engine to use | [optional] [default to null]
**ForSite** | **bool** |  | [optional] [default to null]
**Height** | **int32** | when type&#x3D;image, height of the image map | [optional] [default to null]
**HeightM** | **float64** |  | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**LatlngBr** | [***AllOfMapLatlngBr**](AllOfMapLatlngBr.md) |  | [optional] [default to null]
**LatlngTl** | [***AllOfMapLatlngTl**](AllOfMapLatlngTl.md) |  | [optional] [default to null]
**Locked** | **bool** | whether this map is considered locked down | [optional] [default to false]
**ModifiedTime** | **float64** |  | [optional] [default to null]
**Name** | **string** | The name of the map | [optional] [default to null]
**OccupancyLimit** | **int32** |  | [optional] [default to null]
**OrgId** | **string** |  | [optional] [default to null]
**Orientation** | **int32** | orientation of the map, 0 means up is north, 90 means up is west | [optional] [default to 0]
**OriginX** | **int32** | the user-annotated x origin, pixels | [optional] [default to null]
**OriginY** | **int32** | the user-annotated y origin, pixels | [optional] [default to null]
**Ppm** | **float64** | when type&#x3D;image, pixels per meter | [optional] [default to null]
**SiteId** | **string** |  | [optional] [default to null]
**SitesurveyPath** | [**[]MapSitesurveyPathItems**](map_sitesurvey_path_items.md) | sitesurvey_path | [optional] [default to null]
**ThumbnailUrl** | **string** | when type&#x3D;image, the url for the thumbnail image / preview | [optional] [default to null]
**Type_** | [***AllOfMapType_**](AllOfMapType_.md) |  | [optional] [default to null]
**Url** | **string** | when type&#x3D;image, the url | [optional] [default to null]
**UseAutoOrientation** | **bool** | whether this map uses autooreintation values or ignores them | [optional] [default to false]
**UseAutoPlacement** | **bool** | whether this map uses autoplacement values or ignores them | [optional] [default to false]
**View** | [***AllOfMapView**](AllOfMapView.md) |  | [optional] [default to null]
**WallPath** | [***AllOfMapWallPath**](AllOfMapWallPath.md) |  | [optional] [default to null]
**Wayfinding** | [***AllOfMapWayfinding**](AllOfMapWayfinding.md) |  | [optional] [default to null]
**WayfindingPath** | [***AllOfMapWayfindingPath**](AllOfMapWayfindingPath.md) |  | [optional] [default to null]
**Width** | **int32** | when type&#x3D;image, width of the image map | [optional] [default to null]
**WidthM** | **float64** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

