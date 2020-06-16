package host

import (
	"context"
	"strings"
	"time"

	"github.com/filanov/bm-inventory/internal/common"

	"github.com/filanov/bm-inventory/internal/hardware"
	"github.com/filanov/bm-inventory/models"
	"github.com/go-openapi/strfmt"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
)

var _ = Describe("insufficient_state", func() {
	var (
		ctx           = context.Background()
		state         API
		db            *gorm.DB
		currentState  = HostStatusInsufficient
		host          models.Host
		id, clusterId strfmt.UUID
		updateReply   *UpdateReply
		updateErr     error
		expectedReply *expect
		ctrl          *gomock.Controller
		mockValidator *hardware.MockValidator
	)

	BeforeEach(func() {
		db = prepareDB()
		ctrl = gomock.NewController(GinkgoT())
		mockValidator = hardware.NewMockValidator(ctrl)
		state = &Manager{insufficient: NewInsufficientState(getTestLog(), db, mockValidator)}

		id = strfmt.UUID(uuid.New().String())
		clusterId = strfmt.UUID(uuid.New().String())
		host = getTestHost(id, clusterId, currentState)
		Expect(db.Create(&host).Error).ShouldNot(HaveOccurred())
		expectedReply = &expect{expectedState: currentState}
	})

	Context("update hw info", func() {
		It("update", func() {
			mockValidator.EXPECT().IsSufficient(gomock.Any()).
				Return(&common.IsSufficientReply{IsSufficient: true}, nil).AnyTimes()
			updateReply, updateErr = state.UpdateHwInfo(ctx, &host, "some hw info")
			updateReply, updateErr = state.RefreshStatus(ctx, &host)
			expectedReply.expectedState = HostStatusKnown
			expectedReply.postCheck = func() {
				h := getHost(id, clusterId, db)
				Expect(h.Inventory).Should(Equal(""))
				Expect(h.HardwareInfo).Should(Equal("some hw info"))
			}
		})
	})

	Context("update_inventory", func() {
		It("sufficient_hw", func() {
			mockValidator.EXPECT().IsSufficient(gomock.Any()).
				Return(&common.IsSufficientReply{IsSufficient: true}, nil).AnyTimes()
			updateReply, updateErr = state.UpdateInventory(ctx, &host, "some hw info")
			updateReply, updateErr = state.RefreshStatus(ctx, &host)
			expectedReply.expectedState = HostStatusKnown
			expectedReply.postCheck = func() {
				h := getHost(id, clusterId, db)
				Expect(h.HardwareInfo).Should(Equal(defaultHwInfo))
				Expect(h.Inventory).Should(Equal("some hw info"))
			}
		})
		It("insufficient_hw", func() {
			mockValidator.EXPECT().IsSufficient(gomock.Any()).
				Return(&common.IsSufficientReply{IsSufficient: false, Reason: "because"}, nil).AnyTimes()
			updateReply, updateErr = state.UpdateInventory(ctx, &host, "some hw info")
			updateReply, updateErr = state.RefreshStatus(ctx, &host)
			expectedReply.expectedState = HostStatusInsufficient
			expectedReply.postCheck = func() {
				h := getHost(id, clusterId, db)
				Expect(h.Inventory).Should(Equal("some hw info"))
				Expect(strings.Contains(*h.StatusInfo, "because")).To(Equal(true))
			}
		})
		It("hw_validation_error", func() {
			mockValidator.EXPECT().IsSufficient(gomock.Any()).
				Return(nil, errors.New("error")).AnyTimes()
			updateReply, updateErr = state.UpdateInventory(ctx, &host, "some hw info")
			updateReply, updateErr = state.RefreshStatus(ctx, &host)
			expectedReply.expectedState = HostStatusInsufficient
			expectedReply.postCheck = func() {
				h := getHost(id, clusterId, db)
				Expect(h.HardwareInfo).Should(Equal(defaultHwInfo))
			}
		})
	})

	Context("update_role", func() {
		It("sufficient_hw", func() {
			updateReply, updateErr = state.UpdateRole(ctx, &host, "master", nil)
			mockValidator.EXPECT().IsSufficient(gomock.Any()).
				Return(&common.IsSufficientReply{IsSufficient: true}, nil).AnyTimes()
			updateReply, updateErr = state.RefreshStatus(ctx, &host)
			expectedReply.expectedState = HostStatusKnown
			expectedReply.postCheck = func() {
				h := getHost(id, clusterId, db)
				Expect(h.Role).Should(Equal("master"))
			}
		})
		It("insufficient_hw", func() {
			mockValidator.EXPECT().IsSufficient(gomock.Any()).
				Return(&common.IsSufficientReply{IsSufficient: false, Reason: "because"}, nil).AnyTimes()
			updateReply, updateErr = state.UpdateRole(ctx, &host, "master", nil)
			updateReply, updateErr = state.RefreshStatus(ctx, &host)
			expectedReply.postCheck = func() {
				h := getHost(id, clusterId, db)
				Expect(h.Role).Should(Equal("master"))
				Expect(strings.Contains(*h.StatusInfo, "because")).To(Equal(true))
			}
		})
		It("hw_validation_error", func() {
			mockValidator.EXPECT().IsSufficient(gomock.Any()).
				Return(nil, errors.New("error")).AnyTimes()
			updateReply, updateErr = state.UpdateInventory(ctx, &host, "some hw info")
			updateReply, updateErr = state.UpdateRole(ctx, &host, "master", nil)
			updateReply, updateErr = state.RefreshStatus(ctx, &host)
			expectedReply.postCheck = func() {
				h := getHost(id, clusterId, db)
				Expect(h.Role).Should(Equal("master"))
			}
		})
		It("master_with_tx", func() {
			tx := db.Begin()
			Expect(tx.Error).ShouldNot(HaveOccurred())
			mockValidator.EXPECT().IsSufficient(gomock.Any()).
				Return(&common.IsSufficientReply{IsSufficient: false}, nil).AnyTimes()
			updateReply, updateErr = state.UpdateRole(ctx, &host, "master", tx)
			Expect(tx.Rollback().Error).ShouldNot(HaveOccurred())
			expectedReply.postCheck = func() {
				h := getHost(id, clusterId, db)
				Expect(h.Role).Should(Equal("worker"))
			}
		})
	})

	Context("refresh_status", func() {
		It("keep_alive", func() {
			host.CheckedInAt = strfmt.DateTime(time.Now().Add(-time.Minute))
			mockValidator.EXPECT().IsSufficient(gomock.Any()).
				Return(&common.IsSufficientReply{Type: "hardware", IsSufficient: true}, nil).AnyTimes()
			updateReply, updateErr = state.RefreshStatus(ctx, &host)
			expectedReply.expectedState = HostStatusKnown
		})
		It("keep_alive_timeout", func() {
			host.CheckedInAt = strfmt.DateTime(time.Now().Add(-time.Hour))
			mockValidator.EXPECT().IsSufficient(gomock.Any()).
				Return(&common.IsSufficientReply{Type: "hardware", IsSufficient: true}, nil).AnyTimes()
			updateReply, updateErr = state.RefreshStatus(ctx, &host)
			expectedReply.expectedState = HostStatusDisconnected
		})
	})

	It("install", func() {
		updateReply, updateErr = state.Install(ctx, &host, nil)
		expectedReply.expectError = true
	})

	It("enable_host", func() {
		updateReply, updateErr = state.EnableHost(ctx, &host)
	})

	It("disable_host", func() {
		updateReply, updateErr = state.DisableHost(ctx, &host)
		expectedReply.expectedState = HostStatusDisabled
	})

	AfterEach(func() {
		ctrl.Finish()
		postValidation(expectedReply, currentState, db, id, clusterId, updateReply, updateErr)

		// cleanup
		db.Close()
		expectedReply = nil
		updateReply = nil
		updateErr = nil
	})
})
