
HW 1
=====

1. How many files are there::

   ls -l | wc -l

2. Which file has "hello" spelled with two ones (he11o) somewhere in it::

  "grep he11o ./*.txt" 

   or

  "grep he11o *.txt"

3. How many times does the word "the" appear in the output of "grep --help" (including things like oTHEr)::

   grep --help | grep the | wc -l

4. Move all files that have a "111" in the filename into a new directory named "have-ones"::

   mkdir have-ones
   mv *111* have-ones/

5. Move all files that have a "333" in the filename into a new directory named "have-threes"::

   mkdir have-threes
   cp *333* have-threes/

6. Combine all files that have a "222" in the filename into a file named "twos-all.txt"::

   cat *222* > twos-all.txt

7. How many files have a "11" and a "22" in the file name::

   ls -l | grep 11 | grep 22 | wc -l

8. Which of the last three words in every file is misspelled in one file (hint - the last three words are "banana", "apple", and "orange")::

   tail -n=3 *.txt | cat | grep banana | wc -l
   tail -n=3 *.txt | cat | grep apple | wc -l
   tail -n=3 *.txt | cat | grep orange | wc -l

   (or a slower way is to not use tail)

