package api

import (
    "encoding/json"
)

type Response map[string]interface{}

func (r Response) String() string {
    bytes, err := json.Marshal(r)

    if err != nil {
        return ""
    }

    return string(bytes)
}