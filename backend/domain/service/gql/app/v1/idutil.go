package appv1

import (
	"github.com/samber/oops"

	"app/library/idutil"
)

const (
	ResourceNameTask = "Task"
)

func EncodeTaskID(id uint64) string {
	return idutil.Encode(ResourceNameTask, id)
}

func DecodeTaskID(encodedID string) (uint64, error) {
	return decodeResourceID(ResourceNameTask, encodedID)
}

func decodeResourceID(resourceNameExpected string, encodedID string) (uint64, error) {
	resourceName, id, err := idutil.Decode(encodedID)
	if err != nil {
		return 0, err
	}
	if resourceName != resourceNameExpected {
		return 0, oops.Errorf("unexpected resource name: %s", resourceName)
	}

	return id, nil
}
