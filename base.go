package alog

// Print creates informational logs based on the inputs
func Print(v ...interface{}) { Instance.Print(v) }

// Println prints the data coming in on individual lines
func Println(v ...interface{}) { Instance.Println(v) }

// Printf creates an informational log using the format and values
func Printf(format string, v ...interface{}) { Instance.Printf(format, v) }

// Warn creates a warning log using the error passed in along with the
// values passed in
func Warn(err error, v ...interface{}) { Instance.Warn(err, v) }

// Warnln creates a warning log using the error and values passed in.
// Each error and value is printed on a different line
func Warnln(err error, v ...interface{}) { Instance.Warnln(err, v) }

// Warnf creates a warning log using the error passed in, along with the string
// formatting and values
func Warnf(err error, format string, v ...interface{}) { Instance.Warnf(err, format, v) }

// Error creates an error log using the error and other values passed in
func Error(err error, v ...interface{}) { Instance.Error(err, v) }

// Errorln creates error logs using the error and other values passed in.
// Each error and value is printed on a different line
func Errorln(err error, v ...interface{}) { Instance.Errorln(err, v) }

// Errorf creates an error log using the error passed in, along with the string
// formatting and values
func Errorf(err error, format string, v ...interface{}) { Instance.Errorf(err, format, v) }

// Crit creates critical logs using the error and other values passed in
func Crit(err error, v ...interface{}) { Instance.Crit(err, v) }

// Critln creates critical logs using the error and other values passed in.
// Each error and value is printed on a different line
func Critln(err error, v ...interface{}) { Instance.Critln(err, v) }

// Critf creates a critical log using the error passed in, along with the string
// formatting and values
func Critf(err error, format string, v ...interface{}) { Instance.Critf(err, format, v) }

// Fatal creates a fatal log using the error and values passed into the method
// After logging the fatal log the Fatal method throws a panic to crash the application
func Fatal(err error, v ...interface{}) { Instance.Fatal(err, v) }

// Fatalln creates fatal logs using the error and other values passed in.
// Each error and value is printed on a different line
// After logging the fatal log the Fatalln method throws a panic to crash the application
func Fatalln(err error, v ...interface{}) { Instance.Fatalln(err, v) }

// Fatalf creates an error log using the error passed in, along with the string
// formatting and values
// After logging the fatal log the Fatalf method throws a panic to crash the application
func Fatalf(err error, format string, v ...interface{}) { Instance.Fatalf(err, format, v) }
