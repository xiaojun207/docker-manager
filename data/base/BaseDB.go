package base

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/xiaojun207/go-base-utils/utils"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"xorm.io/core"
	"xorm.io/xorm"
	"xorm.io/xorm/caches"
)

// https://www.kancloud.cn/kancloud/xorm-manual-zh-cn/56013
var (
	DBEngine   *xorm.Engine
	DriverName = ""
)

var (
	SUPPORT_DRIVER = []string{"mysql", "sqlite3"}
)

/*
eg.:
 Init("mysql", "root:123456@(127.0.0.1:3306)/test?charset=utf8")
or
 Init("sqlite3", "./test.db")
*/
func InitDB(driverName, dataSourceUrl string, useCache bool) *xorm.Engine {
	if !utils.ArrayContains(SUPPORT_DRIVER, driverName) {
		log.Println("driver \"" + driverName + "\" is unsupport!")
		panic("driver \"" + driverName + "\" is unsupport!")
	}
	DriverName = driverName
	if driverName == "sqlite3" {
		os.MkdirAll(filepath.Dir(dataSourceUrl), os.ModePerm)
	}

	// 创建一个引擎，对于xorm，一个引擎对应一个数据库
	// 但是创建引擎时不会检查连接有效性，只是解析为xorm数据结构
	// 当执行的操作需要用到连接时，xorm才会发现错误并返回错误
	var err error
	DBEngine, err = xorm.NewEngine(driverName, dataSourceUrl)
	if err != nil {
		log.Println("dataSourceUrl:", dataSourceUrl)
		panic(err)
	}
	err = DBEngine.Ping()
	if err != nil {
		log.Println("dataSourceUrl:", dataSourceUrl)
		panic(err)
	}
	log.Println(driverName + " 数据库配置成功")

	DBConfig(DBEngine)
	if useCache {
		DBEngine.SetMapper(core.GonicMapper{})
		// 启用内存缓存
		cacher := caches.NewLRUCacher(caches.NewMemoryStore(), 1000)
		DBEngine.SetDefaultCacher(cacher)
	}
	//GetDBVersion()
	return DBEngine
}

func IsSqlite3() bool {
	return DriverName == "sqlite3"
}

func GetDBVersion() {
	log.Println("DBEngine.DataSourceName():", DBEngine.DataSourceName())
	result, err := DBEngine.QueryInterface("select version();")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	if len(result) != 1 {
		panic("GetDBVersion, QueryOne Expectation 1, reality " + strconv.Itoa(len(result)))
	}

	log.Println("DB Version:", string(result[0]["version()"].([]byte)))
}

func CloseDBEngine() {
	DBEngine.Close()
}

// Config xorm一些可选设置
func DBConfig(engine *xorm.Engine) {
	// 设置日志等级，设置显示sql，设置显示执行时间
	//engine.SetLogLevel(xorm.DEFAULT_LOG_LEVEL)
	//engine.ShowSQL(true)
	//engine.ShowExecTime(true)

	// 指定结构体字段到数据库字段的转换器
	// 默认为core.SnakeMapper
	// 但是我们通常在struct中使用"ID"
	// 而SnakeMapper将"ID"转换为"i_d"
	// 因此我们需要手动指定转换器为core.GonicMapper{}
}

func Transaction(engine *xorm.Engine, f func(*xorm.Session) (interface{}, error)) (interface{}, error) {
	session := engine.NewSession()
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			session.Rollback()
		}
		session.Close()
	}()

	if err := session.Begin(); err != nil {
		return nil, err
	}

	result, err := f(session)
	if err != nil {
		session.Rollback()
		return nil, err
	}

	if err := session.Commit(); err != nil {
		//session.Rollback()
		return nil, err
	}

	return result, nil
}

func ExecSQL(engine *xorm.Engine, sql string, args ...interface{}) sql.Result {
	res, err := engine.Exec(sql)
	if err != nil {
		log.Printf("执行SQL失败：%v", err)
		panic(err)
	}
	affected, err := res.RowsAffected()
	log.Println("执行SQL成功:", affected, err)
	return res
}

func QuerySql(engine *xorm.Engine, sqlOrArgs ...interface{}) ([]map[string]interface{}, error) {
	return engine.QueryInterface(sqlOrArgs)
}

func QueryOne(engine *xorm.Engine, sqlOrArgs ...interface{}) (map[string]interface{}, error) {
	result, err := engine.QueryInterface(sqlOrArgs)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	if len(result) != 1 {
		panic("QueryOne Expectation 1, reality " + strconv.Itoa(len(result)))
	}
	return result[0], err
}

func QueryOneStr(engine *xorm.Engine, sqlOrArgs string) (map[string]string, error) {
	result, err := engine.QueryString(sqlOrArgs)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	if len(result) != 1 {
		panic("QueryOne Expectation 1, reality " + strconv.Itoa(len(result)))
	}
	return result[0], err
}

func ArrayParams(s []string) (sql string, params []interface{}) {
	sql = ""
	for i := 0; i < len(s); i++ {
		if i == 0 {
			sql += "?"
		} else {
			sql += "?,"
		}
		params = append(params, s[i])
	}
	return sql, params
}
