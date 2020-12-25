package main
import "fmt"

func getLikesAndShares(postId int) (int,int) {//here we have 1 parameter and 2 return types in (..,..) format.
  var likesForPost, sharesForPost int
  switch postId {
  case 1:
    likesForPost = 5
    sharesForPost = 7
  case 2:
    likesForPost = 3
    sharesForPost = 11
  case 3:
    likesForPost = 22
    sharesForPost = 1
  case 4:
    likesForPost = 7
    sharesForPost = 9
  }
  fmt.Println("Likes: ", likesForPost, "Shares: ", sharesForPost)
  return likesForPost, sharesForPost//now, here we define which var's to use to return the required 2 return values.
}

func main() {
  var likes, shares int
  
  likes, shares = getLikesAndShares(4)//here as we have two return values, the getLikesAndShares() will return those two consecatively
 //so we store them in variables we defined in main() itself so as to use those reutrn values in the main(). 
  
 if likes > 5 {
    fmt.Println("Woohoo! We got some likes.")
  }
  if shares > 10 {
    fmt.Println("We went viral!")
  }
}