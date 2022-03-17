package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	pl := fmt.Println //shorthand statements
	p := fmt.Print

	pl("Hello World!")
	var name string
	p("What is your name? : ")
	scanner := bufio.NewScanner(os.Stdin) //initialize a scanner to receive input
	scanner.Scan()
	name = scanner.Text()                                                      //set scanner to name
	now := time.Now()                                                          //get current time
	pl("Hello", name, "! It is", now.Format("3:15PM, Monday January 2, 2006")) //tell what time it is in a readable format.
	var valid = false
	var bday string
	var age int
	for !valid { //loop until they give valid input
		p("What is your birthday? Please enter it in YYYY-MM-DD format: ")
		scanner.Scan() //scan again
		bday = scanner.Text()
		bday = bday + "T15:04:05.000Z"                             //add needed but unused time data to convert to time value
		parse, err := time.Parse("2006-01-02T15:04:05.000Z", bday) //convert using format string of non-significant values (https://stackoverflow.com/questions/25845172/parsing-rfc-3339-iso-8601-date-time-string-in-go)
		if err != nil {
			pl("Try again!") //didnt give a valid format
		} else {
			age = now.Year() - parse.Year() //calculate what age they'll be this year
			pl("You were born on", parse.Format("Monday January 2, 2006"), ". You turn", age, "this year!")
			valid = true //move on
		}
	}
	valid = false
	for !valid {
		var num string
		p("Give me your favorite integer. ")
		scanner.Scan() //scan again
		num = scanner.Text()
		intnum, err := strconv.Atoi(num) //convert to int
		if err != nil {
			pl("Not an integer!") //invalid input
		} else {
			var newage = age + intnum         //find out what age they'll be int number of years from now
			var newyear = now.Year() + intnum //figure out what year that will be
			if intnum >= 0 {                  //is it positive? print this
				pl("Nice int! You know, in", intnum, "years, in", newyear, ", you will be", newage, "years old!")
			} else { //this will return a past year, so give them an appropriate message
				pl("Nice int! You know,", intnum*-1, "years ago, in", newyear, ", you were", newage, "years old!")
			}
			valid = true //move on
		}
	}
	pl("That's all folks!")
}
