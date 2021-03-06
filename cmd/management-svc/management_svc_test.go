package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/go-chi/chi"
	"github.com/kelseyhightower/envconfig"
	management_server "github.com/library/cmd/management-svc/management-server"
	data_store "github.com/library/data-store"
	"github.com/library/envConfig"
	"github.com/library/middleware"
	"github.com/library/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Management-Service", func() {
	var (
		r          *chi.Mux
		adminToken string
		userToken  string
		err        error
	)

	BeforeSuite(func() {
		env = &envConfig.Env{}
		err = envconfig.Process("LIBRARY", env)
		Expect(err).To(BeNil())
		dataStore = data_store.DbConnect(env, true)
		srv = management_server.NewServer(env, dataStore, nil)
		Expect(err).To(BeNil())
		srv.TestRun = true
		middleware.SetJwtSigningKey(srv.Env.JwtSigningKey)
		adminToken, userToken, err = setupAuthInfo(env, dataStore.Db)
		Expect(err).To(BeNil())
		r = management_server.SetupRouter(srv, nil)
		err = setupTestData(dataStore.Db)
		Expect(err).To(BeNil())
	})

	Describe("Handlers test", func() {
		Describe("Check availability", func() {
			It("Should return the availability of specified book", func() {
				req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/user/check-availability/%s", "1010"), nil)
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", "Bearer "+userToken)
				rec := httptest.NewRecorder()
				r.ServeHTTP(rec, req)
				resp := rec.Result()
				Expect(resp.StatusCode).To(BeEquivalentTo(http.StatusOK))
				var avail bool
				err = json.NewDecoder(resp.Body).Decode(&avail)
				Expect(err).To(BeNil())
				Expect(avail).To(BeEquivalentTo(true))
			})
		})

		Describe("Issue book", func() {
			It("Should issue book to specified user", func() {
				formData := url.Values{
					"userId": {"1010"},
					"bookId": {"1010"},
				}
				req := httptest.NewRequest(http.MethodPost, "/admin/issue-book", strings.NewReader(formData.Encode()))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				req.Header.Set("Authorization", "Bearer "+adminToken)

				rec := httptest.NewRecorder()
				r.ServeHTTP(rec, req)
				resp := rec.Result()
				Expect(resp.StatusCode).To(BeEquivalentTo(http.StatusOK))
				var book models.Book
				err = dataStore.Db.Where("id = ?", "1010").Find(&book).Error
				Expect(err).To(BeNil())
			})
		})

		Describe("Get History", func() {
			It("Should return the complete book issue history", func() {
				req := httptest.NewRequest(http.MethodGet, "/admin/get-history/1010", nil)
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", "Bearer "+adminToken)
				rec := httptest.NewRecorder()
				r.ServeHTTP(rec, req)
				resp := rec.Result()
				Expect(resp.StatusCode).To(BeEquivalentTo(http.StatusOK))
				var history []map[string]interface{}
				err = json.NewDecoder(resp.Body).Decode(&history)
				Expect(err).To(BeNil())
				Expect(history[0]["bookId"]).To(BeEquivalentTo(1010))
			})
		})

		Describe("Return Book", func() {
			It("Should change the availability status of that book", func() {
				req := httptest.NewRequest(http.MethodGet, "/admin/return-book/1010", nil)
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", "Bearer "+adminToken)
				rec := httptest.NewRecorder()
				r.ServeHTTP(rec, req)
				resp := rec.Result()
				Expect(resp.StatusCode).To(BeEquivalentTo(http.StatusOK))
				var book models.Book
				err = dataStore.Db.Where("id = ?", "1010").Find(&book).Error
				Expect(err).To(BeNil())
			})
		})
	})

	AfterSuite(func() {
		err = cleanTestData(dataStore.Db)
		Expect(err).To(BeNil())
	})
})
