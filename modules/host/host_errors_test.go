package host

import (
	"net"
	"os"
	"path/filepath"
	"testing"

	"github.com/NebulousLabs/Sia/modules"
	"github.com/NebulousLabs/Sia/persist"
	"github.com/NebulousLabs/bolt"
)

// dependencyErrMkdirAll is a dependency set that returns an error when MkdirAll
// is called.
type dependencyErrMkdirAll struct {
	productionDependencies
}

func (dependencyErrMkdirAll) MkdirAll(_ string, _ os.FileMode) error {
	return mockErrMkdirAll
}

// TestHostFailedMkdirAll initializes the host using a call to MkdirAll that
// will fail.
func TestHostFailedMkdirAll(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	ht, err := blankHostTester("TestHostFailedMkdirAll")
	if err != nil {
		t.Fatal(err)
	}
	err = ht.host.Close()
	if err != nil {
		t.Fatal(err)
	}
	_, err = newHost(dependencyErrMkdirAll{}, ht.cs, ht.tpool, ht.wallet, ":0", filepath.Join(ht.persistDir, modules.HostDir))
	if err != mockErrMkdirAll {
		t.Fatal(err)
	}
}

// dependencyErrNewLogger is a dependency set that returns an error when
// NewLogger is called.
type dependencyErrNewLogger struct {
	productionDependencies
}

func (dependencyErrNewLogger) NewLogger(_ string) (*persist.Logger, error) {
	return nil, mockErrNewLogger
}

// TestHostFailedNewLogger initializes the host using a call to NewLogger that
// will fail.
func TestHostFailedNewLogger(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	ht, err := blankHostTester("TestHostFailedNewLogger")
	if err != nil {
		t.Fatal(err)
	}
	err = ht.host.Close()
	if err != nil {
		t.Fatal(err)
	}
	_, err = newHost(dependencyErrNewLogger{}, ht.cs, ht.tpool, ht.wallet, ":0", filepath.Join(ht.persistDir, modules.HostDir))
	if err != mockErrNewLogger {
		t.Fatal(err)
	}
}

// dependencyErrOpenDatabase is a dependency that returns an error when
// OpenDatabase is called.
type dependencyErrOpenDatabase struct {
	productionDependencies
}

func (dependencyErrOpenDatabase) OpenDatabase(_ persist.Metadata, _ string) (*persist.BoltDatabase, error) {
	return nil, mockErrOpenDatabase
}

// TestHostFailedOpenDatabase initializes the host using a call to OpenDatabase
// that has been mocked to fail.
func TestHostFailedOpenDatabase(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	ht, err := blankHostTester("TestHostFailedOpenDatabase")
	if err != nil {
		t.Fatal(err)
	}
	err = ht.host.Close()
	if err != nil {
		t.Fatal(err)
	}
	_, err = newHost(dependencyErrOpenDatabase{}, ht.cs, ht.tpool, ht.wallet, ":0", filepath.Join(ht.persistDir, modules.HostDir))
	if err != mockErrOpenDatabase {
		t.Fatal(err)
	}
}

// TestUnsuccessfulDBInit sets the stage for an error to be triggered when the
// host tries to initialize the database. The host should return the error.
func TestUnsuccessfulDBInit(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	// Create a blank host tester so that all the host dependencies are
	// available.
	ht, err := blankHostTester("TestSetPersistentSettings")
	if err != nil {
		t.Fatal(err)
	}
	// Corrupt the host database by deleting BucketStorageObligations, which is
	// used to tell whether the host has been initialized or not. This will
	// cause errors to be returned when initialization tries to create existing
	// buckets.
	err = ht.host.db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket(BucketStorageObligations)
	})
	if err != nil {
		t.Fatal(err)
	}
	// Close the host so a new host can be created after the database has been
	// corrupted.
	err = ht.host.Close()
	if err != nil {
		t.Fatal(err)
	}
	_, err = New(ht.cs, ht.tpool, ht.wallet, ":0", filepath.Join(ht.persistDir, modules.HostDir))
	if err == nil {
		t.Fatal("expecting initDB to fail")
	}
}

// dependencyErrLoadFile is a dependency that returns an error when
// LoadFile is called.
type dependencyErrLoadFile struct {
	productionDependencies
}

func (dependencyErrLoadFile) LoadFile(_ persist.Metadata, _ interface{}, _ string) error {
	return mockErrLoadFile
}

// TestHostFailedLoadFile initializes the host using a call to LoadFile that
// has been mocked to fail.
func TestHostFailedLoadFile(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	ht, err := blankHostTester("TestHostFailedLoadFile")
	if err != nil {
		t.Fatal(err)
	}
	err = ht.host.Close()
	if err != nil {
		t.Fatal(err)
	}
	_, err = newHost(dependencyErrLoadFile{}, ht.cs, ht.tpool, ht.wallet, ":0", filepath.Join(ht.persistDir, modules.HostDir))
	if err != mockErrLoadFile {
		t.Fatal(err)
	}
}

// dependencyErrListen is a dependency that returns an error when Listen is
// called.
type dependencyErrListen struct {
	productionDependencies
}

func (dependencyErrListen) Listen(_, _ string) (net.Listener, error) {
	return nil, mockErrListen
}

// TestHostFailedListen initializes the host using a call to Listen that
// has been mocked to fail.
func TestHostFailedListen(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	ht, err := blankHostTester("TestHostFailedListen")
	if err != nil {
		t.Fatal(err)
	}
	err = ht.host.Close()
	if err != nil {
		t.Fatal(err)
	}
	_, err = newHost(dependencyErrListen{}, ht.cs, ht.tpool, ht.wallet, ":0", filepath.Join(ht.persistDir, modules.HostDir))
	if err != mockErrListen {
		t.Fatal(err)
	}
}
