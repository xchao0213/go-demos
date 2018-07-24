package model


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

// 	// name := "chenyichao"
// 	// birthday := "1981-02-13"
// 	// sex := "m"
// 	// identity_card := "310226198102132017"
// 	// mobile := "18221659878"
	
// 	name1 := "panjia"
// 	birthday1 := "1981-02-06"
// 	sex1 := "m"
// 	identity_card1 := "310226198102062021"
//     mobile1 := "13482844330"

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