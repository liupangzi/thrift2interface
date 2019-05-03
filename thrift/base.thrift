namespace go thrift.test
namespace java thrift.test
namespace js ThriftTest
namespace * thrift.test

include 'common/struct/struct.thrift'
include "common/exception/exception.thrift"

typedef string Json

enum HTTPCode {
    OK = 200,
    NOT_FOUND = 404,
    INTERNAL_SERVER_ERROR = 500,
}

const HTTPCode CodeOK = HTTPCode.OK;

service BaseOne {
    void ServiceOne() throws (1: exception.pException e1);

    void ServiceTwo(1: struct.XStruct p2);

    struct.YStruct ServiceThree();

    struct.YStruct ServiceFive(1: struct.XStruct p5) throws (1: exception.pException e5);

    oneway void ServiceSeven() throws (1: exception.pException e7);

    oneway void ServiceEleven(1: struct.XStruct p13);
}
