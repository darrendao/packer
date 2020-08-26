package iso

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/packer/builder/vsphere/driver"
	"github.com/hashicorp/packer/helper/multistep"
	"github.com/hashicorp/packer/packer"
)

func TestStepRemoteUpload_Run(t *testing.T) {
	state := new(multistep.BasicStateBag)
	state.Put("ui", &packer.BasicUi{
		Reader: new(bytes.Buffer),
		Writer: new(bytes.Buffer),
	})
	driverMock := driver.NewDriverMock()
	state.Put("driver", driverMock)
	state.Put("iso_path", "[datastore] iso/path")

	step := &StepRemoteUpload{
		Datastore:                  "datastore",
		Host:                       "host",
		SetHostForDatastoreUploads: false,
	}

	if action := step.Run(context.TODO(), state); action == multistep.ActionHalt {
		t.Fatalf("Should not halt.")
	}

	if !driverMock.FindDatastoreCalled {
		t.Fatalf("driver.FindDatastore should be called.")
	}
	if !driverMock.DatastoreMock.FileExistsCalled {
		t.Fatalf("datastore.FindDatastore should be called.")
	}
	if !driverMock.DatastoreMock.MakeDirectoryCalled {
		t.Fatalf("datastore.MakeDirectory should be called.")
	}
	if !driverMock.DatastoreMock.UploadFileCalled {
		t.Fatalf("datastore.UploadFile should be called.")
	}
	remotePath, ok := state.GetOk("iso_remote_path")
	if !ok {
		t.Fatalf("state should contain iso_remote_path")
	}
	expectedRemovePath := fmt.Sprintf("[%s] packer_cache//path", driverMock.DatastoreMock.Name())
	if remotePath != expectedRemovePath {
		t.Fatalf("iso_remote_path expected to be %s but was %s", expectedRemovePath, remotePath)
	}
}
