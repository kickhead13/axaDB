package dberrs

func DB_C01(host string) AxaErr{
  return AxaErr{`(axa err : db-c01) connection proccess : server is unavailable at ` + host + `
  ? input the correct ip / port / host for the axaDB server
  ? look up the status of the server (is it even on?)
  ! axa connect failed!`}
}

func DB_C02() AxaErr{
  return AxaErr{`(axa err : db-c02) connection process : login credentials are not valid
  ? either no username or no password
  ? or bad username and password combination
  ! axa connect failed!`}
}
