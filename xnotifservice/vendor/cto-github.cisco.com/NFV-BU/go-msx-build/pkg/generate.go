package pkg

import (
	"cto-github.cisco.com/NFV-BU/go-msx/exec"
	"cto-github.cisco.com/NFV-BU/go-msx/fs"
	"github.com/shurcooL/vfsgen"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

var commandRegexp = regexp.MustCompile(`[^\s"']+|"([^"]*)"|'([^']*)'`)

func init() {
	AddTarget("generate", "Generate code", GenerateCode)
}

func GenerateCode(args []string) error {
	for _, p := range BuildConfig.Generate {
		var err error
		logger.Infof("Generating path '%s'", p.Path)
		if p.VfsGen != nil {
			err = generateCodePathVfs(p)
		} else if p.Command != "" {
			err = generateCodePathCommand(p)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func generateCodePathCommand(p Generate) error {
	if p.Command == "" {
		p.Command = "go generate"
	}

	args := commandRegexp.FindAllString(p.Command, -1)
	return exec.ExecutePipes(
		exec.WithDir(p.Path,
			exec.Exec(args[0], args[1:])))
}

func generateCodePathVfs(p Generate) (err error) {
	root, err := getGenerateRootDir(p)
	if err != nil {
		return err
	}

	targetFs, err := fs.NewGlobFileSystem(http.Dir(root), p.VfsGen.Includes, p.VfsGen.Excludes)
	if err != nil {
		return err
	}

	projectRoot, err := os.Getwd()
	if err != nil {
		return err
	}

	fileName := path.Join(projectRoot, p.Path, p.VfsGen.Filename)

	packageName := path.Base(p.Path)
	return vfsgen.Generate(targetFs, vfsgen.Options{
		Filename:     fileName,
		PackageName:  packageName,
		VariableName: p.VfsGen.VariableName,
	})
}

func getGenerateRootDir(p Generate) (string, error) {
	abs, err := filepath.Abs(p.Path)
	if err != nil {
		return "", err
	}

	if p.VfsGen.Root != "" {
		abs, err = filepath.Abs(filepath.Join(abs, p.VfsGen.Root))
		if err != nil {
			return "", err
		}
	}

	return abs, nil
}
