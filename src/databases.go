package main

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	_ "github.com/go-sql-driver/mysql"
)

var store = map[string]*Recipe{}

type RecipeResponce struct {
	Message string `json:"message"`
	Recipes [1]Recipe `json:"recipe"`
}

type RecipesResponce struct {
	Recipes []Recipe `json:"recipes"`
}

type ErrorResponceWithRequrired struct {
	Message string `json:"message"`
	Required string `json:"required"`
}

type ErrorResponce struct {
	Message string `json:"message"`
}

func (i *Impl) GetAllRecipes(w rest.ResponseWriter, r *rest.Request){
	recipes := []Recipe{}
	i.DB.Find(&recipes)
	var res RecipesResponce
	res.Recipes = recipes
	w.WriteJson(&res)
}

func (i *Impl) PostRecipe(w rest.ResponseWriter, r *rest.Request){
	recipe := Recipe{}
	err := r.DecodeJsonPayload(&recipe)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if recipe.Title == "" || recipe.Making_time == "" || recipe.Serves == "" || recipe.Ingredients == "" || recipe.Cost == 0 {
		var res ErrorResponceWithRequrired
		res.Message = "Recipe creation failed!"
		res.Required = "title, making_time, serves, ingredients, cost"
		w.WriteHeader(500)
		w.WriteJson(&res)
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
	id := r.PathParam("id")
	recipe := Recipe{}
	if i.DB.First(&recipe, id).Error != nil{
		rest.NotFound(w, r)
		return
	}
	var res RecipeResponce
	var recipes [1]Recipe = [1]Recipe{recipe}
	res.Message = "Recipe details by id"
	res.Recipes = recipes
	w.WriteJson(&res)
}
func (i *Impl) PatchRecipe(w rest.ResponseWriter, r *rest.Request){
	id := r.PathParam("id")
	recipe := Recipe{}
	if i.DB.First(&recipe, id).Error != nil{
		rest.NotFound(w, r)
		return
	}

	updated := Recipe{}
	if err := r.DecodeJsonPayload(&updated); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	if err := i.DB.Model(&recipe).Updates(updated).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var res RecipeResponce
	var recipes [1]Recipe = [1]Recipe{updated}
	res.Message = "Recipe successfully updated!"
	res.Recipes = recipes
	w.WriteJson(&res)
}
func (i *Impl) DeleteRecipe(w rest.ResponseWriter, r *rest.Request){
	id := r.PathParam("id")
	recipe := Recipe{}
	if i.DB.First(&recipe, id).Error != nil {
		var res ErrorResponce
		res.Message = "No Recipe found"
		w.WriteHeader(500)
		w.WriteJson(&res)
		return
	}

	if err := i.DB.Delete(&recipe).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteJson(map[string]string{"message":"Recipe successfully removed!"})
}