package twitter

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFriendship_Lookup(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/friendships/lookup.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"screen_name": "dghubble"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"name":"Dalton Hubble","screen_name":"dghubble","id":623265148,"id_str":"623265148","connections":["none"]}]`)
	})
	expected := &[]FriendshipLookupStatus{
		FriendshipLookupStatus{
			Name:        "Dalton Hubble",
			ScreenName:  "dghubble",
			ID:          623265148,
			IDStr:       "623265148",
			Connections: []string{"none"},
		},
	}

	client := NewClient(httpClient)
	params := &FriendshipLookupParams{
		ScreenName: "dghubble",
	}
	friendshipStatus, _, err := client.Friendships.Lookup(params)
	assert.Nil(t, err)
	assert.Equal(t, expected, friendshipStatus)
}

func TestFriendship_Show(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/friendships/show.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"source_screen_name": "thejokersthief", "target_screen_name": "dghubble"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"relationship":{"source":{"id":54655541,"id_str":"54655541","screen_name":"TheJokersThief","following":false,"followed_by":false,"live_following":false,"following_received":false,"following_requested":false,"notifications_enabled":false,"can_dm":false,"blocking":false,"blocked_by":false,"muting":false,"want_retweets":false,"all_replies":false,"marked_spam":false},"target":{"id":623265148,"id_str":"623265148","screen_name":"dghubble","following":false,"followed_by":false,"following_received":false,"following_requested":false}}}`)
	})
	expected := &FriendshipShowResult{
		Relationship: FriendshipRelationship{
			Target: FriendshipRelationshipTarget{
				IDStr:      "623265148",
				ID:         623265148,
				ScreenName: "dghubble",
				Following:  false,
				FollowedBy: false,
			},
			Source: FriendshipRelationshipSource{
				CanDM:                false,
				Blocking:             false,
				Muting:               false,
				IDStr:                "54655541",
				AllReplies:           false,
				WantRetweets:         false,
				ID:                   54655541,
				MarkedSpam:           false,
				ScreenName:           "TheJokersThief",
				Following:            false,
				FollowedBy:           false,
				NotificationsEnabled: false,
			},
		},
	}

	client := NewClient(httpClient)
	params := &FriendshipShowParams{
		SourceScreenName: "thejokersthief",
		TargetScreenName: "dghubble",
	}
	friendshipStatus, _, err := client.Friendships.Show(params)
	assert.Nil(t, err)
	assert.Equal(t, expected, friendshipStatus)
}

func TestFriendship_Create(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/friendships/create.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertQuery(t, map[string]string{"screen_name": "dghubble"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"screen_name":"dghubble","id":623265148}`)
	})
	expected := &User{
		ScreenName: "dghubble",
		ID:         623265148,
	}

	client := NewClient(httpClient)
	params := &FriendshipLookupParams{
		ScreenName: "dghubble",
	}
	user, _, err := client.Friendships.Create(params)
	assert.Nil(t, err)
	assert.Equal(t, expected, user)
}

func TestFriendship_Destroy(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/friendships/destroy.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertQuery(t, map[string]string{"screen_name": "dghubble"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"screen_name":"dghubble","id":623265148}`)
	})
	expected := &User{
		ScreenName: "dghubble",
		ID:         623265148,
	}

	client := NewClient(httpClient)
	params := &FriendshipLookupParams{
		ScreenName: "dghubble",
	}
	user, _, err := client.Friendships.Destroy(params)
	assert.Nil(t, err)
	assert.Equal(t, expected, user)
}

func TestFriendship_Update(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/friendships/update.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertQuery(t, map[string]string{"screen_name": "thejokersthief", "retweets": "true"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"relationship":{"source":{"id":54655541,"id_str":"54655541","screen_name":"TheJokersThief","following":false,"followed_by":false,"live_following":false,"following_received":false,"following_requested":false,"notifications_enabled":false,"can_dm":false,"blocking":false,"blocked_by":false,"muting":false,"want_retweets":true,"all_replies":false,"marked_spam":false},"target":{"id":623265148,"id_str":"623265148","screen_name":"dghubble","following":false,"followed_by":false,"following_received":false,"following_requested":false}}}`)
	})
	expected := &FriendshipShowResult{
		Relationship: FriendshipRelationship{
			Target: FriendshipRelationshipTarget{
				IDStr:      "623265148",
				ID:         623265148,
				ScreenName: "dghubble",
				Following:  false,
				FollowedBy: false,
			},
			Source: FriendshipRelationshipSource{
				CanDM:                false,
				Blocking:             false,
				Muting:               false,
				IDStr:                "54655541",
				AllReplies:           false,
				WantRetweets:         true,
				ID:                   54655541,
				MarkedSpam:           false,
				ScreenName:           "TheJokersThief",
				Following:            false,
				FollowedBy:           false,
				NotificationsEnabled: false,
			},
		},
	}

	client := NewClient(httpClient)
	params := &FriendshipUpdateParams{
		ScreenName: "thejokersthief",
		Retweets:   true,
	}
	friendshipStatus, _, err := client.Friendships.Update(params)
	assert.Nil(t, err)
	assert.Equal(t, expected, friendshipStatus)
}

func TestFriendship_Incoming(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/friendships/incoming.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"cursor": "-1"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"previous_cursor":0,"ids": [1, 2, 3], "next_cursor": 123}`)
	})
	expected := &FriendshipIncomingResult{
		PreviousCursor: 0,
		IDs:            []int64{1, 2, 3},
		NextCursor:     123,
	}

	client := NewClient(httpClient)
	params := &FriendshipIncomingParams{
		Cursor: -1,
	}
	user, _, err := client.Friendships.Incoming(params)
	assert.Nil(t, err)
	assert.Equal(t, expected, user)
}

func TestFriendship_NoRetweets(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/friendships/no_retweets/ids.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[777925,732321]`)
	})
	expected := &[]int64{
		777925,
		732321,
	}

	client := NewClient(httpClient)
	user, _, err := client.Friendships.NoRetweets()
	assert.Nil(t, err)
	assert.Equal(t, expected, user)
}
