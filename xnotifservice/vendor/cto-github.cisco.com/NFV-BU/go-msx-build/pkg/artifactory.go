package pkg

import (
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
)

func init() {
	AddTarget("publish-binaries", "Publishes Binaries", PublishBinaries)
}

func PublishBinaries(args []string) error {
	// Ensure we can log into artifactory
	if BuildConfig.Binaries.Username == "" || BuildConfig.Binaries.Password == "" {
		return errors.New("Artifactory username or password unset: Please supply " +
			"ARTIFACTORY_USERNAME/ARTIFACTORY_PASSWORD environment variables")
	}

	// Ensure we know where artifactory is
	if BuildConfig.Binaries.Repository == "" {
		return errors.New("Artifactory repository unset: Please supply " +
			"artifactory.repository configuration setting in build.yml")
	}

	// Ensure we know what subpath to use
	if BuildConfig.Msx.DeploymentGroup == "" {
		return errors.New("Deployment Group name not defined: Please supply " +
			"msx.deploymentGroup configuration setting in build.yml")
	}

	logger.Info("Publishing binary artifacts")

	if err := publishAssemblies(); err != nil {
		return err
	}

	if err := publishBinaries(); err != nil {
		return err
	}

	logger.Info("Successfully published binary artifacts.")

	return nil
}

func publishBinaries() error {
	if BuildConfig.Binaries.Installer == "" {
		logger.Warn("Installer artifact directory (artifactory.installer) not configured.  Skipping.")
		return nil
	}

	err := filepath.Walk(BuildConfig.Binaries.Installer, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.Mode().IsRegular() {
			logger.Warn("Skipping non-regular file %q", path)
			return nil
		}

		return publishBinary(path)
	})

	if err != nil {
		return errors.Wrap(err, "Failed to publish binaries")
	}

	return nil
}

func publishBinary(binaryPath string) error {
	uploadUrl := path.Join(BuildConfig.BinariesUrl(), binaryPath)
	return uploadArtifactory(binaryPath, uploadUrl)
}

func publishAssemblies() (err error) {
	assemblies, err := getAllAssemblies()
	if err != nil {
		return err
	}

	for _, assembly := range assemblies {
		err = publishAssembly(assembly)
		if err != nil {
			return errors.Wrapf(err, "Failed to publish assembly %q", assembly.ManifestPrefix)
		}
	}

	return nil
}

func publishAssembly(assembly Assembly) (err error) {
	outputFile := assembly.OutputFile()
	publishUrl := assembly.PublishUrl()
	return uploadArtifactory(outputFile, publishUrl)
}

func uploadArtifactory(sourceFile string, uploadUrl string) (err error) {
	logger.Infof("Uploading %q to %q", sourceFile, uploadUrl)

	var req = new(http.Request)
	req.URL, err = url.Parse(uploadUrl)
	if err != nil {
		return err
	}

	req.Method = http.MethodPut
	req.Header = make(http.Header)
	req.Header.Set("Authorization", BuildConfig.Binaries.Authorization())

	req.Body, err = os.Open(sourceFile)
	if err != nil {
		return err
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrapf(err, "Failed to upload binary %q", filepath.Base(sourceFile))
	}

	return nil
}

