package controller

import (
	"io"
	"membuattasktodo/helpers"
	"membuattasktodo/model"
	"membuattasktodo/service"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	service service.Service
}

func NewController(service service.Service) Controller {
	return Controller{service}
}

func (c *Controller) UserController(ctx echo.Context) error {
	// Perform appropriate processing here
	user, err := c.service.GetUsers()

	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Berhasil Menampilkan Data User",
		"data":    user,
	})
}

func (c *Controller) GetAlltaskController(ctx echo.Context) error {
	Claims := helpers.ClaimToken(ctx)
	id := Claims.ID

	task, err := c.service.GetAlltask(id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Succesfully Get All Task Data",
		"data":    task,
	})
}

func (c *Controller) GetTaskById(ctx echo.Context) error {
	Claims := helpers.ClaimToken(ctx)
	id := Claims.ID
	taskId := ctx.Param("id")
	taskidtrue, err := strconv.Atoi(taskId)
	if err != nil {
		return err
	}

	task, err := c.service.GetTaskById(id, taskidtrue)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Success Get Task Detail",
		"data":    task,
	})
}
func (c *Controller) DeleteTasksController(ctx echo.Context) error {
	Id := ctx.Param("id")
	IdAsli, err := strconv.Atoi(Id)
	if err != nil {
		return err
	}

	err = c.service.DeleteTask(IdAsli)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Task deleted successfully",
	})
}

func (c *Controller) BulkDeleteTask(ctx echo.Context) error {
	Claims := helpers.ClaimToken(ctx)
	Id := Claims.ID

	var req model.BulkDeleteRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	taskIds := req.ID

	err := c.service.BulkDeleteTask(taskIds, Id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success delete multiple tasks",
	})
}

func (c *Controller) CreateTasksController(ctx echo.Context) error {
	var req model.TaskReq

	err := ctx.Bind(&req)
	if err != nil {
		return err
	}
	// Menerima file gambar dari form dengan nama "image"
	image, err := ctx.FormFile("image")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Tidak dapat memproses file gambar"})
	}

	// Buka file yang diunggah
	src, err := image.Open()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal membuka file gambar"})
	}
	defer src.Close()

	// Lokasi penyimpanan file gambar lokal
	uploadDir := "uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal membuat direktori penyimpanan"})
	}

	// Generate nama file unik
	dstPath := filepath.Join(uploadDir, image.Filename)

	// Membuka file tujuan untuk penyimpanan
	dst, err := os.Create(dstPath)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal membuat file gambar"})
	}
	defer dst.Close()

	// Salin isi file dari file asal ke file tujuan
	if _, err = io.Copy(dst, src); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal menyalin file gambar"})
	}

	// Membuat URL ke gambar yang diunggah
	imageURL := "http://localhost:8080/uploads/" + image.Filename

	Claims := helpers.ClaimToken(ctx)
	Id := Claims.ID

	// validate .......
	task, err := c.service.CreateTask(req, Id, imageURL)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"Message": "successfully created task",
		"task":    task,
	})
}

func (c *Controller) UpdateTaskController(ctx echo.Context) error {
	var req model.TaskReq

	err := ctx.Bind(&req)
	if err != nil {
		return err
	}
	// Menerima file gambar dari form dengan nama "image"
	image, err := ctx.FormFile("image")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Tidak dapat memproses file gambar"})
	}

	// Buka file yang diunggah
	src, err := image.Open()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal membuka file gambar"})
	}
	defer src.Close()

	// Lokasi penyimpanan file gambar lokal
	uploadDir := "uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal membuat direktori penyimpanan"})
	}

	// Generate nama file unik
	dstPath := filepath.Join(uploadDir, image.Filename)

	// Membuka file tujuan untuk penyimpanan
	dst, err := os.Create(dstPath)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal membuat file gambar"})
	}
	defer dst.Close()

	// Salin isi file dari file asal ke file tujuan
	if _, err = io.Copy(dst, src); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal menyalin file gambar"})
	}

	// Membuat URL ke gambar yang diunggah
	imageURL := "http://localhost:8080/uploads/" + image.Filename

	Claims := helpers.ClaimToken(ctx)
	Id := Claims.ID

	// validate .......
	task, err := c.service.UpdateTask(req, imageURL, Id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"task":    task,
		"Message": "successfully updated task",
	})
}

func (c *Controller) SearchTasksFormController(ctx echo.Context) (err error) {
	Claims := helpers.ClaimToken(ctx)
	id := Claims.ID
	keywoard := ctx.QueryParam("search")
	dateStr := ctx.QueryParam("date")
	var parsedDate time.Time
	if dateStr != "" {
		parsedDate, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			return err
		}
	}

	limit, err := strconv.Atoi(ctx.QueryParam("limit"))
	if err != nil {
		limit = 10
	}

	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil {
		page = 1
	}

	offset := (page - 1) * limit

	users, err := c.service.SearchTasks(id, keywoard, parsedDate, limit, offset)
	if err != nil {
		return err
	}

	count, err := c.service.CountTasks(id, keywoard, parsedDate)
	if err != nil {
		return err
	}

	totalPages := count / limit
	if count%limit != 0 {
		totalPages++
	}

	if len(users) == 0 {
		users = []model.TaskRes{}
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"Message":     "Success Search Tasks for User",
		"data":        users,
		"page":        page,
		"limit_page":  limit,
		"total_data":  count,
		"total_pages": totalPages,
	})
}

func (c *Controller) CountTask(ctx echo.Context) error {
	claims := helpers.ClaimToken(ctx)
	id := claims.ID

	data, err := c.service.CountTask(id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"Message": "successfully displays status data",
		"data":    data,
	})
}

// AUTH
func (c *Controller) Login(ctx echo.Context) error {
	var req model.UserLogin
	err := ctx.Bind(&req)
	if err != nil {
		return err
	}

	data, err := c.service.Login(req.Email, req.Password)
	if err != nil {
		return err
	}

	var (
		jwtToken  *jwt.Token
		secretKey = []byte("secret")
	)

	jwtClaims := &model.Claims{
		ID:    data.ID,
		Email: data.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	jwtToken = jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	token, err := jwtToken.SignedString(secretKey)
	if err != nil {
		return err
	}

	err = c.service.SaveToken(token, int(data.ID))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Login Successful",
		"token":   token,
		"data":    data,
	})
}

func (c *Controller) Logout(ctx echo.Context) error {
	var reqToken string
	headerDataToken := ctx.Request().Header.Get("Authorization")

	splitToken := strings.Split(headerDataToken, "Bearer ")
	if len(splitToken) > 1 {
		reqToken = splitToken[1]
	} else {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	err := c.service.Logout(reqToken)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"messgae": "logout successfully",
	})
}
func (c *Controller) RegisterController(ctx echo.Context) error {
	var req model.UserRegis
	err := ctx.Bind(&req)
	if err != nil {
		return err
	}

	data, err := c.service.Regis(req.Email, req.Password)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Register Successful",
		"data":    data,
	})
}

// KATEGORI
func (c *Controller) GetKategoriController(ctx echo.Context) error {

	kategori, err := c.service.GetAllKategori()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Succesfully Get All Category Data",
		"data":    kategori,
	})
}
func (c *Controller) AddKategoriController(ctx echo.Context) error {
	var katreq model.KategoriReq

	err := ctx.Bind(&katreq)
	if err != nil {
		return err
	}

	kategori, err := c.service.CreateKategori(katreq)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Successfully Added New category Data",
		"data":    kategori,
	})
}
func (c *Controller) EditKategoriController(ctx echo.Context) error {
	var katreq model.KategoriReq

	err := ctx.Bind(&katreq)
	if err != nil {
		return err
	}

	Id := ctx.Param("id")
	IdAsli, err := strconv.Atoi(Id)
	if err != nil {
		return err
	}

	kategori, err := c.service.EditKategori(IdAsli, katreq)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Successfully edited Data ",
		"data":    kategori,
	})
}
func (c *Controller) DeleteKategoriController(ctx echo.Context) error {
	Id := ctx.Param("id")
	IdAsli, err := strconv.Atoi(Id)
	if err != nil {
		return err
	}

	err = c.service.DeleteKategori(IdAsli)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Category deleted successfully",
	})
}


func ForgotPasswordHandler(ctx echo.Context) error {
	
	email := ctx.FormValue("email")

	
	if email == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Email tidak valid"})
	}

	if err := service.sendResetPasswordEmail(email); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal mengirim email"})
	}

	
	return ctx.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("Tautan reset password dikirim ke %s", email)})
}