/*e. Logger usage
1. Create a package named "logger". Inside this package, create a struct named
"LoggerStore" with one unexported field named "log *log.Logger". Create a function
named as New which returns the instance of LoggerStore.
2. Call this logger.New() function from main function
Note:- you can get an instance of *log.Logger using the log.New() function
*/

package main

import (
	"assignment8-e/logger"
	"log"
)

func main() {
	log.Println(logger.New()) //calling new func by logger
}
