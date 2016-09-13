package twitter

import (
	"net/http"

	"github.com/dghubble/sling"
)

// FriendshipService provides methods for accessing Twitter friendship endpoints.
type FriendshipService struct {
	sling *sling.Sling
}

// Creates a new friendship service
func newFriendshipService(sling *sling.Sling) *FriendshipService {
	return &FriendshipService{
		sling: sling.Path("friendships/"),
	}
}

// FriendshipLookupStatus is The relationship status between the authenticated user and the target
type FriendshipLookupStatus struct {
	Name        string   `json:"name"`
	ScreenName  string   `json:"screen_name"`
	ID          int64    `json:"id"`
	IDStr       string   `json:"id_str"`
	Connections []string `json:"connections"`
}

// FriendshipLookupParams are Basic parameters for friendship requests
type FriendshipLookupParams struct {
	UserID     string `url:"user_id,omitempty"`
	ScreenName string `url:"screen_name,omitempty"`
}

// Lookup returns the relationships of the authenticating user to target user
func (s *FriendshipService) Lookup(params *FriendshipLookupParams) (*[]FriendshipLookupStatus, *http.Response, error) {
	friendships := new([]FriendshipLookupStatus)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("lookup.json").QueryStruct(params).Receive(friendships, apiError)
	return friendships, resp, relevantError(err, *apiError)
}

// FriendshipShowResult is the result from the Friendship show function
type FriendshipShowResult struct {
	Relationship FriendshipRelationship `json:"relationship"`
}

// FriendshipRelationship is the underlying relationship of the show function
type FriendshipRelationship struct {
	Target FriendshipRelationshipTarget `json:"target"`
	Source FriendshipRelationshipSource `json:"source"`
}

// FriendshipRelationshipTarget is the target's attributes from the show function
type FriendshipRelationshipTarget struct {
	IDStr      string `json:"id_str"`
	ID         int64  `json:"id"`
	ScreenName string `json:"screen_name"`
	Following  bool   `json:"following"`
	FollowedBy bool   `json:"followed_by"`
}

// FriendshipRelationshipSource is the source's attributes from the show function
type FriendshipRelationshipSource struct {
	CanDM                bool   `json:"can_dm"`
	Blocking             bool   `json:"blocking"`
	Muting               bool   `json:"muting"`
	IDStr                string `json:"id_str"`
	AllReplies           bool   `json:"all_replies"`
	WantRetweets         bool   `json:"want_retweets"`
	ID                   int64  `json:"id"`
	MarkedSpam           bool   `json:"marked_spam"`
	ScreenName           string `json:"screen_name"`
	Following            bool   `json:"following"`
	FollowedBy           bool   `json:"followed_by"`
	NotificationsEnabled bool   `json:"notifications_enabled"`
}

// FriendshipShowParams are the parameters given to the show function
type FriendshipShowParams struct {
	SourceScreenName string `url:"source_screen_name,omitempty"`
	SourceID         string `url:"source_id,omitempty"`
	TargetScreenName string `url:"target_screen_name,omitempty"`
	TargetID         string `url:"target_id,omitempty"`
}

// Show Returns detailed information about the relationship between two
// arbitrary users.
func (s *FriendshipService) Show(params *FriendshipShowParams) (*FriendshipShowResult, *http.Response, error) {
	friendships := new(FriendshipShowResult)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("show.json").QueryStruct(params).Receive(friendships, apiError)
	return friendships, resp, relevantError(err, *apiError)
}

// FriendshipIncomingParams are the parameters given to
// FriendshipService.Incoming
type FriendshipIncomingParams struct {
	Cursor       int64 `url:"cursor,omitempty"`
	StringifyIDs bool  `url:"stringify_ids,omitempty"`
}

// FriendshipIncomingResult is the result from FriendshipService.Incoming
type FriendshipIncomingResult struct {
	NextCursor        int64   `json:"next_cursor"`
	NextCursorStr     string  `json:"next_cursor_str"`
	PreviousCursor    int64   `json:"previous_cursor"`
	PreviousCursorStr string  `json:"previous_cursor_str"`
	IDs               []int64 `json:"ids"`
}

// Incoming returns a collection of numeric IDs for every user who has a
// pending request to follow the authenticating user.
// https://dev.twitter.com/rest/reference/get/friendships/incoming
func (s *FriendshipService) Incoming(params *FriendshipIncomingParams) (*FriendshipIncomingResult, *http.Response, error) {
	result := new(FriendshipIncomingResult)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("incoming.json").QueryStruct(params).Receive(result, apiError)
	return result, resp, relevantError(err, *apiError)
}

// Outgoing returns a collection of numeric IDs for every protected user for
// whom the authenticating user has a pending follow request.
// https://dev.twitter.com/rest/reference/get/friendships/outgoing
func (s *FriendshipService) Outgoing(params *FriendshipIncomingParams) (*FriendshipIncomingResult, *http.Response, error) {
	result := new(FriendshipIncomingResult)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("outgoing.json").QueryStruct(params).Receive(result, apiError)
	return result, resp, relevantError(err, *apiError)
}

// NoRetweets reeturns a collection of user_ids that the currently
// authenticated user does not want to receive retweets from.
// https://dev.twitter.com/rest/reference/get/friendships/no_retweets/ids
func (s *FriendshipService) NoRetweets() (*[]int64, *http.Response, error) {
	result := new([]int64)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("no_retweets/ids.json").Receive(result, apiError)
	return result, resp, relevantError(err, *apiError)
}

// Destroy allows the authenticating user to unfollow the user specified in
// the ID/ScreenName parameter.
func (s *FriendshipService) Destroy(params *FriendshipLookupParams) (*User, *http.Response, error) {
	friendships := new(User)

	apiError := new(APIError)
	resp, err := s.sling.New().Post("destroy.json").QueryStruct(params).Receive(friendships, apiError)
	return friendships, resp, relevantError(err, *apiError)
}

// Create allows the authenticating users to follow the user specified in
// the ID/ScreenName parameter.
func (s *FriendshipService) Create(params *FriendshipLookupParams) (*User, *http.Response, error) {
	friendships := new(User)

	apiError := new(APIError)
	resp, err := s.sling.New().Post("create.json").QueryStruct(params).Receive(friendships, apiError)
	return friendships, resp, relevantError(err, *apiError)
}

// FriendshipUpdateParams are the parameters given to the Update function
type FriendshipUpdateParams struct {
	ScreenName string `url:"screen_name,omitempty"` // The screen name of the user for whom to befriend.
	SourceID   string `url:"user_id,omitempty"`     // The ID of the user for whom to befriend.
	Device     bool   `url:"device,omitempty"`      // Enable/disable device notifications from the target user.
	Retweets   bool   `url:"retweets,omitempty"`    // Enable/disable retweets from the target user.
}

// Update allows one to enable or disable retweets and device notifications
// from the specified user.
func (s *FriendshipService) Update(params *FriendshipUpdateParams) (*FriendshipShowResult, *http.Response, error) {
	friendship := new(FriendshipShowResult)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("update.json").QueryStruct(params).Receive(friendship, apiError)
	return friendship, resp, relevantError(err, *apiError)
}
