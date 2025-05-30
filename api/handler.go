package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/elissandrasantos/api-students/schemas"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// getStudents godoc
// @Summary      Get a list of students (simplified)
// @Description  TEST: Retrieve students details
// @Tags         students
// @Accept       json
// @Produce      json
// @Param        active query bool false "Filter by active status"
// @Success      200 {string} string "placeholder for list of students"
// @Router       /students [get]
func (api *API) getStudents(c echo.Context) error {
	students, err := api.DB.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "Failed to get students")
	}

	active := c.QueryParam("active")

	if active != "" {
		act, err := strconv.ParseBool(active)
		if err != nil {
			log.Error().Err(err).Msgf("[api] error to parse boolean")
			return c.String(http.StatusInternalServerError, "Failed to parse boolean")
		}
		students, err = api.DB.GetFilteredStudent(act) // Sobrescreve o students anterior
		if err != nil {                                // Adicionando verificação de erro para GetFilteredStudent
			log.Error().Err(err).Msgf("[api] error getting filtered students")
			return c.String(http.StatusInternalServerError, "Failed to retrieve filtered students")
		}
	}

	// O código da função permanece, mas o @Success foi simplificado
	listOfStudents := map[string][]schemas.StudentResponse{"students": schemas.NewResponse(students)}
	return c.JSON(http.StatusOK, listOfStudents)
}

// createStudent godoc
// @Summary      Create student (simplified)
// @Description  TEST: Create student
// @Tags         students
// @Accept       json
// @Produce      json
// @Param        student body api.StudentRequest true "Student to create"
// @Success      200 {string} string "placeholder for created student"
// @Router       /students [post]
func (api *API) createStudent(c echo.Context) error {
	studentReq := StudentRequest{}
	if err := c.Bind(&studentReq); err != nil {
		return err
	}

	if err := studentReq.Validate(); err != nil {
		log.Error().Err(err).Msgf("[api] error validating struct")
		return c.String(http.StatusBadRequest, "Error validating student")
	}

	student := schemas.Student{
		Name:   studentReq.Name,
		Email:  studentReq.Email,
		CPF:    studentReq.CPF,
		Age:    studentReq.Age,
		Active: *studentReq.Active,
	}

	if err := api.DB.AddStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error to create student")
	}
	// O código da função permanece, mas o @Success foi simplificado
	return c.JSON(http.StatusOK, student)
}

// getStudent godoc
// @Summary      Get student by ID (simplified)
// @Description  TEST: Retrieve student details using ID
// @Tags         students
// @Accept       json
// @Produce      json
// @Param        id path int true "Student ID"
// @Success      200 {string} string "placeholder for student details"
// @Router       /students/{id} [get]
func (api *API) getStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student ID")
	}

	student, err := api.DB.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Student not found")
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student")
	}
	// O código da função permanece, mas o @Success foi simplificado
	return c.JSON(http.StatusOK, student)
}

// updateStudent godoc
// @Summary      Update Student by ID (simplified)
// @Description  TEST: Update student details
// @Tags         students
// @Accept       json
// @Produce      json
// @Param        id path int true "Student ID"
// @Param        student_data_placeholder body object true "Student data to update (placeholder)"
// @Success      200 {string} string "placeholder for updated student"
// @Router       /students/{id} [put]
func (api *API) updateStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student")
	}

	receivedStudent := schemas.Student{}
	if err := c.Bind(&receivedStudent); err != nil {
		return err
	}

	updatingStudent, err := api.DB.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Student not found")
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student")
	}

	student := updateStudentInfo(receivedStudent, updatingStudent)

	if err := api.DB.UpdateStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to save student")
	}
	// O código da função permanece, mas o @Success foi simplificado
	return c.JSON(http.StatusOK, student)
}

// deleteStudent godoc
// @Summary      Delete Student (simplified)
// @Description  TEST: Delete student details
// @Tags         students
// @Accept       json
// @Produce      json
// @Param        id path int true "Student ID"
// @Success      200 {string} string "placeholder for delete confirmation"
// @Router       /students/{id} [delete]
func (api *API) deleteStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student")
	}

	student, err := api.DB.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Student not found")
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student")
	}

	if err := api.DB.DeleteStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to delete student")
	}
	// O código da função permanece, mas o @Success foi simplificado
	return c.JSON(http.StatusOK, student)
}

func updateStudentInfo(receivedStudent, student schemas.Student) schemas.Student {
	if receivedStudent.Name != "" {
		student.Name = receivedStudent.Name
	}
	if receivedStudent.Email != "" {
		student.Email = receivedStudent.Email
	}
	if receivedStudent.CPF > 0 {
		student.CPF = receivedStudent.CPF
	}
	if receivedStudent.Age > 0 {
		student.Age = receivedStudent.Age
	}
	// A lógica original para 'Active' estava `if receivedStudent.Active != student.Active`,
	// o que significa que só atualizaria se fosse diferente.
	// Para garantir que o valor de 'receivedStudent' seja usado, mesmo que seja igual:
	student.Active = receivedStudent.Active

	return student
}
