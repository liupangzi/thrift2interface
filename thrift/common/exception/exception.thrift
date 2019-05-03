namespace go root.common_exception.exception

include '../../base.thrift'

exception pException {
    41: optional base.HTTPCode code,
    43: required base.Json message,
}

exception qException {
    53: optional base.HTTPCode code,
    59: required base.Json message,
    61: optional set<base.Json> qSet,
}
