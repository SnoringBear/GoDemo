package agollo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

type Namespace struct {
	Name                string `json:"name"`
	AppID               string `json:"appId"`
	Format              string `json:"format"`
	IsPublic            bool   `json:"isPublic"`
	Comment             string `json:"comment,omitempty"`
	DataChangeCreatedBy string `json:"dataChangeCreatedBy"`
}

func TestCreate02(t *testing.T) {
	// --- Configuration ---
	portalAddress := "http://localhost:8070"
	token := "948c81e4cfd942a0012cb80a9a817ea996fb73421faf0de129ee996ae7b2d5d9"
	creator := "apollo"

	// --- Create a Private Namespace ---
	privateNamespace := Namespace{
		Name:                "my-private-namespace",
		AppID:               "SampleApp",
		Format:              "properties",
		IsPublic:            false,
		Comment:             "This is a private namespace for my application.",
		DataChangeCreatedBy: creator,
	}

	err := CreateApolloNamespace(portalAddress, token, privateNamespace)
	if err != nil {
		fmt.Println("Error creating private namespace:", err)
	}

	fmt.Println("--------------------")

	// --- Create a Public Namespace ---
	publicNamespace := Namespace{
		Name:                "my-public-namespace",
		AppID:               "SampleApp",
		Format:              "json",
		IsPublic:            true,
		Comment:             "This is a public namespace accessible by other applications.",
		DataChangeCreatedBy: creator,
	}

	err = CreateApolloNamespace(portalAddress, token, publicNamespace)
	if err != nil {
		fmt.Println("Error creating public namespace:", err)
	}

}

func CreateApolloNamespace(portalAddress, token string, namespace Namespace) error {
	// Marshal the namespace object to a JSON byte slice.
	payload, err := json.Marshal(namespace)
	if err != nil {
		return fmt.Errorf("failed to marshal namespace payload: %w", err)
	}

	// Construct the API endpoint.
	url := fmt.Sprintf("%s/openapi/v1/apps/%s/appnamespaces", portalAddress, namespace.AppID)

	// Create a new HTTP request.
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to create new HTTP request: %w", err)
	}

	// Set the required headers.
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	// Execute the request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Check the response status code.
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create namespace, status code: %d, response: %s", resp.StatusCode, string(body))
	}

	fmt.Println("Successfully created namespace:", namespace.Name)
	fmt.Println("Response:", string(body))

	return nil
}
