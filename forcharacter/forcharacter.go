package forcharacter

import "fmt"

func PrintCharacter()  {
  fmt.Print("\nfor_character 练习:\n\n")

  var str string = "G"
  for i:= 0; i < 25; i++ {
    fmt.Println(str)
    str += "G"
  }

}

