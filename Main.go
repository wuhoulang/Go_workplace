package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
)

type Ret1 struct {
	Code  int
	Param string
	Msg   string
}

var (
	dbhostsip  = "127.0.0.1:3306" //IP地址
	dbusername = "root"           //用户名
	dbpassword = "123456"         //密码
	dbname     = "shot"           //表名
	id         int
	name1      string
	name       string
	sex        string
	home       string
	password   string
	password1   string
	account   string
	extra      string
	db         *sql.DB
)

func main() {

	db, _ = sql.Open("mysql", "root:123456@tcp(192.168.0.44:3306)/shot?charset=utf8")
	defer db.Close() //关闭数据库
	err := db.Ping() //连接数据库
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}


	go func(){
		mux := http.NewServeMux()
		mux.HandleFunc("/", HelloServer1)
		http.ListenAndServe(":9092", mux)
	}()
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloServer2)
	http.ListenAndServe(":9093", mux)

	//sqlselect(db)
	//insert(db)
	////delete(db)
}

func HelloServer1(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name1 = r.FormValue("name")
	password = r.FormValue("password")
	fmt.Println("name", name1)
	fmt.Println("password", password)
	if name1 != "" {
		insert(db, w, name1, password)
		//ret.Code = 0
		//ret.Msg = "success"
		//ret_json,_ := json.Marshal(ret)
		//io.WriteString(w, string(ret_json))
	} else {
		//ret.Code = 1
		//ret.Msg = "error"
		//ret_json,_ := json.Marshal(ret)
		//io.WriteString(w, string(ret_json))
	}
}

func HelloServer2(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name1 = r.FormValue("name")
	password1 = r.FormValue("password")
	if name1 != "" {
		sqlselect(db, w, name1, password1)
	}
}

func sqlselect(db *sql.DB,w http.ResponseWriter,name1 string,password1 string) {
	var ret_json1 []byte
	var ret2 Ret1
	fmt.Println("name1",name1)
	fmt.Println("password1",password1)
	rows, err := db.Query("SELECT * FROM login where account = ? and password = ? ",name1 ,password1)
	// rows相当于执行迭代器，到末尾后自动关闭
	if rows ==nil {
		return
	}
	if  err !=nil {
		fmt.Println("error",err)
	}
	for rows.Next() {

		// 把数据读取出到相应的变量中
		err = rows.Scan(&id, &account, &password, )
		if err != nil {
			fmt.Println("查询出错了")
		}
		fmt.Println(id)
		fmt.Println("account",account)
		fmt.Println("password",password)
	}

	if account==name1 && password==password1 {
		ret2 =  Ret1{}
		ret2.Code = 0
		ret2.Msg = "success"
		ret_json1, _ = json.Marshal(ret2)
		io.WriteString(w, string(ret_json1))
	}else {
		ret2 =  Ret1{}
		ret2.Code = 1
		ret2.Msg = "fail"
		ret_json1, _ = json.Marshal(ret2)
		io.WriteString(w, string(ret_json1))
	}

}

func insert(db *sql.DB, w http.ResponseWriter, name1 string, password1 string) {

	var err2 error
	var ret1 Ret1
	var ret_json []byte

	result, err2 := db.Exec("INSERT INTO login (account, password) VALUES (?, ?)",name1,password1)
	id, err1 := result.LastInsertId()
	if err1 != nil {
		return
	}
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("result",result)
	fmt.Println("id",id)
	//fmt.Println(id)
	if result != nil {
		ret1 =  Ret1{}
		ret1.Code = 0
		ret1.Msg = "success"
		ret_json, _ = json.Marshal(ret1)
		io.WriteString(w, string(ret_json))
	}
}

func delete(db *sql.DB) {
	//准备sql语句
	stmt, err := db.Prepare("DELETE FROM price WHERE id = ?")
	if err != nil {
		fmt.Println("Prepare fail")
		return
	}
	//设置参数以及执行sql语句
	res, err := stmt.Exec(5)
	if err != nil {
		fmt.Println("Exec fail")
	}
	fmt.Println(res)
}

func UpdateUser(db *sql.DB) {
	//准备sql语句
	stmt, err := db.Prepare("UPDATE price SET name = ?, password = ? WHERE id = ?")
	if err != nil {
		fmt.Println("Prepare fail")
	}
	//设置参数以及执行sql语句
	res, err := stmt.Exec(name, password, id)
	if err != nil {
		fmt.Println("Exec fail")
	}
	//提交事务
	fmt.Println(res.LastInsertId())

}
