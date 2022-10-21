MATCH
  (story:Story {id: $storyID})
DETACH DELETE story
