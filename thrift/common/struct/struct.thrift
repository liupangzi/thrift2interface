namespace go root_common.struct.struct

include "../../base.thrift"

struct XStruct {
    2: required base.HTTPCode code,
    3: optional string msg,
    7: required map<i32, i8> xMap,
    11: optional list<i16> xList,
}

struct YStruct {
    13: optional map<i64, list<XStruct>> yMap,
    17: required list<base.HTTPCode> yList,
    19: optional map<string, base.Json> jsonMap,
}
