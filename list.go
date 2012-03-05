package main

import "flag"
import "io/ioutil"
import "path/filepath"
import "strings"

func (gpkg *Gpkg) listPackageVersions(pkg *Package) {
	versions := pkg.GetVersions()
	version_str := ""
	for n, version := range versions {
		version_str += version
		if n < len(versions) - 1 {
			version_str += ", "
		}
	}
	gpkg.logger.Info(pkg.name, "(" + version_str + ")")
}

func (gpkg *Gpkg) listPackage(pkg *Package) {
	logger := gpkg.logger
	logger.Message("Package Info:", pkg.name)
	logger.Info("  version:", pkg.version)
	data, err := ioutil.ReadFile(filepath.Join(pkg.root, pkg.version, "manifest"))
	if err == nil {
		logger.Info("  deps:")
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			if len(line) > 3 && line[0:3] == "pkg" {
				logger.Info("   ", line[4:])
			}
		}
	}
}

func (gpkg *Gpkg) list() {
	logger := gpkg.logger
	gvm := gpkg.gvm
	pkgname := readCommand()
	if pkgname != "" {
		var pkg *Package
		pkg = gvm.FindPackage(pkgname)
		if pkg != nil {
			version := flag.String("version", "", "Package version to install")
			flag.Parse()
			if *version != "" {
				pkg = gvm.FindPackageByVersion(pkgname, *version)
				if pkg != nil {
					gpkg.listPackage(pkg)
				} else {
					logger.Fatal("Package version not found")
				}
				return
			} else {
				logger.Message("\ngpkg list", pkg.name, "in", gvm.go_name + "@" + gvm.pkgset_name, "\n")
				gpkg.listPackageVersions(pkg)
			}
		} else {
			logger.Fatal("Package not found")
		}
	} else {
		logger.Message("\ngpkg list", gvm.go_name + "@" + gvm.pkgset_name, "\n")
		pkgs := gvm.PackageList()
		for _, pkg := range pkgs {
			gpkg.listPackageVersions(pkg)
		}
	}
	logger.Info()
}
