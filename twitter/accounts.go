package twitter

import (
	"net/http"

	"github.com/dghubble/sling"
)

// AccountService provides a method for account credential verification.
type AccountService struct {
	sling *sling.Sling
}

// newAccountService returns a new AccountService.
func newAccountService(sling *sling.Sling) *AccountService {
	return &AccountService{
		sling: sling.Path("account/"),
	}
}

// AccountVerifyParams are the params for AccountService.VerifyCredentials.
type AccountVerifyParams struct {
	IncludeEntities *bool `url:"include_entities,omitempty"`
	SkipStatus      *bool `url:"skip_status,omitempty"`
	IncludeEmail    *bool `url:"include_email,omitempty"`
}

// VerifyCredentials returns the authorized user if credentials are valid and
// returns an error otherwise.
// Requires a user auth context.
// https://dev.twitter.com/rest/reference/get/account/verify_credentials
func (s *AccountService) VerifyCredentials(params *AccountVerifyParams) (*User, *http.Response, error) {
	user := new(User)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("verify_credentials.json").QueryStruct(params).Receive(user, apiError)
	return user, resp, relevantError(err, *apiError)
}

// AccountSettingsResult is the result from AccountService.Settings
type AccountSettingsResult struct {
	AlwaysUseHTTPS           bool                                 `json:"always_use_https" url:"always_use_https,omitempty"`
	DiscoverableByEmail      bool                                 `json:"discoverable_by_email" url:"discoverable_by_email,omitempty"`
	GeoEnabled               bool                                 `json:"geo_enabled" url:"geo_enabled,omitempty"`
	Language                 string                               `json:"language" url:"language,omitempty"`
	Protected                bool                                 `json:"protected" url:"protected,omitempty"`
	ScreenName               string                               `json:"screen_name" url:"screen_name,omitempty"`
	ShowAllInlineMedia       bool                                 `json:"show_all_inline_media" url:"show_all_inline_media,omitempty"`
	UseCookiePersonalization bool                                 `json:"use_cookie_personalization" url:"use_cookie_personalization,omitempty"`
	AllowContributorRequest  string                               `json:"allow_contributor_request" url:"allow_contributor_request,omitempty"`
	SleepTime                AccountSettingsResultSleepTime       `json:"sleep_time" url:"sleep_time,omitempty"`
	TimeZone                 AccountSettingsResultTimeZone        `json:"time_zone" url:"time_zone,omitempty"`
	TrendLocation            []AccountSettingsResultTrendLocation `json:"trend_location" url:"trend_location,omitempty"`
}

// AccountSettingsResultSleepTime is part of the result from
// AccountService.Settings
type AccountSettingsResultSleepTime struct {
	Enabled   bool   `json:"enabled" url:"enabled"`
	EndTime   string `json:"end_time" url:"end_time"`
	StartTime string `json:"start_time" url:"start_time"`
}

// AccountSettingsResultTimeZone is part of the result from
// AccountService.Settings
type AccountSettingsResultTimeZone struct {
	Name       string `json:"name" url:"name"`
	TzinfoName string `json:"tzinfo_name" url:"tzinfo_name"`
	UtcOffset  int64  `json:"utc_offset" url:"utc_offset"`
}

// AccountSettingsResultTrendLocation is part of the result from
// AccountService.Settings
type AccountSettingsResultTrendLocation struct {
	Country     string                         `json:"country" url:"country"`
	CountryCode string                         `json:"countryCode" url:"countryCode"`
	Name        string                         `json:"name" url:"name"`
	ParentID    int64                          `json:"parentid" url:"parentid"`
	URL         string                         `json:"url" url:"url"`
	WoeID       int64                          `json:"woeid" url:"woeid"`
	PlaceType   AccountSettingsResultPlaceType `json:"placeType" url:"placeType"`
}

// AccountSettingsResultPlaceType is part of the result from
// AccountService.Settings
type AccountSettingsResultPlaceType struct {
	Code int64  `json:"code" url:"code"`
	Name string `json:"name" url:"name"`
}

// Settings returns settings (including current trend, geo and sleep time
// information) for the authenticating user.
// https://dev.twitter.com/rest/reference/get/account/settings
func (s *AccountService) Settings() (*AccountSettingsResult, *http.Response, error) {
	settings := new(AccountSettingsResult)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("settings.json").Receive(settings, apiError)
	return settings, resp, relevantError(err, *apiError)
}

// UpdateSettings updates the authenticating user’s settings.
// https://dev.twitter.com/rest/reference/post/account/settings
func (s *AccountService) UpdateSettings(params *AccountSettingsResult) (*AccountSettingsResult, *http.Response, error) {
	settings := new(AccountSettingsResult)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("settings.json").QueryStruct(params).Receive(settings, apiError)
	return settings, resp, relevantError(err, *apiError)
}

// AccountSettingsUpdateProfileParams are the parameters
// for AccountService.UpdateProfile
type AccountSettingsUpdateProfileParams struct {
	Name             string `url:"name"`
	URL              string `url:"url"`
	Location         string `url:"location"`
	Description      string `url:"description"`
	ProfileLinkColor string `url:"profile_link_color"`
	IncludeEntities  bool   `url:"include_entities"`
	SkipStatus       bool   `url:"skip_status"`
}

// AccountSettingsUpdateProfileResult is the result
// for AccountService.UpdateProfile
type AccountSettingsUpdateProfileResult struct {
	CreatedAt                      string `json:"created_at"`
	Description                    string `json:"description"`
	FavouritesCount                int64  `json:"favourites_count"`
	FollowersCount                 int64  `json:"followers_count"`
	FriendsCount                   int64  `json:"friends_count"`
	GeoEnabled                     bool   `json:"geo_enabled"`
	ID                             int64  `json:"id"`
	IDStr                          string `json:"id_str"`
	Lang                           string `json:"lang"`
	ListedCount                    int64  `json:"listed_count"`
	Location                       string `json:"location"`
	Name                           string `json:"name"`
	ProfileBackgroundColor         string `json:"profile_background_color"`
	ProfileBackgroundImageURL      string `json:"profile_background_image_url"`
	ProfileBackgroundImageURLHTTPS string `json:"profile_background_image_url_https"`
	ProfileImageURL                string `json:"profile_image_url"`
	ProfileImageURLHTTPS           string `json:"profile_image_url_https"`
	ProfileLinkColor               string `json:"profile_link_color"`
	ProfileSidebarBorderColor      string `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor        string `json:"profile_sidebar_fill_color"`
	ProfileTextColor               string `json:"profile_text_color"`
	ScreenName                     string `json:"screen_name"`
	StatusesCount                  int64  `json:"statuses_count"`
	TimeZone                       string `json:"time_zone"`
	URL                            string `json:"url"`
	UtcOffset                      int64  `json:"utc_offset"`
	ProfileUseBackgroundImage      bool   `json:"profile_use_background_image"`
	ContributorsEnabled            bool   `json:"contributors_enabled"`
	DefaultProfile                 bool   `json:"default_profile"`
	DefaultProfileImage            bool   `json:"default_profile_image"`
	FollowRequestSent              bool   `json:"follow_request_sent"`
	Following                      bool   `json:"following"`
	IsTranslator                   bool   `json:"is_translator"`
	Notifications                  bool   `json:"notifications"`
	ProfileBackgroundTile          bool   `json:"profile_background_tile"`
	Protected                      bool   `json:"protected"`
	ShowAllInlineMedia             bool   `json:"show_all_inline_media"`
	Verified                       bool   `json:"verified"`
}

// UpdateProfile sets some values that users are able to set under the
// “Account” tab of their settings page. Only the parameters specified will
// be updated.
// https://dev.twitter.com/rest/reference/post/account/update_profile
func (s *AccountService) UpdateProfile(params *AccountSettingsUpdateProfileParams) (*AccountSettingsUpdateProfileResult, *http.Response, error) {
	settings := new(AccountSettingsUpdateProfileResult)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("update_profile.json").QueryStruct(params).Receive(settings, apiError)
	return settings, resp, relevantError(err, *apiError)
}
