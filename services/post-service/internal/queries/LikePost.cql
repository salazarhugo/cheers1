MATCH
  (u:User {id: $userID})
MATCH
  (creator:User)-[:POSTED]->(post:Post {id: $postID})
MERGE
  (u)-[:LIKED]->(post)
RETURN
  {
    photoUrl: u.picture,
    username: u.username,
    name:     u.name,
    tokens:   creator.registrationTokens,
    authorId: creator.id,
    drink: post.drink
  }
