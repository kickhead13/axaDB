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

func BoolParse(args []string, params []string) bool{
  for _, arg := range args {
    for _, param := range params {
      if arg == param {
        return true
      }
    }
  }
  return false
}

func ConnectParse(args []string, params []string) (string, error){
	return InitParse(args, params)
}
