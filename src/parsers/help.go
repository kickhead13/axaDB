package parsers

func AxaHelp() string{
	return "Usage: axa [PROGRAM] ...\nPROGRAM:\n   init		initializes database, creates the mandatory\n		files for the database\n   start	boots up database server\n   stop		shuts database server down\n   exec		executes axa lang code on the database"
}

func InitHelp() string{
	help1  := "Usage: axa init --at [DATABASE_DIRECTORY] [OPTIONS...]" 
	help2  := "\nInitialize axaDB database at a given location on the current server's disk."
	help3  := "\nCreates database directory, with given name and init.cfg config file.\n"
	help4  := "\nDATABASE_DIRECTORY:          non-relative path to database directory"
	help5  := "\nOPTIONS:"
	help6  := "\n   --cpuCores, -cc:          sets the amount of cores available on the"
	help7  := "\n                             machine, to be used by dispacher"
	help8  := "\n                             (default is 4)"
	help9  := "\n   --possibleBackups, -pb:   sets the amount of backups the database will"
	help10 := "\n                             hold on the machine"
	help11 := "\n                             (default is 4)"
	help12 := "\n   --maxDataFileSize, -pb:   sets the maximum size a single data file can"
	help13 := "\n                             have before flush (default is 1024KB)"
	return help1 + help2 + help3 + help4 + help5 + help6 + help7 + help8 + help9 + help10 + help11 + help12 + help13
}