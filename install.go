package main

import "flag"
import "os"
import "path/filepath"

func (gpkg *Gpkg) install() {
	logger := gpkg.logger
	gvm := gpkg.gvm
	pkgname := readCommand()
	if pkgname == "" {
		logger.Fatal("Please specify package name")
	} else if pkgname == "." {
		wd, _ := os.Getwd()
		p := gvm.NewPackage(filepath.Base(wd), "")
		p.source = wd
		p.Install(gpkg.tmpdir)
		return
	}

	version := flag.String("version", "", "Package version to install")
	flag.Parse()
	if *version == "" {
		p := gvm.NewPackage(pkgname, "")
		p.Install(gpkg.tmpdir)
	} else {
		p := gvm.NewPackage(pkgname, *version)
		p.Install(gpkg.tmpdir)
	}
}