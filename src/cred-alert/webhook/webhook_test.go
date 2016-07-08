package webhook_test

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/google/go-github/github"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/pivotal-golang/lager/lagertest"

	"cred-alert/webhook"
	"cred-alert/webhook/webhookfakes"
)

var _ = Describe("Webhook", func() {
	var (
		logger *lagertest.TestLogger

		handler  http.Handler
		ingestor *webhookfakes.FakeIngestor

		fakeRequest *http.Request
		recorder    *httptest.ResponseRecorder

		token string
	)

	BeforeEach(func() {
		logger = lagertest.NewTestLogger("webhook")
		recorder = httptest.NewRecorder()
		ingestor = &webhookfakes.FakeIngestor{}
		token = "example-key"

		handler = webhook.Handler(logger, ingestor, token)
	})

	pushEvent := github.PushEvent{
		Before: github.String("abc123bef04e"),
		After:  github.String("def456af4e4"),
		Repo: &github.PushEventRepository{
			Name:     github.String("repository-name"),
			FullName: github.String("repository-owner/repository-name"),
			Owner: &github.PushEventRepoOwner{
				Name: github.String("repository-owner"),
			},
		},
		Commits: []github.PushEventCommit{
			{ID: github.String("commit-sha-1")},
			{ID: github.String("commit-sha-2")},
			{ID: github.String("commit-sha-3")},
			{ID: github.String("commit-sha-4")},
			{ID: github.String("commit-sha-5")},
		},
	}

	Context("when the request is properly formed", func() {
		BeforeEach(func() {
			body := &bytes.Buffer{}
			err := json.NewEncoder(body).Encode(pushEvent)
			Expect(err).NotTo(HaveOccurred())

			macHeader := fmt.Sprintf("sha1=%s", messageMAC(token, body.Bytes()))

			fakeRequest, _ = http.NewRequest("POST", "http://example.com/webhook", body)
			fakeRequest.Header.Set("X-Hub-Signature", macHeader)
		})

		It("responds with 200", func() {
			handler.ServeHTTP(recorder, fakeRequest)

			Expect(recorder.Code).To(Equal(http.StatusOK))
		})

		It("handles and scans the event directly", func() {
			handler.ServeHTTP(recorder, fakeRequest)

			Eventually(ingestor.IngestPushScanCallCount).Should(Equal(1))
		})
	})

	Context("when the signature is invalid", func() {
		BeforeEach(func() {
			body := &bytes.Buffer{}
			err := json.NewEncoder(body).Encode(pushEvent)
			Expect(err).NotTo(HaveOccurred())

			fakeRequest, _ = http.NewRequest("POST", "http://example.com/webhook", body)
			fakeRequest.Header.Set("X-Hub-Signature", "thisaintnohmacsignature")
		})

		It("responds with 403", func() {
			handler.ServeHTTP(recorder, fakeRequest)

			Expect(recorder.Code).To(Equal(http.StatusForbidden))
		})

		It("does not directly handle the event", func() {
			handler.ServeHTTP(recorder, fakeRequest)

			Consistently(ingestor.IngestPushScanCallCount()).Should(BeZero())
		})
	})

	Context("when the payload is not valid JSON", func() {
		BeforeEach(func() {
			badJSON := bytes.NewBufferString("{'ooops:---")

			fakeRequest, _ = http.NewRequest("POST", "http://example.com/webhook", badJSON)
			fakeRequest.Header.Set("X-Hub-Signature", fmt.Sprintf("sha1=%s", messageMAC(token, badJSON.Bytes())))
		})

		It("responds with 400", func() {
			handler.ServeHTTP(recorder, fakeRequest)

			Expect(recorder.Code).To(Equal(http.StatusBadRequest))
		})

		It("does not directly handle the event", func() {
			handler.ServeHTTP(recorder, fakeRequest)

			Consistently(ingestor.IngestPushScanCallCount).Should(BeZero())
		})
	})
})

func messageMAC(key string, body []byte) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write(body)
	return fmt.Sprintf("%x", mac.Sum(nil))
}
