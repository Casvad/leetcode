package main

//
//Write a function in Go that takes a list of (prerequisite, course) pairs and returns the order in which the courses should be taken by a student.
//
//prereqsCourses := [][]string{
//{"Data Structures", "Algorithms"},
//{"Foundations of Computer Science", "Operating Systems"},
//{"Computer Networks", "Computer Architecture"},
//{"Algorithms", "Foundations of Computer Science"},
//{"Computer Architecture", "Data Structures"},
//{"Software Design", "Computer Networks"},
//}
//has context menu

func ResolveOrder(courses [][]string) []string {

	mapCourses := make(map[string]string)
	ans := make([]string, 0)

	for _, course := range courses {
		mapCourses[course[0]] = course[1]
	}

	return ans
}

func main() {

}
