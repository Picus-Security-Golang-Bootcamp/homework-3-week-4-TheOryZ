package service

import (
	"flag"
	"strconv"
	"strings"
)

//SetFlags Added Flags informations [-help]
func SetFlags() {
	flag.String("list", "", "Returns the entire book list")
	flag.String("search", "", "Returns the detail information of the searched book [search <value>]")
	flag.String("get", "", "Returns the detail information of the entered Id book [get <id>]")
	flag.String("delete", "", "Delete the entered Id book [delete <id>]")
	flag.String("buy", "", "Buying the entered Id book with quantity [buy <id> <quantity>]")
	flag.Parse()
}

//ValidationCommand Function that checks the validity of the command entered by the user. Returns (true/false).
func ValidationCommand(commands []string) bool {
	if len(commands) < 1 {
		return false
	}
	if strings.ToLower(commands[0]) == "list" && len(commands) == 1 {
		return true
	} else if strings.ToLower(commands[0]) == "search" && len(commands) >= 2 {
		return true
	} else if strings.ToLower(commands[0]) == "get" && len(commands) == 2 {
		_, err := strconv.Atoi(commands[1])
		return err == nil
	} else if strings.ToLower(commands[0]) == "delete" && len(commands) == 2 {
		_, err := strconv.Atoi(commands[1])
		return err == nil
	} else if strings.ToLower(commands[0]) == "buy" && len(commands) == 3 {
		for i := 1; i < 3; i++ {
			_, err := strconv.Atoi(commands[i])
			if err != nil {
				return false
			}
		}
		return true
	}
	return false
}
