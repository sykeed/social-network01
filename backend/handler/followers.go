package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	db "social-network/db/cration"
	"social-network/utils"
)

func Followreq(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.JsonResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	
	token, _ := r.Cookie("SessionToken")
	userid := db.GetId("sessionToken", token.Value)
	
	type followReqwest struct {
		FollowingId string `json:"following_id"`
	}
	
	var targetID followReqwest
	
	if err := json.NewDecoder(r.Body).Decode(&targetID); err != nil {
		utils.JsonResponse(w, http.StatusBadRequest, "invalid Profile id")
		fmt.Println("wslat follow", err)
		return
		
	}
	
	
	idInt, _ := strconv.Atoi(targetID.FollowingId)
	
	if targetID.FollowingId == "" || idInt == userid {
		fmt.Println("hnaaa")
		utils.JsonResponse(w, http.StatusBadRequest, "Invalid followeing ID")
		return
	}

	public, err := db.CheckPublic(idInt)
	if err != nil {
		log.Println("error when checking profile status", err)
	}
	
	// check if already on my FOLLOWING list
	exist := db.BeforInsertion(userid, idInt)
	if exist {
		utils.JsonResponse(w, http.StatusBadRequest, "already on my following list")
		log.Println("already on my following list")
		return
	}
	
	if public {
		db.InsertFOllow(userid, idInt, "accepted")
		fmt.Println( exist)
		log.Println("follow succesfullyy")
	} else {
		db.InsertFOllow(userid, idInt, "pending")
		log.Println("follow succesfully")
	}
}


func Unfollowreq(w http.ResponseWriter, r *http.Request) {
	fmt.Println("unfoloooow ")
	if r.Method != http.MethodPost {
		utils.JsonResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	token, _ := r.Cookie("SessionToken")
	userid := db.GetId("sessionToken", token.Value)

	type unfollowRequest struct {
		FollowingId string `json:"following_id"`
	}

	var targetID unfollowRequest

	if err := json.NewDecoder(r.Body).Decode(&targetID); err != nil {
		utils.JsonResponse(w, http.StatusBadRequest, "Invalid Profile ID")
		return
	}

	idInt, _ := strconv.Atoi(targetID.FollowingId)

	if targetID.FollowingId == "" || idInt == userid {
		utils.JsonResponse(w, http.StatusBadRequest, "Invalid following ID")
		return
	}

	 
	exist := db.BeforInsertion(userid, idInt)
	if !exist {
		utils.JsonResponse(w, http.StatusBadRequest, "User is not in your following list")
		log.Println("User is not in your following list")
		return
	}

	 
	db.DeleteFollow(userid, idInt)
	log.Println("Unfollow successfully")
	utils.JsonResponse(w, http.StatusOK, "Unfollow successfully")
}
