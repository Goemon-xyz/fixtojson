package fixtojson

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/quickfixgo/field"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/datadictionary"
)

type Converter struct {
	dictionary *datadictionary.DataDictionary
}

func NewConverter(dictionaryPath string) (*Converter, error) {
	dictionary, err := datadictionary.Parse(dictionaryPath)
	if err != nil {
		return nil, fmt.Errorf("datadictionary.Parse() -> ERROR: %v", err)
	}
	return &Converter{dictionary: dictionary}, nil
}

func (c *Converter) FIXToJSON(raw []byte) ([]byte, error) {
	msg := quickfix.NewMessage()

	err := quickfix.ParseMessage(msg, bytes.NewBuffer(raw))
	if err != nil {
		return nil, fmt.Errorf("error parsing message: %v", err)
	}

	var msgType field.MsgTypeField
	msg.Header.Get(&msgType)

	MessageDesc := c.dictionary.Messages[string(msgType.Value())]

	messageMap := make(map[string]interface{})

	messageMap["Header"] = c.FieldMapToJSON(msg.Header.FieldMap, NewFieldDefs(c.dictionary.Header.Parts, c.dictionary.Header.Fields))
	messageMap["Body"] = c.FieldMapToJSON(msg.Body.FieldMap, NewFieldDefs(MessageDesc.Parts, MessageDesc.Fields))
	messageMap["Trailer"] = c.FieldMapToJSON(msg.Trailer.FieldMap, NewFieldDefs(c.dictionary.Trailer.Parts, c.dictionary.Trailer.Fields))

	jsonData, err := json.MarshalIndent(messageMap, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("error converting to JSON: %v", err)
	}

	return jsonData, nil
}

func (c *Converter) FieldMapToJSON(fm quickfix.FieldMap, fieldDefs FieldDefs) map[string]interface{} {
	jsonMap := make(map[string]interface{})

	for _, gFieldDef := range fieldDefs {
		var value quickfix.FIXBytes
		tag := quickfix.Tag(gFieldDef.Tag())
		err := fm.GetField(tag, &value)
		if err != nil {
			continue
		}

		fieldDesc, found := c.dictionary.FieldTypeByTag[int(tag)]
		if found {
			strValue := string(value)
			switch fieldDesc.Type {
			case "NUMINGROUP":
				groupField := fieldDefs.Find(tag)
				if groupField == nil {
					fmt.Printf("ERROR: Missing group field definition for tag [%d]\n", tag)
					continue
				}
				groupTemplate := quickfix.GroupTemplate{}
				for _, subField := range groupField.Fields {
					groupTemplate = append(groupTemplate, quickfix.GroupElement(quickfix.Tag(subField.Tag())))
				}
				group := quickfix.NewRepeatingGroup(tag, groupTemplate)
				err := fm.GetGroup(group)
				if err != nil {
					fmt.Printf("ERROR: Unable to parse repeating group [%d]: %v\n", tag, err)
					continue
				}
				groupArray := []map[string]interface{}{}
				for i := 0; i < group.Len(); i++ {
					g := group.Get(i)
					groupArray = append(groupArray, c.FieldMapToJSON(g.FieldMap, NewFieldDefsFromArr(groupField.Fields)))
				}
				jsonMap[fieldDesc.Name()] = groupArray
			default:
				if isJSON(strValue) {
					var nestedJSON map[string]interface{}
					if err := json.Unmarshal([]byte(strValue), &nestedJSON); err == nil {
						jsonMap[fieldDesc.Name()] = nestedJSON
					} else {
						fmt.Printf("ERROR: Unable to parse JSON for tag [%d]: %v\n", tag, err)
					}
				} else {
					jsonMap[fieldDesc.Name()] = strValue
				}
			}
		} else {
			jsonMap[fmt.Sprintf("Unknown_%d", tag)] = string(value)
		}
	}

	return jsonMap
}

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}
