package parsers

import "errors"

func InitParse(args []string, params []string) (string, error) {
	for iter, arg := range args {
		for _, param := range params {
			if arg == param && len(args) > iter+1 {
				if len(args) > iter+1 {
					return args[iter+1], nil
				} else {
					return args[0], errors.New("params: not enough params for init")
				}
			}
		}
	}
	return "~", errors.New("params: param not found")
}