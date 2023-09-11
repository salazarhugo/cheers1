MATCH (party:Party {slug: $slug})
RETURN PROPERTIES(party)
