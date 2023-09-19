# To-dos:
* fix global dirs, make it extensible


# Commands:
* calc:
    * check if a course has a rec.
    * fetch points from given course + group 
    * given a file/stdin of a list of courses -> return total points.
    * given a course code -> course points
* tests:
    * give all tests
    * give all tests in a spesific semester
    * given a course code -> course test dates

## e.g.
```
shoam info 89220
shoam info 89220 --group=05 -> course info for spesific group
shoam calc [file] or cat file | shoam calc -> spits out total points
shoam test -> show tests
shoam test --course=89230 -> show tests for course 89230
```
