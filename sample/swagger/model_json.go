// generated by jwg -output model_json.go -transcripttag swagger .; DO NOT EDIT

package main

import (
	"encoding/json"
	"time"
)

// TodoJSON is jsonized struct for Todo.
type TodoJSON struct {
	ID        int64     `json:"id,omitempty,string"`
	Text      string    `json:"text,omitempty" swagger:",req"`
	Done      bool      `json:"done,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// TodoJSONList is synonym about []*TodoJSON.
type TodoJSONList []*TodoJSON

// TodoPropertyEncoder is property encoder for [1]sJSON.
type TodoPropertyEncoder func(src *Todo, dest *TodoJSON) error

// TodoPropertyDecoder is property decoder for [1]sJSON.
type TodoPropertyDecoder func(src *TodoJSON, dest *Todo) error

// TodoPropertyInfo stores property information.
type TodoPropertyInfo struct {
	name    string
	Encoder TodoPropertyEncoder
	Decoder TodoPropertyDecoder
}

// TodoJSONBuilder convert between Todo to TodoJSON mutually.
type TodoJSONBuilder struct {
	_properties map[string]*TodoPropertyInfo
	ID          *TodoPropertyInfo
	Text        *TodoPropertyInfo
	Done        *TodoPropertyInfo
	CreatedAt   *TodoPropertyInfo
}

// NewTodoJSONBuilder make new TodoJSONBuilder.
func NewTodoJSONBuilder() *TodoJSONBuilder {
	return &TodoJSONBuilder{
		_properties: map[string]*TodoPropertyInfo{},
		ID: &TodoPropertyInfo{
			name: "ID",
			Encoder: func(src *Todo, dest *TodoJSON) error {
				if src == nil {
					return nil
				}
				dest.ID = src.ID
				return nil
			},
			Decoder: func(src *TodoJSON, dest *Todo) error {
				if src == nil {
					return nil
				}
				dest.ID = src.ID
				return nil
			},
		},
		Text: &TodoPropertyInfo{
			name: "Text",
			Encoder: func(src *Todo, dest *TodoJSON) error {
				if src == nil {
					return nil
				}
				dest.Text = src.Text
				return nil
			},
			Decoder: func(src *TodoJSON, dest *Todo) error {
				if src == nil {
					return nil
				}
				dest.Text = src.Text
				return nil
			},
		},
		Done: &TodoPropertyInfo{
			name: "Done",
			Encoder: func(src *Todo, dest *TodoJSON) error {
				if src == nil {
					return nil
				}
				dest.Done = src.Done
				return nil
			},
			Decoder: func(src *TodoJSON, dest *Todo) error {
				if src == nil {
					return nil
				}
				dest.Done = src.Done
				return nil
			},
		},
		CreatedAt: &TodoPropertyInfo{
			name: "CreatedAt",
			Encoder: func(src *Todo, dest *TodoJSON) error {
				if src == nil {
					return nil
				}
				dest.CreatedAt = src.CreatedAt
				return nil
			},
			Decoder: func(src *TodoJSON, dest *Todo) error {
				if src == nil {
					return nil
				}
				dest.CreatedAt = src.CreatedAt
				return nil
			},
		},
	}
}

// AddAll adds all property to TodoJSONBuilder.
func (b *TodoJSONBuilder) AddAll() *TodoJSONBuilder {
	b._properties["ID"] = b.ID
	b._properties["Text"] = b.Text
	b._properties["Done"] = b.Done
	b._properties["CreatedAt"] = b.CreatedAt
	return b
}

// Add specified property to TodoJSONBuilder.
func (b *TodoJSONBuilder) Add(info *TodoPropertyInfo) *TodoJSONBuilder {
	b._properties[info.name] = info
	return b
}

// Remove specified property to TodoJSONBuilder.
func (b *TodoJSONBuilder) Remove(info *TodoPropertyInfo) *TodoJSONBuilder {
	delete(b._properties, info.name)
	return b
}

// Convert specified non-JSON object to JSON object.
func (b *TodoJSONBuilder) Convert(orig *Todo) (*TodoJSON, error) {
	if orig == nil {
		return nil, nil
	}
	ret := &TodoJSON{}

	for _, info := range b._properties {
		if err := info.Encoder(orig, ret); err != nil {
			return nil, err
		}
	}

	return ret, nil
}

// ConvertList specified non-JSON slice to JSONList.
func (b *TodoJSONBuilder) ConvertList(orig []*Todo) (TodoJSONList, error) {
	if orig == nil {
		return nil, nil
	}

	list := make(TodoJSONList, len(orig))
	for idx, or := range orig {
		json, err := b.Convert(or)
		if err != nil {
			return nil, err
		}
		list[idx] = json
	}

	return list, nil
}

// Convert specified JSON object to non-JSON object.
func (orig *TodoJSON) Convert() (*Todo, error) {
	ret := &Todo{}

	b := NewTodoJSONBuilder().AddAll()
	for _, info := range b._properties {
		if err := info.Decoder(orig, ret); err != nil {
			return nil, err
		}
	}

	return ret, nil
}

// Convert specified JSONList to non-JSON slice.
func (jsonList TodoJSONList) Convert() ([]*Todo, error) {
	orig := ([]*TodoJSON)(jsonList)

	list := make([]*Todo, len(orig))
	for idx, or := range orig {
		obj, err := or.Convert()
		if err != nil {
			return nil, err
		}
		list[idx] = obj
	}

	return list, nil
}

// Marshal non-JSON object to JSON string.
func (b *TodoJSONBuilder) Marshal(orig *Todo) ([]byte, error) {
	ret, err := b.Convert(orig)
	if err != nil {
		return nil, err
	}
	return json.Marshal(ret)
}
