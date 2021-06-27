package main

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	_ "github.com/go-sql-driver/mysql"
)

var store = map[string]*Recipe{}

func (i *Impl) GetAllRecipes(w rest.ResponseWriter, r *rest.Request){
	recipes := []Recipe{}
	i.DB.Find(&recipes)
	w.WriteJson(&recipes)
}

type RecipeResponce struct {
	Message string `json:"message"`
	Recipes [1]Recipe `json:"recipe"`
}

type RecipesResponce struct {
	Recipes []Recipe `json:"recipes"`
}

func (i *Impl) PostRecipe(w rest.ResponseWriter, r *rest.Request){
	recipe := Recipe{}
	err := r.DecodeJsonPayload(&recipe)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if recipe.Title == "" || recipe.Making_time == "" || recipe.Serves == "" || recipe.Ingredients == "" || recipe.Cost == 0 {
		w.WriteJson(map[string]string{"message": "Recipe creation failed!","required": "title, making_time, serves, ingredients, cost"})
		return
	}

	if err := i.DB.Save(&recipe).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var res RecipeResponce
	var recipes [1]Recipe = [1]Recipe{recipe}
	res.Message = "Recipe successfully created!"
	res.Recipes = recipes
	w.WriteJson(&res)
}

func (i *Impl) GetRecipe(w rest.ResponseWriter, r *rest.Request){
	recipe := []Recipe{}
	i.DB.Find(&recipe)
	w.WriteJson(&recipe)
}
func (i *Impl) PatchRecipe(w rest.ResponseWriter, r *rest.Request){
	recipe := []Recipe{}
	i.DB.Find(&recipe)
	w.WriteJson(&recipe)
}
func (i *Impl) DeleteRecipe(w rest.ResponseWriter, r *rest.Request){
	recipe := []Recipe{}
	i.DB.Find(&recipe)
	w.WriteJson(&recipe)
}
