package clients

import (
	"github.com/avinetworks/sdk/go/models"
	"github.com/avinetworks/sdk/go/session"
	"testing"
	"github.com/golang/glog"
)

func getSessions(t *testing.T) (
	credentialsSession *session.AviSession, authTokenSession *session.AviSession) {
	/* Test username/password authentication */
	credentialsSession, err := session.NewAviSession("10.10.25.201",
		"admin", session.SetPassword("avi123"), session.SetInsecure)
	if err != nil {
		t.Fatalf("Session Creation failed: %s", err)
	}

	/* Test token authentication */
	authTokenSession, err = session.NewAviSession("localhost", "admin",
		session.SetTokenAuth, session.SetInsecure)

	if err != nil {
		t.Fatalf("Session Creation failed: %s", err)
	}

	return credentialsSession, authTokenSession
}

// Create Pool
// Get All
// Update Pool
// Delete Pool
func testAviPoolClient(t * testing.T, aviSession *session.AviSession) {
	pclient := NewPoolClient(aviSession)

	obj := models.Pool{}
	obj.Name = "testpool"
	objp, err := pclient.Create(&obj)
	if err != nil {
		t.Fatalf("ERROR: %s", err)
	}
	glog.Infof("res: %+v; err: %+v", *objp, err)

	aserver := models.Server{}
	aserver.IP = &models.IPAddr{Addr: "10.10.10.10", Type: "V4"}
	objp.Servers = append(objp.Servers, &aserver)

	objp, err = pclient.Update(objp)
	glog.Infof("After update -- res: %+v; err: %+v", *objp, err)
	glog.Infof("After update -- server: %+v", objp.Servers[0])
	if err != nil {
		t.Fatalf("ERROR: %s", err)
	}

	lobjp, err := pclient.GetAll()
	glog.Infof("res: %+v; err: %+v", *lobjp[0], err)
	if err != nil {
		t.Fatalf("ERROR: %s", err)
	}

	objp, err = pclient.GetByName("testpool")
	glog.Infof("res: %+v; err: %+v", *objp, err)
	if err != nil {
		t.Fatalf("ERROR: %s", err)
	}

	objp, err = pclient.Get(objp.UUID)
	glog.Infof("res: %+v; err: %+v", *objp, err)
	objp.Enabled = false
	if err != nil {
		t.Fatalf("ERROR: %s", err)
	}

	objp, err = pclient.Update(objp)
	glog.Infof("res: %+v; err: %+v", *objp, err)
	if err != nil {
		t.Fatalf("ERROR: %s", err)
	}

	err = pclient.Delete(objp.UUID)
	glog.Infof("err: %+v", err)
	if err != nil {
		t.Fatalf("ERROR: %s", err)
	}
}

func TestAviPoolClient(t *testing.T) {
	credentialsSession, authTokenSession := getSessions(t)
	testAviPoolClient(t, credentialsSession)
	testAviPoolClient(t, authTokenSession)
}
