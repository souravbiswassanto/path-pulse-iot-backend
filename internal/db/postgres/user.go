package postgres

import (
	"context"
	"fmt"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
	custom_error "github.com/souravbiswassanto/path-pulse-iot-backend/pkg/lib"
)

type UserSqlDB struct {
	*PostgresClient
}

func NewUserSqlDB(pc *PostgresClient) (*UserSqlDB, error) {
	if pc == nil || pc.xc == nil {
		return nil, fmt.Errorf("client can't be nil")
	}
	return &UserSqlDB{pc}, nil
}

func (db *UserSqlDB) SetupTable() error {
	// this should create the table
	return db.xc.Sync(new(models.User))
}

func (db *UserSqlDB) Get(ctx context.Context, v interface{}) (interface{}, error) {
	userid, ok := v.(models.UserID)
	if !ok {
		return nil, fmt.Errorf("expected %T object, but got %T object", uint64(0), v)
	}
	if userid == 0 {
		return nil, fmt.Errorf("userid can't be 0")
	}
	user := new(models.User)
	//selectQuery := fmt.Sprintf(`
	//    SELECT
	//        userid, name, age, gender, email, phone, address, created_at, updated_at
	//    FROM
	//        USER
	//    WHERE
	//        userid='%v'`, userid)
	//
	//results, err := db.xc.QueryString(selectQuery)
	//if err != nil {
	//	return nil, err
	//}
	//
	//if len(results) == 0 {
	//	return nil, custom_error.ErrUserNotFound
	//}
	//row := results[0]
	//
	//user := &models.User{
	//	Name:      row["name"],
	//	Gender:    row["gender"],
	//	Email:     row["email"],
	//	Phone:     row["phone"],
	//	Address:   row["address"],
	//	CreatedAt: pointer.StringP(row["created_at"]),
	//	UpdatedAt: pointer.StringP(row["updated_at"]),
	//}
	//
	//// Convert age to int32
	//age, err := strconv.ParseInt(row["age"], 10, 32)
	//if err != nil {
	//	log.Fatalf("Failed to parse age: %v", err)
	//}
	//user.Age = int32(age)
	db.xc.SetDefaultContext(ctx)

	exists, err := db.xc.ID(userid).Get(user)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, custom_error.ErrUserNotFound
	}
	return user, nil
}

func (db *UserSqlDB) Create(ctx context.Context, v interface{}) error {
	user, ok := v.(*models.User)
	if !ok {
		return fmt.Errorf("expected %T object, but got %T object", &models.User{}, v)
	}
	if user.ID == 0 {
		return fmt.Errorf("userid can't be 0")
	}
	_, err := db.Get(ctx, user.ID)
	if err == nil {
		return custom_error.ErrUserAlreadyExists
	}

	db.xc.SetDefaultContext(ctx)
	_, err = db.xc.Insert(user)
	//createQuery := fmt.Sprintf(`insert into USER(userid,name,age,gender,email,phone,address,created_at,updated_at) values(%v, '%v', %v, '%v', '%v', '%v', '%v','%v', '%v')`, user.ID, user.Name, user.Age, user.Gender, user.Email, user.Phone, user.Address, user.CreatedAt, user.UpdatedAt)
	//_, err = db.xc.QueryString(createQuery)
	return err
}

func (db *UserSqlDB) Update(ctx context.Context, v interface{}) error {
	user, ok := v.(*models.User)
	if !ok {
		return fmt.Errorf("expected %T object, but got %T object", &models.User{}, v)
	}
	if user.ID == 0 {
		return fmt.Errorf("userid can't be 0")
	}
	_, err := db.Get(ctx, user.ID)
	if err != nil && custom_error.IsUserNotFoundErr(err) {
		return db.Create(ctx, user)
	} else if err != nil {
		return fmt.Errorf("can't get the record with %v, err: %v", user.ID, err)
	}
	db.xc.SetDefaultContext(ctx)
	//updateQuery := fmt.Sprintf(`UPDATE USER Set name='%v',
	//            age=%v,
	//            gender='%v',email='%v',
	//            phone='%v',address='%v',
	//            created_at='%v',updated_at='%v' where userid=%v`, user.Name, user.Age, user.Gender, user.ContactInfo.Email, user.ContactInfo.Phone, user.ContactInfo.Address, user.CreatedAt, user.UpdatedAt, user.ID)
	_, err = db.xc.ID(user.ID).Update(user)

	return err
}

func (db *UserSqlDB) Delete(ctx context.Context, v interface{}) error {
	userid, ok := v.(models.UserID)
	if !ok {
		return fmt.Errorf("expected %T object, but got %T object", uint64(0), v)
	}
	if userid == 0 {
		return fmt.Errorf("userid can't be 0")
	}
	user := new(models.User)
	db.xc.SetDefaultContext(ctx)

	_, err := db.xc.ID(userid).Delete(user)
	return err
}
