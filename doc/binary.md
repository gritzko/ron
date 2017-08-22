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
    * ;,.! (len=1 => ?)
 
2 - 2 (length: 4, 16)

prefixes: 4567

uuid: ttvvllll rrppnnnn nnnnnnnn ...
