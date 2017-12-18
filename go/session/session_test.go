package session

import (
	"github.com/avinetworks/sdk/go/models"
	"reflect"
	"testing"
	"github.com/golang/glog"
)

func getSessions(t *testing.T) (
	credentialsSession *AviSession, authTokenSession *AviSession) {
	/* Test username/password authentication */
	credentialsSession, err := NewAviSession("10.10.25.201",
		"admin", SetPassword("avi123"), SetInsecure)
	if err != nil {
		t.Fatalf("Session Creation failed: %s", err)
	}

	/* Test token authentication */
	authTokenSession, err = NewAviSession("localhost", "admin",
		SetTokenAuth, SetInsecure)

	if err != nil {
		t.Fatalf("Session Creation failed: %s", err)
	}

	return credentialsSession, authTokenSession
}

func testAviSession(t *testing.T, avisess *AviSession) {

	var res interface{}
	err := avisess.Get("api/tenant", &res)
	glog.Infof("res: %s, err: %s", res, err)
	resp := res.(map[string]interface{})
	glog.Infof("count: %s", resp["count"])

	// create a tenant
	tenant := make(map[string]string)
	tenant["name"] = "testtenant"
	var tres interface{}
	err = avisess.Post("api/tenant", &tenant, &tres)
	glog.Infof("res: %s, err: %s", tres, err)
	if err != nil {
		t.Error("Tenant Creation failed: ", err)
		return
	}

	// check tenant is created well
	err = avisess.Get("api/tenant?name=testtenant", &res)
	glog.Infof("res: %s, err: %s", res, err)
	if reflect.TypeOf(res).Kind() == reflect.String {
		t.Errorf("Got string instead of json!")
		return
	}
	resp = res.(map[string]interface{})
	glog.Infof("count: %s", resp["count"])
	currCount := resp["count"].(float64)
	if currCount != 1.0 {
		t.Errorf("could not find a tenant with name testtenant")
		return
	}
	tenant["uuid"] = resp["results"].([]interface{})[0].(map[string]interface{})["uuid"].(string)

	// delete the tenant
	err = avisess.Delete("api/tenant/" + tenant["uuid"])
	glog.Infof("err: %s", err)
	if err != nil {
		t.Error("Deletion failed")
		return
	}

	// check to make sure that the tenant is not there any more
	// check tenant is created well
	err = avisess.Get("api/tenant?name=testtenant", &res)
	glog.Infof("res: %s, err: %s", res, err)
	resp = res.(map[string]interface{})
	glog.Infof("count: ", resp["count"])
	currCount = resp["count"].(float64)
	if currCount != 0.0 {
		t.Errorf("Expecting no tenant with that name")
		return
	}

}

func testAviPool(t *testing.T, avisess *AviSession) {
	tpool := models.Pool{}
	pname := "testpool"
	tpool.Name = pname
	var res models.Pool
	err := avisess.Post("api/pool", tpool, &res)
	glog.Infof("res: %s, err: %s", res, err)
	if err != nil {
		t.Fatalf("Pool Creation failed: %s", err)
	}

	var npool2 models.Pool
	err = avisess.GetObjectByName("pool", "testpool", &npool2)

	glog.Infof("npool: %+v err: %+v", npool2, err)
	glog.Infof("name %s: ", npool2.Name)

	err = avisess.Delete("api/pool/" + npool2.UUID)
	if err != nil {
		t.Fatalf("Pool deletion failed: %s", err)
	}
}

func TestAviSession(t *testing.T) {
	credentialsSession, authTokenSession := getSessions(t)
	testAviSession(t, credentialsSession)
	testAviSession(t, authTokenSession)
}

func TestAviPool(t *testing.T) {
	credentialsSession, authTokenSession := getSessions(t)
	testAviPool(t, credentialsSession)
	testAviPool(t, authTokenSession)
}
