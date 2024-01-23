import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";

const UpdateUser = () => {
  const { userId } = useParams();
  const [formData, setFormData] = useState({
    first_name: "",
    last_name: "",
    thumbnail: "",
    phone_number: "",
    birth_date: "",
  });

  useEffect(() => {
    fetchUserData();
  }, []);

  const fetchUserData = async () => {
    try {
      const response = await axios.get(
        `http://localhost:8080/users/profile/${userId}`
      );
      const user = response.data;

      setFormData({
        first_name: user.first_name,
        last_name: user.last_name,
        thumbnail: user.thumbnail || "",
        phone_number: user.profile?.phone_number || "",
        birth_date: user.profile?.birth_date?.substring(0, 10) || "",
      });
    } catch (error) {
      console.error("Error fetching user information:", error);
    }
  };

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      await axios.put(
        `http://localhost:8080/users/profile/update-user/${userId}`,
        formData
      );
      console.log("User updated successfully");
      alert("User updated successfully");
      // You may want to redirect or show a success message here
    } catch (error) {
      console.error("Error updating user:", error);
      // Handle error, e.g., show an error message
    }
  };

  return (
    <div className="container mt-5">
      <h1 className="mb-4">Update User</h1>
      <form onSubmit={handleSubmit}>
        <div className="mb-3">
          <label htmlFor="firstName" className="form-label">
            First Name
          </label>
          <input
            type="text"
            className="form-control"
            id="firstName"
            name="first_name"
            value={formData.first_name}
            onChange={handleInputChange}
          />
        </div>
        <div className="mb-3">
          <label htmlFor="lastName" className="form-label">
            Last Name
          </label>
          <input
            type="text"
            className="form-control"
            id="lastName"
            name="last_name"
            value={formData.last_name}
            onChange={handleInputChange}
          />
        </div>
        <div className="mb-3">
          <label htmlFor="thumbnail" className="form-label">
            User Name
          </label>
          <input
            type="text"
            className="form-control"
            id="thumbnail"
            name="thumbnail"
            value={formData.thumbnail}
            onChange={handleInputChange}
          />
        </div>
        <div className="mb-3">
          <label htmlFor="phoneNumber" className="form-label">
            Phone Number
          </label>
          <input
            type="text"
            className="form-control"
            id="phoneNumber"
            name="phone_number"
            value={formData.phone_number}
            onChange={handleInputChange}
          />
        </div>
        <div className="mb-3">
          <label htmlFor="birthDate" className="form-label">
            Birth Date
          </label>
          <input
            type="date"
            className="form-control"
            id="birthDate"
            name="birth_date"
            value={formData.birth_date}
            onChange={handleInputChange}
          />
        </div>
        <button type="submit" className="btn btn-primary">
          Update User
        </button>
      </form>
    </div>
  );
};

export default UpdateUser;
