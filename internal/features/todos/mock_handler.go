// Code generated by mockery v2.34.2. DO NOT EDIT.

package todos

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// MockHandler is an autogenerated mock type for the Handler type
type MockHandler struct {
	mock.Mock
}

type MockHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockHandler) EXPECT() *MockHandler_Expecter {
	return &MockHandler_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: w, r
func (_m *MockHandler) Create(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// MockHandler_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockHandler_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *MockHandler_Expecter) Create(w interface{}, r interface{}) *MockHandler_Create_Call {
	return &MockHandler_Create_Call{Call: _e.mock.On("Create", w, r)}
}

func (_c *MockHandler_Create_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *MockHandler_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *MockHandler_Create_Call) Return() *MockHandler_Create_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockHandler_Create_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *MockHandler_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: w, r
func (_m *MockHandler) Delete(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// MockHandler_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockHandler_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *MockHandler_Expecter) Delete(w interface{}, r interface{}) *MockHandler_Delete_Call {
	return &MockHandler_Delete_Call{Call: _e.mock.On("Delete", w, r)}
}

func (_c *MockHandler_Delete_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *MockHandler_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *MockHandler_Delete_Call) Return() *MockHandler_Delete_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockHandler_Delete_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *MockHandler_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: w, r
func (_m *MockHandler) Get(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// MockHandler_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockHandler_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *MockHandler_Expecter) Get(w interface{}, r interface{}) *MockHandler_Get_Call {
	return &MockHandler_Get_Call{Call: _e.mock.On("Get", w, r)}
}

func (_c *MockHandler_Get_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *MockHandler_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *MockHandler_Get_Call) Return() *MockHandler_Get_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockHandler_Get_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *MockHandler_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Search provides a mock function with given fields: w, r
func (_m *MockHandler) Search(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// MockHandler_Search_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Search'
type MockHandler_Search_Call struct {
	*mock.Call
}

// Search is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *MockHandler_Expecter) Search(w interface{}, r interface{}) *MockHandler_Search_Call {
	return &MockHandler_Search_Call{Call: _e.mock.On("Search", w, r)}
}

func (_c *MockHandler_Search_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *MockHandler_Search_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *MockHandler_Search_Call) Return() *MockHandler_Search_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockHandler_Search_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *MockHandler_Search_Call {
	_c.Call.Return(run)
	return _c
}

// Sort provides a mock function with given fields: w, r
func (_m *MockHandler) Sort(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// MockHandler_Sort_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Sort'
type MockHandler_Sort_Call struct {
	*mock.Call
}

// Sort is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *MockHandler_Expecter) Sort(w interface{}, r interface{}) *MockHandler_Sort_Call {
	return &MockHandler_Sort_Call{Call: _e.mock.On("Sort", w, r)}
}

func (_c *MockHandler_Sort_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *MockHandler_Sort_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *MockHandler_Sort_Call) Return() *MockHandler_Sort_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockHandler_Sort_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *MockHandler_Sort_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: w, r
func (_m *MockHandler) Update(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// MockHandler_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockHandler_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *MockHandler_Expecter) Update(w interface{}, r interface{}) *MockHandler_Update_Call {
	return &MockHandler_Update_Call{Call: _e.mock.On("Update", w, r)}
}

func (_c *MockHandler_Update_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *MockHandler_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *MockHandler_Update_Call) Return() *MockHandler_Update_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockHandler_Update_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *MockHandler_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockHandler creates a new instance of MockHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockHandler {
	mock := &MockHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
