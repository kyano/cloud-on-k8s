// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package volume

import (
	corev1 "k8s.io/api/core/v1"
)

// NewConfigMapVolume creates a new ConfigMapVolume struct
func NewConfigMapVolume(name, mountPath string) ConfigMapVolume {
	return ConfigMapVolume{
		name:      name,
		mountPath: mountPath,
	}
}

// ConfigMapVolume defines a volume to expose a configmap
type ConfigMapVolume struct {
	name      string
	mountPath string
	items     []corev1.KeyToPath
}

// VolumeMount returns the k8s volume mount.
func (cm ConfigMapVolume) VolumeMount() corev1.VolumeMount {
	return corev1.VolumeMount{
		Name:      cm.name,
		MountPath: cm.mountPath,
		ReadOnly:  true,
	}
}

// Volume returns the k8s volume.
func (cm ConfigMapVolume) Volume() corev1.Volume {
	return corev1.Volume{
		Name: cm.name,
		VolumeSource: corev1.VolumeSource{
			ConfigMap: &corev1.ConfigMapVolumeSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: cm.name,
				},
				Items:    cm.items,
				Optional: &defaultOptional,
			},
		},
	}
}

// Name returns the name of the volume
func (cm ConfigMapVolume) Name() string {
	return cm.name
}

var _ VolumeLike = ConfigMapVolume{}
