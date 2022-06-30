package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"../entities"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

const conn = "root:pari@tcp(34.196.253.50:3306)/YogaStudentdb"

func GetStudents(w http.ResponseWriter, r *http.Request) {
	db, err := sqlx.Open("mysql", conn)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Queryx(`call GetAllStudents()`)
	if err != nil {
		log.Fatal(err)
	} else {
		var students []entities.Students
		for rows.Next() {
			var studentId string
			var name string
			var address string
			var mobile string
			var email string
			var emergencyName string
			var emergencyContact string
			var referredBy string

			var gender string
			var dateOfBirth string
			var occupation string
			var education string
			var relationship string
			var dateOfJoining string
			var file string

			var height string
			var weight string
			var currentLevelOfActivity string
			var currentRoutine string
			var mostStress string
			var pranayamaPractice string
			var anySurgery string
			var anyMedical string
			var smoked string

			var reasonForYoga string
			var musculoskeletal string
			var respiratory string
			var cardiovascular string
			var circulatory string
			var neurological string
			var gastrointestinal string
			var endocrinological string
			var gynecologicalOrUrological string
			var otherMedicalConditions string
			var otherInformations string

			// var UserId string
			// var CreatedBy string
			// var CreatedDate string
			// var ModifiedBy string
			// var ModifiedDate string
			// var ActiveStatus bool
			err2 := rows.Scan(&studentId, &name, &address, &mobile, &email, &emergencyName, &emergencyContact, &referredBy,
				&gender, &dateOfBirth, &occupation, &education, &relationship, &dateOfJoining, &file,
				&height, &weight, &currentLevelOfActivity, &currentRoutine, &mostStress, &pranayamaPractice, &anySurgery, &anyMedical, &smoked,
				&reasonForYoga, &musculoskeletal, &respiratory, &cardiovascular, &circulatory, &neurological, &gastrointestinal, &endocrinological, &gynecologicalOrUrological, &otherMedicalConditions, &otherInformations)
			// &UserId, &CreatedBy, &CreatedDate, &ModifiedBy, &ModifiedDate, &ActiveStatus)

			if err2 != nil {
				log.Fatal(err)
			} else {
				student := entities.Students{studentId, name, address, mobile, email, emergencyName, emergencyContact, referredBy,
					gender, dateOfBirth, occupation, education, relationship, dateOfJoining, file,
					height, weight, currentLevelOfActivity, currentRoutine, mostStress, pranayamaPractice, anySurgery, anyMedical, smoked,
					reasonForYoga, musculoskeletal, respiratory, cardiovascular, circulatory, neurological, gastrointestinal, endocrinological, gynecologicalOrUrological, otherMedicalConditions, otherInformations}
				//

				students = append(students, student)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(students)
	}

}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db, err := sqlx.Open("mysql", conn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	students := entities.Students{}
	students.StudentId = params["studentId"]
	fmt.Println("this is the id:", students.StudentId)

	stmt := fmt.Sprintf(` call GetOneStudents('%s')`, students.StudentId)

	rows, err := db.Queryx(stmt)
	if err != nil {
		log.Fatal(err)
	} else {
		var students []entities.Students
		for rows.Next() {
			var studentId string
			var name string
			var address string
			var mobile string
			var email string
			var emergencyName string
			var emergencyContact string
			var referredBy string

			var gender string
			var dateOfBirth string
			var occupation string
			var education string
			var relationship string
			var dateOfJoining string
			var file string
			// f,err:= os.Open("file")
			// if err != nil {
			// 	w.WriteHeader(500)
			// 	return
			// }

			var height string
			var weight string
			var currentLevelOfActivity string
			var currentRoutine string
			var mostStress string
			var pranayamaPractice string
			var anySurgery string
			var anyMedical string
			var smoked string

			var reasonForYoga string
			var musculoskeletal string
			var respiratory string
			var cardiovascular string
			var circulatory string
			var neurological string
			var gastrointestinal string
			var endocrinological string
			var gynecologicalOrUrological string
			var otherMedicalConditions string
			var otherInformations string

			// var UserId string
			// var CreatedBy string
			// var CreatedDate string
			// var ModifiedBy string
			// var ModifiedDate string
			// var ActiveStatus bool

			err2 := rows.Scan(&studentId, &name, &address, &mobile, &email, &emergencyName, &emergencyContact, &referredBy,
				&gender, &dateOfBirth, &occupation, &education, &relationship, &dateOfJoining, &file,
				&height, &weight, &currentLevelOfActivity, &currentRoutine, &mostStress, &pranayamaPractice, &anySurgery, &anyMedical, &smoked,
				&reasonForYoga, &musculoskeletal, &respiratory, &cardiovascular, &circulatory, &neurological, &gastrointestinal, &endocrinological, &gynecologicalOrUrological, &otherMedicalConditions, &otherInformations)
			// &UserId, &CreatedBy, &CreatedDate, &ModifiedBy, &ModifiedDate, &ActiveStatus)

			if err2 != nil {
				log.Fatal(err2)
			} else {
				student := entities.Students{studentId, name, address, mobile, email, emergencyName, emergencyContact, referredBy,
					gender, dateOfBirth, occupation, education, relationship, dateOfJoining, file,
					height, weight, currentLevelOfActivity, currentRoutine, mostStress, pranayamaPractice, anySurgery, anyMedical, smoked,
					reasonForYoga, musculoskeletal, respiratory, cardiovascular, circulatory, neurological, gastrointestinal, endocrinological, gynecologicalOrUrological, otherMedicalConditions, otherInformations,
					// UserId, CreatedBy, CreatedDate, ModifiedBy, ModifiedDate, ActiveStatus}
				}
				students = append(students, student)
			}
		}
		w.Header().Set("Content-Type", "multipart/form-data")
		json.NewEncoder(w).Encode(students)
		fmt.Println(students)
	}

}

// func CreateStudents(w http.ResponseWriter, r *http.Request) {
// 	db, err := sqlx.Open("mysql", conn)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// defer db.Close()
// 	var student *Students = new(Students)
// 	_ = json.NewDecoder(r.Body).Decode(&student)

// 	fmt.Println(student)
// 	stmt := fmt.Sprintf(` call InsertStudent('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s')`, student.Name, student.Address,
// 		student.Mobile, student.Email, student.EmergencyName, student.EmergencyContact, student.ReferredBy, student.Gender, student.DateOfBirth, student.Occupation, student.Education, student.Relationship, student.DateOfJoining, student.File,
// 		student.Height, student.Weight, student.CurrentLevelOfActivity, student.CurrentRoutine, student.MostStress, student.PranayamaPractice, student.AnySurgery, student.AnyMedical, student.Smoked,
// 		student.ReasonForYoga, student.Musculoskeletal, student.Respiratory, student.Cardiovascular, student.Circulatory, student.Neurological, student.Gastrointestinal, student.Endocrinological, student.GynecologicalOrUrological, student.OtherMedicalConditions, student.OtherInformations,
// 		// student.UserId, student.CreatedBy, student.CreatedDate, student.ModifiedBy, student.ModifiedDate, student.ActiveStatus)
// 	)

// 	res, err := db.Queryx(stmt)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Print("student was created")
// 	w.WriteHeader(200)
// 	defer res.Close()

// }

func CreateStudents(w http.ResponseWriter, r *http.Request) {
	db, err := sqlx.Open("mysql", conn)
	if err != nil {
		log.Fatal(err)
	}

	var student *entities.Students = new(entities.Students)

	student.Name = r.FormValue("name")
	student.Address = r.FormValue("address")
	student.Mobile = r.FormValue("mobile")
	student.Email = r.FormValue("email")
	student.EmergencyName = r.FormValue("emergencyName")
	student.EmergencyContact = r.FormValue("emergencyContact")
	student.ReferredBy = r.FormValue("referredBy")
	student.Gender = r.FormValue("gender")
	student.DateOfBirth = r.FormValue("dateOfBirth")
	student.Occupation = r.FormValue("occupation")
	student.Education = r.FormValue("education")
	student.Relationship = r.FormValue("relationship")
	student.DateOfJoining = r.FormValue("dateOfJoining")
	File, handler, err := r.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}
	defer File.Close()

	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, _ = io.Copy(f, File)

	fmt.Println(f)

	student.Height = r.FormValue("height")
	student.Weight = r.FormValue("weight")
	student.CurrentLevelOfActivity = r.FormValue("currentLevelOfActivity")
	student.CurrentRoutine = r.FormValue("currentRoutine")
	student.MostStress = r.FormValue("mostStress")
	student.PranayamaPractice = r.FormValue("pranayamaPractice")
	student.AnySurgery = r.FormValue("anySurgery")
	student.AnyMedical = r.FormValue("anyMedical")
	student.Smoked = r.FormValue("smoked")
	student.ReasonForYoga = r.FormValue("reasonForYoga")
	student.Musculoskeletal = r.FormValue("musculoskeletal")
	student.Respiratory = r.FormValue("respiratory")
	student.Cardiovascular = r.FormValue("cardiovascular")
	student.Circulatory = r.FormValue("circulatory")
	student.Neurological = r.FormValue("neurological")
	student.Gastrointestinal = r.FormValue("gastrointestinal")
	student.Endocrinological = r.FormValue("endocrinological")
	student.GynecologicalOrUrological = r.FormValue("gynecologicalOrUrological")
	student.OtherMedicalConditions = r.FormValue("otherMedicalConditions")
	student.OtherInformations = r.FormValue("otherInformations")

	// fmt.Println(student)
	stmt := fmt.Sprintf(` call InsertStudent('%v','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%v','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s')`, 0, student.Name, student.Address,
		student.Mobile, student.Email, student.EmergencyName, student.EmergencyContact, student.ReferredBy, student.Gender, student.DateOfBirth, student.Occupation, student.Education, student.Relationship, student.DateOfJoining, f,
		student.Height, student.Weight, student.CurrentLevelOfActivity, student.CurrentRoutine, student.MostStress, student.PranayamaPractice, student.AnySurgery, student.AnyMedical, student.Smoked,
		student.ReasonForYoga, student.Musculoskeletal, student.Respiratory, student.Cardiovascular, student.Circulatory, student.Neurological, student.Gastrointestinal, student.Endocrinological, student.GynecologicalOrUrological, student.OtherMedicalConditions, student.OtherInformations,
		// student.UserId, student.CreatedBy, student.CreatedDate, student.ModifiedBy, student.ModifiedDate, student.ActiveStatus)
	)

	res, err := db.Queryx(stmt)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("student was created")
	w.WriteHeader(200)
	defer res.Close()

}

// func CreateStudents(w http.ResponseWriter, r *http.Request) {
// 	db, err := sqlx.Open("mysql", conn)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
// 	var student *entities.Students = new(entities.Students)
// 	_ = json.NewDecoder(r.Body).Decode(&student)

// 	// fmt.Println(student)
// 	stmt := fmt.Sprintf(` call InsertStudent('%v','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%v','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s')`, 0, student.Name, student.Address,
// 		student.Mobile, student.Email, student.EmergencyName, student.EmergencyContact, student.ReferredBy, student.Gender, student.DateOfBirth, student.Occupation, student.Education, student.Relationship, student.DateOfJoining, student.File,
// 		student.Height, student.Weight, student.CurrentLevelOfActivity, student.CurrentRoutine, student.MostStress, student.PranayamaPractice, student.AnySurgery, student.AnyMedical, student.Smoked,
// 		student.ReasonForYoga, student.Musculoskeletal, student.Respiratory, student.Cardiovascular, student.Circulatory, student.Neurological, student.Gastrointestinal, student.Endocrinological, student.GynecologicalOrUrological, student.OtherMedicalConditions, student.OtherInformations,
// 		// student.UserId, student.CreatedBy, student.CreatedDate, student.ModifiedBy, student.ModifiedDate, student.ActiveStatus)
// 	)

// 	res, err := db.Queryx(stmt)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Print("student was created")
// 	w.WriteHeader(200)
// 	defer res.Close()

// }

func DeleteStudents(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db, err := sqlx.Open("mysql", conn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	students := entities.Students{}
	students.StudentId = params["studentId"]

	fmt.Println("this is the id:", students.StudentId)

	stmt := fmt.Sprintf(`call DeleteOneStudent('%s')`, students.StudentId)

	res, err := db.Queryx(stmt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Employee was deleted")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "Employee was deleted")
	w.WriteHeader(200)
	defer res.Close()

}

func UpdateStudents(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db, err := sqlx.Open("mysql", conn)
	if err != nil {
		log.Fatal(err)
	}
	students := entities.Students{}
	students.StudentId = params["studentId"]

	var student *entities.Students = new(entities.Students)

	student.StudentId = r.FormValue("studentId")
	student.Name = r.FormValue("name")
	student.Address = r.FormValue("address")
	student.Mobile = r.FormValue("mobile")
	student.Email = r.FormValue("email")
	student.EmergencyName = r.FormValue("emergencyName")
	student.EmergencyContact = r.FormValue("emergencyContact")
	student.ReferredBy = r.FormValue("referredBy")
	student.Gender = r.FormValue("gender")
	student.DateOfBirth = r.FormValue("dateOfBirth")
	student.Occupation = r.FormValue("occupation")
	student.Education = r.FormValue("education")
	student.Relationship = r.FormValue("relationship")
	student.DateOfJoining = r.FormValue("dateOfJoining")
	File, handler, err := r.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}
	defer File.Close()

	fmt.Println("\nfile:", File, "\nheader:", handler, "\nerr", err)
	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, _ = io.Copy(f, File)

	fmt.Println(f)

	// res1, err := ioutil.ReadAll(File)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	student.Height = r.FormValue("height")
	student.Weight = r.FormValue("weight")
	student.CurrentLevelOfActivity = r.FormValue("currentLevelOfActivity")
	student.CurrentRoutine = r.FormValue("currentRoutine")
	student.MostStress = r.FormValue("mostStress")
	student.PranayamaPractice = r.FormValue("pranayamaPractice")
	student.AnySurgery = r.FormValue("anySurgery")
	student.AnyMedical = r.FormValue("anyMedical")
	student.Smoked = r.FormValue("smoked")
	student.ReasonForYoga = r.FormValue("reasonForYoga")
	student.Musculoskeletal = r.FormValue("musculoskeletal")
	student.Respiratory = r.FormValue("respiratory")
	student.Cardiovascular = r.FormValue("cardiovascular")
	student.Circulatory = r.FormValue("circulatory")
	student.Neurological = r.FormValue("neurological")
	student.Gastrointestinal = r.FormValue("gastrointestinal")
	student.Endocrinological = r.FormValue("endocrinological")
	student.GynecologicalOrUrological = r.FormValue("gynecologicalOrUrological")
	student.OtherMedicalConditions = r.FormValue("otherMedicalConditions")
	student.OtherInformations = r.FormValue("otherInformations")

	// fmt.Println(student)
	stmt := fmt.Sprintf(` call UpdateStudent('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%v','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s')`, students.StudentId, student.Name, student.Address,
		student.Mobile, student.Email, student.EmergencyName, student.EmergencyContact, student.ReferredBy, student.Gender, student.DateOfBirth, student.Occupation, student.Education, student.Relationship, student.DateOfJoining, f,
		student.Height, student.Weight, student.CurrentLevelOfActivity, student.CurrentRoutine, student.MostStress, student.PranayamaPractice, student.AnySurgery, student.AnyMedical, student.Smoked,
		student.ReasonForYoga, student.Musculoskeletal, student.Respiratory, student.Cardiovascular, student.Circulatory, student.Neurological, student.Gastrointestinal, student.Endocrinological, student.GynecologicalOrUrological, student.OtherMedicalConditions, student.OtherInformations,
		// student.UserId, student.CreatedBy, student.CreatedDate, student.ModifiedBy, student.ModifiedDate, student.ActiveStatus)
	)

	res, err := db.Queryx(stmt)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("student was Updated")
	w.WriteHeader(200)
	defer res.Close()

}

// func UpdateStudents(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	db, err := sqlx.Open("mysql", conn)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
// 	var student *entities.Students = new(entities.Students)
// 	_ = json.NewDecoder(r.Body).Decode(&student)
// 	students := entities.Students{}
// 	students.StudentId = params["studentId"]

// 	// fmt.Println(student)
// 	stmt := fmt.Sprintf(` call UpdateStudent('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%v','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s')`, students.StudentId, student.Name, student.Address,
// 		student.Mobile, student.Email, student.EmergencyName, student.EmergencyContact, student.ReferredBy, student.Gender, student.DateOfBirth, student.Occupation, student.Education, student.Relationship, student.DateOfJoining, student.File,
// 		student.Height, student.Weight, student.CurrentLevelOfActivity, student.CurrentRoutine, student.MostStress, student.PranayamaPractice, student.AnySurgery, student.AnyMedical, student.Smoked,
// 		student.ReasonForYoga, student.Musculoskeletal, student.Respiratory, student.Cardiovascular, student.Circulatory, student.Neurological, student.Gastrointestinal, student.Endocrinological, student.GynecologicalOrUrological, student.OtherMedicalConditions, student.OtherInformations,
// 		// student.UserId, student.CreatedBy, student.CreatedDate, student.ModifiedBy, student.ModifiedDate, student.ActiveStatus)
// 	)

// 	res, err := db.Queryx(stmt)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Print("student was created")
// 	w.WriteHeader(200)
// 	defer res.Close()

// }

// const SecretKey = "secret"

// func Login(c *fiber.Ctx) error {

// 	var data map[string]string

// 	if err := c.BodyParser(&data); err != nil {
// 		log.Fatal(err)
// 	}

// 	var user entities.Users

// 	DB.Where("userName= ?", data["userName"]).First(&user)

// 	if user.UserName == "" {
// 		c.Status(fiber.StatusNotFound)
// 		return c.JSON(fiber.Map{
// 			"message": "user not found",
// 		})
// 	}

// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["passWord"])); err != nil {
// 		c.Status(fiber.StatusBadRequest)
// 		return c.JSON(fiber.Map{
// 			"message": "incorrect password",
// 		})
// 	}

// 	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
// 		Issuer:    user.UserId,
// 		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
// 	})

// 	token, err := claims.SignedString([]byte(SecretKey))

// 	if err != nil {
// 		c.Status(fiber.StatusInternalServerError)
// 		return c.JSON(fiber.Map{
// 			"message": "could not login",
// 		})
// 	}

// 	cookie := fiber.Cookie{
// 		Name:     "jwt",
// 		Value:    token,
// 		Expires:  time.Now().Add(time.Hour * 24),
// 		HTTPOnly: true,
// 	}

// 	c.Cookie(&cookie)

// 	return c.JSON(fiber.Map{
// 		"message": "success",
// 	})

// }

type PartiFile struct {
	file []byte `json:file`
}

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db, err := sqlx.Open("mysql", conn)
	if err != nil {
		log.Fatal(err)
	}

	students := entities.Students{}
	students.StudentId = params["studentId"]
	// fmt.Println(students.StudentId)
	stmt := fmt.Sprintf(` SELECT File FROM studentdb.tblstudents where StudentId=%s;`, students.StudentId)
	rows, err := db.Queryx(stmt)
	if err != nil {
		log.Fatal(err)
	} else {
		var f PartiFile
		for rows.Next() {
			rows.Scan(&f)

		}
		fmt.Println(f)
		jsn, err := json.Marshal(f)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v", jsn)
		fmt.Println(f)
		w.Write(jsn)

		// File, err := os.Open(jsn)
		// if err != nil {
		// 	fmt.Println(err)
		// 	w.WriteHeader(500)
		// 	return
		// }
		// defer File.Close()

		// if _, err := io.Copy(w, ); err != nil {
		// 	fmt.Println(err)
		// 	w.WriteHeader(500)
		// }

		// ctx := context.Background()
		// b, err := blob.OpenBucket(ctx)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// defer b.Close()
		// data, err := b.ReadAll(ctx, "my-key")
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Println(string(data))

	}

	// f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_RDONLY, 0666)
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// _, _ = io.Copy(f, rows)

}

// } else {
// 	var partiFile []PartiFile
// 	for rows.Next() {
// 		var file string

// 		err2 := rows.Scan(&file)
// 		if err2 != nil {
// 			log.Fatal(err2)
// 		} else {
// 			File := PartiFile{file}
// 			partiFile = append(partiFile, File)
// 			fmt.Println(partiFile)
// 		}

// 	}
// 	json.NewEncoder(w).Encode(partiFile)
// }
