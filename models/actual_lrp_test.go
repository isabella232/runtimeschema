package models_test

import (
	. "github.com/cloudfoundry-incubator/runtime-schema/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ActualLRP", func() {
	var lrp ActualLRP

	lrpPayload := `{
    "process_guid":"some-guid",
    "instance_guid":"some-instance-guid",
    "host": "1.2.3.4",
    "ports": [
      { "container_port": 8080 },
      { "container_port": 8081, "host_port": 1234 }
    ],
    "index": 2,
    "state": 0,
    "since": 1138,
    "executor_id":"some-executor-id",
    "domain":"some-domain"
  }`

	BeforeEach(func() {
		lrp = ActualLRP{
			ProcessGuid:  "some-guid",
			InstanceGuid: "some-instance-guid",
			Host:         "1.2.3.4",
			Ports: []PortMapping{
				{ContainerPort: 8080},
				{ContainerPort: 8081, HostPort: 1234},
			},
			Index:      2,
			Since:      1138,
			ExecutorID: "some-executor-id",
			Domain:     "some-domain",
		}
	})

	Describe("ToJSON", func() {
		It("should JSONify", func() {
			json := lrp.ToJSON()
			Ω(string(json)).Should(MatchJSON(lrpPayload))
		})
	})

	Describe("NewActualLRPFromJSON", func() {
		It("returns a LRP with correct fields", func() {
			decodedStartAuction, err := NewActualLRPFromJSON([]byte(lrpPayload))
			Ω(err).ShouldNot(HaveOccurred())

			Ω(decodedStartAuction).Should(Equal(lrp))
		})

		Context("with an invalid payload", func() {
			It("returns the error", func() {
				decodedStartAuction, err := NewActualLRPFromJSON([]byte("something lol"))
				Ω(err).Should(HaveOccurred())

				Ω(decodedStartAuction).Should(BeZero())
			})
		})

		for field, payload := range map[string]string{
			"process_guid":  `{"instance_guid": "instance_guid", "executor_id": "executor_id", "domain": "domain"}`,
			"instance_guid": `{"process_guid": "process-guid", "executor_id": "executor_id", "domain": "domain"}`,
			"executor_id":   `{"process_guid": "process-guid", "instance_guid": "instance_guid", "domain": "domain"}`,
			"domain":        `{"process_guid": "process-guid", "executor_id": "executor_id", "instance_guid": "instance_guid"}`,
		} {
			missingField := field
			json := payload

			Context("when the json is missing a "+missingField, func() {
				It("returns an error indicating so", func() {
					decodedStartAuction, err := NewActualLRPFromJSON([]byte(json))
					Ω(err).Should(HaveOccurred())
					Ω(err.Error()).Should(Equal("JSON has missing/invalid field: " + missingField))

					Ω(decodedStartAuction).Should(BeZero())
				})
			})
		}
	})

	Describe("NewActualLRP", func() {
		It("returns a LRP with correct fields", func() {
			actualLrp, err := NewActualLRP(
				"processGuid",
				"instanceGuid",
				"executorID",
				"domain",
				0,
			)
			Ω(err).ShouldNot(HaveOccurred())

			Ω(actualLrp.ProcessGuid).Should(Equal("processGuid"))
			Ω(actualLrp.InstanceGuid).Should(Equal("instanceGuid"))
			Ω(actualLrp.ExecutorID).Should(Equal("executorID"))
			Ω(actualLrp.Domain).Should(Equal("domain"))
			Ω(actualLrp.Index).Should(BeZero())
		})

		Context("When given a blank process guid", func() {
			It("returns an error indicating so", func() {
				_, err := NewActualLRP("", "instanceGuid", "executorID", "domain", 0)
				Ω(err).Should(HaveOccurred())
				Ω(err.Error()).Should(Equal("Cannot construct Actual LRP with empty process guid"))
			})
		})

		Context("When given a blank instance guid", func() {
			It("returns an error indicating so", func() {
				_, err := NewActualLRP("processGuid", "", "executorID", "domain", 0)
				Ω(err).Should(HaveOccurred())
				Ω(err.Error()).Should(Equal("Cannot construct Actual LRP with empty instance guid"))
			})
		})

		Context("When given a blank executor ID", func() {
			It("returns an error indicating so", func() {
				_, err := NewActualLRP("processGuid", "instanceGuid", "", "domain", 0)
				Ω(err).Should(HaveOccurred())
				Ω(err.Error()).Should(Equal("Cannot construct Actual LRP with empty executor ID"))
			})
		})

		Context("When given a blank domain", func() {
			It("returns an error indicating so", func() {
				_, err := NewActualLRP("processGuid", "instanceGuid", "executorID", "", 0)
				Ω(err).Should(HaveOccurred())
				Ω(err.Error()).Should(Equal("Cannot construct Actual LRP with empty domain"))
			})
		})
	})
})
