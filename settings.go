package gani

import (
	"fmt"
	"strings"
)

// Settings defines the Gani file settings
type Settings struct {
	Looping         bool
	Continuous      bool
	SingleDirection bool
	SetBackTo       string
	DefaultParams   [3]string
	DefaultAttrs    [3]string
	DefaultHead     string
	DefaultBody     string
}

func (s *Settings) String() string {
	builder := strings.Builder{}

	if s.Looping {
		builder.WriteString("LOOP")
		builder.WriteString("\n")
	}

	if s.Continuous {
		builder.WriteString("CONTINUOUS")
		builder.WriteString("\n")
	}

	if s.SingleDirection {
		builder.WriteString("SINGLEDIRECTION")
		builder.WriteString("\n")
	}

	if !strings.EqualFold(s.SetBackTo, "") {
		builder.WriteString(fmt.Sprintf("SETBACKTO %s", s.SetBackTo))
		builder.WriteString("\n")
	}

	for i, param := range s.DefaultParams {
		if !strings.EqualFold(param, "") {
			builder.WriteString(fmt.Sprintf("DEFAULTPARAM%d %s", i+1, param))
			builder.WriteString("\n")
		}
	}

	for i, attr := range s.DefaultAttrs {
		if !strings.EqualFold(attr, "") {
			builder.WriteString(fmt.Sprintf("DEFAULTATTR%d %s", i+1, attr))
			builder.WriteString("\n")
		}
	}

	if !strings.EqualFold(s.DefaultHead, "") {
		builder.WriteString(fmt.Sprintf("DEFAULTHEAD %s", s.DefaultHead))
		builder.WriteString("\n")
	}

	if !strings.EqualFold(s.DefaultBody, "") {
		builder.WriteString(fmt.Sprintf("DEFAULTBODY %s", s.DefaultBody))
		builder.WriteString("\n")
	}

	return builder.String()
}

// Parse attempts to parse the given line to a Settings value
func (s *Settings) Parse(line string) error {
	strs := strings.Split(line, " ")

	if len(strs) == 1 {
		return s.parseBool(strs[0])
	}

	return s.parseKeyValuePair(strs[0], strs[1])
}

// parseBool attempts to determine a boolean Settings value from the given string
func (s *Settings) parseBool(str string) error {
	if strings.EqualFold(str, "LOOP") {
		s.Looping = true
		return nil
	}

	if strings.EqualFold(str, "CONTINUOUS") {
		s.Continuous = true
		return nil
	}

	if strings.EqualFold(str, "SINGLEDIRECTION") {
		s.SingleDirection = true
		return nil
	}

	return fmt.Errorf("unknown boolean field: %s", str)
}

// parseKeyValuePair attempts to determine a key value pair Settings value for the given
// key and value
func (s *Settings) parseKeyValuePair(key, value string) error {
	switch key {
	case "SETBACKTO":
		s.SetBackTo = value
	case "DEFAULTPARAM1":
		s.DefaultParams[0] = value
	case "DEFAULTPARAM2":
		s.DefaultParams[1] = value
	case "DEFAULTPARAM3":
		s.DefaultParams[2] = value
	case "DEFAULTATTR1":
		s.DefaultAttrs[0] = value
	case "DEFAULTATTR2":
		s.DefaultAttrs[1] = value
	case "DEFAULTATTR3":
		s.DefaultAttrs[2] = value
	case "DEFAULTHEAD":
		s.DefaultHead = value
	case "DEFAULTBODY":
		s.DefaultBody = value
	default:
		return fmt.Errorf("unknown kvp")
	}

	return nil
}

// NewSettings returns the default settings object found when creating a new Gani
func NewSettings() *Settings {
	return &Settings{
		Looping:         false,
		Continuous:      false,
		SingleDirection: false,
		SetBackTo:       "",
		DefaultParams:   [3]string{},
		DefaultAttrs:    [3]string{"hat0.png"},
		DefaultHead:     "head19.png",
		DefaultBody:     "body.png",
	}
}
