package pkgmgr

import (
	"regexp"
	"os/exec"
)

const DPKG_CMD_NAME = "dpkg"

func DpkgGetInstalledPackages() ([]PackageInstallation, error) {
	//Get regex ready
	installedPackageRegexString := `(?m)(^[iuprh]\S)\s+(\S+)\s+(\S+)\s+(\S+)\s+(.+)`
	regex, _ := regexp.Compile(installedPackageRegexString)

	// dpkg command
	cmdName := DPKG_CMD_NAME
	cmdArgs := []string{"--list"}

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
		//fmt.Printf("[%d]: Got Install Package: %s, Version:%s\n", i, iPackage.Name, iPackage.Version)
	}
	return packages, nil

}