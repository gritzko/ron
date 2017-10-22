# Binary RON

Idea: stream of explicit-width chunks, either UUIDs or atoms

* UUID
    * type
    * object
    * event
    * reference
* uuid scheme
    * event
    * derived
    * hash
    * name
* atoms
    * string
    * int
    * float
    * uuid
* terms
    * ;,?!
    * options: length to skip atoms (;,) or entire frame (!;)
 
2 - 2 (length: 4, 16)

uuid: ttvvllll fpppnnnn nnnnnnnn ...
prefix: 0..7 (7 bytes = 56 bit + 4 bits in the head)
perfect for 5 stable symbols, 6th symbol changing (pref 4 bytes = 32, 4 in the head)

If zipped, a binary RON CT document will be size of the plain text :)
