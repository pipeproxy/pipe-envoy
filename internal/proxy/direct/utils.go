package direct

import (
	"fmt"

	"google.golang.org/protobuf/types/known/structpb"
)

func MapToProtoStruct(m map[string]interface{}) (*structpb.Struct, error) {
	fields := map[string]*structpb.Value{}
	for k, v := range m {
		val, err := ValueToStructValue(v)
		if err != nil {
			return nil, err
		}
		fields[k] = val
	}
	return &structpb.Struct{Fields: fields}, nil
}

func ValueToStructValue(v interface{}) (*structpb.Value, error) {
	switch x := v.(type) {
	case nil:
		return &structpb.Value{Kind: &structpb.Value_NullValue{}}, nil
	case bool:
		return &structpb.Value{Kind: &structpb.Value_BoolValue{BoolValue: x}}, nil
	case float64:
		return &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: x}}, nil
	case float32:
		return &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: float64(x)}}, nil
	case int:
		return &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: float64(x)}}, nil
	case int8:
		return &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: float64(x)}}, nil
	case int16:
		return &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: float64(x)}}, nil
	case int32:
		return &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: float64(x)}}, nil
	case int64:
		return &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: float64(x)}}, nil
	case uint:
		return &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: float64(x)}}, nil
	case uint8:
		return &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: float64(x)}}, nil
	case uint16:
		return &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: float64(x)}}, nil
	case uint32:
		return &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: float64(x)}}, nil
	case uint64:
		return &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: float64(x)}}, nil
	case string:
		return &structpb.Value{Kind: &structpb.Value_StringValue{StringValue: x}}, nil
	case map[string]interface{}:
		m, err := MapToProtoStruct(x)
		if err != nil {
			return nil, err
		}
		return &structpb.Value{Kind: &structpb.Value_StructValue{StructValue: m}}, nil
	case []interface{}:
		var vals []*structpb.Value
		for _, e := range x {
			val, err := ValueToStructValue(e)
			if err != nil {
				return nil, err
			}
			vals = append(vals, val)
		}
		return &structpb.Value{Kind: &structpb.Value_ListValue{ListValue: &structpb.ListValue{Values: vals}}}, nil
	default:
		return nil, fmt.Errorf("bad type %T for JSON value", v)
	}
}
