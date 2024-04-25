# axa init utility

## Scope
Creates database at given location on disk, with the above mentioned specifications

## Usage
```sh
$ axa init
Usage: axa init --at [DATABASE_DIRECTORY] [OPTIONS...]
Initialize axaDB database at a given location on the current server's disk.
Creates database directory, with given name and init.cfg config file.

DATABASE_DIRECTORY:          non-relative path to database directory
OPTIONS:
   --cpuCores, -cc:          sets the amount of cores available on the
                             machine, to be used by dispacher
                             (default is 4)
   --possibleBackups, -pb:   sets the amount of backups the database will
                             hold on the machine
                             (default is 4)
   --maxDataFileSize, -pb:   sets the maximum size a single data file can
                             have before flush (default is 1024KB)
```