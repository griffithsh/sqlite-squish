# sqlite-squish
A tool to turn one or more plain-text *.sql files into a sqlite database.

## The problem
I work on a game that stores its game data (ie, level geometry, animations,
item type definitions, skills, etc) in a sqlite database. As part of that work,
I often commit changes to the sqlite database. Because sqlite databases are
binary files, no meaningful diffs can be committed, and merges cannot be
performed.

My first approach to solving this was to use the sqlite dump command to export
the database, which worked reasonably well. The place where I'm unsatisfied
with this approach though is that I have no control over how the data is
exported; I just get a single output file, which is difficult to read and edit.

I want a way to create the database from plaintext *.sql files, with meaningful
names, and commit those files, and not the binary representation of the
database at all. To do that, I need a tool that will take a directory of *.sql
files, and create a sqlite database out of them.

This is fairly easy to achieve, but I also want to create foreign keys so some
tables depend on others. This requires a tool that understands the dependencies
of e.g. a `CREATE TABLE` statement, and makes sure the statements that create
those dependencies are run first.

