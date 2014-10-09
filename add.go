package dynamodb

import (
	ep "github.com/smugmug/godynamo/endpoint"
	update_item "github.com/smugmug/godynamo/endpoints/update_item"
)

const (
	ACTION_PUT = "PUT"
	ACTION_DEL = "DELETE"
	ACTION_ADD = "ADD"
)

type AttrUpdate struct {
	Value  string
	Action string
}

func (t *Table) Add(attributes map[string]AttrUpdate,
	keys ...string) (row Row, err error) {
	var up1 update_item.Request

	up1.TableName = t.TableName
	up1.Key = make(ep.Item)
	up1.Key[t.HashKey] = ep.AttributeValue{S: keys[0]}
	if len(t.RangeKey) > 0 && len(keys) > 1 {
		up1.Key[t.RangeKey] = ep.AttributeValue{S: keys[1]}
	}
	up1.AttributeUpdates = make(update_item.AttributeUpdates)

	for key, attr := range attributes {
		up1.AttributeUpdates[key] = update_item.AttributeAction{
			Value:  ep.AttributeValue{S: attr.Value},
			Action: attr.Action,
		}
	}

	up1.ReturnValues = update_item.RETVAL_ALL_NEW

	body, code, err := up1.EndpointReq()
	if code != 200 || err != nil {
		return
	}

	row, err = t.JsonToRow(body)
	if err != nil {
		return
	}

	return
}
