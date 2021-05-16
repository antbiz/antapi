package execute

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/gogf/gf/os/gcmd"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/util/grand"
)

// Help .
func Help() {
	fmt.Println(`
USAGE
    bench execute PATH FUNC
ARGUMENT
	PATH execute path.
	FUNC function name.
    OPTION  
OPTION
EXAMPLES
	go run cmd/bench/bench.go execute patches/system InitAdminAccount
DESCRIPTION
    The "execute" command is used for execute patch`)
}

// Run .
func Run() {
	filePath, funcName := gcmd.GetArg(2), gcmd.GetArg(3)
	if filePath == "" {
		panic("execute file path cannot empty")
	}
	if funcName == "" {
		panic("execute function name cannot empty")
	}
	fileAbs := gfile.Abs(filePath)

	pkgName := extractPackage(fileAbs)
	if pkgName == "" {
		return
	}

	randStr := grand.S(10)
	genFileName := randStr + "_test.go"
	fileContent := fmt.Sprintf(`package %v
import (
	"testing"
	"github.com/gogf/gf/frame/g"
)
func Test_%v(t *testing.T) {
	g.Cfg().SetPath("%s")
	%v()
}`, pkgName, randStr, gfile.Pwd(), funcName,
	)

	genPath := gfile.Join(fileAbs, genFileName)
	defer gfile.Remove(genPath)

	if err := gfile.PutContents(genPath, fileContent); err != nil {
		panic(err)
	}

	if output, err := exec.Command("goimports", "-w=true", fileAbs).CombinedOutput(); err != nil {
		panic(err)
	} else {
		fmt.Print(string(output))
	}

	output, err := exec.Command("go", "test", fileAbs, "-v", "--run", fmt.Sprintf("^Test_%s$", randStr)).CombinedOutput()
	stdout := string(output)
	stdoutLines := strings.Split(stdout, "\n")
	if err != nil {
		stdout = strings.Join(stdoutLines[:len(stdoutLines)-2], "\n")
		fmt.Println(stdout)
	} else {
		stdout = strings.Join(stdoutLines[1:len(stdoutLines)-4], "\n")
		fmt.Println(stdout)
	}
}

func extractPackage(path string) string {
	dirFiles, err := gfile.ScanDirFile(path, "*.go")
	if err != nil {
		panic(err)
	}
	if len(dirFiles) == 0 {
		return ""
	}

	filePath := dirFiles[0]
	regex, err := regexp.Compile("^ *package (.*)$")
	if err != nil {
		panic(err)
	}

	fh, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	f := bufio.NewReader(fh)

	defer fh.Close()

	buf := make([]byte, 1024)
	for {
		buf, _, err = f.ReadLine()
		if err != nil {
			panic(err)
		}
		s := string(buf)
		if regex.MatchString(s) {
			return strings.Split(string(buf), " ")[1]
		}
	}
}
