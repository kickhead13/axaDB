package parsers

func AxaHelp() string{
	return `Usage: axa [PROGRAM] ...

PROGRAM:
   init              initializes database, creates the mandatory
                     files for the database
   start             boots up database server
   stop              shuts database server down
   exec              executes axa lang code on the database
ENVIRONMENT VARIABLES:
For the DBMS to run as it should, you will have to define two
environment variables. These variables are incredibly important.
They are the keys to encrypting / decrypting database data.
LOSING / FORGETTING / CHANGING them at any time will result in
the distruction of the integrity of your database. Be very
careful not to LOSE / FORGET / CHANGE any of these two after they
are set, and the database is created.
   AXADB_AES_KEY     A.E.S. algorithm key
   AXADB_AES_IV      A.E.S. algorithm I.V.
   `
}

func InitHelp() string{

	return `Usage: axa init --at[or -@] [DATABASE_DIRECTORY] [OPTIONS...]
Initialize axaDB database at a given location on the current server's disk.
Creates database directory, with given name and init.cfg config file.
	
DATABASE_DIRECTORY:          non-relative path to database directory
OPTIONS:
   --databaseName, -dbn       sets the detabase name (default is the name
                             of the DATABASE_DIRECTORY)
   --cpuCores, -cc:          sets the amount of cores available on the
                             machine, to be used by dispacher
                             (default is 4)
   --possibleBackups, -pb:   sets the amount of backups the database will
                             hold on the machine (default is 4)
   --maxDataFileSize, -pb:   sets the maximum size a single data file can
                             have before flush (default is 1024KB)
   --sysPassword, -sp        sets the password of 'sys' user 
                             (default is 'veryBadPassowrd')`
}