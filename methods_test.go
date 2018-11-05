package easycron

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	Status string
}

func TestList(t *testing.T) {
	tr := testStruct{Status: "success"}
	tw, _ := json.Marshal(tr)

	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "/list?token=mytoken")
		// Send response to be tested
		w.Write(tw)
	}))
	// Close the server when test finishess
	defer server.Close()

	// Use Client & URL from our local test server
	api := Client{&Config{server.Client(), "mytoken", server.URL + "/"}}
	testOutput, err := api.List(nil)

	assert.NoError(t, err)
	assert.Equal(t, tr.Status, testOutput.Status)
}

func TestDetail(t *testing.T) {
	tr := testStruct{Status: "success"}
	tw, _ := json.Marshal(tr)

	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "/detail?id=12345&token=mytoken")
		// Send response to be tested
		w.Write(tw)
	}))
	// Close the server when test finishess
	defer server.Close()

	// Use Client & URL from our local test server
	api := Client{&Config{server.Client(), "mytoken", server.URL + "/"}}
	testOutput, err := api.Detail(12345)

	assert.NoError(t, err)
	assert.Equal(t, tr.Status, testOutput.Status)
}

func TestAdd(t *testing.T) {
	tr := testStruct{Status: "success"}
	tw, _ := json.Marshal(tr)

	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "/add?cron_expression=mycron&token=mytoken&url=myurl")
		// Send response to be tested
		w.Write(tw)
	}))
	// Close the server when test finishess
	defer server.Close()

	// Use Client & URL from our local test server
	api := Client{&Config{server.Client(), "mytoken", server.URL + "/"}}
	testOutput, err := api.Add("myurl", "mycron", nil)

	assert.NoError(t, err)
	assert.Equal(t, tr.Status, testOutput.Status)
}

func TestTimezone(t *testing.T) {
	tr := testStruct{Status: "success"}
	tw, _ := json.Marshal(tr)

	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "/timezone?token=mytoken")
		// Send response to be tested
		w.Write(tw)
	}))
	// Close the server when test finishess
	defer server.Close()

	// Use Client & URL from our local test server
	api := Client{&Config{server.Client(), "mytoken", server.URL + "/"}}
	testOutput, err := api.Timezone()

	assert.NoError(t, err)
	assert.Equal(t, tr.Status, testOutput.Status)
}

func TestEdit(t *testing.T) {
	tr := testStruct{Status: "success"}
	tw, _ := json.Marshal(tr)

	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "/edit?id=12345&token=mytoken")
		// Send response to be tested
		w.Write(tw)
	}))
	// Close the server when test finishess
	defer server.Close()

	// Use Client & URL from our local test server
	api := Client{&Config{server.Client(), "mytoken", server.URL + "/"}}
	testOutput, err := api.Edit(12345, nil)

	assert.NoError(t, err)
	assert.Equal(t, tr.Status, testOutput.Status)
}

func TestEnable(t *testing.T) {
	tr := testStruct{Status: "success"}
	tw, _ := json.Marshal(tr)

	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "/enable?id=12345&token=mytoken")
		// Send response to be tested
		w.Write(tw)
	}))
	// Close the server when test finishess
	defer server.Close()

	// Use Client & URL from our local test server
	api := Client{&Config{server.Client(), "mytoken", server.URL + "/"}}
	testOutput, err := api.Enable(12345)

	assert.NoError(t, err)
	assert.Equal(t, tr.Status, testOutput.Status)
}

func TestDisable(t *testing.T) {
	tr := testStruct{Status: "success"}
	tw, _ := json.Marshal(tr)

	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "/disable?id=12345&token=mytoken")
		// Send response to be tested
		w.Write(tw)
	}))
	// Close the server when test finishess
	defer server.Close()

	// Use Client & URL from our local test server
	api := Client{&Config{server.Client(), "mytoken", server.URL + "/"}}
	testOutput, err := api.Disable(12345)

	assert.NoError(t, err)
	assert.Equal(t, tr.Status, testOutput.Status)
}

func TestDelete(t *testing.T) {
	tr := testStruct{Status: "success"}
	tw, _ := json.Marshal(tr)

	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "/delete?id=12345&token=mytoken")
		// Send response to be tested
		w.Write(tw)
	}))
	// Close the server when test finishess
	defer server.Close()

	// Use Client & URL from our local test server
	api := Client{&Config{server.Client(), "mytoken", server.URL + "/"}}
	testOutput, err := api.Delete(12345)

	assert.NoError(t, err)
	assert.Equal(t, tr.Status, testOutput.Status)
}

func TestLogs(t *testing.T) {
	tr := testStruct{Status: "success"}
	tw, _ := json.Marshal(tr)

	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "/logs?id=12345&token=mytoken")
		// Send response to be tested
		w.Write(tw)
	}))
	// Close the server when test finishess
	defer server.Close()

	// Use Client & URL from our local test server
	api := Client{&Config{server.Client(), "mytoken", server.URL + "/"}}
	testOutput, err := api.Logs(12345)

	assert.NoError(t, err)
	assert.Equal(t, tr.Status, testOutput.Status)
}

func TestLog(t *testing.T) {
	tr := testStruct{Status: "success"}
	tw, _ := json.Marshal(tr)

	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "/log?id=12345&log_id=logid&token=mytoken")
		// Send response to be tested
		w.Write(tw)
	}))
	// Close the server when test finishess
	defer server.Close()

	// Use Client & URL from our local test server
	api := Client{&Config{server.Client(), "mytoken", server.URL + "/"}}
	testOutput, err := api.Log(12345, "logid")

	assert.NoError(t, err)
	assert.Equal(t, tr.Status, testOutput.Status)
}
