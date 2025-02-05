// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package student

import (
	"sync"
)

// Ensure, that StudentRepositoryInterfaceMock does implement StudentRepositoryInterface.
// If this is not the case, regenerate this file with moq.
var _ StudentRepositoryInterface = &StudentRepositoryInterfaceMock{}

// StudentRepositoryInterfaceMock is a mock implementation of StudentRepositoryInterface.
//
//	func TestSomethingThatUsesStudentRepositoryInterface(t *testing.T) {
//
//		// make and configure a mocked StudentRepositoryInterface
//		mockedStudentRepositoryInterface := &StudentRepositoryInterfaceMock{
//			GetAllStudentsFunc: func() ([]Student, error) {
//				panic("mock out the GetAllStudents method")
//			},
//		}
//
//		// use mockedStudentRepositoryInterface in code that requires StudentRepositoryInterface
//		// and then make assertions.
//
//	}
type StudentRepositoryInterfaceMock struct {
	// GetAllStudentsFunc mocks the GetAllStudents method.
	GetAllStudentsFunc func() ([]Student, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetAllStudents holds details about calls to the GetAllStudents method.
		GetAllStudents []struct {
		}
	}
	lockGetAllStudents sync.RWMutex
}

// GetAllStudents calls GetAllStudentsFunc.
func (mock *StudentRepositoryInterfaceMock) GetAllStudents() ([]Student, error) {
	if mock.GetAllStudentsFunc == nil {
		panic("StudentRepositoryInterfaceMock.GetAllStudentsFunc: method is nil but StudentRepositoryInterface.GetAllStudents was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetAllStudents.Lock()
	mock.calls.GetAllStudents = append(mock.calls.GetAllStudents, callInfo)
	mock.lockGetAllStudents.Unlock()
	return mock.GetAllStudentsFunc()
}

// GetAllStudentsCalls gets all the calls that were made to GetAllStudents.
// Check the length with:
//
//	len(mockedStudentRepositoryInterface.GetAllStudentsCalls())
func (mock *StudentRepositoryInterfaceMock) GetAllStudentsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetAllStudents.RLock()
	calls = mock.calls.GetAllStudents
	mock.lockGetAllStudents.RUnlock()
	return calls
}
