package ohgi

import (
	"fmt"
	"github.com/tcnksm/go-latest"
)

func Version(version string) string {
	var result []byte
	result = append(result, fmt.Sprintf("ohgi version %s\n", version)...)

	fixFunc := latest.DeleteFrontV()
	githubTag := &latest.GithubTag{
		Owner:             "hico-horiuchi",
		Repository:        "ohgi",
		FixVersionStrFunc: fixFunc,
	}

	res, _ := latest.Check(githubTag, fixFunc(version))
	if res.Outdated {
		result = append(result, fmt.Sprintf("Latest version of ohgi is %s, please update it\n", res.Current)...)
	}

	return string(result)
}
