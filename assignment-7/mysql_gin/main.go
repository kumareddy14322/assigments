package main

//go get -u github.com/go-sql-driver/mysql
//go run RestApi.go
// cd GinGonic
import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@kumar(localhost:3306)/product_db")
	if err != nil {
		// fmt.Print(err.Error())
		fmt.Println("Error creating DB:", err)
		fmt.Println("To verify, db is:", db)
	}
	defer db.Close()
	fmt.Println("Successfully  Connected to MYSQl")
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}

	type Person struct {
		Id         int    `db:"ID" json:"id"`
		First_Name string `db:"first_name" json:"first_name"`
		Last_Name  string `db:"last_name" json:"last_name"`
		Address    string `db:"Address" json:"Address"`
	}

	router := gin.Default()

	// GET a person detail
	router.GET("/person/:id", func(c *gin.Context) {
		var (
			person Person
		)
		//strconv.Atoi("-42")
		// id := c.Query("id")
		// id1,err = strconv.Atoi(id)
		// id := c.Query("id")
		// id := c.PostForm("id")
		// id := c.Params.ByName("id")
		// id := c.PostForm("id")
		id := c.Param("id")

		rows, err := db.Query("select * from person where id = ?;", id)
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&person.Id, &person.First_Name, &person.Last_Name, &person.Address)
			// persons = append(persons, person)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": person,
			"count":  1,
		})
	})

	// GET all persons
	router.GET("/persons", func(c *gin.Context) {
		var (
			person  Person
			persons []Person
		)
		rows, err := db.Query("select * from person;")
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&person.Id, &person.First_Name, &person.Last_Name, &person.Address)
			persons = append(persons, person)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": persons,
			"count":  len(persons),
		})
	})

	// POST new person details
	router.POST("/person", func(c *gin.Context) {
		var buffer bytes.Buffer
		var person Person
		c.Bind(&person)
		// id, err := strconv.Atoi(c.PostForm("id"))
		// fmt.Println("hello", id)
		// //id := c.PostForm("id")
		// first_name := c.PostForm("first_name")
		// last_name := c.PostForm("last_name")
		// Address, err := strconv.Atoi(c.PostForm("Address"))
		id := person.Id

		//id := c.PostForm("id")
		first_name := person.First_Name
		last_name := person.Last_Name
		Address := person.Address
		//Address := c.PostForm("Address")
		stmt, err := db.Prepare("insert into person (id,first_name, last_name,Address) values(?,?,?,?);")
		if err != nil {
			fmt.Print(err.Error())
		}
		// _, err = stmt.Exec(&id, &first_name, &last_name, &Address)
		_, err = stmt.Exec(id, first_name, last_name, Address)

		if err != nil {
			fmt.Print(err.Error())
		}

		// Fastest way to append strings
		//buffer.WriteString(id)
		buffer.WriteString(" ")
		buffer.WriteString(first_name)
		buffer.WriteString(" ")
		buffer.WriteString(last_name)
		buffer.WriteString(" ")

		// buffer.WriteString(strconv.Itoa(Address))
		//buffer.WriteString(Address)
		defer stmt.Close()
		name := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"messAddress": fmt.Sprintf(" %s %ssuccessfully created", first_name, name),
		})
	})

	// PUT - update a person details
	router.PUT("/person/:id", func(c *gin.Context) {
		var buffer bytes.Buffer
		// id := c.Query("id")
		// first_name := c.Query("first_name")
		// last_name := c.Query("last_name")
		// Address := c.Query("Address")

		id := c.Param("id")
		var person Person
		c.Bind(&person)
		// id := person.Id

		//id := c.PostForm("id")
		first_name := person.First_Name
		last_name := person.Last_Name
		Address := person.Address
		stmt, err := db.Prepare("update person set first_name= ?, last_name= ?,Address=? where id= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(first_name, last_name, Address, id)
		if err != nil {
			fmt.Print(err.Error())
		}

		// Fastest way to append strings
		buffer.WriteString(first_name)
		buffer.WriteString(" ")
		buffer.WriteString(last_name)
		// buffer.WriteString(" ")
		// buffer.WriteString(Address)
		defer stmt.Close()
		name := buffer.String()

		c.JSON(http.StatusOK, gin.H{
			"messAddress": fmt.Sprintf("Successfully updated to %s", name),
		})
	})

	// Delete resources
	router.DELETE("/person/:id", func(c *gin.Context) {
		// id := c.Query("id")

		var person Person
		c.Bind(&person)
		// id := person.Id
		id := c.Param("id")
		stmt, err := db.Prepare("delete from person where id= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(id)
		if err != nil {
			fmt.Print(err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"messAddress": fmt.Sprintf("Successfully deleted user: %s", id),
		})
	})
	router.Run(":9000")
}

//go run gin_api.go