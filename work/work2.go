package work

//第二周作业 错误信息上抛，业务层来打印原始错误和
import (
	"database/sql"
	"fmt"
	"github.com/beego/beego/v2/client/httplib"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

const (
	userName = "root"
	password = "123456"
	ip = "120.24.43.187"
	port = "3306"
	dbName = "onlinemall_bn"
)

//Db数据库连接池
var DB *sql.DB
func InitDB()  {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil{
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("opon database success")
}


type user struct {
	id int
	user_name string
	password string
	nike_name string
}

func NewUser() *user {
	return &user{}
}

//查询一行
//查询到错误，封装错误信息向上抛出去由业务层来处理
func QueryOneRow(u *user) (err error) {
	httplib.Post()
	str_sql:=fmt.Sprintf("se lect id,user_name,password,nike_name from user where id=%d",1)
	row := DB.QueryRow(str_sql)
	if err = row.Scan(&u.id, &u.user_name, &u.password,&u.nike_name);err!=nil {
		err = fmt.Errorf("查询一行出错,sql=[%s] %w",str_sql,err)
	}
	return
}

func init()  {
	InitDB()
}



