/*a. Error handling
1. Create 2 files named a.txt b.txt in your current project directory.
2. Create your main.go file.
3. Inside main.go file, Write a function that takes the fileName as an input.
4. Inside the function, you have to call os.Remove() which would remove a file from the
current directory.
5. Now call this function from your main function & pass the above created file names to
this function.
Note: The os.remove function removes the file if it exists, or gives you an error if it does
not exist. You need to check & handle the returned error.
*/

package main

import (
	"fmt"
	"os"
)

func RemoveFile(fileName string) { //defining RemoveFile func to delete file
	err := os.Remove(fileName) //calling remove func to delete file
	if err != nil {            //check if any error
		fmt.Printf("File %s does not exists.\n", fileName)
		fmt.Println(err)
		return
	}
	fmt.Printf("File %s has been removed successfully.\n", fileName)
}

func main() {
	fileNameA := "a.txt" //creating file
	fileNameB := "c.txt" //creating file

	RemoveFile(fileNameA) //calling RemoveFile
	RemoveFile(fileNameB)

}
