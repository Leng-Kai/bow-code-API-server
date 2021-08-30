package submit

import (
	"fmt"
	"encoding/json"
	// "log"
	"os"

	. "github.com/Leng-Kai/bow-code-API-server/schemas"
	"github.com/Leng-Kai/bow-code-API-server/util"

	_ "github.com/joho/godotenv/autoload"
)

type judgeRequest struct {
	Source_code     string `json:"source_code"`
	Language_id     int    `json:"language_id"`
	Stdin           string `json:"stdin"`
	Expected_output string `json:"expected_output"`
	Callback_url    string `json:"callback_url"`
}

func init() {

}

func SendJudgeRequests(problem Problem, us UserSubmission, sid SubmissionID, crid string) error {
	var err error
	source_code := us.SourceCode
	language_id := us.LanguageID
	self_url := os.Getenv("SELF_URL")

	inputs := problem.Testcase.Input
	outputs := problem.Testcase.ExpectedOutput

	for i, _ := range inputs {
		stdin := inputs[i]
		expected_output := outputs[i]
		callback_url := fmt.Sprintf("%s/submit/%s/%d", self_url, sid.Hex(), i)
		if len(crid) > 0 {
			callback_url = fmt.Sprintf("%s/submit/%s/%s/%d", self_url, crid, sid.Hex(), i)
		}
		url := fmt.Sprintf("%s/%s", os.Getenv("JUDGE0_URL"), "submissions")
		body := judgeRequest{
			source_code, language_id, stdin, expected_output, callback_url,
		}

		jsonData, _ := json.Marshal(body)
		err = util.SendHTTPRequest("POST", url, jsonData)
		if err != nil {
			break
		}
	}

	return err
}
