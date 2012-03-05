package main

import "os"
import "io/ioutil"
import "path/filepath"
import "strings"
import "exec"

type Gvm struct {
	root string
	go_name string
	go_root string
	pkgset_name string
	pkgset_root string
	sources []string
	logger *Logger
}

func NewGvm(logger *Logger) *Gvm {
	gvm := &Gvm{logger: logger}
	gvm.root = os.Getenv("GVM_ROOT")
	gvm.go_name = os.Getenv("gvm_go_name")
	gvm.go_root = filepath.Join(gvm.root, "gos", gvm.go_name)
	gvm.pkgset_name = os.Getenv("gvm_pkgset_name")
	gvm.pkgset_root = filepath.Join(gvm.root, "pkgsets", gvm.go_name, gvm.pkgset_name)

	data, err := ioutil.ReadFile(filepath.Join(gvm.root, "config", "sources"))
	if err != nil {
		panic(err)
	}

	gvm.sources = strings.Split(string(data), "\n")
	return gvm
}

func (gvm *Gvm) NewPackage(name string, version string) *Package {
	p := &Package{
		gvm: gvm,
		logger: gvm.logger,
		name: name,
		version: version,
	}
	p.root = filepath.Join(p.gvm.pkgset_root, "pkg.gvm", p.name)
	return p
}

func (gvm *Gvm) InstallPackageByVersion(name string, version string) *Package {
	p := gvm.NewPackage(name, version)
	p.Install()
	return p
}

func (gvm *Gvm) InstallPackage(name string) *Package {
	gvm.logger.Trace("name", name)
	p := gvm.NewPackage(name, "")
	p.Install()
	return p
}

func (gvm *Gvm) FindPackageByVersion(name string, version string) *Package {
	gvm.logger.Trace("name", name)
	gvm.logger.Trace("version", version)
	_, err := os.Open(filepath.Join(gvm.pkgset_root, "pkg.gvm", name, version, "pkg"))
	if err == nil {
		p := gvm.NewPackage(name, version)
		return p
	}
	return nil
}

func (gvm *Gvm) FindPackage(name string) *Package {
	gvm.logger.Trace("name", name)
	_, err := os.Open(filepath.Join(gvm.pkgset_root, "pkg.gvm", name))
	if err == nil {
		verions, _ := ioutil.ReadDir(filepath.Join(gvm.pkgset_root, "pkg.gvm", name))
		for _, version := range verions {
			p := gvm.NewPackage(name, version.Name)
			return p
		}
		os.Exit(1)
	}
	return nil
}

func (gvm *Gvm) PackageList() (pkglist[] *Package) {
	out, err := exec.Command("ls", filepath.Join(gvm.pkgset_root, "pkg.gvm")).CombinedOutput()
	if err == nil {
		pkgs := strings.Split(string(out), "\n")
		pkgs = pkgs[0:len(pkgs)-1]
		pkglist = make([]*Package, len(pkgs))
		for n, pkg := range pkgs {
			pkglist[n] = gvm.NewPackage(pkg, "")
		}
		return pkglist
	}
	return []*Package{}
}
