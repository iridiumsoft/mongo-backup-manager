package controllers

import (
	"time"
	"os"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/iridiumsoft/mongo-backup-manager/util"
	"github.com/iridiumsoft/mongo-backup-manager/models"

	"github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"log"
)

const (
	TempPath       = "./tmp/"
	S3DBBackUpPath = "db-backups/"
)

func (c *Controllers) ManualBackup(context *gin.Context) {
	c.DBDump(false)
	context.String(http.StatusOK, "Backup Job Started")
}

func (c *Controllers) DBDump(UpdatesOnly bool) {

	var files []string
	var where = bson.M{}

	DateTime := time.Now().Format("2006-01-02-150405")
	logrus.Println("Backup Bot started", DateTime)

	// get list of collections
	collections, err := c.App.DB.CollectionNames()
	if err != nil {
		logrus.Error("Error while getting collections", err.Error())
	}

	// if we need to get latest updates only
	if UpdatesOnly {
		where = c.GetUpdatesWhereClause()
		logrus.Print("Where values")
	}

	for _, collection := range collections {

		if util.InArray(collection, c.Config.ExcludeCollectionsInBackup) {
			continue
		}

		var docs []bson.M
		// find all Documents of a collection
		c.App.DB.C(collection).Find(where).All(&docs)

		// check if values are more than 0, then create a json file
		if len(docs) == 0 {
			continue
		}

		filePath := TempPath + collection + ".json"

		// Check if ObjectId Exists
		docs = util.Map(docs, func(ms bson.M) bson.M {
			for k, v := range ms {
				if c.App.DbService.IsObjectId(v) {
					v = bson.M{"$id": v}
				}
				ms[k] = v
			}
			return ms
		})

		if util.StoreJsonFile(filePath, docs) {
			files = append(files, filePath)
		}
	}

	if len(files) > 0 {
		util.ZipAndRemoveFiles(files, DateTime)
		c.UploadZipToS3(TempPath, DateTime)
		// Add Entry to database
		c.App.DB.C("backups").Insert(&models.Backup{
			Name: DateTime,
			Date: time.Now(),
		})
		logrus.Print("Backup Completed")
	} else {
		logrus.Println("No files Found")
	}

	// Update DB with last_update date
	c.App.DbService.UpdateSetting("last_db_update", bson.M{"date": time.Now()})

}

func (c *Controllers) GetUpdatesWhereClause() bson.M {
	var Where = bson.M{}

	setting := c.App.DbService.GetSetting("last_db_update")

	if setting["date"] != nil {
		lastUpdate, err := time.Parse(setting["date"].(string), time.RFC3339)
		if err == nil && !lastUpdate.IsZero() {
			Where["updated_on"] = lastUpdate
		}
	}

	return Where
}

func (c *Controllers) UploadZipToS3(path string, name string) {
	FilePath := path + name + ".zip"
	c.App.S3Service.PutObject(FilePath, S3DBBackUpPath+name+".zip", "private")
	os.Remove(FilePath)
}

func (c *Controllers) ExportCollection(context *gin.Context) {
	var docs []bson.M
	collection := context.Param("collection")
	// find all Documents of a collection
	c.App.DB.C(collection).Find(nil).All(&docs)
	filePath := TempPath + collection + ".json"
	util.StoreJsonFile(filePath, docs)
	util.DownloadAndRemoveFile(collection+".json", filePath, context)
}

func (c *Controllers) ImportCollection(context *gin.Context) {

}

func (c *Controllers) ImportView(context *gin.Context) {
	context.HTML(http.StatusOK, "import-view.html", gin.H{})
}

func (c *Controllers) RestoreDB(context *gin.Context) {

	id := bson.ObjectId(context.Param("id"))

	filterCollections := []string{"users"}

	var Backup models.Backup

	c.App.DB.C("backups").Find(bson.M{"_id": id}).One(&Backup)
	if Backup.Id == "" {
		context.JSON(http.StatusNotFound, bson.M{"error": "Invalid Id", "Backup": Backup, "Id": id,})
		return
	}

	// Get Zip folder from S3 Bucket
	filePath := TempPath + Backup.Name + ".zip"

	err := c.App.S3Service.GetObject(S3DBBackUpPath+Backup.Name+".zip", filePath)

	if err != nil {
		context.JSON(http.StatusNotFound, bson.M{"error": err.Error(), "step": "GetObject",})
		return
	}

	// Create Temporary Folder
	TempFilePath := TempPath + Backup.Name
	os.Mkdir(TempFilePath, 0755)

	// Unzip to Temporary Folder
	util.Unzip(filePath, TempFilePath)

	// Read Folder
	files, err := ioutil.ReadDir(TempFilePath)
	if err != nil {
		context.JSON(http.StatusNotFound, bson.M{"error": err.Error(), "step": "ReadDir",})
		return
	}

	for _, file := range files {
		name := file.Name()
		collectionName := strings.Replace(name, ".json", "", -1)
		if !util.InArray(collectionName, filterCollections) {
			c.UpdateFileInDb(TempFilePath+"/"+name, collectionName)
		}
	}

	context.JSON(http.StatusOK, bson.M{"Status": "Completed"})

}

func (c *Controllers) UpdateFileInDb(path string, collectionName string) {

	f, _ := ioutil.ReadFile(path)

	var docs []bson.M

	err := json.Unmarshal(f, &docs)

	if err != nil {
		log.Print("Error while Unmarshal doc: ", err.Error())
		return
	}

	for _, item := range docs {

		// get ID
		ID := item["_id"].(bson.M)["$id"]
		// Delete _id from item
		delete(item, "_id")
		for k, i := range item {
			c := i.(bson.M)
			if c["$id"] != nil {
				// check if $id exists so that we can convert it to object
				item[k] = c["$id"]
			}
		}

		var doc bson.M
		c.App.DB.C(collectionName).Find(bson.M{"_id": ID}).One(&doc)
		if doc["_id"] == "" {
			// Insert new Record
			c.App.DB.C(collectionName).Insert(item)
			continue
		}

		// update exiting
		c.App.DB.C(collectionName).Update(bson.M{"_id": ID}, item)

	}

}
