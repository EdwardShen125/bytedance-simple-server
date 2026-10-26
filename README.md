# bytedance-simple-server

A short-video app backend — feed, publish, user, comment, favorite, relation, message — built in Go as a microservice set: a Gin HTTP gateway in front of seven Kitex RPC services defined by Thrift IDL.

## Layout

| Path | What it is |
|---|---|
| `cmd/api` | Gin HTTP gateway: handlers, JWT auth, Kitex clients for downstream calls |
| `cmd/{feed,user,publish,comment,favorite,relation,message}` | Seven Kitex RPC services |
| `idl/` | Thrift IDL per service — the contract |
| `kitex_gen/` | Kitex-generated stubs |
| `pkg/` | Shared infrastructure (below) |
| `app/` | Companion Android client (release APK) |
| `script/` | `bootstrap.sh` to bring the set up |

## Shared packages (`pkg/`)

- `jwt` — token issue and verification (gin-jwt)
- `middleware` — cross-cutting handler chain
- `errno` — unified error codes across services
- `oss` — object-storage client for media
- `tracer` — OpenTracing/Jaeger wiring, including GORM spans

## Runtime dependencies

etcd (service registration and discovery via registry-etcd) · MySQL (GORM) · Jaeger · object storage.

## Regenerating stubs

```
kitex -module github.com/EdwardShen125/bytedance-simple-server -service comment idl/comment.thrift
kitex -module github.com/EdwardShen125/bytedance-simple-server -service favorite idl/favorite.thrift
kitex -module github.com/EdwardShen125/bytedance-simple-server -service feed idl/feed.thrift
kitex -module github.com/EdwardShen125/bytedance-simple-server -service message idl/message.thrift
kitex -module github.com/EdwardShen125/bytedance-simple-server -service publish idl/publish.thrift
kitex -module github.com/EdwardShen125/bytedance-simple-server -service relation idl/relation.thrift
kitex -module github.com/EdwardShen125/bytedance-simple-server -service user idl/user.thrift
```
