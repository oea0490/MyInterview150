package Graph

var (
	courseEdges [][]int // 表示课程i的后继课程
	inDegree    []int   // 表示课程i的入度
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 初始化courseEdges和inDegree
	courseEdges = make([][]int, numCourses)
	inDegree = make([]int, numCourses)
	for _, p := range prerequisites {
		courseEdges[p[1]] = append(courseEdges[p[1]], p[0])
		inDegree[p[0]]++
	}
	// count表示已经学习的课程数
	count := 0
	// 初始化队列,将入度为0的课程加入队列
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	// 学习课程
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		count++
		// 学习完cur课程后,更新cur课程的后继课程的入度,将入度为0的课程加入队列
		for _, c := range courseEdges[cur] {
			inDegree[c]--
			if inDegree[c] == 0 {
				queue = append(queue, c)
			}
		}
	}
	return count == numCourses
}
