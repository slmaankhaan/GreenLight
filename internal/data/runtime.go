package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {

	// Generate  a string containing the movie runtime in the required format

	jsonValue := fmt.Sprintf("%d mins", r)

	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil

}
