package parsers

func AxaHelp() string{
	return `Usage: axa [PROGRAM] ...

PROGRAM:
   init              initializes database, creates the mandatory
                     files for the database
   start             boots up database server
   halt              shuts database server down
   connect           connects user to specified axa server instance,
                     drops user to axa lang terminal from which to
                     run axa lang scriptss 
ENVIRONMENT VARIABLES:
For the DBMS to run as it should, you will have to define two
environment variables. These variables are incredibly important.
They are the keys to encrypting / decrypting database data.
LOSING / FORGETTING / CHANGING them at any time will result in
the distruction of the integrity of your database. Be very
careful not to LOSE / FORGET / CHANGE any of these two after they
are set, and the database is created. DO NOT SHARE any of the 
following to people you - or your org - do not trust.
   AXADB_AES_KEY     A.E.S. algorithm key (must be 32 characters)
   AXADB_AES_IV      A.E.S. algorithm I.V. (must be 16 characters)
On Unix based systems:
*-----------------------------------------------------------*
|$ export AXADB_AES_KEY="my32characterslongverysecretkey="\ |
| && export AXADB_AES_IV="my16charslongiv="                 |
*-----------------------------------------------------------*`
}

func InitHelp() string{

	return `Usage: axa init --at[or -@] [DATABASE_DIRECTORY] [OPTIONS...]
Initialize axaDB database at a given location on the current server's disk.
Creates database directory, with given name and init.cfg config file.
	
DATABASE_DIRECTORY:                non-relative path to database directory
OPTIONS:
   --databaseName, -dbn <NAME>:    sets the detabase name (default is the name
                                   of the DATABASE_DIRECTORY)
   --cpuCores, -cc <N>:            sets the amount of cores available on the
                                   machine, to be used by dispacher
                                   (default is 4)
   --possibleBackups, -pb <N>:     sets the amount of backups the database will
                                   hold on the machine (default is 4)
   --maxDataFileSize, -pb <N>:     sets the maximum size a single data file can
                                   have before flush (default is 1024KB)
   --sysPassword, -sp <N>:         sets the password of 'sys' user 
                                   (default is 'veryBadPassowrd')`
}

func ConnectHelp() string{
  return `Usage: axa connect [OPTIONS...] [--no-exec-term <COMMAND>]
Connects / logs (in) user to the given axaDB server, and opens axa exec 
terminal (if the user desires it) so that they can perform commands on the
database.

OPTIONS:
   --ip, -i <IP>:          server's ip (default is 127.0.0.1 / localhost)
   --port, -p <PORT>:              server's port (default is 13131)
   --host, -h <HOST> (<IP:PORT):   server's ip:host combination 
                                   (default is localhost:13131, this
                                   does not have to be mentioned if --ip 
                                   and --port are given)
   --login, -l <USERNAME>:         user's username (default: user)
   --password, -pass <PASSWORD>:   user's paswword (default: pass)
   --no-exec-term <COMMAND>:       does not drop user to axa exec term but
                                   instead executes command and then logs
                                   user out
COMMAND:
   * C.R.U.D. commands are of the form:
     [KEYWORD] [in/from] <COLLECTION> [[at <key1>], [at <key2], ...] <JSON>;
KEYWORD:
   * feed (in <COLLECTION_NAME> [at ...] <JSON>):
      "feeds" the database the JSON given by user at the sepcified path
      in the specified collection
   * fetch (from <COLLECTION_NAME> [at ...])
      "fetches" the json found within the database in the COLLECTION at the
      specified path
   * delete (from <COLLECTION> [at ...])
      "deleted" from the database all the information found within the 
      COLLECTION at the specified path
OTHER COMMANDS:
   * login <USERNAME> <PASSWORD>
      sends a message to server and receives either "TRUE" or "FALSE" 
      depending on the validity of the credentials given
   * adduser <USERNAME> <PASSWORD>
      if logged in as user with AXA_ADMIN privileges, it will add new 
      user to the AXA_USER collection
   * addrole <ROLE>
      if logged in as user with AXA_ADMIN privileges, it will create new 
      role 
   * new <COLLECTION> <ROLES_JSON>
      creates new collection (COLLECTION) with the rules provided via the
      RULES_JSON json`
}

