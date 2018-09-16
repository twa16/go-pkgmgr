package pkgmgr

type PackageManager interface {
	GetInstalledPackages() ([]PackageInstallation, error)
	GetPackageUpdates()    ([]PackageUpdate, error)
}

//PackageInstallation Represents an installed package
type PackageInstallation struct {
	PackageName    string //Name of the package
	PackageVersion string //Version of the package
	PackageManager string //Package manager that reported this package
}

//PackageListing Represents a package version that available for installation. This may be the same as the installed version.
type PackageListing struct {
	PackageName        string   //Name of the package
	PackageDescription string   //Description of package
	PackageVersion     string   //Version of the package
	PackageManager     string   //Package manager that reported this package
	Repository         string   //Repository that contains this package. If applicable.
	Group              string   //Group that this package belongs to
	Dependencies       []string //Packages required for installation
}

type PackageUpdate struct {
	PackageName       string
	InstalledVersions []string
	NewestVersion     string
	PackageManager    string
}

//BrewInfoJSONResponse Brew-specific JSON struct that used to process the output of `brew list --json=1`
type BrewInfoJSONResponse []struct {
	Name     string      `json:"name"`
	FullName string      `json:"full_name"`
	Desc     string      `json:"desc"`
	Homepage string      `json:"homepage"`
	Oldname  interface{} `json:"oldname"`
	Aliases  []string    `json:"aliases"`
	Versions struct {
		Stable string      `json:"stable"`
		Bottle bool        `json:"bottle"`
		Devel  interface{} `json:"devel"`
		Head   string      `json:"head"`
	} `json:"versions"`
	Revision      int `json:"revision"`
	VersionScheme int `json:"version_scheme"`
	Installed     []struct {
		Version               string        `json:"version"`
		UsedOptions           []interface{} `json:"used_options"`
		BuiltAsBottle         bool          `json:"built_as_bottle"`
		PouredFromBottle      bool          `json:"poured_from_bottle"`
		RuntimeDependencies   interface{}   `json:"runtime_dependencies"`
		InstalledAsDependency interface{}   `json:"installed_as_dependency"`
		InstalledOnRequest    interface{}   `json:"installed_on_request"`
	} `json:"installed"`
	LinkedKeg               string        `json:"linked_keg"`
	Pinned                  bool          `json:"pinned"`
	Outdated                bool          `json:"outdated"`
	KegOnly                 interface{}   `json:"keg_only"`
	Dependencies            []string      `json:"dependencies"`
	RecommendedDependencies []interface{} `json:"recommended_dependencies"`
	OptionalDependencies    []interface{} `json:"optional_dependencies"`
	BuildDependencies       []string      `json:"build_dependencies"`
	ConflictsWith           []interface{} `json:"conflicts_with"`
	Caveats                 interface{}   `json:"caveats"`
	Requirements            []struct {
		Name           string      `json:"name"`
		DefaultFormula string      `json:"default_formula"`
		Cask           interface{} `json:"cask"`
		Download       interface{} `json:"download"`
	} `json:"requirements"`
	Options []interface{} `json:"options"`
	Bottle  struct {
		Stable struct {
			Rebuild int    `json:"rebuild"`
			Cellar  string `json:"cellar"`
			Prefix  string `json:"prefix"`
			RootURL string `json:"root_url"`
			Files   struct {
				Sierra struct {
					URL    string `json:"url"`
					Sha256 string `json:"sha256"`
				} `json:"sierra"`
				ElCapitan struct {
					URL    string `json:"url"`
					Sha256 string `json:"sha256"`
				} `json:"el_capitan"`
				Yosemite struct {
					URL    string `json:"url"`
					Sha256 string `json:"sha256"`
				} `json:"yosemite"`
			} `json:"files"`
		} `json:"stable"`
	} `json:"bottle"`
}
