# RS-backend

基于 Go + Gin 的抽油机监控后端。系统通过 Modbus TCP 轮询 PLC 数据，提供 REST API、WebSocket 和静态页面服务。

## 当前已实现能力

- Modbus TCP 客户端连接 PLC
- 以固定间隔轮询 11 个保持寄存器
- 内存中维护最近一次实时数据
- 将 `Position`、`Load` 写入 SQLite
- 提供实时数据 REST API
- 提供历史单点查询 REST API
- 提供实时数据 WebSocket 推送
- 托管 `static/` 下前端页面

## 项目结构

```text
.
├── internal/
│   ├── database/      # SQLite 初始化与历史数据存储
│   ├── handlers/      # HTTP / WebSocket 入口
│   ├── modbus/        # Modbus 客户端与轮询逻辑
│   ├── models/        # 请求 / 响应模型
│   ├── routes/        # 路由注册
│   └── services/      # 业务层
├── static/            # 前端静态资源
├── main.go            # 程序入口
└── README.md
```

## 已实现 API

服务默认监听 `http://localhost:8080`。

### 1. `GET /api/realtime`

返回最近一次 Modbus 轮询得到的实时数据。

示例响应：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "timestamp": 1772356354,
    "realtime": "2026-03-01 17:12:34",
    "position": 981,
    "load": 445,
    "motorSpeed": 1500,
    "strokesNumber": 5,
    "distance": 3000,
    "time": "2026-03-01 17:12:33",
    "rodDensity": 7850,
    "transmissionRatio": 2000,
    "area": 500,
    "inclination": 30,
    "pumpInsertionDepth": 1200,
    "oilDensity": 850
  }
}
```

说明：
- `timestamp` 和 `realtime` 是接口返回时的当前时间。
- `time` 是最近一次轮询数据写入内存时的时间。

### 2. `GET /api/history?timestamp=<unix>`

按 Unix 秒级时间戳查询历史点位，查询范围是该秒内的第一条记录。

请求示例：

```http
GET /api/history?timestamp=1760871111
```

成功响应：

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "id": 3960,
    "timestamp": 1772356945,
    "time": "2026-03-01 17:22:25",
    "position": 1129,
    "load": 812,
    "inclination": 0,
    "motorSpeed": 0,
    "oilDensity": 0,
    "pumpInsertionDepth": 0,
    "rodDensity": 0,
    "strokesNumber": 0,
    "transmissionRatio": 0,
    "area": 0,
    "distance": 0
  }
}
```

错误响应：

```json
{
  "code": 1,
  "msg": "缺少时间戳参数"
}
```

```json
{
  "code": 1,
  "msg": "时间戳格式错误，应为Unix时间戳格式（例如：1760869745）"
}
```

```json
{
  "code": 1,
  "msg": "未找到指定时间的数据"
}
```

说明：
- 当前数据库只落库了 `Position` 和 `Load`。
- 其余字段会出现在响应结构里，但默认是 `0`。

### 3. `POST /api/user`

当前实现为“设备连接信息提交接口”，会校验 JSON 格式并返回一组 mock 设备参数。

请求体：

```json
{
  "id": "device-001",
  "ip": "127.0.0.1",
  "port": "5020"
}
```

成功响应：

```json
{
  "status": 200,
  "data": {
    "motorSpeed": "1200",
    "strokesNumber": "5",
    "distance": "3.5",
    "rodDensity": "7850",
    "transmissionRatio": "50",
    "area": "0.001",
    "inclination": "30",
    "pumpInsertionDepth": "1000",
    "oilDensity": "850"
  }
}
```

失败响应：

```json
{
  "status": 400,
  "error": "Invalid connection info"
}
```

说明：
- 该接口目前没有真正修改 Modbus 连接地址。
- 返回参数来自业务层中的 mock 数据，不会影响 `internal/modbus/client.go` 当前运行中的客户端。

### 4. `GET /ws`

WebSocket 实时推送最近一次轮询数据，推送周期为 `500ms`。

连接地址：

```text
ws://localhost:8080/ws
```

消息示例：

```json
{
  "Time": "2026-03-01 17:12:33",
  "Position": 981,
  "Load": 445,
  "MotorSpeed": 1500,
  "StrokesNumber": 5,
  "Distance": 3000,
  "RodDensity": 7850,
  "TransmissionRatio": 2000,
  "Area": 500,
  "Inclination": 30,
  "PumpInsertionDepth": 1200,
  "OilDensity": 850
}
```

## 静态页面入口

- `GET /`：返回 `static/index.html`
- `GET /static/*filepath`：托管 `static/` 目录下全部静态资源

## Modbus 采集说明

当前 [client.go](/home/way/GolandProjects/RS-backend/internal/modbus/client.go) 的实现行为：

- 默认连接地址：`127.0.0.1:5020`
- 连接协议：`tcp://`
- `Unit ID` 固定为 `1`
- 轮询周期：`500ms`
- 读取寄存器范围：从地址 `0` 开始连续读取 `11` 个保持寄存器

寄存器映射：

1. `0` -> `Position`
2. `1` -> `Load`
3. `2` -> `MotorSpeed`
4. `3` -> `StrokesNumber`
5. `4` -> `Distance`
6. `5` -> `RodDensity`
7. `6` -> `TransmissionRatio`
8. `7` -> `Area`
9. `8` -> `Inclination`
10. `9` -> `PumpInsertionDepth`
11. `10` -> `OilDensity`

异常处理：

- 启动时首次连接失败不会退出，会记录日志并继续重试
- 轮询失败时会关闭当前连接，等待 `2s` 后重连
- 成功重连后继续轮询

## 运行方式

环境要求：

- Go 1.16+
- 可访问的 Modbus TCP 设备或模拟器

启动：

```bash
go build -o rs-backend main.go
./rs-backend
```

## 当前限制

- `POST /api/user` 还是 mock，没有驱动真实设备连接
- 历史数据只保存了 `Position` 和 `Load`
- Modbus 地址写死在 [main.go](/home/way/GolandProjects/RS-backend/main.go)
- 缺少配置文件、鉴权、结构化日志和测试

## 测试建议

如需本地联调，可使用 fake Modbus 服务，例如：

- [fake-modbus-server](https://github.com/WAYYYAW/fake-modbus-server)
