package postgres

import (
	"context"
	"fmt"
	custom_error "github.com/souravbiswassanto/path-pulse-iot-backend/internal/custom-error"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
	"gomodules.xyz/pointer"
	"log"
	"strconv"
)

type UserDB struct {
	*PostgresClient
}

func (db *UserDB) SetupDatabase() {
	// this should create the table
	db.xc.Sync(new(models.User))
	_, _ = db.xc.IsTableExist(models.User{})
	_ = db.xc.CreateTables(models.User{})
	return
}

func (db *UserDB) Get(ctx context.Context, v interface{}) (interface{}, error) {
	userid, ok := v.(models.UserID)
	if !ok {
		return nil, fmt.Errorf("expected %T object, but got %T object", uint64(0), v)
	}
	selectQuery := fmt.Sprintf(`
        SELECT 
            userid, name, age, gender, email, phone, address, created_at, updated_at 
        FROM 
            USER 
        WHERE 
            userid='%v'`, userid)

	results, err := db.xc.QueryString(selectQuery)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, custom_error.ErrUserNotFound
	}
	row := results[0]

	user := &models.User{
		Name:      row["name"],
		Gender:    row["gender"],
		Email:     row["email"],
		Phone:     row["phone"],
		Address:   row["address"],
		CreatedAt: pointer.StringP(row["created_at"]),
		UpdatedAt: pointer.StringP(row["updated_at"]),
	}

	// Convert age to int32
	age, err := strconv.ParseInt(row["age"], 10, 32)
	if err != nil {
		log.Fatalf("Failed to parse age: %v", err)
	}
	user.Age = int32(age)
	return user, nil
}

func (db *UserDB) Create(ctx context.Context, v interface{}) error {
	user, ok := v.(*models.User)
	if !ok {
		return fmt.Errorf("expected %T object, but got %T object", &models.User{}, v)
	}
	_, err := db.Get(ctx, user.ID)
	if err == nil {
		return custom_error.ErrUserAlreadyExists
	}

	db.xc.SetDefaultContext(ctx)
	createQuery := fmt.Sprintf(`insert into USER(userid,name,age,gender,email,phone,address,created_at,updated_at) values(%v, '%v', %v, '%v', '%v', '%v', '%v','%v', '%v')`, user.ID, user.Name, user.Age, user.Gender, user.Email, user.Phone, user.Address, user.CreatedAt, user.UpdatedAt)
	_, err = db.xc.QueryString(createQuery)
	return err
}

func (db *UserDB) Update(ctx context.Context, v interface{}) error {
	user, ok := v.(*models.User)
	if !ok {
		return fmt.Errorf("expected %T object, but got %T object", &models.User{}, v)
	}
	_, err := db.Get(ctx, user.ID)
	if err != nil && custom_error.IsUserNotFoundErr(err) {
		return db.Create(ctx, user)
	} else if err != nil {
		return fmt.Errorf("can't get the record with %v, err: %v", user.ID, err)
	}
	db.xc.SetDefaultContext(ctx)
	updateQuery := fmt.Sprintf(`UPDATE USER Set name='%v',
                age=%v,
                gender='%v',email='%v',
                phone='%v',address='%v',
                created_at='%v',updated_at='%v' where userid=%v`, user.Name, user.Age, user.Gender, user.ContactInfo.Email, user.ContactInfo.Phone, user.ContactInfo.Address, user.CreatedAt, user.UpdatedAt, user.ID)
	_, err = db.xc.ID(user.ID).Update(user)

	return nil
}
