# ATOM

## 框架介绍
### 技术栈
- 命令行工具 [cobra](https://github.com/spf13/cobra)
- 配置管理 [viper](https://github.com/spf13/viper)
- ORM [gorm](https://github.com/go-gorm/gorm)
- 文档 [swagger](https://github.com/swaggo/swag)
- 数据生成 [gofakeit](github.com/brianvoe/gofakeit)
- 框架增强 [gen](github.com/rogeecn/gen)
- 依赖注入 [dig](go.uber.org/dig)

### 命令行工具 atomctl

#### 安装 

    go install github.com/rogeecn/atomctl@latest

#### 工具使用
- gen 生成类功能
    - [crud](#gen:curd) 生成crud模板
    - [routes](#gen:routes) 生成业务路由
- new 创建类功能
    - [controller](#new:controller) 创建controller
    - [service](#new:service) 创建 service
    - [dao](#new:dao) 创建 dao 文件 
    - [migration](#new:migration) 创建 migration
    - [seeder](#new:seeder) 创建填充数据 seeder 
    - [http](#new:http) 创建 http 项目
    - [module](#new:module) 创建 http module
    - [suite](#new:suite) 创建测试用例文件

##### gen:curd
    atomctl gen crud [表名] [module]
例: 为 system 模块生成表 users 的 crud 操作
    atomctl gen crud users system
命令会分别生成文件:
```
module/
    system/
        controllers/
            users.go
        service/
            users.go
        dao/
            users.go
```
> 生成的文件不会被自动注册,需要手动添加新生成的方法至各目录的 provider 中.

#### gen:routes
1. 为整个项目生成 routes
    ```
    atomctl gen routes 
    ```
2. 为指定controller 生成 routes
    ```
    atomctl gen routes [controller_file]
    ```
> 生成的路由需要手动注册至`routes.go`中

#### new:controller
    atomctl new controller [module] [name]
#### new:service
    atomctl new service [module] [name]
#### new:dao
    atomctl new dao [module] [name]
#### new:migration
    atomctl new migration [migration_name]

创建文件路径 : `project_path/database/migrations/migration_name.go`
> migration 需要创建后手动在当前项目中执行 `go run . migrate` 才可生效
#### new:seeder
    atomctl new seeder [seeder_name]

创建文件路径 : `project_path/database/seeders/seeder_name.go`
> 注意：新添加的seeder需要手动在 `project_path/database/seeders/seeder.go` 中添加才可正常使用
使用方法：
    go run . seed

#### new:http
    atomctl new http [pkg] [project_name]
生成一个新的http类型项目

#### new:module
    atomctl new module [module_name]
为项目生成新的 module

#### new:suite
    atomctl new suite [filename]

## 快速开始
下面将会创建一个用户管理应用示例程序

1. 创建项目 atom-project
    ```
    atomctl new http atom/http atom-project
    cd atom-project
    go mod tidy
    ```
2. 添加  user module
    ```
    atomctl new module users

    # 生成目录 modules/users
    ```
3. 配置项目启动需要的 providers
    打开入口文件  `main.go`。下面我们需要4个基本的provider
    - sqlite 数据库
    - swagger api文档 
    - faker seeder数据生成
    - query dao需要
    - users 新建立的module
    - boot 应用数据初始化相关
    
    编辑后的 main 文件相关内容如下
    ```go
    import (
        "atom/http/database/migrations"
        "atom/http/database/query"
        "atom/http/database/seeders"
        "atom/http/modules/boot"
        "atom/http/modules/users"

        "github.com/rogeecn/atom"
        "github.com/rogeecn/atom-addons/providers/database/sqlite"
        "github.com/rogeecn/atom-addons/providers/faker"
        "github.com/rogeecn/atom-addons/providers/swagger"
        "github.com/rogeecn/atom-addons/services/http"
        "github.com/spf13/cobra"
    )
    // ...
	providers := http.Default(
		sqlite.DefaultProvider(),
		swagger.DefaultProvider(),
		faker.DefaultProvider(),
		query.DefaultProvider(),
	).With(
		boot.Providers(),
		users.Providers(),
	)
    // ...
    ```
4. 创建 boot module provider
    因为业务中需要用到业务内初始化内容，所以引入 boot provider
    编辑文件 `modules/boot/provider.go`
    ```
    package boot

    import (
        "atom/http/docs"

        "github.com/rogeecn/atom"
        "github.com/rogeecn/atom-addons/providers/swagger"
        "github.com/rogeecn/atom/container"
        "github.com/rogeecn/atom/contracts"
        "github.com/rogeecn/atom/utils/opt"
    )

    func Providers() container.Providers {
        return container.Providers{
            {Provider: provideSwagger},
        }
    }

    func provideSwagger(opts ...opt.Option) error {
        return container.Container.Provide(func(swagger *swagger.Swagger) contracts.Initial {
            swagger.Load(docs.SwaggerSpec)
            return nil
        }, atom.GroupInitial)
    }

    ```
5. 执行 `go mod tidy`
6. 添加 migration 
    ```
    atomctl new migration create_user
    ```
    编辑新建的文件，内容如下：
    ```
    func (m *xxxx) table() interface{} {
        type User struct {
            gorm.Model // 注意，新建表必须引入此字段，更新表不需要
            Username string
            Age      int
            Sex      string
        }

        return &User{}
    }
    ```
    配置文件默认使用sqlite引擎，其它引擎同理
8. 运行migration
    ```
    go run . migrate up
    ```
    直接执行，会运行失败，你会得到类似错误
    ```
    2023/06/15 17:50:54 load config file: : config file read error: Config File "http.toml" Not Found in "[/Users/rogee /Users/rogee/http /Users/rogee/.config /Users/rogee/.config/http /etc /etc/http /usr/local/etc /usr/local/etc/http]"
    ```
    把项目目录中的配置文件转移到项目目录中去
    ```
    ln -s  $PWD/config.toml ~/.config/http.toml  
    ```
    再次执行如下输出 
    ```
    2023/06/15 17:53:12 config file: /Users/rogee/.config/http.toml
    2023/06/15 17:53:12 BINGO! migrate up done
    ```
    此时会在项目根目录下看到 `sqlite.db` 文件, 打开文件可见包信2个表，migrations 和 users, migrations 是已执行migration的id有序记录
    ```
    sqlite3 sqlite.db
    SQLite version 3.39.4 2022-09-07 20:51:41
    Enter ".help" for usage hints.
    sqlite> .tables
    migrations  users     
    ```
9. 生成model
    go run . model

    model 会在新表添加后再次执行生成，所以不要编辑任何 `database/models`下的文件内容 

9. 添加假数据
    atomctl new seeder User 
    > 注意User是单数，表示 model.User 区别于数据库表名

    编辑文件 `database/seeders/users.go`

    ```
    func (s *UsersSeeder) Generate(faker *gofakeit.Faker, idx int) models.User {
        return models.User{
            Username: faker.Name(),
            Age:      int32(faker.Number(18, 60)),
            Sex:      faker.RandomString([]string{"male", "female"}),
        }
    }
    ```
    编辑文件 `database/seeders/seeder.go`, 把新添加的seeder注册到执行列表中，添加的seeder会被有序执行
    ```
    var Seeders = []contracts.SeederProvider{
        NewUsersSeeder,
    }
    ```
    运行 `go run . seed `, 些时 users 表会被写入10条随机数据
10. 添加 crud
    atomctl gen crud users users
    两个users第一个为crud的名称,第二个为模块名称，执行输出如下
    ```
    2023/06/15 18:06:42 generate:  modules/users/service/users.go
    2023/06/15 18:06:42 generate:  modules/users/dto/users.go
    2023/06/15 18:06:42 generate:  modules/users/dao/users.go
    2023/06/15 18:06:42 generate:  modules/users/controller/users.go
    2023/06/15 18:06:42 generate crud success
    2023/06/15 18:06:42 REMEMBER TO ADD NEW PROVIDERS
    ```
11. 生成 api
    atomctl gen routes
    输出如下
    ```
    2023/06/15 18:07:43 route path:  /Users/rogee/tmp/atom-project
    2023/06/15 18:07:43 go mod package:  atom/http
    2023/06/15 18:07:43 generate routes for dir
    2023/06/15 18:07:43 generate route: /Users/rogee/tmp/atom-project/modules/users/routes/route_user_controller.go @ routeUserController(group, userController)
    ```
    把文件后的提示方法写到 `modules/users/routes/routes.go`, 些时新建的 route provider 还没有注入 userController ，我们手动添加进去，编辑后的文件内容如下:
    ```
    func newRoute(svc http.Service, userController *controller.UserController) http.Route {
        engine := svc.GetEngine().(*gin.Engine)
        group := engine.Group("users")
        log.Info("register route group: %s", group)

        routeUserController(group, userController)

        return nil
    }

    ```
13. 注册服务 
    ```
    modules/users/
            controller/ provider.go
            dao/provider.go
            service/provider.go
        provider.go
    ```
    注册crud provider对象
    内容类似
    ```
    func Provide(opts ...opt.Option) error {
        _ = container.Container.Provide(NewUserController)
        return nil
    }   
    ```
12. 运行项目
    ```
    go generate ./...
    go run .
    ```
13. 访问
    1. [localhost:9800/doc/index.html](localhost:9800/doc/index.html)