package models

import (
	"encoding/json"
	"strings"
)

type WsType int32

const (
	MESSAGE  WsType = iota
	CHAT     WsType = iota
	PRESENCE WsType = iota
	TYPING   WsType = iota
	UNKNOWN  WsType = iota
)

// String - Creating common behavior - give the type a String function
func (w WsType) String() string {
	switch w {
	case CHAT:
		return "CHAT"
	case MESSAGE:
		return "MESSAGE"
	case PRESENCE:
		return "PRESENCE"
	case TYPING:
		return "TYPING"
	}

	return "UNKNOWN"
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (w WsType) EnumIndex() int {
	return int(w)
}

func (w WsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(w.String())
}

// UnmarshalJSON must be a *pointer receiver* to ensure that the indirect from the
// parsed value can be set on the unmarshaling object. This means that the
// ParseType function must return a *value* and not a pointer.
func (w *WsType) UnmarshalJSON(data []byte) (err error) {
	var suits string
	if err := json.Unmarshal(data, &suits); err != nil {
		return err
	}

	if *w, err = ParseType(suits); err != nil {
		return err
	}

	return nil
}

func ParseType(s string) (WsType, error) {
	s = strings.TrimSpace(strings.ToLower(s))
	switch s {
	case "chat":
		return CHAT, nil
	case "message":
		return MESSAGE, nil
	case "presence":
		return PRESENCE, nil
	case "typing":
		return TYPING, nil

	}

	return UNKNOWN, nil
}
