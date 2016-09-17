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

func TestAccountService_UpdateProfile(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/account/update_profile.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertQuery(t, map[string]string{"name": "Sean Cook", "url": "http://cnn.com", "location": "San Francisco, CA", "description": "Keep calm and rock on", "profile_link_color": "0084B4", "include_entities": "true", "skip_status": "false"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"contributors_enabled":false,"created_at":"Thu Aug 23 19:45:07 +0000 2012","default_profile":false,"default_profile_image":false,"description":"Keep calm and rock on.","favourites_count":0,"follow_request_sent":false,"followers_count":0,"following":false,"friends_count":10,"geo_enabled":true,"id":776627022,"id_str":"776627022","is_translator":false,"lang":"en","listed_count":0,"location":"San Francisco, CA","name":"Sean Cook","notifications":false,"profile_background_color":"9AE4E8","profile_background_image_url":"http://a0.twimg.com/images/themes/theme16/bg.gif","profile_background_image_url_https":"https://si0.twimg.com/images/themes/theme16/bg.gif","profile_background_tile":false,"profile_image_url":"http://a0.twimg.com/profile_images/2550256790/hv5rtkvistn50nvcuydl_normal.jpeg","profile_image_url_https":"https://si0.twimg.com/profile_images/2550256790/hv5rtkvistn50nvcuydl_normal.jpeg","profile_link_color":"0084B4","profile_sidebar_border_color":"BDDCAD","profile_sidebar_fill_color":"DDFFCC","profile_text_color":"333333","profile_use_background_image":true,"protected":false,"screen_name":"s0c1alm3dia","show_all_inline_media":false,"statuses_count":0,"time_zone":"Pacific Time (US & Canada)","url":"http://cnn.com","utc_offset":-28800,"verified":false}`)
	})

	client := NewClient(httpClient)

	params := &AccountSettingsUpdateProfileParams{
		Name:             "Sean Cook",
		URL:              "http://cnn.com",
		Location:         "San Francisco, CA",
		Description:      "Keep calm and rock on",
		ProfileLinkColor: "0084B4",
		IncludeEntities:  true,
		SkipStatus:       false,
	}

	settings, _, err := client.Accounts.UpdateProfile(params)
	expected := &User{
		CreatedAt:       "Thu Aug 23 19:45:07 +0000 2012",
		Description:     "Keep calm and rock on.",
		FavouritesCount: 0,
		FollowersCount:  0,
		FriendsCount:    10,
		GeoEnabled:      true,
		ID:              776627022,
		IDStr:           "776627022",
		Lang:            "en",
		ListedCount:     0,
		Location:        "San Francisco, CA",
		Name:            "Sean Cook",
		ProfileBackgroundColor:         "9AE4E8",
		ProfileBackgroundImageURL:      "http://a0.twimg.com/images/themes/theme16/bg.gif",
		ProfileBackgroundImageURLHttps: "https://si0.twimg.com/images/themes/theme16/bg.gif",
		ProfileImageURL:                "http://a0.twimg.com/profile_images/2550256790/hv5rtkvistn50nvcuydl_normal.jpeg",
		ProfileImageURLHttps:           "https://si0.twimg.com/profile_images/2550256790/hv5rtkvistn50nvcuydl_normal.jpeg",
		ProfileLinkColor:               "0084B4",
		ProfileSidebarBorderColor:      "BDDCAD",
		ProfileSidebarFillColor:        "DDFFCC",
		ProfileTextColor:               "333333",
		ScreenName:                     "s0c1alm3dia",
		StatusesCount:                  0,
		Timezone:                       "Pacific Time (US & Canada)",
		URL:                            "http://cnn.com",
		UtcOffset:                      -28800,
		ProfileUseBackgroundImage:      true,
		ContributorsEnabled:            false,
		DefaultProfile:                 false,
		DefaultProfileImage:            false,
		FollowRequestSent:              false,
		Following:                      false,
		IsTranslator:                   false,
		Notifications:                  false,
		ProfileBackgroundTile:          false,
		Protected:                      false,
		ShowAllInlineMedia:             false,
		Verified:                       false,
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, settings)
}

func TestAccountService_UpdateProfileImage(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/account/update_profile_image.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertQuery(t, map[string]string{"image": "ABCDEFGH", "include_entities": "true"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"contributors_enabled":false,"created_at":"Thu Aug 23 19:45:07 +0000 2012","default_profile":false,"default_profile_image":false,"description":"Keep calm and rock on.","favourites_count":0,"follow_request_sent":false,"followers_count":0,"following":false,"friends_count":10,"geo_enabled":true,"id":776627022,"id_str":"776627022","is_translator":false,"lang":"en","listed_count":0,"location":"San Francisco, CA","name":"Sean Cook","notifications":false,"profile_background_color":"9AE4E8","profile_background_image_url":"http://a0.twimg.com/images/themes/theme16/bg.gif","profile_background_image_url_https":"https://si0.twimg.com/images/themes/theme16/bg.gif","profile_background_tile":false,"profile_image_url":"http://a0.twimg.com/profile_images/2550256790/hv5rtkvistn50nvcuydl_normal.jpeg","profile_image_url_https":"https://si0.twimg.com/profile_images/2550256790/hv5rtkvistn50nvcuydl_normal.jpeg","profile_link_color":"0084B4","profile_sidebar_border_color":"BDDCAD","profile_sidebar_fill_color":"DDFFCC","profile_text_color":"333333","profile_use_background_image":true,"protected":false,"screen_name":"s0c1alm3dia","show_all_inline_media":false,"statuses_count":0,"time_zone":"Pacific Time (US & Canada)","url":"http://cnn.com","utc_offset":-28800,"verified":false}`)
	})

	client := NewClient(httpClient)

	params := &AccountSettingsUpdateProfileImageParams{
		Image:           "ABCDEFGH",
		IncludeEntities: true,
		SkipStatus:      false,
	}

	settings, _, err := client.Accounts.UpdateProfileImage(params)
	expected := &User{
		CreatedAt:       "Thu Aug 23 19:45:07 +0000 2012",
		Description:     "Keep calm and rock on.",
		FavouritesCount: 0,
		FollowersCount:  0,
		FriendsCount:    10,
		GeoEnabled:      true,
		ID:              776627022,
		IDStr:           "776627022",
		Lang:            "en",
		ListedCount:     0,
		Location:        "San Francisco, CA",
		Name:            "Sean Cook",
		ProfileBackgroundColor:         "9AE4E8",
		ProfileBackgroundImageURL:      "http://a0.twimg.com/images/themes/theme16/bg.gif",
		ProfileBackgroundImageURLHttps: "https://si0.twimg.com/images/themes/theme16/bg.gif",
		ProfileImageURL:                "http://a0.twimg.com/profile_images/2550256790/hv5rtkvistn50nvcuydl_normal.jpeg",
		ProfileImageURLHttps:           "https://si0.twimg.com/profile_images/2550256790/hv5rtkvistn50nvcuydl_normal.jpeg",
		ProfileLinkColor:               "0084B4",
		ProfileSidebarBorderColor:      "BDDCAD",
		ProfileSidebarFillColor:        "DDFFCC",
		ProfileTextColor:               "333333",
		ScreenName:                     "s0c1alm3dia",
		StatusesCount:                  0,
		Timezone:                       "Pacific Time (US & Canada)",
		URL:                            "http://cnn.com",
		UtcOffset:                      -28800,
		ProfileUseBackgroundImage:      true,
		ContributorsEnabled:            false,
		DefaultProfile:                 false,
		DefaultProfileImage:            false,
		FollowRequestSent:              false,
		Following:                      false,
		IsTranslator:                   false,
		Notifications:                  false,
		ProfileBackgroundTile:          false,
		Protected:                      false,
		ShowAllInlineMedia:             false,
		Verified:                       false,
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, settings)
}

func TestAccountService_UpdateProfileBanner(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/account/update_profile_banner.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertQuery(t, map[string]string{"banner": "ABCDEFGH", "width": "600", "height": "600", "offset_left": "0", "offset_top": "0"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"contributors_enabled":false,"created_at":"Thu Aug 23 19:45:07 +0000 2012","default_profile":false,"default_profile_image":false,"description":"Keep calm and rock on.","favourites_count":0,"follow_request_sent":false,"followers_count":0,"following":false,"friends_count":10,"geo_enabled":true,"id":776627022,"id_str":"776627022","is_translator":false,"lang":"en","listed_count":0,"location":"San Francisco, CA","name":"Sean Cook","notifications":false,"profile_background_color":"9AE4E8","profile_background_image_url":"http://a0.twimg.com/images/themes/theme16/bg.gif","profile_background_image_url_https":"https://si0.twimg.com/images/themes/theme16/bg.gif","profile_background_tile":false,"profile_image_url":"http://a0.twimg.com/profile_images/2550256790/hv5rtkvistn50nvcuydl_normal.jpeg","profile_image_url_https":"https://si0.twimg.com/profile_images/2550256790/hv5rtkvistn50nvcuydl_normal.jpeg","profile_link_color":"0084B4","profile_sidebar_border_color":"BDDCAD","profile_sidebar_fill_color":"DDFFCC","profile_text_color":"333333","profile_use_background_image":true,"protected":false,"screen_name":"s0c1alm3dia","show_all_inline_media":false,"statuses_count":0,"time_zone":"Pacific Time (US & Canada)","url":"http://cnn.com","utc_offset":-28800,"verified":false}`)
	})

	client := NewClient(httpClient)

	params := &AccountSettingsUpdateProfileBannerParams{
		Banner:     "ABCDEFGH",
		Width:      600,
		Height:     600,
		OffsetLeft: 0,
		OffsetTop:  0,
	}

	settings, _, err := client.Accounts.UpdateProfileBanner(params)
	expected := &User{
		CreatedAt:       "Thu Aug 23 19:45:07 +0000 2012",
		Description:     "Keep calm and rock on.",
		FavouritesCount: 0,
		FollowersCount:  0,
		FriendsCount:    10,
		GeoEnabled:      true,
		ID:              776627022,
		IDStr:           "776627022",
		Lang:            "en",
		ListedCount:     0,
		Location:        "San Francisco, CA",
		Name:            "Sean Cook",
		ProfileBackgroundColor:         "9AE4E8",
		ProfileBackgroundImageURL:      "http://a0.twimg.com/images/themes/theme16/bg.gif",
		ProfileBackgroundImageURLHttps: "https://si0.twimg.com/images/themes/theme16/bg.gif",
		ProfileImageURL:                "http://a0.twimg.com/profile_images/2550256790/hv5rtkvistn50nvcuydl_normal.jpeg",
		ProfileImageURLHttps:           "https://si0.twimg.com/profile_images/2550256790/hv5rtkvistn50nvcuydl_normal.jpeg",
		ProfileLinkColor:               "0084B4",
		ProfileSidebarBorderColor:      "BDDCAD",
		ProfileSidebarFillColor:        "DDFFCC",
		ProfileTextColor:               "333333",
		ScreenName:                     "s0c1alm3dia",
		StatusesCount:                  0,
		Timezone:                       "Pacific Time (US & Canada)",
		URL:                            "http://cnn.com",
		UtcOffset:                      -28800,
		ProfileUseBackgroundImage:      true,
		ContributorsEnabled:            false,
		DefaultProfile:                 false,
		DefaultProfileImage:            false,
		FollowRequestSent:              false,
		Following:                      false,
		IsTranslator:                   false,
		Notifications:                  false,
		ProfileBackgroundTile:          false,
		Protected:                      false,
		ShowAllInlineMedia:             false,
		Verified:                       false,
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, settings)
}

func TestAccountService_RemoveProfileBanner(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/account/remove_profile_banner.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"contributors_enabled":false,"created_at":"Thu Aug 23 19:45:07 +0000 2012","default_profile":false,"default_profile_image":false,"description":"Keep calm and rock on.","favourites_count":0,"follow_request_sent":false,"followers_count":0,"following":false,"friends_count":10,"geo_enabled":true,"id":776627022,"id_str":"776627022","is_translator":false,"lang":"en","listed_count":0,"location":"San Francisco, CA","name":"Sean Cook","notifications":false,"profile_background_color":"9AE4E8","profile_background_image_url":"http://a0.twimg.com/images/themes/theme16/bg.gif","profile_background_image_url_https":"https://si0.twimg.com/images/themes/theme16/bg.gif","profile_background_tile":false,"profile_image_url":"http://a0.twimg.com/profile_images/2550256790/hv5rtkvistn50nvcuydl_normal.jpeg","profile_image_url_https":"https://si0.twimg.com/profile_images/2550256790/hv5rtkvistn50nvcuydl_normal.jpeg","profile_link_color":"0084B4","profile_sidebar_border_color":"BDDCAD","profile_sidebar_fill_color":"DDFFCC","profile_text_color":"333333","profile_use_background_image":true,"protected":false,"screen_name":"s0c1alm3dia","show_all_inline_media":false,"statuses_count":0,"time_zone":"Pacific Time (US & Canada)","url":"http://cnn.com","utc_offset":-28800,"verified":false}`)
	})

	client := NewClient(httpClient)

	settings, _, err := client.Accounts.RemoveProfileBanner()
	expected := &User{
		CreatedAt:       "Thu Aug 23 19:45:07 +0000 2012",
		Description:     "Keep calm and rock on.",
		FavouritesCount: 0,
		FollowersCount:  0,
		FriendsCount:    10,
		GeoEnabled:      true,
		ID:              776627022,
		IDStr:           "776627022",
		Lang:            "en",
		ListedCount:     0,
		Location:        "San Francisco, CA",
		Name:            "Sean Cook",
		ProfileBackgroundColor:         "9AE4E8",
		ProfileBackgroundImageURL:      "http://a0.twimg.com/images/themes/theme16/bg.gif",
		ProfileBackgroundImageURLHttps: "https://si0.twimg.com/images/themes/theme16/bg.gif",
		ProfileImageURL:                "http://a0.twimg.com/profile_images/2550256790/hv5rtkvistn50nvcuydl_normal.jpeg",
		ProfileImageURLHttps:           "https://si0.twimg.com/profile_images/2550256790/hv5rtkvistn50nvcuydl_normal.jpeg",
		ProfileLinkColor:               "0084B4",
		ProfileSidebarBorderColor:      "BDDCAD",
		ProfileSidebarFillColor:        "DDFFCC",
		ProfileTextColor:               "333333",
		ScreenName:                     "s0c1alm3dia",
		StatusesCount:                  0,
		Timezone:                       "Pacific Time (US & Canada)",
		URL:                            "http://cnn.com",
		UtcOffset:                      -28800,
		ProfileUseBackgroundImage:      true,
		ContributorsEnabled:            false,
		DefaultProfile:                 false,
		DefaultProfileImage:            false,
		FollowRequestSent:              false,
		Following:                      false,
		IsTranslator:                   false,
		Notifications:                  false,
		ProfileBackgroundTile:          false,
		Protected:                      false,
		ShowAllInlineMedia:             false,
		Verified:                       false,
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, settings)
}
