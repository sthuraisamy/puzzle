# puzzle
## One
### Question
You have 100 doors in a row that are all initially closed. you make 100 passes by the doors
starting with the first door every time. the first time through you visit every door and toggle the
door (if the door is closed, you open it, if its open, you close it). the second time you only visit
every 2nd door (door #2, #4, #6). the third time, every 3rd door (door #3, #6, #9), etc, until you
only visit the 100th door. What state are the doors in after the last pass? Which are open which
are closed?

Following sample command is for MacOS. The windows binary (puzzle_win) and linux binary (puzzle_lnx) are included in bin directory. There are two parameters.

number : number of doors.

state : initial door state. open=0 and close=1.
```
bin/puzzle_osx door --number 100 --state 1
```