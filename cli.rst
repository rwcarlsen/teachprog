
Command Line Interface (or Terminal)
=====================================

It is a program that runs programs.  You type in commands, and it runs
programs accordingly and prints the results to the screen.

Terminal Concepts:

  * File System: the terminal provides access to data stored on the
    computer's file system.  Generally, data are grouped into "files" that
    each have a unique handle consisting of textual characters called a
    "file path".  It is often convenient to think of a file system as a
    hierarchy of folders+files where slashes "/" in the file handle act as
    delimiters for subfolder (stored within) relationships.

    - Current working directory (cwd): the "location" or path that all terminal
      commands are executed from.

    - File Paths: a series of slash-separated words containing
      alpha-numeric and a few other characters (_, -, etc.) that uniquely
      identify every file on a file-system.  

      + Relative: path with respect to the cwd. Does NOT start with a slash::

        .                         (cwd)
        ./                        (trailing slash is optional)
        ../                       (parent directory - going out of or undoing the most recent dir)
        Pictures/2012/January
        Pictures/../Pictures/..   (cwd)


      + Absolute: path with respect to the root of the file system. DOES
        start with a slash::

        /                           (file-system root)
        /home/robert
        /home/robert/Pictures

      + Wildcards: Use the asterisk (*) to "match" zero or more of any
        character.  The terminal will expand your entry to all files in the
        cwd that match the pattern::

        cwd looks like:

          dir1  dir2  file1.txt  file2.txt  pic1.jpg  pic2.jpg  cheese.txt

        Wildcard expression               Expands to
        file*                             file1.txt file2.txt
        *.txt                             file1.txt file2.txt cheese.txt
        *2.*                              file2.txt pic2.txt

  * Piping and Redirection: http://www.datasavantconsulting.com/roland/piping.html

  * Environment

    - PATH
    - Variables

Running programs
-----------------

* Each command you give to the terminal is generally a single line of text.

* Hitting enter executes the command

* Many commands take "arguments" or extra parameters that instruct the
  program to act/run in a certain way.  A command and each specified
  argument are separated by one or more spaces.

  - Positional/normal arguments: Start without a dash - e.g. file names

  - option/flag arguments: Start with a dash

* Useful programs include:

  * ls - list directory contents
  * cd - change directory
  * mkdir - make a new directory
  * pwd - print the current working directory
  * cp - copy files from one location to another
  * rm - remove a file (cannot be undone!)
  * mv - move files (between directories - or renaming)
  * find - find files on the file system
  * grep - search for lines in files that match a search string
  * wc - output the line, word, char count of files
  * cat - concatenate the contents of named files
  * sort
  * tail - show the last few lines of named files
  * less - show content in scrollable/searchable document form (often piped into)
  * nano - a simple command line text editor

* Most programs can print out help information about how to use them when
  you run it with the "--help" flag argument::

    ls --help
    wc --help
    grep --help
    mkdir --help

  Or you can search on google for examples.

