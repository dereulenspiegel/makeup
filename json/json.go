package json

import (
	"encoding/json"
	"os"
)

type JsonMakeup struct {
}

func (j *JsonMakeup) prettifyJson(temp interface{}) (out []byte, err error) {
	out, err = json.MarshalIndent(temp, "", "  ")
	return
}

func (j *JsonMakeup) Prettify(data []byte) (out []byte, err error) {
	var temp interface{}
	err = json.Unmarshal(data, &temp)
	if err != nil {
		return
	}
	return j.prettifyJson(temp)
}

func (j *JsonMakeup) PrettifyFile(path string) (out []byte, err error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return
	}
	var temp interface{}
	jsonParser := json.NewDecoder(jsonFile)
	if err = jsonParser.Decode(&temp); err != nil {
		return
	}
	return j.prettifyJson(temp)
}
