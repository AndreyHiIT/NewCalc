package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv" 
)
var romanTOarabic = map[string]int{
	"C": 100,
	"L": 50, 
  "X": 10, 
  "V": 5, 
  "I": 1,
  }

func main() {
  reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите пример: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, " ", "")
  
  operator, err := findOp(text)
		if err != nil {
		panic(err)
  }
   
  bool1:=format(text)
  if bool1==false {
    panic("Неверный формат")
  }
  
  arguments := strings.Split(text, operator)
  argRim1:=inRoman(arguments[0])
  argRim2:=inRoman(arguments[1])
  if argRim1 && argRim2 {
    a:=toInt(arguments[0])
    b:=toInt(arguments[1])
    resultRom,_:=calculation(a,b, operator)
    if resultRom<=0 {
      panic("Римские значения не могут быть отрицательными")
    }
    finalResult:=intToRoman(resultRom)
    fmt.Println(finalResult)
  } else if argRim1 != argRim2 {
    panic("Разный формат значений недопустим")
  } else {
  argInt1, _ := strconv.Atoi(arguments[0])
  argInt2, _ := strconv.Atoi(arguments[1])
  if err := check(argInt1, argInt2); 
  err != nil {
    panic(err)
  }
  result, _:=calculation(argInt1, argInt2, operator)
  fmt.Println(result)
  } 
}

func check(a1,a2 int)error{
  if a1 < 1 || a1 > 10 || a2 < 1 || a2 > 10 {
    return fmt.Errorf("Одно или оба числа выходят за пределы диапазона от 1 до 10.")
  }
  return nil
}
func findOp(text string) (string, error) {
	switch {
	  case strings.Contains(text, "+"):
		return "+", nil
	  case strings.Contains(text, "-"):
		return "-", nil
	  case strings.Contains(text, "*"):
		return "*", nil
  	case strings.Contains(text, "/"):
		return "/", nil
  	default:
		return "", fmt.Errorf("Неверная операция.Доступные операции: +, -, /, *.")
	}
}
func calculation(a1, a2 int, op string)(result int, err error){
  switch op{
    case "+":
    result=a1+a2
    case "-":
    result=a1-a2
    case "/":
    result=a1/a2
    case "*":
    result=a1*a2
    default: err=fmt.Errorf("Неизвестный %s", op)
  }
  return
}
//проверка на наличие римских
func inRoman(arg string) bool{
  roman:=map[string]int{
  "I":1,
  "II":2,
  "III":3,
  "IV":4,
  "V":5,
  "VI":6,
  "VII":7,
  "VIII":8,
  "IX":9,
  "X":10,
  }
  _,err := roman[arg]
  if !err {
    return false
  } 
  return true
}

func toInt(arg string) int{
  roman:=map[string]int{
  "I":1,
  "II":2,
  "III":3,
  "IV":4,
  "V":5,
  "VI":6,
  "VII":7,
  "VIII":8,
  "IX":9,
  "X":10,
  }
  result, _ := roman[arg]
  return result
  
}
func intToRoman(num int) string {
    val := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
    sym := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    roman:=""
    i := 0
    for num > 0 {
        for num >= val[i] {
            roman += sym[i]
            num -= val[i]
        }
        i++
    }
    return roman
}
   
func format(text string)bool{
   op:=[]rune{'-','+','*','/'}
   count:=0
   for _,valText:=range text{
     for _,valOp:=range op{
       if valText==valOp {
         count++
       }
     }
   }
   if count==1 {
    return true
   } else {
    return false
   } 
}
