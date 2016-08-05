package service_test

import (
	"github.com/cloudfoundry/cli/cf/api/apifakes"
	"github.com/cloudfoundry/cli/cf/commandregistry"
	"github.com/cloudfoundry/cli/cf/configuration/coreconfig"
	"github.com/cloudfoundry/cli/cf/models"
	"github.com/cloudfoundry/cli/cf/requirements"
	"github.com/cloudfoundry/cli/cf/requirements/requirementsfakes"
	testcmd "github.com/cloudfoundry/cli/testhelpers/commands"
	testconfig "github.com/cloudfoundry/cli/testhelpers/configuration"
	testterm "github.com/cloudfoundry/cli/testhelpers/terminal"

	. "github.com/cloudfoundry/cli/testhelpers/matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("rename-service command", func() {
	var (
		ui                  *testterm.FakeUI
		config              coreconfig.Repository
		serviceRepo         *apifakes.FakeServiceRepository
		requirementsFactory *requirementsfakes.FakeFactory
		deps                commandregistry.Dependency

		serviceInstance models.ServiceInstance
	)

	updateCommandDependency := func(pluginCall bool) {
		deps.UI = ui
		deps.RepoLocator = deps.RepoLocator.SetServiceRepository(serviceRepo)
		deps.Config = config
		commandregistry.Commands.SetCommand(commandregistry.Commands.FindCommand("rename-service").SetDependency(deps, pluginCall))
	}

	BeforeEach(func() {
		ui = &testterm.FakeUI{}
		config = testconfig.NewRepositoryWithDefaults()
		serviceRepo = new(apifakes.FakeServiceRepository)
		requirementsFactory = new(requirementsfakes.FakeFactory)

		serviceInstance = models.ServiceInstance{}
		serviceInstance.Name = "different-name"
		serviceInstance.GUID = "different-name-guid"
		serviceReq := new(requirementsfakes.FakeServiceInstanceRequirement)
		serviceReq.GetServiceInstanceReturns(serviceInstance)
		requirementsFactory.NewServiceInstanceRequirementReturns(serviceReq)
	})

	runCommand := func(args ...string) bool {
		return testcmd.RunCLICommand("rename-service", args, requirementsFactory, updateCommandDependency, false, ui)
	}

	Describe("requirements", func() {
		It("Fails with usage when exactly two parameters not passed", func() {
			runCommand("whatever")
			Expect(ui.Outputs()).To(ContainSubstrings(
				[]string{"Incorrect Usage", "Requires", "arguments"},
			))
		})

		It("fails when not logged in", func() {
			requirementsFactory.NewTargetedSpaceRequirementReturns(requirements.Passing{})
			requirementsFactory.NewLoginRequirementReturns(requirements.Failing{Message: "not logged in"})
			Expect(runCommand("banana", "fppants")).To(BeFalse())
		})

		It("fails when a space is not targeted", func() {
			requirementsFactory.NewLoginRequirementReturns(requirements.Passing{})
			requirementsFactory.NewTargetedSpaceRequirementReturns(requirements.Failing{Message: "not targeting space"})

			Expect(runCommand("banana", "faaaaasdf")).To(BeFalse())
		})
	})

	Context("when logged in and a space is targeted", func() {
		BeforeEach(func() {
			requirementsFactory.NewLoginRequirementReturns(requirements.Passing{})
			requirementsFactory.NewTargetedSpaceRequirementReturns(requirements.Passing{})
		})

		It("renames the service, obviously", func() {
			runCommand("my-service", "new-name")

			Expect(ui.Outputs()).To(ContainSubstrings(
				[]string{"Renaming service", "different-name", "new-name", "my-org", "my-space", "my-user"},
				[]string{"OK"},
			))

			actualServiceInstance, actualServiceName := serviceRepo.RenameServiceArgsForCall(0)
			Expect(actualServiceInstance).To(Equal(serviceInstance))
			Expect(actualServiceName).To(Equal("new-name"))
		})
	})
})
