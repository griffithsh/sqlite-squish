/*
command sqlite-squish provides an interface to convert between SQL statements
in plain-text *.sql files and sqlite databases.

Given a sqlite database called `database.sqlite3`, it is possible to
decompose it to a series of text files like this:

	sqlite3 database.sqlite3 .dump | sqlite-squish -v -out-dir ./database.sql

Given a directory of .sql files called `database.sql`, it is possible to
compose those files into a sqlite database like this:

	sqlite-squish -in database.sql | sqlite3 database.sqlite3

*/
package main
