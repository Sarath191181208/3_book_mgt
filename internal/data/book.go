package data

import "fmt"

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

type Book struct {
	Id     json.Number `json:"id"`
	Name   string      `json:"name"`
	Author string      `json:"author"`
}

type BookModel struct {
	conn *sql.DB
}

func (db *BookModel) Insert(book *Book) error {
	query := `
    INSERT INTO 
      Books (name, author)
    VALUES (?, ?);
  `
	args := []interface{}{book.Name, book.Author}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	err := db.conn.QueryRowContext(ctx, query, args...).Err()
	if err != nil {
		return err
	}

	return nil
}

func (db *BookModel) GetById(id_json json.Number) (*Book, error) {
	id, err := id_json.Int64()
	if err != nil {
		return nil, err
	}

	book := &Book{}
	query := `SELECT id, name, author FROM Books WHERE id = ? `
	args := []interface{}{id}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = db.conn.QueryRowContext(ctx, query, args...).Scan(&id, &book.Name, &book.Author)
	book.Id = json.Number(fmt.Sprint(id))

	if err != nil {
		return nil, err
	}
	return book, nil
}

func (db *BookModel) GetAll() ([]*Book, error) {
	query := `SELECT id, name, author FROM Books LIMIT 10;`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := db.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []*Book
	for rows.Next() {
		book := &Book{}
		var id int
		err := rows.Scan(&id, &book.Name, &book.Author)
		book.Id = json.Number(fmt.Sprint(id))
		if err != nil {
			return nil, err
		}

		data = append(data, book)
	}

	return data, nil
}

func (db *BookModel) Delete(id_number json.Number) error {
	id, err := id_number.Int64()
	if err != nil {
		return err
	}

	query := `DELETE FROM Books WHERE id = ?;`
	args := []interface{}{id}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err = db.conn.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (db *BookModel) Update(book Book) error {
	query := `UPDATE Books 
    SET name = ? , author = ?
    WHERE id = ?;
  `
	id, err := book.Id.Int64()
	if err != nil {
		return err
	}
	args := []interface{}{book.Name, book.Author, id}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err = db.conn.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
