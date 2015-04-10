package ohgi

import (
	"fmt"
	"github.com/tcnksm/go-latest"
	"time"
)

func verCheck(version string) <-chan *latest.CheckResponse {
	verCheckCh := make(chan *latest.CheckResponse)

	go func() {
		fixFunc := latest.DeleteFrontV()
		githubTag := &latest.GithubTag{
			Owner:             "hico-horiuchi",
			Repository:        "ohgi",
			FixVersionStrFunc: fixFunc,
		}

		res, _ := latest.Check(githubTag, fixFunc(version))
		verCheckCh <- res
	}()

	return verCheckCh
}

func Version(version string) string {
	var result []byte
	result = append(result, fmt.Sprintf("ohgi version %s\n", version)...)
	verCheckCh := verCheck(version)

	for {
		select {
		case res := <-verCheckCh:
			if res != nil && res.Outdated {
				result = append(result, fmt.Sprintf("Latest version of ohgi is %s, please update it\n", res.Current)...)
			}
			return string(result)
		case <-time.After(defaultTimeout):
			return string(result)
		}
	}
}
