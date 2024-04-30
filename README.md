# axaDB

<p align="center">
  <img src="/resources/logo.png" width=300px>
</p>

## Short Description

axaDB is a NoSQL Document-Oriented Data Base Management System written exclusively in the GO programming language. This DBMS takes some optimisational inspiration from Oracle's Oracle Database to give you the best runtime for your C.R.U.D. programs.

## Speed and Security

axaDB is a fast and secure DBMS that encrypts all of your / your client's data, following the Advanced Encryption Standard (A.E.S.)

## Downloading and Installing

Installing axaDB is trivial. You only need two terminal utility programs (git and go). The steps to download the DBMS are:

 - Clone the git repository (I'm going to use https)
```sh
 $ git clone "https://github.com/kickhead13/axaDB.git"
```
 - Enter the cloned repo and build the project
```sh
 $ cd axaDB
 $ go build -o /path/to/your/axa/executable/axa ./src
```
 - Now you can access the executable at the (above) specified path

## Usage

You can run the axa executable with no parameters to get a useful help message.

```sh 
 $ axa
Usage: axa [PROGRAM] ...

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
*-----------------------------------------------------------*
```

As you can see the (very important) first step is setting up the AXADB_AEX_KEY (must be 32 characters long) and the AXADB_AES_IV (must be 16 characters long) environment variables. The way you'd do this is explained in the help message (at least for Unix based systems):

```sh
*-----------------------------------------------------------*
|$ export AXADB_AES_KEY="my32characterslongverysecretkey="\ |
| && export AXADB_AES_IV="my16charslongiv="                 |
*-----------------------------------------------------------*
```

Now you should run
```sh
 $ axa init ...
```
to initialize your database. Then run
```sh
 $ axa start ...
```
to boot up your database. Now, multiple clients can connect to your database and run querries / execute changes into your DB, using:
```sh
 $ axa connect ...
```
When you're done / want to close your DB for maintanance you should run
```sh
 $ axa halt ...
```

To get the help message of the init / connect / start / halt subprograms you must run ```axa <subprogram>``` with no parameters.

