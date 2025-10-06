package common

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/book-service/api/app/utils/common/constants"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

func OTPGenerator(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	otp := make([]byte, length)
	for i := 0; i < length; i++ {
		otp[i] = byte('0' + rand.Intn(10))
	}
	return string(otp)
}

func CreateJSONResponse(c *gin.Context, status constants.ResponseType, statusCode int, message string, data interface{}) {
	var jsonResponse gin.H

	if status == constants.ResponseOK {
		jsonResponse = gin.H{
			"status": status,
			"data":   data,
		}
	} else {
		jsonResponse = gin.H{
			"status":  status,
			"message": message,
		}
	}
	c.JSON(statusCode, jsonResponse)
}

// CustomDate wraps time.Time to allow custom serialization
type CustomDate struct {
	time.Time
}

// MarshalJSON converts CustomDate to a JSON string in "YYYY-MM-DD" format
func (d CustomDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format("2006-01-02"))
}

// UnmarshalJSON parses CustomDate from a JSON string
func (d *CustomDate) UnmarshalJSON(data []byte) error {
	var dateStr string
	if err := json.Unmarshal(data, &dateStr); err != nil {
		return err
	}
	parsedTime, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return err
	}
	d.Time = parsedTime
	return nil
}

// MarshalBSONValue converts CustomDate to BSON DateTime
func (d CustomDate) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(d.Time)
}

// UnmarshalBSONValue parses CustomDate from BSON DateTime
func (d *CustomDate) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	parsedTime := time.Unix(0, int64(bson.RawValue{Type: t, Value: data}.DateTime())*int64(time.Millisecond))
	d.Time = parsedTime
	return nil
}
