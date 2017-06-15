package pkgmgr

import (
	"strings"
	"os/exec"
	"encoding/json"
	"regexp"
)

const PACKAGE_MANAGER_BREW = "brew"

//BrewGetInstalledBrewPackages Calls brew to get installed packages and their versions
func BrewGetInstalledPackages() ([]PackageInstallation, error){
	//Run `brew list --full-name --versions`
	res, err := exec.Command("brew", "list", "--full-name", "--versions").Output()
	if err != nil {
		return nil, err
	}

	//Variable to hold packages
	var packages []PackageInstallation
	//Split response from brew into lines
	packageLines := strings.Split(string(res), "\n")
	//Process the lines
	for _, packageLine := range packageLines {
		//Split the line to get package and versions
		lineParts := strings.Split(packageLine, " ")
		//Iterate through versions
		for _, packageVersion := range lineParts[1:] {
			//Create a package listing
			packageInstallation := PackageInstallation{
				PackageName: lineParts[0],
				PackageVersion: packageVersion,
			}
			//Add our new package
			packages = append(packages, packageInstallation)
		}

	}
	//Return the result
	return packages, nil
}

//BrewGetAvailablePackages Calls Brew to get available local packages. This does not pull version information.
func BrewGetAvailablePackages() ([]string, error){
	//Run `brew list --full-name --versions`
	res, err := exec.Command("brew", "search").Output()
	if err != nil {
		return nil, err
	}
	packageLines := strings.Split(string(res), "\n")
	return packageLines, nil
}

func BrewGetPackageInformation(packageName string) (*PackageListing, error) {
	//Run `brew list --full-name --versions`
	_, err := exec.Command("brew", "info", packageName).Output()
	if err != nil {
		return nil, err
	}

	//Split response from brew into lines
	//packageLines := strings.Split(string(res), "\n")

	//Result Object
	return nil, nil
}

func BrewGetInstalledPackagesDetailed() ([]PackageInstallation, error) {
	//Run `brew list --full-name --versions`
	res, err := exec.Command("brew", "info", "--json=v1", "--installed").Output()
	if err != nil {
		return nil, err
	}

	//Get output
	var output BrewInfoJSONResponse
	err = json.Unmarshal(res, &output)
	if err != nil {
		return nil, err
	}

	//Create variable to store package installations
	var installations []PackageInstallation

	for _, pkg := range output {
		pkgInstall := PackageInstallation{
			PackageName: pkg.Name,
			PackageVersion: pkg.Versions.Stable,
			PackageManager: PACKAGE_MANAGER_BREW,
		}
		installations = append(installations, pkgInstall)
	}

	return installations, nil
}

func BrewGetAvailablePackagesDetailed() ([]PackageListing, error) {
	//Run `brew list --full-name --versions`
	res, err := exec.Command("brew", "info", "--json=v1", "--all").Output()
	if err != nil {
		return nil, err
	}

	//Get output
	var output BrewInfoJSONResponse
	err = json.Unmarshal(res, &output)
	if err != nil {
		return nil, err
	}

	//Create variable to store package listings
	var listings []PackageListing

	for _, pkg := range output {
		pkgListing := PackageListing{
			PackageName: pkg.Name,
			PackageDescription: pkg.Desc,
			PackageVersion: pkg.Versions.Stable,
			PackageManager: PACKAGE_MANAGER_BREW,
			Repository: pkg.FullName,
			Group: "",
			Dependencies: pkg.Dependencies,
		}
		listings = append(listings, pkgListing)
	}

	return listings, nil
}

func BrewCheckForUpdates() ([]PackageUpdate, error) {
	//Result Object
	var updates []PackageUpdate

	//Execute Command
	res, err := exec.Command("brew", "outdated", "--verbose").Output()
	if err != nil {
		return nil, err
	}
	//Put together regex
	regexPattern := regexp.MustCompile(`^(\S+)\s+(\(.+\))\s+<\s+(.+)`)
	existingVersionRegex := regexp.MustCompile(`([^,\s()]+)`)

	//Split output into lines
	for _, line := range strings.Split(string(res), "\n") {
		parts := regexPattern.FindStringSubmatch(line)
		if len(parts) > 0 {
			versions := existingVersionRegex.FindAllString(parts[2], -1)
			update := PackageUpdate{
				PackageName:       parts[1],
				InstalledVersions: versions,
				NewestVersion:     parts[3],
				PackageManager:    PACKAGE_MANAGER_BREW,
			}
			updates = append(updates, update)
		}
	}
	return updates, nil
}