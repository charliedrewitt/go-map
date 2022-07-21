package mapper_test

import (
	"github.com/charliedrewitt/go-map/mapper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRealisticDtoMap(t *testing.T) {
	// Arrange
	source := policySrc{
		Version: "1.0.0",
		Kind: "MyKind",
		Metadata: metadataSrc{
			Labels: labelsSrc {
				Tier: "web",
			},
		},
		Specification: specSrc{
			Ingress: []ingressSrc{
				{ 
					Action: "Deny", 
					Destination: destinationSrc{
						Ports: []int{ 1, 2, 3 },
						Protocol: "tcp",
					},
					Source: sourceSrc{
						NamespaceSelector: "foobar",
					},
				},
			},
			Order: 1,
			Tier: "application",
			Selector: "fubar",
			Types: []string{ "foo", "bar" },
		},
	}

	// Act
	result := mapper.Map[policyTgt](source)

	// Assert
	assert.Equal(t, source.Version, result.Version)
	assert.Equal(t, source.Kind, result.Kind)
	assert.Equal(t, source.Specification.Tier, result.Specification.Tier)
	assert.Equal(t, source.Specification.Order, result.Specification.Order)
	assert.Equal(t, source.Specification.Selector, result.Specification.Selector)
	assert.Equal(t, source.Specification.Types[0], result.Specification.Types[0])
	assert.Equal(t, source.Specification.Types[1], result.Specification.Types[1])
	assert.Equal(t, source.Metadata.Labels.Tier, result.Metadata.Labels.Tier)
	assert.Equal(t, source.Specification.Ingress[0].Action, result.Specification.Ingress[0].Action)
	assert.Equal(t, source.Specification.Ingress[0].Destination.Ports[0], result.Specification.Ingress[0].Destination.Ports[0])
	assert.Equal(t, source.Specification.Ingress[0].Destination.Ports[1], result.Specification.Ingress[0].Destination.Ports[1])
	assert.Equal(t, source.Specification.Ingress[0].Destination.Ports[2], result.Specification.Ingress[0].Destination.Ports[2])
	assert.Equal(t, source.Specification.Ingress[0].Destination.Protocol, result.Specification.Ingress[0].Destination.Protocol)
	assert.Equal(t, source.Specification.Ingress[0].Source.NamespaceSelector, result.Specification.Ingress[0].Source.NamespaceSelector)
}

type policySrc struct {
	Version       string   `json:"apiVersion"`
	Kind          string   `json:"kind"`
	Metadata      metadataSrc `json:"metadata"`
	Specification specSrc     `json:"spec"`
}

type specSrc struct {
	Ingress  []ingressSrc `json:"ingress"`
	Order    int       `json:"order" default:"101"`
	Selector string    `json:"selector" default:"app-id=''"`
	Tier     string    `json:"tier" default:"application"`
	Types    []string  `json:"types" default:"[\"Ingress\"]"`
}

type metadataSrc struct {
	Labels    labelsSrc `json:"labels"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type ingressSrc struct {
	Action      string      `json:"action" default:"Allow"`
	Destination destinationSrc `json:"destination"`
	Source      sourceSrc      `json:"source"`
}

type labelsSrc struct {
	Tier string `json:"projectcalico.org/tier" default:"application"`
}

type sourceSrc struct {
	NamespaceSelector string `json:"namespaceSelector" default:"app-id==''"`
	Selector          string `json:"selector" default:"app-id==''"`
}

type destinationSrc struct {
	Ports    []int  `json:"ports" default:"[80, 8080, 443]"`
	Protocol string `json:"protocol" default:"TCP"`
}

type policyTgt struct {
	Version       string   `json:"apiVersion"`
	Kind          string   `json:"kind"`
	Metadata      metadataTgt `json:"metadata"`
	Specification specTgt     `json:"spec"`
}

type metadataTgt struct {
	Labels    labelsTgt `json:"labels"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type ingressTgt struct {
	Action      string      `json:"action"`
	Destination destinationTgt `json:"destination"`
	Source      sourceTgt      `json:"source"`
}

type labelsTgt struct {
	Tier string `json:"tier" default:"application"`
}

type specTgt struct {
	Ingress  []ingressTgt `json:"ingress"`
	Order    int       `json:"order"`
	Selector string    `json:"selector"`
	Tier     string    `json:"tier"`
	Types    []string  `json:"types"`
}

type sourceTgt struct {
	NamespaceSelector string `json:"namespaceSelector"`
	Selector          string `json:"selector"`
}

type destinationTgt struct {
	Ports    []int  `json:"ports"`
	Protocol string `json:"protocol"`
}

