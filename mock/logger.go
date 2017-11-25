// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package mock

import (
	"github.com/go-mixins/log"
	"sync"
)

var (
	lockContextLoggerMockDebug       sync.RWMutex
	lockContextLoggerMockDebugf      sync.RWMutex
	lockContextLoggerMockError       sync.RWMutex
	lockContextLoggerMockErrorf      sync.RWMutex
	lockContextLoggerMockFatal       sync.RWMutex
	lockContextLoggerMockFatalf      sync.RWMutex
	lockContextLoggerMockInfo        sync.RWMutex
	lockContextLoggerMockInfof       sync.RWMutex
	lockContextLoggerMockPanic       sync.RWMutex
	lockContextLoggerMockPanicf      sync.RWMutex
	lockContextLoggerMockWarn        sync.RWMutex
	lockContextLoggerMockWarnf       sync.RWMutex
	lockContextLoggerMockWithContext sync.RWMutex
)

// ContextLoggerMock is a mock implementation of ContextLogger.
//
//     func TestSomethingThatUsesContextLogger(t *testing.T) {
//
//         // make and configure a mocked ContextLogger
//         mockedContextLogger := &ContextLoggerMock{
//             DebugFunc: func(in1 ...interface{})  {
// 	               panic("TODO: mock out the Debug method")
//             },
//             DebugfFunc: func(in1 string, in2 ...interface{})  {
// 	               panic("TODO: mock out the Debugf method")
//             },
//             ErrorFunc: func(in1 ...interface{})  {
// 	               panic("TODO: mock out the Error method")
//             },
//             ErrorfFunc: func(in1 string, in2 ...interface{})  {
// 	               panic("TODO: mock out the Errorf method")
//             },
//             FatalFunc: func(in1 ...interface{})  {
// 	               panic("TODO: mock out the Fatal method")
//             },
//             FatalfFunc: func(in1 string, in2 ...interface{})  {
// 	               panic("TODO: mock out the Fatalf method")
//             },
//             InfoFunc: func(in1 ...interface{})  {
// 	               panic("TODO: mock out the Info method")
//             },
//             InfofFunc: func(in1 string, in2 ...interface{})  {
// 	               panic("TODO: mock out the Infof method")
//             },
//             PanicFunc: func(in1 ...interface{})  {
// 	               panic("TODO: mock out the Panic method")
//             },
//             PanicfFunc: func(in1 string, in2 ...interface{})  {
// 	               panic("TODO: mock out the Panicf method")
//             },
//             WarnFunc: func(in1 ...interface{})  {
// 	               panic("TODO: mock out the Warn method")
//             },
//             WarnfFunc: func(in1 string, in2 ...interface{})  {
// 	               panic("TODO: mock out the Warnf method")
//             },
//             WithContextFunc: func(in1 log.M) log.ContextLogger {
// 	               panic("TODO: mock out the WithContext method")
//             },
//         }
//
//         // TODO: use mockedContextLogger in code that requires ContextLogger
//         //       and then make assertions.
//
//     }
type ContextLoggerMock struct {
	// DebugFunc mocks the Debug method.
	DebugFunc func(in1 ...interface{})

	// DebugfFunc mocks the Debugf method.
	DebugfFunc func(in1 string, in2 ...interface{})

	// ErrorFunc mocks the Error method.
	ErrorFunc func(in1 ...interface{})

	// ErrorfFunc mocks the Errorf method.
	ErrorfFunc func(in1 string, in2 ...interface{})

	// FatalFunc mocks the Fatal method.
	FatalFunc func(in1 ...interface{})

	// FatalfFunc mocks the Fatalf method.
	FatalfFunc func(in1 string, in2 ...interface{})

	// InfoFunc mocks the Info method.
	InfoFunc func(in1 ...interface{})

	// InfofFunc mocks the Infof method.
	InfofFunc func(in1 string, in2 ...interface{})

	// PanicFunc mocks the Panic method.
	PanicFunc func(in1 ...interface{})

	// PanicfFunc mocks the Panicf method.
	PanicfFunc func(in1 string, in2 ...interface{})

	// WarnFunc mocks the Warn method.
	WarnFunc func(in1 ...interface{})

	// WarnfFunc mocks the Warnf method.
	WarnfFunc func(in1 string, in2 ...interface{})

	// WithContextFunc mocks the WithContext method.
	WithContextFunc func(in1 log.M) log.ContextLogger

	// calls tracks calls to the methods.
	calls struct {
		// Debug holds details about calls to the Debug method.
		Debug []struct {
			// In1 is the in1 argument value.
			In1 []interface{}
		}
		// Debugf holds details about calls to the Debugf method.
		Debugf []struct {
			// In1 is the in1 argument value.
			In1 string
			// In2 is the in2 argument value.
			In2 []interface{}
		}
		// Error holds details about calls to the Error method.
		Error []struct {
			// In1 is the in1 argument value.
			In1 []interface{}
		}
		// Errorf holds details about calls to the Errorf method.
		Errorf []struct {
			// In1 is the in1 argument value.
			In1 string
			// In2 is the in2 argument value.
			In2 []interface{}
		}
		// Fatal holds details about calls to the Fatal method.
		Fatal []struct {
			// In1 is the in1 argument value.
			In1 []interface{}
		}
		// Fatalf holds details about calls to the Fatalf method.
		Fatalf []struct {
			// In1 is the in1 argument value.
			In1 string
			// In2 is the in2 argument value.
			In2 []interface{}
		}
		// Info holds details about calls to the Info method.
		Info []struct {
			// In1 is the in1 argument value.
			In1 []interface{}
		}
		// Infof holds details about calls to the Infof method.
		Infof []struct {
			// In1 is the in1 argument value.
			In1 string
			// In2 is the in2 argument value.
			In2 []interface{}
		}
		// Panic holds details about calls to the Panic method.
		Panic []struct {
			// In1 is the in1 argument value.
			In1 []interface{}
		}
		// Panicf holds details about calls to the Panicf method.
		Panicf []struct {
			// In1 is the in1 argument value.
			In1 string
			// In2 is the in2 argument value.
			In2 []interface{}
		}
		// Warn holds details about calls to the Warn method.
		Warn []struct {
			// In1 is the in1 argument value.
			In1 []interface{}
		}
		// Warnf holds details about calls to the Warnf method.
		Warnf []struct {
			// In1 is the in1 argument value.
			In1 string
			// In2 is the in2 argument value.
			In2 []interface{}
		}
		// WithContext holds details about calls to the WithContext method.
		WithContext []struct {
			// In1 is the in1 argument value.
			In1 log.M
		}
	}
}

// Debug calls DebugFunc.
func (mock *ContextLoggerMock) Debug(in1 ...interface{}) {
	if mock.DebugFunc == nil {
		panic("moq: ContextLoggerMock.DebugFunc is nil but ContextLogger.Debug was just called")
	}
	callInfo := struct {
		In1 []interface{}
	}{
		In1: in1,
	}
	lockContextLoggerMockDebug.Lock()
	mock.calls.Debug = append(mock.calls.Debug, callInfo)
	lockContextLoggerMockDebug.Unlock()
	mock.DebugFunc(in1...)
}

// DebugCalls gets all the calls that were made to Debug.
// Check the length with:
//     len(mockedContextLogger.DebugCalls())
func (mock *ContextLoggerMock) DebugCalls() []struct {
	In1 []interface{}
} {
	var calls []struct {
		In1 []interface{}
	}
	lockContextLoggerMockDebug.RLock()
	calls = mock.calls.Debug
	lockContextLoggerMockDebug.RUnlock()
	return calls
}

// Debugf calls DebugfFunc.
func (mock *ContextLoggerMock) Debugf(in1 string, in2 ...interface{}) {
	if mock.DebugfFunc == nil {
		panic("moq: ContextLoggerMock.DebugfFunc is nil but ContextLogger.Debugf was just called")
	}
	callInfo := struct {
		In1 string
		In2 []interface{}
	}{
		In1: in1,
		In2: in2,
	}
	lockContextLoggerMockDebugf.Lock()
	mock.calls.Debugf = append(mock.calls.Debugf, callInfo)
	lockContextLoggerMockDebugf.Unlock()
	mock.DebugfFunc(in1, in2...)
}

// DebugfCalls gets all the calls that were made to Debugf.
// Check the length with:
//     len(mockedContextLogger.DebugfCalls())
func (mock *ContextLoggerMock) DebugfCalls() []struct {
	In1 string
	In2 []interface{}
} {
	var calls []struct {
		In1 string
		In2 []interface{}
	}
	lockContextLoggerMockDebugf.RLock()
	calls = mock.calls.Debugf
	lockContextLoggerMockDebugf.RUnlock()
	return calls
}

// Error calls ErrorFunc.
func (mock *ContextLoggerMock) Error(in1 ...interface{}) {
	if mock.ErrorFunc == nil {
		panic("moq: ContextLoggerMock.ErrorFunc is nil but ContextLogger.Error was just called")
	}
	callInfo := struct {
		In1 []interface{}
	}{
		In1: in1,
	}
	lockContextLoggerMockError.Lock()
	mock.calls.Error = append(mock.calls.Error, callInfo)
	lockContextLoggerMockError.Unlock()
	mock.ErrorFunc(in1...)
}

// ErrorCalls gets all the calls that were made to Error.
// Check the length with:
//     len(mockedContextLogger.ErrorCalls())
func (mock *ContextLoggerMock) ErrorCalls() []struct {
	In1 []interface{}
} {
	var calls []struct {
		In1 []interface{}
	}
	lockContextLoggerMockError.RLock()
	calls = mock.calls.Error
	lockContextLoggerMockError.RUnlock()
	return calls
}

// Errorf calls ErrorfFunc.
func (mock *ContextLoggerMock) Errorf(in1 string, in2 ...interface{}) {
	if mock.ErrorfFunc == nil {
		panic("moq: ContextLoggerMock.ErrorfFunc is nil but ContextLogger.Errorf was just called")
	}
	callInfo := struct {
		In1 string
		In2 []interface{}
	}{
		In1: in1,
		In2: in2,
	}
	lockContextLoggerMockErrorf.Lock()
	mock.calls.Errorf = append(mock.calls.Errorf, callInfo)
	lockContextLoggerMockErrorf.Unlock()
	mock.ErrorfFunc(in1, in2...)
}

// ErrorfCalls gets all the calls that were made to Errorf.
// Check the length with:
//     len(mockedContextLogger.ErrorfCalls())
func (mock *ContextLoggerMock) ErrorfCalls() []struct {
	In1 string
	In2 []interface{}
} {
	var calls []struct {
		In1 string
		In2 []interface{}
	}
	lockContextLoggerMockErrorf.RLock()
	calls = mock.calls.Errorf
	lockContextLoggerMockErrorf.RUnlock()
	return calls
}

// Fatal calls FatalFunc.
func (mock *ContextLoggerMock) Fatal(in1 ...interface{}) {
	if mock.FatalFunc == nil {
		panic("moq: ContextLoggerMock.FatalFunc is nil but ContextLogger.Fatal was just called")
	}
	callInfo := struct {
		In1 []interface{}
	}{
		In1: in1,
	}
	lockContextLoggerMockFatal.Lock()
	mock.calls.Fatal = append(mock.calls.Fatal, callInfo)
	lockContextLoggerMockFatal.Unlock()
	mock.FatalFunc(in1...)
}

// FatalCalls gets all the calls that were made to Fatal.
// Check the length with:
//     len(mockedContextLogger.FatalCalls())
func (mock *ContextLoggerMock) FatalCalls() []struct {
	In1 []interface{}
} {
	var calls []struct {
		In1 []interface{}
	}
	lockContextLoggerMockFatal.RLock()
	calls = mock.calls.Fatal
	lockContextLoggerMockFatal.RUnlock()
	return calls
}

// Fatalf calls FatalfFunc.
func (mock *ContextLoggerMock) Fatalf(in1 string, in2 ...interface{}) {
	if mock.FatalfFunc == nil {
		panic("moq: ContextLoggerMock.FatalfFunc is nil but ContextLogger.Fatalf was just called")
	}
	callInfo := struct {
		In1 string
		In2 []interface{}
	}{
		In1: in1,
		In2: in2,
	}
	lockContextLoggerMockFatalf.Lock()
	mock.calls.Fatalf = append(mock.calls.Fatalf, callInfo)
	lockContextLoggerMockFatalf.Unlock()
	mock.FatalfFunc(in1, in2...)
}

// FatalfCalls gets all the calls that were made to Fatalf.
// Check the length with:
//     len(mockedContextLogger.FatalfCalls())
func (mock *ContextLoggerMock) FatalfCalls() []struct {
	In1 string
	In2 []interface{}
} {
	var calls []struct {
		In1 string
		In2 []interface{}
	}
	lockContextLoggerMockFatalf.RLock()
	calls = mock.calls.Fatalf
	lockContextLoggerMockFatalf.RUnlock()
	return calls
}

// Info calls InfoFunc.
func (mock *ContextLoggerMock) Info(in1 ...interface{}) {
	if mock.InfoFunc == nil {
		panic("moq: ContextLoggerMock.InfoFunc is nil but ContextLogger.Info was just called")
	}
	callInfo := struct {
		In1 []interface{}
	}{
		In1: in1,
	}
	lockContextLoggerMockInfo.Lock()
	mock.calls.Info = append(mock.calls.Info, callInfo)
	lockContextLoggerMockInfo.Unlock()
	mock.InfoFunc(in1...)
}

// InfoCalls gets all the calls that were made to Info.
// Check the length with:
//     len(mockedContextLogger.InfoCalls())
func (mock *ContextLoggerMock) InfoCalls() []struct {
	In1 []interface{}
} {
	var calls []struct {
		In1 []interface{}
	}
	lockContextLoggerMockInfo.RLock()
	calls = mock.calls.Info
	lockContextLoggerMockInfo.RUnlock()
	return calls
}

// Infof calls InfofFunc.
func (mock *ContextLoggerMock) Infof(in1 string, in2 ...interface{}) {
	if mock.InfofFunc == nil {
		panic("moq: ContextLoggerMock.InfofFunc is nil but ContextLogger.Infof was just called")
	}
	callInfo := struct {
		In1 string
		In2 []interface{}
	}{
		In1: in1,
		In2: in2,
	}
	lockContextLoggerMockInfof.Lock()
	mock.calls.Infof = append(mock.calls.Infof, callInfo)
	lockContextLoggerMockInfof.Unlock()
	mock.InfofFunc(in1, in2...)
}

// InfofCalls gets all the calls that were made to Infof.
// Check the length with:
//     len(mockedContextLogger.InfofCalls())
func (mock *ContextLoggerMock) InfofCalls() []struct {
	In1 string
	In2 []interface{}
} {
	var calls []struct {
		In1 string
		In2 []interface{}
	}
	lockContextLoggerMockInfof.RLock()
	calls = mock.calls.Infof
	lockContextLoggerMockInfof.RUnlock()
	return calls
}

// Panic calls PanicFunc.
func (mock *ContextLoggerMock) Panic(in1 ...interface{}) {
	if mock.PanicFunc == nil {
		panic("moq: ContextLoggerMock.PanicFunc is nil but ContextLogger.Panic was just called")
	}
	callInfo := struct {
		In1 []interface{}
	}{
		In1: in1,
	}
	lockContextLoggerMockPanic.Lock()
	mock.calls.Panic = append(mock.calls.Panic, callInfo)
	lockContextLoggerMockPanic.Unlock()
	mock.PanicFunc(in1...)
}

// PanicCalls gets all the calls that were made to Panic.
// Check the length with:
//     len(mockedContextLogger.PanicCalls())
func (mock *ContextLoggerMock) PanicCalls() []struct {
	In1 []interface{}
} {
	var calls []struct {
		In1 []interface{}
	}
	lockContextLoggerMockPanic.RLock()
	calls = mock.calls.Panic
	lockContextLoggerMockPanic.RUnlock()
	return calls
}

// Panicf calls PanicfFunc.
func (mock *ContextLoggerMock) Panicf(in1 string, in2 ...interface{}) {
	if mock.PanicfFunc == nil {
		panic("moq: ContextLoggerMock.PanicfFunc is nil but ContextLogger.Panicf was just called")
	}
	callInfo := struct {
		In1 string
		In2 []interface{}
	}{
		In1: in1,
		In2: in2,
	}
	lockContextLoggerMockPanicf.Lock()
	mock.calls.Panicf = append(mock.calls.Panicf, callInfo)
	lockContextLoggerMockPanicf.Unlock()
	mock.PanicfFunc(in1, in2...)
}

// PanicfCalls gets all the calls that were made to Panicf.
// Check the length with:
//     len(mockedContextLogger.PanicfCalls())
func (mock *ContextLoggerMock) PanicfCalls() []struct {
	In1 string
	In2 []interface{}
} {
	var calls []struct {
		In1 string
		In2 []interface{}
	}
	lockContextLoggerMockPanicf.RLock()
	calls = mock.calls.Panicf
	lockContextLoggerMockPanicf.RUnlock()
	return calls
}

// Warn calls WarnFunc.
func (mock *ContextLoggerMock) Warn(in1 ...interface{}) {
	if mock.WarnFunc == nil {
		panic("moq: ContextLoggerMock.WarnFunc is nil but ContextLogger.Warn was just called")
	}
	callInfo := struct {
		In1 []interface{}
	}{
		In1: in1,
	}
	lockContextLoggerMockWarn.Lock()
	mock.calls.Warn = append(mock.calls.Warn, callInfo)
	lockContextLoggerMockWarn.Unlock()
	mock.WarnFunc(in1...)
}

// WarnCalls gets all the calls that were made to Warn.
// Check the length with:
//     len(mockedContextLogger.WarnCalls())
func (mock *ContextLoggerMock) WarnCalls() []struct {
	In1 []interface{}
} {
	var calls []struct {
		In1 []interface{}
	}
	lockContextLoggerMockWarn.RLock()
	calls = mock.calls.Warn
	lockContextLoggerMockWarn.RUnlock()
	return calls
}

// Warnf calls WarnfFunc.
func (mock *ContextLoggerMock) Warnf(in1 string, in2 ...interface{}) {
	if mock.WarnfFunc == nil {
		panic("moq: ContextLoggerMock.WarnfFunc is nil but ContextLogger.Warnf was just called")
	}
	callInfo := struct {
		In1 string
		In2 []interface{}
	}{
		In1: in1,
		In2: in2,
	}
	lockContextLoggerMockWarnf.Lock()
	mock.calls.Warnf = append(mock.calls.Warnf, callInfo)
	lockContextLoggerMockWarnf.Unlock()
	mock.WarnfFunc(in1, in2...)
}

// WarnfCalls gets all the calls that were made to Warnf.
// Check the length with:
//     len(mockedContextLogger.WarnfCalls())
func (mock *ContextLoggerMock) WarnfCalls() []struct {
	In1 string
	In2 []interface{}
} {
	var calls []struct {
		In1 string
		In2 []interface{}
	}
	lockContextLoggerMockWarnf.RLock()
	calls = mock.calls.Warnf
	lockContextLoggerMockWarnf.RUnlock()
	return calls
}

// WithContext calls WithContextFunc.
func (mock *ContextLoggerMock) WithContext(in1 log.M) log.ContextLogger {
	if mock.WithContextFunc == nil {
		panic("moq: ContextLoggerMock.WithContextFunc is nil but ContextLogger.WithContext was just called")
	}
	callInfo := struct {
		In1 log.M
	}{
		In1: in1,
	}
	lockContextLoggerMockWithContext.Lock()
	mock.calls.WithContext = append(mock.calls.WithContext, callInfo)
	lockContextLoggerMockWithContext.Unlock()
	return mock.WithContextFunc(in1)
}

// WithContextCalls gets all the calls that were made to WithContext.
// Check the length with:
//     len(mockedContextLogger.WithContextCalls())
func (mock *ContextLoggerMock) WithContextCalls() []struct {
	In1 log.M
} {
	var calls []struct {
		In1 log.M
	}
	lockContextLoggerMockWithContext.RLock()
	calls = mock.calls.WithContext
	lockContextLoggerMockWithContext.RUnlock()
	return calls
}
