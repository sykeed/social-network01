'use client';
import { useState } from 'react';
import { FaUserPlus } from "react-icons/fa6";
import { FaUserXmark } from "react-icons/fa6";

export default function FollowButton({ targetUserid }) {
  const [following, setFollowing] = useState(false);
  const [loading, setLoading] = useState(false);

  const handleClick = async () => {
    setLoading(true);
    try {
      const url = following ? "/api/unfollowRequest" : "/api/followRequest";

      const response = await fetch(url, {
        credentials: 'include',
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          following_id: "1"
        })
      });

      if (!response.ok) {
        throw new Error('Request failed');
      } else {
        setFollowing(!following);
      }
    } catch (err) {
      console.error('Error:', err);
    }
    setLoading(false);
  };

  return (
    <button onClick={handleClick} disabled={loading}>
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
