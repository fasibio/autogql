// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

func (d *TodoType) MergeToType() TodoType {
	return *d
}

func (d *CatPatch) MergeToType() map[string]interface{} {
	res := make(map[string]interface{})
	if d.Name != nil {
		res["name"] = *d.Name
	}
	if d.BirthDay != nil {
		res["birth_day"] = *d.BirthDay
	}
	if d.UserID != nil {
		res["user_id"] = *d.UserID
	}
	if d.Alive != nil {
		res["alive"] = d.Alive
	}
	return res
}

func (d *CatInput) MergeToType() Cat {

	tmpName := d.Name

	tmpBirthDay := d.BirthDay

	tmpUserID := d.UserID

	var tmpAlive *bool
	if d.Alive != nil {
		tmpAlive = d.Alive
	}
	return Cat{
		Name:     tmpName,
		BirthDay: tmpBirthDay,
		UserID:   tmpUserID,
		Alive:    tmpAlive,
	}
}

func (d *CompanyPatch) MergeToType() map[string]interface{} {
	res := make(map[string]interface{})
	if d.Name != nil {
		res["name"] = *d.Name
	}
	if d.Description != nil {
		res["description"] = d.Description
	}
	if d.MotherCompanyID != nil {
		res["mother_company_id"] = d.MotherCompanyID
	}
	if d.MotherCompany != nil {
		res["mother_company"] = d.MotherCompany.MergeToType()
	}
	return res
}

func (d *CompanyInput) MergeToType() Company {

	tmpName := d.Name

	var tmpDescription *string
	if d.Description != nil {
		tmpDescription = d.Description
	}

	var tmpMotherCompanyID *int
	if d.MotherCompanyID != nil {
		tmpMotherCompanyID = d.MotherCompanyID
	}

	var tmpMotherCompany Company
	if d.MotherCompany != nil {
		tmpMotherCompany = d.MotherCompany.MergeToType()
	}
	return Company{
		Name:            tmpName,
		Description:     tmpDescription,
		MotherCompanyID: tmpMotherCompanyID,
		MotherCompany:   &tmpMotherCompany,
	}
}

func (d *SmartPhonePatch) MergeToType() map[string]interface{} {
	res := make(map[string]interface{})
	if d.Brand != nil {
		res["brand"] = *d.Brand
	}
	if d.Phonenumber != nil {
		res["phonenumber"] = *d.Phonenumber
	}
	if d.UserID != nil {
		res["user_id"] = *d.UserID
	}
	return res
}

func (d *SmartPhoneInput) MergeToType() SmartPhone {

	tmpBrand := d.Brand

	tmpPhonenumber := d.Phonenumber

	tmpUserID := d.UserID
	return SmartPhone{
		Brand:       tmpBrand,
		Phonenumber: tmpPhonenumber,
		UserID:      tmpUserID,
	}
}

func (d *TodoPatch) MergeToType() map[string]interface{} {
	res := make(map[string]interface{})
	if d.Name != nil {
		res["name"] = *d.Name
	}
	if d.Etype1 != nil {
		res["etype1"] = d.Etype1.MergeToType()
	}
	if d.Etype5 != nil {
		res["etype5"] = d.Etype5.MergeToType()
	}
	if d.Test123 != nil {
		res["test123"] = *d.Test123
	}
	return res
}

func (d *TodoInput) MergeToType() Todo {

	tmpName := d.Name

	var tmpEtype1 TodoType
	if d.Etype1 != nil {
		tmpEtype1 = d.Etype1.MergeToType()
	}

	var tmpEtype5 TodoType
	tmpEtype5 = d.Etype5.MergeToType()

	var tmpTest123 string
	if d.Test123 != nil {
		tmpTest123 = *d.Test123
	}
	return Todo{
		Name:    tmpName,
		Etype1:  &tmpEtype1,
		Etype5:  tmpEtype5,
		Test123: &tmpTest123,
	}
}

func (d *UserPatch) MergeToType() map[string]interface{} {
	res := make(map[string]interface{})
	if d.Name != nil {
		res["name"] = *d.Name
	}
	if d.Cat != nil {
		res["cat"] = d.Cat.MergeToType()
	}
	if d.CompanyID != nil {
		res["company_id"] = d.CompanyID
	}
	if d.Company != nil {
		res["company"] = d.Company.MergeToType()
	}
	if d.SmartPhones != nil {
		tmpSmartPhones := make([]map[string]interface{}, len(d.SmartPhones))
		for _, v := range d.SmartPhones {
			tmp := v.MergeToType()
			tmpSmartPhones = append(tmpSmartPhones, tmp)
		}
		res["smart_phones"] = tmpSmartPhones
	}
	return res
}

func (d *UserInput) MergeToType() User {

	tmpName := d.Name

	var tmpCat Cat
	if d.Cat != nil {
		tmpCat = d.Cat.MergeToType()
	}

	var tmpCompanyID *int
	if d.CompanyID != nil {
		tmpCompanyID = d.CompanyID
	}

	var tmpCompany Company
	if d.Company != nil {
		tmpCompany = d.Company.MergeToType()
	}

	var tmpSmartPhones []*SmartPhone
	if d.SmartPhones != nil {
		tmpSmartPhones = make([]*SmartPhone, len(d.SmartPhones))
		for _, v := range d.SmartPhones {
			tmp := v.MergeToType()
			tmpSmartPhones = append(tmpSmartPhones, &tmp)
		}
	}
	return User{
		Name:        tmpName,
		Cat:         &tmpCat,
		CompanyID:   tmpCompanyID,
		Company:     &tmpCompany,
		SmartPhones: tmpSmartPhones,
	}
}
