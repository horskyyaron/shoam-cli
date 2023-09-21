# CLI tool for the courses catalog.

In BIU there is a course catalog called "shoam".
Since I'm tired of going online each time I want to figure out something about a specific course,
or even what is the course with the code 89220. (algo 1 of course, it cannot be forgotten).

The process wasn't fast enough for me, thus "shoam" came to be!

* [Installation](installation)
* [Environment variables](#environment-variables)
* [Features](#features)
* [Requirements](#requirements)
* [Usage](#usage)
    * [help](#help)
    * [db](#db)
    * [search](#search)
    * [info](#info)
    * [calc](#calc)
* [note on the data fetching](#note-on-the-data-fetching)

## Installation

clone this repo.
```
git clone https://github.com/horskyyaron/shoam
```

## Environment variables

For shoam to work, you will need to export an env variable: SHOAM_DIR with the value of the path
where the this repo is cloned to.

full set up:

```
git clone https://github.com/horskyyaron/shoam.git
cd ./shoam
export SHOAM_DIR=$(pwd) 
```

Add SHOAM_DIR to your zshrc/bashrc for shoam to persist.

## Features

* Search for a course by code, or by name. (can also use partial code/name)
* Get course info. (tests, professor, credit points etc.. as shown on the course page on the shoam system)
* calculate credits (by Nakaz! not hours)

## Requirements

Mostly linux command line utilities, and some Go packages.

* curl
* grep
* awk
* sed
* pup https://github.com/ericchiang/pup/tree/master

## Usage

### help

You can always use the built in helper to read about the different options.
```
$ shoam -h
BIU courses cli tool for the savvy student

Usage:
  shoam [command]

Available Commands:
  calc        calculates total credit points for a list of courses
  completion  Generate the autocompletion script for the specified shell
  db          creates and updates the courses db based on the Shoam system of BIU
  help        Help about any command
  info        Get information for the input course
  search      Searching for a course using a string pattern

Flags:
  -h, --help   help for shoam

Use "shoam [command] --help" for more information about a command.
```


### db

```
shoam db create
```

this command will fetch the courses information from the Shoam system,
it will store it locally.

* (!) When you clone the repo, **the DB is already set in the repo**, if you want to re-fetch
the data you can use this command to be sure the db is up to date.

Explanation on how db create works is at the end of the readme. [click here](#note-on-the-data-fetching)
     

### search

* find all courses that have "calc" in their names.
```
$ shoam search "calc"
89118     Introduction to Calculus I
89133     Calculus II
89218     Introduction to Calculus II
```

* Can't remember if algo 1 is 89220 or 89226? search using only a part of the code. 

```
$ shoam search 8922
89220     Algorithms 1
892226    Computability and Complexity

```

search is case insensitive.

## info

get information regarding a given course.

```
$ shoam info 89230
https://shoham.biu.ac.il/BiuCoursesViewer/ENCourseDetails.aspx?lid=764650
Course Number
89230-01
Course Name
Computer Architecture
Department
Department of Computer Science
Faculty
Exact Sciences 
Meeting Type
Lecture
Lecturer&#39;s Name
Dr. Marina Kogan-Sadetsky
Course Academic Hours
Semester A - 3.00
Credit(s)
1.50
Day
 Mon
Time
15:00 - 18:00 
Exam dates



First Exam
15/02/2024
16:00


Second Exam
14/03/2024
16:00


Syllabus
See Syllabus
```

* working on better output for this command :).

## calc 

calculate total credit points using one of the following options:

1. course(s) code as arguments

```
$ shoam calc 89230 // one course
total points (1 courses): 2.50

$ shoam calc 89230 89220 // multiple courses
total points (2 courses): 5.00

$ shoam calc 89230 89220 -v // verbose output
89230, Computer Architecture
---------------------------------------
1.50, Lecture points
1.00, Recitation points
2.50, total

89220, Algorithms 1
---------------------------------------
1.50, Lecture points
1.00, Recitation points
2.50, total

total points (2 courses): 5.00
```

2. Using the -f flag for using a course list file.

```
/path/to/somefile
------
89230
89220


$ shoam calc -f /path/to/somefile 
total points (2 courses): 5.00

```

3. Using pipe and stdin.
```
$ echo 89230 | shoam calc
total points (1 courses): 2.50


$ shoam calc
89230
^d //signaling end of stdin input.
total points (1 courses): 2.50
```

### Note on the data fetching
Each course page in the shoam system is of the form:
```
https://shoham.biu.ac.il/BiuCoursesViewer/CourseDetails.aspx?lid=XXXXXX
```

e.g. https://shoham.biu.ac.il/BiuCoursesViewer/CourseDetails.aspx?lid=764564 (Intro to computers)

The current way of getting the "lids" part for every CS department course, is by going to the Shoam system,
enter the current year + CS department, there will be around 20~ pages of lists of courses, each course entry with a link to
the course page.

The links are what we are interested in.
So I went to each of the 20 pages, copied the html of the page, from there created a 
file that includes all of the "lids". (the file inside "links" folder)

when using db create, the program will curl each of these lids and get the information.

