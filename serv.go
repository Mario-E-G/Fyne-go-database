package main

import (
	// "context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	_ "github.com/go-sql-driver/mysql"
)

var ip string = "127.0.0.1"
var port int = 3306
var userName string = "root"
var password string = "rootroot"
var DataBaseName string = "hospital"

func main() {
	Options := [...]string{"Get all databases", "Show tables", "Insert", "Update", "Select", "Delete Record", "Search By Id", "Quit"}

	fmt.Println("App is Running....")
	a := app.New()
	w := a.NewWindow("Database Access")
	w.Resize(fyne.NewSize(1200, 600))

	ListView := widget.NewList(func() int {
		return len(Options)
	}, func() fyne.CanvasObject {
		return widget.NewLabel("Templates")
	}, func(lii widget.ListItemID, co fyne.CanvasObject) {
		co.(*widget.Label).SetText(Options[lii])
	})

	contentText := widget.NewLabel("Please Choose An Option....")
	msg := widget.NewLabel("")
	searchMsg := widget.NewLabel("Please Choose Table To Search In....")

	ListView.OnSelected = func(id widget.ListItemID) {

		if id == 0 { // show all databases
			var databaseList []string
			var table string
			db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+":"+strconv.Itoa(port)+")/"+DataBaseName)
			res, err := db.Query("SHOW DATABASES")
			// contentText.SetText("DATABASES ARE : ")

			var i int = 1
			for res.Next() {
				res.Scan(&table)
				i++
				databaseList = append(databaseList, table)
			}
			if err != nil {
				fmt.Println("Erorr : ", err)
			}
			Lv := widget.NewList(func() int {
				return len(databaseList)
			}, func() fyne.CanvasObject {
				return widget.NewLabel("Templates")
			}, func(lii widget.ListItemID, co fyne.CanvasObject) {
				co.(*widget.Label).SetText(databaseList[lii])
			})
			fmt.Println("Connection Succeeded ", res)
			w.SetContent(
				container.NewGridWithColumns(
					2,
					ListView,
					Lv,
				),
			)
			Lv.Refresh()
		} else if id == 1 { //show Tables
			var databaseList []string

			var table string
			db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+":"+strconv.Itoa(port)+")/"+DataBaseName)
			res, err := db.Query("SHOW TABLES")
			for res.Next() {
				res.Scan(&table)
				databaseList = append(databaseList, table)
			}
			if err != nil {
				fmt.Println("Erorr : ", err)
			}
			tablelist := widget.NewList(func() int { return len(databaseList) },
				func() fyne.CanvasObject { return widget.NewLabel("Templates") },
				func(lii widget.ListItemID, co fyne.CanvasObject) {
					co.(*widget.Label).SetText(databaseList[lii])
				})

			w.SetContent(container.NewGridWithColumns(
				2,
				ListView,
				tablelist,
			))
		} else if id == 2 { // INSERT
			var databaseList []string

			var table string
			db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+":"+strconv.Itoa(port)+")/"+DataBaseName)
			res, err := db.Query("SHOW TABLES")
			// fmt.Println("tables:")
			var i int = 1
			for res.Next() {
				res.Scan(&table)
				i++
				databaseList = append(databaseList, table)
			}
			if err != nil {
				fmt.Println("Erorr : ", err)
			}
			tablelist := widget.NewList(func() int { return len(databaseList) },
				func() fyne.CanvasObject { return widget.NewLabel("Templates") },
				func(lii widget.ListItemID, co fyne.CanvasObject) {
					co.(*widget.Label).SetText(databaseList[lii])
				})
			var content *fyne.Container = container.NewWithoutLayout()
			tablelist.OnSelected = func(id widget.ListItemID) {
				if id == 0 { //bill
					msg.SetText("")
					BillNO := widget.NewEntry()
					BillNO.SetPlaceHolder("Enter Bill Number ")
					BillNO.Resize(fyne.NewSize(200, 40))
					BillNO.Move(fyne.NewPos(300, 140))

					PatientID := widget.NewEntry()
					PatientID.SetPlaceHolder("Enter Patient Id ")
					PatientID.Resize(fyne.NewSize(200, 40))
					PatientID.Move(fyne.NewPos(300, 180))

					RoomCharge := widget.NewEntry()
					RoomCharge.SetPlaceHolder("Enter Room Charge ")
					RoomCharge.Resize(fyne.NewSize(200, 40))
					RoomCharge.Move(fyne.NewPos(300, 180))

					DoctorCharge := widget.NewEntry()
					DoctorCharge.SetPlaceHolder("Enter Doctor Charge ")
					DoctorCharge.Resize(fyne.NewSize(200, 40))
					DoctorCharge.Move(fyne.NewPos(300, 180))

					NoOfDays := widget.NewEntry()
					NoOfDays.SetPlaceHolder("Enter Number Of Days ")
					NoOfDays.Resize(fyne.NewSize(200, 40))
					NoOfDays.Move(fyne.NewPos(300, 180))

					LabChargeBill := widget.NewEntry()
					LabChargeBill.SetPlaceHolder("Enter Lab Charge Bill ")
					LabChargeBill.Resize(fyne.NewSize(200, 40))
					LabChargeBill.Move(fyne.NewPos(300, 180))

					if err != nil {
						fmt.Println(err)
					}
					btn := widget.NewButton("Insert", func() {
						if BillNO.Text == "" || PatientID.Text == "" || RoomCharge.Text == "" || DoctorCharge.Text == "" || NoOfDays.Text == "" || LabChargeBill.Text == "" {
							msg.SetText("All inputs must be filled")
						} else {
							db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+":"+strconv.Itoa(port)+")/"+DataBaseName)
							dd, err := db.Query("insert into bill (BillNO, PatientID,RoomCharge,DoctorCharge,NoOfDays,LabChargeBill) Values " + "(" + BillNO.Text + ", '" + PatientID.Text + "', '" + RoomCharge.Text + "', '" + DoctorCharge.Text + "','" + NoOfDays.Text + "','" + LabChargeBill.Text + "')")
							if err != nil {
								fmt.Println(err, dd)
							}
							msg.SetText("Row inserted Successfully!!! ")
						}
					})
					btn.Resize(fyne.NewSize(200, 40))
					btn.Move(fyne.NewPos(300, 225))

					content = container.NewGridWithRows(8, BillNO, PatientID, RoomCharge, DoctorCharge, NoOfDays, LabChargeBill, btn, msg)
					w.SetContent(container.NewGridWithColumns(
						3,
						ListView,
						tablelist,
						content,
					),
					)

				} else if id == 1 { //doctor
					msg.SetText("")

					DoctorID := widget.NewEntry()
					DoctorID.SetPlaceHolder("Enter Doctor Id ")
					DoctorID.Resize(fyne.NewSize(200, 40))
					DoctorID.Move(fyne.NewPos(300, 140))

					DoctorName := widget.NewEntry()
					DoctorName.SetPlaceHolder("Enter Doctor Name ")
					DoctorName.Resize(fyne.NewSize(200, 40))
					DoctorName.Move(fyne.NewPos(300, 180))

					Age := widget.NewEntry()
					Age.SetPlaceHolder("Enter Age ")
					Age.Resize(fyne.NewSize(200, 40))
					Age.Move(fyne.NewPos(300, 180))

					Gender := widget.NewEntry()
					Gender.SetPlaceHolder("Enter Gender ")
					Gender.Resize(fyne.NewSize(200, 40))
					Gender.Move(fyne.NewPos(300, 180))

					Address := widget.NewEntry()
					Address.SetPlaceHolder("Enter Address ")
					Address.Resize(fyne.NewSize(200, 40))
					Address.Move(fyne.NewPos(300, 180))

					Speciality := widget.NewEntry()
					Speciality.SetPlaceHolder("Enter Speciality ")
					Speciality.Resize(fyne.NewSize(200, 40))
					Speciality.Move(fyne.NewPos(300, 180))

					if err != nil {
						fmt.Println(err)
					}
					btn := widget.NewButton("Insert", func() {
						if DoctorID.Text == "" || DoctorName.Text == "" || Age.Text == "" || Gender.Text == "" || Address.Text == "" || Speciality.Text == "" {
							msg.SetText("All inputs must be filled")
						} else {
							db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+":"+strconv.Itoa(port)+")/"+DataBaseName)
							dd, err := db.Query("insert into doctor (DoctorID, DoctorName,Age,Gender,Address,Speciality) Values " + "(" + DoctorID.Text + ", '" + DoctorName.Text + "', '" + Age.Text + "', '" + Gender.Text + "','" + Address.Text + "','" + Speciality.Text + "')")
							if err != nil {
								fmt.Println(err, dd)
							}
							msg.SetText(" Row inserted Successfully!!! ")
						}
					})
					btn.Resize(fyne.NewSize(200, 40))
					btn.Move(fyne.NewPos(300, 225))

					content = container.NewGridWithRows(8, DoctorID, DoctorName, Age, Gender, Address, Speciality, btn, msg)
					w.SetContent(container.NewGridWithColumns(
						3,
						ListView,
						tablelist,
						content,
					))
				} else if id == 2 { //inpatient
					msg.SetText("")

					PatientID := widget.NewEntry()
					PatientID.SetPlaceHolder("Enter Patient ID ")
					PatientID.Resize(fyne.NewSize(200, 40))
					PatientID.Move(fyne.NewPos(300, 140))

					RoomNo := widget.NewEntry()
					RoomNo.SetPlaceHolder("Enter Room Number ")
					RoomNo.Resize(fyne.NewSize(200, 40))
					RoomNo.Move(fyne.NewPos(300, 180))

					LabNo := widget.NewEntry()
					LabNo.SetPlaceHolder("Enter Lab Number ")
					LabNo.Resize(fyne.NewSize(200, 40))
					LabNo.Move(fyne.NewPos(300, 180))

					DoctorID := widget.NewEntry()
					DoctorID.SetPlaceHolder("Enter Doctor ID ")
					DoctorID.Resize(fyne.NewSize(200, 40))
					DoctorID.Move(fyne.NewPos(300, 180))

					DateOfADM := widget.NewEntry()
					DateOfADM.SetPlaceHolder("Enter Date Of ADM ")
					DateOfADM.Resize(fyne.NewSize(200, 40))
					DateOfADM.Move(fyne.NewPos(300, 180))

					DateOfDIS := widget.NewEntry()
					DateOfDIS.SetPlaceHolder("Enter Date Of DIS ")
					DateOfDIS.Resize(fyne.NewSize(200, 40))
					DateOfDIS.Move(fyne.NewPos(300, 180))

					if err != nil {
						fmt.Println(err)
					}
					btn := widget.NewButton("Insert", func() {
						if PatientID.Text == "" || RoomNo.Text == "" || LabNo.Text == "" || DoctorID.Text == "" || DateOfADM.Text == "" || DateOfDIS.Text == "" {
							msg.SetText("All inputs must be filled")
						} else {
							db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+":"+strconv.Itoa(port)+")/"+DataBaseName)
							dd, err := db.Query("insert into inpatient (PatientID, RoomNo, LabNo, DoctorID, DateOfADM, DateOfDIS) Values " + "(" + PatientID.Text + ", '" + RoomNo.Text + "', '" + LabNo.Text + "', '" + DoctorID.Text + "','" + DateOfADM.Text + "','" + DateOfDIS.Text + "')")
							if err != nil {
								fmt.Println(err, dd)
							}
							msg.SetText(" Row inserted Successfully!!! ")
						}
					})
					btn.Resize(fyne.NewSize(200, 40))
					btn.Move(fyne.NewPos(300, 225))

					content = container.NewGridWithRows(8, PatientID, RoomNo, LabNo, DoctorID, DateOfADM, DateOfDIS, btn, msg)
					w.SetContent(container.NewGridWithColumns(
						3,
						ListView,
						tablelist,
						content,
					))
				} else if id == 3 { //laboratory
					msg.SetText("")

					LabNo := widget.NewEntry()
					LabNo.SetPlaceHolder("Enter Lab Number ")
					LabNo.Resize(fyne.NewSize(200, 40))
					LabNo.Move(fyne.NewPos(300, 140))

					DoctorID := widget.NewEntry()
					DoctorID.SetPlaceHolder("Enter Doctor ID ")
					DoctorID.Resize(fyne.NewSize(200, 40))
					DoctorID.Move(fyne.NewPos(300, 180))

					PatientID := widget.NewEntry()
					PatientID.SetPlaceHolder("Enter Patient ID ")
					PatientID.Resize(fyne.NewSize(200, 40))
					PatientID.Move(fyne.NewPos(300, 180))

					labDate := widget.NewEntry()
					labDate.SetPlaceHolder("Enter lab Date ")
					labDate.Resize(fyne.NewSize(200, 40))
					labDate.Move(fyne.NewPos(300, 180))

					Amount := widget.NewEntry()
					Amount.SetPlaceHolder("Enter Amount ")
					Amount.Resize(fyne.NewSize(200, 40))
					Amount.Move(fyne.NewPos(300, 180))
					if err != nil {
						fmt.Println(err)
					}
					btn := widget.NewButton("Insert", func() {
						if LabNo.Text == "" || DoctorID.Text == "" || PatientID.Text == "" || labDate.Text == "" || Amount.Text == "" {
							msg.SetText("All inputs must be filled!!!")
						} else {
							db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+":"+strconv.Itoa(port)+")/"+DataBaseName)
							dd, err := db.Query("insert into laboratory (LabNo, DoctorID, PatientID, labDate, Amount) Values " + "(" + LabNo.Text + ", '" + DoctorID.Text + "', '" + PatientID.Text + "', '" + labDate.Text + "','" + Amount.Text + "')")
							if err != nil {
								fmt.Println(err, dd)
							}
							msg.SetText(" Row inserted Successfully!!! ")
						}
					})
					btn.Resize(fyne.NewSize(200, 40))
					btn.Move(fyne.NewPos(300, 225))

					content = container.NewGridWithRows(7, LabNo, DoctorID, PatientID, labDate, Amount, btn, msg)
					w.SetContent(container.NewGridWithColumns(
						3,
						ListView,
						tablelist,
						content,
					))
				} else if id == 4 { // outpatient
					msg.SetText("")

					PatientID := widget.NewEntry()
					PatientID.SetPlaceHolder("Enter Patient ID ")
					PatientID.Resize(fyne.NewSize(200, 40))
					PatientID.Move(fyne.NewPos(300, 180))

					Pdate := widget.NewEntry()
					Pdate.SetPlaceHolder("Enter P Date ")
					Pdate.Resize(fyne.NewSize(200, 40))
					Pdate.Move(fyne.NewPos(300, 140))

					LabNo := widget.NewEntry()
					LabNo.SetPlaceHolder("Enter Lab Number ")
					LabNo.Resize(fyne.NewSize(200, 40))
					LabNo.Move(fyne.NewPos(300, 180))

					DoctorID := widget.NewEntry()
					DoctorID.SetPlaceHolder("Enter Doctor ID ")
					DoctorID.Resize(fyne.NewSize(200, 40))
					DoctorID.Move(fyne.NewPos(300, 180))

					if err != nil {
						fmt.Println(err)
					}
					btn := widget.NewButton("Insert", func() {
						if PatientID.Text == "" || Pdate.Text == "" || LabNo.Text == "" || DoctorID.Text == "" {
							msg.SetText("All inputs must be filled!!!")
						} else {
							db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+":"+strconv.Itoa(port)+")/"+DataBaseName)
							dd, err := db.Query("insert into outpatient (PatientID, Pdate, LabNo, DoctorID) Values " + "(" + PatientID.Text + ", '" + Pdate.Text + "', '" + LabNo.Text + "'," + DoctorID.Text + ")")
							if err != nil {
								fmt.Println(err, dd)
							}
							msg.SetText(" Row inserted Successfully!!! ")
						}
					})
					btn.Resize(fyne.NewSize(200, 40))
					btn.Move(fyne.NewPos(300, 225))

					content = container.NewGridWithRows(6, PatientID, Pdate, LabNo, DoctorID, btn, msg)
					w.SetContent(container.NewGridWithColumns(
						3,
						ListView,
						tablelist,
						content,
					))
				} else if id == 5 { //patient
					msg.SetText("")

					PatientID := widget.NewEntry()
					PatientID.SetPlaceHolder("Enter Patient ID ")
					PatientID.Resize(fyne.NewSize(200, 40))
					PatientID.Move(fyne.NewPos(300, 180))

					PatientName := widget.NewEntry()
					PatientName.SetPlaceHolder("Enter Patient Name ")
					PatientName.Resize(fyne.NewSize(200, 40))
					PatientName.Move(fyne.NewPos(300, 140))

					Age := widget.NewEntry()
					Age.SetPlaceHolder("Enter Age ")
					Age.Resize(fyne.NewSize(200, 40))
					Age.Move(fyne.NewPos(300, 180))

					Gender := widget.NewEntry()
					Gender.SetPlaceHolder("Enter Gender ")
					Gender.Resize(fyne.NewSize(200, 40))
					Gender.Move(fyne.NewPos(300, 180))

					Address := widget.NewEntry()
					Address.SetPlaceHolder("Enter Address ")
					Address.Resize(fyne.NewSize(200, 40))
					Address.Move(fyne.NewPos(300, 180))

					Disease := widget.NewEntry()
					Disease.SetPlaceHolder("Enter Disease ")
					Disease.Resize(fyne.NewSize(200, 40))
					Disease.Move(fyne.NewPos(300, 180))

					DoctorID := widget.NewEntry()
					DoctorID.SetPlaceHolder("Enter Doctor ID ")
					DoctorID.Resize(fyne.NewSize(200, 40))
					DoctorID.Move(fyne.NewPos(300, 180))

					if err != nil {
						fmt.Println(err)
					}
					btn := widget.NewButton("Insert", func() {
						if PatientID.Text == "" || PatientName.Text == "" || Age.Text == "" || Gender.Text == "" || Address.Text == "" || Disease.Text == "" || DoctorID.Text == "" {
							msg.SetText("All inputs must be filled!!!")
						} else {
							db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+":"+strconv.Itoa(port)+")/"+DataBaseName)
							dd, err := db.Query("insert into patient (PatientID, PatientName, Age, Gender, Address, Disease, DoctorID) Values " + "(" + PatientID.Text + ", '" + PatientName.Text + "', '" + Age.Text + "', '" + Gender.Text + "','" + Address.Text + "','" + Disease.Text + "'," + DoctorID.Text + ")")
							if err != nil {
								fmt.Println(err, dd)
							}
							msg.SetText(" Row inserted Successfully!!! ")
						}

					})
					btn.Resize(fyne.NewSize(200, 40))
					btn.Move(fyne.NewPos(300, 225))

					content = container.NewGridWithRows(9, PatientID, PatientName, Age, Gender, Address, Disease, DoctorID, btn, msg)
					w.SetContent(container.NewGridWithColumns(
						3,
						ListView,
						tablelist,
						content,
					))
				} else if id == 6 { //room
					msg.SetText("")

					room_number := widget.NewEntry()
					room_number.SetPlaceHolder("Enter Room Number ")
					room_number.Resize(fyne.NewSize(200, 40))
					room_number.Move(fyne.NewPos(300, 140))

					room_type := widget.NewEntry()
					room_type.SetPlaceHolder("Enter Room Type ")
					room_type.Resize(fyne.NewSize(200, 40))
					room_type.Move(fyne.NewPos(300, 140))

					room_status := widget.NewEntry()
					room_status.SetPlaceHolder("Enter Room Status ")
					room_status.Resize(fyne.NewSize(200, 40))
					room_status.Move(fyne.NewPos(300, 140))

					PatientID := widget.NewEntry()
					PatientID.SetPlaceHolder("Enter Patient Id ")
					PatientID.Resize(fyne.NewSize(200, 40))
					PatientID.Move(fyne.NewPos(300, 180))

					if err != nil {
						fmt.Println(err)
					}
					btn := widget.NewButton("Insert", func() {
						if room_number.Text == "" || room_type.Text == "" || room_status.Text == "" || PatientID.Text == "" {
							msg.SetText(" All inputs must be filled!!! ")
						} else {
							db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+":"+strconv.Itoa(port)+")/"+DataBaseName)
							dd, err := db.Query("insert into Room (RoomNo, RoomType, RoomStatus, PatientID) Values " + "('" + room_number.Text + "', '" + room_type.Text + "' , '" + room_status.Text + "'," + PatientID.Text + ")")
							if err != nil {
								fmt.Println(err, dd)
							}
							msg.SetText(" Row inserted Successfully!!! ")
						}
					})
					btn.Resize(fyne.NewSize(200, 40))
					btn.Move(fyne.NewPos(300, 225))

					content = container.NewGridWithRows(6, room_number, room_type, room_status, PatientID, btn, msg)
					w.SetContent(container.NewGridWithColumns(
						3,
						ListView,
						tablelist,
						content,
					))
				} else {
					w.SetContent(container.NewGridWithColumns(
						3,
						ListView,
						tablelist,
					))
				}
			}
			w.SetContent(container.NewGridWithColumns(
				3,
				ListView,
				tablelist,
				content,
			))
		} else if id == 3 { // Update

			var databaseList []string

			var table string
			db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+":"+strconv.Itoa(port)+")/"+DataBaseName)
			res, err := db.Query("SHOW TABLES")
			// fmt.Println("tables:")
			var i int = 1
			for res.Next() {
				res.Scan(&table)
				i++
				databaseList = append(databaseList, table)
			}
			if err != nil {
				fmt.Println("Erorr : ", err)
			}
			tablelist := widget.NewList(func() int { return len(databaseList) },
				func() fyne.CanvasObject { return widget.NewLabel("Templates") },
				func(lii widget.ListItemID, co fyne.CanvasObject) {
					co.(*widget.Label).SetText(databaseList[lii])
				})
			var content *fyne.Container = container.NewWithoutLayout()
			tablelist.OnSelected = func(id widget.ListItemID) {
				if id == 0 { // update bill
					msg.SetText("")

					BillNO := widget.NewEntry()
					BillNO.SetPlaceHolder("Enter Bill Number ")
					BillNO.Resize(fyne.NewSize(200, 40))
					BillNO.Move(fyne.NewPos(300, 140))

					PatientID := widget.NewEntry()
					PatientID.SetPlaceHolder("Enter Patient Id ")
					PatientID.Resize(fyne.NewSize(200, 40))
					PatientID.Move(fyne.NewPos(300, 180))

					RoomCharge := widget.NewEntry()
					RoomCharge.SetPlaceHolder("Enter Room Charge ")
					RoomCharge.Resize(fyne.NewSize(200, 40))
					RoomCharge.Move(fyne.NewPos(300, 180))

					DoctorCharge := widget.NewEntry()
					DoctorCharge.SetPlaceHolder("Enter Doctor Charge ")
					DoctorCharge.Resize(fyne.NewSize(200, 40))
					DoctorCharge.Move(fyne.NewPos(300, 180))

					NoOfDays := widget.NewEntry()
					NoOfDays.SetPlaceHolder("Enter Number Of Days ")
					NoOfDays.Resize(fyne.NewSize(200, 40))
					NoOfDays.Move(fyne.NewPos(300, 180))

					LabChargeBill := widget.NewEntry()
					LabChargeBill.SetPlaceHolder("Enter Lab Charge Bill ")
					LabChargeBill.Resize(fyne.NewSize(200, 40))
					LabChargeBill.Move(fyne.NewPos(300, 180))

					if err != nil {
						fmt.Println(err)
					}
					btn := widget.NewButton("Update", func() {
						if BillNO.Text == "" || PatientID.Text == "" || RoomCharge.Text == "" || DoctorCharge.Text == "" || NoOfDays.Text == "" || LabChargeBill.Text == "" {
							msg.SetText("All Fields must be filled!!!")
						} else {
							db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+":"+strconv.Itoa(port)+")/"+DataBaseName)
							dd, err := db.Query("update bill set PatientID = " + PatientID.Text + "," + " RoomCharge = " + RoomCharge.Text + "," + " DoctorCharge = " + DoctorCharge.Text + "," + " NoOfDays = '" + NoOfDays.Text + "'," + " LabChargeBill = '" + LabChargeBill.Text + "' where BillNO = " + BillNO.Text)
							if err != nil {
								fmt.Println(err, dd)
							}
							msg.SetText(" Row Updated Successfully!!! ")
						}
					})
					btn.Resize(fyne.NewSize(200, 40))
					btn.Move(fyne.NewPos(300, 225))

					content = container.NewGridWithRows(8, BillNO, PatientID, RoomCharge, DoctorCharge, NoOfDays, LabChargeBill, btn, msg)
					w.SetContent(container.NewGridWithColumns(
						3,
						ListView,
						tablelist,
						content,
					))

				} else if id == 1 { // Update Doctor
					msg.SetText("")

					DoctorID := widget.NewEntry()
					DoctorID.SetPlaceHolder("Enter Doctor Id ")
					DoctorID.Resize(fyne.NewSize(200, 40))
					DoctorID.Move(fyne.NewPos(300, 140))

					DoctorName := widget.NewEntry()
					DoctorName.SetPlaceHolder("Enter Doctor Name ")
					DoctorName.Resize(fyne.NewSize(200, 40))
					DoctorName.Move(fyne.NewPos(300, 180))

					Age := widget.NewEntry()
					Age.SetPlaceHolder("Enter Age ")
					Age.Resize(fyne.NewSize(200, 40))
					Age.Move(fyne.NewPos(300, 180))

					Gender := widget.NewEntry()
					Gender.SetPlaceHolder("Enter Gender ")
					Gender.Resize(fyne.NewSize(200, 40))
					Gender.Move(fyne.NewPos(300, 180))

					Address := widget.NewEntry()
					Address.SetPlaceHolder("Enter Address ")
					Address.Resize(fyne.NewSize(200, 40))
					Address.Move(fyne.NewPos(300, 180))

					Speciality := widget.NewEntry()
					Speciality.SetPlaceHolder("Enter Speciality ")
					Speciality.Resize(fyne.NewSize(200, 40))
					Speciality.Move(fyne.NewPos(300, 180))

					if err != nil {
						fmt.Println(err)
					}
					btn := widget.NewButton("Update", func() {
						if DoctorID.Text == "" || DoctorName.Text == "" || Age.Text == "" || Gender.Text == "" || Address.Text == "" || Speciality.Text == "" {
							msg.SetText("All Fields must be filled!!!")
						} else {
							db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+":"+strconv.Itoa(port)+")/"+DataBaseName)
							dd, err := db.Query("update doctor set DoctorName = '" + DoctorName.Text + "'," + " Age = '" + Age.Text + "'," + "Gender = '" + Gender.Text + "'," + " Address = '" + Address.Text + "'," + "Speciality = '" + Speciality.Text + "' where DoctorID = '" + DoctorID.Text + " '")
							if err != nil {
								fmt.Println(err, dd)
							}
							msg.SetText(" Row Updated Successfully!!! ")
						}
					})
					btn.Resize(fyne.NewSize(200, 40))
					btn.Move(fyne.NewPos(300, 225))

					content = container.NewGridWithRows(8, DoctorID, DoctorName, Age, Gender, Address, Speciality, btn, msg)
					w.SetContent(container.NewGridWithColumns(
						3,
						ListView,
						tablelist,
						content,
					))
				} else if id == 2 { // inpatient
					msg.SetText("")

					PatientID := widget.NewEntry()
					PatientID.SetPlaceHolder("Enter Patient ID ")
					PatientID.Resize(fyne.NewSize(200, 40))
					PatientID.Move(fyne.NewPos(300, 140))

					RoomNo := widget.NewEntry()
					RoomNo.SetPlaceHolder("Enter Room Number ")
					RoomNo.Resize(fyne.NewSize(200, 40))
					RoomNo.Move(fyne.NewPos(300, 180))

					LabNo := widget.NewEntry()
					LabNo.SetPlaceHolder("Enter Lab Number ")
					LabNo.Resize(fyne.NewSize(200, 40))
					LabNo.Move(fyne.NewPos(300, 180))

					DoctorID := widget.NewEntry()
					DoctorID.SetPlaceHolder("Enter Doctor ID ")
					DoctorID.Resize(fyne.NewSize(200, 40))
					DoctorID.Move(fyne.NewPos(300, 180))

					DateOfADM := widget.NewEntry()
					DateOfADM.SetPlaceHolder("Enter Date Of ADM ")
					DateOfADM.Resize(fyne.NewSize(200, 40))
					DateOfADM.Move(fyne.NewPos(300, 180))

					DateOfDIS := widget.NewEntry()
					DateOfDIS.SetPlaceHolder("Enter Date Of DIS ")
					DateOfDIS.Resize(fyne.NewSize(200, 40))
					DateOfDIS.Move(fyne.NewPos(300, 180))

					if err != nil {
						fmt.Println(err)
					}
					btn := widget.NewButton("Update", func() {
						if PatientID.Text == "" || RoomNo.Text == "" || LabNo.Text == "" || DoctorID.Text == "" || DateOfADM.Text == "" || DateOfDIS.Text == "" {
							msg.SetText("All Fields must be filled!!!")

						} else {
							db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+":"+strconv.Itoa(port)+")/"+DataBaseName)
							dd, err := db.Query("update inpatient set RoomNo = '" + RoomNo.Text + "'," + " LabNo = '" + LabNo.Text + "'," + "DoctorID = '" + DoctorID.Text + "'," + " DateOfADM = '" + DateOfADM.Text + "'," + "DateOfDIS = '" + DateOfDIS.Text + "' where PatientID = '" + PatientID.Text + " '")
							if err != nil {
								fmt.Println(err, dd)
							}
							msg.SetText(" Row Updated Successfully!!! ")
						}

					})
					btn.Resize(fyne.NewSize(200, 40))
					btn.Move(fyne.NewPos(300, 225))

					content = container.NewGridWithRows(8, PatientID, RoomNo, LabNo, DoctorID, DateOfADM, DateOfDIS, btn, msg)
					w.SetContent(container.NewGridWithColumns(
						3,
						ListView,
						tablelist,
						content,
					))
				} else if id == 3 { // laboratory
					msg.SetText("")

					LabNo := widget.NewEntry()
					LabNo.SetPlaceHolder("Enter Lab Number ")
					LabNo.Resize(fyne.NewSize(200, 40))
					LabNo.Move(fyne.NewPos(300, 140))

					DoctorID := widget.NewEntry()
					DoctorID.SetPlaceHolder("Enter Doctor ID ")
					DoctorID.Resize(fyne.NewSize(200, 40))
					DoctorID.Move(fyne.NewPos(300, 180))

					PatientID := widget.NewEntry()
					PatientID.SetPlaceHolder("Enter Patient ID ")
					PatientID.Resize(fyne.NewSize(200, 40))
					PatientID.Move(fyne.NewPos(300, 180))

					labDate := widget.NewEntry()
					labDate.SetPlaceHolder("Enter lab Date ")
					labDate.Resize(fyne.NewSize(200, 40))
					labDate.Move(fyne.NewPos(300, 180))

					Amount := widget.NewEntry()
					Amount.SetPlaceHolder("Enter Amount ")
					Amount.Resize(fyne.NewSize(200, 40))
					Amount.Move(fyne.NewPos(300, 180))
					if err != nil {
						fmt.Println(err)
					}
					btn := widget.NewButton("Update", func() {
						if LabNo.Text == "" || DoctorID.Text == "" || PatientID.Text == "" || labDate.Text == "" || Amount.Text == "" {
							msg.SetText("All Fields must be filled!!!")

						} else {
							db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+":"+strconv.Itoa(port)+")/"+DataBaseName)
							dd, err := db.Query("update laboratory set DoctorID = '" + DoctorID.Text + "'," + " PatientID = '" + PatientID.Text + "'," + "labDate = '" + labDate.Text + "'," + " Amount = '" + Amount.Text + "' where LabNo = '" + LabNo.Text + " '")
							if err != nil {
								fmt.Println(err, dd)
							}
							msg.SetText(" Row Updated Successfully!!! ")
						}

					})
					btn.Resize(fyne.NewSize(200, 40))
					btn.Move(fyne.NewPos(300, 225))

					content = container.NewGridWithRows(7, LabNo, DoctorID, PatientID, labDate, Amount, btn, msg)
					w.SetContent(container.NewGridWithColumns(
						3,
						ListView,
						tablelist,
						content,
					))
				} else if id == 4 { //outpateint
					msg.SetText("")

					PatientID := widget.NewEntry()
					PatientID.SetPlaceHolder("Enter Patient ID ")
					PatientID.Resize(fyne.NewSize(200, 40))
					PatientID.Move(fyne.NewPos(300, 180))

					Pdate := widget.NewEntry()
					Pdate.SetPlaceHolder("Enter P Date ")
					Pdate.Resize(fyne.NewSize(200, 40))
					Pdate.Move(fyne.NewPos(300, 140))

					LabNo := widget.NewEntry()
					LabNo.SetPlaceHolder("Enter Lab Number ")
					LabNo.Resize(fyne.NewSize(200, 40))
					LabNo.Move(fyne.NewPos(300, 180))

					DoctorID := widget.NewEntry()
					DoctorID.SetPlaceHolder("Enter Doctor ID ")
					DoctorID.Resize(fyne.NewSize(200, 40))
					DoctorID.Move(fyne.NewPos(300, 180))

					if err != nil {
						fmt.Println(err)
					}
					btn := widget.NewButton("Update", func() {
						if PatientID.Text == "" || Pdate.Text == "" || LabNo.Text == "" || DoctorID.Text == "" {
							msg.SetText("All Fields must be filled!!!")

						} else {
							db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+":"+strconv.Itoa(port)+")/"+DataBaseName)
							dd, err := db.Query("update outpatient set Pdate = '" + Pdate.Text + "'," + " LabNo = '" + LabNo.Text + "'," + "DoctorID = '" + DoctorID.Text + "' where PatientID = '" + PatientID.Text + " '")
							if err != nil {
								fmt.Println(err, dd)
							}
							msg.SetText(" Row Updated Successfully!!! ")

						}

					})
					btn.Resize(fyne.NewSize(200, 40))
					btn.Move(fyne.NewPos(300, 225))

					content = container.NewGridWithRows(6, PatientID, Pdate, LabNo, DoctorID, btn, msg)
					w.SetContent(container.NewGridWithColumns(
						3,
						ListView,
						tablelist,
						content,
					))
				} else if id == 5 { //pateint
					msg.SetText("")

					PatientID := widget.NewEntry()
					PatientID.SetPlaceHolder("Enter Patient ID ")
					PatientID.Resize(fyne.NewSize(200, 40))
					PatientID.Move(fyne.NewPos(300, 180))

					PatientName := widget.NewEntry()
					PatientName.SetPlaceHolder("Enter Patient Name ")
					PatientName.Resize(fyne.NewSize(200, 40))
					PatientName.Move(fyne.NewPos(300, 140))

					Age := widget.NewEntry()
					Age.SetPlaceHolder("Enter Age ")
					Age.Resize(fyne.NewSize(200, 40))
					Age.Move(fyne.NewPos(300, 180))

					Gender := widget.NewEntry()
					Gender.SetPlaceHolder("Enter Gender ")
					Gender.Resize(fyne.NewSize(200, 40))
					Gender.Move(fyne.NewPos(300, 180))

					Address := widget.NewEntry()
					Address.SetPlaceHolder("Enter Address ")
					Address.Resize(fyne.NewSize(200, 40))
					Address.Move(fyne.NewPos(300, 180))

					Disease := widget.NewEntry()
					Disease.SetPlaceHolder("Enter Disease ")
					Disease.Resize(fyne.NewSize(200, 40))
					Disease.Move(fyne.NewPos(300, 180))

					DoctorID := widget.NewEntry()
					DoctorID.SetPlaceHolder("Enter Doctor ID ")
					DoctorID.Resize(fyne.NewSize(200, 40))
					DoctorID.Move(fyne.NewPos(300, 180))

					if err != nil {
						fmt.Println(err)
					}
					btn := widget.NewButton("Update", func() {
						if PatientID.Text == "" || PatientName.Text == "" || Age.Text == "" || Gender.Text == "" || Address.Text == "" || Disease.Text == "" || DoctorID.Text == "" {
							msg.SetText("All Fields must be filled!!!")
						} else {
							db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+":"+strconv.Itoa(port)+")/"+DataBaseName)
							dd, err := db.Query("update patient set PatientName = '" + PatientName.Text + "'," + " Age = '" + Age.Text + "'," + "Gender = '" + Gender.Text + "'," + " Address = '" + Address.Text + "'," + "Disease = '" + Disease.Text + "' where PatientID = '" + PatientID.Text + " '")
							if err != nil {
								fmt.Println(err, dd)
							}
							msg.SetText(" Row Updated Successfully!!! ")
						}
					})
					btn.Resize(fyne.NewSize(200, 40))
					btn.Move(fyne.NewPos(300, 225))

					content = container.NewGridWithRows(9, PatientID, PatientName, Age, Gender, Address, Disease, DoctorID, btn, msg)
					w.SetContent(container.NewGridWithColumns(
						3,
						ListView,
						tablelist,
						content,
					))
				} else if id == 6 { //room
					msg.SetText("")

					RoomNo := widget.NewEntry()
					RoomNo.SetPlaceHolder("Enter Room Number ")
					RoomNo.Resize(fyne.NewSize(200, 40))
					RoomNo.Move(fyne.NewPos(300, 140))

					RoomType := widget.NewEntry()
					RoomType.SetPlaceHolder("Enter Room Type ")
					RoomType.Resize(fyne.NewSize(200, 40))
					RoomType.Move(fyne.NewPos(300, 140))

					RoomStatus := widget.NewEntry()
					RoomStatus.SetPlaceHolder("Enter Room Status ")
					RoomStatus.Resize(fyne.NewSize(200, 40))
					RoomStatus.Move(fyne.NewPos(300, 140))

					PatientID := widget.NewEntry()
					PatientID.SetPlaceHolder("Enter Patient Id ")
					PatientID.Resize(fyne.NewSize(200, 40))
					PatientID.Move(fyne.NewPos(300, 180))

					if err != nil {
						fmt.Println(err)
					}
					btn := widget.NewButton("Update", func() {
						if RoomNo.Text == "" || RoomType.Text == "" || RoomStatus.Text == "" || PatientID.Text == "" {
							msg.SetText("All Fields must be filled!!!")
						} else {
							db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+":"+strconv.Itoa(port)+")/"+DataBaseName)
							dd, err := db.Query("update room set RoomType = '" + RoomType.Text + "'," + " RoomStatus = '" + RoomStatus.Text + "'," + "PatientID = '" + PatientID.Text + "' where RoomNo = " + RoomNo.Text + " ")
							if err != nil {
								fmt.Println(err, dd)
							}
							msg.SetText(" Row Updated Successfully!!! ")
						}
					})
					btn.Resize(fyne.NewSize(200, 40))
					btn.Move(fyne.NewPos(300, 225))

					content = container.NewGridWithRows(6, RoomNo, RoomType, RoomStatus, PatientID, btn, msg)
					w.SetContent(container.NewGridWithColumns(
						3,
						ListView,
						tablelist,
						content,
					))
				} else {
					w.SetContent(container.NewGridWithColumns(
						3,
						ListView,
						tablelist,
					))
				}
			}
			w.SetContent(container.NewGridWithColumns(
				3,
				ListView,
				tablelist,
				content,
			))
		} else if id == 4 { //select
			var databaseList []string
			// var recordsList []string
			// var alldata string
			// multilineEntry := widget.NewMultiLineEntry()

			var table string
			db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+":"+strconv.Itoa(port)+")/"+DataBaseName)
			res, err := db.Query("SHOW TABLES")
			var i int = 1
			for res.Next() {
				res.Scan(&table)
				i++
				databaseList = append(databaseList, table)

			}
			if err != nil {
				fmt.Println("Erorr : ", err)
			}
			tablelist := widget.NewList(func() int { return len(databaseList) },
				func() fyne.CanvasObject { return widget.NewLabel("Templates") },
				func(lii widget.ListItemID, co fyne.CanvasObject) {
					co.(*widget.Label).SetText(databaseList[lii])
				})
			tablelist.OnSelected = func(id widget.ListItemID) {
				if id == 0 {
					type billstruct struct {
						billno        int
						patientId     int
						roomcharge    string
						DoctorCharge  string
						NoOfDays      string
						LabChargeBill string
					}
					res, err := db.Query("select * from bill")
					if err != nil {
						panic(err)
					}
					// recordsList = nil

					f, err := os.Create("C:/Users/Mario Ehab/Desktop/data.txt")
					var all string
					for res.Next() {
						var billrecord billstruct
						err = res.Scan(&billrecord.billno, &billrecord.patientId, &billrecord.roomcharge, &billrecord.DoctorCharge, &billrecord.LabChargeBill, &billrecord.NoOfDays)

						all = "Bill Number : " + strconv.Itoa(billrecord.billno) + " \nPatient ID :  " + strconv.Itoa(billrecord.patientId) + " \nRoom Charge : " + billrecord.roomcharge + " \nDoctor Charge : " + billrecord.DoctorCharge + " \nLab Charge Bill : " + billrecord.LabChargeBill + " \nNumber Of Days : " + billrecord.NoOfDays
						_, err = f.WriteString(all + "\n" + "\n")
					}

					data, _ := ioutil.ReadFile("C:/Users/Mario Ehab/Desktop/data.txt")

					result := fyne.NewStaticResource("Data in Tables", data)
					entry := widget.NewMultiLineEntry()
					entry.SetText(string(result.StaticContent))
					w := fyne.CurrentApp().NewWindow(
						string(result.StaticName))
					w.SetContent(container.NewScroll(entry))
					w.Resize(fyne.NewSize(400, 400))
					w.Show()
				} else if id == 1 {
					type doctorstruct struct {
						DoctorID   int
						DoctorName string
						Age        string
						Gender     string
						Address    string
						Speciality string
					}
					res, err := db.Query("select * from doctor")
					if err != nil {
						panic(err)
					}
					// recordsList = nil

					f, err := os.Create("C:/Users/Mario Ehab/Desktop/data.txt")
					var all string
					for res.Next() {
						var doctrorecord doctorstruct
						err = res.Scan(&doctrorecord.DoctorID, &doctrorecord.DoctorName, &doctrorecord.Age, &doctrorecord.Gender, &doctrorecord.Address, &doctrorecord.Speciality)

						all = " Doctor ID : " + strconv.Itoa(doctrorecord.DoctorID) + " \nDoctor Name : " + doctrorecord.DoctorName + " \nDoctor Age " + doctrorecord.Age + " \nGender : " + doctrorecord.Gender + " \nAddress :  " + doctrorecord.Address + " \nSpeciality : " + doctrorecord.Speciality
						_, err = f.WriteString(all + "\n" + "\n")
					}

					data, _ := ioutil.ReadFile("C:/Users/Mario Ehab/Desktop/data.txt")

					result := fyne.NewStaticResource("Data in Tables", data)
					entry := widget.NewMultiLineEntry()
					entry.SetText(string(result.StaticContent))
					w := fyne.CurrentApp().NewWindow(
						string(result.StaticName))
					w.SetContent(container.NewScroll(entry))
					w.Resize(fyne.NewSize(400, 400))
					w.Show()

				} else if id == 2 {
					type inpatientstruct struct {
						PatientID   int
						RoomNo      int
						LabNo       int
						DoctorID    int
						DateOfADM   string
						DateOfDIS   string
						inpatientID int
					}
					res, err := db.Query("select * from inpatient")
					if err != nil {
						panic(err)
					}
					// recordsList = nil

					f, err := os.Create("C:/Users/Mario Ehab/Desktop/data.txt")
					var all string
					// details = "PatientID	RoomNo   LabNo   DoctorID	DateOfADM	DateOfDIS"
					for res.Next() {
						var inpatientrecord inpatientstruct
						err = res.Scan(&inpatientrecord.PatientID, &inpatientrecord.RoomNo, &inpatientrecord.LabNo, &inpatientrecord.DoctorID, &inpatientrecord.DateOfADM, &inpatientrecord.DateOfDIS, &inpatientrecord.inpatientID)

						all = "PatientID : " + strconv.Itoa(inpatientrecord.PatientID) + "\n RoomNo: " + strconv.Itoa(inpatientrecord.RoomNo) + "\n LabNo : " + strconv.Itoa(inpatientrecord.LabNo) + "\n DoctorID: " + strconv.Itoa(inpatientrecord.DoctorID) + " \n DateOfADM : " + inpatientrecord.DateOfADM + "\n DateOfDIS : " + inpatientrecord.DateOfDIS + "\n inpatient ID : " + strconv.Itoa(inpatientrecord.inpatientID)
						_, err = f.WriteString("\n" + all + "\n" + "\n")
					}

					data, _ := ioutil.ReadFile("C:/Users/Mario Ehab/Desktop/data.txt")

					result := fyne.NewStaticResource("Data in Tables", data)
					entry := widget.NewMultiLineEntry()
					entry.SetText(string(result.StaticContent))
					w := fyne.CurrentApp().NewWindow(
						string(result.StaticName))
					w.SetContent(container.NewScroll(entry))
					w.Resize(fyne.NewSize(400, 400))
					w.Show()
				} else if id == 3 {
					type labtstruct struct {
						LabNo     int
						DoctorID  int
						PatientID int
						labDate   string
						Amount    string
					}
					res, err := db.Query("select * from laboratory")
					if err != nil {
						panic(err)
					}
					// recordsList = nil

					f, err := os.Create("C:/Users/Mario Ehab/Desktop/data.txt")
					var all string
					for res.Next() {
						var labrecord labtstruct
						err = res.Scan(&labrecord.LabNo, &labrecord.DoctorID, &labrecord.PatientID, &labrecord.labDate, &labrecord.Amount)

						all = "Lab Number : " + strconv.Itoa(labrecord.LabNo) + "\n Doctor ID: " + strconv.Itoa(labrecord.DoctorID) + "\n Patient ID : " + strconv.Itoa(labrecord.PatientID) + "\n lab Date: " + labrecord.labDate + " \n Amount : " + labrecord.Amount
						_, err = f.WriteString("\n" + all + "\n" + "\n")
					}

					data, _ := ioutil.ReadFile("C:/Users/Mario Ehab/Desktop/data.txt")

					result := fyne.NewStaticResource("Data in Tables", data)
					entry := widget.NewMultiLineEntry()
					entry.SetText(string(result.StaticContent))
					w := fyne.CurrentApp().NewWindow(
						string(result.StaticName))
					w.SetContent(container.NewScroll(entry))
					w.Resize(fyne.NewSize(400, 400))
					w.Show()
				} else if id == 4 {
					type outpatientstruct struct {
						PatientID    int
						Pdate        string
						LabNo        int
						DoctorID     int
						OutpatientID int
					}
					res, err := db.Query("select * from outpatient")
					if err != nil {
						panic(err)
					}
					// recordsList = nil

					f, err := os.Create("C:/Users/Mario Ehab/Desktop/data.txt")
					var all string
					// details = "PatientID	RoomNo   LabNo   DoctorID	DateOfADM	DateOfDIS"
					for res.Next() {
						var outpatientrecord outpatientstruct
						err = res.Scan(&outpatientrecord.PatientID, &outpatientrecord.Pdate, &outpatientrecord.LabNo, &outpatientrecord.DoctorID, &outpatientrecord.OutpatientID)

						all = "PatientID : " + strconv.Itoa(outpatientrecord.PatientID) + "\n P date: " + outpatientrecord.Pdate + "\n LabNo : " + strconv.Itoa(outpatientrecord.LabNo) + "\n DoctorID : " + strconv.Itoa(outpatientrecord.DoctorID) + "\n OutPatient ID : " + strconv.Itoa(outpatientrecord.OutpatientID)
						_, err = f.WriteString("\n" + all + "\n" + "\n")
					}

					data, _ := ioutil.ReadFile("C:/Users/Mario Ehab/Desktop/data.txt")

					result := fyne.NewStaticResource("Data in Tables", data)
					entry := widget.NewMultiLineEntry()
					entry.SetText(string(result.StaticContent))
					w := fyne.CurrentApp().NewWindow(
						string(result.StaticName))
					w.SetContent(container.NewScroll(entry))
					w.Resize(fyne.NewSize(400, 400))
					w.Show()
				} else if id == 5 {
					type patientstruct struct {
						PatientID   int
						PatientName string
						Age         string
						Gender      string
						Address     string
						Disease     string
						DoctorID    int
					}
					res, err := db.Query("select * from patient")
					if err != nil {
						panic(err)
					}
					// recordsList = nil

					f, err := os.Create("C:/Users/Mario Ehab/Desktop/data.txt")
					var all string
					for res.Next() {
						var patientrecord patientstruct
						err = res.Scan(&patientrecord.PatientID, &patientrecord.PatientName, &patientrecord.Age, &patientrecord.Gender, &patientrecord.Address, &patientrecord.Disease, &patientrecord.DoctorID)

						all = "PatientID : " + strconv.Itoa(patientrecord.PatientID) + "\n Patient Name: " + patientrecord.PatientName + "\n Age : " + patientrecord.Age + "\n Gender: " + patientrecord.Gender + "\n Address: " + patientrecord.Address + "\n Doctor ID: " + strconv.Itoa(patientrecord.DoctorID)
						_, err = f.WriteString("\n" + all + "\n" + "\n")
					}

					data, _ := ioutil.ReadFile("C:/Users/Mario Ehab/Desktop/data.txt")

					result := fyne.NewStaticResource("Data in Tables", data)
					entry := widget.NewMultiLineEntry()
					entry.SetText(string(result.StaticContent))
					w := fyne.CurrentApp().NewWindow(
						string(result.StaticName))
					w.SetContent(container.NewScroll(entry))
					w.Resize(fyne.NewSize(400, 400))
					w.Show()
				} else if id == 6 {
					type roomstruct struct {
						RoomNo     int
						RoomType   string
						RoomStatus string
						PatientID  int
					}
					res, err := db.Query("select * from room")
					if err != nil {
						panic(err)
					}
					// recordsList = nil

					f, err := os.Create("C:/Users/Mario Ehab/Desktop/data.txt")
					var all string
					for res.Next() {
						var roomrecord roomstruct
						err = res.Scan(&roomrecord.RoomNo, &roomrecord.RoomType, &roomrecord.RoomStatus, &roomrecord.PatientID)

						all = "Room Number : " + strconv.Itoa(roomrecord.RoomNo) + "\n Room Type: " + roomrecord.RoomType + "\n Room Status : " + roomrecord.RoomStatus + "\n Patient ID : " + strconv.Itoa(roomrecord.PatientID)
						_, err = f.WriteString("\n" + all + "\n" + "\n")
					}

					data, _ := ioutil.ReadFile("C:/Users/Mario Ehab/Desktop/data.txt")

					result := fyne.NewStaticResource("Data in Tables", data)
					entry := widget.NewMultiLineEntry()
					entry.SetText(string(result.StaticContent))
					w := fyne.CurrentApp().NewWindow(
						string(result.StaticName))
					w.SetContent(container.NewScroll(entry))
					w.Resize(fyne.NewSize(400, 400))
					w.Show()
				}
			}
			w.SetContent(container.NewGridWithColumns(
				2,
				ListView,
				tablelist,
			))
		} else if id == 5 { // Delete
			var databaseList []string
			var table string
			db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+":"+strconv.Itoa(port)+")/"+DataBaseName)
			res, err := db.Query("SHOW TABLES")
			// fmt.Println("tables:")
			var i int = 1
			for res.Next() {
				res.Scan(&table)
				i++
				databaseList = append(databaseList, table)
			}
			if err != nil {
				fmt.Println("Erorr : ", err)
			}
			tablelist := widget.NewList(func() int { return len(databaseList) },
				func() fyne.CanvasObject { return widget.NewLabel("Templates") },
				func(lii widget.ListItemID, co fyne.CanvasObject) {
					co.(*widget.Label).SetText(databaseList[lii])
				})
			var content *fyne.Container = container.NewWithoutLayout()
			tablelist.OnSelected = func(id widget.ListItemID) {
				recordID := widget.NewEntry()
				recordID.SetPlaceHolder("Enter record ID ")
				recordID.Resize(fyne.NewSize(200, 40))
				recordID.Move(fyne.NewPos(300, 180))
				if id == 0 {
					msg.SetText("")
					delbtn := widget.NewButton("Delete", func() {
						if recordID.Text == "" {
							msg.SetText(" Input must be filled!!!")

						} else {
							db, err := sql.Open("mysql", userName+":"+password+"@tcp(127.0.0.1:"+strconv.Itoa(port)+")/"+DataBaseName)
							if err != nil {
								panic(err)
							}
							db.Query("delete from bill where BillNO = " + recordID.Text)
							msg.SetText(" RECORD Deleted Successfully!!...")
						}

					})
					delbtn.Resize(fyne.NewSize(200, 40))
					delbtn.Move(fyne.NewPos(300, 180))

					content = container.NewGridWithRows(3, recordID, delbtn, msg)
				} else if id == 1 {
					msg.SetText("")
					delbtn := widget.NewButton("Delete", func() {
						if recordID.Text == "" {
							msg.SetText("Input must be filled!!!")
						} else {
							db, err := sql.Open("mysql", userName+":"+password+"@tcp(127.0.0.1:"+strconv.Itoa(port)+")/"+DataBaseName)
							if err != nil {
								panic(err)
							}

							db.Query("delete from doctor where DoctorID = " + recordID.Text)
							msg.SetText(" RECORD Deleted Successfully!!...")
						}

					})
					delbtn.Resize(fyne.NewSize(200, 40))
					delbtn.Move(fyne.NewPos(300, 180))

					content = container.NewGridWithRows(3, recordID, delbtn, msg)
				} else if id == 2 {
					msg.SetText("")
					delbtn := widget.NewButton("Delete", func() {
						if recordID.Text == "" {
							msg.SetText("Input must be filled!!!")
						} else {
							db, err := sql.Open("mysql", userName+":"+password+"@tcp(127.0.0.1:"+strconv.Itoa(port)+")/"+DataBaseName)
							if err != nil {
								panic(err)
							}

							db.Query("delete from inpatient where inpatientID = " + recordID.Text)
							msg.SetText(" RECORD Deleted Successfully!!...")
						}

					})
					delbtn.Resize(fyne.NewSize(200, 40))
					delbtn.Move(fyne.NewPos(300, 180))

					content = container.NewGridWithRows(3, recordID, delbtn, msg)
				} else if id == 3 {
					msg.SetText("")

					delbtn := widget.NewButton("Delete", func() {
						if recordID.Text == "" {
							msg.SetText("Input must be filled!!!")

						} else {
							db, err := sql.Open("mysql", userName+":"+password+"@tcp(127.0.0.1:"+strconv.Itoa(port)+")/"+DataBaseName)
							if err != nil {
								panic(err)
							}

							db.Query("delete from laboratory where LabNo = " + recordID.Text)
							msg.SetText(" RECORD Deleted Successfully!!...")
						}

					})
					delbtn.Resize(fyne.NewSize(200, 40))
					delbtn.Move(fyne.NewPos(300, 180))

					content = container.NewGridWithRows(3, recordID, delbtn, msg)
				} else if id == 4 {
					msg.SetText("")

					delbtn := widget.NewButton("Delete", func() {
						if recordID.Text == "" {
							msg.SetText("Input must be filled!!!")

						} else {
							db, err := sql.Open("mysql", userName+":"+password+"@tcp(127.0.0.1:"+strconv.Itoa(port)+")/"+DataBaseName)
							if err != nil {
								panic(err)
							}

							db.Query("delete from outpatient where OutpatientID = " + recordID.Text)
							msg.SetText(" RECORD Deleted Successfully!!...")
						}

					})
					delbtn.Resize(fyne.NewSize(200, 40))
					delbtn.Move(fyne.NewPos(300, 180))

					content = container.NewGridWithRows(3, recordID, delbtn, msg)
				} else if id == 5 {
					msg.SetText("")
					delbtn := widget.NewButton("Delete", func() {
						if recordID.Text == "" {
							msg.SetText("Input must be filled!!!")

						} else {
							db, err := sql.Open("mysql", userName+":"+password+"@tcp(127.0.0.1:"+strconv.Itoa(port)+")/"+DataBaseName)
							if err != nil {
								panic(err)
							}

							db.Query("delete from patient where PatientID = " + recordID.Text)
							msg.SetText(" RECORD Deleted Successfully!!...")
						}
					})
					delbtn.Resize(fyne.NewSize(200, 40))
					delbtn.Move(fyne.NewPos(300, 180))

					content = container.NewGridWithRows(3, recordID, delbtn, msg)
				} else {
					msg.SetText("")
					delbtn := widget.NewButton("Delete", func() {
						if recordID.Text == "" {
							msg.SetText("Input must be filled!!!")
						} else {
							db, err := sql.Open("mysql", userName+":"+password+"@tcp(127.0.0.1:"+strconv.Itoa(port)+")/"+DataBaseName)
							if err != nil {
								panic(err)
							}
							db.Query("delete from room where RoomNo = " + recordID.Text)
							fmt.Println("----> RECORD Deleted Successfully!!...")
						}
					})
					delbtn.Resize(fyne.NewSize(200, 40))
					delbtn.Move(fyne.NewPos(300, 180))

					content = container.NewGridWithRows(3, recordID, delbtn, msg)
				}
				w.SetContent(
					container.NewGridWithColumns(
						3,
						ListView,
						tablelist,
						content,
					))
			}
			w.SetContent(
				container.NewGridWithColumns(
					3,
					ListView,
					tablelist,
					content,
				))
		} else if id == 6 { //searchbyid
			var databaseList []string

			var table string

			db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+":"+strconv.Itoa(port)+")/"+DataBaseName)

			res, err := db.Query("SHOW TABLES")
			for res.Next() {
				res.Scan(&table)
				databaseList = append(databaseList, table)
			}
			if err != nil {
				fmt.Println("Erorr : ", err)
			}
			tablelist := widget.NewList(func() int { return len(databaseList) },
				func() fyne.CanvasObject { return widget.NewLabel("Templates") },
				func(lii widget.ListItemID, co fyne.CanvasObject) {
					co.(*widget.Label).SetText(databaseList[lii])
				})
			w.SetContent(container.NewGridWithColumns(
				3,
				ListView,
				tablelist,
				searchMsg,
			))
			var content *fyne.Container = container.NewWithoutLayout()
			tablelist.OnSelected = func(id widget.ListItemID) {
				recordID := widget.NewEntry()
				recordID.SetPlaceHolder("Enter record ID ")
				recordID.Resize(fyne.NewSize(200, 40))
				recordID.Move(fyne.NewPos(300, 180))
				if id == 0 {
					msg.SetText("")
					type billstruct struct {
						billno        int
						patientId     int
						roomcharge    string
						DoctorCharge  string
						NoOfDays      string
						LabChargeBill string
					}

					searchbtn := widget.NewButton("Search", func() {
						if recordID.Text == "" {
							msg.SetText("Input must be filled!!!")

						} else {
							db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+":"+strconv.Itoa(port)+")/"+DataBaseName)
							if err != nil {
								panic(err)
							}
							res, err = db.Query("select * from bill where BillNO = " + recordID.Text)
							f, err := os.Create("C:/Users/Mario Ehab/Desktop/data.txt")
							var all string
							for res.Next() {
								var billrecord billstruct
								err = res.Scan(&billrecord.billno, &billrecord.patientId, &billrecord.roomcharge, &billrecord.DoctorCharge, &billrecord.LabChargeBill, &billrecord.NoOfDays)

								all = "Bill Number : " + strconv.Itoa(billrecord.billno) + " \nPatient ID :  " + strconv.Itoa(billrecord.patientId) + " \nRoom Charge : " + billrecord.roomcharge + " \nDoctor Charge : " + billrecord.DoctorCharge + " \nLab Charge Bill : " + billrecord.LabChargeBill + " \nNumber Of Days : " + billrecord.NoOfDays
								_, err = f.WriteString(all + "\n" + "\n")
							}
							if err != nil {
								panic(err)
							}
							data, _ := ioutil.ReadFile("C:/Users/Mario Ehab/Desktop/data.txt")

							result := fyne.NewStaticResource("Data in Tables", data)
							entry := widget.NewMultiLineEntry()
							entry.SetText(string(result.StaticContent))
							w := fyne.CurrentApp().NewWindow(
								string(result.StaticName))
							w.SetContent(container.NewScroll(entry))
							w.Resize(fyne.NewSize(400, 400))
							w.Show()
							msg.SetText("DONE")
						}
					})
					searchbtn.Resize(fyne.NewSize(200, 40))
					searchbtn.Move(fyne.NewPos(300, 180))

					content = container.NewGridWithRows(3, recordID, searchbtn, msg)
				} else if id == 1 {
					msg.SetText("")
					type doctorstruct struct {
						DoctorID   int
						DoctorName string
						Age        string
						Gender     string
						Address    string
						Speciality string
					}
					searchbtn := widget.NewButton("Search", func() {
						if recordID.Text == "" {
							msg.SetText("Input must be filled!!!")
						} else {
							db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+":"+strconv.Itoa(port)+")/"+DataBaseName)
							if err != nil {
								panic(err)
							}
							res, err = db.Query("select * from doctor where DoctorID = " + recordID.Text)
							f, err := os.Create("C:/Users/Mario Ehab/Desktop/data.txt")
							var all string
							for res.Next() {
								var doctrorecord doctorstruct
								err = res.Scan(&doctrorecord.DoctorID, &doctrorecord.DoctorName, &doctrorecord.Age, &doctrorecord.Gender, &doctrorecord.Address, &doctrorecord.Speciality)

								all = " Doctor ID : " + strconv.Itoa(doctrorecord.DoctorID) + " \nDoctor Name : " + doctrorecord.DoctorName + " \nDoctor Age " + doctrorecord.Age + " \nGender : " + doctrorecord.Gender + " \nAddress :  " + doctrorecord.Address + " \nSpeciality : " + doctrorecord.Speciality
								_, err = f.WriteString(all + "\n" + "\n")
							}
							if err != nil {
								panic(err)
							}
							data, _ := ioutil.ReadFile("C:/Users/Mario Ehab/Desktop/data.txt")

							result := fyne.NewStaticResource("Data in Tables", data)
							entry := widget.NewMultiLineEntry()
							entry.SetText(string(result.StaticContent))
							w := fyne.CurrentApp().NewWindow(
								string(result.StaticName))
							w.SetContent(container.NewScroll(entry))
							w.Resize(fyne.NewSize(400, 400))
							w.Show()

							fmt.Println("----> RECORD Searched Successfully!!...")
						}

					})
					searchbtn.Resize(fyne.NewSize(200, 40))
					searchbtn.Move(fyne.NewPos(300, 180))

					content = container.NewGridWithRows(3, recordID, searchbtn, msg)
				} else if id == 2 {
					msg.SetText("")
					type inpatientstruct struct {
						PatientID   int
						RoomNo      int
						LabNo       int
						DoctorID    int
						DateOfADM   string
						DateOfDIS   string
						inpatientID int
					}

					searchbtn := widget.NewButton("Search", func() {
						if recordID.Text == "" {
							msg.SetText("Input must be filled!!!")
						} else {
							db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+":"+strconv.Itoa(port)+")/"+DataBaseName)
							if err != nil {
								panic(err)
							}
							res, err = db.Query("select * from inpatient where inpatientID = " + recordID.Text)
							f, err := os.Create("C:/Users/Mario Ehab/Desktop/data.txt")
							var all string
							for res.Next() {
								var inpatientrecord inpatientstruct
								err = res.Scan(&inpatientrecord.PatientID, &inpatientrecord.RoomNo, &inpatientrecord.LabNo, &inpatientrecord.DoctorID, &inpatientrecord.DateOfADM, &inpatientrecord.DateOfDIS, &inpatientrecord.inpatientID)

								all = "PatientID : " + strconv.Itoa(inpatientrecord.PatientID) + "\n RoomNo: " + strconv.Itoa(inpatientrecord.RoomNo) + "\n LabNo : " + strconv.Itoa(inpatientrecord.LabNo) + "\n DoctorID: " + strconv.Itoa(inpatientrecord.DoctorID) + " \n DateOfADM : " + inpatientrecord.DateOfADM + "\n DateOfDIS : " + inpatientrecord.DateOfDIS + "\n Inpatient ID : " + strconv.Itoa(inpatientrecord.inpatientID)
								_, err = f.WriteString("\n" + all + "\n" + "\n")
							}
							if err != nil {
								panic(err)
							}
							data, _ := ioutil.ReadFile("C:/Users/Mario Ehab/Desktop/data.txt")

							result := fyne.NewStaticResource("Data in Tables", data)
							entry := widget.NewMultiLineEntry()
							entry.SetText(string(result.StaticContent))
							w := fyne.CurrentApp().NewWindow(
								string(result.StaticName))
							w.SetContent(container.NewScroll(entry))
							w.Resize(fyne.NewSize(400, 400))
							w.Show()

							fmt.Println("----> RECORD Searched Successfully!!...")
						}

					})
					searchbtn.Resize(fyne.NewSize(200, 40))
					searchbtn.Move(fyne.NewPos(300, 180))

					content = container.NewGridWithRows(3, recordID, searchbtn, msg)
				} else if id == 3 {
					msg.SetText("")
					type labtstruct struct {
						LabNo     int
						DoctorID  int
						PatientID int
						labDate   string
						Amount    string
					}
					searchbtn := widget.NewButton("Search", func() {
						if recordID.Text == "" {
							msg.SetText("Input must be filled!!!")
						} else {
							db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+":"+strconv.Itoa(port)+")/"+DataBaseName)
							if err != nil {
								panic(err)
							}
							res, err = db.Query("select * from laboratory where LabNo = " + recordID.Text)
							f, err := os.Create("C:/Users/Mario Ehab/Desktop/data.txt")
							var all string
							for res.Next() {
								var labrecord labtstruct
								err = res.Scan(&labrecord.LabNo, &labrecord.DoctorID, &labrecord.PatientID, &labrecord.labDate, &labrecord.Amount)

								all = "Lab Number : " + strconv.Itoa(labrecord.LabNo) + "\n Doctor ID: " + strconv.Itoa(labrecord.DoctorID) + "\n Patient ID : " + strconv.Itoa(labrecord.PatientID) + "\n lab Date: " + labrecord.labDate + " \n Amount : " + labrecord.Amount
								_, err = f.WriteString("\n" + all + "\n" + "\n")
							}
							if err != nil {
								panic(err)
							}
							data, _ := ioutil.ReadFile("C:/Users/Mario Ehab/Desktop/data.txt")

							result := fyne.NewStaticResource("Data in Tables", data)
							entry := widget.NewMultiLineEntry()
							entry.SetText(string(result.StaticContent))
							w := fyne.CurrentApp().NewWindow(
								string(result.StaticName))
							w.SetContent(container.NewScroll(entry))
							w.Resize(fyne.NewSize(400, 400))
							w.Show()

							fmt.Println("----> RECORD Searched Successfully!!...")
						}

					})
					searchbtn.Resize(fyne.NewSize(200, 40))
					searchbtn.Move(fyne.NewPos(300, 180))

					content = container.NewGridWithRows(3, recordID, searchbtn, msg)
				} else if id == 4 {
					msg.SetText("")

					type outpatientstruct struct {
						PatientID    int
						Pdate        string
						LabNo        int
						DoctorID     int
						OutpatientID int
					}
					searchbtn := widget.NewButton("Search", func() {
						if recordID.Text == "" {
							msg.SetText("Input must be filled")
						} else {
							db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+":"+strconv.Itoa(port)+")/"+DataBaseName)
							if err != nil {
								panic(err)
							}
							res, err = db.Query("select * from outpatient where OutpatientID = " + recordID.Text)
							f, err := os.Create("C:/Users/Mario Ehab/Desktop/data.txt")
							var all string
							for res.Next() {
								var outpatientrecord outpatientstruct
								err = res.Scan(&outpatientrecord.PatientID, &outpatientrecord.Pdate, &outpatientrecord.LabNo, &outpatientrecord.DoctorID, &outpatientrecord.OutpatientID)

								all = "PatientID : " + strconv.Itoa(outpatientrecord.PatientID) + "\n P date: " + outpatientrecord.Pdate + "\n LabNo : " + strconv.Itoa(outpatientrecord.LabNo) + "\n DoctorID: " + strconv.Itoa(outpatientrecord.DoctorID) + "\n outPatient ID : " + strconv.Itoa(outpatientrecord.OutpatientID)
								_, err = f.WriteString("\n" + all + "\n" + "\n")
							}

							if err != nil {
								panic(err)
							}
							data, _ := ioutil.ReadFile("C:/Users/Mario Ehab/Desktop/data.txt")

							result := fyne.NewStaticResource("Data in Tables", data)
							entry := widget.NewMultiLineEntry()
							entry.SetText(string(result.StaticContent))
							w := fyne.CurrentApp().NewWindow(
								string(result.StaticName))
							w.SetContent(container.NewScroll(entry))
							w.Resize(fyne.NewSize(400, 400))
							w.Show()

							fmt.Println("----> RECORD Searched Successfully!!...")
						}

					})
					searchbtn.Resize(fyne.NewSize(200, 40))
					searchbtn.Move(fyne.NewPos(300, 180))

					content = container.NewGridWithRows(3, recordID, searchbtn, msg)
				} else if id == 5 {
					msg.SetText("")

					type patientstruct struct {
						PatientID   int
						PatientName string
						Age         string
						Gender      string
						Address     string
						Disease     string
						DoctorID    int
					}
					searchbtn := widget.NewButton("Search", func() {
						if recordID.Text == "" {
							msg.SetText("Input must be filled!!!")
						} else {
							db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+":"+strconv.Itoa(port)+")/"+DataBaseName)
							if err != nil {
								panic(err)
							}
							res, err = db.Query("select * from patient where PatientID = " + recordID.Text)
							f, err := os.Create("C:/Users/Mario Ehab/Desktop/data.txt")
							var all string
							for res.Next() {
								var patientrecord patientstruct
								err = res.Scan(&patientrecord.PatientID, &patientrecord.PatientName, &patientrecord.Age, &patientrecord.Gender, &patientrecord.Address, &patientrecord.Disease, &patientrecord.DoctorID)

								all = "PatientID : " + strconv.Itoa(patientrecord.PatientID) + "\n Patient Name: " + patientrecord.PatientName + "\n Age : " + patientrecord.Age + "\n Gender: " + patientrecord.Gender + "\n Address: " + patientrecord.Address + "\n Doctor ID: " + strconv.Itoa(patientrecord.DoctorID)
								_, err = f.WriteString("\n" + all + "\n" + "\n")
							}

							if err != nil {
								panic(err)
							}
							data, _ := ioutil.ReadFile("C:/Users/Mario Ehab/Desktop/data.txt")

							result := fyne.NewStaticResource("Data in Tables", data)
							entry := widget.NewMultiLineEntry()
							entry.SetText(string(result.StaticContent))
							w := fyne.CurrentApp().NewWindow(
								string(result.StaticName))
							w.SetContent(container.NewScroll(entry))
							w.Resize(fyne.NewSize(400, 400))
							w.Show()

							fmt.Println("----> RECORD Searched Successfully!!...")
						}

					})
					searchbtn.Resize(fyne.NewSize(200, 40))
					searchbtn.Move(fyne.NewPos(300, 180))

					content = container.NewGridWithRows(3, recordID, searchbtn, msg)
				} else {
					msg.SetText("")
					type roomstruct struct {
						RoomNo     int
						RoomType   string
						RoomStatus string
						PatientID  int
					}
					searchbtn := widget.NewButton("Search", func() {
						if recordID.Text == "" {
							msg.SetText("Input must be filled!!!")
						} else {
							db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+":"+strconv.Itoa(port)+")/"+DataBaseName)
							if err != nil {
								panic(err)
							}
							res, err = db.Query("select * from room where RoomNo = " + recordID.Text)
							f, err := os.Create("C:/Users/Mario Ehab/Desktop/data.txt")
							var all string
							for res.Next() {
								var roomrecord roomstruct
								err = res.Scan(&roomrecord.RoomNo, &roomrecord.RoomType, &roomrecord.RoomStatus, &roomrecord.PatientID)

								all = "Room Number : " + strconv.Itoa(roomrecord.RoomNo) + "\n Room Type: " + roomrecord.RoomType + "\n Room Status : " + roomrecord.RoomStatus + "\n Patient ID : " + strconv.Itoa(roomrecord.PatientID)
								_, err = f.WriteString("\n" + all + "\n" + "\n")
							}

							if err != nil {
								panic(err)
							}
							data, _ := ioutil.ReadFile("C:/Users/Mario Ehab/Desktop/data.txt")

							result := fyne.NewStaticResource("Data in Tables", data)
							entry := widget.NewMultiLineEntry()
							entry.SetText(string(result.StaticContent))
							w := fyne.CurrentApp().NewWindow(
								string(result.StaticName))
							w.SetContent(container.NewScroll(entry))
							w.Resize(fyne.NewSize(400, 400))
							w.Show()

							fmt.Println("----> RECORD Searched Successfully!!...")
						}

					})
					searchbtn.Resize(fyne.NewSize(200, 40))
					searchbtn.Move(fyne.NewPos(300, 180))

					content = container.NewGridWithRows(3, recordID, searchbtn, msg)
				}
				w.SetContent(
					container.NewGridWithColumns(
						3,
						ListView,
						tablelist,
						content,
					),
				)
			}
		} else {
			w.Close()
		}
	}
	split := container.NewHSplit(
		ListView,
		container.NewMax(contentText),
	)
	split.Offset = 0.2

	w.SetContent(split)
	w.ShowAndRun()
	fmt.Println("App is Closed....")
}
