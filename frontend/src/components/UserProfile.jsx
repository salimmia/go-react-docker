// UserProfile.js
import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";
import Profile from "./Profile";
import UpdateUser from "./UpdateUser";

const UserProfile = () => {
  const { userId } = useParams();
  const [userData, setUserData] = useState({});
  const [editMode, setEditMode] = useState(false);

  useEffect(() => {
    fetchUserData();
  }, []);

  const fetchUserData = async () => {
    try {
      const response = await axios.get(
        `http://localhost:8080/users/profile/${userId}`
      );
      setUserData(response.data);
    } catch (error) {
      console.error("Error fetching user information:", error);
    }
  };

  const toggleEditMode = () => {
    setEditMode(!editMode);
  };

  return (
    <div className="container mt-5">
      {editMode ? (
        <UpdateUser
          userId={userId}
          fetchUserData={fetchUserData}
          toggleEditMode={toggleEditMode}
        />
      ) : (
        <Profile userData={userData} />
      )}

      <div className="mt-3">
        <button className="btn btn-primary" onClick={toggleEditMode}>
          {editMode ? "Cancel Edit" : "Edit Profile"}
        </button>
      </div>
    </div>
  );
};

export default UserProfile;
