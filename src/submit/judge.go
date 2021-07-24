package submit

import (
	"fmt"
	"os"
	// "net/http"

	. "github.com/Leng-Kai/bow-code-API-server/schemas"
	"github.com/Leng-Kai/bow-code-API-server/util"

	_ "github.com/joho/godotenv/autoload"
)

type judgeRequest struct {
	source_code     string `json:"source_code"`
	language_id     int    `json:"language_id"`
	stdin           string `json:"stdin"`
	expected_output string `json:"expected_output"`
	callback_url    string `json:"callback_url"`
}

func init() {

}

func SendJudgeRequests(problem Problem, us UserSubmission, sid SubmissionID) error {
	var err error
	source_code := us.SourceCode
	language_id := us.LanguageID
	self_url := os.Getenv("SELF_URL")

	inputs := problem.Testcase.Input
	outputs := problem.Testcase.ExpectedOutput

	for i, _ := range inputs {
		stdin := inputs[i]
		expected_output := outputs[i]
		callback_url := fmt.Sprintf("%s/submit/%s/%d", self_url, sid, i)

		url := os.Getenv("JUDGE0_URL")
		body := judgeRequest{
			source_code, language_id, stdin, expected_output, callback_url,
		}

		err = util.SendHTTPRequest(url, body)
		if err != nil {
			break
		}
	}

	return err
}
