package storage

import (
	"encoding/csv"
	"errors"
	"os"
	"strings"
)

func AuthenticateUser(username, password string) (string, error) {
	// Assume this matches against hardcoded or file-stored user credentials.
	// Returning "admin" or "regular" as userType.
	if username == "admin" && password == "admin123" {
		return "admin", nil
	} else if username == "user" && password == "user123" {
		return "regular", nil
	}
	return "", errors.New("user not found")
}

type User struct {
	Username string
	Password string
	UserType string
}

type Books struct {
	BookName        string `json:"bookName"`
	Author          string `json:"author"`
	PublicationYear string `json:"publicationYear"`
}

func GetBooks(userType string) ([]Books, error) {
	var (
		books []Books
		path  string
	)

	if userType == "admin" {
		path = "/Users/shubhamprajapati/Documents/p-shubh/scalex.club/internal/data/adminUser.csv"
	} else {
		path = "/Users/shubhamprajapati/Documents/p-shubh/scalex.club/internal/data/regularUser.csv"
	}

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Skip header row
	records = records[1:]

	for _, record := range records {
		book := Books{
			BookName:        record[0],
			Author:          record[1],
			PublicationYear: record[2],
		}
		books = append(books, book)
	}

	if userType == "admin" {
		path = "/Users/shubhamprajapati/Documents/p-shubh/scalex.club/internal/data/regularUser.csv"
		file, err := os.Open(path)

		if err != nil {
			return nil, err
		}
		defer file.Close()

		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			return nil, err
		}

		// Skip header row
		records = records[1:]

		for _, record := range records {
			book := Books{
				BookName:        record[0],
				Author:          record[1],
				PublicationYear: record[2],
			}
			books = append(books, book)
		}
	}

	return books, nil
}

// AddBook and DeleteBook to be implemented according to CRUD operations on CSV files.
const (
	regularUserCSV = "/Users/shubhamprajapati/Documents/p-shubh/scalex.club/internal/data/regularUser.csv"
	adminUserCSV   = "/Users/shubhamprajapati/Documents/p-shubh/scalex.club/internal/data/adminUser.csv"
)

func AddBook(userType, bookName, author, publicationYear string) error {
	var csvFile string = regularUserCSV

	file, err := os.OpenFile(csvFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{bookName, author, publicationYear})
	if err != nil {
		return err
	}

	return nil
}

func DeleteBook(userType, bookName string) error {
	var csvFile string = regularUserCSV

	tempFile, err := os.CreateTemp("", "temp.csv")
	if err != nil {
		return err
	}
	defer tempFile.Close()

	file, err := os.Open(csvFile)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	writer := csv.NewWriter(tempFile)

	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return err
		}

		if strings.EqualFold(record[0], bookName) {
			continue
		}

		err = writer.Write(record)
		if err != nil {
			return err
		}
	}

	writer.Flush()

	err = os.Remove(csvFile)
	if err != nil {
		return err
	}

	err = os.Rename(tempFile.Name(), csvFile)
	if err != nil {
		return err
	}

	return nil
}
