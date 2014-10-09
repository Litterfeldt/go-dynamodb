package dynamodb

import (
	ep "github.com/smugmug/godynamo/endpoint"
	delete_item "github.com/smugmug/godynamo/endpoints/delete_item"
)

// Del("hash_key", "range_key") || Del("hash_key")
func (t *Table) Del(keys ...string) (err error) {
	var del1 delete_item.Request

	del1.Key = make(ep.Item)
	del1.TableName = t.TableName
	del1.Key[t.HashKey] = ep.AttributeValue{S: keys[0]}
	if len(t.RangeKey) > 0 && len(keys) > 1 {
		del1.Key[t.RangeKey] = ep.AttributeValue{S: keys[1]}
	}
	del1.ReturnValues = delete_item.RETVAL_ALL_OLD

	_, _, err = del1.EndpointReq()
	return
}
