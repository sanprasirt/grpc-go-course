package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sanprasirt/grpc-go-course/blog/blogpb"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()
	conn, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	c := blogpb.NewBlogServiceClient(conn)

	fmt.Println("Create the blog")
	blog := &blogpb.Blog{
		AuthorId: "Sanrpasirt",
		Title:    "My second blog.",
		Content:  "This is my second blog content.",
	}
	crtblog, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error %v", err)
	}
	fmt.Printf("Blog has been created: %v\n", crtblog)
	blogID := "5fe937287cd1fc41d6d7a3d5"

	// Read Blog
	fmt.Println("Reading the blog")

	// _, err = c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "aaaaccccc"})
	// if err != nil {
	// 	fmt.Printf("Error while reading: %v\n", err)
	// }

	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogID}
	res, errRes := c.ReadBlog(context.Background(), readBlogReq)
	if errRes != nil {
		fmt.Printf("Error while reading: %v\n", errRes)
	}

	fmt.Printf("Blog was read : %v\n", res)
}
