package function

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	tsapi "github.com/threatstack/ts/api"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

// DismissReason stores different reasons for dismissing an alert
type DismissReason string

const (
	// DismissBusinessOp - Required for Business Operations
	DismissBusinessOp DismissReason = "BUSINESS_OP"
	// DismissCompanyPolicy - Normal per Company Policy
	DismissCompanyPolicy DismissReason = "COMPANY_POLICY"
	// DismissMaintenance - Required Temporarily, for Testing and Maintenance
	DismissMaintenance DismissReason = "MAINTENANCE"
	// DismissOther - Other
	DismissOther DismissReason = "OTHER"
)

// Alert stores all information related to an individual alert
type FullAlert struct {
	Alert        Alert        `json:"alert_info"`
	AlertEvents  AlertEvents  `json:"alert_events"`
	AlertContext AlertContext `json:"alert_context"`
}

type FullAlerts struct {
	Alerts []FullAlert `json:"alerts"`
}

// Alert stores information related to an individual alert
type Alert struct {
	ID                string            `json:"id"`
	Title             string            `json:"title"`
	DataSource        string            `json:"dataSource"`
	CreatedAt         string            `json:"createdAt"`
	IsDismissed       bool              `json:"isDismissed"`
	DismissedAt       string            `json:"dismissedAt"`
	DismissReason     DismissReason     `json:"dismissReason"`
	DismissReasonText string            `json:"dismissReasonText"`
	DismissedBy       string            `json:"dismissedBy"`
	Severity          int               `json:"severity"`
	AgentID           string            `json:"agentId"`
	RuleID            string            `json:"ruleId"`
	RulesetID         string            `json:"rulesetId"`
	Aggregates        map[string]string `json:"aggregates"`
}

// Autogenerated from json response
type AlertEvents struct {
	Events []struct {
		Timestamp        int64 `json:"timestamp"`
		ProviderMetadata struct {
			CloudProvider    string `json:"cloud_provider"`
			AccountID        string `json:"account_id"`
			AvailabilityZone string `json:"availability_zone"`
		} `json:"provider_metadata"`
		NetworkMetadata []struct {
			VpcID            string   `json:"vpc_id"`
			PublicDNSName    []string `json:"public_dns_name"`
			PrivateIPAddress []string `json:"private_ip_address"`
			PrivateDNSName   []string `json:"private_dns_name"`
			NetworkInterface string   `json:"network_interface"`
			PublicIPAddress  []string `json:"public_ip_address"`
			SubnetID         string   `json:"subnet_id"`
		} `json:"network_metadata"`
		AgentMetadata struct {
			AgentID        string `json:"agent_id"`
			DeploymentMode string `json:"deployment_mode"`
			Version        string `json:"version"`
		} `json:"agent_metadata"`
		HostMetadata struct {
			Kernel       string `json:"kernel"`
			Hostname     string `json:"hostname"`
			ImageID      string `json:"image_id"`
			OsVersion    string `json:"os_version"`
			InstanceID   string `json:"instance_id"`
			InstanceType string `json:"instance_type"`
		} `json:"host_metadata"`
		PodName        string   `json:"pod_name"`
		Success        bool     `json:"success"`
		Path           []string `json:"path"`
		Gid            int      `json:"gid"`
		ContainerImage string   `json:"container_image"`
		ID             string   `json:"_id"`
		PodUID         string   `json:"pod_uid"`
		EventClass     string   `json:"eventClass"`
		ContainerID    string   `json:"container_id"`
		Command        string   `json:"command"`
		ParentMetadata struct {
			Exe       string `json:"exe"`
			Pid       int    `json:"pid"`
			Timestamp int64  `json:"timestamp"`
		} `json:"parent_metadata"`
		Auid           int64  `json:"auid"`
		Pid            int    `json:"pid"`
		TsEventType    string `json:"ts_event_type"`
		EventTime      int64  `json:"event_time"`
		EventType      string `json:"event_type"`
		OrganizationID string `json:"organization_id"`
		Session        int64  `json:"session"`
		Exit           string `json:"exit"`
		UID            int    `json:"uid"`
		Cwd            string `json:"cwd"`
		Header         struct {
			ID           int `json:"id"`
			Milliseconds int `json:"milliseconds"`
			SequenceID   int `json:"sequence_id"`
			Timestamp    int `json:"timestamp"`
		} `json:"header"`
		Ppid            int           `json:"ppid"`
		Args            []string      `json:"args"`
		IngestTime      int64         `json:"ingest_time"`
		Tty             string        `json:"tty"`
		Syscall         string        `json:"syscall"`
		Type            string        `json:"type"`
		EventID         string        `json:"event_id"`
		Group           []string      `json:"group"`
		User            string        `json:"user"`
		AgentID         string        `json:"agent_id"`
		Exe             string        `json:"exe"`
		TimeID          string        `json:"time_id"`
		Tags            []interface{} `json:"tags"`
		ContainerLabels struct {
			IoCriContainerdKind       string `json:"io.cri-containerd.kind"`
			IoKubernetesPodName       string `json:"io.kubernetes.pod.name"`
			IoKubernetesPodUID        string `json:"io.kubernetes.pod.uid"`
			IoKubernetesPodNamespace  string `json:"io.kubernetes.pod.namespace"`
			IoKubernetesContainerName string `json:"io.kubernetes.container.name"`
		} `json:"containerLabels"`
	} `json:"events"`
}

// Autogenerated from json response
type AlertContext struct {
	UserAgents []struct {
		User  string `json:"user"`
		Agent []struct {
			AgentID  string `json:"agentId"`
			Hostname string `json:"hostname"`
			Count    int    `json:"count"`
		} `json:"agent"`
		TopAgents []struct {
			AgentID  string `json:"agentId"`
			Hostname string `json:"hostname"`
			Count    int    `json:"count"`
		} `json:"topAgents"`
	} `json:"userAgents"`
	UserProcesses []struct {
		User      string `json:"user"`
		Processes []struct {
			Exe   string `json:"exe"`
			Count int    `json:"count"`
		} `json:"processes"`
		TopProcesses []struct {
			Exe   string `json:"exe"`
			Count int    `json:"count"`
		} `json:"topProcesses"`
	} `json:"userProcesses"`
	UserSources []struct {
		User       string        `json:"user"`
		Sources    []interface{} `json:"sources"`
		TopSources []struct {
			Source      string      `json:"source"`
			Destination interface{} `json:"destination"`
			IsWan       interface{} `json:"isWan"`
			Count       int         `json:"count"`
		} `json:"topSources"`
	} `json:"userSources"`
	CloudtrailSources     interface{} `json:"cloudtrailSources"`
	CloudtrailTasks       interface{} `json:"cloudtrailTasks"`
	CloudtrailAuthMethods interface{} `json:"cloudtrailAuthMethods"`
	Digests               []struct {
		AlertType string            `json:"alertType"`
		Indicator string            `json:"indicator"`
		Digest    string            `json:"digest"`
		Template  string            `json:"template"`
		Values    map[string]string `json:"values,omitempty"`
	} `json:"digests"`
}

type AlertHookRaw struct {
	Alerts []AlertHookData
}

type AlertHookData struct {
	Id               string
	CreatedAt        int
	OrganizationId   string
	OrganizationName string
	ServerOrRegion   string
	Severity         int
	Source           string
	Title            string
}

// Config configures the API object
type Config struct {
	User string
	Key  string
	Org  string
}

type AlertProcessor struct {
	client         *http.Client
	config         tsapi.Config
	destinationUrl string
}

func init() {
	functions.HTTP("ReceiveAlertHook", receiveAlert)
}

func (ap *AlertProcessor) fetchAlert(id string) (Alert, error) {
	var a Alert
	alertEndpoint := fmt.Sprintf("https://api.threatstack.com/v2/alerts/%s", id)
	req, err := tsapi.Request(ap.config, "GET", alertEndpoint, nil)
	if err != nil {
		return Alert{}, fmt.Errorf("failed to construct alert request: %v", err)
	}
	resp, err := ap.client.Do(req)
	if err != nil {
		return Alert{}, fmt.Errorf("fetch alert failed: %v", err)
	}
	if resp.StatusCode != 200 {
		return Alert{}, fmt.Errorf("fetch alert failed with code %d: %v", resp.StatusCode, err)
	}
	err = json.NewDecoder(resp.Body).Decode(&a)
	if err != nil {
		return Alert{}, fmt.Errorf("failed to decode alert response: %v", err)
	}
	return a, nil
}

func (ap *AlertProcessor) fetchEvents(id string) (AlertEvents, error) {
	var e AlertEvents
	alertEndpoint := fmt.Sprintf("https://api.threatstack.com/v2/alerts/%s/events", id)
	req, err := tsapi.Request(ap.config, "GET", alertEndpoint, nil)
	if err != nil {
		return AlertEvents{}, fmt.Errorf("failed to construct alert request: %v", err)
	}
	resp, err := ap.client.Do(req)
	if err != nil {
		return AlertEvents{}, fmt.Errorf("fetch alert failed: %v", err)
	}
	if resp.StatusCode != 200 {
		return AlertEvents{}, fmt.Errorf("fetch alert failed with code %d: %v", resp.StatusCode, err)
	}
	err = json.NewDecoder(resp.Body).Decode(&e)
	if err != nil {
		return AlertEvents{}, fmt.Errorf("failed to decode alert response: %v", err)
	}
	return e, nil
}

func (ap *AlertProcessor) fetchContext(id string) (AlertContext, error) {
	var c AlertContext
	alertEndpoint := fmt.Sprintf("https://api.threatstack.com/v2/alerts/%s/context", id)
	req, err := tsapi.Request(ap.config, "GET", alertEndpoint, nil)
	if err != nil {
		return AlertContext{}, fmt.Errorf("failed to construct alert request: %v", err)
	}
	resp, err := ap.client.Do(req)
	if err != nil {
		return AlertContext{}, fmt.Errorf("fetch alert failed: %v", err)
	}
	if resp.StatusCode != 200 {
		return AlertContext{}, fmt.Errorf("fetch context failed with code %d: %v", resp.StatusCode, err)
	}
	err = json.NewDecoder(resp.Body).Decode(&c)
	if err != nil {
		return AlertContext{}, fmt.Errorf("failed to decode alert response: %v", err)
	}
	return c, nil
}

func (ap *AlertProcessor) processAlert(id string) (FullAlert, error) {
	a, err := ap.fetchAlert((id))
	if err != nil {
		return FullAlert{}, fmt.Errorf("processing alert %s failed: %b", id, err)
	}
	e, err := ap.fetchEvents((id))
	if err != nil {
		return FullAlert{}, fmt.Errorf("processing events %s failed: %b", id, err)
	}
	c, err := ap.fetchContext((id))
	if err != nil {
		return FullAlert{}, fmt.Errorf("processing context %s failed: %b", id, err)
	}
	return FullAlert{
		Alert:        a,
		AlertEvents:  e,
		AlertContext: c,
	}, nil
}

func (ap *AlertProcessor) forwardData(a FullAlerts) error {
	url := ap.destinationUrl
	jsonBytes, _ := json.Marshal(a)
	fmt.Printf("%s\n", string(jsonBytes))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return fmt.Errorf("failed constructing full alert req: %v", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	response, err := ap.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed forwarding full alert: %v", err)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 300 {
		return fmt.Errorf("error code from forward destination: %d - %s", response.StatusCode, string(body))
	}
	return nil
}

func (ap *AlertProcessor) processAlerts(a AlertHookRaw) error {
	// Nothing to process.
	if len(a.Alerts) == 0 {
		return nil
	}

	allAlerts := FullAlerts{}
	for _, v := range a.Alerts {
		if v.Id == "" {
			return errors.New("alertId not found in data")
		}
		alert, err := ap.processAlert(v.Id)
		if err != nil {
			return err
		}
		allAlerts.Alerts = append(allAlerts.Alerts, alert)
	}
	return ap.forwardData(allAlerts)
}

// receiveAlert handles incoming Webhook Alerts from Threatstack.
func receiveAlert(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405 - Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	apiKey, isSet := os.LookupEnv("API_KEY")
	if !isSet {
		http.Error(w, "500 - InternalServerError", http.StatusInternalServerError)
		log.Printf("API_KEY Not Set")
		return
	}
	m, _ := url.ParseQuery(r.URL.RawQuery)
	if len(m["api_key"]) == 0 || m["api_key"][0] != apiKey {
		http.Error(w, "401 - Unauthorized", http.StatusUnauthorized)
		return
	}
	dst, isSet := os.LookupEnv("DESTINATION_URL")
	if !isSet {
		http.Error(w, "500 - InternalServerError", http.StatusInternalServerError)
		log.Printf("DESTINATION_URL Not Set")
		return
	}
	tsUser, isSet := os.LookupEnv("TS_USER")
	if !isSet {
		http.Error(w, "500 - InternalServerError", http.StatusInternalServerError)
		log.Printf("TS_USER Not Set")
		return
	}
	tsKey, isSet := os.LookupEnv("TS_KEY")
	if !isSet {
		http.Error(w, "500 - InternalServerError", http.StatusInternalServerError)
		log.Printf("TS_KEY Not Set")
		return
	}
	tsOrg, isSet := os.LookupEnv("TS_ORG")
	if !isSet {
		http.Error(w, "500 - InternalServerError", http.StatusInternalServerError)
		log.Printf("TS_ORG Not Set")
		return
	}
	var a AlertHookRaw
	// Try to decode the request body into the Alert struct. Return 400 on bad request.
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		log.Printf("Failed decoding request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ap := AlertProcessor{
		client:         &http.Client{},
		destinationUrl: dst,
		config: tsapi.Config{
			User: tsUser,
			Key:  tsKey,
			Org:  tsOrg,
		},
	}
	err = ap.processAlerts(a)
	if err != nil {
		log.Printf("Failed decoding request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//log.Printf("%v\n", a)
	fmt.Fprintln(w, "OK")
}
