package alog

import "context"

// Printc creates informational logs based on the data coming from the
// concurrency channel that is passed in for processing
func Printc(ctx context.Context, v <-chan interface{}) {
	instance.Printc(ctx, v)
}

// Print creates informational logs based on the inputs
func Print(v ...interface{}) error {
	return instance.Print(v)
}

// Println prints the data coming in as an informational log on individual lines
func Println(v ...interface{}) error {
	return instance.Println(v)
}

// Printf creates an informational log using the format and values
func Printf(format string, v ...interface{}) error {
	return instance.Printf(format, v)
}

// Debugc creates debug logs based on the data coming from the
// concurrency channel that is passed in for processing
func Debugc(ctx context.Context, v <-chan interface{}) {
	instance.Debugc(ctx, v)
}

// Debug creates debugging logs based on the inputs
func Debug(err error, v ...interface{}) error {
	return instance.Debug(err, v)
}

// Debugln prints the data coming in as a debug log on individual lines
func Debugln(err error, v ...interface{}) error {
	return instance.Debugln(err, v)
}

// Debugf creates an debugging log using the format and values
func Debugf(err error, format string, v ...interface{}) error {
	return instance.Debugf(err, format, v)
}

// Warnc creates warning logs based on the data coming from the
// concurrency channel that is passed in for processing
func Warnc(ctx context.Context, v <-chan interface{}) {
	instance.Warnc(ctx, v)
}

// Warn creates a warning log using the error passed in along with the
// values passed in
func Warn(err error, v ...interface{}) error {
	return instance.Warn(err, v)
}

// Warnln creates a warning log using the error and values passed in.
// Each error and value is printed on a different line
func Warnln(err error, v ...interface{}) error {
	return instance.Warnln(err, v)
}

// Warnf creates a warning log using the error passed in, along with the string
// formatting and values
func Warnf(err error, format string, v ...interface{}) error {
	return instance.Warnf(err, format, v)
}

// Errorc creates error logs based on the data coming from the
// concurrency channel that is passed in for processing
func Errorc(ctx context.Context, v <-chan interface{}) {
	instance.Errorc(ctx, v)
}

// Error creates an error log using the error and other values passed in
func Error(err error, v ...interface{}) error {
	return instance.Error(err, v)
}

// Errorln creates error logs using the error and other values passed in.
// Each error and value is printed on a different line
func Errorln(err error, v ...interface{}) error {
	return instance.Errorln(err, v)
}

// Errorf creates an error log using the error passed in, along with the string
// formatting and values
func Errorf(err error, format string, v ...interface{}) error {
	return instance.Errorf(err, format, v)
}

// Critc creates critical logs based on the data coming from the
// concurrency channel that is passed in for processing
func Critc(ctx context.Context, v <-chan interface{}) {
	instance.Critc(ctx, v)
}

// Crit creates critical logs using the error and other values passed in
func Crit(err error, v ...interface{}) error {
	return instance.Crit(err, v)
}

// Critln creates critical logs using the error and other values passed in.
// Each error and value is printed on a different line
func Critln(err error, v ...interface{}) error {
	return instance.Critln(err, v)
}

// Critf creates a critical log using the error passed in, along with the string
// formatting and values
func Critf(err error, format string, v ...interface{}) error {
	return instance.Critf(err, format, v)
}

// Fatal creates a fatal log using the error and values passed into the method
// After logging the fatal log the Fatal method throws a panic to crash the application
func Fatal(err error, v ...interface{}) {
	instance.Fatal(err, v)
}

// Fatalln creates fatal logs using the error and other values passed in.
// Each error and value is printed on a different line
// After logging the fatal log the Fatalln method throws a panic to crash the application
func Fatalln(err error, v ...interface{}) {
	instance.Fatalln(err, v)
}

// Fatalf creates an error log using the error passed in, along with the string
// formatting and values
// After logging the fatal log the Fatalf method throws a panic to crash the application
func Fatalf(err error, format string, v ...interface{}) {
	instance.Fatalf(err, format, v)
}

// Customc creates custom logs based on the data coming from the
// concurrency channel that is passed in for processing
func Customc(ctx context.Context, v <-chan interface{}, ltype string) {
	instance.Customc(ctx, v, ltype)
}

// Custom creates a custom log using the error and values passed into the method
func Custom(ltype string, err error, v ...interface{}) error {
	return instance.Custom(ltype, err, v)
}

// Customln creates custom logs using the error and other values passed in.
// Each error and value is printed on a different line
func Customln(ltype string, err error, v ...interface{}) error {
	return instance.Customln(ltype, err, v)
}

// Customf creates a custom log using the error passed in, along with the string
// formatting and values
func Customf(ltype string, err error, format string, v ...interface{}) error {
	return instance.Customf(ltype, err, format, v)
}

// Close cancels the context throughout the logger and closes
// all read / write operations across the logger and IO
func Close() {
	instance.Close()
}

// Wait blocks on the logger context until the context is closed
func Wait() {
	instance.Wait()
}
