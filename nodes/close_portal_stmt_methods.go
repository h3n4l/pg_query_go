// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node ClosePortalStmt) MarshalJSON() ([]byte, error) {
	type ClosePortalStmtMarshalAlias ClosePortalStmt
	return json.Marshal(map[string]interface{}{
		"CLOSEPORTALSTMT": (*ClosePortalStmtMarshalAlias)(&node),
	})
}

func (node *ClosePortalStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node ClosePortalStmt) Deparse() string {
	panic("Not Implemented")
}
