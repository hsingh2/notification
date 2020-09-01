package pkg

import (
	"archive/tar"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const manifestFileName = "manifest.json"

func init() {
	AddTarget("build-assemblies", "Builds Assemblies", BuildAssemblies)
}

func BuildAssemblies(args []string) error {
	assemblies, err := getAllAssemblies()
	if err != nil {
		return nil
	}

	if len(assemblies) == 0 {
		logger.Warn("No assemblies defined.")
		return nil
	}

	// Create an assembly for each sub-folder containing a `manifest.json`
	for _, assembly := range assemblies {
		if err = installAssembly(assembly); err != nil {
			return errors.Wrapf(err, "Failed to install assembly %q", assembly.ManifestPrefix)
		}
	}

	return nil
}

func getAllAssemblies() ([]Assembly, error) {
	rootAssemblies, err := getRootAssemblies()
	if err != nil {
		return nil, err
	}

	return append(rootAssemblies, BuildConfig.Assemblies.Custom...), nil
}

func getRootAssemblies() ([]Assembly, error) {
	// Find all sub-directories of assemblies.root
	folders, err := getFolders(BuildConfig.Assemblies.Root)
	if err != nil {
		return nil, err
	}

	var results []Assembly
	for _, folder := range folders {
		_, err := os.Stat(filepath.Join(BuildConfig.Assemblies.Root, folder, manifestFileName))
		if err != nil {
			logger.Warnf("Manifest not found in folder %q.  Skipping assembly.", folder)
			continue
		}

		results = append(results, getFolderAssembly(folder))
	}
	return results, nil
}

func getFolderAssembly(folder string) Assembly {
	return Assembly{
		Path:           filepath.Join(BuildConfig.Assemblies.Root, folder),
		PathPrefix:     "",
		ManifestPrefix: folder + "-templates",
		ManifestKey:    folder + "-templates",
		Includes:       []string{"**/*"},
		Excludes:       nil,
	}
}

func installAssembly(assembly Assembly) error {
	logger.Infof("Creating assembly %q", assembly.OutputFile())

	// Create the output directory
	err := os.MkdirAll(BuildConfig.AssemblyPath(), 0755)
	if err != nil {
		return err
	}

	// Generate the tarball into the output directory
	fw, err := os.Create(assembly.OutputFile())
	if err != nil {
		return err
	}
	defer fw.Close()

	tw := tar.NewWriter(fw)
	err = filepath.Walk(assembly.Path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			logger.WithError(err).Errorf("Failed to process entry at path %q", p)
			return nil
		}

		if info.IsDir() {
			return nil
		} else if !info.Mode().IsRegular() {
			logger.Warnf("Ignoring non-regular file at path %q", p)
			return nil
		}

		subpath := strings.TrimPrefix(p, assembly.Path+"/")
		if assembly.PathPrefix != "" {
			subpath = path.Join(assembly.PathPrefix, subpath)
		}
		logger.Infof("Adding %q", subpath)

		hdr := &tar.Header{
			Name:    subpath,
			Size:    info.Size(),
			Mode:    int64(info.Mode().Perm()),
			ModTime: info.ModTime(),
		}

		if err := tw.WriteHeader(hdr); err != nil {
			return errors.Wrapf(err, "Failed to write tar file header for %q", p)
		}

		data, err := ioutil.ReadFile(p)
		if err != nil {
			return errors.Wrapf(err, "Failed to read file data from %q", p)
		}

		if _, err := tw.Write(data); err != nil {
			return errors.Wrapf(err, "Failed to write file body for %q", p)
		}

		return nil
	})

	if err != nil {
		return err
	}

	err = tw.Close()
	if err != nil {
		return err
	}

	logger.Infof("Successfully created assembly %q", assembly.OutputFile())
	return nil
}

func getFolders(root string) ([]string, error) {
	fileInfos, err := ioutil.ReadDir(root)
	if err != nil {
		return nil, err
	}

	var folders []string
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			continue
		}
		folders = append(folders, fileInfo.Name())
	}

	return folders, nil
}
