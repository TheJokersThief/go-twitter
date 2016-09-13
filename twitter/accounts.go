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
	AlwaysUseHTTPS           bool                                 `json:"always_use_https"`
	DiscoverableByEmail      bool                                 `json:"discoverable_by_email"`
	GeoEnabled               bool                                 `json:"geo_enabled"`
	Language                 string                               `json:"language"`
	Protected                bool                                 `json:"protected"`
	ScreenName               string                               `json:"screen_name"`
	ShowAllInlineMedia       bool                                 `json:"show_all_inline_media"`
	UseCookiePersonalization bool                                 `json:"use_cookie_personalization"`
	AllowContributorRequest  string                               `json:"allow_contributor_request"`
	SleepTime                AccountSettingsResultSleepTime       `json:"sleep_time"`
	TimeZone                 AccountSettingsResultTimeZone        `json:"time_zone"`
	TrendLocation            []AccountSettingsResultTrendLocation `json:"trend_location"`
}

// AccountSettingsResultSleepTime is part of the result from
// AccountService.Settings
type AccountSettingsResultSleepTime struct {
	Enabled   bool   `json:"enabled"`
	EndTime   string `json:"end_time"`
	StartTime string `json:"start_time"`
}

// AccountSettingsResultTimeZone is part of the result from
// AccountService.Settings
type AccountSettingsResultTimeZone struct {
	Name       string `json:"name"`
	TzinfoName string `json:"tzinfo_name"`
	UtcOffset  int64  `json:"utc_offset"`
}

// AccountSettingsResultTrendLocation is part of the result from
// AccountService.Settings
type AccountSettingsResultTrendLocation struct {
	Country     string                         `json:"country"`
	CountryCode string                         `json:"countryCode"`
	Name        string                         `json:"name"`
	ParentID    int64                          `json:"parentid"`
	URL         string                         `json:"url"`
	WoeID       int64                          `json:"woeid"`
	PlaceType   AccountSettingsResultPlaceType `json:"placeType"`
}

// AccountSettingsResultPlaceType is part of the result from
// AccountService.Settings
type AccountSettingsResultPlaceType struct {
	Code int64  `json:"code"`
	Name string `json:"name"`
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
