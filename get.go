package dynamodb

import (
	ep "github.com/smugmug/godynamo/endpoint"
	get "github.com/smugmug/godynamo/endpoints/get_item"
)

// Get("hash_key", "range_key") || Get("hash_key")
func (t *Table) Get(keys ...string) (row Row, err error) {
	var get1 get.Request

	get1.TableName = t.TableName
	get1.Key = make(ep.Item)
	get1.Key[t.HashKey] = ep.AttributeValue{S: keys[0]}

	if len(t.RangeKey) > 0 && len(keys) > 1 {
		get1.Key[t.RangeKey] = ep.AttributeValue{S: keys[1]}
	}

	body, code, err := get1.EndpointReq()
	if code != 200 || err != nil {
		return
	}

	row, err = t.JsonToRow(body)
	if err != nil {
		return
	}

	return
}
