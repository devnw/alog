package alog

// Print creates informational logs based on the inputs
func Print(v ...interface{}) {
	instance.Print(v)
}

// Println prints the data coming in as an informational log on individual lines
func Println(v ...interface{}) {
	instance.Println(v)
}

// Printf creates an informational log using the format and values
func Printf(format string, v ...interface{}) {
	instance.Printf(format, v)
}

// Debug creates debugging logs based on the inputs
func Debug(v ...interface{}) {
	instance.Debug(v)
}

// Debugln prints the data coming in as a debug log on individual lines
func Debugln(v ...interface{}) {
	instance.Debugln(v)
}

// Debugf creates an debugging log using the format and values
func Debugf(format string, v ...interface{}) {
	instance.Debugf(format, v)
}

// Warn creates a warning log using the error passed in along with the
// values passed in
func Warn(err error, v ...interface{}) {
	instance.Warn(err, v)
}

// Warnln creates a warning log using the error and values passed in.
// Each error and value is printed on a different line
func Warnln(err error, v ...interface{}) {
	instance.Warnln(err, v)
}

// Warnf creates a warning log using the error passed in, along with the string
// formatting and values
func Warnf(err error, format string, v ...interface{}) {
	instance.Warnf(err, format, v)
}

// Error creates an error log using the error and other values passed in
func Error(err error, v ...interface{}) {
	instance.Error(err, v)
}

// Errorln creates error logs using the error and other values passed in.
// Each error and value is printed on a different line
func Errorln(err error, v ...interface{}) {
	instance.Errorln(err, v)
}

// Errorf creates an error log using the error passed in, along with the string
// formatting and values
func Errorf(err error, format string, v ...interface{}) {
	instance.Errorf(err, format, v)
}

// Crit creates critical logs using the error and other values passed in
func Crit(err error, v ...interface{}) {
	instance.Crit(err, v)
}

// Critln creates critical logs using the error and other values passed in.
// Each error and value is printed on a different line
func Critln(err error, v ...interface{}) {
	instance.Critln(err, v)
}

// Critf creates a critical log using the error passed in, along with the string
// formatting and values
func Critf(err error, format string, v ...interface{}) {
	instance.Critf(err, format, v)
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

// Custom creates a custom log using the error and values passed into the method
func Custom(ltype string, err error, v ...interface{}) {
	instance.Custom(ltype, err, v)
}

// Customln creates custom logs using the error and other values passed in.
// Each error and value is printed on a different line
func Customln(ltype string, err error, v ...interface{}) {
	instance.Customln(ltype, err, v)
}

// Customf creates a custom log using the error passed in, along with the string
// formatting and values
func Customf(ltype string, err error, format string, v ...interface{}) {
	instance.Customf(ltype, err, format, v)
}
