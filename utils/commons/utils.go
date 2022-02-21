package commons

import (
	"fmt"
	"os"
)

func LoadEnvCommandLineVariables() {
	argsForProg := os.Args[1:]
	fmt.Println("setting env variables ", argsForProg)
	for i, v := range argsForProg {
		fmt.Println("args passed ", i, v)
		switch i {
		case 0:
			os.Setenv("mysql_users_username", v)
		case 1:
			os.Setenv("mysql_users_password", v)
		case 3:
			os.Setenv("mysql_users_host", v)
		case 4:
			os.Setenv("mysql_users_schema", v)
		}
	}
}
