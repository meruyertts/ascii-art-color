# ascii-art-color

Ascii-art is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII. 

A graphic representation using ASCII, is to write the string received using ASCII characters, as you can see in the example below:  
```
@@@@@@BB@@@@``^^``^^``@@BB$$@@BB$$
@@%%$$$$^^^^WW&&8888&&^^""BBBB@@@@
@@@@@@""WW8888&&WW888888WW``@@@@$$
BB$$``&&&&WWWW8888&&&&8888&&``@@@@
$$``&&WW88&&88&&&&8888&&88WW88``$$
@@""&&&&&&&&88888888&&&&&&88&&``$$
``````^^``^^^^^^````""^^``^^``^^``
""WW^^@@@@^^``````^^BB@@^^``^^&&``
^^&&^^@@````^^``&&``@@````^^^^&&``
``WW&&^^""``^^WW&&&&""``^^^^&&88``
^^8888&&&&&&WW88&&88WW&&&&88&&WW``
@@``&&88888888WW&&WW88&&88WW88^^$$
@@""88&&&&&&&&888888&&``^^&&88``$$
@@@@^^&&&&&&""``^^^^^^8888&&^^@@@@
@@@@@@^^888888&&88&&&&MM88^^BB$$$$
@@@@@@BB````&&&&&&&&88""``BB@@BB$$
$$@@$$$$$$$$``````````@@$$@@$$$$$$
```
This project handles an input with numbers, letters, spaces, special characters and \n.  

Some banner files with a specific graphical template representation using ASCII will be given. The files are formatted in a way that is not necessary to change them.

* shadow
* standard
* thinkertoy

The output manipulates colors using the flag --color=<color> <letters to be colored>, in which --color is the flag and <color> is the color desired by the user and <letters to be colored> is the letter or letters that you can chose to be colored.
* You are able to choose between coloring a single letter or a set of letters.
* If the letter is not specified, the whole string is colored.
* The flag must have exactly the same format as above, any other formats returns the following usage message:
```
Usage: go run . [STRING] [OPTION]
EX: go run . something --color=<color>
```
# Usage
```
$ go run . "hello" --color=cyan
$ go run . "hello" --color=cyan "[1:2]"
$ go run . "hello" --color=cyan "[1,2]"
$ go run . "hello" --color=cyan "[:]" 
```