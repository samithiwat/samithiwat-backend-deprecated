package icon

import (
	"github.com/gofiber/fiber/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/samithiwat/samithiwat-backend/src/service"
	"gorm.io/gorm"
	"time"
)

var mockIcons []*model.Icon

type TestDB struct{}

func (TestDB) FindAll(icons *[]*model.Icon) error {
	*icons = mockIcons
	return nil
}

func (TestDB) FindOne(id int64, icon *model.Icon) error {
	if int(id) > len(mockIcons) || int(id) < 1 {
		return fiber.ErrNotFound
	}

	*icon = *mockIcons[id-1]
	return nil
}

func (TestDB) Create(icon *model.Icon) error {
	createdDate := time.Date(2022, time.Month(time.March), 7, 19, 16, 52, 0, time.UTC)
	icon.ID = int64(len(mockIcons)) + 1
	icon.CreatedAt = createdDate
	icon.UpdatedAt = createdDate
	mockIcons = append(mockIcons, icon)
	return nil
}

func (TestDB) Update(id int64, icon *model.Icon) error {
	updatedDate := time.Date(2022, time.Month(time.March), 7, 19, 50, 12, 0, time.UTC)

	if int(id) > len(mockIcons) || int(id) < 1 {
		return fiber.ErrNotFound
	}

	existed := mockIcons[id-1]
	existed.Name = icon.Name
	existed.BgColor = icon.BgColor
	existed.IconType = icon.IconType
	existed.OwnerID = icon.OwnerID
	existed.OwnerType = icon.OwnerType
	existed.UpdatedAt = updatedDate

	*icon = *existed

	return nil
}

func (TestDB) Delete(id int64, icon *model.Icon) error {
	deletedDate := time.Date(2022, time.Month(time.March), 7, 20, 1, 12, 0, time.UTC)

	if int(id) > len(mockIcons) || int(id) < 1 {
		return fiber.ErrNotFound
	}

	*icon = *mockIcons[id-1]
	mockIcons[id-1].DeletedAt = gorm.DeletedAt{Time: deletedDate}

	return nil
}

var _ = Describe("Icon Service", func() {
	var iconService Service
	var icon1 *model.Icon
	var icon2 *model.Icon
	var icon3 *model.Icon

	BeforeEach(func() {
		validator := service.NewValidatorService()
		iconService = NewIconService(TestDB{}, validator)
	})

	Describe("Find All Icon", func() {
		When("Not Empty", func() {
			BeforeEach(func() {
				icon1 = &model.Icon{
					ID:        1,
					Name:      "A",
					BgColor:   "#ffffff",
					IconType:  "icon",
					OwnerID:   0,
					OwnerType: "",
				}

				icon2 = &model.Icon{
					ID:        2,
					Name:      "B",
					BgColor:   "#ffffff",
					IconType:  "svg",
					OwnerID:   0,
					OwnerType: "",
				}

				icon3 = &model.Icon{
					ID:        3,
					Name:      "C",
					BgColor:   "#000000",
					IconType:  "svg",
					OwnerID:   0,
					OwnerType: "",
				}

				mockIcons = append(mockIcons, icon1, icon2, icon3)
			})

			AfterEach(func() {
				mockIcons = []*model.Icon{}
			})

			It("Should get 3 icons", func() {
				icons, err := iconService.FindAll()
				Expect(err).ShouldNot(HaveOccurred())

				var want []*model.Icon
				want = append(want, icon1, icon2, icon3)

				for i, icon := range *icons {
					Expect(icon).Should(Equal(want[i]))
				}
			})
		})

		When("Empty", func() {
			BeforeEach(func() {
				mockIcons = []*model.Icon{}
			})

			AfterEach(func() {
				mockIcons = []*model.Icon{}
			})

			It("Should get an empty slice", func() {
				icons, err := iconService.FindAll()
				Expect(err).ShouldNot(HaveOccurred())

				var want []*model.Icon

				Expect(len(*icons)).Should(Equal(len(want)))
			})
		})
	})

	Describe("Find One Icon", func() {
		When("Not Empty", func() {
			BeforeEach(func() {
				icon1 = &model.Icon{
					ID:        1,
					Name:      "Facebook",
					BgColor:   "#ffffff",
					IconType:  "svg",
					OwnerID:   1,
					OwnerType: "badges",
				}

				icon2 = &model.Icon{
					ID:        2,
					Name:      "Google",
					BgColor:   "#ffffff",
					IconType:  "icon",
					OwnerID:   1,
					OwnerType: "badges",
				}

				icon3 = &model.Icon{
					ID:        3,
					Name:      "Instagram",
					BgColor:   "#000000",
					IconType:  "icon",
					OwnerID:   1,
					OwnerType: "badges",
				}

				mockIcons = append(mockIcons, icon1, icon2, icon3)
			})

			AfterEach(func() {
				mockIcons = []*model.Icon{}
			})

			It("Should get facebook icon", func() {
				icon, err := iconService.FindOne(1)
				Expect(err).ShouldNot(HaveOccurred())

				want := icon1

				Expect(icon).Should(Equal(want))
			})

			It("Should not found", func() {
				icon, err := iconService.FindOne(5)

				want := fiber.ErrNotFound

				Expect(icon).Should(BeNil())
				Expect(want.Code).Should(Equal(err.(*fiber.Error).Code))
			})
		})

		When("Empty", func() {
			BeforeEach(func() {
				mockIcons = []*model.Icon{}
			})

			AfterEach(func() {
				mockIcons = []*model.Icon{}
			})

			It("Should not found", func() {
				icon, err := iconService.FindOne(1)

				want := fiber.ErrNotFound

				Expect(icon).Should(BeNil())
				Expect(want.Code).Should(Equal(err.(*fiber.Error).Code))
			})
		})
	})

	Describe("Create Icon", func() {

		BeforeEach(func() {
			icon1 = &model.Icon{
				ID:        1,
				Name:      "Facebook",
				BgColor:   "#ffffff",
				IconType:  "svg",
				OwnerID:   1,
				OwnerType: "badges",
			}

			icon2 = &model.Icon{
				ID:        2,
				Name:      "Google",
				BgColor:   "#ffffff",
				IconType:  "icon",
				OwnerID:   1,
				OwnerType: "badges",
			}

			icon3 = &model.Icon{
				ID:        3,
				Name:      "Instagram",
				BgColor:   "#000000",
				IconType:  "svg",
				OwnerID:   1,
				OwnerType: "badges",
			}

			mockIcons = append(mockIcons, icon1, icon2, icon3)
		})

		AfterEach(func() {
			mockIcons = []*model.Icon{}
		})

		It("Should create netflix icon", func() {
			iconDto := &model.NewIcon{
				Name:      "Netflix",
				BgColor:   "#000000",
				IconType:  "icon",
				OwnerID:   0,
				OwnerType: "",
			}

			icon, err := iconService.Create(*iconDto)
			Expect(err).ShouldNot(HaveOccurred())

			createdDate := time.Date(2022, time.Month(time.March), 7, 19, 16, 52, 0, time.UTC)
			want := model.Icon{
				ID:        4,
				Name:      "Netflix",
				BgColor:   "#000000",
				IconType:  "icon",
				OwnerID:   0,
				OwnerType: "",
				CreatedAt: createdDate,
				UpdatedAt: createdDate,
			}

			Expect(icon.ID).Should(Equal(want.ID))
			Expect(icon.Name).Should(Equal(want.Name))
			Expect(icon.BgColor).Should(Equal(want.BgColor))
			Expect(icon.IconType).Should(Equal(want.IconType))
			Expect(icon.OwnerID).Should(Equal(want.OwnerID))
			Expect(icon.OwnerType).Should(Equal(want.OwnerType))
			Expect(icon.CreatedAt).Should(Equal(want.CreatedAt))
			Expect(icon.UpdatedAt).Should(Equal(want.UpdatedAt))
		})

	})

	Describe("Update Icon", func() {

		BeforeEach(func() {
			icon1 = &model.Icon{
				ID:        1,
				Name:      "Facebook",
				BgColor:   "#ffffff",
				IconType:  "svg",
				OwnerID:   1,
				OwnerType: "badges",
				CreatedAt: time.Date(2022, time.Month(time.March), 7, 19, 16, 52, 0, time.UTC),
				UpdatedAt: time.Date(2022, time.Month(time.March), 7, 19, 50, 12, 0, time.UTC),
			}

			icon2 = &model.Icon{
				ID:        2,
				Name:      "Google",
				BgColor:   "#ffffff",
				IconType:  "icon",
				OwnerID:   1,
				OwnerType: "badges",
				CreatedAt: time.Date(2022, time.Month(time.March), 7, 19, 16, 52, 0, time.UTC),
				UpdatedAt: time.Date(2022, time.Month(time.March), 7, 19, 50, 12, 0, time.UTC),
			}

			icon3 = &model.Icon{
				ID:        3,
				Name:      "Instagram",
				BgColor:   "#000000",
				IconType:  "svg",
				OwnerID:   1,
				OwnerType: "badges",
				CreatedAt: time.Date(2022, time.Month(time.March), 7, 19, 16, 52, 0, time.UTC),
				UpdatedAt: time.Date(2022, time.Month(time.March), 7, 19, 50, 12, 0, time.UTC),
			}

			mockIcons = append(mockIcons, icon1, icon2, icon3)
		})

		AfterEach(func() {
			mockIcons = []*model.Icon{}
		})

		It("Should update facebook to netflix", func() {
			iconDto := &model.NewIcon{
				Name:      "Netflix",
				BgColor:   "#000000",
				IconType:  "icon",
				OwnerID:   0,
				OwnerType: "",
			}

			icon, err := iconService.Update(1, *iconDto)
			Expect(err).ShouldNot(HaveOccurred())

			createdDate := time.Date(2022, time.Month(time.March), 7, 19, 16, 52, 0, time.UTC)
			updatedDate := time.Date(2022, time.Month(time.March), 7, 19, 50, 12, 0, time.UTC)
			want := model.Icon{
				ID:        1,
				Name:      "Netflix",
				BgColor:   "#000000",
				IconType:  "icon",
				OwnerID:   0,
				OwnerType: "",
				CreatedAt: createdDate,
				UpdatedAt: updatedDate,
			}

			Expect(icon.ID).Should(Equal(want.ID))
			Expect(icon.Name).Should(Equal(want.Name))
			Expect(icon.BgColor).Should(Equal(want.BgColor))
			Expect(icon.IconType).Should(Equal(want.IconType))
			Expect(icon.OwnerID).Should(Equal(want.OwnerID))
			Expect(icon.OwnerType).Should(Equal(want.OwnerType))
			Expect(icon.CreatedAt).Should(Equal(want.CreatedAt))
			Expect(icon.UpdatedAt).Should(Equal(want.UpdatedAt))
		})

		It("Should not found", func() {
			iconDto := &model.NewIcon{
				Name:      "Netflix",
				BgColor:   "#000000",
				IconType:  "icon",
				OwnerID:   0,
				OwnerType: "",
			}

			icon, err := iconService.Update(8, *iconDto)

			want := fiber.ErrNotFound

			Expect(icon).Should(BeNil())
			Expect(want.Code).Should(Equal(err.(*fiber.Error).Code))
		})

	})

	Describe("Delete Icon", func() {

		BeforeEach(func() {
			icon1 = &model.Icon{
				ID:        1,
				Name:      "Facebook",
				BgColor:   "#ffffff",
				IconType:  "svg",
				OwnerID:   1,
				OwnerType: "badges",
				CreatedAt: time.Date(2022, time.Month(time.March), 7, 19, 16, 52, 0, time.UTC),
				UpdatedAt: time.Date(2022, time.Month(time.March), 7, 19, 50, 12, 0, time.UTC),
			}

			icon2 = &model.Icon{
				ID:        2,
				Name:      "Google",
				BgColor:   "#ffffff",
				IconType:  "icon",
				OwnerID:   1,
				OwnerType: "badges",
				CreatedAt: time.Date(2022, time.Month(time.March), 7, 19, 16, 52, 0, time.UTC),
				UpdatedAt: time.Date(2022, time.Month(time.March), 7, 19, 16, 52, 0, time.UTC),
			}

			icon3 = &model.Icon{
				ID:        3,
				Name:      "Instagram",
				BgColor:   "#000000",
				IconType:  "svg",
				OwnerID:   1,
				OwnerType: "badges",
				CreatedAt: time.Date(2022, time.Month(time.March), 7, 19, 16, 52, 0, time.UTC),
				UpdatedAt: time.Date(2022, time.Month(time.March), 7, 19, 50, 12, 0, time.UTC),
			}

			mockIcons = append(mockIcons, icon1, icon2, icon3)
		})

		AfterEach(func() {
			mockIcons = []*model.Icon{}
		})

		It("Should delete google icon", func() {
			icon, err := iconService.Delete(2)
			Expect(err).ShouldNot(HaveOccurred())

			createdDate := time.Date(2022, time.Month(time.March), 7, 19, 16, 52, 0, time.UTC)
			want := model.Icon{
				ID:        2,
				Name:      "Google",
				BgColor:   "#ffffff",
				IconType:  "icon",
				OwnerID:   1,
				OwnerType: "badges",
				CreatedAt: createdDate,
				UpdatedAt: createdDate,
			}

			Expect(icon.ID).Should(Equal(want.ID))
			Expect(icon.Name).Should(Equal(want.Name))
			Expect(icon.BgColor).Should(Equal(want.BgColor))
			Expect(icon.IconType).Should(Equal(want.IconType))
			Expect(icon.OwnerID).Should(Equal(want.OwnerID))
			Expect(icon.OwnerType).Should(Equal(want.OwnerType))
			Expect(icon.CreatedAt).Should(Equal(want.CreatedAt))
			Expect(icon.UpdatedAt).Should(Equal(want.UpdatedAt))
		})

		It("Should not found", func() {
			icon, err := iconService.Delete(100)

			want := fiber.ErrNotFound

			Expect(icon).Should(BeNil())
			Expect(want.Code).Should(Equal(err.(*fiber.Error).Code))
		})

	})

})
