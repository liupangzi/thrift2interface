namespace go service_twitter.test

include "../common/struct/struct.thrift"
include '../base.thrift'

struct TStruct {
    2: required base.HTTPCode code,
    3: optional string msg,
    7: required map<i32, i8> xMap,
    11: optional list<i16> xList,
}

service Twitter {
    void TwitterOne();

    void TwitterTwo(1: struct.XStruct p2);

    TStruct TwitterThree();

    struct.YStruct TwitterFive(1: TStruct p5);

    oneway void TwitterSeven();

    oneway void TwitterEleven(1: struct.XStruct p11);
}
