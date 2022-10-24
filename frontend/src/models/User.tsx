import { UserTypeInterface } from "./UserType"
export interface UserInterface {
    ID: number,
	Name:     string
	User_PASSWORD: string
	ISNAME:        string
	UserType_ID: Number
	UserType:    UserTypeInterface
  }