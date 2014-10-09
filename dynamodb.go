package dynamodb

import "encoding/json"

type Table struct {
	TableName string
	HashKey   string
	RangeKey  string
}

type Row struct {
	HashKey    string
	RangeKey   string
	Attributes map[string]string
}

func (t *Table) JsonToRow(j string) (row Row, err error) {
	var f map[string]map[string]map[string]string

	err = json.Unmarshal([]byte(j), &f)
	if err != nil {
		return
	}

	attrs := f["Attributes"]
	if attrs == nil {
		attrs = f["Item"]
	}
	if attrs == nil {
		err = DynamoDbError{"Item not found"}
		return
	}

	row.Attributes = make(map[string]string)

	for key, value := range attrs {
		if key == t.HashKey {
			row.HashKey = value["S"]
		} else if len(t.RangeKey) > 0 && key == t.RangeKey {
			row.RangeKey = value["S"]
		} else {
			row.Attributes[key] = value["S"]
		}
	}
	return
}
