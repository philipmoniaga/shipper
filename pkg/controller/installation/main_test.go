package installation

import (
	"os"
	"testing"

	"k8s.io/helm/pkg/repo/repotest"
)

var (
	repoUrl string
	repoPwd string
)

func TestMain(m *testing.M) {
	srv, hh, err := repotest.NewTempServer("testdata/chart-cache/localhost/*.tgz")
	if err != nil {
		panic(err.Error())
	}
	repoUrl = srv.URL()
	repoPwd = hh.String()
	status := m.Run()
	srv.Stop()
	os.RemoveAll(hh.String())
	os.Exit(status)
}
