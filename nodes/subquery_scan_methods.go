// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node SubqueryScan) MarshalJSON() ([]byte, error) {
	type SubqueryScanMarshalAlias SubqueryScan
	return json.Marshal(map[string]interface{}{
		"SUBQUERYSCAN": (*SubqueryScanMarshalAlias)(&node),
	})
}

func (node *SubqueryScan) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node SubqueryScan) Deparse() string {
	panic("Not Implemented")
}
