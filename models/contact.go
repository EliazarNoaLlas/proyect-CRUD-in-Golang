/*
* File              : contact.go
* Author            : Eliazar
* Creation date     : 17-03-2024
* Last modified by  : Eliazar
* Last modified date: 17-03-2024
* Description       : This file contains the models for contact
 */

package models

type Contact struct {
	// Description: identifier of the contact
	Id int `json:"id" binding:"required" example:"1"`
	// Description: name of the contact
	Name string `json:"name" binding:"required" example:"John Doe"`
	// Description: email address of the contact
	Email string `json:"email" binding:"required" example:"john@example.com"`
	// Description: phone number of the contact
	Phone string `json:"phone" binding:"required" example:"+123456789"`
}
