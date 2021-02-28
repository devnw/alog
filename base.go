// Copyright © 2019 Developer Network, LLC
//
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of this source code package.

package alog

import "context"

// Printc creates informational logs based on the data coming from the
// concurrency channel that is passed in for processing
func Printc(ctx context.Context, v <-chan interface{}) {
	global.Printc(ctx, v)
}

// Print creates informational logs based on the inputs
func Print(v ...interface{}) {
	global.Print(v...)
}

// Println prints the data coming in as an informational log on individual lines
func Println(v ...interface{}) {
	global.Println(v...)
}

// Printf creates an informational log using the format and values
func Printf(format string, v ...interface{}) {
	global.Printf(format, v...)
}

// Debugc creates debug logs based on the data coming from the
// concurrency channel that is passed in for processing
func Debugc(ctx context.Context, v <-chan interface{}) {
	global.Debugc(ctx, v)
}

// Debug creates debugging logs based on the inputs
func Debug(err error, v ...interface{}) {
	global.Debug(err, v...)
}

// Debugln prints the data coming in as a debug log on individual lines
func Debugln(err error, v ...interface{}) {
	global.Debugln(err, v...)
}

// Debugf creates an debugging log using the format and values
func Debugf(err error, format string, v ...interface{}) {
	global.Debugf(err, format, v...)
}

// Tracec creates trace logs based on the data coming from the
// concurrency channel that is passed in for processing
func Tracec(ctx context.Context, v <-chan interface{}) {
	global.Tracec(ctx, v)
}

// Trace creates trace logs based on the inputs
func Trace(err error, v ...interface{}) {
	global.Trace(err, v...)
}

// Traceln prints the data coming in as a trace log on individual lines
func Traceln(err error, v ...interface{}) {
	global.Traceln(err, v...)
}

// Tracef creates an trace log using the format and values
func Tracef(err error, format string, v ...interface{}) {
	global.Tracef(err, format, v...)
}

// Warnc creates warning logs based on the data coming from the
// concurrency channel that is passed in for processing
func Warnc(ctx context.Context, v <-chan interface{}) {
	global.Warnc(ctx, v)
}

// Warn creates a warning log using the error passed in along with the
// values passed in
func Warn(err error, v ...interface{}) {
	global.Warn(err, v...)
}

// Warnln creates a warning log using the error and values passed in.
// Each error and value is printed on a different line
func Warnln(err error, v ...interface{}) {
	global.Warnln(err, v...)
}

// Warnf creates a warning log using the error passed in, along with the string
// formatting and values
func Warnf(err error, format string, v ...interface{}) {
	global.Warnf(err, format, v...)
}

// Errorc creates error logs based on the data coming from the
// concurrency channel that is passed in for processing
func Errorc(ctx context.Context, v <-chan interface{}) {
	global.Errorc(ctx, v)
}

// Error creates an error log using the error and other values passed in
func Error(err error, v ...interface{}) {
	global.Error(err, v...)
}

// Errorln creates error logs using the error and other values passed in.
// Each error and value is printed on a different line
func Errorln(err error, v ...interface{}) {
	global.Errorln(err, v...)
}

// Errorf creates an error log using the error passed in, along with the string
// formatting and values
func Errorf(err error, format string, v ...interface{}) {
	global.Errorf(err, format, v...)
}

// Critc creates critical logs based on the data coming from the
// concurrency channel that is passed in for processing
func Critc(ctx context.Context, v <-chan interface{}) {
	global.Critc(ctx, v)
}

// Crit creates critical logs using the error and other values passed in
func Crit(err error, v ...interface{}) {
	global.Crit(err, v...)
}

// Critln creates critical logs using the error and other values passed in.
// Each error and value is printed on a different line
func Critln(err error, v ...interface{}) {
	global.Critln(err, v...)
}

// Critf creates a critical log using the error passed in, along with the string
// formatting and values
func Critf(err error, format string, v ...interface{}) {
	global.Critf(err, format, v...)
}

// Fatalc creates fatal logs based on the data coming from the
// concurrency channel that is passed in for processing
func Fatalc(ctx context.Context, v <-chan interface{}) {
	global.Fatalc(ctx, v)
}

// Fatal creates a fatal log using the error and values passed into the method
func Fatal(err error, v ...interface{}) {
	global.Fatal(err, v...)
}

// Fatalln creates fatal logs using the error and other values passed in.
// Each error and value is printed on a different line
func Fatalln(err error, v ...interface{}) {
	global.Fatalln(err, v...)
}

// Fatalf creates an error log using the error passed in, along with the string
// formatting and values
func Fatalf(err error, format string, v ...interface{}) {
	global.Fatalf(err, format, v...)
}

// Customc creates custom logs based on the data coming from the
// concurrency channel that is passed in for processing
func Customc(ctx context.Context, v <-chan interface{}, ltype string) {
	global.Customc(ctx, v, ltype)
}

// Custom creates a custom log using the error and values passed into the method
func Custom(ltype string, err error, v ...interface{}) {
	global.Custom(ltype, err, v...)
}

// Customln creates custom logs using the error and other values passed in.
// Each error and value is printed on a different line
func Customln(ltype string, err error, v ...interface{}) {
	global.Customln(ltype, err, v...)
}

// Customf creates a custom log using the error passed in, along with the string
// formatting and values
func Customf(ltype string, err error, format string, v ...interface{}) {
	global.Customf(ltype, err, format, v...)
}

// Close cancels the context throughout the logger and closes
// all read / write operations across the logger and IO
func Close() {
	global.Close()
}

// Wait blocks on the logger context until the context is closed
func Wait(close bool) {
	global.Wait(close)
}
