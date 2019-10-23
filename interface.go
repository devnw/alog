package alog

import "io"

// Logger provides the ability to log using different methods asynchronously
type Logger interface {
	// Print creates informational logs based on the inputs
	Print(v ...interface{})

	// Println prints the data coming in as an informational log on individual lines
	Println(v ...interface{})

	// Printf creates an informational log using the format and values
	Printf(format string, v ...interface{})

	// Debug creates debugging logs based on the inputs
	Debug(v ...interface{})

	// Debugln prints the data coming in as a debug log on individual lines
	Debugln(v ...interface{})

	// Debugf creates an debugging log using the format and values
	Debugf(format string, v ...interface{})

	// Warn creates a warning log using the error passed in along with the
	// values passed in
	Warn(err error, v ...interface{})

	// Warnln creates a warning log using the error and values passed in.
	// Each error and value is printed on a different line
	Warnln(err error, v ...interface{})

	// Warnf creates a warning log using the error passed in, along with the string
	// formatting and values
	Warnf(err error, format string, v ...interface{})

	// Error creates an error log using the error and other values passed in
	Error(err error, v ...interface{})

	// Error creates error logs using the error and other values passed in.
	// Each error and value is printed on a different line
	Errorln(err error, v ...interface{})

	// Errorf creates an error log using the error passed in, along with the string
	// formatting and values
	Errorf(err error, format string, v ...interface{})

	// Crit creates critical logs using the error and other values passed in
	Crit(err error, v ...interface{})

	// Critln creates critical logs using the error and other values passed in.
	// Each error and value is printed on a different line
	Critln(err error, v ...interface{})

	// Critf creates a critical log using the error passed in, along with the string
	// formatting and values
	Critf(err error, format string, v ...interface{})

	// Fatal creates a fatal log using the error and values passed into the method
	// After logging the fatal log the Fatal method throws a panic to crash the application
	Fatal(err error, v ...interface{})

	// Fatalln creates fatal logs using the error and other values passed in.
	// Each error and value is printed on a different line
	// After logging the fatal log the Fatalln method throws a panic to crash the application
	Fatalln(err error, v ...interface{})

	// Fatalf creates an error log using the error passed in, along with the string
	// formatting and values
	// After logging the fatal log the Fatalf method throws a panic to crash the application
	Fatalf(err error, format string, v ...interface{})

	// AddOutput adds an additional logging source to the logger which
	// will be added to the different logging outputs for this logger
	AddOutput(out io.Writer)

	// Close cancels the context throughout the logger and closes
	// all read / write operations accross the logger and IO
	Close()
}
