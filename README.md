# 功能
1. 读取指定的配置文件
2. 读取配置文件到内存变量(map)中
3. 读取系统环境变量
  ` 主要用来读取环境变量或者配置文件 `

# 使用方法
 go get github.com/bugfan/goini
## 读取指定的配置文件 (此方式每次从文件读取)
  ` 支持以[空格][等于号][冒号]分隔key与value `
1. goini.LoadConfig("./test.conf")
2. goini.Config.GetString("key")
3. goini.Config.GetInt64("key")
## 读取配置文件到自定义变量 (此方式只读取一次配置文件并放到变量)
1. myMap, err := goini.ReadFile(".env")   //读取文件配置到自定义map
2. log.Println("Key's Value is:",myMap[key])  //从自定义map获取value
## 读取单个或者多个配置文件到内存变量 (此方式支持读取多个配置文件,支持读取系统环境变量，并把所有读到的配置放到goini.Env对象，并返回)
1. goini.Env:=goini.NewMyEnv(".env","../path/env","myenv.txt") //读取配置文件到goini.Env
2. goini.Env:=goini.NewMyEnv()    //不指定配置文件则从系统环境变量(程序执行的进程空间)读取
3. goini.Env.Getenv(key)  //获取key的vlaue 返回string类型
4. goini.Env.Getenvd(key,def) //获取不到则使用def作为值 返回string类型
5. goini.Env.GetAllenv()  //获取所有环境变量,返回一个map

# 使用须知
  ` 可以参照go test 文件来使用 `
  ` 配置文件支持以[空格][等于号][冒号]分隔key与value ，支持[#号]注释配置文件`
