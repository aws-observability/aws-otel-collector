package cadvisor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsContainerInContainer(t *testing.T) {
	cases := []struct {
		name string
		path string
		dind bool
	}{
		{
			"guranteed",
			"/kubepods.slice/kubepods-podc8f7bb69_65f2_4b61_ae5a_9b19ac47a239.slice/docker-523b624a86a2a74c2bedf586d8448c86887ef7858a8dec037d6559e5ad3fccb5.scope",
			false,
		},
		{
			"guranteed-dind",
			"/kubepods.slice/kubepods-burstable-podc9adcee4_c874_4dad_8bc8_accdbd67ac3a.slice/docker-e58cfbc8b67f6e1af458efdd31cb2a8abdbf9f95db64f4c852b701285a09d40e.scope/docker/fb651068cfbd4bf3d45fb092ec9451f8d1a36b3753687bbaa0a9920617eae5b9",
			true,
		},
		{
			"burstable",
			"/kubepods.slice/kubepods-besteffort.slice/kubepods-besteffort-podab0e310c_0bdb_48e8_ac87_81a701514645.slice/docker-caa8a5e51cd6610f8f0110b491e8187d23488b9635acccf0355a7975fd3ff158.scope",
			false,
		},
		{
			"burstable-dind",
			"/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-podc9adcee4_c874_4dad_8bc8_accdbd67ac3a.slice/docker-e58cfbc8b67f6e1af458efdd31cb2a8abdbf9f95db64f4c852b701285a09d40e.scope/docker/fb651068cfbd4bf3d45fb092ec9451f8d1a36b3753687bbaa0a9920617eae5b9",
			true,
		},
	}
	for _, c := range cases {
		assert.Equal(t, c.dind, isContainerInContainer(c.path), c.name)
	}
}
