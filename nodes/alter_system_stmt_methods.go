// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node AlterSystemStmt) MarshalJSON() ([]byte, error) {
	type AlterSystemStmtMarshalAlias AlterSystemStmt
	return json.Marshal(map[string]interface{}{
		"ALTERSYSTEMSTMT": (*AlterSystemStmtMarshalAlias)(&node),
	})
}

func (node *AlterSystemStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node AlterSystemStmt) Deparse() string {
	panic("Not Implemented")
}
