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
