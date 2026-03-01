# RS-backend 抽油机监控系统后端

这是一个用于监控抽油机运行状态的后端系统，通过Modbus协议读取设备数据并通过WebSocket实时推送。

## 功能特性

- 通过Modbus TCP协议连接PLC设备
- 实时读取抽油机的Position（位置）和Load（载荷）数据
- 提供REST API接口获取实时数据
- 通过WebSocket推送实时数据到前端
- 静态文件服务，可托管前端页面

## 项目结构

```
.
├── internal/
│   ├── modbus/          # Modbus客户端实现
│   ├── handlers/        # HTTP请求处理器
│   └── routes/          # 路由配置
├── static/              # 静态文件（前端页面）
├── main.go             # 程序入口
└── README.md           # 项目说明文档
```

## API接口

### REST API

- `GET /api/realtime` - 获取当前实时数据（完整设备参数）

返回示例：
```json
{
  "code": 0,
  "data": {
    "area": 500.0,
    "distance": 3000.0,
    "inclination": 3000.0,
    "load": 445.0,
    "motorSpeed": 1500.0,
    "oilDensity": 850.0,
    "position": 981.0,
    "pumpInsertionDepth": 1200.0,
    "realtime": "2026-03-01 17:12:34",
    "rodDensity": 7850.0,
    "strokesNumber": 5.0,
    "time": "2026-03-01 17:12:33",
    "timestamp": 1772356354,
    "transmissionRatio": 2000.0
  },
  "msg": "success"
}
```
- `GET /api/history` - 获取历史数据（Position和Load）

请求示例：
```http
GET /api/history?timestamp=1760871111
```

返回示例：
```json
{
  "code": 0,
  "data": {
    "area": 0,
    "distance": 0,
    "id": 3960,
    "inclination": 0,
    "load": 812,
    "motorSpeed": 0,
    "oilDensity": 0,
    "position": 1129,
    "pumpInsertionDepth": 0,
    "rodDensity": 0,
    "strokesNumber": 0,
    "time": "2026-03-01 17:22:25",
    "timestamp": 1772356945,
    "transmissionRatio": 0
  },
  "msg": "success"
}
```

### WebSocket

- `GET /ws` - WebSocket连接端点，实时推送Position和Load数据

推送数据格式：
```json
{
  "Time": "2026-03-01 17:12:33",
  "Position": 981.0,
  "Load": 445.0,
  "MotorSpeed": 1500.0,
  "StrokesNumber": 5.0,
  "Distance": 3000.0,
  "RodDensity": 7850.0,
  "TransmissionRatio": 2000.0,
  "Area": 500.0,
  "Inclination": 3000.0,
  "PumpInsertionDepth": 1200.0,
  "OilDensity": 850.0
}
```

## 静态页面

- `GET /` - 主页面（重定向到index.html）__（尚未完成）__
- `GET /static/test.html` - 测试页面，显示实时数据

## 环境要求

- Go 1.16+
- PLC设备或Modbus模拟器

## 配置

默认连接地址：`127.0.0.1:5020`

如需修改连接地址，请在 [main.go](file:///home/way/GolandProjects/RS-backend/main.go) 中修改：


```
client := modbus.NewClient("your-plc-address:port")
```
## 构建和运行

```bash
# 构建
go build -o rs-backend main.go

# 运行
./rs-backend
```
## TODO
- [ ] 添加配置文件读取
- [ ] 添加日志记录
- [x] 将数据存储到数据库
- [x] 添加数据可视化
- [x] 修复大小写命名规范问题
## 测试模式

如果需要使用模拟数据进行测试，可以使用[fake-modbus-server](https://github.com/WAYYYAW/fake-modbus-server)项目启动一个模拟的Modbus服务器。

## 前端页面

- test.html - 简化的实时数据显示页面，显示所有设备参数
- index.html - 主页面（基础版本）
- home2.html - 完整功能的主页面，包含设备配置和历史数据查询
- home.html - 备用主页

## 日志

程序运行时会输出读取到的数据日志：
```
读取到数据: Position=1256.000000, Load=2406.000000
```

如果无法连接到PLC设备，程序会直接终止并输出错误日志。