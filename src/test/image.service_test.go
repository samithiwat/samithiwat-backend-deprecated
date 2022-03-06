package test_test

import (
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/samithiwat/samithiwat-backend/src/database"
	"github.com/samithiwat/samithiwat-backend/src/graph/model"
	service "github.com/samithiwat/samithiwat-backend/src/graph/services"
	"regexp"
	"time"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

var _ = Describe("Image.Service", func() {
	var imageService service.ImageService
	var mock sqlmock.Sqlmock

	BeforeEach(func() {
		var db database.Database
		var err error

		db, mock, err = database.MockDatabase()
		Expect(err).ShouldNot(HaveOccurred())

		validateService := service.NewValidatorService()
		imageService = service.NewImageService(db, validateService)
	})

	AfterEach(func() {
		err := mock.ExpectationsWereMet()
		Expect(err).ShouldNot(HaveOccurred())
	})

	Context("get All()", func() {
		It("empty", func() {
			const sqlSelectAll = `SELECT * FROM "images" WHERE "images"."deleted_at" IS NULL`

			mock.ExpectQuery(regexp.QuoteMeta(sqlSelectAll)).WillReturnRows(sqlmock.NewRows(nil))

			images, err := imageService.GetAll()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(images).Should(BeEmpty())
		})

		It("return all", func() {
			const sqlSelectAll = `SELECT * FROM "images" WHERE "images"."deleted_at" IS NULL`

			image1 := model.Image{
				ID:          1,
				Name:        "A",
				Description: "Image A",
				ImgUrl:      "https://imgurl.com/A",
			}

			image2 := model.Image{
				ID:          2,
				Name:        "B",
				Description: "Image B",
				ImgUrl:      "https://imgurl.com/B",
			}

			image3 := model.Image{
				ID:          3,
				Name:        "C",
				Description: "Image C",
				ImgUrl:      "https://imgurl.com/C",
			}
			var images []*model.Image

			images = append(images, &image1, &image2, &image3)

			rows := sqlmock.
				NewRows([]string{"id", "name", "description", "img_url"}).
				AddRow(image1.ID, image1.Name, image1.Description, image1.ImgUrl).
				AddRow(image2.ID, image2.Name, image2.Description, image2.ImgUrl).
				AddRow(image3.ID, image3.Name, image3.Description, image3.ImgUrl)

			mock.ExpectQuery(regexp.QuoteMeta(sqlSelectAll)).WillReturnRows(rows)

			queriedImaged, err := imageService.GetAll()
			Expect(err).ShouldNot(HaveOccurred())

			for i, image := range queriedImaged {
				Expect(image).Should(Equal(images[i]))
			}
		})
	})

	Context("getOne()", func() {
		It("found", func() {
			const sqlSelectOne = `SELECT * FROM "images" WHERE "images"."id" = $1 AND "images"."deleted_at" IS NULL ORDER BY "images"."id" LIMIT 1 `

			image1 := &model.Image{
				ID:          1,
				Name:        "A",
				Description: "Image A",
				ImgUrl:      "https://imgurl.com/A",
			}

			rows := sqlmock.
				NewRows([]string{"id", "name", "description", "img_url"}).
				AddRow(image1.ID, image1.Name, image1.Description, image1.ImgUrl)

			mock.ExpectQuery(regexp.QuoteMeta(sqlSelectOne)).WithArgs(image1.ID).WillReturnRows(rows)

			queriedImage, err := imageService.GetOne(image1.ID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(queriedImage).Should(Equal(image1))
		})

		It("Not found", func() {
			mock.ExpectQuery(`.+`).WillReturnRows(sqlmock.NewRows(nil))

			_, err := imageService.GetOne(1)
			Expect(err).Should(Equal(fiber.NewError(fiber.StatusNotFound, "Not found")))
		})
	})

	Context("save", func() {
		var image *model.Image
		var imageDto *model.NewImage

		BeforeEach(func() {
			imageDto = &model.NewImage{
				Name:        "A",
				Description: "Image A",
				ImgURL:      "https://imgurl.com/A",
			}

			image = &model.Image{
				Name:        imageDto.Name,
				Description: imageDto.Description,
				ImgUrl:      imageDto.ImgURL,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}
		})

		It("insert", func() {
			const sqlInsert = `INSERT INTO "images" ("name","description","img_url","owner_type","created_at","updated_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "owner_id","id"`

			const newId = 1
			mock.ExpectBegin()
			mock.ExpectQuery(regexp.QuoteMeta(sqlInsert)).
				WithArgs(image.Name, image.Description, image.ImgUrl, "", AnyTime{}, AnyTime{}, nil).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(newId))
			mock.ExpectCommit()

			Expect(image.ID).Should(BeZero())

			image, err := imageService.Create(imageDto)
			Expect(err).ShouldNot(HaveOccurred())

			Expect(image.ID).Should(BeEquivalentTo(newId))
		})

	})
})
