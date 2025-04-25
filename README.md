# Go_Test_food-delivery-backend

这是一个使用 **Go + Gin + MySQL + Redis** 构建的外卖平台后端服务，支持用户下单、商家上架菜品、订单管理等核心功能。

## 🌟 项目功能

- 用户注册与登录（JWT 鉴权）
- 商家注册与登录
- 菜品管理（上架、查看）
- 用户下单（通过菜品名称）
- 订单管理（修改、查看、取消）
- 热门菜品排行榜（使用 Redis 实现）

## 🛠 技术栈

- 编程语言：Golang
- Web框架：Gin
- 数据库：MySQL（使用 GORM 作为 ORM）
- 缓存：Redis（实现热门菜品排行）
- 工具：Postman（测试接口）、Git（版本控制）

## 📂 项目结构
```bash
.
├── controller         # 控制器：处理业务逻辑
├── middleware         # 中间件：鉴权
├── models             # 数据模型
├── routes             # 路由设置
├── utils              # 工具类：数据库连接等
├── main.go            # 启动入口
```
## 📦 接口说明（部分示例）

### 1. 用户注册
POST /register
{
  "username": "test",
  "password": "123456"
}

### 2. 用户下单（根据菜品名）
POST /orders/Place
Header: Authorization: Bearer <token>
{
  "dish_name": "香辣鸡腿堡",
  "quantity": 2
}

### 3. 查看订单
GET /orders/View?page=1
Header: Authorization: Bearer <token>

## ⚙️ 启动方式

克隆仓库：

git clone https://github.com/ryowaso/Go_Test_food-delivery-backend.git
cd Go_Test_food-delivery-backend

配置 .env 文件或在 utils/DB.go 中修改数据库连接：

dsn := "用户名:密码@tcp(127.0.0.1:3306)/数据库名?charset=utf8mb4&parseTime=True&loc=Local"

启动服务：

go run main.go

## 📋 TODO（可扩展功能）

添加商家菜品编辑、删除接口

支持订单评价、评分

增加订单分页、排序筛选

引入 Swagger 文档或前端页面联调

## 📄 License

MIT License

欢迎 star ⭐ 或 fork 🔁 本项目！如有问题请提 issue 🙌
