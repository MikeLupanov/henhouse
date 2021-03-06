/**
 * @file category_test.go
 * @author Mikhail Klementyev jollheef<AT>riseup.net
 * @license GNU AGPLv3
 * @date November, 2015
 * @brief test work with category table
 */

package db

import (
	"errors"
	"fmt"
	"testing"
)

func TestCreateCategoryTable(*testing.T) {

	db, err := InitDatabase(dbPath)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = createCategoryTable(db)
	if err != nil {
		panic(err)
	}
}

func TestAddCategory(*testing.T) {

	db, err := InitDatabase(dbPath)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	category := Category{ID: 255, Name: "test"}

	err = AddCategory(db, &category)
	if err != nil {
		panic(err)
	}

	if category.ID != 1 {
		panic(errors.New("Category id not correct"))
	}
}

// Test add category with closed database
func TestFailAddCategory(*testing.T) {

	db, err := InitDatabase(dbPath)
	if err != nil {
		panic(err)
	}

	db.Close()

	err = AddCategory(db, &Category{})
	if err == nil {
		panic(err)
	}
}

func TestGetCategories(*testing.T) {

	db, err := InitDatabase(dbPath)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	categories := 150

	for i := 0; i < categories; i++ {

		err = AddCategory(db, &Category{Name: fmt.Sprintf("%d", i)})
		if err != nil {
			panic(err)
		}
	}

	cats, err := GetCategories(db)
	if err != nil {
		panic(err)
	}

	if len(cats) != categories {
		panic(errors.New("Mismatch get categories length"))
	}

	for i := 0; i < categories; i++ {

		if cats[i].Name != fmt.Sprintf("%d", i) && cats[i].ID != i {
			panic(errors.New("Get invalid category"))
		}
	}
}

// Test get categories with closed database
func TestFailGetCategories(*testing.T) {

	db, err := InitDatabase(dbPath)
	if err != nil {
		panic(err)
	}

	db.Close()

	_, err = GetCategories(db)
	if err == nil {
		panic(err)
	}
}
