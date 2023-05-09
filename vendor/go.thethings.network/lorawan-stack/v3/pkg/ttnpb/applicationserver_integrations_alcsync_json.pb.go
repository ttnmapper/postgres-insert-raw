// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.4.2
// - protoc             v3.21.1
// source: lorawan-stack/api/applicationserver_integrations_alcsync.proto

package ttnpb

import (
	golang "github.com/TheThingsIndustries/protoc-gen-go-json/golang"
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
)

// MarshalProtoJSON marshals the ALCSyncCommandIdentifier to JSON.
func (x ALCSyncCommandIdentifier) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	s.WriteEnumString(int32(x), ALCSyncCommandIdentifier_name)
}

// MarshalText marshals the ALCSyncCommandIdentifier to text.
func (x ALCSyncCommandIdentifier) MarshalText() ([]byte, error) {
	return []byte(jsonplugin.GetEnumString(int32(x), ALCSyncCommandIdentifier_name)), nil
}

// MarshalJSON marshals the ALCSyncCommandIdentifier to JSON.
func (x ALCSyncCommandIdentifier) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// ALCSyncCommandIdentifier_customvalue contains custom string values that extend ALCSyncCommandIdentifier_value.
var ALCSyncCommandIdentifier_customvalue = map[string]int32{
	"PKG_VERSION":              0,
	"APP_TIME":                 1,
	"APP_DEV_TIME_PERIODICITY": 2,
	"FORCE_DEV_RESYNC":         3,
}

// UnmarshalProtoJSON unmarshals the ALCSyncCommandIdentifier from JSON.
func (x *ALCSyncCommandIdentifier) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	v := s.ReadEnum(ALCSyncCommandIdentifier_value, ALCSyncCommandIdentifier_customvalue)
	if err := s.Err(); err != nil {
		s.SetErrorf("could not read ALCSyncCommandIdentifier enum: %v", err)
		return
	}
	*x = ALCSyncCommandIdentifier(v)
}

// UnmarshalText unmarshals the ALCSyncCommandIdentifier from text.
func (x *ALCSyncCommandIdentifier) UnmarshalText(b []byte) error {
	i, err := jsonplugin.ParseEnumString(string(b), ALCSyncCommandIdentifier_customvalue, ALCSyncCommandIdentifier_value)
	if err != nil {
		return err
	}
	*x = ALCSyncCommandIdentifier(i)
	return nil
}

// UnmarshalJSON unmarshals the ALCSyncCommandIdentifier from JSON.
func (x *ALCSyncCommandIdentifier) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the ALCSyncCommand message to JSON.
func (x *ALCSyncCommand) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Cid != 0 || s.HasField("cid") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("cid")
		x.Cid.MarshalProtoJSON(s)
	}
	if x.Payload != nil {
		switch ov := x.Payload.(type) {
		case *ALCSyncCommand_AppTimeReq_:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("app_time_req")
			// NOTE: ALCSyncCommand_AppTimeReq does not seem to implement MarshalProtoJSON.
			golang.MarshalMessage(s, ov.AppTimeReq)
		case *ALCSyncCommand_AppTimeAns_:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("app_time_ans")
			// NOTE: ALCSyncCommand_AppTimeAns does not seem to implement MarshalProtoJSON.
			golang.MarshalMessage(s, ov.AppTimeAns)
		}
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the ALCSyncCommand to JSON.
func (x *ALCSyncCommand) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the ALCSyncCommand message from JSON.
func (x *ALCSyncCommand) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "cid":
			s.AddField("cid")
			x.Cid.UnmarshalProtoJSON(s)
		case "app_time_req", "appTimeReq":
			s.AddField("app_time_req")
			ov := &ALCSyncCommand_AppTimeReq_{}
			x.Payload = ov
			if s.ReadNil() {
				ov.AppTimeReq = nil
				return
			}
			// NOTE: ALCSyncCommand_AppTimeReq does not seem to implement UnmarshalProtoJSON.
			var v ALCSyncCommand_AppTimeReq
			golang.UnmarshalMessage(s, &v)
			ov.AppTimeReq = &v
		case "app_time_ans", "appTimeAns":
			s.AddField("app_time_ans")
			ov := &ALCSyncCommand_AppTimeAns_{}
			x.Payload = ov
			if s.ReadNil() {
				ov.AppTimeAns = nil
				return
			}
			// NOTE: ALCSyncCommand_AppTimeAns does not seem to implement UnmarshalProtoJSON.
			var v ALCSyncCommand_AppTimeAns
			golang.UnmarshalMessage(s, &v)
			ov.AppTimeAns = &v
		}
	})
}

// UnmarshalJSON unmarshals the ALCSyncCommand from JSON.
func (x *ALCSyncCommand) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}