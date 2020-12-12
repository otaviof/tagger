package services

import (
	"context"
	"reflect"
	"strings"
	"testing"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	coreinf "k8s.io/client-go/informers"
	corfake "k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/tools/cache"

	"github.com/mattbaird/jsonpatch"

	tagfake "github.com/ricardomaraschini/tagger/imagetags/generated/clientset/versioned/fake"
	taginf "github.com/ricardomaraschini/tagger/imagetags/generated/informers/externalversions"
	imagtagv1 "github.com/ricardomaraschini/tagger/imagetags/v1"
)

func TestCurrentReferenceForTagByName(t *testing.T) {
	for _, tt := range []struct {
		name    string
		itname  string
		objects []runtime.Object
		expref  string
		err     string
	}{
		{
			name: "no image tag",
		},
		{
			name:   "generation not existent",
			itname: "tag",
			objects: []runtime.Object{
				&imagtagv1.Tag{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "tag",
						Namespace: "default",
					},
				},
			},
		},
		{
			name:   "happy path",
			expref: "my ref",
			itname: "tag",
			objects: []runtime.Object{
				&imagtagv1.Tag{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "tag",
						Namespace: "default",
					},
					Status: imagtagv1.TagStatus{
						Generation: 60,
						References: []imagtagv1.HashReference{
							{Generation: 63},
							{Generation: 62},
							{Generation: 61},
							{
								Generation:     60,
								ImageReference: "my ref",
							},
						},
					},
				},
			},
		},
		{
			name:   "tag in different namespace",
			itname: "tag",
			objects: []runtime.Object{
				&imagtagv1.Tag{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "tag",
						Namespace: "another-namespace",
					},
					Status: imagtagv1.TagStatus{
						Generation: 0,
						References: []imagtagv1.HashReference{
							{
								Generation:     0,
								ImageReference: "ref",
							},
						},
					},
				},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			fakecli := tagfake.NewSimpleClientset(tt.objects...)
			informer := taginf.NewSharedInformerFactory(fakecli, time.Minute)
			taglis := informer.Images().V1().Tags().Lister()
			informer.Start(ctx.Done())
			if !cache.WaitForCacheSync(
				ctx.Done(),
				informer.Images().V1().Tags().Informer().HasSynced,
			) {
				t.Fatal("errors waiting for caches to sync")
			}

			svc := NewTag(nil, nil, taglis, nil, nil, nil)
			ref, err := svc.CurrentReferenceForTagByName("default", tt.itname)
			if err != nil {
				if len(tt.err) == 0 {
					t.Errorf("unexpected error: %s", err)
				} else if !strings.Contains(err.Error(), tt.err) {
					t.Errorf("expecting %q, %q received instead", tt.err, err)
				}
			} else if len(tt.err) > 0 {
				t.Errorf("expecting error %q, nil received instead", tt.err)
			}

			if ref != tt.expref {
				t.Errorf("expecting ref %q, received %q instead", tt.expref, ref)
			}
		})
	}
}

func TestPatchForPod(t *testing.T) {
	for _, tt := range []struct {
		name     string
		pod      corev1.Pod
		tags     []runtime.Object
		replicas []runtime.Object
		expected []jsonpatch.JsonPatchOperation
		err      string
	}{
		{
			name: "pod without owner",
			pod:  corev1.Pod{},
		},
		{
			name: "job pod",
			pod: corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "default",
					Name:      "my-pod",
					OwnerReferences: []metav1.OwnerReference{
						{
							Kind: "Job",
						},
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Image: "imagetag",
						},
					},
				},
			},
			tags: []runtime.Object{
				&imagtagv1.Tag{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "imagetag",
						Namespace: "default",
					},
					Status: imagtagv1.TagStatus{
						Generation: 1,
						References: []imagtagv1.HashReference{
							{
								Generation: 0,
							},
						},
					},
				},
			},
		},
		{
			name: "happy path",
			expected: []jsonpatch.JsonPatchOperation{
				{
					Operation: "replace",
					Path:      "/spec/containers/0/image",
					Value:     "image ref",
				},
			},
			pod: corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "default",
					Name:      "my-pod",
					OwnerReferences: []metav1.OwnerReference{
						{
							Kind: "ReplicaSet",
							Name: "replicaset",
						},
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Image: "imagetag",
						},
					},
				},
			},
			tags: []runtime.Object{
				&imagtagv1.Tag{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "imagetag",
						Namespace: "default",
					},
					Status: imagtagv1.TagStatus{
						Generation: 0,
						References: []imagtagv1.HashReference{
							{
								Generation:     0,
								ImageReference: "image ref",
							},
						},
					},
				},
			},
			replicas: []runtime.Object{
				&appsv1.ReplicaSet{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "replicaset",
						Namespace: "default",
						Annotations: map[string]string{
							"image-tag": "true",
						},
					},
				},
			},
		},
		{
			name: "replica without annotation",
			pod: corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "default",
					Name:      "my-pod",
					OwnerReferences: []metav1.OwnerReference{
						{
							Kind: "ReplicaSet",
							Name: "replicaset",
						},
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Image: "imagetag",
						},
					},
				},
			},
			replicas: []runtime.Object{
				&appsv1.ReplicaSet{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "replicaset",
						Namespace: "default",
					},
				},
			},
		},
		{
			name: "non existing replica set",
			err:  "not found",
			pod: corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "default",
					Name:      "my-pod",
					OwnerReferences: []metav1.OwnerReference{
						{
							Kind: "ReplicaSet",
							Name: "replicaset",
						},
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Image: "imagetag",
						},
					},
				},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			tagcli := tagfake.NewSimpleClientset(tt.tags...)
			taginf := taginf.NewSharedInformerFactory(tagcli, time.Minute)
			taglis := taginf.Images().V1().Tags().Lister()

			corcli := corfake.NewSimpleClientset(tt.replicas...)
			corinf := coreinf.NewSharedInformerFactory(corcli, time.Minute)
			rslist := corinf.Apps().V1().ReplicaSets().Lister()

			taginf.Start(ctx.Done())
			corinf.Start(ctx.Done())
			if !cache.WaitForCacheSync(
				ctx.Done(),
				taginf.Images().V1().Tags().Informer().HasSynced,
				corinf.Apps().V1().ReplicaSets().Informer().HasSynced,
			) {
				t.Fatal("errors waiting for caches to sync")
			}

			svc := NewTag(nil, nil, taglis, rslist, nil, nil)
			patch, err := svc.PatchForPod(tt.pod)
			if err != nil {
				if len(tt.err) == 0 {
					t.Errorf("unexpected error: %s", err)
				} else if !strings.Contains(err.Error(), tt.err) {
					t.Errorf("expecting %q, %q received instead", tt.err, err)
				}
			} else if len(tt.err) > 0 {
				t.Errorf("expecting error %q, nil received instead", tt.err)
			}

			if !reflect.DeepEqual(tt.expected, patch) {
				t.Errorf("patch mismatch: %v, %v", tt.expected, patch)
			}
		})
	}
}

func Test_prependHashReference(t *testing.T) {
	for _, tt := range []struct {
		name      string
		current   []imagtagv1.HashReference
		reference imagtagv1.HashReference
		expected  []imagtagv1.HashReference
	}{
		{
			name:    "nil current generations slice",
			current: nil,
			reference: imagtagv1.HashReference{
				Generation: 1,
			},
			expected: []imagtagv1.HashReference{
				{Generation: 1},
			},
		},
		{
			name:    "empty current generations slice",
			current: []imagtagv1.HashReference{},
			reference: imagtagv1.HashReference{
				Generation: 1,
			},
			expected: []imagtagv1.HashReference{
				{Generation: 1},
			},
		},
		{
			name: "full current generations slice",
			current: []imagtagv1.HashReference{
				{Generation: 4},
				{Generation: 3},
				{Generation: 2},
				{Generation: 1},
				{Generation: 0},
			},
			reference: imagtagv1.HashReference{
				Generation: 5,
			},
			expected: []imagtagv1.HashReference{
				{Generation: 5},
				{Generation: 4},
				{Generation: 3},
				{Generation: 2},
				{Generation: 1},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewTag(nil, nil, nil, nil, nil, nil)
			res := svc.prependHashReference(tt.reference, tt.current)
			if reflect.DeepEqual(res, tt.expected) {
				return
			}
			t.Errorf("expected %+v, %+v received", tt.expected, res)
		})
	}
}

func TestUpdate(t *testing.T) {
	for _, tt := range []struct {
		name       string
		tag        *imagtagv1.Tag
		err        string
		corObjects []runtime.Object
		tagObjects []runtime.Object
	}{
		{
			name: "empty tag",
			err:  "empty tag reference",
			tag: &imagtagv1.Tag{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "default",
					Name:      "empty-tag",
				},
			},
		},
		{
			name: "import of non existing tag",
			err:  "not found",
			tag: &imagtagv1.Tag{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "default",
					Name:      "new-tag",
				},
				Spec: imagtagv1.TagSpec{
					From:       "centos:latest",
					Generation: 0,
				},
			},
		},
		{
			name: "first import (happy path)",
			tag: &imagtagv1.Tag{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "default",
					Name:      "new-tag",
				},
				Spec: imagtagv1.TagSpec{
					From:       "centos:latest",
					Generation: 0,
				},
			},
			tagObjects: []runtime.Object{
				&imagtagv1.Tag{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default",
						Name:      "new-tag",
					},
					Spec: imagtagv1.TagSpec{
						From:       "centos:latest",
						Generation: 0,
					},
				},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			corcli := corfake.NewSimpleClientset(tt.corObjects...)
			corinf := coreinf.NewSharedInformerFactory(corcli, time.Minute)
			cmlist := corinf.Core().V1().ConfigMaps().Lister()
			seclis := corinf.Core().V1().Secrets().Lister()
			replis := corinf.Apps().V1().ReplicaSets().Lister()
			deplis := corinf.Apps().V1().Deployments().Lister()

			tagcli := tagfake.NewSimpleClientset(tt.tagObjects...)
			taginf := taginf.NewSharedInformerFactory(tagcli, time.Minute)
			taglis := taginf.Images().V1().Tags().Lister()

			corinf.Start(ctx.Done())
			taginf.Start(ctx.Done())
			if !cache.WaitForCacheSync(
				ctx.Done(),
				corinf.Core().V1().ConfigMaps().Informer().HasSynced,
				corinf.Core().V1().Secrets().Informer().HasSynced,
				corinf.Apps().V1().ReplicaSets().Informer().HasSynced,
				corinf.Apps().V1().Deployments().Informer().HasSynced,
				taginf.Images().V1().Tags().Informer().HasSynced,
			) {
				t.Fatal("errors waiting for caches to sync")
			}

			syssvc := NewSysContext(cmlist, seclis)
			impsvc := NewImporter(syssvc)
			svc := NewTag(corcli, tagcli, taglis, replis, deplis, impsvc)

			err := svc.Update(ctx, tt.tag)
			if err != nil {
				if len(tt.err) == 0 {
					t.Errorf("unexpected error: %s", err)
				} else if !strings.Contains(err.Error(), tt.err) {
					t.Errorf("expecting %q, %q received instead", tt.err, err)
				}
			} else if len(tt.err) > 0 {
				t.Errorf("expecting error %q, nil received instead", tt.err)
			}
		})
	}
}
