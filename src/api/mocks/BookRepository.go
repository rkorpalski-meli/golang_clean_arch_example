// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/rkorpalski-meli/golang_clean_arch_example/src/api/entity"
	mock "github.com/stretchr/testify/mock"
)

// BookRepository is an autogenerated mock type for the BookRepository type
type BookRepository struct {
	mock.Mock
}

// Get provides a mock function with given fields: title
func (_m *BookRepository) Get(title string) (entity.Book, error) {
	ret := _m.Called(title)

	var r0 entity.Book
	if rf, ok := ret.Get(0).(func(string) entity.Book); ok {
		r0 = rf(title)
	} else {
		r0 = ret.Get(0).(entity.Book)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: book
func (_m *BookRepository) Insert(book entity.Book) error {
	ret := _m.Called(book)

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.Book) error); ok {
		r0 = rf(book)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
