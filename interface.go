// Copyright © 2019 Developer Network, LLC
//
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of this source code package.

package alog

import (
	"context"
)

// Logger provides the ability to log using different methods asynchronously
type Logger interface {

	// Printc creates informational logs based on the data coming from the
	// concurrency channel that is passed in for processing
	Printc(ctx context.Context, v <-chan interface{})

	// Print creates informational logs based on the inputs
	Print(v ...interface{})

	// Println prints the data coming in as an informational log on individual lines
	Println(v ...interface{})

	// Printf creates an informational log using the format and values
	Printf(format string, v ...interface{})

	// Debugc creates debug logs based on the data coming from the
	// concurrency channel that is passed in for processing
	Debugc(ctx context.Context, v <-chan interface{})

	// Debug creates debugging logs based on the inputs
	Debug(err error, v ...interface{})

	// Debugln prints the data coming in as a debug log on individual lines
	Debugln(err error, v ...interface{})

	// Debugf creates an debugging log using the format and values
	Debugf(err error, format string, v ...interface{})

	// Tracec creates trace logs based on the data coming from the
	// concurrency channel that is passed in for processing
	Tracec(ctx context.Context, v <-chan interface{})

	// Trace creates trace logs based on the inputs
	Trace(err error, v ...interface{})

	// Traceln prints the data coming in as a trace log on individual lines
	Traceln(err error, v ...interface{})

	// Tracef creates an trace log using the format and values
	Tracef(err error, format string, v ...interface{})

	// Warnc creates warning logs based on the data coming from the
	// concurrency channel that is passed in for processing
	Warnc(ctx context.Context, v <-chan interface{})

	// Warn creates a warning log using the error passed in along with the
	// values passed in
	Warn(err error, v ...interface{})

	// Warnln creates a warning log using the error and values passed in.
	// Each error and value is printed on a different line
	Warnln(err error, v ...interface{})

	// Warnf creates a warning log using the error passed in, along with the string
	// formatting and values
	Warnf(err error, format string, v ...interface{})

	// Errorc creates error logs based on the data coming from the
	// concurrency channel that is passed in for processing
	Errorc(ctx context.Context, v <-chan interface{})

	// Error creates an error log using the error and other values passed in
	Error(err error, v ...interface{})

	// Error creates error logs using the error and other values passed in.
	// Each error and value is printed on a different line
	Errorln(err error, v ...interface{})

	// Errorf creates an error log using the error passed in, along with the string
	// formatting and values
	Errorf(err error, format string, v ...interface{})

	// Critc creates critical logs based on the data coming from the
	// concurrency channel that is passed in for processing
	Critc(ctx context.Context, v <-chan interface{})

	// Crit creates critical logs using the error and other values passed in
	Crit(err error, v ...interface{})

	// Critln creates critical logs using the error and other values passed in.
	// Each error and value is printed on a different line
	Critln(err error, v ...interface{})

	// Critf creates a critical log using the error passed in, along with the string
	// formatting and values
	Critf(err error, format string, v ...interface{})

	// Fatalc creates fatal logs based on the data coming from the
	// concurrency channel that is passed in for processing
	Fatalc(ctx context.Context, v <-chan interface{})

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

	// Customc creates custom logs based on the data coming from the
	// concurrency channel that is passed in for processing
	Customc(ctx context.Context, v <-chan interface{}, ltype string)

	// Custom creates a custom log using the error and values passed into the method
	Custom(ltype string, err error, v ...interface{})

	// Customln creates custom logs using the error and other values passed in.
	// Each error and value is printed on a different line
	Customln(ltype string, err error, v ...interface{})

	// Customf creates a custom log using the error passed in, along with the string
	// formatting and values
	Customf(ltype string, err error, format string, v ...interface{})

	// Close cancels the context throughout the logger and closes
	// all read / write operations across the logger and IO
	Close()

	// Wait blocks on the logger context until the context is closed
	Wait(close bool)
}
