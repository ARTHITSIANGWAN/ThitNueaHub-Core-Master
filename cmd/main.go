package main
import (
	"crypto/hmac	"
	"crypto/sha256	"
	"encoding/hex	"
	"encoding/json	"
	"fmt	"
	"net/http	"
	"github.com/syumai/workers	"
)
const secretKey = 	"thitnueahub_secret_token_v86	"
type job struct {
JobID string `json:	"job_id	"`
	Action string `json:	"action	"`
	Topic string `json:	"topic	"`
	Ts int64 `json:	"ts	"`
	Ttl int `json:	"ttl	"`
	Sig string `json:	"sig	"`
}
func validate	Signature(j job) bool {
message := fmt.	Sprintf(	"%s|%s|%s	", j.	JobID, j.	Action, j.	Topic)
mac := hmac.New(sha256.New, []byte(secretKey))
mac.Write([]byte(message))
return j.	Sig == hex.Encode	To	String(mac.	Sum(nil))
}
func processHandler(w http.ResponseWriter, r *http.Request) {
defer func() {
if err := recover(); err != nil {
w.WriteHeader(http.	StatusInternal	ServerError)
fmt.Fprintf(w, `{	"status	":	"edge_panic_recovered	"}`)
}
}()
if r.Method != http.MethodPost {
http.Error(w, 	"method not allowed	", http.	StatusMethodNot	Allowed)
return
}
var j job
if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
http.Error(w, 	"bad request	", http.	StatusBadRequest)
return
}
if !validate	Signature(j) {
http.Error(w, 	"unauthorized_garbage_detected	", http.	StatusUnauthorized)
return
}
w.Header().	Set(	"Content-	Type	", 	"application/json	")
w.WriteHeader(http.	Status	Accepted)
switch j.	Action {
case 	"scrape	":
fmt.Fprintf(w, `{	"status	":	"dispatched	",	"target	":	"services/
scraper	",	"job_id	":	"%s	"}`, j.	JobID)
case 	"schedule	":
fmt.Fprintf(w, `{	"status	":	"dispatched	",	"target	":	"services/
scheduler	",	"job_id	":	"%s	"}`, j.	JobID)
case 	"agent_call	":
fmt.Fprintf(w, `{	"status	":	"dispatched	",	"target	":	"my_agent/app	",	"job_id	":	"%s	"}`, 
j.	JobID)
default:
fmt.Fprintf(w, `{	"status	":	"dispatched	",	"target	":	"general_spoke	",	"job_id	":	"%s	"}
`, j.	JobID)
}
}
func main() {
	http.HandleFunc(	"/process	", processHandler)
workers.	Serve(nil)
}
