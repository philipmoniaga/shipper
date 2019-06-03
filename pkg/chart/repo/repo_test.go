package repo

import (
	"testing"
)

const (
	IndexYamlResp = `
---
apiVersion: v1
entries:
  alpine:
    - created: 2016-10-06T16:23:20.499814565-06:00
      description: Deploy a basic Alpine Linux pod
      digest: 99c76e403d752c84ead610644d4b1c2f2b453a74b921f422b9dcb8a7c8b559cd
      home: https://k8s.io/helm
      name: alpine
      sources:
      - https://github.com/helm/helm
      urls:
      - https://technosophos.github.io/tscharts/alpine-0.2.0.tgz
      version: 0.2.0
    - created: 2016-10-06T16:23:20.499543808-06:00
      description: Deploy a basic Alpine Linux pod
      digest: 515c58e5f79d8b2913a10cb400ebb6fa9c77fe813287afbacf1a0b897cd78727
      home: https://k8s.io/helm
      name: alpine
      sources:
      - https://github.com/helm/helm
      urls:
      - https://technosophos.github.io/tscharts/alpine-0.1.0.tgz
      version: 0.1.0
  nginx:
    - created: 2016-10-06T16:23:20.499543808-06:00
      description: Create a basic nginx HTTP server
      digest: aaff4545f79d8b2913a10cb400ebb6fa9c77fe813287afbacf1a0b897cdffffff
      home: https://k8s.io/helm
      name: nginx
      sources:
      - https://github.com/helm/charts
      urls:
      - https://technosophos.github.io/tscharts/nginx-1.1.0.tgz
      version: 1.1.0
generated: 2016-10-06T16:23:20.499029981-06:00
`
)

func TestRefreshIndex(t *testing.T) {
	tests := []struct {
		name             string
		fetchBody        string
		fetchErr         error
		repoUrl          string
		expectedFetchUrl string
		expectedErr      error
	}{
		{
			name:             "Plain fetch",
			fetchBody:        IndexYamlResp,
			fetchErr:         nil,
			repoUrl:          "https://registry.example.com/charts",
			expectedFetchUrl: "https://registry.example.com/charts/index.yaml",
			expectedErr:      nil,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			var fetchedUrl string

			oldFetch := fetch
			fetch = func(url string) ([]byte, error) {
				fetchedUrl = url
				return []byte(testCase.fetchBody), testCase.fetchErr
			}
			defer func() { fetch = oldFetch }()

			cache := NewTestCache(testCase.name)

			repo := NewRepo(testCase.repoUrl, cache)

			err := repo.RefreshIndex()

			if !equivalent(err, testCase.expectedErr) {
				t.Fatalf("Unexpected error: %q, want: %q", err, testCase.expectedErr)
			}

			if fetchedUrl != testCase.expectedFetchUrl {
				t.Fatalf("Unexpected fetch URL: %q, want: %q", fetchedUrl, testCase.expectedFetchUrl)
			}
		})
	}
}

func equivalent(err1, err2 error) bool {
	if err1 == nil && err2 == nil {
		return true
	}
	if err1 != nil && err2 != nil {
		return err1.Error() == err2.Error()
	}
	return false
}
