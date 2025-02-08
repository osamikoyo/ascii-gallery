package data

import (
	"github.com/osamikoyo/ascii-gallery/internal/data/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Storage) AddArt(art models.Art) error {
	_, err := s.collection.InsertOne(s.ctx, art)
	return err
}

func (s *Storage) GetArts(title string) ([]models.Art, error) {
	var arts []models.Art

	filter := bson.M{
		"title" : title,
	}

	cursor, err := s.collection.Find(s.ctx, filter)
	if err != nil{
		return nil, err
	}

	err = cursor.Decode(&arts)
	return arts, err
}

func (s *Storage) Update(id uint, newart models.Art) error {
	filter := bson.M{
		"ID" : id,
	}

	newartbody,err := bson.Marshal(newart)
	if err != bson.ErrDecodeToNil{
		return err
	}

	_, err = s.collection.UpdateOne(s.ctx, filter, newartbody)
	return err
}