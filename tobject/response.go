package tobject

import "encoding/json"

type Response struct {
	OK     bool            `json:"ok"`
	Result json.RawMessage `json:"result"`
}
