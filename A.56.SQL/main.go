package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type Student struct {
	Id    string
	Name  string
	Age   int
	Grade int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./dev.db")

	if err != nil {
		return nil, err
	}

	return db, nil
}

func init() {
	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	schema := `
		CREATE TABLE IF NOT EXISTS tb_student (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		age INTEGER NOT NULL,
		grade INTEGER NOT NULL
		);

		INSERT INTO tb_student (id, name, age, grade) VALUES
		('B001', 'Jason Bourne', 29, 1),
		('B002', 'James Bond', 27, 1),
		('E001', 'Ethan Hunt', 27, 2),
		('W001', 'John Wick', 28, 2)
		ON CONFLICT(id) DO NOTHING;
	`
	_, err = db.Exec(schema)
	if err != nil {
		log.Fatalf("Failed to init database: %v", err)
	}
}

func getAllStudent(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM tb_student")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var result []Student
	for rows.Next() {
		var each = Student{}
		var err = rows.Scan(&each.Id, &each.Name, &each.Age, &each.Grade)

		if err != nil {
			log.Fatal(err)
		}

		result = append(result, each)
	}

	if rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, each := range result {
		fmt.Println(each.Id, each.Name, each.Age, each.Grade)
	}
}

func getStudent(db *sql.DB, id string) {
	var result = Student{}

	err := db.
		QueryRow("SELECT * FROM tb_student WHERE id = ?", id).
		Scan(&result.Id, &result.Name, &result.Age, &result.Grade)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("name: %s\ngrade: %d\n", result.Name, result.Grade)
}

func prepareQuery(db *sql.DB) {
	stmt, err := db.Prepare("SELECT name from tb_student WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	var result string
	stmt.QueryRow("E001").Scan(&result)

	fmt.Println(result)
}

func insertNewStudent(db *sql.DB) {
	var newStudent = Student{
		Id:    "A010",
		Name:  "Indra",
		Age:   22,
		Grade: 2,
	}

	result, err := db.Exec(`
		INSERT INTO tb_student 
		VALUES (?, ?, ?, ?) 
		ON CONFLICT(id) DO NOTHING`, newStudent.Id, newStudent.Name, newStudent.Age, newStudent.Grade)

	if err != nil {
		log.Fatal(err)
	}

	row, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Row RowsAffected %v\n", row)
}

func updateNameStudent(db *sql.DB, id string) {
	result, err := db.Exec(`
		UPDATE tb_student
		SET name = ?
		WHERE id = ?
	`, "updated name", id)

	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Row updated %d\n", rowsAffected)
}

func deleteStudent(db *sql.DB, id string) {
	result, err := db.Exec(`
		DELETE FROM tb_student
		WHERE id = ?
	`, id)
	if err != nil {
		log.Fatal()
	}

	deleted, err := result.RowsAffected()
	if err != nil {
		log.Fatal()
	}

	if deleted > 0 {
		fmt.Println("DELETE SUCCESS")
	}
}

func main() {
	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	getAllStudent(db)
	getStudent(db, "B002")
	prepareQuery(db)
	insertNewStudent(db)
	updateNameStudent(db, "W001")
	deleteStudent(db, "E001")
}
