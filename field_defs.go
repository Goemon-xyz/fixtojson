package fixtojson

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/datadictionary"
)

type FieldDefs []*datadictionary.FieldDef

func (t FieldDefs) Find(tag quickfix.Tag) *datadictionary.FieldDef {
	for _, fieldDef := range t {
		if fieldDef.Tag() == int(tag) {
			return fieldDef
		}
	}
	return nil
}

func NewFieldDefs(Parts []datadictionary.MessagePart, MapFields map[int]*datadictionary.FieldDef) FieldDefs {
	values := make([]*datadictionary.FieldDef, 0)
	for _, part := range Parts {
		switch rpart := part.(type) {
		case datadictionary.Component:
			{
				values = append(values, rpart.Fields()...)
			}
		default:
			{
				for _, v := range MapFields {
					if v.Name() == part.Name() {
						values = append(values, v)
						break
					}
				}
			}
		}
	}
	return values
}

func NewFieldDefsFromArr(MapFields []*datadictionary.FieldDef) FieldDefs {
	return MapFields
}
