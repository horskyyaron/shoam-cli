# To-dos:
* fix global dirs, make it extensible


# Commands:
* info [course_name]/[course_code]   - gives information regarding a course.
* calc:
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
