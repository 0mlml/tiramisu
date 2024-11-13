package backend

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/boltdb/bolt"
	"github.com/google/uuid"
)

var (
	usersBucket       = []byte("users")
	questionsBucket   = []byte("questions")
	submissionsBucket = []byte("submissions")
)

type DB struct {
	*bolt.DB
}

func newDB(path string) (*DB, error) {
	db, err := bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		for _, bucket := range [][]byte{usersBucket, questionsBucket, submissionsBucket} {
			_, err := tx.CreateBucketIfNotExists(bucket)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return &DB{db}, err
}

// User methods
func (db *DB) CreateUser(user *User) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(usersBucket)

		// Check if email already exists
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var existingUser User
			if err := json.Unmarshal(v, &existingUser); err != nil {
				return err
			}
			if existingUser.Email == user.Email {
				return errors.New("email already exists")
			}
		}

		// Generate UUID if not provided
		if user.ID == "" {
			user.ID = uuid.New().String()
		}

		user.Created = time.Now()

		buf, err := json.Marshal(user)
		if err != nil {
			return err
		}

		return b.Put([]byte(user.ID), buf)
	})
}

func (db *DB) GetUser(id string) (*User, error) {
	var user User
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(usersBucket)
		v := b.Get([]byte(id))
		if v == nil {
			return errors.New("user not found")
		}
		return json.Unmarshal(v, &user)
	})
	return &user, err
}

func (db *DB) GetUserByEmail(email string) (*User, error) {
	var user User
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(usersBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var u User
			if err := json.Unmarshal(v, &u); err != nil {
				return err
			}
			if u.Email == email {
				user = u
				return nil
			}
		}
		return errors.New("user not found")
	})
	return &user, err
}

// Question methods
func (db *DB) CreateQuestion(question *Question) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(questionsBucket)

		if question.ID == "" {
			question.ID = uuid.New().String()
		}

		buf, err := json.Marshal(question)
		if err != nil {
			return err
		}

		return b.Put([]byte(question.ID), buf)
	})
}

func (db *DB) GetQuestions() ([]Question, error) {
	var questions []Question
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(questionsBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var question Question
			if err := json.Unmarshal(v, &question); err != nil {
				return err
			}
			questions = append(questions, question)
		}
		return nil
	})
	return questions, err
}

// Submission methods
func (db *DB) CreateSubmission(submission *Submission) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(submissionsBucket)

		if submission.ID == "" {
			submission.ID = uuid.New().String()
		}

		submission.CreatedAt = time.Now()

		buf, err := json.Marshal(submission)
		if err != nil {
			return err
		}

		return b.Put([]byte(submission.ID), buf)
	})
}

func (db *DB) GetUserSubmissions(userID string) ([]Submission, error) {
	var submissions []Submission
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(submissionsBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var submission Submission
			if err := json.Unmarshal(v, &submission); err != nil {
				return err
			}
			if submission.UserID == userID {
				submissions = append(submissions, submission)
			}
		}
		return nil
	})
	return submissions, err
}

func (db *DB) GetAllSubmissions() ([]Submission, error) {
	var submissions []Submission
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(submissionsBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var submission Submission
			if err := json.Unmarshal(v, &submission); err != nil {
				return err
			}
			submissions = append(submissions, submission)
		}
		return nil
	})
	return submissions, err
}
