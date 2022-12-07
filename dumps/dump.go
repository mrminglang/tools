package dumps

import (
	"bytes"
	"encoding/json"
	"log"
)

func Dump(data interface{}) {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.SetIndent("", "\t")
	_ = jsonEncoder.Encode(data)
	log.Printf("%s", bf.String())
}
