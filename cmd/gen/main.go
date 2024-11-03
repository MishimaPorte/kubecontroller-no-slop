package main

import (
	"kek/codegen/clientset"
	"kek/codegen/deepcopy"
	"kek/codegen/informer"
	"kek/codegen/lister"

	clargs "k8s.io/code-generator/cmd/client-gen/args"
	"k8s.io/code-generator/cmd/client-gen/types"
	dcargs "k8s.io/code-generator/cmd/deepcopy-gen/args"
	infargs "k8s.io/code-generator/cmd/informer-gen/args"
	largs "k8s.io/code-generator/cmd/lister-gen/args"
)

func must(err error) {
	if err != nil {
		panic(err.Error())
	}
}
func main() {
	must(deepcopy.Generate(&dcargs.Args{
		OutputFile: "zz_generated.deepcopy.go",
	},
		"kek/apis/modulepodset/v1alpha1",
	))
	must(clientset.Generate(&clargs.Args{
		OutputDir: "generated",
		OutputPkg: "kek/generated",
		Groups: []types.GroupVersions{{
			PackageName: "modulepodsets",
			Group:       "sp.aps",
			Versions: []types.PackageVersion{{
				Version: "v1alpha1",
				Package: "kek/apis/modulepodset/v1alpha1",
			}},
		}},
		ClientsetName:    "clientset",
		ClientsetAPIPath: "/apis",
		FakeClient:       true,
	}))

	must(lister.Generate(&largs.Args{
		OutputDir:    "generated/listers",
		OutputPkg:    "kek/generated/listers",
		GoHeaderFile: "",
	}, "kek/apis/modulepodset/v1alpha1"))

	must(informer.Generate(&infargs.Args{
		OutputDir:                 "generated/informers",
		OutputPkg:                 "kek/generated/informers",
		VersionedClientSetPackage: "kek/generated/clientset",
		ListersPackage:            "kek/generated/listers",
		SingleDirectory:           false,
	}, "kek/apis/modulepodset/v1alpha1"))
}
