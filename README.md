# sqlite-squish

Command line tools to turn one or more plain-text \*.sql files into a sqlite
database and back again into text.

## Problem Statement

I work on a game that stores its game data (ie, level geometry, animations,
item type definitions, skills, etc) in a sqlite database. As part of that work,
I often commit changes to the sqlite database. Because sqlite databases are
binary files, no meaningful diffs can be committed, and merges cannot be
performed.

My first approach to solving this was to use the sqlite .dump command to
export the database, which worked reasonably well. The place where I'm
unsatisfied with this approach though is that I have no control over how the
data is exported; I just get a single output file, which is difficult to read
and edit. As the database size grew, the dumped file became more and more
unwieldy. Ideally there would be one output file per table.

I also need to be able to compose one or more .sql files into a sqlite
database, to regenerate the database after cloning, merging or pulling.

This repository holds my solution to that problem.

## Commands

[sqlite-squish](cmd/sqlite-squish) This is a fully featured tool to convert _from_ one of:

- one or more sql statements piped in via stdin
- a sqlite database
- or a directory of *.sql files

and _to_ one of:

- stdout
- a sqlite database
- or a directory of *.sql files.

It has a several flags to drive behaviour due to it's more complicated feature set.

[compose-sqlite](cmd/compose-sqlite) This is a simple command to compose a
directory of *.sql files into a sqlite database. It takes one argument to
specify the directory to compose.

[decompose-sqlite](cmd/decompose/sqlite) This is a simple command to
decompose a sqlite database into a directory of *.sql files. It takes one
argument to specify the name of the database to decompose.
