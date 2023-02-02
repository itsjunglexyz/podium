package service_test

import (
	"context"
	"fmt"
	"github.com/itsjunglexyz/podium/leaderboard/v2/database"

	"github.com/golang/mock/gomock"
	"github.com/itsjunglexyz/podium/leaderboard/v2/service"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service TotalPages", func() {
	var ctrl *gomock.Controller
	var mock *database.MockDatabase
	var svc *service.Service

	var leaderboard string = "leaderboardTest"
	var totalMembers int = 11
	var pageSize int = 10
	var pageSize2 int = 3

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mock = database.NewMockDatabase(ctrl)

		svc = &service.Service{mock}
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("Should return number of pages if all is OK", func() {
		mock.EXPECT().GetTotalMembers(gomock.Any(), gomock.Eq(leaderboard)).Return(totalMembers, nil)
		mock.EXPECT().GetTotalMembers(gomock.Any(), gomock.Eq(leaderboard)).Return(totalMembers, nil)

		pages, err := svc.TotalPages(context.Background(), leaderboard, pageSize)
		Expect(err).NotTo(HaveOccurred())

		calculatedPages := 2

		Expect(pages).To(Equal(calculatedPages))

		pages, err = svc.TotalPages(context.Background(), leaderboard, pageSize2)
		Expect(err).NotTo(HaveOccurred())

		calculatedPages = 4

		Expect(pages).To(Equal(calculatedPages))
	})

	It("Should return error if database return in error", func() {
		mock.EXPECT().GetTotalMembers(gomock.Any(), gomock.Eq(leaderboard)).Return(-1, fmt.Errorf("Database error example"))

		_, err := svc.TotalPages(context.Background(), leaderboard, pageSize)
		Expect(err).To(Equal(service.NewGeneralError("total pages", "Database error example")))
	})
})
