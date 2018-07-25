package model

import "fmt"


const tblUsers = "users"

type User struct {
	Id           string `db:"id"`
	Name         string `db:"name"`
	Birthday     string `db:"birthday"`
	Sex          string `db:"sex"`
	IdentifyCard string `db:"identify_card"`
	Mobile       string `db:"mobile"`
	Photo        string `db:"photo"`
	CreatedAt    int64  `db:"created_at"`
	UpdatedAt    int64  `db:"updated_at"`
}


func UserInsert(u *User) (int64,error){
  query, err := DB.Prepare(fmt.Sprintf("INSERT INTO %s%s (`id`,`name`,`birthday`,`sex`,`identify_card`,`mobile`,`photo`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?,?,?)",prefix,tblUsers))
  if err != nil {
	  return 0, err
  }

  ret, err := query.Exec(u.Id,u.Name,u.Birthday,u.Sex,u.IdentifyCard,u.Mobile,u.Photo,u.CreatedAt,u.UpdatedAt)
  if err != nil {
	  return 0, err
  }

  return ret.RowsAffected()
}

// func DoInsert() {
// 	db,err := sql.Open("mysql","root:123456@tcp(localhost:3306)/home?charset=latin1")

// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	defer db.Close()

// 	inser_sql := "INSERT INTO `user` VALUES (?,?,?,?,?)"

// 	stmt,err := db.Prepare(inser_sql)
 
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	res,err := stmt.Exec(name1,birthday1,sex1,identity_card1,mobile1)

// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	fmt.Println(res.RowsAffected)
// }

// func DoQuery() {
// 	db,err := sql.Open("mysql","root:123456@tcp(localhost:3306)/home?charset=latin1")

// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	defer db.Close()

// 	rows,err := db.Query("SELECT * FROM user")

// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	for rows.Next() {
// 		var name string
// 		var birthday string
// 		var sex string
// 		var identity_card string
// 		var mobile string

// 		err = rows.Scan(&name,&birthday,&sex,&identity_card,&mobile)

// 		if err != nil {
// 			panic(err.Error())
// 		}

// 		fmt.Println("name",name,"birthday",birthday,"sex",sex,"identity_card",identity_card,"mobile",mobile)
// 	}
// }