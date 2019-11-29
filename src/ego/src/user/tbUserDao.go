package user

import (
	"ego/src/commons"
	"fmt"
)

//根据用户名和密码查询
func SelByUnPwdDao(un,pwd string) *TbUser {
	sql :="select * from tb_user where username =? and password=?  or email =? and  password=?"
	rows,err:=commons.Dql(sql,un,pwd,un,pwd)
	//fmt.Println(rows)
	if err !=nil{
		fmt.Println(err)
		return nil
	}
	if rows.Next(){
		user:=new(TbUser)
		rows.Scan(&user.Id,&user.Username,&user.Password,&user.Phone,&user.Email,&user.Created,&user.Updated)

		commons.CloseConn()
		return user
	}
	return nil
}
