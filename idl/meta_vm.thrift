namespace cpp meta_rpc
namespace go meta_rpc

struct VMInfoReq {
    1: i64 LastUpdateTimestamp,
}

struct VMInfoResp {
    1: binary VMInfo,
    2: i64 UpdateTimestamp,
}

struct VersionInfo{
    1: string version_name,
    2: string params,
}

struct VersionResp{
    1: required VersionInfo info,
    2: required string msg,
}

struct VersionReq{
    1: required string req,
}

service MetaVMService {
    VMInfoResp GetVMInfo(1: VMInfoReq req),
    VersionResp GetAbversion(1: VersionReq req),
}
