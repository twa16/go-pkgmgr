package pkgmgr

import "testing"

func TestDpkgGetInstalled(t *testing.T) {
	pkgs, err := DpkgGetInstalledPackages()
	if err != nil {
		t.Errorf("Error Getting Installed Packages: %s\n", err.Error())
		t.FailNow()
	}
	t.Logf("Got %d Packages\n", len(pkgs))
}
