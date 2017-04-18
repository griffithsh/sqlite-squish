# sqlite-squish
A tool to turn one or more plain-text \*.sql files into a sqlite database.

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

I want a way to create the database from plaintext \*.sql files, with meaningful
names, and commit those files, and not the binary representation of the
database at all. To do that, I need a tool that will take a directory of \*.sql
files, and create a sqlite database out of them.

This is fairly easy to achieve, but I also want to create foreign keys so some
tables depend on others. This requires a tool that understands the dependencies
of e.g. a `CREATE TABLE` statement, and makes sure the statements that create
those dependencies are run first.

## A solution?
Binary sqlite database files cannot be merged in git, but a plaintext set of
`CREATE TABLE`, `INSERT INTO` statements *can* be.

Changes to the game data and the structure that data is stored in could come
from an editing tool or sqlite GUI modifying the binary database file, or from
a text editor modifying the plaintext representation.

Therefore basic functionality must include:

- Turn a collection of text files containing sql statements into a sqlite
  database.
  - Must understand foreign keys and other dependencies and run the statements
    in the correct order.
- Export the schema and data from a sqlite database into a set of text files.
  - Must always output the same format and order, to eliminate diff noise.
- For any given sqlite database, the tool must be able to export, import, then
  export again without diffs from the first export, and the imported database
  must be logically equivalent to the original database.

## Further work

- [ ] Extend sqlite-squish to open and read from a sqlite database directly
- [ ] Extend sqlite-squish to create and write to a sqlite database directly
