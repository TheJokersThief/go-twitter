package twitter

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/blocks/list.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"skip_status": "true", "cursor": "-1"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"next_cursor":0,"next_cursor_str":"0","previous_cursor":0,"previous_cursor_str":"0","users":[{"contributors_enabled":false,"created_at":"Thu Mar 01 00:16:47 +0000 2012","default_profile":true,"default_profile_image":true,"description":"","favourites_count":0,"follow_request_sent":false,"followers_count":0,"following":false,"friends_count":0,"geo_enabled":false,"id":509466276,"id_str":"509466276","is_translator":false,"lang":"en","listed_count":0,"location":"","name":"Javier Heady \r","notifications":false,"profile_background_color":"C0DEED","profile_background_image_url":"http://a0.twimg.com/images/themes/theme1/bg.png","profile_background_image_url_https":"https://si0.twimg.com/images/themes/theme1/bg.png","profile_background_tile":false,"profile_image_url":"http://a0.twimg.com/sticky/default_profile_images/default_profile_4_normal.png","profile_image_url_https":"https://si0.twimg.com/sticky/default_profile_images/default_profile_4_normal.png","profile_link_color":"0084B4","profile_sidebar_border_color":"C0DEED","profile_sidebar_fill_color":"DDEEF6","profile_text_color":"333333","profile_use_background_image":true,"protected":false,"screen_name":"javierg3ong","statuses_count":4,"time_zone":null,"url":null,"utc_offset":null,"verified":false}]}`)
	})

	client := NewClient(httpClient)

	params := &BlockServiceListParams{
		IncludeEntities: false,
		SkipStatus:      true,
		Cursor:          -1,
	}
	result, _, err := client.Block.List(params)
	expected := &BlockServiceListResult{
		PreviousCursor:    0,
		PreviousCursorStr: "0",
		NextCursor:        0,
		NextCursorStr:     "0",
		Users: []User{
			User{
				ProfileSidebarFillColor:   "DDEEF6",
				ProfileBackgroundTile:     false,
				ProfileSidebarBorderColor: "C0DEED",
				Name:                           "Javier Heady \r",
				CreatedAt:                      "Thu Mar 01 00:16:47 +0000 2012",
				ProfileImageURL:                "http://a0.twimg.com/sticky/default_profile_images/default_profile_4_normal.png",
				Location:                       "",
				IsTranslator:                   false,
				FollowRequestSent:              false,
				ProfileLinkColor:               "0084B4",
				IDStr:                          "509466276",
				ContributorsEnabled:            false,
				FavouritesCount:                0,
				URL:                            "",
				DefaultProfile:                 true,
				UtcOffset:                      0,
				ProfileImageURLHttps:           "https://si0.twimg.com/sticky/default_profile_images/default_profile_4_normal.png",
				ID:                             509466276,
				ListedCount:                    0,
				ProfileUseBackgroundImage:      true,
				FollowersCount:                 0,
				Protected:                      false,
				Lang:                           "en",
				ProfileTextColor:               "333333",
				ProfileBackgroundColor:         "C0DEED",
				Notifications:                  false,
				Verified:                       false,
				Description:                    "",
				GeoEnabled:                     false,
				Timezone:                       "",
				ProfileBackgroundImageURLHttps: "https://si0.twimg.com/images/themes/theme1/bg.png",
				FriendsCount:                   0,
				DefaultProfileImage:            true,
				StatusesCount:                  4,
				ProfileBackgroundImageURL:      "http://a0.twimg.com/images/themes/theme1/bg.png",
				Following:                      false,
				ScreenName:                     "javierg3ong",
			},
		},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}
