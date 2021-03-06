package twitter

import (
	"net/http"

	"github.com/dghubble/sling"
)

// BlockService provides methods for accessing Twitter status API endpoints.
type BlockService struct {
	sling *sling.Sling
}

// newBlockService returns a new BlockService.
func newBlockService(sling *sling.Sling) *BlockService {
	return &BlockService{
		sling: sling.Path("blocks/"),
	}
}

// BlockServiceListResult is the result for BlockService.List
type BlockServiceListResult struct {
	PreviousCursor    int64  `json:"previous_cursor"`
	PreviousCursorStr string `json:"previous_cursor_str"`
	NextCursor        int64  `json:"next_cursor"`
	NextCursorStr     string `json:"next_cursor_str"`
	Users             []User `json:"users"`
}

// BlockServiceListParams are the parameters for BlockService.List
type BlockServiceListParams struct {
	IncludeEntities bool  `url:"block_entities,omitempty"`
	SkipStatus      bool  `url:"skip_status,omitempty"`
	Cursor          int64 `url:"cursor,omitempty"`
}

// List returns a collection of user objects that the authenticating user
// is blocking.
func (s *BlockService) List(params *BlockServiceListParams) (*BlockServiceListResult, *http.Response, error) {
	result := new(BlockServiceListResult)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("list.json").QueryStruct(params).Receive(result, apiError)
	return result, resp, relevantError(err, *apiError)
}

// BlockServiceIDsParams are the params for BlockService.IDs
type BlockServiceIDsParams struct {
	StringifyIDs bool  `url:"stringify_ids,omitempty"`
	Cursor       int64 `url:"cursor,omitempty"`
}

// BlockServiceIDsResult is the result for BlockService.IDs
type BlockServiceIDsResult struct {
	IDs               []int64 `json:"ids"`
	PreviousCursor    int64   `json:"previous_cursor"`
	PreviousCursorStr string  `json:"previous_cursor_str"`
	NextCursor        int64   `json:"next_cursor"`
	NextCursorStr     string  `json:"next_cursor_str"`
}

// IDs returns an array of numeric user ids the authenticating user is blocking.
// https://dev.twitter.com/rest/reference/get/blocks/ids
func (s *BlockService) IDs(params *BlockServiceIDsParams) (*BlockServiceIDsResult, *http.Response, error) {
	result := new(BlockServiceIDsResult)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("ids.json").QueryStruct(params).Receive(result, apiError)
	return result, resp, relevantError(err, *apiError)
}

// BlockServiceCreateParams are the params for BlockService.Create
type BlockServiceCreateParams struct {
	UserID          int64  `url:"user_id,omitempty"`
	ScreenName      string `url:"screen_name,omitempty"`
	IncludeEntities bool   `url:"block_entities,omitempty"`
	SkipStatus      bool   `url:"skip_status,omitempty"`
}

// Create blocks the specified user from following the authenticating user.
// In addition the blocked user will not show in the authenticating users
// mentions or timeline (unless retweeted by another user).
// https://dev.twitter.com/rest/reference/post/blocks/create
func (s *BlockService) Create(params *BlockServiceCreateParams) (*User, *http.Response, error) {
	result := new(User)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("create.json").QueryStruct(params).Receive(result, apiError)
	return result, resp, relevantError(err, *apiError)
}

// Destroy un-blocks the user specified in the ID parameter for the
// authenticating user. Returns the un-blocked user in the requested
// format when successful.
// https://dev.twitter.com/rest/reference/post/blocks/destroy
func (s *BlockService) Destroy(params *BlockServiceCreateParams) (*User, *http.Response, error) {
	result := new(User)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("destroy.json").QueryStruct(params).Receive(result, apiError)
	return result, resp, relevantError(err, *apiError)
}
