package handler

import (
)

func extract_movies_list( body []interface{} ) []uint32 {

  // Get all MovieIds
  movieIds := make( []uint32, 0 )

  for _, t := range body {
    m := t.(map[string]interface{})
    movieIds = append( movieIds, (uint32)(m["MovieId"].(float64)) )
  }

  return movieIds
}

func extract_comments( comments_body []interface{} ) map[uint32]string {

  // Get all Comments
  mComments := make( map[uint32]string, 0 )

  for _, t := range comments_body {
    m := t.(map[string]interface{})
    movieId := (uint32)(m["MovieId"].(float64))

    str, ok := m["Comment"]
    if !ok {
      continue
    }

    comment := str.(string)
    if len(comment) == 0 {
      continue
    }

    mComments[ movieId ] = comment
  }

	return mComments
}

func set_comments( body []interface{}, mComments map[uint32]string ) []interface{} {

  for i, t := range body {
    m := t.(map[string]interface{})
    movieId := (uint32)(m["MovieId"].(float64))
    commnt, ok := mComments[movieId]
    if !ok {
      continue
    }
    m["Comment"] = commnt
    body[i] = m
  }

	return body
}

