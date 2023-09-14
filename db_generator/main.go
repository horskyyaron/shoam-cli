package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"sync"
)

const (
	THREADS   = 1000
	ADDRESSES = 50
)

func fetchData(url string, base int) {
	fmt.Println(base)
	for i := base; i < base+ADDRESSES; i++ {
		url = "https://shoham.biu.ac.il/BiuCoursesViewer/ENCourseDetails.aspx?lid=" + strconv.Itoa(
			i,
		)
		fmt.Printf("trying %s\n", url)
		cmd := exec.Command("/home/yaron/projects/shoam_bash/course_info", url)
		err := cmd.Run()
		if err == nil {
			fmt.Println("course added")
		}
	}
}

// this script will scan the Shoam system for CS courses.
// it will save the data locally in a folder called "db" (see 'course_info' program in the parent folder)
// each course file is of the following form:
// ------- [course #]-[group #]-[lid #] (lid is explained below)
// since the urls do not corresponeds to the course number or group,
// this script will try 50,000 different urls, and will extract the data for every CS course it finds.
// the url is of the form:
// ------ https://shoham.biu.ac.il/BiuCoursesViewer/ENCourseDetails.aspx?lid=[...]
// this script can scan more addresses by changing the number of ADDRESSES in the global variable above.
// can also change the # of threads running for efficiency.
// results for 8RAM memory and Intel CPU i5-8250U - about 25 minutes.

func main() {
	var wg sync.WaitGroup

	base := 760000

	wg.Add(THREADS)
	url := "https://shoham.biu.ac.il/BiuCoursesViewer/ENCourseDetails.aspx?lid=" + strconv.Itoa(
		base,
	)

	for i := 0; i < THREADS; i++ {
		go fetchData(url, base)
		base = base + ADDRESSES
	}

	wg.Wait()
}
