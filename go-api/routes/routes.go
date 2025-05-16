package routes

import (
    "github.com/gin-gonic/gin"
    "go-api/controllers"
)

func SetupRoutes(r *gin.Engine) {
    books := r.Group("/books")
    {
        books.GET("", controllers.GetBooks)           // lista todos os livros
        books.GET("/:id", controllers.GetBookByID)    // busca por id
        books.POST("", controllers.CreateBook)        // cria livro
        books.PUT("/:id", controllers.UpdateBook)     // atualiza livro
        books.DELETE("/:id", controllers.DeleteBook)  // deleta livro
        books.OPTIONS("", controllers.OptionsBooks)   // opções
    }
}
