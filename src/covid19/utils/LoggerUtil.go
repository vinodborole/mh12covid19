package utils

import (
	log "github.com/sirupsen/logrus"
)

//Infoln log info msg
func Infoln(args ...interface{}) {

	log.Infoln(args...)

}

//Debugln log debug msg
func Debugln(args ...interface{}) {

	log.Debugln(args...)

}

//Errorln log error msg
func Errorln(args ...interface{}) {

	log.Errorln(args...)

}

//Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {

	log.Errorf(format, args...)

}

//Warnln log warn msg
func Warnln(args ...interface{}) {

	log.Warnln(args...)

}
