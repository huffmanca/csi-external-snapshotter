/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/kubernetes-csi/external-snapshotter/v2/pkg/apis/volumesnapshot/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VolumeSnapshotClassLister helps list VolumeSnapshotClasses.
type VolumeSnapshotClassLister interface {
	// List lists all VolumeSnapshotClasses in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.VolumeSnapshotClass, err error)
	// Get retrieves the VolumeSnapshotClass from the index for a given name.
	Get(name string) (*v1beta1.VolumeSnapshotClass, error)
	VolumeSnapshotClassListerExpansion
}

// volumeSnapshotClassLister implements the VolumeSnapshotClassLister interface.
type volumeSnapshotClassLister struct {
	indexer cache.Indexer
}

// NewVolumeSnapshotClassLister returns a new VolumeSnapshotClassLister.
func NewVolumeSnapshotClassLister(indexer cache.Indexer) VolumeSnapshotClassLister {
	return &volumeSnapshotClassLister{indexer: indexer}
}

// List lists all VolumeSnapshotClasses in the indexer.
func (s *volumeSnapshotClassLister) List(selector labels.Selector) (ret []*v1beta1.VolumeSnapshotClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.VolumeSnapshotClass))
	})
	return ret, err
}

// Get retrieves the VolumeSnapshotClass from the index for a given name.
func (s *volumeSnapshotClassLister) Get(name string) (*v1beta1.VolumeSnapshotClass, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("volumesnapshotclass"), name)
	}
	return obj.(*v1beta1.VolumeSnapshotClass), nil
}
