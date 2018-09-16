package pkgmgr

import (
	"os/exec"
	"regexp"
	"strings"
)
const (
	YumPackageRegex = `(\S+)\s+(\d\S+)\s+(\S+)`
	)

type YumPackageManager struct {

}

func (YumPackageManager) GetInstalledPackages() ([]PackageInstallation, error) {
	regex, _ := regexp.Compile(YumPackageRegex)

	// dpkg command
	cmdName := DPKG_CMD_NAME
	cmdArgs := []string{"list", "installed"}

	//Run the command
	cmdOutput, err := exec.Command(cmdName, cmdArgs...).Output()
	if err != nil {
		return nil, err
	}

	// Only output the commands stdout
	var packages []PackageInstallation
	res := regex.FindAllStringSubmatch(string(cmdOutput), -1)
	for _, line := range res {
		iPackage := PackageInstallation{
			PackageName: line[2],
			PackageVersion: line[3],
			PackageManager: PACKAGE_MANAGER_DPKG,
		}
		packages = append(packages, iPackage)
	}
	return packages, nil
}

func (YumPackageManager) GetPackageUpdates() ([]PackageUpdate, error) {
	//Result Object
	var updates []PackageUpdate
	//Compile Regex
	regex, _ := regexp.Compile(YumPackageRegex)

	// dpkg command
	cmdName := DPKG_CMD_NAME
	cmdArgs := []string{"list", "updates"}

	//Run the command
	cmdOutput, err := exec.Command(cmdName, cmdArgs...).Output()
	if err != nil {
		return nil, err
	}

	//Split output into lines
	for _, line := range strings.Split(string(cmdOutput), "\n") {
		parts := regex.FindStringSubmatch(line)
		if len(parts) > 0 {
			update := PackageUpdate{
				PackageName:       parts[1],
				NewestVersion:     parts[2],
				PackageManager:    PACKAGE_MANAGER_BREW,
			}
			updates = append(updates, update)
		}
	}
	return updates, nil
}

