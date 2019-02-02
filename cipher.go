package main
/*
 * A Caeser cipher that prompts the user for an offset and message to encrpt,
 * and prints the encoded message
 *
 */
import (
  "fmt"
  "strings"
  "os"
  "bufio"
  "strconv"
)

var MAX_UPPER rune = 90
var MAX_LOWER rune = 122
var MIN_UPPER rune = 65
var MIN_LOWER rune = 97

func validateRune(ascii rune) bool {
  if (ascii < MIN_UPPER || ascii > MAX_LOWER) ||
     (ascii < MIN_LOWER && ascii > MAX_UPPER) {
       return false
  }
  return true
}

// calculates offset ascii value. Only offsets alphabetic characters (uppercase and lowercase)
func getAscii(ascii rune, offset rune) rune {
  var result rune
  isValid := validateRune(ascii)
  if (isValid) {
    var newAscii rune = ascii + offset
    if (newAscii >= MIN_UPPER && newAscii <= MAX_UPPER) ||
       (newAscii >= MIN_LOWER && newAscii <= MAX_LOWER) {
         result = newAscii
    } else if (newAscii > MAX_LOWER) {
      difference := newAscii - MAX_LOWER
      result = (MIN_LOWER - 1) + difference
    } else if (newAscii > MAX_UPPER) {
      difference := newAscii - MAX_UPPER
      result = (MIN_UPPER - 1) + difference
    }
  } else {
    result = ascii
  }

  return result
}

//
func cipher() {
  var offsetInput string
  reader := bufio.NewReader(os.Stdin)

  fmt.Println("Enter an offset (or type 'quit' to exit): ")
  offsetInput, _ = reader.ReadString('\n')

  offsetInput = strings.TrimSuffix(offsetInput, "\n")

  if (offsetInput == "quit") {
    os.Exit(2)
  }

  offsetInt, err := strconv.ParseInt(offsetInput, 10, 32)
  if (err != nil) {
    fmt.Println("Invalid offset. Please enter an integer")
    cipher()
  }

  offsetRune := int32(offsetInt)

  fmt.Println("Your offset is: ", offsetRune)

  var str strings.Builder

  fmt.Println("Enter a message to encode (or type 'quit' to exit): ")
  var messageInput string
  messageInput, _ = reader.ReadString('\n')

  messageInput = strings.TrimSuffix(messageInput, "\n")

  if (messageInput == "quit") {
    os.Exit(2)
  }

  i := 0
  for _, r := range messageInput {
    str.WriteString(string(getAscii(r, offsetRune)))
    i++
  }

  fmt.Println("Encoded string: ", str.String())
}

func main() {
  for {
    cipher()
  }
}
