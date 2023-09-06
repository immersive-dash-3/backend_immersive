package service

import (
	"errors"
	"immersive_project/klp3/features/menteelogs"
	"immersive_project/klp3/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	mockData := new(mocks.MenteeLogData)

	t.Run("success delete data",func(t *testing.T){
		mockData.On("Delete",uint(1)).Return(nil).Once()
		srv:=New(mockData)
		err:=srv.Delete(uint(1))
		assert.Nil(t,err)
		mockData.AssertExpectations(t)
	})

	t.Run("failed delete data",func(t *testing.T){
		mockData.On("Delete",uint(1)).Return(errors.New("error delete data")).Once()
		srv:=New(mockData)
		err:=srv.Delete(uint(1))
		assert.NotNil(t,err)
		mockData.AssertExpectations(t)		
	})
}

func TestEdit(t *testing.T){
	mockData := new(mocks.MenteeLogData)

	t.Run("success edit data",func(t *testing.T){
		inputData:=menteelogs.MenteeLogEntity{
			MenteeID:  uint(1),
			UserID:    uint(1),
			Status:    "join new",
			Log:       "keren banget",
		}
		mockData.On("Update",uint(1),inputData).Return(nil).Once()
		srv:=New(mockData)
		err:= srv.Edit(uint(1),inputData)
		assert.Nil(t,err)
		mockData.AssertExpectations(t)
	})

	t.Run("failed edit data",func(t *testing.T){
		inputData:=menteelogs.MenteeLogEntity{
			MenteeID:  uint(1),
			UserID:    uint(1),
			Status:    "join new",
			Log:       "keren banget",
		}
		mockData.On("Update",uint(1),inputData).Return(errors.New("error edit data")).Once()
		srv:=New(mockData)
		err:= srv.Edit(uint(1),inputData)
		assert.NotNil(t,err)
		mockData.AssertExpectations(t)		
	})
}

func TestGet(t *testing.T){
	mockData := new(mocks.MenteeLogData)
	returnData:=menteelogs.MenteeEntity{
		Id:              uint(1),
		FirstName:       "Santi",
		LastName:        "Maria",
		Gender:          "Perempuan",
		Email:           "santi@gmail.com",
		PhoneNumber:     "081738472873",
		Telegram:        "@jsanti",
		ClassID:         uint(1),
		StatusID:        uint(1),
		EducationType:   "Informatics",
	}

	t.Run("success get mentee log",func(t *testing.T){
		mockData.On("Select",uint(1)).Return(returnData,nil).Once()
		srv:=New(mockData)
		response,err:=srv.Get(uint(1))
		assert.Nil(t,err)
		assert.Equal(t,returnData,response)
		mockData.AssertExpectations(t)
	})

	t.Run("failed get menteelog",func(t *testing.T){
		mockData.On("Select",uint(1)).Return(menteelogs.MenteeEntity{},errors.New("failed get mentee log")).Once()
		srv:=New(mockData)
		response,err:=srv.Get(uint(1))
		assert.NotNil(t,err)
		assert.Equal(t,menteelogs.MenteeEntity{},response)
		mockData.AssertExpectations(t)		
	})
}

func TestAdd(t *testing.T){
	mockData := new(mocks.MenteeLogData)

	t.Run("success insert feedback",func(t *testing.T){
		inputData:=menteelogs.MenteeLogEntity{
			MenteeID:  uint(1),
			UserID:    uint(1),
			Status:    "join new",
			Log:       "keren banget",
		}
		mockData.On("Insert",inputData).Return(inputData.Status,nil).Once()	
		mockData.On("InsertStatus",inputData.Status).Return(uint(1),nil).Once()
		mockData.On("UpdateMentee",uint(1),uint(1)).Return(nil).Once()
		srv:=New(mockData)
		err:=srv.Add(inputData)
		assert.Nil(t,err)
		mockData.AssertExpectations(t)
	})

	t.Run("failed insert log",func(t *testing.T){
		inputData:=menteelogs.MenteeLogEntity{
			MenteeID:  uint(1),
			UserID:    uint(1),
			Status:    "join new",
			Log:       "keren banget",
		}
		mockData.On("Insert",inputData).Return("",errors.New("insert failed")).Once()	
		srv:=New(mockData)
		err:=srv.Add(inputData)
		assert.NotNil(t,err)
		mockData.AssertExpectations(t)		
	})
	t.Run("failed insert status",func(t *testing.T){
		inputData:=menteelogs.MenteeLogEntity{
			MenteeID:  uint(1),
			UserID:    uint(1),
			Status:    "join new",
			Log:       "keren banget",
		}
		mockData.On("Insert",inputData).Return(inputData.Status,nil).Once()	
		mockData.On("InsertStatus",inputData.Status).Return(uint(0),errors.New("failed insert status")).Once()	
		srv:=New(mockData)
		err:=srv.Add(inputData)
		assert.NotNil(t,err)
		mockData.AssertExpectations(t)		
	})
	t.Run("failed update status mentee",func(t *testing.T){
		inputData:=menteelogs.MenteeLogEntity{
			MenteeID:  uint(1),
			UserID:    uint(1),
			Status:    "join new",
			Log:       "keren banget",
		}
		mockData.On("Insert",inputData).Return(inputData.Status,nil).Once()	
		mockData.On("InsertStatus",inputData.Status).Return(uint(1),nil).Once()
		mockData.On("UpdateMentee",uint(1),uint(1)).Return(errors.New("failed status update")).Once()
		srv:=New(mockData)
		err:=srv.Add(inputData)
		assert.NotNil(t,err)
		mockData.AssertExpectations(t)		
	})
}