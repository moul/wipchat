package wipchat

import "encoding/json"

func typeToType(in, out interface{}) error {
	buf, err := json.Marshal(in)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, out)
}
