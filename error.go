package dynamodb

type DynamoDbError struct {
	s string
}

func (e DynamoDbError) Error() string {
	return e.s
}
