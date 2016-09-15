package twitter

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountService_VerifyCredentials(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/account/verify_credentials.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"include_entities": "false", "include_email": "true"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"name": "Dalton Hubble", "id": 623265148}`)
	})

	client := NewClient(httpClient)
	user, _, err := client.Accounts.VerifyCredentials(&AccountVerifyParams{IncludeEntities: Bool(false), IncludeEmail: Bool(true)})
	expected := &User{Name: "Dalton Hubble", ID: 623265148}
	assert.Nil(t, err)
	assert.Equal(t, expected, user)
}

func TestAccountService_Settings(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/account/settings.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"always_use_https":true,"discoverable_by_email":true,"geo_enabled":true,"language":"en","protected":false,"screen_name":"theSeanCook","show_all_inline_media":false,"sleep_time":{"enabled":false,"end_time":null,"start_time":null},"time_zone":{"name":"Pacific Time (US & Canada)","tzinfo_name":"America/Los_Angeles","utc_offset":-28800},"trend_location":[{"country":"United States","countryCode":"US","name":"Atlanta","parentid":23424977,"placeType":{"code":7,"name":"Town"},"url":"http://where.yahooapis.com/v1/place/2357024","woeid":2357024}],"use_cookie_personalization":true,"allow_contributor_request":"all"}`)
	})

	client := NewClient(httpClient)
	settings, _, err := client.Accounts.Settings()
	expected := &AccountSettingsResult{
		AlwaysUseHTTPS:           true,
		DiscoverableByEmail:      true,
		GeoEnabled:               true,
		Language:                 "en",
		Protected:                false,
		ScreenName:               "theSeanCook",
		ShowAllInlineMedia:       false,
		UseCookiePersonalization: true,
		AllowContributorRequest:  "all",
		SleepTime: AccountSettingsResultSleepTime{
			Enabled:   false,
			EndTime:   "",
			StartTime: "",
		},
		TimeZone: AccountSettingsResultTimeZone{
			Name:       "Pacific Time (US & Canada)",
			TzinfoName: "America/Los_Angeles",
			UtcOffset:  -28800,
		},
		TrendLocation: []AccountSettingsResultTrendLocation{
			AccountSettingsResultTrendLocation{
				Country:     "United States",
				CountryCode: "US",
				Name:        "Atlanta",
				ParentID:    23424977,
				PlaceType: AccountSettingsResultPlaceType{
					Code: 7,
					Name: "Town",
				},
				URL:   "http://where.yahooapis.com/v1/place/2357024",
				WoeID: 2357024,
			},
		},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, settings)
}

func TestAccountService_UpdateSettings(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/account/settings.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"always_use_https":true,"discoverable_by_email":false,"geo_enabled":false,"language":"en","protected":false,"screen_name":"theSeanCook","show_all_inline_media":false,"sleep_time":{"enabled":false,"end_time":null,"start_time":null},"time_zone":{"name":"Pacific Time (US & Canada)","tzinfo_name":"America/Los_Angeles","utc_offset":-28800},"trend_location":[{"country":"United States","countryCode":"US","name":"Atlanta","parentid":23424977,"placeType":{"code":7,"name":"Town"},"url":"http://where.yahooapis.com/v1/place/2357024","woeid":2357024}],"use_cookie_personalization":true,"allow_contributor_request":"all"}`)
	})

	client := NewClient(httpClient)
	expected := &AccountSettingsResult{
		AlwaysUseHTTPS:           true,
		DiscoverableByEmail:      false,
		GeoEnabled:               false,
		Language:                 "en",
		Protected:                false,
		ScreenName:               "theSeanCook",
		ShowAllInlineMedia:       false,
		UseCookiePersonalization: true,
		AllowContributorRequest:  "all",
		SleepTime: AccountSettingsResultSleepTime{
			Enabled:   false,
			EndTime:   "",
			StartTime: "",
		},
		TimeZone: AccountSettingsResultTimeZone{
			Name:       "Pacific Time (US & Canada)",
			TzinfoName: "America/Los_Angeles",
			UtcOffset:  -28800,
		},
		TrendLocation: []AccountSettingsResultTrendLocation{
			AccountSettingsResultTrendLocation{
				Country:     "United States",
				CountryCode: "US",
				Name:        "Atlanta",
				ParentID:    23424977,
				PlaceType: AccountSettingsResultPlaceType{
					Code: 7,
					Name: "Town",
				},
				URL:   "http://where.yahooapis.com/v1/place/2357024",
				WoeID: 2357024,
			},
		},
	}

	settings, _, err := client.Accounts.UpdateSettings(expected)
	assert.Nil(t, err)
	assert.Equal(t, expected, settings)
}
