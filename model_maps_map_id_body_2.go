/*
 * Mist API
 *
 * > Version: **2409.1.9** > > Date: **September 27, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates] (https://www.juniper.net/documentation/us/en/software/mist/api/http/getting-started/how-to-get-started)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 
 *
 * API version: 2409.1.9
 * Contact: tmunzer@juniper.net
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package mist

type MapsMapIdBody2 struct {
	CreatedTime float64 `json:"created_time,omitempty"`
	// name/val pair objects for location engine to use
	Flags map[string]int32 `json:"flags,omitempty"`
	ForSite bool `json:"for_site,omitempty"`
	// when type=image, height of the image map
	Height int32 `json:"height,omitempty"`
	HeightM float64 `json:"height_m,omitempty"`
	Id string `json:"id,omitempty"`
	LatlngBr *any `json:"latlng_br,omitempty"`
	LatlngTl *any `json:"latlng_tl,omitempty"`
	// whether this map is considered locked down
	Locked bool `json:"locked,omitempty"`
	ModifiedTime float64 `json:"modified_time,omitempty"`
	// The name of the map
	Name string `json:"name,omitempty"`
	OccupancyLimit int32 `json:"occupancy_limit,omitempty"`
	OrgId string `json:"org_id,omitempty"`
	// orientation of the map, 0 means up is north, 90 means up is west
	Orientation int32 `json:"orientation,omitempty"`
	// the user-annotated x origin, pixels
	OriginX int32 `json:"origin_x,omitempty"`
	// the user-annotated y origin, pixels
	OriginY int32 `json:"origin_y,omitempty"`
	// when type=image, pixels per meter
	Ppm float64 `json:"ppm,omitempty"`
	SiteId string `json:"site_id,omitempty"`
	// sitesurvey_path
	SitesurveyPath []MapSitesurveyPathItems `json:"sitesurvey_path,omitempty"`
	// when type=image, the url for the thumbnail image / preview
	ThumbnailUrl string `json:"thumbnail_url,omitempty"`
	Type_ *any `json:"type,omitempty"`
	// when type=image, the url
	Url string `json:"url,omitempty"`
	// whether this map uses autooreintation values or ignores them
	UseAutoOrientation bool `json:"use_auto_orientation,omitempty"`
	// whether this map uses autoplacement values or ignores them
	UseAutoPlacement bool `json:"use_auto_placement,omitempty"`
	View *any `json:"view,omitempty"`
	WallPath *any `json:"wall_path,omitempty"`
	Wayfinding *any `json:"wayfinding,omitempty"`
	WayfindingPath *any `json:"wayfinding_path,omitempty"`
	// when type=image, width of the image map
	Width int32 `json:"width,omitempty"`
	WidthM float64 `json:"width_m,omitempty"`
}
