package pkgmgr

import (
	"testing"
	"fmt"
	"github.com/k0kubun/pp"
)

func TestBrewInstalledInfo(t *testing.T) {
	pkgs, err := BrewGetInstalledPackagesDetailed()
	if err != nil {
		fmt.Errorf("Error: %s\n", err.Error())
		pp.Println(err)
		t.FailNow()
	}
	pp.Println(pkgs)
}

func TestBrewAvailablePackages(t *testing.T) {
	pkgs, err := BrewGetAvailablePackagesDetailed()
	if err != nil {
		fmt.Errorf("Error: %s\n", err.Error())
		pp.Println(err)
		t.FailNow()
	}
	pp.Println(pkgs)
}

func TestBrewPackageUpdate(t *testing.T) {
	updates, err := BrewCheckForUpdates()
	if err != nil {
		fmt.Errorf("Error: %s\n", err.Error())
		pp.Println(err)
		t.FailNow()
	}
	pp.Println(updates)
}