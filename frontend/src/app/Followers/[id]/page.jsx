'use client';
import { useEffect, useState } from 'react';
import { FaUserPlus } from "react-icons/fa6";
import { FaUserXmark } from "react-icons/fa6";

export default function FollowButton({ targetUserid }) {
  const [following, setFollowing] = useState(false);


  useEffect(()=> {

      const checkfollow = async () => {
        try {
          const res = await fetch ("/api/isFollowing?targetId=${targetUserid}",{
            credentials : 'include'
          })

          const data = await res.json()

          setFollowing(data.isFollowing)
        }catch (err){
          console.log("error while cheking follow satus:",err);
          
        }
      }
    checkfollow()
  },[targetUserid])

  const handleClick = async () => {
    try {
      const url = following ? "/api/unfollowRequest" : "/api/followRequest";

      const response = await fetch(url, {
        credentials: 'include',
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          following_id: "2"
        })
      });
      
      if (!response.ok) {
       
        console.log("response:",await response.json());
        throw new Error('Request failed');
      } else {
        setFollowing(!following);
      }
    } catch (err) {
      console.error('Error:', err);
    }

  };

  return (
    <button onClick={handleClick}>
      {following ? (
        <>
          <FaUserXmark /> Unfollow
        </>
      ) : (
        <>
          <FaUserPlus /> Follow
        </>
      )}
    </button>
  );
}
