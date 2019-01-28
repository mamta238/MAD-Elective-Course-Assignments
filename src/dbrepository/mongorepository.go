package dbrepository

import (
	"fmt"
	"domain"
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
)

//MongoRepository mongodb repo
type MongoRepository struct {
	mongoSession *mgo.Session
	db           string
}

var collectionName = "restaurant"

//NewMongoRepository create new repository
func NewMongoRepository(mongoSession *mgo.Session, db string) *MongoRepository {
	return &MongoRepository{
		mongoSession: mongoSession,
		db:           db,
	}
}

//Reader Method:Find a Restaurant
func (r *MongoRepository) Get(id domain.ID) (*domain.Restaurant, error) {

	result := domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"_id": bson.ObjectId(id)}).One(&result)
	switch err {
	case nil:
		return &result, nil
	case mgo.ErrNotFound:
		return nil, domain.ErrNotFound
	default:
		return nil, err
	}
}

//Reader Method:Find By Restaurant Name
func (r *MongoRepository) FindByName(name string) ([] *domain.Restaurant, error) {
	
	result := [] *domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"name": bson.RegEx{name , ""} }).All(&result)
	switch err {
	case nil:
		return result, nil
	case mgo.ErrNotFound:
		return nil, domain.ErrNotFound
	default:
		return nil , err
	}
}

//Reader Method:Get list of all Restaurants

func (r *MongoRepository) GetAll() ([] *domain.Restaurant, error){

	result := [] *domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	
	err := coll.Find(nil).All(&result)
	
	switch err {
	case nil:
		return result, nil
	case mgo.ErrNotFound:
		return nil, domain.ErrNotFound
	default:
		return nil , err
	}
}


//Writer method:Store a Restaurantrecord
func (r *MongoRepository) Store(b *domain.Restaurant) (domain.ID, error) {

	
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	
	if  "" == b.DBID {
		b.DBID = domain.NewID()
		fmt.Println(b)
	}

	_, err := coll.UpsertId(b.DBID, b)

	if err != nil {
		return domain.ID(0), err
	}
	return b.DBID, nil
}

//Writer method:Delete a record on basis of Id

func (r *MongoRepository) Delete(id domain.ID) error{
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Remove(bson.M{"_id": bson.ObjectId(id)})
	return err		
}

//Filter Method:Filter restaurants with given food type 
func (r *MongoRepository) FindByTypeOfFood(foodType string) ([] *domain.Restaurant,error){
	
	result := [] *domain.Restaurant{}
	
	session := r.mongoSession.Clone()
	coll := session.DB(r.db).C(collectionName)
	
	err := coll.Find(bson.M{"typeOfFood":foodType}).All(&result)
	
	switch err{
	
		case nil : 
			return result,err
		case mgo.ErrNotFound :
			return nil,domain.ErrNotFound
		default :
			return nil,err 	
	}
	
}

//Filter Method:Filter restaurants with given post code
func (r *MongoRepository) FindByTypeOfPostCode(postcode string) ([] *domain.Restaurant,error){
	
	result := [] *domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	
	err := coll.Find(bson.M{"postcode":postcode}).All(&result)
	
	switch err{
	
		case nil : 
			return result,err
		case mgo.ErrNotFound :
			return nil,domain.ErrNotFound
		default :
			return nil,err		
	}	
}

//Filter Method: Search on all string type(Using text search by creating index) 

func (r *MongoRepository) Search(query string) ([] *domain.Restaurant,error){
	
	
	result := []*domain.Restaurant{} 
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	
	index := mgo.Index{ Key : []string					{"$text:name","$text:address","$text:addressLine2","$text:url","$text:outcode","$text:postcode","$text:typeOfFood"},
	}	
	coll.EnsureIndex(index)
	
	err := coll.Find(bson.M{"$text":bson.M{"$search":query}}).All(&result)
	
	return result,err
	
}
