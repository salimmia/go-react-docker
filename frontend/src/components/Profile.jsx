import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";

const Profile = () => {
  const { userId } = useParams();
  const [user, setUser] = useState({ profile: {} });

  useEffect(() => {
    fetchUserData();
  }, [userId]);

  const fetchUserData = async () => {
    try {
      const response = await axios.get(
        `http://localhost:8080/users/profile/${userId}`
      );
      setUser(response.data);
    } catch (error) {
      console.error("Error fetching user information:", error);
    }
  };

  return (
    <div className="container mt-5">
      <h1 className="mb-4">User Profile</h1>
      <div>
        <p>
          <strong>Name:</strong> {user.first_name} {user.last_name}
        </p>
        <p>
          <strong>Email:</strong> {user.email}
        </p>
        <p>
          <strong>Phone Number:</strong>{" "}
          {user.profile && user.profile.phone_number
            ? user.profile.phone_number
            : "N/A"}
        </p>
        <p>
          <strong>Birth Date:</strong>{" "}
          {user.profile && user.profile.birth_date
            ? user.profile.birth_date.substring(0, 10)
            : "N/A"}
        </p>
        <p>
          <strong>User Name:</strong> {user.thumbnail}
        </p>
      </div>
    </div>
  );
};

export default Profile;
