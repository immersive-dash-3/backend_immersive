package service

import (
	"errors"
	"immersive_project/klp3/features/classes"
	"immersive_project/klp3/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T){
	mockData := new(mocks.ClassData)
	
	t.Run("success delete class",func(t *testing.T){
		mockData.On("Delete",uint(1)).Return(nil).Once()
		srv:=New(mockData)
		err:=srv.Delete(uint(1))
		assert.Nil(t,err)
		mockData.AssertExpectations(t)	
	})

	t.Run("failed delete class",func(t *testing.T){
		mockData.On("Delete",uint(1)).Return(errors.New("delete failed")).Once()
		srv:=New(mockData)
		err:=srv.Delete(uint(1))
		assert.NotNil(t,err)
		mockData.AssertExpectations(t)		
	})
}

func TestGetById(t *testing.T){
	mockData := new(mocks.ClassData)
	returnData:=classes.ClassessEntity{
		Id: uint(1),Name:"BE 17",
	}

	t.Run("success get by id", func(t *testing.T){
		mockData.On("SelectById",uint(1)).Return(returnData,nil).Once()
		srv:=New(mockData)
		response,err:=srv.GetById(uint(1))
		assert.Nil(t,err)
		assert.Equal(t,returnData,response)
		mockData.AssertExpectations(t)
	})

	t.Run("failed get by id",func(t *testing.T){
		mockData.On("SelectById",uint(1)).Return(classes.ClassessEntity{},errors.New("error get data class")).Once()
		srv:=New(mockData)
		response,err:=srv.GetById(uint(1))
		assert.NotNil(t,err)
		assert.Equal(t,classes.ClassessEntity{},response)
		mockData.AssertExpectations(t)
	})
}

func TestEdit(t *testing.T){
	mockData := new(mocks.ClassData)
	
	t.Run("success update class",func(t *testing.T){
		inputData:=classes.ClassessEntity{
			Name:"BE 17",
		}
		mockData.On("Update",uint(1),inputData).Return(uint(1),nil).Once()
		srv:=New(mockData)
		err:=srv.Edit(uint(1),inputData)
		assert.Nil(t,err)
		mockData.AssertExpectations(t)	
	})

	t.Run("failed update class",func(t *testing.T){
		inputData:=classes.ClassessEntity{
			Name:"BE 17",
		}
		mockData.On("Update",uint(1),inputData).Return(uint(0), errors.New("update failed")).Once()
		srv:=New(mockData)
		err:=srv.Edit(uint(1),inputData)
		assert.NotNil(t, err)
		mockData.AssertExpectations(t)

	})
}

func TestGetAll(t *testing.T){
	mockData := new(mocks.ClassData)
	returnData:=[]classes.ClassessEntity{
		{Id: uint(1),Name:"BE 17",},
		{Id: uint(2),Name:"BE 18",},
	}

	t.Run("success get all",func(t *testing.T){
		mockData.On("SelectAll",int(1),int(1)).Return(returnData,nil).Once()
		srv:=New(mockData)
		response,err:=srv.GetAll(int(1),int(1))
		assert.Nil(t,err)
		assert.Equal(t,returnData,response)
		mockData.AssertExpectations(t)
	})

	t.Run("failed get all",func(t *testing.T){
		mockData.On("SelectAll",int(1),int(1)).Return(nil,errors.New("error get all")).Once()
		srv:=New(mockData)
		response,err:=srv.GetAll(int(1),int(1))
		assert.NotNil(t,err)
		assert.Nil(t,response)
		mockData.AssertExpectations(t)		
	})
}

func TestAdd(t *testing.T){
	mockData := new(mocks.ClassData)

	t.Run("success add class",func(t *testing.T){
		inputData:=classes.ClassessEntity{
			Name:"BE 17",
		}
		mockData.On("Insert",inputData).Return(uint(1),nil).Once()
		srv:=New(mockData)
		err:=srv.Add(inputData)
		assert.Nil(t,err)
		mockData.AssertExpectations(t)	

	})

	t.Run("failed add class",func(t *testing.T){
		inputData:=classes.ClassessEntity{
			Name:"BE 17",
		}
		mockData.On("Insert",inputData).Return(uint(0),errors.New("error insert class")).Once()
		srv:=New(mockData)
		err:=srv.Add(inputData)
		assert.NotNil(t,err)
		mockData.AssertExpectations(t)			
	})

	t.Run("validate add class",func(t *testing.T){
		inputData:=classes.ClassessEntity{}
		srv:=New(mockData)
		err:=srv.Add(inputData)
		assert.NotNil(t,err)
		mockData.AssertExpectations(t)		
	})
}