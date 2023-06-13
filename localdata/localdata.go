package localdata

import (
	"encoding/json"
	"fmt"
	"proofofaccess/database"
	"strings"
	"time"
)

var Time = time.Now()

const Layout = "2006-01-02 15:04:05.999999 -0700 MST m=+0.000000000"

type Message struct {
	Type   string `json:"type"`
	Hash   string `json:"hash"`
	CID    string `json:"CID"`
	Status string `json:"status"`
}

var Synced = false
var NodeName = ""
var PeerNames = []string{}
var ValidatorNames = []string{}
var PeerStats = map[string][]string{}
var PeerSize = map[string]int{}
var Validators = map[string]bool{}
var ValidatorsStatus = map[string]string{}
var NodesStatus = map[string]string{}
var NodeType int
var PinFileCids = map[string][]string{}
var SavedRefs = map[string][]string{}
var NodeCount = 0
var SyncedPercentage = float32(0)

// SaveTime
// Saves the time to the database
func SaveTime(salt string) {
	Time = time.Now()
	timeStr := Time.Format(Layout)
	database.Update([]byte(salt+"time"), []byte(timeStr))
}

// GetTime
// Gets the time from the database
func GetTime(hash string) time.Time {
	data := database.Read([]byte(hash + "time"))
	parsedTime, _ := time.Parse(Layout, string(data))
	return parsedTime
}

// SetElapsed
// Sets the elapsed time to the database
func SetElapsed(hash string, elapsed time.Duration) {
	database.Update([]byte(hash+"elapsed"), []byte(elapsed.String()))
}

// GetElapsed
// Gets the elapsed time from the database
func GetElapsed(hash string) time.Duration {
	data := database.Read([]byte(hash + "elapsed"))
	parsedTime, _ := time.ParseDuration(string(data))
	return parsedTime
}

// GetStatus
// Gets the status from the database
func GetStatus(seed string) Message {
	data := database.Read([]byte(seed))
	var message Message
	err := json.Unmarshal([]byte(string(data)), &message)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
	return message
}

// SetStatus
// Sets the status to the database
func SetStatus(seed string, cid string, status string) {
	fmt.Println("SetStatus", seed, cid, status)
	jsonString := `{"type": "ProofOfAccess", "CID":"` + cid + `", "seed":"` + seed + `", "status":"` + status + `"}`
	jsonString = strings.TrimSpace(jsonString)
	database.Update([]byte(seed), []byte(jsonString))
}

// SetNodeName
// Sets the node name in localdata
func SetNodeName(name string) {
	NodeName = name
}

func GetNodeName() string {
	return NodeName
}
func RemoveDuplicates(peerNames []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for v := range peerNames {
		if encountered[peerNames[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as encountered.
			encountered[peerNames[v]] = true
			// Append to result slice.
			result = append(result, peerNames[v])
		}
	}

	return result
}
