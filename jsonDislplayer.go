package jsonDisplayer

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

// show file from json into terminal
func MapShow(param map[string]interface{}, space string) {

	for index, value := range param {

		if CheckType(value, "map") { // map into value

			fmt.Printf(space+" [%s] => \n", index)

			intrF, ok := value.(map[string]interface{})
			if ok {
				// add spaces
				addSpaceString := ""
				lenSpace := len(index)
				for j := 0; j < (lenSpace + 9); j++ {
					addSpaceString = addSpaceString + " "
				}

				MapShow(intrF, space+addSpaceString) //+"    "

			}

		} else {

			fmt.Println(space+"[", index, "] : ", "[", value, "]")
		}

	}

}

// create format text string from json
func MapIntoTextString(param map[string]interface{}, space string, text string) string {

	for index, value := range param {

		if CheckType(value, "map") { // map into value

			text += fmt.Sprintf(space+"[%s] => \n", index)

			intrF, ok := value.(map[string]interface{})
			if ok {
				// add spaces
				addSpaceString := ""
				lenSpace := len(index)
				for j := 0; j < (lenSpace + 6); j++ {
					addSpaceString = addSpaceString + " "
				}

				// mapShow(intrF, space+addSpaceString) //+"    "
				text = MapIntoTextString(intrF, space+addSpaceString, text)

			}

		} else {

			text += fmt.Sprint(space+"[", index, "] : ", "[", value, "]\n")

		}

	}

	return text
}

// write json string into file
func WriteBytesIntoFile(fileName string, stringForWrite string) {

	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(stringForWrite)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "\n bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}

// interface can be used like undefine type
// for check type of parameter by string value
func CheckType(param interface{}, stringName string) bool {
	reflectString := reflect.TypeOf(param)
	valueString := fmt.Sprint(reflectString)
	contain := strings.Contains(valueString, stringName)
	return contain
}
