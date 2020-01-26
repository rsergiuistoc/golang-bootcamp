package main


import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"github.com/rsergiuistoc/golang-bootcamp/models"
)

var contacts = models.Contacts

func main(){

	router := gin.Default()

	v1 := router.Group("/api/v1/contacts")
	{
		v1.GET("/", getAllContacts)
		v1.GET("/:id", getContact)
		v1.POST("/", createContact)
		v1.PUT("/:id", updateContact)
		v1.DELETE("/:id", deleteContact)
	}
	router.Run()
}

func getAllContacts(c *gin.Context){

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": contacts})
}

func getContact(c *gin.Context){

	contactId := c.Param("id")

	if contactId == ""{
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Contact not found!"})
		return
	}

	for _, con := range contacts{
		if con.Id == contactId{
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": con})
		}
	}
}

func createContact(c *gin.Context){

	var contact models.Contact

	err := c.BindJSON(&contact)
	if err != nil{
		log.Fatalln(err)
		c.JSON(http.StatusBadRequest, gin.H{"statuts": http.StatusBadRequest,
				"message": "Invalid json body"})
		return
	}

	contacts = append(contacts, contact)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK,
					"message": "Contant succesfully created", "data": contact})

}

func updateContact(c *gin.Context){

	contactId := c.Param("id")

	var contact models.Contact

	err := c.BindJSON(&contact)
	if err != nil{
		log.Fatalln(err)
		c.JSON(http.StatusBadRequest, gin.H{"statuts": http.StatusBadRequest,
			"message": "Invalid json body"})
		return
	}

	if contactId == ""{
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Contact found!"})
		return
	}

	for _ , con := range contacts{
		if con.Id == contactId{
			con.Id = contact.Id
			con.Name = contact.Name
			con.PhoneNumber = contact.PhoneNumber
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK,
			"message": "Contant succesfully updated"})

}

func deleteContact(c *gin.Context){

	contactId := c.Param("id")

	if contactId == ""{
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Contact found!"})
		return
	}

	for i, con := range contacts{
		if con.Id == contactId{
			copy(contacts[i:], contacts[i+1:])
			contacts[len(contacts) - 1] = models.Contact{}
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Contant succesfully deleted"})
}