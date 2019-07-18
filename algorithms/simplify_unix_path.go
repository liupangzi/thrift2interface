// +build linux darwin
// +build go1.12

package algorithms

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var logger logrus.Logger

// See https://www.geeksforgeeks.org/simplify-directory-path-unix-like/
func SimplifyUnixDirectoryPath(path string) (simplifiedPath string) {
	stack := []string{""} // dummy head
	for _, dir := range strings.Split(path, "/") {
		if len(dir) == 0 || dir == "." {
			continue
		} else if dir == ".." {
			if len(stack) != 1 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, dir)
		}
	}
	simplifiedPath = strings.Join(stack, string(os.PathSeparator))

	logger.Infof("%s => %s", path, simplifiedPath)
	return
}
