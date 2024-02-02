package ran_du_system_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/openshift-kni/eco-goinfra/pkg/polarion"
	"github.com/openshift-kni/eco-gosystem/tests/ran-du/internal/find"
	. "github.com/openshift-kni/eco-gosystem/tests/ran-du/internal/randuinittools"
)

var _ = Describe(
	"PTPLogs",
	Label("PTPLogs"),
	Ordered,
	ContinueOnFailure,
	func() {
		It("Clock Jumps", polarion.ID("OCP-48755"), Label("ClockJumps"), func() {
			pod, err := find.LinuxPTPDaemonPod(APIClient)
			Expect(err).ToNot(HaveOccurred())
			logs, err := pod.GetFullLog("linuxptp-daemon-container")
			Expect(err).ToNot(HaveOccurred())
			Expect(logs).NotTo(ContainSubstring("jump"), "Found 'jump' in linuxptp-daemon-container container logs")
		})
	},
)
